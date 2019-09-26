// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.ibm.com/jdunham/csi-scale-operator/pkg/apis/scale/v1alpha1.CSIClusterSpec":         schema_pkg_apis_scale_v1alpha1_CSIClusterSpec(ref),
		"github.ibm.com/jdunham/csi-scale-operator/pkg/apis/scale/v1alpha1.CSIPrimarySpec":         schema_pkg_apis_scale_v1alpha1_CSIPrimarySpec(ref),
		"github.ibm.com/jdunham/csi-scale-operator/pkg/apis/scale/v1alpha1.CSIRestApiSpec":         schema_pkg_apis_scale_v1alpha1_CSIRestApiSpec(ref),
		"github.ibm.com/jdunham/csi-scale-operator/pkg/apis/scale/v1alpha1.CSIScaleOperator":       schema_pkg_apis_scale_v1alpha1_CSIScaleOperator(ref),
		"github.ibm.com/jdunham/csi-scale-operator/pkg/apis/scale/v1alpha1.CSIScaleOperatorSpec":   schema_pkg_apis_scale_v1alpha1_CSIScaleOperatorSpec(ref),
		"github.ibm.com/jdunham/csi-scale-operator/pkg/apis/scale/v1alpha1.CSIScaleOperatorStatus": schema_pkg_apis_scale_v1alpha1_CSIScaleOperatorStatus(ref),
	}
}

func schema_pkg_apis_scale_v1alpha1_CSIClusterSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "\n CSIClusterSpec  defines the desired state of CSIi Scale Cluster",
				Properties: map[string]spec.Schema{
					"id": {
						SchemaProps: spec.SchemaProps{
							Description: "The cluster id of the gpfs cluster specified (mandatory).",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"secureSslMode": {
						SchemaProps: spec.SchemaProps{
							Description: "Require a secure SSL connection to connect to GPFS.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"secrets": {
						SchemaProps: spec.SchemaProps{
							Description: "A string specifying a secret resource name.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"cacert": {
						SchemaProps: spec.SchemaProps{
							Description: "A string specifying a cacert resource name.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"primary": {
						SchemaProps: spec.SchemaProps{
							Description: "The primary file system for the GPFS cluster.",
							Ref:         ref("github.ibm.com/jdunham/csi-scale-operator/pkg/apis/scale/v1alpha1.CSIPrimarySpec"),
						},
					},
					"restApi": {
						SchemaProps: spec.SchemaProps{
							Description: "A collection of targets for REST calls.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.ibm.com/jdunham/csi-scale-operator/pkg/apis/scale/v1alpha1.CSIRestApiSpec"),
									},
								},
							},
						},
					},
				},
				Required: []string{"id"},
			},
		},
		Dependencies: []string{
			"github.ibm.com/jdunham/csi-scale-operator/pkg/apis/scale/v1alpha1.CSIPrimarySpec", "github.ibm.com/jdunham/csi-scale-operator/pkg/apis/scale/v1alpha1.CSIRestApiSpec"},
	}
}

func schema_pkg_apis_scale_v1alpha1_CSIPrimarySpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Defines the primary filesystem.",
				Properties: map[string]spec.Schema{
					"primaryFS": {
						SchemaProps: spec.SchemaProps{
							Description: "The name of the primary filesystem.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"primaryFset": {
						SchemaProps: spec.SchemaProps{
							Description: "The name of the primary fileset, created in primaryFS.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_scale_v1alpha1_CSIRestApiSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Defines the desired REST API access info.",
				Properties: map[string]spec.Schema{
					"guiHost": {
						SchemaProps: spec.SchemaProps{
							Description: "The hostname of the REST server.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"guiPort": {
						SchemaProps: spec.SchemaProps{
							Description: "The port number running the REST server.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_scale_v1alpha1_CSIScaleOperator(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CSIScaleOperator is the Schema for the csiscaleoperators API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.ibm.com/jdunham/csi-scale-operator/pkg/apis/scale/v1alpha1.CSIScaleOperatorSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.ibm.com/jdunham/csi-scale-operator/pkg/apis/scale/v1alpha1.CSIScaleOperatorStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.ibm.com/jdunham/csi-scale-operator/pkg/apis/scale/v1alpha1.CSIScaleOperatorSpec", "github.ibm.com/jdunham/csi-scale-operator/pkg/apis/scale/v1alpha1.CSIScaleOperatorStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_scale_v1alpha1_CSIScaleOperatorSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CSIScaleOperatorSpec defines the desired state of CSIScaleOperator",
				Properties: map[string]spec.Schema{
					"attacher": {
						SchemaProps: spec.SchemaProps{
							Description: "Attacher image for csi (actually attaches to the storage).",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"provisioner": {
						SchemaProps: spec.SchemaProps{
							Description: "Provisioner image for csi (actually issues provision requests).",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"driverRegistrar": {
						SchemaProps: spec.SchemaProps{
							Description: "Sidecar container image for the csi spectrum scale plugin pods.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"spectrumScale": {
						SchemaProps: spec.SchemaProps{
							Description: "Image name for the csi spectrum scale plugin container.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"scaleHostpath": {
						SchemaProps: spec.SchemaProps{
							Description: "The path to the gpfs file system mounted on the host machine.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"clusters": {
						SchemaProps: spec.SchemaProps{
							Description: "A collection of gpfs cluster properties for the csi driver to mount.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.ibm.com/jdunham/csi-scale-operator/pkg/apis/scale/v1alpha1.CSIClusterSpec"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.ibm.com/jdunham/csi-scale-operator/pkg/apis/scale/v1alpha1.CSIClusterSpec"},
	}
}

func schema_pkg_apis_scale_v1alpha1_CSIScaleOperatorStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CSIScaleOperatorStatus defines the observed state of CSIScaleOperator",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}
