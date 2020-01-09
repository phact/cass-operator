// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/riptano/dse-operator/operator/pkg/apis/datastax/v1alpha1.DseDatacenter":       schema_pkg_apis_datastax_v1alpha1_DseDatacenter(ref),
		"github.com/riptano/dse-operator/operator/pkg/apis/datastax/v1alpha1.DseDatacenterSpec":   schema_pkg_apis_datastax_v1alpha1_DseDatacenterSpec(ref),
		"github.com/riptano/dse-operator/operator/pkg/apis/datastax/v1alpha1.DseDatacenterStatus": schema_pkg_apis_datastax_v1alpha1_DseDatacenterStatus(ref),
	}
}

func schema_pkg_apis_datastax_v1alpha1_DseDatacenter(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "DseDatacenter is the Schema for the dsedatacenters API",
				Type:        []string{"object"},
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
							Ref: ref("github.com/riptano/dse-operator/operator/pkg/apis/datastax/v1alpha1.DseDatacenterSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/riptano/dse-operator/operator/pkg/apis/datastax/v1alpha1.DseDatacenterStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/riptano/dse-operator/operator/pkg/apis/datastax/v1alpha1.DseDatacenterSpec", "github.com/riptano/dse-operator/operator/pkg/apis/datastax/v1alpha1.DseDatacenterStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_datastax_v1alpha1_DseDatacenterSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "DseDatacenterSpec defines the desired state of DseDatacenter",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"size": {
						SchemaProps: spec.SchemaProps{
							Description: "Desired number of DSE server nodes",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"dseVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "DSE version number",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"dseImage": {
						SchemaProps: spec.SchemaProps{
							Description: "DSE container image name. More info: https://kubernetes.io/docs/concepts/containers/images",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"config": {
						SchemaProps: spec.SchemaProps{
							Description: "Config for DSE, in YAML format",
							Type:        []string{"string"},
							Format:      "byte",
						},
					},
					"managementApiAuth": {
						SchemaProps: spec.SchemaProps{
							Description: "Config for DSE Management API certificates",
							Ref:         ref("github.com/riptano/dse-operator/operator/pkg/apis/datastax/v1alpha1.ManagementApiAuthConfig"),
						},
					},
					"resources": {
						SchemaProps: spec.SchemaProps{
							Description: "Kubernetes resource requests and limits, per DSE pod",
							Ref:         ref("k8s.io/api/core/v1.ResourceRequirements"),
						},
					},
					"racks": {
						SchemaProps: spec.SchemaProps{
							Description: "A list of the named racks in the datacenter, representing independent failure domains. The number of racks should match the replication factor in the keyspaces you plan to create, and the number of racks cannot easily be changed once a datacenter is deployed.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/riptano/dse-operator/operator/pkg/apis/datastax/v1alpha1.DseRack"),
									},
								},
							},
						},
					},
					"storageClaim": {
						SchemaProps: spec.SchemaProps{
							Description: "Describes the persistent storage request of each DSE node",
							Ref:         ref("github.com/riptano/dse-operator/operator/pkg/apis/datastax/v1alpha1.DseStorageClaim"),
						},
					},
					"dseClusterName": {
						SchemaProps: spec.SchemaProps{
							Description: "The name by which CQL clients and DSE instances will know the DSE cluster. If the same cluster name is shared by multiple DseDatacenters in the same Kubernetes namespace, they will join together in a multi-datacenter DSE cluster.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"parked": {
						SchemaProps: spec.SchemaProps{
							Description: "Indicates no DSE nodes should run, like powering down bare metal servers. Volume resources will be left intact in Kubernetes and re-attached when the cluster is unparked. This is an experimental feature that requires that pod ip addresses do not change on restart.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"configBuilderImage": {
						SchemaProps: spec.SchemaProps{
							Description: "Container image for the DSE config builder init container, with host, path, and tag",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"canaryUpgrade": {
						SchemaProps: spec.SchemaProps{
							Description: "Indicates DSE configuration and container image changes should only be pushed to the first rack of the datacenter",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"allowMultipleNodesPerWorker": {
						SchemaProps: spec.SchemaProps{
							Description: "Turning this option on allows multiple DSE pods to be created on a k8s worker node. By default the operator creates just one DSE pod per k8s worker node using k8s podAntiAffinity and requiredDuringSchedulingIgnoredDuringExecution.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"dseSuperuserSecret": {
						SchemaProps: spec.SchemaProps{
							Description: "This secret defines the username and password for the DSE superuser.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"serviceAccount": {
						SchemaProps: spec.SchemaProps{
							Description: "The k8s service account to use for DSE pods",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"size", "dseVersion", "dseClusterName"},
			},
		},
		Dependencies: []string{
			"github.com/riptano/dse-operator/operator/pkg/apis/datastax/v1alpha1.DseRack", "github.com/riptano/dse-operator/operator/pkg/apis/datastax/v1alpha1.DseStorageClaim", "github.com/riptano/dse-operator/operator/pkg/apis/datastax/v1alpha1.ManagementApiAuthConfig", "k8s.io/api/core/v1.ResourceRequirements"},
	}
}

func schema_pkg_apis_datastax_v1alpha1_DseDatacenterStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "DseDatacenterStatus defines the observed state of DseDatacenter",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"nodes": {
						SchemaProps: spec.SchemaProps{
							Description: "The number of the DSE server nodes",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"superUserUpserted": {
						SchemaProps: spec.SchemaProps{
							Description: "The timestamp at which CQL superuser credentials were last upserted to the DSE management API",
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
					"lastDseNodeStarted": {
						SchemaProps: spec.SchemaProps{
							Description: "The timestamp when the operator last started a DSE node with the management API",
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
				},
				Required: []string{"nodes"},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.Time"},
	}
}
