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

import (
	"encoding/json"
	"errors"
	"os"
	"path"
	"strings"

	terraformjson "github.com/hashicorp/terraform-json"
	"github.com/spf13/cobra"
)

// ExtractFromTerraformState extracts the resources from a Terraform state file.
func ExtractFromTerraformState(file string) (*InfrastructureState, error) {
	data, err := os.ReadFile(file)
	cobra.CheckErr(err)
	// Parse the JSON representation using the terraformjson.StateFromJSON function.
	var state terraformjson.State
	err = json.Unmarshal(data, &state)
	cobra.CheckErr(err)
	_, fileName := path.Split(file)
	return convertToInfrastructureStateHolder(&state, strings.Split(fileName, ".")[0])
}

// Takes a terraformjson.State and converts it into an InfrastructureState representation.
func convertToInfrastructureStateHolder(inputState *terraformjson.State, fileName string) (*InfrastructureState, error) {
	// Create an empty InfrastructureState with the filename and input state
	infrastructureState := InfrastructureState{
		FileName:      fileName,
		OriginalState: inputState,
	}

	for _, resource := range collectResources(inputState.Values.RootModule) {
		infrastructureComponent := InfrastructureComponent{}
		infrastructureComponentDependencies := make([]*InfrastructureComponentDependency, 0)
		if resource.AttributeValues["id"] != nil {
			// The Terraform "id" representation contains "/" symbols, this is not allowed in ArchiMate,
			// so we have to replace them with "-"
			infrastructureComponent.ID = ConvertTerraformResourceIDToGenericID(resource.AttributeValues["id"].(string))
		} else {
			// This should not be the case, if we get this error there is a high chance,
			// that the Terraform state file is corrupted.
			return nil, errors.New("resource does not have an id")
		}
		if resource.AttributeValues["name"] != nil {
			infrastructureComponent.Name = resource.AttributeValues["name"].(string)
		} else {
			infrastructureComponent.Name = resource.Name
		}
		infrastructureComponent.Type = resource.Type
		infrastructureComponent.Provider = resource.ProviderName
		infrastructureComponent.Address = resource.Address
		infrastructureComponent.Attributes = resource.AttributeValues

		for _, dependent := range resource.DependsOn {
			// TODO: fix this
			infrastructureComponentDependencies = append(infrastructureComponentDependencies, &InfrastructureComponentDependency{
				Address: dependent,
			})
		}
		infrastructureComponent.Dependencies = infrastructureComponentDependencies
		infrastructureState.Components = append(infrastructureState.Components, &infrastructureComponent)
	}
	// We need to add the type of the infrastructure dependency. This is done in a second loop to ensure that all components are available.
	// We need the types, because we only want to create relations for certain dependencies, so we have to filter by type.
	for _, component := range infrastructureState.Components {
		for _, componentDependency := range component.Dependencies {
			dependencyType, err := infrastructureState.GetTypeForComponent(componentDependency.Address)
			if err != nil {
				return nil, err
			}
			componentDependency.Type = dependencyType
		}
	}
	return &infrastructureState, nil
}

// Recursively collect all resources from a *terraformjson.StateModule pointer
func collectResources(module *terraformjson.StateModule) []*terraformjson.StateResource {
	resources := make([]*terraformjson.StateResource, 0)
	resources = append(resources, module.Resources...)
	// Recursively collect resources from child modules
	for _, childModule := range module.ChildModules {
		resources = append(resources, collectResources(childModule)...)
	}
	return resources
}

// Takes an id from a terraformjson.StateResource and converts it into a generic id.
func ConvertTerraformResourceIDToGenericID(input string) string {
	// take input and replace all occurences of "/" with "-"
	return strings.ReplaceAll(strings.TrimPrefix(input, "/"), "/", "-")
}
