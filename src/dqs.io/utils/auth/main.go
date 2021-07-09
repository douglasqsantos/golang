package auth

import (
	"flag"
	"fmt"
	"path/filepath"

	"dqs.io/utils/errors"
	"dqs.io/utils/msgs"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func Auth() *kubernetes.Clientset {

	// Setting up the kubeconfig
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}

	flag.Parse()

	// Showing information about the kubeconfig
	message := fmt.Sprintf("Using the kubeconfig %s\n", msgs.Focus(*kubeconfig))
	msgs.Warn("Auth", message)

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	errors.HandError(err)

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	errors.HandError(err)

	// returning the clientset
	return clientset

}
