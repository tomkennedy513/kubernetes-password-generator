package reconciler

import (
	"context"
	"github.com/go-logr/logr"
	passwordgenv1 "github.com/tomkennedy513/password-gen/pkg/apis/passwordgen/v1"
	"github.com/tomkennedy513/password-gen/pkg/password"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

type ReplicaSetReconciler struct {
	client.Client
	*runtime.Scheme
	logr.Logger
}

func (r *ReplicaSetReconciler) Reconcile(req reconcile.Request) (reconcile.Result, error) {
	var pw passwordgenv1.Password
	err := r.Get(context.TODO(), req.NamespacedName, &pw)
	if err != nil {
		return reconcile.Result{}, err
	}

	spec := pw.Spec
	result := map[string]string{}

	config, err := password.NewPasswordGenerationConfig(
		password.SetPasswordLength(spec.GenerationParameters.Length),
		password.SetCharacterSet([]rune(spec.GenerationParameters.CharacterSet)),
	)
	if err != nil {
		return reconcile.Result{}, err
	}

	generator := password.NewPasswordGenerator(config)
	generatedPassword, err := generator.Generate()
	if err != nil {
		return reconcile.Result{}, err
	}

	result["password"] = generatedPassword

	username := spec.Username
	if username == "" {
		username, err = generator.Generate()
		if err != nil {
			return reconcile.Result{}, err
		}
	}
	result["username"] = username

	for _, secret := range spec.Secrets {
		namespace := "default"
		if secret.Namespace != "" {
			namespace = secret.Namespace
		}
		err := r.CreateOrUpdateSecret(types.NamespacedName{Namespace: namespace, Name: secret.SecretName}, result)
		if err != nil {
			return reconcile.Result{}, err
		}
	}

	return reconcile.Result{}, nil
}

func (r *ReplicaSetReconciler) InjectClient(c client.Client) error {
	r.Client = c
	return nil
}

func (r *ReplicaSetReconciler) CreateOrUpdateSecret(name types.NamespacedName, credential map[string]string) error {
	secret, err := r.getSecret(name)
	if err != nil {
		if errors.IsNotFound(err) {
			return r.createSecret(name, credential)
		} else {
			return err
		}
	}

	return r.updateSecret(secret, credential)
}

func (r *ReplicaSetReconciler) createSecret(name types.NamespacedName, credential map[string]string) error {
	secret := corev1.Secret{
		TypeMeta: v1.TypeMeta{
			Kind: "Secret",
		},
		ObjectMeta: v1.ObjectMeta{
			Name: name.Name,
			Annotations: map[string]string{
				"generated": "true",
			},
			Namespace: name.Namespace,
		},
		StringData: credential,
		Type:       "kubernetes.io/basic-auth",
	}

	err := r.Client.Create(context.Background(), &secret, &client.CreateOptions{})
	if err != nil {
		return err
	}

	r.Logger.Info("successfully create secret", name)
	return nil
}

func (r *ReplicaSetReconciler) updateSecret(old *corev1.Secret, credential map[string]string) error {
	old.Data = map[string][]byte{}
	old.StringData = credential

	err := r.Client.Update(context.Background(), old, &client.UpdateOptions{})
	if err != nil {
		return err
	}

	r.Logger.Info("successfully updated secret", map[string]string{old.Namespace: old.Name})

	return nil
}

func (r *ReplicaSetReconciler) getSecret(name types.NamespacedName) (*corev1.Secret, error) {
	var secret corev1.Secret
	err := r.Client.Get(context.Background(), name, &secret)
	if err != nil {
		return nil, err
	}

	return &secret, nil
}
