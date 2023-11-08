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

type LangStringType struct {
	XMLLangAttr string `xml:"xml:lang,attr,omitempty"`
	Value       string `xml:",chardata"`
}

// NameGroup ...
type NameGroup struct {
	Name *LangStringType `xml:"name"`
}

type IdentifierGroup struct {
	IdentifierAttr string `xml:"identifier,attr"`
}

// ReferenceableType represents something that can be referenced in the model.
type ReferenceableType struct {
	*IdentifierGroup
}

// NamedReferenceableType represents something that can be referenced in the model and has a name.
type NamedReferenceableType struct {
	*ReferenceableType
	*NameGroup
}
