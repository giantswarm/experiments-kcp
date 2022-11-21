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
	kcpcache "github.com/kcp-dev/apimachinery/pkg/cache"
	"github.com/kcp-dev/logicalcluster/v2"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"

	apisv1alpha1 "github.com/kcp-dev/kcp/pkg/apis/apis/v1alpha1"
)

// APIBindingClusterLister can list APIBindings across all workspaces, or scope down to a APIBindingLister for one workspace.
// All objects returned here must be treated as read-only.
type APIBindingClusterLister interface {
	// List lists all APIBindings in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*apisv1alpha1.APIBinding, err error)
	// Cluster returns a lister that can list and get APIBindings in one workspace.
	Cluster(cluster logicalcluster.Name) APIBindingLister
	APIBindingClusterListerExpansion
}

type aPIBindingClusterLister struct {
	indexer cache.Indexer
}

// NewAPIBindingClusterLister returns a new APIBindingClusterLister.
// We assume that the indexer:
// - is fed by a cross-workspace LIST+WATCH
// - uses kcpcache.MetaClusterNamespaceKeyFunc as the key function
// - has the kcpcache.ClusterIndex as an index
func NewAPIBindingClusterLister(indexer cache.Indexer) *aPIBindingClusterLister {
	return &aPIBindingClusterLister{indexer: indexer}
}

// List lists all APIBindings in the indexer across all workspaces.
func (s *aPIBindingClusterLister) List(selector labels.Selector) (ret []*apisv1alpha1.APIBinding, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*apisv1alpha1.APIBinding))
	})
	return ret, err
}

// Cluster scopes the lister to one workspace, allowing users to list and get APIBindings.
func (s *aPIBindingClusterLister) Cluster(cluster logicalcluster.Name) APIBindingLister {
	return &aPIBindingLister{indexer: s.indexer, cluster: cluster}
}

// APIBindingLister can list all APIBindings, or get one in particular.
// All objects returned here must be treated as read-only.
type APIBindingLister interface {
	// List lists all APIBindings in the workspace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*apisv1alpha1.APIBinding, err error)
	// Get retrieves the APIBinding from the indexer for a given workspace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*apisv1alpha1.APIBinding, error)
	APIBindingListerExpansion
}

// aPIBindingLister can list all APIBindings inside a workspace.
type aPIBindingLister struct {
	indexer cache.Indexer
	cluster logicalcluster.Name
}

// List lists all APIBindings in the indexer for a workspace.
func (s *aPIBindingLister) List(selector labels.Selector) (ret []*apisv1alpha1.APIBinding, err error) {
	err = kcpcache.ListAllByCluster(s.indexer, s.cluster, selector, func(i interface{}) {
		ret = append(ret, i.(*apisv1alpha1.APIBinding))
	})
	return ret, err
}

// Get retrieves the APIBinding from the indexer for a given workspace and name.
func (s *aPIBindingLister) Get(name string) (*apisv1alpha1.APIBinding, error) {
	key := kcpcache.ToClusterAwareKey(s.cluster.String(), "", name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(apisv1alpha1.Resource("APIBinding"), name)
	}
	return obj.(*apisv1alpha1.APIBinding), nil
}

// NewAPIBindingLister returns a new APIBindingLister.
// We assume that the indexer:
// - is fed by a workspace-scoped LIST+WATCH
// - uses cache.MetaNamespaceKeyFunc as the key function
func NewAPIBindingLister(indexer cache.Indexer) *aPIBindingScopedLister {
	return &aPIBindingScopedLister{indexer: indexer}
}

// aPIBindingScopedLister can list all APIBindings inside a workspace.
type aPIBindingScopedLister struct {
	indexer cache.Indexer
}

// List lists all APIBindings in the indexer for a workspace.
func (s *aPIBindingScopedLister) List(selector labels.Selector) (ret []*apisv1alpha1.APIBinding, err error) {
	err = cache.ListAll(s.indexer, selector, func(i interface{}) {
		ret = append(ret, i.(*apisv1alpha1.APIBinding))
	})
	return ret, err
}

// Get retrieves the APIBinding from the indexer for a given workspace and name.
func (s *aPIBindingScopedLister) Get(name string) (*apisv1alpha1.APIBinding, error) {
	key := name
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(apisv1alpha1.Resource("APIBinding"), name)
	}
	return obj.(*apisv1alpha1.APIBinding), nil
}
