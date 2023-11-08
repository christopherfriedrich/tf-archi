/*
Copyright 2023 The tf-archi Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package transformer provides the API for using and implementing concrete transformations for different providers.
package transformer

import (
	"fmt"

	"github.com/christopherfriedrich/tf-archi/internal/extraction"
	"github.com/christopherfriedrich/tf-archi/internal/util"
	"github.com/christopherfriedrich/tf-archi/pkg/archimate"
)

// TransformationManager is an interface for transforming resource types of one specific IaC provider into an internal representation.
type TransformationManager interface {
	// Provider returns the name of the IaC provider of this TransformationManager.
	Provider() ProviderName
	// TransformForType transforms a given resource type from a IaC provider into an internal representation.
	// If the resource type is not supported, an error is returned.
	TransformForType(resourceType string) (Transformer, error)
	// RegisterTransformer registers a new transformer for a specific resource type.
	RegisterTransformer(Transformer)

	Init()
}

type Transformer interface {
	Type() string
	Transform(*extraction.InfrastructureComponent, *extraction.InfrastructureState) ([]*archimate.Element, []*archimate.Relationship)
}

// ProviderName is a type alias for the name of a terraform provider.
//
// See the official terraform provider registry for a list of providers:
// https://registry.terraform.io/browse/providers
type ProviderName string

// Transformers is a map of all registered TransformationManagers.
var Transformers = make(map[ProviderName]TransformationManager)

func TransformState(infrastructureState *extraction.InfrastructureState, archimateModelName string) (*archimate.Model, []error) {
	// create a slice for all errors
	errors := make([]error, 0)

	// create an empty model named after the input file for the infrastructure state
	archimateModel := archimate.NewModel(infrastructureState.FileName)
	// first extract all the infrastructure components, transform them into their
	// ArchiMate representation and add them to the model.
	for _, component := range infrastructureState.Components {
		// check if Transformers has an entry for component.Type
		if Transformers[ProviderName(component.Provider)] != nil {
			fmt.Printf("Transforming component %s%s%s\n", util.ColorYellow, component.Name, util.ColorNone)
			transformer, err := Transformers[ProviderName(component.Provider)].TransformForType(component.Type)
			if err != nil {
				errors = append(errors, err)
			} else {
				elements, relationships := transformer.Transform(component, infrastructureState)
				if len(elements) > 0 {
					// add all elements to the model
					for _, element := range elements {
						archimateModel.AddElement(element)
					}
				}
				if len(relationships) > 0 {
					// add all relationships to the model
					for _, relationship := range relationships {
						archimateModel.AddRelationship(relationship)
					}
				}
			}

		}
	}
	return archimateModel, errors
}
