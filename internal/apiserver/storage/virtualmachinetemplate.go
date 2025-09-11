package storage

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/rest"

	templateapi "kubevirt.io/virt-template/api"
	"kubevirt.io/virt-template/api/subresourcesv1alpha1"
)

// VirtualMachineTemplateDummyREST is required to satisfy the k8s.io/apiserver conventions.
// A subresource cannot be served without a storage for its parent resource.
type VirtualMachineTemplateDummyREST struct{}

func NewVirtualMachineTemplateDummyREST() *VirtualMachineTemplateDummyREST {
	return &VirtualMachineTemplateDummyREST{}
}

var (
	_ = rest.Storage(&VirtualMachineTemplateDummyREST{})
	_ = rest.Scoper(&VirtualMachineTemplateDummyREST{})
	_ = rest.SingularNameProvider(&VirtualMachineTemplateDummyREST{})
)

func (r *VirtualMachineTemplateDummyREST) New() runtime.Object {
	return &subresourcesv1alpha1.VirtualMachineTemplate{}
}

func (r *VirtualMachineTemplateDummyREST) Destroy() {}

func (r *VirtualMachineTemplateDummyREST) NamespaceScoped() bool { return true }

func (r *VirtualMachineTemplateDummyREST) GetSingularName() string {
	return templateapi.SingularResourceName
}
