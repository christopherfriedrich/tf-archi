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

package archimate

// Element is the base element type that can be extended by concrete ArchiMate types.
//
// Note that Element is abstract, so one must have derived types of this type. this is indicated in xml
// by having a tag name of "element" and an attribute of xsi:type="BusinessRole" where BusinessRole is
// a derived type from ElementType.
type Element struct {
	*NamedReferenceableType
	*ElementType
}

// ElementType is the type of the archimate element.
type ElementType struct {
	Type LayerElement `xml:"xsi:type,attr"`
}

// NewElement creates a new element with the given id, name and element type.
func NewElement(id, name string, element LayerElement) *Element {
	return &Element{
		NamedReferenceableType: &NamedReferenceableType{
			ReferenceableType: &ReferenceableType{
				IdentifierGroup: &IdentifierGroup{
					IdentifierAttr: id,
				},
			},
			NameGroup: &NameGroup{
				Name: &LangStringType{
					XMLLangAttr: "en",
					Value:       name,
				},
			},
		},
		ElementType: &ElementType{
			Type: element,
		},
	}
}
