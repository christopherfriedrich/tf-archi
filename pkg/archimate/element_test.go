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

func TestNewElement(t *testing.T) {
	id := "test-id"
	name := "test-name"
	element := archimate.LayerElement("test-element")

	expected := &archimate.Element{
		NamedReferenceableType: &archimate.NamedReferenceableType{
			ReferenceableType: &archimate.ReferenceableType{
				IdentifierGroup: &archimate.IdentifierGroup{
					IdentifierAttr: id,
				},
			},
			NameGroup: &archimate.NameGroup{
				Name: &archimate.LangStringType{
					XMLLangAttr: "en",
					Value:       name,
				},
			},
		},
		ElementType: &archimate.ElementType{
			Type: element,
		},
	}

	actual := archimate.NewElement(id, name, element)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("NewElement() = %v, expected %v", actual, expected)
	}
}
