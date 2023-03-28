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

package v1alpha1

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"

	v1alpha1 "github.com/kcp-dev/kcp/sdk/apis/apis/v1alpha1"
	scheme "github.com/kcp-dev/kcp/sdk/client/clientset/versioned/scheme"
)

// APIExportEndpointSlicesGetter has a method to return a APIExportEndpointSliceInterface.
// A group's client should implement this interface.
type APIExportEndpointSlicesGetter interface {
	APIExportEndpointSlices() APIExportEndpointSliceInterface
}

// APIExportEndpointSliceInterface has methods to work with APIExportEndpointSlice resources.
type APIExportEndpointSliceInterface interface {
	Create(ctx context.Context, aPIExportEndpointSlice *v1alpha1.APIExportEndpointSlice, opts v1.CreateOptions) (*v1alpha1.APIExportEndpointSlice, error)
	Update(ctx context.Context, aPIExportEndpointSlice *v1alpha1.APIExportEndpointSlice, opts v1.UpdateOptions) (*v1alpha1.APIExportEndpointSlice, error)
	UpdateStatus(ctx context.Context, aPIExportEndpointSlice *v1alpha1.APIExportEndpointSlice, opts v1.UpdateOptions) (*v1alpha1.APIExportEndpointSlice, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.APIExportEndpointSlice, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.APIExportEndpointSliceList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.APIExportEndpointSlice, err error)
	APIExportEndpointSliceExpansion
}

// aPIExportEndpointSlices implements APIExportEndpointSliceInterface
type aPIExportEndpointSlices struct {
	client rest.Interface
}

// newAPIExportEndpointSlices returns a APIExportEndpointSlices
func newAPIExportEndpointSlices(c *ApisV1alpha1Client) *aPIExportEndpointSlices {
	return &aPIExportEndpointSlices{
		client: c.RESTClient(),
	}
}

// Get takes name of the aPIExportEndpointSlice, and returns the corresponding aPIExportEndpointSlice object, and an error if there is any.
func (c *aPIExportEndpointSlices) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.APIExportEndpointSlice, err error) {
	result = &v1alpha1.APIExportEndpointSlice{}
	err = c.client.Get().
		Resource("apiexportendpointslices").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of APIExportEndpointSlices that match those selectors.
func (c *aPIExportEndpointSlices) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.APIExportEndpointSliceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.APIExportEndpointSliceList{}
	err = c.client.Get().
		Resource("apiexportendpointslices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested aPIExportEndpointSlices.
func (c *aPIExportEndpointSlices) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("apiexportendpointslices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a aPIExportEndpointSlice and creates it.  Returns the server's representation of the aPIExportEndpointSlice, and an error, if there is any.
func (c *aPIExportEndpointSlices) Create(ctx context.Context, aPIExportEndpointSlice *v1alpha1.APIExportEndpointSlice, opts v1.CreateOptions) (result *v1alpha1.APIExportEndpointSlice, err error) {
	result = &v1alpha1.APIExportEndpointSlice{}
	err = c.client.Post().
		Resource("apiexportendpointslices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(aPIExportEndpointSlice).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a aPIExportEndpointSlice and updates it. Returns the server's representation of the aPIExportEndpointSlice, and an error, if there is any.
func (c *aPIExportEndpointSlices) Update(ctx context.Context, aPIExportEndpointSlice *v1alpha1.APIExportEndpointSlice, opts v1.UpdateOptions) (result *v1alpha1.APIExportEndpointSlice, err error) {
	result = &v1alpha1.APIExportEndpointSlice{}
	err = c.client.Put().
		Resource("apiexportendpointslices").
		Name(aPIExportEndpointSlice.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(aPIExportEndpointSlice).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *aPIExportEndpointSlices) UpdateStatus(ctx context.Context, aPIExportEndpointSlice *v1alpha1.APIExportEndpointSlice, opts v1.UpdateOptions) (result *v1alpha1.APIExportEndpointSlice, err error) {
	result = &v1alpha1.APIExportEndpointSlice{}
	err = c.client.Put().
		Resource("apiexportendpointslices").
		Name(aPIExportEndpointSlice.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(aPIExportEndpointSlice).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the aPIExportEndpointSlice and deletes it. Returns an error if one occurs.
func (c *aPIExportEndpointSlices) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("apiexportendpointslices").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *aPIExportEndpointSlices) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("apiexportendpointslices").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched aPIExportEndpointSlice.
func (c *aPIExportEndpointSlices) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.APIExportEndpointSlice, err error) {
	result = &v1alpha1.APIExportEndpointSlice{}
	err = c.client.Patch(pt).
		Resource("apiexportendpointslices").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}