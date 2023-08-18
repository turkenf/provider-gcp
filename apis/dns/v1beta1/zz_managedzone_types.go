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

type CloudLoggingConfigInitParameters struct {

	// If set, enable query logging for this ManagedZone. False by default, making logging opt-in.
	EnableLogging *bool `json:"enableLogging,omitempty" tf:"enable_logging,omitempty"`
}

type CloudLoggingConfigObservation struct {

	// If set, enable query logging for this ManagedZone. False by default, making logging opt-in.
	EnableLogging *bool `json:"enableLogging,omitempty" tf:"enable_logging,omitempty"`
}

type CloudLoggingConfigParameters struct {

	// If set, enable query logging for this ManagedZone. False by default, making logging opt-in.
	// +kubebuilder:validation:Optional
	EnableLogging *bool `json:"enableLogging" tf:"enable_logging,omitempty"`
}

type DNSSECConfigInitParameters struct {

	// Specifies parameters that will be used for generating initial DnsKeys
	// for this ManagedZone. If you provide a spec for keySigning or zoneSigning,
	// you must also provide one for the other.
	// default_key_specs can only be updated when the state is off.
	// Structure is documented below.
	DefaultKeySpecs []DefaultKeySpecsInitParameters `json:"defaultKeySpecs,omitempty" tf:"default_key_specs,omitempty"`

	// Identifies what kind of resource this is
	Kind *string `json:"kind,omitempty" tf:"kind,omitempty"`

	// Specifies the mechanism used to provide authenticated denial-of-existence responses.
	// non_existence can only be updated when the state is off.
	// Possible values are: nsec, nsec3.
	NonExistence *string `json:"nonExistence,omitempty" tf:"non_existence,omitempty"`

	// Specifies whether DNSSEC is enabled, and what mode it is in
	// Possible values are: off, on, transfer.
	State *string `json:"state,omitempty" tf:"state,omitempty"`
}

type DNSSECConfigObservation struct {

	// Specifies parameters that will be used for generating initial DnsKeys
	// for this ManagedZone. If you provide a spec for keySigning or zoneSigning,
	// you must also provide one for the other.
	// default_key_specs can only be updated when the state is off.
	// Structure is documented below.
	DefaultKeySpecs []DefaultKeySpecsObservation `json:"defaultKeySpecs,omitempty" tf:"default_key_specs,omitempty"`

	// Identifies what kind of resource this is
	Kind *string `json:"kind,omitempty" tf:"kind,omitempty"`

	// Specifies the mechanism used to provide authenticated denial-of-existence responses.
	// non_existence can only be updated when the state is off.
	// Possible values are: nsec, nsec3.
	NonExistence *string `json:"nonExistence,omitempty" tf:"non_existence,omitempty"`

	// Specifies whether DNSSEC is enabled, and what mode it is in
	// Possible values are: off, on, transfer.
	State *string `json:"state,omitempty" tf:"state,omitempty"`
}

type DNSSECConfigParameters struct {

	// Specifies parameters that will be used for generating initial DnsKeys
	// for this ManagedZone. If you provide a spec for keySigning or zoneSigning,
	// you must also provide one for the other.
	// default_key_specs can only be updated when the state is off.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	DefaultKeySpecs []DefaultKeySpecsParameters `json:"defaultKeySpecs,omitempty" tf:"default_key_specs,omitempty"`

	// Identifies what kind of resource this is
	// +kubebuilder:validation:Optional
	Kind *string `json:"kind,omitempty" tf:"kind,omitempty"`

	// Specifies the mechanism used to provide authenticated denial-of-existence responses.
	// non_existence can only be updated when the state is off.
	// Possible values are: nsec, nsec3.
	// +kubebuilder:validation:Optional
	NonExistence *string `json:"nonExistence,omitempty" tf:"non_existence,omitempty"`

	// Specifies whether DNSSEC is enabled, and what mode it is in
	// Possible values are: off, on, transfer.
	// +kubebuilder:validation:Optional
	State *string `json:"state,omitempty" tf:"state,omitempty"`
}

type DefaultKeySpecsInitParameters struct {

	// String mnemonic specifying the DNSSEC algorithm of this key
	// Possible values are: ecdsap256sha256, ecdsap384sha384, rsasha1, rsasha256, rsasha512.
	Algorithm *string `json:"algorithm,omitempty" tf:"algorithm,omitempty"`

	// Length of the keys in bits
	KeyLength *float64 `json:"keyLength,omitempty" tf:"key_length,omitempty"`

	// Specifies whether this is a key signing key (KSK) or a zone
	// signing key (ZSK). Key signing keys have the Secure Entry
	// Point flag set and, when active, will only be used to sign
	// resource record sets of type DNSKEY. Zone signing keys do
	// not have the Secure Entry Point flag set and will be used
	// to sign all other types of resource record sets.
	// Possible values are: keySigning, zoneSigning.
	KeyType *string `json:"keyType,omitempty" tf:"key_type,omitempty"`

	// Identifies what kind of resource this is
	Kind *string `json:"kind,omitempty" tf:"kind,omitempty"`
}

type DefaultKeySpecsObservation struct {

	// String mnemonic specifying the DNSSEC algorithm of this key
	// Possible values are: ecdsap256sha256, ecdsap384sha384, rsasha1, rsasha256, rsasha512.
	Algorithm *string `json:"algorithm,omitempty" tf:"algorithm,omitempty"`

	// Length of the keys in bits
	KeyLength *float64 `json:"keyLength,omitempty" tf:"key_length,omitempty"`

	// Specifies whether this is a key signing key (KSK) or a zone
	// signing key (ZSK). Key signing keys have the Secure Entry
	// Point flag set and, when active, will only be used to sign
	// resource record sets of type DNSKEY. Zone signing keys do
	// not have the Secure Entry Point flag set and will be used
	// to sign all other types of resource record sets.
	// Possible values are: keySigning, zoneSigning.
	KeyType *string `json:"keyType,omitempty" tf:"key_type,omitempty"`

	// Identifies what kind of resource this is
	Kind *string `json:"kind,omitempty" tf:"kind,omitempty"`
}

type DefaultKeySpecsParameters struct {

	// String mnemonic specifying the DNSSEC algorithm of this key
	// Possible values are: ecdsap256sha256, ecdsap384sha384, rsasha1, rsasha256, rsasha512.
	// +kubebuilder:validation:Optional
	Algorithm *string `json:"algorithm,omitempty" tf:"algorithm,omitempty"`

	// Length of the keys in bits
	// +kubebuilder:validation:Optional
	KeyLength *float64 `json:"keyLength,omitempty" tf:"key_length,omitempty"`

	// Specifies whether this is a key signing key (KSK) or a zone
	// signing key (ZSK). Key signing keys have the Secure Entry
	// Point flag set and, when active, will only be used to sign
	// resource record sets of type DNSKEY. Zone signing keys do
	// not have the Secure Entry Point flag set and will be used
	// to sign all other types of resource record sets.
	// Possible values are: keySigning, zoneSigning.
	// +kubebuilder:validation:Optional
	KeyType *string `json:"keyType,omitempty" tf:"key_type,omitempty"`

	// Identifies what kind of resource this is
	// +kubebuilder:validation:Optional
	Kind *string `json:"kind,omitempty" tf:"kind,omitempty"`
}

type ForwardingConfigInitParameters struct {

	// List of target name servers to forward to. Cloud DNS will
	// select the best available name server if more than
	// one target is given.
	// Structure is documented below.
	TargetNameServers []TargetNameServersInitParameters `json:"targetNameServers,omitempty" tf:"target_name_servers,omitempty"`
}

type ForwardingConfigObservation struct {

	// List of target name servers to forward to. Cloud DNS will
	// select the best available name server if more than
	// one target is given.
	// Structure is documented below.
	TargetNameServers []TargetNameServersObservation `json:"targetNameServers,omitempty" tf:"target_name_servers,omitempty"`
}

type ForwardingConfigParameters struct {

	// List of target name servers to forward to. Cloud DNS will
	// select the best available name server if more than
	// one target is given.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	TargetNameServers []TargetNameServersParameters `json:"targetNameServers" tf:"target_name_servers,omitempty"`
}

type GkeClustersInitParameters struct {
}

type GkeClustersObservation struct {

	// The resource name of the cluster to bind this ManagedZone to.
	// This should be specified in the format like
	// projects/*/locations/*/clusters/*
	GkeClusterName *string `json:"gkeClusterName,omitempty" tf:"gke_cluster_name,omitempty"`
}

type GkeClustersParameters struct {

	// The resource name of the cluster to bind this ManagedZone to.
	// This should be specified in the format like
	// projects/*/locations/*/clusters/*
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/container/v1beta1.Cluster
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	GkeClusterName *string `json:"gkeClusterName,omitempty" tf:"gke_cluster_name,omitempty"`

	// Reference to a Cluster in container to populate gkeClusterName.
	// +kubebuilder:validation:Optional
	GkeClusterNameRef *v1.Reference `json:"gkeClusterNameRef,omitempty" tf:"-"`

	// Selector for a Cluster in container to populate gkeClusterName.
	// +kubebuilder:validation:Optional
	GkeClusterNameSelector *v1.Selector `json:"gkeClusterNameSelector,omitempty" tf:"-"`
}

type ManagedZoneInitParameters struct {

	// Cloud logging configuration
	// Structure is documented below.
	CloudLoggingConfig []CloudLoggingConfigInitParameters `json:"cloudLoggingConfig,omitempty" tf:"cloud_logging_config,omitempty"`

	// The DNS name of this managed zone, for instance "example.com.".
	DNSName *string `json:"dnsName,omitempty" tf:"dns_name,omitempty"`

	// DNSSEC configuration
	// Structure is documented below.
	DNSSECConfig []DNSSECConfigInitParameters `json:"dnssecConfig,omitempty" tf:"dnssec_config,omitempty"`

	// A textual description field.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Set this true to delete all records in the zone.
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// The presence for this field indicates that outbound forwarding is enabled
	// for this zone. The value of this field contains the set of destinations
	// to forward to.
	// Structure is documented below.
	ForwardingConfig []ForwardingConfigInitParameters `json:"forwardingConfig,omitempty" tf:"forwarding_config,omitempty"`

	// A set of key/value label pairs to assign to this ManagedZone.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The presence of this field indicates that DNS Peering is enabled for this
	// zone. The value of this field contains the network to peer with.
	// Structure is documented below.
	PeeringConfig []PeeringConfigInitParameters `json:"peeringConfig,omitempty" tf:"peering_config,omitempty"`

	// For privately visible zones, the set of Virtual Private Cloud
	// resources that the zone is visible from.
	// Structure is documented below.
	PrivateVisibilityConfig []PrivateVisibilityConfigInitParameters `json:"privateVisibilityConfig,omitempty" tf:"private_visibility_config,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The zone's visibility: public zones are exposed to the Internet,
	// while private zones are visible only to Virtual Private Cloud resources.
	// Default value is public.
	// Possible values are: private, public.
	Visibility *string `json:"visibility,omitempty" tf:"visibility,omitempty"`
}

type ManagedZoneObservation struct {

	// Cloud logging configuration
	// Structure is documented below.
	CloudLoggingConfig []CloudLoggingConfigObservation `json:"cloudLoggingConfig,omitempty" tf:"cloud_logging_config,omitempty"`

	// The time that this resource was created on the server.
	// This is in RFC3339 text format.
	CreationTime *string `json:"creationTime,omitempty" tf:"creation_time,omitempty"`

	// The DNS name of this managed zone, for instance "example.com.".
	DNSName *string `json:"dnsName,omitempty" tf:"dns_name,omitempty"`

	// DNSSEC configuration
	// Structure is documented below.
	DNSSECConfig []DNSSECConfigObservation `json:"dnssecConfig,omitempty" tf:"dnssec_config,omitempty"`

	// A textual description field.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Set this true to delete all records in the zone.
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// The presence for this field indicates that outbound forwarding is enabled
	// for this zone. The value of this field contains the set of destinations
	// to forward to.
	// Structure is documented below.
	ForwardingConfig []ForwardingConfigObservation `json:"forwardingConfig,omitempty" tf:"forwarding_config,omitempty"`

	// an identifier for the resource with format projects/{{project}}/managedZones/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A set of key/value label pairs to assign to this ManagedZone.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Unique identifier for the resource; defined by the server.
	ManagedZoneID *float64 `json:"managedZoneId,omitempty" tf:"managed_zone_id,omitempty"`

	// Delegate your managed_zone to these virtual name servers;
	// defined by the server
	NameServers []*string `json:"nameServers,omitempty" tf:"name_servers,omitempty"`

	// The presence of this field indicates that DNS Peering is enabled for this
	// zone. The value of this field contains the network to peer with.
	// Structure is documented below.
	PeeringConfig []PeeringConfigObservation `json:"peeringConfig,omitempty" tf:"peering_config,omitempty"`

	// For privately visible zones, the set of Virtual Private Cloud
	// resources that the zone is visible from.
	// Structure is documented below.
	PrivateVisibilityConfig []PrivateVisibilityConfigObservation `json:"privateVisibilityConfig,omitempty" tf:"private_visibility_config,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The zone's visibility: public zones are exposed to the Internet,
	// while private zones are visible only to Virtual Private Cloud resources.
	// Default value is public.
	// Possible values are: private, public.
	Visibility *string `json:"visibility,omitempty" tf:"visibility,omitempty"`
}

type ManagedZoneParameters struct {

	// Cloud logging configuration
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	CloudLoggingConfig []CloudLoggingConfigParameters `json:"cloudLoggingConfig,omitempty" tf:"cloud_logging_config,omitempty"`

	// The DNS name of this managed zone, for instance "example.com.".
	// +kubebuilder:validation:Optional
	DNSName *string `json:"dnsName,omitempty" tf:"dns_name,omitempty"`

	// DNSSEC configuration
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	DNSSECConfig []DNSSECConfigParameters `json:"dnssecConfig,omitempty" tf:"dnssec_config,omitempty"`

	// A textual description field.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Set this true to delete all records in the zone.
	// +kubebuilder:validation:Optional
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// The presence for this field indicates that outbound forwarding is enabled
	// for this zone. The value of this field contains the set of destinations
	// to forward to.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	ForwardingConfig []ForwardingConfigParameters `json:"forwardingConfig,omitempty" tf:"forwarding_config,omitempty"`

	// A set of key/value label pairs to assign to this ManagedZone.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The presence of this field indicates that DNS Peering is enabled for this
	// zone. The value of this field contains the network to peer with.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	PeeringConfig []PeeringConfigParameters `json:"peeringConfig,omitempty" tf:"peering_config,omitempty"`

	// For privately visible zones, the set of Virtual Private Cloud
	// resources that the zone is visible from.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	PrivateVisibilityConfig []PrivateVisibilityConfigParameters `json:"privateVisibilityConfig,omitempty" tf:"private_visibility_config,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The zone's visibility: public zones are exposed to the Internet,
	// while private zones are visible only to Virtual Private Cloud resources.
	// Default value is public.
	// Possible values are: private, public.
	// +kubebuilder:validation:Optional
	Visibility *string `json:"visibility,omitempty" tf:"visibility,omitempty"`
}

type NetworksInitParameters struct {
}

type NetworksObservation struct {

	// The id or fully qualified URL of the VPC network to forward queries to.
	// This should be formatted like projects/{project}/global/networks/{network} or
	// https://www.googleapis.com/compute/v1/projects/{project}/global/networks/{network}
	NetworkURL *string `json:"networkUrl,omitempty" tf:"network_url,omitempty"`
}

type NetworksParameters struct {

	// The id or fully qualified URL of the VPC network to forward queries to.
	// This should be formatted like projects/{project}/global/networks/{network} or
	// https://www.googleapis.com/compute/v1/projects/{project}/global/networks/{network}
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.Network
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-gcp/config/common.SelfLinkExtractor()
	// +kubebuilder:validation:Optional
	NetworkURL *string `json:"networkUrl,omitempty" tf:"network_url,omitempty"`

	// Reference to a Network in compute to populate networkUrl.
	// +kubebuilder:validation:Optional
	NetworkURLRef *v1.Reference `json:"networkUrlRef,omitempty" tf:"-"`

	// Selector for a Network in compute to populate networkUrl.
	// +kubebuilder:validation:Optional
	NetworkURLSelector *v1.Selector `json:"networkUrlSelector,omitempty" tf:"-"`
}

type PeeringConfigInitParameters struct {

	// The network with which to peer.
	// Structure is documented below.
	TargetNetwork []TargetNetworkInitParameters `json:"targetNetwork,omitempty" tf:"target_network,omitempty"`
}

type PeeringConfigObservation struct {

	// The network with which to peer.
	// Structure is documented below.
	TargetNetwork []TargetNetworkObservation `json:"targetNetwork,omitempty" tf:"target_network,omitempty"`
}

type PeeringConfigParameters struct {

	// The network with which to peer.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	TargetNetwork []TargetNetworkParameters `json:"targetNetwork" tf:"target_network,omitempty"`
}

type PrivateVisibilityConfigInitParameters struct {

	// The list of Google Kubernetes Engine clusters that can see this zone.
	// Structure is documented below.
	GkeClusters []GkeClustersInitParameters `json:"gkeClusters,omitempty" tf:"gke_clusters,omitempty"`

	// The list of VPC networks that can see this zone.12 SDK in a future release, you
	// may experience issues with this resource while updating. If you encounter this issue, remove all networks
	// blocks in an update and then apply another update adding all of them back simultaneously.
	// Structure is documented below.
	Networks []NetworksInitParameters `json:"networks,omitempty" tf:"networks,omitempty"`
}

type PrivateVisibilityConfigObservation struct {

	// The list of Google Kubernetes Engine clusters that can see this zone.
	// Structure is documented below.
	GkeClusters []GkeClustersObservation `json:"gkeClusters,omitempty" tf:"gke_clusters,omitempty"`

	// The list of VPC networks that can see this zone.12 SDK in a future release, you
	// may experience issues with this resource while updating. If you encounter this issue, remove all networks
	// blocks in an update and then apply another update adding all of them back simultaneously.
	// Structure is documented below.
	Networks []NetworksObservation `json:"networks,omitempty" tf:"networks,omitempty"`
}

type PrivateVisibilityConfigParameters struct {

	// The list of Google Kubernetes Engine clusters that can see this zone.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	GkeClusters []GkeClustersParameters `json:"gkeClusters,omitempty" tf:"gke_clusters,omitempty"`

	// The list of VPC networks that can see this zone.12 SDK in a future release, you
	// may experience issues with this resource while updating. If you encounter this issue, remove all networks
	// blocks in an update and then apply another update adding all of them back simultaneously.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Networks []NetworksParameters `json:"networks" tf:"networks,omitempty"`
}

type TargetNameServersInitParameters struct {

	// Forwarding path for this TargetNameServer. If unset or default Cloud DNS will make forwarding
	// decision based on address ranges, i.e. RFC1918 addresses go to the VPC, Non-RFC1918 addresses go
	// to the Internet. When set to private, Cloud DNS will always send queries through VPC for this target
	// Possible values are: default, private.
	ForwardingPath *string `json:"forwardingPath,omitempty" tf:"forwarding_path,omitempty"`

	// IPv4 address of a target name server.
	IPv4Address *string `json:"ipv4Address,omitempty" tf:"ipv4_address,omitempty"`
}

type TargetNameServersObservation struct {

	// Forwarding path for this TargetNameServer. If unset or default Cloud DNS will make forwarding
	// decision based on address ranges, i.e. RFC1918 addresses go to the VPC, Non-RFC1918 addresses go
	// to the Internet. When set to private, Cloud DNS will always send queries through VPC for this target
	// Possible values are: default, private.
	ForwardingPath *string `json:"forwardingPath,omitempty" tf:"forwarding_path,omitempty"`

	// IPv4 address of a target name server.
	IPv4Address *string `json:"ipv4Address,omitempty" tf:"ipv4_address,omitempty"`
}

type TargetNameServersParameters struct {

	// Forwarding path for this TargetNameServer. If unset or default Cloud DNS will make forwarding
	// decision based on address ranges, i.e. RFC1918 addresses go to the VPC, Non-RFC1918 addresses go
	// to the Internet. When set to private, Cloud DNS will always send queries through VPC for this target
	// Possible values are: default, private.
	// +kubebuilder:validation:Optional
	ForwardingPath *string `json:"forwardingPath,omitempty" tf:"forwarding_path,omitempty"`

	// IPv4 address of a target name server.
	// +kubebuilder:validation:Optional
	IPv4Address *string `json:"ipv4Address" tf:"ipv4_address,omitempty"`
}

type TargetNetworkInitParameters struct {
}

type TargetNetworkObservation struct {

	// The id or fully qualified URL of the VPC network to forward queries to.
	// This should be formatted like projects/{project}/global/networks/{network} or
	// https://www.googleapis.com/compute/v1/projects/{project}/global/networks/{network}
	NetworkURL *string `json:"networkUrl,omitempty" tf:"network_url,omitempty"`
}

type TargetNetworkParameters struct {

	// The id or fully qualified URL of the VPC network to forward queries to.
	// This should be formatted like projects/{project}/global/networks/{network} or
	// https://www.googleapis.com/compute/v1/projects/{project}/global/networks/{network}
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.Network
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-gcp/config/common.SelfLinkExtractor()
	// +kubebuilder:validation:Optional
	NetworkURL *string `json:"networkUrl,omitempty" tf:"network_url,omitempty"`

	// Reference to a Network in compute to populate networkUrl.
	// +kubebuilder:validation:Optional
	NetworkURLRef *v1.Reference `json:"networkUrlRef,omitempty" tf:"-"`

	// Selector for a Network in compute to populate networkUrl.
	// +kubebuilder:validation:Optional
	NetworkURLSelector *v1.Selector `json:"networkUrlSelector,omitempty" tf:"-"`
}

// ManagedZoneSpec defines the desired state of ManagedZone
type ManagedZoneSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ManagedZoneParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider ManagedZoneInitParameters `json:"initProvider,omitempty"`
}

// ManagedZoneStatus defines the observed state of ManagedZone.
type ManagedZoneStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ManagedZoneObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ManagedZone is the Schema for the ManagedZones API. A zone is a subtree of the DNS namespace under one administrative responsibility.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type ManagedZone struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.dnsName) || has(self.initProvider.dnsName)",message="dnsName is a required parameter"
	Spec   ManagedZoneSpec   `json:"spec"`
	Status ManagedZoneStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ManagedZoneList contains a list of ManagedZones
type ManagedZoneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManagedZone `json:"items"`
}

// Repository type metadata.
var (
	ManagedZone_Kind             = "ManagedZone"
	ManagedZone_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ManagedZone_Kind}.String()
	ManagedZone_KindAPIVersion   = ManagedZone_Kind + "." + CRDGroupVersion.String()
	ManagedZone_GroupVersionKind = CRDGroupVersion.WithKind(ManagedZone_Kind)
)

func init() {
	SchemeBuilder.Register(&ManagedZone{}, &ManagedZoneList{})
}
