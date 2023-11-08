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

import (
	"encoding/xml"

	"github.com/google/uuid"
)

// Model is the root model type for an ArchiMate Enterprise Architecture Model.
// It is a container for the elements, relationships, diagrams and organizations of the model.
//
// See the official XSD for more information: https://www.opengroup.org/xsd/archimate/3.1/archimate3_Model.xsd
type Model struct {
	XMLName           xml.Name `xml:"model"`
	Xmlns             string   `xml:"xmlns,attr"`
	XmlnsXsi          string   `xml:"xmlns:xsi,attr"`
	XsiSchemaLocation string   `xml:"xsi:schemaLocation,attr"`
	Identifier        string   `xml:"identifier,attr"`
	*NamedReferenceableType
	Elements      []*Element      `xml:"elements>element"`
	Relationships []*Relationship `xml:"relationships>relationship,omitempty"`
}

func NewModel(name string) *Model {
	identifier := "m" + uuid.New().String()
	return &Model{
		Xmlns:             "http://www.opengroup.org/xsd/archimate/3.0/",
		XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
		XsiSchemaLocation: "http://www.opengroup.org/xsd/archimate/3.0/ http://www.opengroup.org/xsd/archimate/3.1/archimate3_Diagram.xsd",
		Identifier:        identifier,
		NamedReferenceableType: &NamedReferenceableType{
			ReferenceableType: &ReferenceableType{
				IdentifierGroup: &IdentifierGroup{
					IdentifierAttr: identifier,
				},
			},
			NameGroup: &NameGroup{
				Name: &LangStringType{
					XMLLangAttr: "en",
					Value:       name,
				},
			},
		},
		Elements:      make([]*Element, 0),
		Relationships: make([]*Relationship, 0),
	}
}

// Adds an element to the model.
func (m *Model) AddElement(element *Element) {
	m.Elements = append(m.Elements, element)
}

func (m *Model) AddRelationship(relationship *Relationship) {
	m.Relationships = append(m.Relationships, relationship)
}
