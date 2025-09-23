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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// VirtualMachineTemplateSpec defines the desired state of VirtualMachineTemplate
type VirtualMachineTemplateSpec struct {
	// message is an optional instructional message for this template.
	// This field should inform the user how to utilize the newly created resource.
	Message string `json:"message,omitempty,omitzero" protobuf:"bytes,1,opt,name=message"`

	// virtualMachine is the KubeVirt VirtualMachine to include in this template.
	// If a namespace value is hardcoded in the object, it will be removed
	// during template instantiation, however if the namespace value
	// is, or contains, a ${PARAMETER_REFERENCE}, the resolved
	// value after parameter substitution will be respected and the object
	// will be created in that namespace.
	// +kubebuilder:pruning:PreserveUnknownFields
	VirtualMachine runtime.RawExtension `json:"virtualMachine" protobuf:"bytes,2,name=virtualMachine"`

	// parameters is an optional array of Parameters used during the
	// Template to Config transformation.
	Parameters []Parameter `json:"parameters,omitempty" protobuf:"bytes,3,rep,name=parameters"`
}

// Parameter defines a name/value variable that is to be processed during
// the Template to Config transformation.
type Parameter struct {
	// name must be set and it can be referenced in Template
	// Items using ${PARAMETER_NAME}. Required.
	Name string `json:"name" protobuf:"bytes,1,opt,name=name"`

	// Optional: The name that will show in UI instead of parameter 'Name'
	DisplayName string `json:"displayName,omitempty,omitzero" protobuf:"bytes,2,opt,name=displayName"`

	// description of a parameter. Optional.
	Description string `json:"description,omitempty,omitzero" protobuf:"bytes,3,opt,name=description"`

	// value holds the Parameter data. If specified, the generator will be
	// ignored. The value replaces all occurrences of the Parameter ${Name}
	// expression during the Template to Config transformation. Optional.
	Value string `json:"value,omitempty,omitzero" protobuf:"bytes,4,opt,name=value"`

	// generate specifies the generator to be used to generate random string
	// from an input value specified by From field. The result string is
	// stored into Value field. If empty, no generator is being used, leaving
	// the result Value untouched. Optional.
	//
	// The only supported generator is "expression", which accepts a "from"
	// value in the form of a simple regular expression containing the
	// range expression "[a-zA-Z0-9]", and the length expression "a{length}".
	//
	// Examples:
	//
	// from             | value
	// -----------------------------
	// "test[0-9]{1}x"  | "test7x"
	// "[0-1]{8}"       | "01001100"
	// "0x[A-F0-9]{4}"  | "0xB3AF"
	// "[a-zA-Z0-9]{8}" | "hW4yQU5i"
	//
	Generate string `json:"generate,omitempty,omitzero" protobuf:"bytes,5,opt,name=generate"`

	// from is an input value for the generator. Optional.
	From string `json:"from,omitempty,omitzero" protobuf:"bytes,6,opt,name=from"`

	// Optional: Indicates the parameter must have a value.  Defaults to false.
	Required bool `json:"required,omitempty" protobuf:"varint,7,opt,name=required"`
}

// VirtualMachineTemplateStatus defines the observed state of VirtualMachineTemplate.
type VirtualMachineTemplateStatus struct {
	// conditions represent the current state of the VirtualMachineTemplate resource.
	// Each condition has a unique type and reflects the status of a specific aspect of the resource.
	//
	// Condition types include:
	// - "Ready": the resource is fully functional
	//
	// The status of each condition is one of True, False, or Unknown.
	// +listType=map
	// +listMapKey=type
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty" protobuf:"bytes,1,rep,name=conditions"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +genclient

// VirtualMachineTemplate is the Schema for the virtualmachinetemplates API
type VirtualMachineTemplate struct {
	metav1.TypeMeta `json:",inline"`

	// metadata is a standard object metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty,omitzero" protobuf:"bytes,1,opt,name=metadata"`

	// spec defines the desired state of VirtualMachineTemplate
	// +required
	Spec VirtualMachineTemplateSpec `json:"spec" protobuf:"bytes,2,name=spec"`

	// status defines the observed state of VirtualMachineTemplate
	// +optional
	Status VirtualMachineTemplateStatus `json:"status,omitempty,omitzero" protobuf:"bytes,3,opt,name=status"`
}

// +kubebuilder:object:root=true

// VirtualMachineTemplateList contains a list of VirtualMachineTemplate
type VirtualMachineTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty,omitzero" protobuf:"bytes,1,opt,name=metadata"`
	Items           []VirtualMachineTemplate `json:"items" protobuf:"bytes,2,rep,name=items"`
}

func init() {
	SchemeBuilder.Register(&VirtualMachineTemplate{}, &VirtualMachineTemplateList{})
}
