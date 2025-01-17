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

type StreamConsumerObservation struct {

	// Amazon Resource Name (ARN) of the stream consumer.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Approximate timestamp in RFC3339 format of when the stream consumer was created.
	CreationTimestamp *string `json:"creationTimestamp,omitempty" tf:"creation_timestamp,omitempty"`

	// Amazon Resource Name (ARN) of the stream consumer.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type StreamConsumerParameters struct {

	// Name of the stream consumer.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// –  Amazon Resource Name (ARN) of the data stream the consumer is registered with.
	// +crossplane:generate:reference:type=Stream
	// +kubebuilder:validation:Optional
	StreamArn *string `json:"streamArn,omitempty" tf:"stream_arn,omitempty"`

	// Reference to a Stream to populate streamArn.
	// +kubebuilder:validation:Optional
	StreamArnRef *v1.Reference `json:"streamArnRef,omitempty" tf:"-"`

	// Selector for a Stream to populate streamArn.
	// +kubebuilder:validation:Optional
	StreamArnSelector *v1.Selector `json:"streamArnSelector,omitempty" tf:"-"`
}

// StreamConsumerSpec defines the desired state of StreamConsumer
type StreamConsumerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     StreamConsumerParameters `json:"forProvider"`
}

// StreamConsumerStatus defines the observed state of StreamConsumer.
type StreamConsumerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        StreamConsumerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// StreamConsumer is the Schema for the StreamConsumers API. Manages a Kinesis Stream Consumer.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type StreamConsumer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StreamConsumerSpec   `json:"spec"`
	Status            StreamConsumerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StreamConsumerList contains a list of StreamConsumers
type StreamConsumerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StreamConsumer `json:"items"`
}

// Repository type metadata.
var (
	StreamConsumer_Kind             = "StreamConsumer"
	StreamConsumer_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: StreamConsumer_Kind}.String()
	StreamConsumer_KindAPIVersion   = StreamConsumer_Kind + "." + CRDGroupVersion.String()
	StreamConsumer_GroupVersionKind = CRDGroupVersion.WithKind(StreamConsumer_Kind)
)

func init() {
	SchemeBuilder.Register(&StreamConsumer{}, &StreamConsumerList{})
}
