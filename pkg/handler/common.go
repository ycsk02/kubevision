package handler

import (
	"flag"
	"github.com/ycsk02/kubevision/pkg/clustermanager"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"os"
	"path/filepath"
)

type ObjectClient struct {
	restClient rest.Interface
	resource   *metav1.APIResource
	gvk        schema.GroupVersionKind
	ns         string
}

type Store struct {
	*kubernetes.Clientset
}

func GetControllerClusterConfig() *rest.Config {
	var (
		config     *rest.Config
		kubeconfig *string
		err        error
	)
	if _, inCluster := os.LookupEnv("KUBERNETES_SERVICE_HOST"); inCluster == true {
		config, err = rest.InClusterConfig()
		if err != nil {
			panic(err.Error())
		}
	}

	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err = clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	return config
}

func GetKubeClient() *kubernetes.Clientset {
	config := GetControllerClusterConfig()
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return clientset
}

func GetClusterManagerClient() *clustermanager.ClusterManagerV1Client {
	config := GetControllerClusterConfig()
	clientset, err := clustermanager.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return clientset
}

//
// func NewStore() *Store{
// 	return &Store{
// 		GetKubeClient(),
// 	}
// }
//
// func (s *Store) GetWorkerClusterConfig(clusterName string) {
// }
