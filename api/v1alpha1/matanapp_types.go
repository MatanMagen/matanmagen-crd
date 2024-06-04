/*
Copyright 2024 MatanMagen.

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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MatanAppSpec defines the desired state of MatanApp
type MatanAppSpec struct {
	// Secret defines the secret spec
	Secret corev1.Secret `json:"secret"`

	// Job defines the job spec
	ImageName string `json:"imageName"`

	// ConfigMap defines the configmap spec
	ConfigMap corev1.ConfigMap `json:"configMap"`
}

// MatanAppStatus defines the observed state of MatanApp
type MatanAppStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// MatanApp is the Schema for the matanapps API
type MatanApp struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MatanAppSpec   `json:"spec,omitempty"`
	Status MatanAppStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MatanAppList contains a list of MatanApp
type MatanAppList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MatanApp `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MatanApp{}, &MatanAppList{})
}
