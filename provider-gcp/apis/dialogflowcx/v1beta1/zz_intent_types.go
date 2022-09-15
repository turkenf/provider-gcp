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

type IntentObservation struct {

	// The unique identifier of the training phrase.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The unique identifier of the intent.
	// Format: projects//locations//agents//intents/.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The collection of training phrases the agent is trained on to identify the intent.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	TrainingPhrases []TrainingPhrasesObservation `json:"trainingPhrases,omitempty" tf:"training_phrases,omitempty"`
}

type IntentParameters struct {

	// Human readable description for better understanding an intent like its scope, content, result etc. Maximum character limit: 140 characters.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The human-readable name of the intent, unique within the agent.
	// +kubebuilder:validation:Required
	DisplayName *string `json:"displayName" tf:"display_name,omitempty"`

	// Indicates whether this is a fallback intent. Currently only default fallback intent is allowed in the agent, which is added upon agent creation.
	// Adding training phrases to fallback intent is useful in the case of requests that are mistakenly matched, since training phrases assigned to fallback intents act as negative examples that triggers no-match event.
	// +kubebuilder:validation:Optional
	IsFallback *bool `json:"isFallback,omitempty" tf:"is_fallback,omitempty"`

	// The key/value metadata to label an intent. Labels can contain lowercase letters, digits and the symbols '-' and '_'. International characters are allowed, including letters from unicase alphabets. Keys must start with a letter. Keys and values can be no longer than 63 characters and no more than 128 bytes.
	// Prefix "sys-" is reserved for Dialogflow defined labels. Currently allowed Dialogflow defined labels include: * sys-head * sys-contextual The above labels do not require value. "sys-head" means the intent is a head intent. "sys.contextual" means the intent is a contextual intent.
	// An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The language of the following fields in intent:
	// Intent.training_phrases.parts.text
	// If not specified, the agent's default language is used. Many languages are supported. Note: languages must be enabled in the agent before they can be used.
	// +kubebuilder:validation:Optional
	LanguageCode *string `json:"languageCode,omitempty" tf:"language_code,omitempty"`

	// The collection of parameters associated with the intent.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Parameters []ParametersParameters `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// The agent to create an intent for.
	// Format: projects//locations//agents/.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-gcp/apis/dialogflowcx/v1beta1.Agent
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Parent *string `json:"parent,omitempty" tf:"parent,omitempty"`

	// Reference to a Agent in dialogflowcx to populate parent.
	// +kubebuilder:validation:Optional
	ParentRef *v1.Reference `json:"parentRef,omitempty" tf:"-"`

	// Selector for a Agent in dialogflowcx to populate parent.
	// +kubebuilder:validation:Optional
	ParentSelector *v1.Selector `json:"parentSelector,omitempty" tf:"-"`

	// The priority of this intent. Higher numbers represent higher priorities.
	// If the supplied value is unspecified or 0, the service translates the value to 500,000, which corresponds to the Normal priority in the console.
	// If the supplied value is negative, the intent is ignored in runtime detect intent requests.
	// +kubebuilder:validation:Optional
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// The collection of training phrases the agent is trained on to identify the intent.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	TrainingPhrases []TrainingPhrasesParameters `json:"trainingPhrases,omitempty" tf:"training_phrases,omitempty"`
}

type ParametersObservation struct {
}

type ParametersParameters struct {

	// The entity type of the parameter.
	// Format: projects/-/locations/-/agents/-/entityTypes/ for system entity types (for example, projects/-/locations/-/agents/-/entityTypes/sys.date), or projects//locations//agents//entityTypes/ for developer entity types.
	// +kubebuilder:validation:Required
	EntityType *string `json:"entityType" tf:"entity_type,omitempty"`

	// The unique identifier of the training phrase.
	// +kubebuilder:validation:Required
	ID *string `json:"id" tf:"id,omitempty"`

	// Indicates whether the parameter represents a list of values.
	// +kubebuilder:validation:Optional
	IsList *bool `json:"isList,omitempty" tf:"is_list,omitempty"`

	// Indicates whether the parameter content should be redacted in log. If redaction is enabled, the parameter content will be replaced by parameter name during logging.
	// Note: the parameter content is subject to redaction if either parameter level redaction or entity type level redaction is enabled.
	// +kubebuilder:validation:Optional
	Redact *bool `json:"redact,omitempty" tf:"redact,omitempty"`
}

type PartsObservation struct {
}

type PartsParameters struct {

	// The parameter used to annotate this part of the training phrase. This field is required for annotated parts of the training phrase.
	// +kubebuilder:validation:Optional
	ParameterID *string `json:"parameterId,omitempty" tf:"parameter_id,omitempty"`

	// The text for this part.
	// +kubebuilder:validation:Required
	Text *string `json:"text" tf:"text,omitempty"`
}

type TrainingPhrasesObservation struct {

	// The unique identifier of the training phrase.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type TrainingPhrasesParameters struct {

	// The ordered list of training phrase parts. The parts are concatenated in order to form the training phrase.
	// Note: The API does not automatically annotate training phrases like the Dialogflow Console does.
	// Note: Do not forget to include whitespace at part boundaries, so the training phrase is well formatted when the parts are concatenated.
	// If the training phrase does not need to be annotated with parameters, you just need a single part with only the Part.text field set.
	// If you want to annotate the training phrase, you must create multiple parts, where the fields of each part are populated in one of two ways:
	// Part.text is set to a part of the phrase that has no parameters.
	// Part.text is set to a part of the phrase that you want to annotate, and the parameterId field is set.
	// Structure is documented below.
	// +kubebuilder:validation:Required
	Parts []PartsParameters `json:"parts" tf:"parts,omitempty"`

	// Indicates how many times this example was added to the intent.
	// +kubebuilder:validation:Optional
	RepeatCount *float64 `json:"repeatCount,omitempty" tf:"repeat_count,omitempty"`
}

// IntentSpec defines the desired state of Intent
type IntentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IntentParameters `json:"forProvider"`
}

// IntentStatus defines the observed state of Intent.
type IntentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IntentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Intent is the Schema for the Intents API. An intent represents a user's intent to interact with a conversational agent.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Intent struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IntentSpec   `json:"spec"`
	Status            IntentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IntentList contains a list of Intents
type IntentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Intent `json:"items"`
}

// Repository type metadata.
var (
	Intent_Kind             = "Intent"
	Intent_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Intent_Kind}.String()
	Intent_KindAPIVersion   = Intent_Kind + "." + CRDGroupVersion.String()
	Intent_GroupVersionKind = CRDGroupVersion.WithKind(Intent_Kind)
)

func init() {
	SchemeBuilder.Register(&Intent{}, &IntentList{})
}
