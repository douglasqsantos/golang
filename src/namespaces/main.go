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

	ns, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	errors.HandError(err)
	for _, n := range ns.Items {
		message := fmt.Sprintf("Namespace [%s]\n", n.Name)
		msgs.Ok("NS", message)
	}
	fmt.Println()

}
