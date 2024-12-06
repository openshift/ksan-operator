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
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type DevicesSelectionMode string

const (
	DevicesSelectionModeAll      DevicesSelectionMode = "AllDevices"
	DevicesSelectionModeSelected DevicesSelectionMode = "SelectedDevices"
)

type DevicesConfig struct {
	// +kubebuilder:validation:Enum="AllDevices;SelectedDevices"
	Mode DevicesSelectionMode `json:"mode"`
	// +kubebuilder:validation:MinItems:=1
	// +kubebuilder:validation:UniqueItems:=true
	Devices []string `json:"devices"`
}

// KSANStorageSpec defines the desired state of KSANStorage
type KSANStorageSpec struct {
	// +kubebuilder:default:="KsanCluster"
	Name string `json:"name"`

	DevicesConfig DevicesConfig `json:"devicesConfig"`

	// +kubebuilder:validation:Optional
	KubesanParams map[string]string `json:"kubesanParams"`

	// +kubebuilder:validation:Optional
	Affinity v1.Affinity `json:"affinity,omitempty"`
}

// KSANStorageStatus defines the observed state of KSANStorage
type KSANStorageStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// KSANStorage is the Schema for the ksanstorages API
type KSANStorage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KSANStorageSpec   `json:"spec,omitempty"`
	Status KSANStorageStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KSANStorageList contains a list of KSANStorage
type KSANStorageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KSANStorage `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KSANStorage{}, &KSANStorageList{})
}
