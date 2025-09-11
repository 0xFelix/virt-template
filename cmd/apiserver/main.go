package main

import (
	"os"

	"github.com/spf13/pflag"
	"k8s.io/klog/v2"

	"kubevirt.io/client-go/kubecli"

	templateapi "kubevirt.io/virt-template/api"
	"kubevirt.io/virt-template/client-go/template"

	"kubevirt.io/virt-template/api/subresourcesv1alpha1"
	"kubevirt.io/virt-template/internal/apiserver"
	"kubevirt.io/virt-template/internal/apiserver/openapi"
	"kubevirt.io/virt-template/internal/apiserver/storage"
	templatescheme "kubevirt.io/virt-template/internal/scheme"
)

func main() {
	s := apiserver.New()

	s.AddFlags(pflag.CommandLine)
	pflag.CommandLine.AddGoFlagSet(kubecli.FlagSet())
	pflag.Parse()

	virtClient, err := kubecli.GetKubevirtClient()
	if err != nil {
		klog.Fatalf("Failed to get virtClient: %v", err)
	}
	templateClient, err := template.NewForConfig(virtClient.Config())
	if err != nil {
		klog.Fatalf("Failed to get templateClient: %v", err)
	}

	scheme := templatescheme.New()
	apiGroups := apiserver.APIGroups{
		subresourcesv1alpha1.GroupVersion: {
			templateapi.PluralResourceName:              storage.NewVirtualMachineTemplateDummyREST(),
			templateapi.PluralResourceName + "/process": storage.NewVirtualMachineTemplateProcessREST(virtClient, templateClient),
		},
	}

	if err := s.Run(
		"virt-template-apiserver",
		scheme, openapi.NewConfig(scheme), openapi.NewV3Config(scheme), apiGroups,
	); err != nil {
		os.Exit(1)
	}
}
