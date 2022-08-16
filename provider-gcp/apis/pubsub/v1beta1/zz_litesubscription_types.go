/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type DeliveryConfigObservation struct {
}

type DeliveryConfigParameters struct {

	// When this subscription should send messages to subscribers relative to messages persistence in storage.
	// Possible values are DELIVER_IMMEDIATELY, DELIVER_AFTER_STORED, and DELIVERY_REQUIREMENT_UNSPECIFIED.
	// When this subscription should send messages to subscribers relative to messages persistence in storage. Possible values: ["DELIVER_IMMEDIATELY", "DELIVER_AFTER_STORED", "DELIVERY_REQUIREMENT_UNSPECIFIED"]
	// +kubebuilder:validation:Required
	DeliveryRequirement *string `json:"deliveryRequirement" tf:"delivery_requirement,omitempty"`
}

type LiteSubscriptionObservation struct {

	// an identifier for the resource with format projects/{{project}}/locations/{{zone}}/subscriptions/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type LiteSubscriptionParameters struct {

	// The settings for this subscription's message delivery.
	// Structure is documented below.
	// The settings for this subscription's message delivery.
	// +kubebuilder:validation:Optional
	DeliveryConfig []DeliveryConfigParameters `json:"deliveryConfig,omitempty" tf:"delivery_config,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The region of the pubsub lite topic.
	// The region of the pubsub lite topic.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// A reference to a Topic resource.
	// A reference to a Topic resource.
	// +crossplane:generate:reference:type=LiteTopic
	// +kubebuilder:validation:Optional
	Topic *string `json:"topic,omitempty" tf:"topic,omitempty"`

	// +kubebuilder:validation:Optional
	TopicRef *v1.Reference `json:"topicRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	TopicSelector *v1.Selector `json:"topicSelector,omitempty" tf:"-"`

	// The zone of the pubsub lite topic.
	// The zone of the pubsub lite topic.
	// +kubebuilder:validation:Required
	Zone *string `json:"zone" tf:"zone,omitempty"`
}

// LiteSubscriptionSpec defines the desired state of LiteSubscription
type LiteSubscriptionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LiteSubscriptionParameters `json:"forProvider"`
}

// LiteSubscriptionStatus defines the observed state of LiteSubscription.
type LiteSubscriptionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LiteSubscriptionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LiteSubscription is the Schema for the LiteSubscriptions API. A named resource representing the stream of messages from a single, specific topic, to be delivered to the subscribing application.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type LiteSubscription struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LiteSubscriptionSpec   `json:"spec"`
	Status            LiteSubscriptionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LiteSubscriptionList contains a list of LiteSubscriptions
type LiteSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LiteSubscription `json:"items"`
}

// Repository type metadata.
var (
	LiteSubscription_Kind             = "LiteSubscription"
	LiteSubscription_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LiteSubscription_Kind}.String()
	LiteSubscription_KindAPIVersion   = LiteSubscription_Kind + "." + CRDGroupVersion.String()
	LiteSubscription_GroupVersionKind = CRDGroupVersion.WithKind(LiteSubscription_Kind)
)

func init() {
	SchemeBuilder.Register(&LiteSubscription{}, &LiteSubscriptionList{})
}
