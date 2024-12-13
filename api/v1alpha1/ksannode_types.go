/*
Copyright 2024.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type KSANNodeStorage struct {
	VolumeGroupName string   `json:"volumeGroupName,omitempty"`
	Devices         []string `json:"devices,omitempty"`
}

// KSANNodeSpec defines the desired state of KSANNode.
type KSANNodeSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	HostID  int               `json:"hostID,omitempty"`
	Storage []KSANNodeStorage `json:"storage,omitempty"`
}

// KSANNodeStatus defines the observed state of KSANNode.
type KSANNodeStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// KSANNode is the Schema for the ksannodes API.
type KSANNode struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KSANNodeSpec   `json:"spec,omitempty"`
	Status KSANNodeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KSANNodeList contains a list of KSANNode.
type KSANNodeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KSANNode `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KSANNode{}, &KSANNodeList{})
}
