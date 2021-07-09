package main

import (
	"context"
	"fmt"

	"dqs.io/utils/auth"
	"dqs.io/utils/errors"
	"dqs.io/utils/msgs"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {

	clientset := auth.Auth()

	nodeList, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	errors.HandError(err)
	for _, n := range nodeList.Items {
		msgs.Warn("Nodes", "############################################\n")
		message := fmt.Sprintf("Node Name: %s\n", msgs.Focus(n.Name))
		msgs.Ok("Nodes", message)
		for k, v := range n.Labels {
			labels := fmt.Sprintf("%s=%s", k, v)
			msg := fmt.Sprintf("Node Labels %s\n", msgs.Focus(labels))
			msgs.Ok("Nodes", msg)
		}
		for _, n := range n.Status.Addresses {
			addr := fmt.Sprintf("Node Address %s\n", msgs.Focus(n.Address))
			msgs.Ok("Nodes", addr)
			atype := fmt.Sprintf("Node Address Type %s\n", msgs.Focus(string(n.Type)))
			msgs.Ok("Nodes", atype)
		}

	}
}
