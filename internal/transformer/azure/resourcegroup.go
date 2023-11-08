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

package azure

import (
	"github.com/christopherfriedrich/tf-archi/internal/extraction"
	"github.com/christopherfriedrich/tf-archi/pkg/archimate"
)

type ResourceGroupTransformer struct{}

func NewResourceGroupTransformer() *ResourceGroupTransformer {
	return &ResourceGroupTransformer{}
}

func (t ResourceGroupTransformer) Type() string {
	return "azurerm_resource_group"
}

func (t ResourceGroupTransformer) Transform(component *extraction.InfrastructureComponent, state *extraction.InfrastructureState) ([]*archimate.Element, []*archimate.Relationship) {
	return []*archimate.Element{
		archimate.NewElement(component.ID, component.Name, archimate.Grouping),
	}, []*archimate.Relationship{}
}
