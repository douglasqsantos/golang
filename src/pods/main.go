package main

import (
	"context"
	"fmt"
	"strconv"

	"dqs.io/utils/auth"
	"dqs.io/utils/errors"
	"dqs.io/utils/msgs"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {

	clientset := auth.Auth()

	// will list all the pods in the namespace default
	// the parameters is just like oc get pod pod-name -o yaml
	// Note: the Attributes needs to be capitalize so name will be Name
	// e.g pod.Name to get the pod name
	// e.g pod.Status.PodIP to get the pod ip and so on
	pods, err := clientset.CoreV1().Pods("default").List(context.TODO(), metav1.ListOptions{})
	errors.HandError(err)

	if len(pods.Items) == 0 {
		msg := fmt.Sprintf("There is no pod in the namespace %s\n", msgs.Focus("default"))
		msgs.Warn("Pods", msg)
	} else if len(pods.Items) == 1 {
		msg := fmt.Sprintf("There is %s pod in the namespace %s\n", msgs.Focus("1"), msgs.Focus("default"))
		msgs.Ok("Pods", msg)
	} else {
		msg := fmt.Sprintf("There are %s pods in the namespace %s\n", msgs.Focus("1"), msgs.Focus("default"))
		msgs.Ok("Pods", msg)
	}

	for _, n := range pods.Items {
		msg := fmt.Sprintf("Pod Name %s\n", msgs.Focus(n.Name))
		msgs.Ok("Pods", msg)
		msg = fmt.Sprintf("Container Status %s\n", msgs.Focus(string(n.Status.Phase)))
		msgs.Ok("Pods", msg)

		msg = fmt.Sprintf("Pod IP Address %s\n", msgs.Focus(n.Status.PodIP))
		msgs.Ok("Pods", msg)
		for k, v := range n.Labels {
			labels := fmt.Sprintf("%s=%s", k, v)
			msg = fmt.Sprintf("Pod Labels %s\n", msgs.Focus(labels))
			msgs.Ok("Pods", msg)
		}

		if len(n.Spec.Containers) == 1 {
			// Convert int to string
			qtd := strconv.Itoa(len(n.Spec.Containers))
			msg = fmt.Sprintf("There is %s container in the Pod: %s\n", msgs.Focus(qtd), msgs.Focus(n.Name))
			msgs.Ok("Pods", msg)
		} else {
			// Convert int to string
			qtd := strconv.Itoa(len(n.Spec.Containers))
			msg = fmt.Sprintf("There are %s containers in the Pod: %s\n", msgs.Focus(qtd), msgs.Focus(n.Name))
			msgs.Ok("Pods", msg)
		}

		for _, c := range n.Spec.Containers {
			msg := fmt.Sprintf("Container Name: %s\n", msgs.Focus(c.Name))
			msgs.Ok("Pods", msg)
			msg = fmt.Sprintf("Container Image: %s\n", msgs.Focus(c.Image))
			msgs.Ok("Pods", msg)
		}

	}
}
