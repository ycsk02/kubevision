package clustermanager

import (
	"context"
	clustermanagerv1 "github.com/ycsk02/cluster-manager/api/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

type KubeConfigInterface interface {
	List(opts metav1.ListOptions) (*clustermanagerv1.KubeConfigList, error)
	Get(name string, options metav1.GetOptions) (*clustermanagerv1.KubeConfig, error)
	Create(*clustermanagerv1.KubeConfig) (*clustermanagerv1.KubeConfig, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	// ...
}

type ClusterManagerV1Client struct {
	restClient	rest.Interface
}

func (c *ClusterManagerV1Client) List(opts metav1.ListOptions) (*clustermanagerv1.KubeConfigList, error) {
	result := clustermanagerv1.KubeConfigList{}
	err := c.restClient.
		Get().
		Resource("kubeconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(context.Background()).
		Into(&result)

	return &result, err
}

func (c *ClusterManagerV1Client) Get(name string, opts metav1.GetOptions) (*clustermanagerv1.KubeConfig, error) {
	result := clustermanagerv1.KubeConfig{}
	err := c.restClient.
		Get().
		Resource("kubeconfigs").
		Name(name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(context.Background()).
		Into(&result)

	return &result, err
}

func (c *ClusterManagerV1Client) Create(project *clustermanagerv1.KubeConfig) (*clustermanagerv1.KubeConfig, error) {
	result := clustermanagerv1.KubeConfig{}
	err := c.restClient.
		Post().
		Resource("kubeconfigs").
		Body(project).
		Do(context.Background()).
		Into(&result)

	return &result, err
}

func (c *ClusterManagerV1Client) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.restClient.
		Get().
		Resource("kubeconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch(context.Background())
}