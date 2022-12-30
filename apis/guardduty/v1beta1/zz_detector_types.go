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

type DatasourcesObservation struct {
}

type DatasourcesParameters struct {

	// Describes whether S3 data event logs are enabled as a data source. See S3 Logs below for more details.
	// +kubebuilder:validation:Optional
	S3Logs []S3LogsParameters `json:"s3Logs,omitempty" tf:"s3_logs,omitempty"`
}

type DetectorObservation struct {

	// The AWS account ID of the GuardDuty detector
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// Amazon Resource Name (ARN) of the GuardDuty detector
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The ID of the GuardDuty detector
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type DetectorParameters struct {

	// Describes which data sources will be enabled for the detector. See Data Sources below for more details.
	// +kubebuilder:validation:Optional
	Datasources []DatasourcesParameters `json:"datasources,omitempty" tf:"datasources,omitempty"`

	// Enable monitoring and feedback reporting. Setting to false is equivalent to "suspending" GuardDuty. Defaults to true.
	// +kubebuilder:validation:Optional
	Enable *bool `json:"enable,omitempty" tf:"enable,omitempty"`

	// Specifies the frequency of notifications sent for subsequent finding occurrences. If the detector is a GuardDuty member account, the value is determined by the GuardDuty primary account and cannot be modified, otherwise defaults to SIX_HOURS. Valid values for standalone and primary accounts: FIFTEEN_MINUTES, ONE_HOUR, SIX_HOURS. See AWS Documentation for more information.
	// +kubebuilder:validation:Optional
	FindingPublishingFrequency *string `json:"findingPublishingFrequency,omitempty" tf:"finding_publishing_frequency,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type S3LogsObservation struct {
}

type S3LogsParameters struct {

	// If true, enables S3 Protection. Defaults to true.
	// +kubebuilder:validation:Required
	Enable *bool `json:"enable" tf:"enable,omitempty"`
}

// DetectorSpec defines the desired state of Detector
type DetectorSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DetectorParameters `json:"forProvider"`
}

// DetectorStatus defines the observed state of Detector.
type DetectorStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DetectorObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Detector is the Schema for the Detectors API. Provides a resource to manage a GuardDuty detector
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Detector struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DetectorSpec   `json:"spec"`
	Status            DetectorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DetectorList contains a list of Detectors
type DetectorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Detector `json:"items"`
}

// Repository type metadata.
var (
	Detector_Kind             = "Detector"
	Detector_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Detector_Kind}.String()
	Detector_KindAPIVersion   = Detector_Kind + "." + CRDGroupVersion.String()
	Detector_GroupVersionKind = CRDGroupVersion.WithKind(Detector_Kind)
)

func init() {
	SchemeBuilder.Register(&Detector{}, &DetectorList{})
}
