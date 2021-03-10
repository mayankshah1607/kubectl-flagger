package k8s

import (
	"fmt"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// Client holds the K8s client set
var Client *kubernetes.Clientset

// RestConfig holds the rest config
var RestConfig *rest.Config

func initNewClient() error {
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	configOverrides := &clientcmd.ConfigOverrides{}
	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides)

	config, err := kubeConfig.ClientConfig()
	if err != nil {
		return err
	}
	RestConfig = config

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}

	Client = clientset
	return nil
}

func init() {
	err := initNewClient()
	if err != nil {
		panic(fmt.Sprintf("failed to initialize k8s client: %s", err.Error()))
	}
}
