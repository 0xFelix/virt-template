/*
Copyright 2025.

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

package subresourcesv1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	virtv1 "kubevirt.io/api/core/v1"
)

// +kubebuilder:object:root=true

// VirtualMachineTemplate is a dummy object to satisfy the k8s.io/apiserver conventions.
// A subresource cannot be served without a storage for its parent resource.
type VirtualMachineTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty,omitzero" protobuf:"bytes,1,opt,name=metadata"`
}

// +kubebuilder:object:root=true

// ProcessedVirtualMachineTemplate is the object served by the /process and /create subresources.
// It's not a standalone resource but represents a process or create action on the parent VirtualMachineTemplate resource.
type ProcessedVirtualMachineTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty,omitzero" protobuf:"bytes,1,opt,name=metadata"`

	// message is an optional instructional message from the processed template.
	// This field should inform the user how to utilize the newly created resource.
	Message string `json:"message,omitempty,omitzero" protobuf:"bytes,2,opt,name=message"`

	VirtualMachine virtv1.VirtualMachine `json:"virtualMachine" protobuf:"bytes,3,name=virtualMachine"`
}

// +kubebuilder:object:root=true

// ProcessOptions are the options used when processing a VirtualMachineTemplate.
type ProcessOptions struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty,omitzero" protobuf:"bytes,1,opt,name=metadata"`

	// parameters is an optional array of Parameters used during the
	// Template to Config transformation.
	Parameters []Parameter `json:"parameters,omitempty" protobuf:"bytes,2,rep,name=parameters"`
}

// Parameter defines a name/value variable that is to be processed during
// the Template to Config transformation.
type Parameter struct {
	// name must be set and it can be referenced in Template
	// Items using ${PARAMETER_NAME}. Required.
	Name string `json:"name" protobuf:"bytes,1,name=name"`

	// value holds the Parameter data. If specified, the generator will be
	// ignored. The value replaces all occurrences of the Parameter ${Name}
	// expression during the Template to Config transformation. Required.
	Value string `json:"value" protobuf:"bytes,2,name=value"`
}

func init() {
	SchemeBuilder.Register(&VirtualMachineTemplate{}, &ProcessOptions{}, &ProcessedVirtualMachineTemplate{})
}
