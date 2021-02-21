package main

import (
	"fmt"
	"github.com/ycsk02/kubevision/pkg/clustermanager"
	"github.com/ycsk02/kubevision/pkg/handler"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
)

func main() {
	clustermanager.AddToScheme(scheme.Scheme)
	clientset := handler.GetClusterManagerClient()
	kubeConfig, err := clientset.KubeConfig().Get("sukai-test60", metav1.GetOptions{})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(kubeConfig)
}
