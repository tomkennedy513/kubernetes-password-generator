package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Password struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PasswordSpec `json:"spec, omitempty"`
}

type Key struct {
	KeyName      string `json:"keyName"`
	Length       int    `json:"length"`
	CharacterSet string `json:"characterSet"`
}

type Secret struct {
	SecretName string `json:"secretName"`
	Namespace  string `json:"namespace"`
}

type PasswordSpec struct {
	Keys []Key `json:"keys"`
	Secrets []Secret `json:"secrets"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type PasswordList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `son:"metadata,omitempty"`

	Items []Password `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Password{}, &PasswordList{})
}