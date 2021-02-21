package clustermanager

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

type ClusterManagerV1Interface interface {
	KubeConfig() KubeConfigInterface
}

type ClusterManagerV1 struct {
	restClient rest.Interface
}

func NewForConfig(c *rest.Config) (*ClusterManagerV1Client, error) {
	config := *c
	config.ContentConfig.GroupVersion = &schema.GroupVersion{Group: "manager.sukai.io", Version: "v1"}
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()
	config.UserAgent = rest.DefaultKubernetesUserAgent()

	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}

	return &ClusterManagerV1Client{restClient: client}, nil
}

func (c *ClusterManagerV1Client) KubeConfig() KubeConfigInterface {
	return &ClusterManagerV1Client{
		restClient: c.restClient,
	}
}