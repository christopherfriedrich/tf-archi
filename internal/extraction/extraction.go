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

package extraction

import "errors"

// InfrastructureState represents the state of the infrastructure as extracted by the extraction function
type InfrastructureState struct {
	// name of the file from which the infrastructure state is extracted
	FileName string
	// stores the original state without any modifications
	OriginalState interface{} `json:"originalState"`
	// stores all the components extracted from the state
	Components []*InfrastructureComponent `json:"components"`
}

// GetComponentsForType returns all components of a given type
func (is *InfrastructureState) GetTypeForComponent(componentAddress string) (string, error) {
	for _, component := range is.Components {
		if component.Address == componentAddress {
			return component.Type, nil
		}
	}
	return "", errors.New("component not found")
}

// GetComponentByAddress returns a component by its address
func (is *InfrastructureState) GetComponentByAddress(componentAddress string) (*InfrastructureComponent, error) {
	for _, component := range is.Components {
		if component.Address == componentAddress {
			return component, nil
		}
	}
	return nil, errors.New("component not found")
}

// InfrastructureComponent is a representation of an infrastructure component
type InfrastructureComponent struct {
	ID           string                               `json:"id"`
	Name         string                               `json:"name"`
	Address      string                               `json:"address"`
	Provider     string                               `json:"provider"`
	Type         string                               `json:"type"`
	Attributes   interface{}                          `json:"attributes"`
	Dependencies []*InfrastructureComponentDependency `json:"dependencies"`
}

func (ic *InfrastructureComponent) GetDependentsForType(componentType string) []*InfrastructureComponentDependency {
	dependents := make([]*InfrastructureComponentDependency, 0)
	for _, dependency := range ic.Dependencies {
		if dependency.Type == componentType {
			dependents = append(dependents, dependency)
		}
	}
	return dependents
}

type InfrastructureComponentDependency struct {
	Address string `json:"address"`
	Type    string `json:"type"`
}
