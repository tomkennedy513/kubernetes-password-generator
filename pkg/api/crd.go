package api

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

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

type PasswordResource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PasswordSpec `json:"spec"`
}
