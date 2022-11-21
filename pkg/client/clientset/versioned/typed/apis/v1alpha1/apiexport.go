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

	v1alpha1 "github.com/kcp-dev/kcp/pkg/apis/apis/v1alpha1"
	scheme "github.com/kcp-dev/kcp/pkg/client/clientset/versioned/scheme"
)

// APIExportsGetter has a method to return a APIExportInterface.
// A group's client should implement this interface.
type APIExportsGetter interface {
	APIExports() APIExportInterface
}

// APIExportInterface has methods to work with APIExport resources.
type APIExportInterface interface {
	Create(ctx context.Context, aPIExport *v1alpha1.APIExport, opts v1.CreateOptions) (*v1alpha1.APIExport, error)
	Update(ctx context.Context, aPIExport *v1alpha1.APIExport, opts v1.UpdateOptions) (*v1alpha1.APIExport, error)
	UpdateStatus(ctx context.Context, aPIExport *v1alpha1.APIExport, opts v1.UpdateOptions) (*v1alpha1.APIExport, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.APIExport, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.APIExportList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.APIExport, err error)
	APIExportExpansion
}

// aPIExports implements APIExportInterface
type aPIExports struct {
	client rest.Interface
}

// newAPIExports returns a APIExports
func newAPIExports(c *ApisV1alpha1Client) *aPIExports {
	return &aPIExports{
		client: c.RESTClient(),
	}
}

// Get takes name of the aPIExport, and returns the corresponding aPIExport object, and an error if there is any.
func (c *aPIExports) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.APIExport, err error) {
	result = &v1alpha1.APIExport{}
	err = c.client.Get().
		Resource("apiexports").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of APIExports that match those selectors.
func (c *aPIExports) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.APIExportList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.APIExportList{}
	err = c.client.Get().
		Resource("apiexports").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested aPIExports.
func (c *aPIExports) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("apiexports").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a aPIExport and creates it.  Returns the server's representation of the aPIExport, and an error, if there is any.
func (c *aPIExports) Create(ctx context.Context, aPIExport *v1alpha1.APIExport, opts v1.CreateOptions) (result *v1alpha1.APIExport, err error) {
	result = &v1alpha1.APIExport{}
	err = c.client.Post().
		Resource("apiexports").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(aPIExport).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a aPIExport and updates it. Returns the server's representation of the aPIExport, and an error, if there is any.
func (c *aPIExports) Update(ctx context.Context, aPIExport *v1alpha1.APIExport, opts v1.UpdateOptions) (result *v1alpha1.APIExport, err error) {
	result = &v1alpha1.APIExport{}
	err = c.client.Put().
		Resource("apiexports").
		Name(aPIExport.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(aPIExport).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *aPIExports) UpdateStatus(ctx context.Context, aPIExport *v1alpha1.APIExport, opts v1.UpdateOptions) (result *v1alpha1.APIExport, err error) {
	result = &v1alpha1.APIExport{}
	err = c.client.Put().
		Resource("apiexports").
		Name(aPIExport.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(aPIExport).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the aPIExport and deletes it. Returns an error if one occurs.
func (c *aPIExports) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("apiexports").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *aPIExports) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("apiexports").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched aPIExport.
func (c *aPIExports) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.APIExport, err error) {
	result = &v1alpha1.APIExport{}
	err = c.client.Patch(pt).
		Resource("apiexports").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
