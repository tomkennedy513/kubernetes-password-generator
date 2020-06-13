package main

import (
	passwordgenv1 "github.com/tomkennedy513/password-gen/pkg/apis/passwordgen/v1"
	"github.com/tomkennedy513/password-gen/pkg/reconciler"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/clientcmd/api"
	"os"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/manager/signals"
)

func main() {
	logf.SetLogger(zap.New())

	var log = logf.Log.WithName("builder-examples")

	err := passwordgenv1.AddToScheme(scheme.Scheme)
	if err != nil {
		log.Error(err, "could not add to scheme")
		os.Exit(1)
	}

	mgrConfig := config.GetConfigOrDie()
	mgrOptions := manager.Options{
	}

	mgr, err := manager.New(mgrConfig, mgrOptions)
	if err != nil {
		log.Error(err, "could not create manager")
		os.Exit(1)
	}

	err = api.AddToScheme(mgr.GetScheme())
	if err != nil {
		log.Error(err, "unable to add scheme")
		os.Exit(1)
	}

	err = builder.
		ControllerManagedBy(mgr).
		For(&passwordgenv1.Password{}).
		Owns(&corev1.Secret{}). // ReplicaSet owns Pods created by it
		Complete(&reconciler.ReplicaSetReconciler{
			Client: mgr.GetClient(),
			Scheme: mgr.GetScheme(),
			Logger: log,
		})
	if err != nil {
		log.Error(err, "could not create controller")
		os.Exit(1)
	}

	if err := mgr.Start(signals.SetupSignalHandler()); err != nil {
		log.Error(err, "could not start manager")
		os.Exit(1)
	}
}

