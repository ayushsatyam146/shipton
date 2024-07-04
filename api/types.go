package api

import (
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ShiptonBuildSpec struct {
	Image      string `json:"image,omitempty"`
	Dockerfile string `json:"dockerfile,omitempty"`
	Context    string `json:"context,omitempty"`
}

type ShiptonBuildStatus struct {
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

type ShiptonBuild struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ShiptonBuildSpec   `json:"spec,omitempty"`
	Status ShiptonBuildStatus `json:"status,omitempty"`
}

type ShiptonBuildList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ShiptonBuild `json:"items"`
}