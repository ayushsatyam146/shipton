package api

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// ShiptonBuildSpec defines the desired state of ShiptonBuild
type ShiptonBuildSpec struct {
	Image      string `json:"image,omitempty"`
	Dockerfile string `json:"dockerfile,omitempty"`
	Context    string `json:"context,omitempty"`
}

// ShiptonBuildStatus defines the observed state of ShiptonBuild
type ShiptonBuildStatus struct {
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// ShiptonBuild is the Schema for the shiptonbuilds API
type ShiptonBuild struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ShiptonBuildSpec   `json:"spec,omitempty"`
	Status ShiptonBuildStatus `json:"status,omitempty"`
}

// DeepCopyObject implements the runtime.Object interface
func (in *ShiptonBuild) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto copies the receiver and writes into out. in must be non-nil.
func (in *ShiptonBuild) DeepCopyInto(out *ShiptonBuild) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ObjectMeta = *in.ObjectMeta.DeepCopy()
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy creates a new DeepCopy of the receiver, creating a new ShiptonBuild.
func (in *ShiptonBuild) DeepCopy() *ShiptonBuild {
	if in == nil {
		return nil
	}
	out := new(ShiptonBuild)
	in.DeepCopyInto(out)
	return out
}

// ShiptonBuildList contains a list of ShiptonBuild
type ShiptonBuildList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ShiptonBuild `json:"items"`
}

// DeepCopyObject implements the runtime.Object interface
func (in *ShiptonBuildList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto copies the receiver and writes into out. in must be non-nil.
func (in *ShiptonBuildList) DeepCopyInto(out *ShiptonBuildList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = *in.ListMeta.DeepCopy()
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ShiptonBuild, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy creates a new DeepCopy of the receiver, creating a new ShiptonBuildList.
func (in *ShiptonBuildList) DeepCopy() *ShiptonBuildList {
	if in == nil {
		return nil
	}
	out := new(ShiptonBuildList)
	in.DeepCopyInto(out)
	return out
}
