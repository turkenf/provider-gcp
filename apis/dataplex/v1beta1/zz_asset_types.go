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

type AssetInitParameters struct {

	// Optional. Description of the asset.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Required. Specification of the discovery feature applied to data referenced by this asset. When this spec is left unset, the asset will use the spec set on the parent zone.
	DiscoverySpec []DiscoverySpecInitParameters `json:"discoverySpec,omitempty" tf:"discovery_spec,omitempty"`

	// Optional. User friendly display name.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Optional. User defined labels for the asset.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The project for the resource
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Required. Immutable. Specification of the resource that is referenced by this asset.
	ResourceSpec []ResourceSpecInitParameters `json:"resourceSpec,omitempty" tf:"resource_spec,omitempty"`
}

type AssetObservation struct {

	// Output only. The time when the asset was created.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// The zone for the resource
	DataplexZone *string `json:"dataplexZone,omitempty" tf:"dataplex_zone,omitempty"`

	// Optional. Description of the asset.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Required. Specification of the discovery feature applied to data referenced by this asset. When this spec is left unset, the asset will use the spec set on the parent zone.
	DiscoverySpec []DiscoverySpecObservation `json:"discoverySpec,omitempty" tf:"discovery_spec,omitempty"`

	// Output only. Status of the discovery feature applied to data referenced by this asset.
	DiscoveryStatus []DiscoveryStatusObservation `json:"discoveryStatus,omitempty" tf:"discovery_status,omitempty"`

	// Optional. User friendly display name.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// an identifier for the resource with format projects/{{project}}/locations/{{location}}/lakes/{{lake}}/zones/{{dataplex_zone}}/assets/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Optional. User defined labels for the asset.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The lake for the resource
	Lake *string `json:"lake,omitempty" tf:"lake,omitempty"`

	// The location for the resource
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The project for the resource
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Required. Immutable. Specification of the resource that is referenced by this asset.
	ResourceSpec []ResourceSpecObservation `json:"resourceSpec,omitempty" tf:"resource_spec,omitempty"`

	// Output only. Status of the resource referenced by this asset.
	ResourceStatus []ResourceStatusObservation `json:"resourceStatus,omitempty" tf:"resource_status,omitempty"`

	// Output only. Status of the security policy applied to resource referenced by this asset.
	SecurityStatus []SecurityStatusObservation `json:"securityStatus,omitempty" tf:"security_status,omitempty"`

	// Output only. Current state of the asset. Possible values: STATE_UNSPECIFIED, ACTIVE, CREATING, DELETING, ACTION_REQUIRED
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// Output only. System generated globally unique ID for the asset. This ID will be different if the asset is deleted and re-created with the same name.
	UID *string `json:"uid,omitempty" tf:"uid,omitempty"`

	// Output only. The time when the asset was last updated.
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`
}

type AssetParameters struct {

	// The zone for the resource
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/dataplex/v1beta1.Zone
	// +kubebuilder:validation:Optional
	DataplexZone *string `json:"dataplexZone,omitempty" tf:"dataplex_zone,omitempty"`

	// Reference to a Zone in dataplex to populate dataplexZone.
	// +kubebuilder:validation:Optional
	DataplexZoneRef *v1.Reference `json:"dataplexZoneRef,omitempty" tf:"-"`

	// Selector for a Zone in dataplex to populate dataplexZone.
	// +kubebuilder:validation:Optional
	DataplexZoneSelector *v1.Selector `json:"dataplexZoneSelector,omitempty" tf:"-"`

	// Optional. Description of the asset.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Required. Specification of the discovery feature applied to data referenced by this asset. When this spec is left unset, the asset will use the spec set on the parent zone.
	// +kubebuilder:validation:Optional
	DiscoverySpec []DiscoverySpecParameters `json:"discoverySpec,omitempty" tf:"discovery_spec,omitempty"`

	// Optional. User friendly display name.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Optional. User defined labels for the asset.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The lake for the resource
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/dataplex/v1beta1.Lake
	// +kubebuilder:validation:Optional
	Lake *string `json:"lake,omitempty" tf:"lake,omitempty"`

	// Reference to a Lake in dataplex to populate lake.
	// +kubebuilder:validation:Optional
	LakeRef *v1.Reference `json:"lakeRef,omitempty" tf:"-"`

	// Selector for a Lake in dataplex to populate lake.
	// +kubebuilder:validation:Optional
	LakeSelector *v1.Selector `json:"lakeSelector,omitempty" tf:"-"`

	// The location for the resource
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// The project for the resource
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Required. Immutable. Specification of the resource that is referenced by this asset.
	// +kubebuilder:validation:Optional
	ResourceSpec []ResourceSpecParameters `json:"resourceSpec,omitempty" tf:"resource_spec,omitempty"`
}

type CsvOptionsInitParameters struct {

	// Optional. The delimiter being used to separate values. This defaults to ','.
	Delimiter *string `json:"delimiter,omitempty" tf:"delimiter,omitempty"`

	// Optional. Whether to disable the inference of data type for Json data. If true, all columns will be registered as their primitive types (strings, number or boolean).
	DisableTypeInference *bool `json:"disableTypeInference,omitempty" tf:"disable_type_inference,omitempty"`

	// Optional. The character encoding of the data. The default is UTF-8.
	Encoding *string `json:"encoding,omitempty" tf:"encoding,omitempty"`

	// Optional. The number of rows to interpret as header rows that should be skipped when reading data rows.
	HeaderRows *float64 `json:"headerRows,omitempty" tf:"header_rows,omitempty"`
}

type CsvOptionsObservation struct {

	// Optional. The delimiter being used to separate values. This defaults to ','.
	Delimiter *string `json:"delimiter,omitempty" tf:"delimiter,omitempty"`

	// Optional. Whether to disable the inference of data type for Json data. If true, all columns will be registered as their primitive types (strings, number or boolean).
	DisableTypeInference *bool `json:"disableTypeInference,omitempty" tf:"disable_type_inference,omitempty"`

	// Optional. The character encoding of the data. The default is UTF-8.
	Encoding *string `json:"encoding,omitempty" tf:"encoding,omitempty"`

	// Optional. The number of rows to interpret as header rows that should be skipped when reading data rows.
	HeaderRows *float64 `json:"headerRows,omitempty" tf:"header_rows,omitempty"`
}

type CsvOptionsParameters struct {

	// Optional. The delimiter being used to separate values. This defaults to ','.
	// +kubebuilder:validation:Optional
	Delimiter *string `json:"delimiter,omitempty" tf:"delimiter,omitempty"`

	// Optional. Whether to disable the inference of data type for Json data. If true, all columns will be registered as their primitive types (strings, number or boolean).
	// +kubebuilder:validation:Optional
	DisableTypeInference *bool `json:"disableTypeInference,omitempty" tf:"disable_type_inference,omitempty"`

	// Optional. The character encoding of the data. The default is UTF-8.
	// +kubebuilder:validation:Optional
	Encoding *string `json:"encoding,omitempty" tf:"encoding,omitempty"`

	// Optional. The number of rows to interpret as header rows that should be skipped when reading data rows.
	// +kubebuilder:validation:Optional
	HeaderRows *float64 `json:"headerRows,omitempty" tf:"header_rows,omitempty"`
}

type DiscoverySpecInitParameters struct {

	// Optional. Configuration for CSV data.
	CsvOptions []CsvOptionsInitParameters `json:"csvOptions,omitempty" tf:"csv_options,omitempty"`

	// Required. Whether discovery is enabled.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Optional. The list of patterns to apply for selecting data to exclude during discovery. For Cloud Storage bucket assets, these are interpreted as glob patterns used to match object names. For BigQuery dataset assets, these are interpreted as patterns to match table names.
	ExcludePatterns []*string `json:"excludePatterns,omitempty" tf:"exclude_patterns,omitempty"`

	// Optional. The list of patterns to apply for selecting data to include during discovery if only a subset of the data should considered. For Cloud Storage bucket assets, these are interpreted as glob patterns used to match object names. For BigQuery dataset assets, these are interpreted as patterns to match table names.
	IncludePatterns []*string `json:"includePatterns,omitempty" tf:"include_patterns,omitempty"`

	// Optional. Configuration for Json data.
	JSONOptions []JSONOptionsInitParameters `json:"jsonOptions,omitempty" tf:"json_options,omitempty"`

	// Optional. Cron schedule (https://en.wikipedia.org/wiki/Cron) for running discovery periodically. Successive discovery runs must be scheduled at least 60 minutes apart. The default value is to run discovery every 60 minutes. To explicitly set a timezone to the cron tab, apply a prefix in the cron tab: "CRON_TZ=${IANA_TIME_ZONE}" or TZ=${IANA_TIME_ZONE}". The ${IANA_TIME_ZONE} may only be a valid string from IANA time zone database. For example, "CRON_TZ=America/New_York 1 * * * *", or "TZ=America/New_York 1 * * * *".
	Schedule *string `json:"schedule,omitempty" tf:"schedule,omitempty"`
}

type DiscoverySpecObservation struct {

	// Optional. Configuration for CSV data.
	CsvOptions []CsvOptionsObservation `json:"csvOptions,omitempty" tf:"csv_options,omitempty"`

	// Required. Whether discovery is enabled.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Optional. The list of patterns to apply for selecting data to exclude during discovery. For Cloud Storage bucket assets, these are interpreted as glob patterns used to match object names. For BigQuery dataset assets, these are interpreted as patterns to match table names.
	ExcludePatterns []*string `json:"excludePatterns,omitempty" tf:"exclude_patterns,omitempty"`

	// Optional. The list of patterns to apply for selecting data to include during discovery if only a subset of the data should considered. For Cloud Storage bucket assets, these are interpreted as glob patterns used to match object names. For BigQuery dataset assets, these are interpreted as patterns to match table names.
	IncludePatterns []*string `json:"includePatterns,omitempty" tf:"include_patterns,omitempty"`

	// Optional. Configuration for Json data.
	JSONOptions []JSONOptionsObservation `json:"jsonOptions,omitempty" tf:"json_options,omitempty"`

	// Optional. Cron schedule (https://en.wikipedia.org/wiki/Cron) for running discovery periodically. Successive discovery runs must be scheduled at least 60 minutes apart. The default value is to run discovery every 60 minutes. To explicitly set a timezone to the cron tab, apply a prefix in the cron tab: "CRON_TZ=${IANA_TIME_ZONE}" or TZ=${IANA_TIME_ZONE}". The ${IANA_TIME_ZONE} may only be a valid string from IANA time zone database. For example, "CRON_TZ=America/New_York 1 * * * *", or "TZ=America/New_York 1 * * * *".
	Schedule *string `json:"schedule,omitempty" tf:"schedule,omitempty"`
}

type DiscoverySpecParameters struct {

	// Optional. Configuration for CSV data.
	// +kubebuilder:validation:Optional
	CsvOptions []CsvOptionsParameters `json:"csvOptions,omitempty" tf:"csv_options,omitempty"`

	// Required. Whether discovery is enabled.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// Optional. The list of patterns to apply for selecting data to exclude during discovery. For Cloud Storage bucket assets, these are interpreted as glob patterns used to match object names. For BigQuery dataset assets, these are interpreted as patterns to match table names.
	// +kubebuilder:validation:Optional
	ExcludePatterns []*string `json:"excludePatterns,omitempty" tf:"exclude_patterns,omitempty"`

	// Optional. The list of patterns to apply for selecting data to include during discovery if only a subset of the data should considered. For Cloud Storage bucket assets, these are interpreted as glob patterns used to match object names. For BigQuery dataset assets, these are interpreted as patterns to match table names.
	// +kubebuilder:validation:Optional
	IncludePatterns []*string `json:"includePatterns,omitempty" tf:"include_patterns,omitempty"`

	// Optional. Configuration for Json data.
	// +kubebuilder:validation:Optional
	JSONOptions []JSONOptionsParameters `json:"jsonOptions,omitempty" tf:"json_options,omitempty"`

	// Optional. Cron schedule (https://en.wikipedia.org/wiki/Cron) for running discovery periodically. Successive discovery runs must be scheduled at least 60 minutes apart. The default value is to run discovery every 60 minutes. To explicitly set a timezone to the cron tab, apply a prefix in the cron tab: "CRON_TZ=${IANA_TIME_ZONE}" or TZ=${IANA_TIME_ZONE}". The ${IANA_TIME_ZONE} may only be a valid string from IANA time zone database. For example, "CRON_TZ=America/New_York 1 * * * *", or "TZ=America/New_York 1 * * * *".
	// +kubebuilder:validation:Optional
	Schedule *string `json:"schedule,omitempty" tf:"schedule,omitempty"`
}

type DiscoveryStatusInitParameters struct {
}

type DiscoveryStatusObservation struct {
	LastRunDuration *string `json:"lastRunDuration,omitempty" tf:"last_run_duration,omitempty"`

	LastRunTime *string `json:"lastRunTime,omitempty" tf:"last_run_time,omitempty"`

	Message *string `json:"message,omitempty" tf:"message,omitempty"`

	// Output only. Current state of the asset. Possible values: STATE_UNSPECIFIED, ACTIVE, CREATING, DELETING, ACTION_REQUIRED
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	Stats []StatsObservation `json:"stats,omitempty" tf:"stats,omitempty"`

	// Output only. The time when the asset was last updated.
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`
}

type DiscoveryStatusParameters struct {
}

type JSONOptionsInitParameters struct {

	// Optional. Whether to disable the inference of data type for Json data. If true, all columns will be registered as their primitive types (strings, number or boolean).
	DisableTypeInference *bool `json:"disableTypeInference,omitempty" tf:"disable_type_inference,omitempty"`

	// Optional. The character encoding of the data. The default is UTF-8.
	Encoding *string `json:"encoding,omitempty" tf:"encoding,omitempty"`
}

type JSONOptionsObservation struct {

	// Optional. Whether to disable the inference of data type for Json data. If true, all columns will be registered as their primitive types (strings, number or boolean).
	DisableTypeInference *bool `json:"disableTypeInference,omitempty" tf:"disable_type_inference,omitempty"`

	// Optional. The character encoding of the data. The default is UTF-8.
	Encoding *string `json:"encoding,omitempty" tf:"encoding,omitempty"`
}

type JSONOptionsParameters struct {

	// Optional. Whether to disable the inference of data type for Json data. If true, all columns will be registered as their primitive types (strings, number or boolean).
	// +kubebuilder:validation:Optional
	DisableTypeInference *bool `json:"disableTypeInference,omitempty" tf:"disable_type_inference,omitempty"`

	// Optional. The character encoding of the data. The default is UTF-8.
	// +kubebuilder:validation:Optional
	Encoding *string `json:"encoding,omitempty" tf:"encoding,omitempty"`
}

type ResourceSpecInitParameters struct {

	// Immutable. Relative name of the cloud resource that contains the data that is being managed within a lake. For example: projects/{project_number}/buckets/{bucket_id} projects/{project_number}/datasets/{dataset_id}
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Required. Immutable. Type of resource. Possible values: STORAGE_BUCKET, BIGQUERY_DATASET
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ResourceSpecObservation struct {

	// Immutable. Relative name of the cloud resource that contains the data that is being managed within a lake. For example: projects/{project_number}/buckets/{bucket_id} projects/{project_number}/datasets/{dataset_id}
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Required. Immutable. Type of resource. Possible values: STORAGE_BUCKET, BIGQUERY_DATASET
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ResourceSpecParameters struct {

	// Immutable. Relative name of the cloud resource that contains the data that is being managed within a lake. For example: projects/{project_number}/buckets/{bucket_id} projects/{project_number}/datasets/{dataset_id}
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Required. Immutable. Type of resource. Possible values: STORAGE_BUCKET, BIGQUERY_DATASET
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type ResourceStatusInitParameters struct {
}

type ResourceStatusObservation struct {
	Message *string `json:"message,omitempty" tf:"message,omitempty"`

	// Output only. Current state of the asset. Possible values: STATE_UNSPECIFIED, ACTIVE, CREATING, DELETING, ACTION_REQUIRED
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// Output only. The time when the asset was last updated.
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`
}

type ResourceStatusParameters struct {
}

type SecurityStatusInitParameters struct {
}

type SecurityStatusObservation struct {
	Message *string `json:"message,omitempty" tf:"message,omitempty"`

	// Output only. Current state of the asset. Possible values: STATE_UNSPECIFIED, ACTIVE, CREATING, DELETING, ACTION_REQUIRED
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// Output only. The time when the asset was last updated.
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`
}

type SecurityStatusParameters struct {
}

type StatsInitParameters struct {
}

type StatsObservation struct {
	DataItems *float64 `json:"dataItems,omitempty" tf:"data_items,omitempty"`

	DataSize *float64 `json:"dataSize,omitempty" tf:"data_size,omitempty"`

	Filesets *float64 `json:"filesets,omitempty" tf:"filesets,omitempty"`

	Tables *float64 `json:"tables,omitempty" tf:"tables,omitempty"`
}

type StatsParameters struct {
}

// AssetSpec defines the desired state of Asset
type AssetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AssetParameters `json:"forProvider"`
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
	InitProvider AssetInitParameters `json:"initProvider,omitempty"`
}

// AssetStatus defines the observed state of Asset.
type AssetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AssetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Asset is the Schema for the Assets API. The Dataplex Asset resource
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Asset struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.discoverySpec) || has(self.initProvider.discoverySpec)",message="discoverySpec is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.resourceSpec) || has(self.initProvider.resourceSpec)",message="resourceSpec is a required parameter"
	Spec   AssetSpec   `json:"spec"`
	Status AssetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AssetList contains a list of Assets
type AssetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Asset `json:"items"`
}

// Repository type metadata.
var (
	Asset_Kind             = "Asset"
	Asset_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Asset_Kind}.String()
	Asset_KindAPIVersion   = Asset_Kind + "." + CRDGroupVersion.String()
	Asset_GroupVersionKind = CRDGroupVersion.WithKind(Asset_Kind)
)

func init() {
	SchemeBuilder.Register(&Asset{}, &AssetList{})
}
