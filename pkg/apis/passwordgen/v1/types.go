package v1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type PasswordType struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PasswordSpec `json:"spec, omitempty"`
}

type PasswordSpec struct {
	Keys []struct {
		KeyName      string `json:"keyName"`
		Length       int    `json:"length"`
		CharacterSet string `json:"characterSet"`
	} `json:"keys"`
	Secrets []struct {
		SecretName string `json:"secretName"`
		Namespace  string `json:"namespace"`
	}
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type PasswordTypeList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `son:"metadata,omitempty"`

	Items []PasswordType `json:"items"`
}
