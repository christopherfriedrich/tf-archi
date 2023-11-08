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

import "github.com/google/uuid"

// Relationship is a base relationship type that can be extended by concrete ArchiMate types.
//
// Note that Relationship is abstract, so one must have derived types of this type. This is indicated in XML
// by having a tag name of "relationship" and an attribute of xsi:type="AccessRelationship" where AccessRelationship is
// a derived type from RelationshipType.
type Relationship struct {
	*IdentifierGroup
	*RelationshipGroup
	*ElementType
}

type RelationshipGroup struct {
	SourceAttr string `xml:"source,attr"`
	TargetAttr string `xml:"target,attr"`
}

type RelationshipType LayerElement

const RelationshipTypeComposition RelationshipType = "Composition"
const RelationshipTypeAssignment RelationshipType = "Assignment"
const RelationshipTypeRealization RelationshipType = "Realization"

func NewRelationship(source, target string, relationship LayerElement) *Relationship {
	return &Relationship{
		IdentifierGroup: &IdentifierGroup{
			// we do not really care about the ID, so we just generate a new one, but
			// since identifiers in the Archimate schema always need to start with a
			// letter, we prefix the generated UUID with a "r"
			IdentifierAttr: "r" + uuid.New().String(),
		},
		RelationshipGroup: &RelationshipGroup{
			SourceAttr: source,
			TargetAttr: target,
		},
		ElementType: &ElementType{
			Type: relationship,
		},
	}
}
