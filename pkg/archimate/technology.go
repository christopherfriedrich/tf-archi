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

// TechnologyLayerElement is a type that represents the elements of the ArchiMate technology layer.
//
// See https://pubs.opengroup.org/architecture/archimate31-doc/chap10.html
type TechnologyLayerElement LayerElement

// TechnologyElementNode represents a computational or physical resource that hosts, manipulates, or interacts with other
// computational or physical resources.
const TechnologyElementNode TechnologyLayerElement = "Node"

// TechnologyElementDevice represents a physical IT resource upon which system software and artifacts may be stored or
// deployed for execution.
const TechnologyElementDevice TechnologyLayerElement = "Device"

// TechnologyElementSystemSoftware represents software that provides or contributes to an environment for storing,
// executing, and using software or data deployed within it.
const TechnologyElementSystemSoftware TechnologyLayerElement = "SystemSoftware"

// TechnologyElementTechnologyCollaboration represents an aggregate of two or more technology internal active structure
// elements that work together to perform collective technology behavior.
const TechnologyElementTechnologyCollaboration TechnologyLayerElement = "TechnologyCollaboration"

// TechnologyElementTechnologyInterface represents a point of access where technology services offered by a node can be
// accessed.
const TechnologyElementTechnologyInterface TechnologyLayerElement = "TechnologyInterface"

// TechnologyElementPath represents a link between two or more nodes, through which these nodes can exchange data,
// energy, or material.
const TechnologyElementPath TechnologyLayerElement = "Path"

// TechnologyElementCommunicationNetwork represents a set of structures that connects nodes for transmission, routing, and
// reception of data.
const TechnologyElementCommunicationNetwork TechnologyLayerElement = "CommunicationNetwork"

// TechnologyElementTechnologyFunction represents a collection of technology behavior that can be performed by a node.
const TechnologyElementTechnologyFunction TechnologyLayerElement = "TechnologyFunction"

// TechnologyElementTechnologyProcess represents a sequence of technology behaviors that achieves a specific result.
const TechnologyElementTechnologyProcess TechnologyLayerElement = "TechnologyProcess"

// TechnologyElementTechnologyInteraction represents a unit of collective technology behavior performed
// by (a collaboration of) two or more nodes.
const TechnologyElementTechnologyInteraction TechnologyLayerElement = "TechnologyInteraction"

// TechnologyElementTechnologyEvent represents a technology state change.
const TechnologyElementTechnologyEvent TechnologyLayerElement = "TechnologyEvent"

// TechnologyElementTechnologyService represents an explicitly defined exposed technology behavior.
const TechnologyElementTechnologyService TechnologyLayerElement = "TechnologyService"

// TechnologyElementArtifact represents a piece of data that is used or produced in a software development process, or by
// deployment and operation of an IT system.
const TechnologyElementArtifact TechnologyLayerElement = "Artifact"
