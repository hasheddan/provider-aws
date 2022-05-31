/*
Copyright 2022 Upbound Inc.
*/

// Code generated by terrajet. DO NOT EDIT.

package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type VPCIPv4CidrBlockAssociationObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type VPCIPv4CidrBlockAssociationParameters struct {

	// +kubebuilder:validation:Optional
	CidrBlock *string `json:"cidrBlock,omitempty" tf:"cidr_block,omitempty"`

	// +kubebuilder:validation:Optional
	IPv4IpamPoolID *string `json:"ipv4IpamPoolId,omitempty" tf:"ipv4_ipam_pool_id,omitempty"`

	// +kubebuilder:validation:Optional
	IPv4NetmaskLength *float64 `json:"ipv4NetmaskLength,omitempty" tf:"ipv4_netmask_length,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/ec2/v1alpha2.VPC
	// +crossplane:generate:reference:refFieldName=VpcIdRef
	// +crossplane:generate:reference:selectorFieldName=VpcIdSelector
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// +kubebuilder:validation:Optional
	VpcIdRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	VpcIdSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

// VPCIPv4CidrBlockAssociationSpec defines the desired state of VPCIPv4CidrBlockAssociation
type VPCIPv4CidrBlockAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VPCIPv4CidrBlockAssociationParameters `json:"forProvider"`
}

// VPCIPv4CidrBlockAssociationStatus defines the observed state of VPCIPv4CidrBlockAssociation.
type VPCIPv4CidrBlockAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VPCIPv4CidrBlockAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VPCIPv4CidrBlockAssociation is the Schema for the VPCIPv4CidrBlockAssociations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type VPCIPv4CidrBlockAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VPCIPv4CidrBlockAssociationSpec   `json:"spec"`
	Status            VPCIPv4CidrBlockAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VPCIPv4CidrBlockAssociationList contains a list of VPCIPv4CidrBlockAssociations
type VPCIPv4CidrBlockAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VPCIPv4CidrBlockAssociation `json:"items"`
}

// Repository type metadata.
var (
	VPCIPv4CidrBlockAssociation_Kind             = "VPCIPv4CidrBlockAssociation"
	VPCIPv4CidrBlockAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VPCIPv4CidrBlockAssociation_Kind}.String()
	VPCIPv4CidrBlockAssociation_KindAPIVersion   = VPCIPv4CidrBlockAssociation_Kind + "." + CRDGroupVersion.String()
	VPCIPv4CidrBlockAssociation_GroupVersionKind = CRDGroupVersion.WithKind(VPCIPv4CidrBlockAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&VPCIPv4CidrBlockAssociation{}, &VPCIPv4CidrBlockAssociationList{})
}
