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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"

	v1alpha1 "github.com/kcp-dev/kcp/contrib/mounts-vw/apis/proxy/v1alpha1"
	proxyv1alpha1 "github.com/kcp-dev/kcp/contrib/mounts-vw/client/applyconfiguration/proxy/v1alpha1"
)

// FakeTargetKubeClusters implements TargetKubeClusterInterface
type FakeTargetKubeClusters struct {
	Fake *FakeProxyV1alpha1
}

var targetkubeclustersResource = v1alpha1.SchemeGroupVersion.WithResource("targetkubeclusters")

var targetkubeclustersKind = v1alpha1.SchemeGroupVersion.WithKind("TargetKubeCluster")

// Get takes name of the targetKubeCluster, and returns the corresponding targetKubeCluster object, and an error if there is any.
func (c *FakeTargetKubeClusters) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.TargetKubeCluster, err error) {
	emptyResult := &v1alpha1.TargetKubeCluster{}
	obj, err := c.Fake.
		Invokes(testing.NewRootGetActionWithOptions(targetkubeclustersResource, name, options), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.TargetKubeCluster), err
}

// List takes label and field selectors, and returns the list of TargetKubeClusters that match those selectors.
func (c *FakeTargetKubeClusters) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.TargetKubeClusterList, err error) {
	emptyResult := &v1alpha1.TargetKubeClusterList{}
	obj, err := c.Fake.
		Invokes(testing.NewRootListActionWithOptions(targetkubeclustersResource, targetkubeclustersKind, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.TargetKubeClusterList{ListMeta: obj.(*v1alpha1.TargetKubeClusterList).ListMeta}
	for _, item := range obj.(*v1alpha1.TargetKubeClusterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested targetKubeClusters.
func (c *FakeTargetKubeClusters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchActionWithOptions(targetkubeclustersResource, opts))
}

// Create takes the representation of a targetKubeCluster and creates it.  Returns the server's representation of the targetKubeCluster, and an error, if there is any.
func (c *FakeTargetKubeClusters) Create(ctx context.Context, targetKubeCluster *v1alpha1.TargetKubeCluster, opts v1.CreateOptions) (result *v1alpha1.TargetKubeCluster, err error) {
	emptyResult := &v1alpha1.TargetKubeCluster{}
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateActionWithOptions(targetkubeclustersResource, targetKubeCluster, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.TargetKubeCluster), err
}

// Update takes the representation of a targetKubeCluster and updates it. Returns the server's representation of the targetKubeCluster, and an error, if there is any.
func (c *FakeTargetKubeClusters) Update(ctx context.Context, targetKubeCluster *v1alpha1.TargetKubeCluster, opts v1.UpdateOptions) (result *v1alpha1.TargetKubeCluster, err error) {
	emptyResult := &v1alpha1.TargetKubeCluster{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateActionWithOptions(targetkubeclustersResource, targetKubeCluster, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.TargetKubeCluster), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeTargetKubeClusters) UpdateStatus(ctx context.Context, targetKubeCluster *v1alpha1.TargetKubeCluster, opts v1.UpdateOptions) (result *v1alpha1.TargetKubeCluster, err error) {
	emptyResult := &v1alpha1.TargetKubeCluster{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceActionWithOptions(targetkubeclustersResource, "status", targetKubeCluster, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.TargetKubeCluster), err
}

// Delete takes name of the targetKubeCluster and deletes it. Returns an error if one occurs.
func (c *FakeTargetKubeClusters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(targetkubeclustersResource, name, opts), &v1alpha1.TargetKubeCluster{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTargetKubeClusters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionActionWithOptions(targetkubeclustersResource, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.TargetKubeClusterList{})
	return err
}

// Patch applies the patch and returns the patched targetKubeCluster.
func (c *FakeTargetKubeClusters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.TargetKubeCluster, err error) {
	emptyResult := &v1alpha1.TargetKubeCluster{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(targetkubeclustersResource, name, pt, data, opts, subresources...), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.TargetKubeCluster), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied targetKubeCluster.
func (c *FakeTargetKubeClusters) Apply(ctx context.Context, targetKubeCluster *proxyv1alpha1.TargetKubeClusterApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.TargetKubeCluster, err error) {
	if targetKubeCluster == nil {
		return nil, fmt.Errorf("targetKubeCluster provided to Apply must not be nil")
	}
	data, err := json.Marshal(targetKubeCluster)
	if err != nil {
		return nil, err
	}
	name := targetKubeCluster.Name
	if name == nil {
		return nil, fmt.Errorf("targetKubeCluster.Name must be provided to Apply")
	}
	emptyResult := &v1alpha1.TargetKubeCluster{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(targetkubeclustersResource, *name, types.ApplyPatchType, data, opts.ToPatchOptions()), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.TargetKubeCluster), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeTargetKubeClusters) ApplyStatus(ctx context.Context, targetKubeCluster *proxyv1alpha1.TargetKubeClusterApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.TargetKubeCluster, err error) {
	if targetKubeCluster == nil {
		return nil, fmt.Errorf("targetKubeCluster provided to Apply must not be nil")
	}
	data, err := json.Marshal(targetKubeCluster)
	if err != nil {
		return nil, err
	}
	name := targetKubeCluster.Name
	if name == nil {
		return nil, fmt.Errorf("targetKubeCluster.Name must be provided to Apply")
	}
	emptyResult := &v1alpha1.TargetKubeCluster{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(targetkubeclustersResource, *name, types.ApplyPatchType, data, opts.ToPatchOptions(), "status"), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.TargetKubeCluster), err
}
