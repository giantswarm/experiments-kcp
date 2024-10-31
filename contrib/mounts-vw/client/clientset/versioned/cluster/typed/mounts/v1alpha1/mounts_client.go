//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The KCP Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by kcp code-generator. DO NOT EDIT.

package v1alpha1

import (
	"net/http"

	kcpclient "github.com/kcp-dev/apimachinery/v2/pkg/client"
	"github.com/kcp-dev/logicalcluster/v3"

	"k8s.io/client-go/rest"

	mountsv1alpha1 "github.com/kcp-dev/kcp/contrib/mounts-vw/client/clientset/versioned/typed/mounts/v1alpha1"
)

type MountsV1alpha1ClusterInterface interface {
	MountsV1alpha1ClusterScoper
	KubeClustersClusterGetter
	VClustersClusterGetter
}

type MountsV1alpha1ClusterScoper interface {
	Cluster(logicalcluster.Path) mountsv1alpha1.MountsV1alpha1Interface
}

type MountsV1alpha1ClusterClient struct {
	clientCache kcpclient.Cache[*mountsv1alpha1.MountsV1alpha1Client]
}

func (c *MountsV1alpha1ClusterClient) Cluster(clusterPath logicalcluster.Path) mountsv1alpha1.MountsV1alpha1Interface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}
	return c.clientCache.ClusterOrDie(clusterPath)
}

func (c *MountsV1alpha1ClusterClient) KubeClusters() KubeClusterClusterInterface {
	return &kubeClustersClusterInterface{clientCache: c.clientCache}
}

func (c *MountsV1alpha1ClusterClient) VClusters() VClusterClusterInterface {
	return &vClustersClusterInterface{clientCache: c.clientCache}
}

// NewForConfig creates a new MountsV1alpha1ClusterClient for the given config.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*MountsV1alpha1ClusterClient, error) {
	client, err := rest.HTTPClientFor(c)
	if err != nil {
		return nil, err
	}
	return NewForConfigAndClient(c, client)
}

// NewForConfigAndClient creates a new MountsV1alpha1ClusterClient for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *rest.Config, h *http.Client) (*MountsV1alpha1ClusterClient, error) {
	cache := kcpclient.NewCache(c, h, &kcpclient.Constructor[*mountsv1alpha1.MountsV1alpha1Client]{
		NewForConfigAndClient: mountsv1alpha1.NewForConfigAndClient,
	})
	if _, err := cache.Cluster(logicalcluster.Name("root").Path()); err != nil {
		return nil, err
	}
	return &MountsV1alpha1ClusterClient{clientCache: cache}, nil
}

// NewForConfigOrDie creates a new MountsV1alpha1ClusterClient for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *MountsV1alpha1ClusterClient {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}
