package storage

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/klog/v2"

	virtv1 "kubevirt.io/api/core/v1"
	"kubevirt.io/client-go/kubecli"

	"kubevirt.io/virt-template/api/subresourcesv1alpha1"
	"kubevirt.io/virt-template/client-go/template"
)

// +kubebuilder:rbac:groups=template.kubevirt.io,resources=virtualmachinetemplates,verbs=get
// +kubebuilder:rbac:groups=template.kubevirt.io,resources=virtualmachinetemplates/status,verbs=get

type ProcessedVirtualMachineTemplateREST struct {
	virtClient     kubecli.KubevirtClient
	templateClient template.Interface
}

func NewVirtualMachineTemplateProcessREST(virtClient kubecli.KubevirtClient, templateClient template.Interface) *ProcessedVirtualMachineTemplateREST {
	return &ProcessedVirtualMachineTemplateREST{
		virtClient:     virtClient,
		templateClient: templateClient,
	}
}

var (
	_ = rest.Storage(&ProcessedVirtualMachineTemplateREST{})
	_ = rest.GetterWithOptions(&ProcessedVirtualMachineTemplateREST{})
)

func (r *ProcessedVirtualMachineTemplateREST) New() runtime.Object {
	return &subresourcesv1alpha1.ProcessedVirtualMachineTemplate{}
}

func (r *ProcessedVirtualMachineTemplateREST) Destroy() {}

func (r *ProcessedVirtualMachineTemplateREST) Get(ctx context.Context, name string, options runtime.Object) (runtime.Object, error) {
	ns, ok := request.NamespaceFrom(ctx)
	if !ok {
		return nil, fmt.Errorf("missing namespace")
	}

	klog.Infof("GET /process for VirtualMachineTemplate %s/%s", ns, name)

	processOptions, ok := options.(*subresourcesv1alpha1.ProcessOptions)
	if !ok {
		return nil, fmt.Errorf("invalid options")
	}

	return &subresourcesv1alpha1.ProcessedVirtualMachineTemplate{
		VirtualMachine: &virtv1.VirtualMachine{
			ObjectMeta: metav1.ObjectMeta{
				Name: processOptions.Foo,
			},
		},
	}, nil
}

func (r *ProcessedVirtualMachineTemplateREST) NewGetOptions() (options runtime.Object, include bool, path string) {
	return &subresourcesv1alpha1.ProcessOptions{}, false, ""
}
