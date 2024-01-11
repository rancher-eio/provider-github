/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type TeamMembershipObservation struct {
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type TeamMembershipParameters struct {

	// The role of the user within the team.
	// Must be one of member or maintainer. Defaults to member.
	// +kubebuilder:validation:Optional
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// The GitHub team id
	// +crossplane:generate:reference:type=github.com/coopnorge/provider-github/apis/team/v1alpha1.Team
	// +kubebuilder:validation:Optional
	TeamID *string `json:"teamId,omitempty" tf:"team_id,omitempty"`

	// Reference to a Team in team to populate teamId.
	// +kubebuilder:validation:Optional
	TeamIDRef *v1.Reference `json:"teamIdRef,omitempty" tf:"-"`

	// Selector for a Team in team to populate teamId.
	// +kubebuilder:validation:Optional
	TeamIDSelector *v1.Selector `json:"teamIdSelector,omitempty" tf:"-"`

	// The user to add to the team.
	// +kubebuilder:validation:Required
	Username *string `json:"username" tf:"username,omitempty"`
}

// TeamMembershipSpec defines the desired state of TeamMembership
type TeamMembershipSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TeamMembershipParameters `json:"forProvider"`
}

// TeamMembershipStatus defines the observed state of TeamMembership.
type TeamMembershipStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TeamMembershipObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TeamMembership is the Schema for the TeamMemberships API. Provides a GitHub team membership resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,github}
type TeamMembership struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TeamMembershipSpec   `json:"spec"`
	Status            TeamMembershipStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TeamMembershipList contains a list of TeamMemberships
type TeamMembershipList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TeamMembership `json:"items"`
}

// Repository type metadata.
var (
	TeamMembership_Kind             = "TeamMembership"
	TeamMembership_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TeamMembership_Kind}.String()
	TeamMembership_KindAPIVersion   = TeamMembership_Kind + "." + CRDGroupVersion.String()
	TeamMembership_GroupVersionKind = CRDGroupVersion.WithKind(TeamMembership_Kind)
)

func init() {
	SchemeBuilder.Register(&TeamMembership{}, &TeamMembershipList{})
}