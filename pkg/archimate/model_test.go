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

package archimate_test

import (
	"reflect"
	"testing"

	"github.com/christopherfriedrich/tf-archi/pkg/archimate"
)

func TestNewModel(t *testing.T) {
	identifier := "test-identifier"
	model := archimate.NewModel(identifier)

	if model.Identifier != identifier {
		t.Errorf("Expected Identifier to be '%s', but got %s", identifier, model.Identifier)
	}

	if model.NamedReferenceableType.ReferenceableType.IdentifierGroup.IdentifierAttr != identifier {
		t.Errorf("Expected IdentifierGroup.IdentifierAttr to be '%s', but got %s", identifier, model.NamedReferenceableType.ReferenceableType.IdentifierGroup.IdentifierAttr)
	}

	if model.NamedReferenceableType.NameGroup.Name.Value != identifier {
		t.Errorf("Expected NameGroup.Name.Value to be '%s', but got %s", identifier, model.NamedReferenceableType.NameGroup.Name.Value)
	}

	if len(model.Elements) != 0 {
		t.Errorf("Expected Elements to be empty, but got %d elements", len(model.Elements))
	}

	if len(model.Relationships) != 0 {
		t.Errorf("Expected Relationships to be empty, but got %d relationships", len(model.Relationships))
	}
}

func TestAddElement(t *testing.T) {
	identifier := "test-identifier"
	model := archimate.NewModel(identifier)

	if len(model.Elements) != 0 {
		t.Errorf("Expected Elements to be empty, but got %d elements", len(model.Elements))
	}
	elemID := "test-id"
	elemName := "test-name"
	elemLayerElement := archimate.LayerElement("test-element")
	model.AddElement(archimate.NewElement(elemID, elemName, elemLayerElement))

	if len(model.Elements) != 1 {
		t.Errorf("Expected Elements to have 1 element, but got %d elements", len(model.Elements))
	}

	expectedElement := &archimate.Element{
		NamedReferenceableType: &archimate.NamedReferenceableType{
			ReferenceableType: &archimate.ReferenceableType{
				IdentifierGroup: &archimate.IdentifierGroup{
					IdentifierAttr: elemID,
				},
			},
			NameGroup: &archimate.NameGroup{
				Name: &archimate.LangStringType{
					XMLLangAttr: "en",
					Value:       elemName,
				},
			},
		},
		ElementType: &archimate.ElementType{
			Type: elemLayerElement,
		},
	}
	actualElement := model.Elements[0]

	if !reflect.DeepEqual(actualElement, expectedElement) {
		t.Errorf("NewElement() = %v, expected %v", actualElement, expectedElement)
	}
}
