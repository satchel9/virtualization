/*
Copyright 2023.

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

package v1

import (
        "k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DiskVolume is the Schema for the diskvolumes API
// +k8s:openapi-gen=true
type DiskVolume struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DiskVolumeSpec   `json:"spec"`
	Status DiskVolumeStatus `json:"status,omitempty"`
}


// DiskRestoreSpec defines the desired state of DiskRestore
type DiskVolumeSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	PVCName      string                  `json:"pvcName"`
	Resources    DiskVolumeSpecResources `json:"resources"`
	Source       DiskVolumeSpecSource    `json:"source"`
	StorageClass string                  `json:"storageClass"`
	VolumeMode   string                  `json:"volumeMode"`
	AccessMode   string                  `json:"accessMode"`
}

const (
	ResourceStorage string = "storage"
)

type ResourceList map[string]resource.Quantity

type DiskVolumeSpecResources struct {
	Limits   ResourceList `json:"limits,omitempty"`
	Requests ResourceList `json:"requests"`
}

type DiskVolumeSpecSource struct {
	Image DiskVolumeSpecImage `json:"image"`
}

type DiskVolumeSpecImage struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

// DiskVolumeStatus defines the observed state of DiskVolume
type DiskVolumeStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	Created      bool   `json:"created"`
	Ready        bool   `json:"ready"`
	StorageClass string `json:"storageClass"`
	Target       string `json:"target"`
	VolumeMode   string `json:"volumeMode"`
	Owner        string `json:"owner"`
	AccessMode   string `json:"accessMode"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DiskVolumeList contains a list of DiskRestore
type DiskVolumeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DiskVolume `json:"items"`
}
