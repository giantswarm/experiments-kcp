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
	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"
	"github.com/kcp-dev/logicalcluster/v3"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"

	mountsv1alpha1 "github.com/kcp-dev/kcp/contrib/mounts-vw/apis/mounts/v1alpha1"
)

// VClusterClusterLister can list VClusters across all workspaces, or scope down to a VClusterLister for one workspace.
// All objects returned here must be treated as read-only.
type VClusterClusterLister interface {
	// List lists all VClusters in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*mountsv1alpha1.VCluster, err error)
	// Cluster returns a lister that can list and get VClusters in one workspace.
	Cluster(clusterName logicalcluster.Name) VClusterLister
	VClusterClusterListerExpansion
}

type vClusterClusterLister struct {
	indexer cache.Indexer
}

// NewVClusterClusterLister returns a new VClusterClusterLister.
// We assume that the indexer:
// - is fed by a cross-workspace LIST+WATCH
// - uses kcpcache.MetaClusterNamespaceKeyFunc as the key function
// - has the kcpcache.ClusterIndex as an index
func NewVClusterClusterLister(indexer cache.Indexer) *vClusterClusterLister {
	return &vClusterClusterLister{indexer: indexer}
}

// List lists all VClusters in the indexer across all workspaces.
func (s *vClusterClusterLister) List(selector labels.Selector) (ret []*mountsv1alpha1.VCluster, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*mountsv1alpha1.VCluster))
	})
	return ret, err
}

// Cluster scopes the lister to one workspace, allowing users to list and get VClusters.
func (s *vClusterClusterLister) Cluster(clusterName logicalcluster.Name) VClusterLister {
	return &vClusterLister{indexer: s.indexer, clusterName: clusterName}
}

// VClusterLister can list all VClusters, or get one in particular.
// All objects returned here must be treated as read-only.
type VClusterLister interface {
	// List lists all VClusters in the workspace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*mountsv1alpha1.VCluster, err error)
	// Get retrieves the VCluster from the indexer for a given workspace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*mountsv1alpha1.VCluster, error)
	VClusterListerExpansion
}

// vClusterLister can list all VClusters inside a workspace.
type vClusterLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
}

// List lists all VClusters in the indexer for a workspace.
func (s *vClusterLister) List(selector labels.Selector) (ret []*mountsv1alpha1.VCluster, err error) {
	err = kcpcache.ListAllByCluster(s.indexer, s.clusterName, selector, func(i interface{}) {
		ret = append(ret, i.(*mountsv1alpha1.VCluster))
	})
	return ret, err
}

// Get retrieves the VCluster from the indexer for a given workspace and name.
func (s *vClusterLister) Get(name string) (*mountsv1alpha1.VCluster, error) {
	key := kcpcache.ToClusterAwareKey(s.clusterName.String(), "", name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(mountsv1alpha1.Resource("vclusters"), name)
	}
	return obj.(*mountsv1alpha1.VCluster), nil
}

// NewVClusterLister returns a new VClusterLister.
// We assume that the indexer:
// - is fed by a workspace-scoped LIST+WATCH
// - uses cache.MetaNamespaceKeyFunc as the key function
func NewVClusterLister(indexer cache.Indexer) *vClusterScopedLister {
	return &vClusterScopedLister{indexer: indexer}
}

// vClusterScopedLister can list all VClusters inside a workspace.
type vClusterScopedLister struct {
	indexer cache.Indexer
}

// List lists all VClusters in the indexer for a workspace.
func (s *vClusterScopedLister) List(selector labels.Selector) (ret []*mountsv1alpha1.VCluster, err error) {
	err = cache.ListAll(s.indexer, selector, func(i interface{}) {
		ret = append(ret, i.(*mountsv1alpha1.VCluster))
	})
	return ret, err
}

// Get retrieves the VCluster from the indexer for a given workspace and name.
func (s *vClusterScopedLister) Get(name string) (*mountsv1alpha1.VCluster, error) {
	key := name
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(mountsv1alpha1.Resource("vclusters"), name)
	}
	return obj.(*mountsv1alpha1.VCluster), nil
}