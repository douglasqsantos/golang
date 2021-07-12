# Golang - Sandbox

Create the directories to store the project

```bash
mkdir -p /Volumes/Data/Git/golang/src
mkdir -p /Volumes/Data/Git/golang/bin
```

Configuring the Path and the GOPATH

```bash
vim ~/.zshrc
[...]
# GO
export GOPATH=/Volumes/Data/Git/golang
export PATH=$PATH:$(go env GOPATH)/bin
```

Creating the directory to start a new module

```bash
mkdir src/pods
go mod init pods
```

Adding the dependencies to initilize the new Module. **Note** Use the correct version in the dependencies, some of them do not work with your current kubernetes version.

```go
vim go.mod 
module pods

go 1.16

require (
	k8s.io/api v0.21.2
	k8s.io/apimachinery v0.21.2
	k8s.io/client-go v0.21.2
)
```

Download the dependencies

```bash
go mod download all
```

## Using

Access the directory with the code so far we have:

- namespaces
- nodes
- pods

```bash
cd src/namespaces
```

Download the modules to be able to use

```bash
go mod download all
```

Now we can run the module

```bash
go run ./main.go
 [ 09/07/2021 - 20:29:04 ][ Auth ] >  [!] Using the kubeconfig [ /Users/douglas/.kube/config ]
 [ 09/07/2021 - 20:29:04 ][ NS ] >  [+] Namespace [default]
 [ 09/07/2021 - 20:29:04 ][ NS ] >  [+] Namespace [kube-node-lease]
 [ 09/07/2021 - 20:29:04 ][ NS ] >  [+] Namespace [kube-public]
 [ 09/07/2021 - 20:29:04 ][ NS ] >  [+] Namespace [kube-system]
```

Running the nodes module

```bash
cd ../nodes
go run ./main.go
 [ 09/07/2021 - 20:29:36 ][ Auth ] >  [!] Using the kubeconfig [ /Users/douglas/.kube/config ]
 [ 09/07/2021 - 20:29:36 ][ Nodes ] >  [!] ############################################
 [ 09/07/2021 - 20:29:36 ][ Nodes ] >  [+] Node Name: [ docker-desktop ]
 [ 09/07/2021 - 20:29:36 ][ Nodes ] >  [+] Node Labels [ beta.kubernetes.io/arch=amd64 ]
 [ 09/07/2021 - 20:29:36 ][ Nodes ] >  [+] Node Labels [ beta.kubernetes.io/os=linux ]
 [ 09/07/2021 - 20:29:36 ][ Nodes ] >  [+] Node Labels [ kubernetes.io/arch=amd64 ]
 [ 09/07/2021 - 20:29:36 ][ Nodes ] >  [+] Node Labels [ kubernetes.io/hostname=docker-desktop ]
 [ 09/07/2021 - 20:29:36 ][ Nodes ] >  [+] Node Labels [ kubernetes.io/os=linux ]
 [ 09/07/2021 - 20:29:36 ][ Nodes ] >  [+] Node Labels [ node-role.kubernetes.io/master= ]
 [ 09/07/2021 - 20:29:36 ][ Nodes ] >  [+] Node Address [ 192.168.65.4 ]
 [ 09/07/2021 - 20:29:36 ][ Nodes ] >  [+] Node Address Type [ InternalIP ]
 [ 09/07/2021 - 20:29:36 ][ Nodes ] >  [+] Node Address [ docker-desktop ]
 [ 09/07/2021 - 20:29:36 ][ Nodes ] >  [+] Node Address Type [ Hostname ]
```

Running the pods module

```bash
cd ../pods
go run ./main.go
 [ 09/07/2021 - 20:30:06 ][ Auth ] >  [!] Using the kubeconfig [ /Users/douglas/.kube/config ]
 [ 09/07/2021 - 20:30:06 ][ Pods ] >  [+] There is [ 1 ] pod in the namespace [ default ]
 [ 09/07/2021 - 20:30:06 ][ Pods ] >  [+] Pod Name [ myapp-pod ]
 [ 09/07/2021 - 20:30:06 ][ Pods ] >  [+] Container Status [ Running ]
 [ 09/07/2021 - 20:30:06 ][ Pods ] >  [+] Pod IP Address [ 10.1.1.71 ]
 [ 09/07/2021 - 20:30:06 ][ Pods ] >  [+] Pod Labels [ app=myapp ]
 [ 09/07/2021 - 20:30:06 ][ Pods ] >  [+] Pod Labels [ type=front-end ]
 [ 09/07/2021 - 20:30:06 ][ Pods ] >  [+] There is [ 1 ] container in the Pod: [ myapp-pod ]
 [ 09/07/2021 - 20:30:06 ][ Pods ] >  [+] Container Name: [ myapp-pod ]
 [ 09/07/2021 - 20:30:06 ][ Pods ] >  [+] Container Image: [ nginx ]
```

## Notes

I spent my all day long figuring out how some features of golang and client-go works and there is a lot of information around the internet talking about how to deal with one or another part of this group of tools and in almost all of them you do not have all the answers that you need to fix the issues. I will try to leave some tips here just to help you guys how to solve some common issues that you can face.

When we try to get the information about a pod e.g

```go
	pods, err := clientset.CoreV1().Pods("default").List(context.TODO(), metav1.ListOptions{})
	errors.HandError(err)

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
```

- The parameters is just like kubectl get pod pod-name -o yaml
- **Note:** The Attributes needs to be capitalize so name will be Name to be able to be accessed as public
- e.g pod.Name to get the pod name
- e.g pod.Status.PodIP to get the pod ip and so on

## Using smaller binary

By default, the produced binary file contains debugging information and the symbol table. This can bloat the size of the file. To reduce the file size, you can include additional flags during the build process to strip this information from the binary. For example, the following command will reduce the binary size by approximately 30 percent:

```bash
go build -ldflags "-w -s"
```

## Accessing API objects with kubectl get

- https://medium.com/programming-kubernetes/building-stuff-with-the-kubernetes-api-1-cc50a3642

## References (pkg.go.dev)

- https://pkg.go.dev/k8s.io/api/core/v1#NodeList
- https://pkg.go.dev/k8s.io/client-go/kubernetes/typed/core/v1#CoreV1Client.Nodes
- https://pkg.go.dev/k8s.io/client-go/tools/clientcmd
- https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#pkg-variables
- https://pkg.go.dev/k8s.io/client-go/tools/clientcmd#pkg-constants
- https://pkg.go.dev/k8s.io/client-go/kubernetes/typed/core/v1#CoreV1Client.Nodes
- https://pkg.go.dev/k8s.io/client-go/kubernetes
- https://pkg.go.dev/search?q=client-go
- https://pkg.go.dev/k8s.io/client-go/kubernetes#Clientset.CoreV1

## References (client-go)

- https://github.com/kubernetes/client-go
- https://github.com/kubernetes/client-go/releases
- https://github.com/kubernetes/client-go/blob/master/examples/out-of-cluster-client-configuration/main.go
- https://github.com/kubernetes/client-go/blob/master/kubernetes/typed/core/v1/pod.go
- https://github.com/kubernetes/client-go/blob/master/kubernetes/typed/core/v1/namespace.go

## References (Kubernetes)

- https://kubernetes.io/docs/reference/kubernetes-api/
- https://kubernetes.io/docs/reference/kubernetes-api/cluster-resources/node-v1/

## References (Examples)

- https://medium.com/programming-kubernetes/building-stuff-with-the-kubernetes-api-part-4-using-go-b1d0e3c1c899
- https://medium.com/@vladimirvivien
- https://github.com/eddiezane/hello-client-go/blob/main/main.go
- https://www.golangprograms.com/assign-default-value-for-struct-field-in-go-programming-language.html
- https://medium.com/programming-kubernetes/building-stuff-with-the-kubernetes-api-toc-84d751876650
- https://searchitoperations.techtarget.com/tutorial/Follow-a-Kubernetes-and-Go-tutorial-in-your-home-lab
- https://golang.github.io/dep/docs/installation.html
- https://medium.com/programming-kubernetes/building-stuff-with-the-kubernetes-api-part-4-using-go-b1d0e3c1c899
- https://medium.com/programming-kubernetes/using-go-modules-with-kubernetes-api-and-client-go-projects-2f3fdd5589a
- https://github.com/vladimirvivien/k8s-client-examples
- https://www.youtube.com/watch?v=jiKwjnlc7Wk&ab_channel=EddieZaneski
- https://medium.com/@adiach3nko/package-management-with-go-modules-the-pragmatic-guide-c831b4eaaf31
- https://gist.github.com/ks888/0a0e0fbf4694d7955999a6f59aa2766d
- https://miminar.fedorapeople.org/_preview/openshift-enterprise/registry-redeploy/go_client/getting_started.html
- https://stackoverflow.com/questions/10105935/how-to-convert-an-int-value-to-string-in-go
- https://www.geeksforgeeks.org/how-to-assign-default-value-for-struct-field-in-golang/
- https://www.geeksforgeeks.org/different-ways-to-find-the-type-of-variable-in-golang/
- https://golangbyexample.com/import-local-module-golang/
- https://blog.golang.org/maps