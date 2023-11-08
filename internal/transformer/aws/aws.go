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

package aws

import (
	"fmt"

	"github.com/christopherfriedrich/tf-archi/internal/transformer"
)

const provider = "registry.terraform.io/hashicorp/aws"

// TransformationManager is a registry of transformers for Azure resources
type TransformationManager struct {
	transformer map[string]transformer.Transformer
}

func NewAWSTransformationManager() *TransformationManager {
	tm := &TransformationManager{
		transformer: make(map[string]transformer.Transformer),
	}
	tm.Init()
	return tm
}

func (t *TransformationManager) Provider() transformer.ProviderName {
	return provider
}

func (t *TransformationManager) TransformForType(resourceType string) (transformer.Transformer, error) {
	if t.transformer[resourceType] != nil {
		return t.transformer[resourceType], nil
	}
	return nil, fmt.Errorf("resource type %s not supported", resourceType)
}

func (t *TransformationManager) RegisterTransformer(transformer transformer.Transformer) {
	t.transformer[transformer.Type()] = transformer
}

func (t *TransformationManager) Init() {
	// TODO: register transformers, e. g.: t.RegisterTransformer(NewEC2Transformer())
}
