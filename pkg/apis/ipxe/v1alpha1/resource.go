/*
 * Copyright 2020 Mandelsoft. All rights reserved.
 *  This file is licensed under the Apache Software License, v. 2 except as noted
 *  otherwise in the LICENSE file
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *       http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

package v1alpha1

import (
	"github.com/gardener/controller-manager-library/pkg/types"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type BootResourceList struct {
	metav1.TypeMeta `json:",inline"`
	// Standard list metadata
	// More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BootResource `json:"items"`
}

// +kubebuilder:storageversion
// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced,path=bootresources,shortName=bresc,singular=bootresource
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name=Plain,JSONPath=".spec.plain",priority=2000,type=bool
// +kubebuilder:printcolumn:name=URL,JSONPath=".spec.URL",priority=2000,type=string
// +kubebuilder:printcolumn:name=Redirect,JSONPath=".spec.redirect",priority=2000,type=bool
// +kubebuilder:printcolumn:name=Volatile,JSONPath=".spec.volatile",priority=2000,type=bool
// +kubebuilder:printcolumn:name=ConfigMap,JSONPath=".spec.configMap",priority=2000,type=string
// +kubebuilder:printcolumn:name=Secret,JSONPath=".spec.secret",priority=2000,type=string
// +kubebuilder:printcolumn:name=Field,JSONPath=".spec.fieldName",priority=2000,type=string
// +kubebuilder:printcolumn:name=State,JSONPath=".status.state",type=string
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type BootResource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BootResourceSpec `json:"spec"`
	// +optional
	Status BootResourceStatus `json:"status,omitempty"`
}

type BootResourceSpec struct {
	// +optional
	MimeType string `json:"mimeType,omitempty"`
	// +optional
	Plain *bool `json:"plainContent,omitempty"`
	// +kubebuilder:validation:XPreserveUnknownFields
	// +kubebuilder:pruning:PreserveUnknownFields
	// +optional
	Mapping types.Values `json:"mapping,omitempty"`
	// +kubebuilder:validation:XPreserveUnknownFields
	// +kubebuilder:pruning:PreserveUnknownFields
	// +optional
	Values types.Values `json:"values,omitempty"`
	// +optional
	URL string `json:"URL,omitempty"`
	// +optional
	Volatile bool `json:"volatile,omitempty"`
	// +optional
	Redirect *bool `json:"redirect,omitempty"`
	// +optional
	Text string `json:"text,omitempty"`
	// +optional
	Binary string `json:"binary,omitempty"`

	// +optional
	ConfigMap string `json:"configMap,omitempty"`
	// +optional
	Secret string `json:"secret,omitempty"`
	// +optional
	FieldName string `json:"fieldName,omitempty"`
}

type BootResourceStatus struct {
	// +optional
	State string `json:"state"`

	// +optional
	Message string `json:"message,omitempty"`
}
