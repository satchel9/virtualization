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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DiskSnapshot is the Schema for the disksnapshots API
// +k8s:openapi-gen=true
type DiskSnapshot struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DiskSnapshotSpec   `json:"spec"`
	Status DiskSnapshotStatus `json:"status,omitempty"`
}

// DiskSnapshotSpec defines the desired state of DiskSnapshot
type DiskSnapshotSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	VolumeSnapshotName string                 `json:"volumeSnapshotName"`
	Source             DiskSnapshotSpecSource `json:"source"`
}

type DiskSnapshotSpecSource struct {
	APIGroup string `json:"apiGroup"`
	Kind     string `json:"kind"`
	Name     string `json:"name"`
}

// DiskSnapshotStatus defines the observed state of DiskSnapshot
type DiskSnapshotStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	AccessMode   string `json:"accessMode"`
	ReadyToUse   bool   `json:"readyToUse"`
	StorageClass string `json:"storageClass"`
	RestoreSize  string `json:"restoreSize"`
	VolumeMode   string `json:"volumeMode"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DiskVolumeList contains a list of DiskSnapshot
type DiskSnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DiskSnapshot `json:"items"`
}
