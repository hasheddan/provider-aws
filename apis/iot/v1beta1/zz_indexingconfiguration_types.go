/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type CustomFieldObservation struct {
}

type CustomFieldParameters struct {

	// The name of the field.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The data type of the field. Valid values: Number, String, Boolean.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type IndexingConfigurationObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type IndexingConfigurationParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Thing group indexing configuration. See below.
	// +kubebuilder:validation:Optional
	ThingGroupIndexingConfiguration []ThingGroupIndexingConfigurationParameters `json:"thingGroupIndexingConfiguration,omitempty" tf:"thing_group_indexing_configuration,omitempty"`

	// Thing indexing configuration. See below.
	// +kubebuilder:validation:Optional
	ThingIndexingConfiguration []ThingIndexingConfigurationParameters `json:"thingIndexingConfiguration,omitempty" tf:"thing_indexing_configuration,omitempty"`
}

type ManagedFieldObservation struct {
}

type ManagedFieldParameters struct {

	// The name of the field.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The data type of the field. Valid values: Number, String, Boolean.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ThingGroupIndexingConfigurationObservation struct {
}

type ThingGroupIndexingConfigurationParameters struct {

	// A list of thing group fields to index. This list cannot contain any managed fields. See below.
	// +kubebuilder:validation:Optional
	CustomField []CustomFieldParameters `json:"customField,omitempty" tf:"custom_field,omitempty"`

	// Contains fields that are indexed and whose types are already known by the Fleet Indexing service. See below.
	// +kubebuilder:validation:Optional
	ManagedField []ManagedFieldParameters `json:"managedField,omitempty" tf:"managed_field,omitempty"`

	// Thing group indexing mode. Valid values: OFF, ON.
	// +kubebuilder:validation:Required
	ThingGroupIndexingMode *string `json:"thingGroupIndexingMode" tf:"thing_group_indexing_mode,omitempty"`
}

type ThingIndexingConfigurationCustomFieldObservation struct {
}

type ThingIndexingConfigurationCustomFieldParameters struct {

	// The name of the field.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The data type of the field. Valid values: Number, String, Boolean.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ThingIndexingConfigurationManagedFieldObservation struct {
}

type ThingIndexingConfigurationManagedFieldParameters struct {

	// The name of the field.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The data type of the field. Valid values: Number, String, Boolean.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ThingIndexingConfigurationObservation struct {
}

type ThingIndexingConfigurationParameters struct {

	// Contains custom field names and their data type. See below.
	// +kubebuilder:validation:Optional
	CustomField []ThingIndexingConfigurationCustomFieldParameters `json:"customField,omitempty" tf:"custom_field,omitempty"`

	// Device Defender indexing mode. Valid values: VIOLATIONS, OFF. Default: OFF.
	// +kubebuilder:validation:Optional
	DeviceDefenderIndexingMode *string `json:"deviceDefenderIndexingMode,omitempty" tf:"device_defender_indexing_mode,omitempty"`

	// Contains fields that are indexed and whose types are already known by the Fleet Indexing service. See below.
	// +kubebuilder:validation:Optional
	ManagedField []ThingIndexingConfigurationManagedFieldParameters `json:"managedField,omitempty" tf:"managed_field,omitempty"`

	// Named shadow indexing mode. Valid values: ON, OFF. Default: OFF.
	// +kubebuilder:validation:Optional
	NamedShadowIndexingMode *string `json:"namedShadowIndexingMode,omitempty" tf:"named_shadow_indexing_mode,omitempty"`

	// Thing connectivity indexing mode. Valid values: STATUS, OFF. Default: OFF.
	// +kubebuilder:validation:Optional
	ThingConnectivityIndexingMode *string `json:"thingConnectivityIndexingMode,omitempty" tf:"thing_connectivity_indexing_mode,omitempty"`

	// Thing indexing mode. Valid values: REGISTRY, REGISTRY_AND_SHADOW, OFF.
	// +kubebuilder:validation:Required
	ThingIndexingMode *string `json:"thingIndexingMode" tf:"thing_indexing_mode,omitempty"`
}

// IndexingConfigurationSpec defines the desired state of IndexingConfiguration
type IndexingConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IndexingConfigurationParameters `json:"forProvider"`
}

// IndexingConfigurationStatus defines the observed state of IndexingConfiguration.
type IndexingConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IndexingConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IndexingConfiguration is the Schema for the IndexingConfigurations API. Managing IoT Thing indexing.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type IndexingConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IndexingConfigurationSpec   `json:"spec"`
	Status            IndexingConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IndexingConfigurationList contains a list of IndexingConfigurations
type IndexingConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IndexingConfiguration `json:"items"`
}

// Repository type metadata.
var (
	IndexingConfiguration_Kind             = "IndexingConfiguration"
	IndexingConfiguration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IndexingConfiguration_Kind}.String()
	IndexingConfiguration_KindAPIVersion   = IndexingConfiguration_Kind + "." + CRDGroupVersion.String()
	IndexingConfiguration_GroupVersionKind = CRDGroupVersion.WithKind(IndexingConfiguration_Kind)
)

func init() {
	SchemeBuilder.Register(&IndexingConfiguration{}, &IndexingConfigurationList{})
}
