/*
Copyright 2023 The Kubernetes Authors.

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

package v1

import (
	"context"
	"time"

	v1 "github.com/satchel9/virtualization/pkg/apis/virtual/v1"
	scheme "github.com/satchel9/virtualization/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DiskRestoresGetter has a method to return a DiskRestoreInterface.
// A group's client should implement this interface.
type DiskRestoresGetter interface {
	DiskRestores(namespace string) DiskRestoreInterface
}

// DiskRestoreInterface has methods to work with DiskRestore resources.
type DiskRestoreInterface interface {
	Create(ctx context.Context, diskRestore *v1.DiskRestore, opts metav1.CreateOptions) (*v1.DiskRestore, error)
	Update(ctx context.Context, diskRestore *v1.DiskRestore, opts metav1.UpdateOptions) (*v1.DiskRestore, error)
	UpdateStatus(ctx context.Context, diskRestore *v1.DiskRestore, opts metav1.UpdateOptions) (*v1.DiskRestore, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.DiskRestore, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.DiskRestoreList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.DiskRestore, err error)
	DiskRestoreExpansion
}

// diskRestores implements DiskRestoreInterface
type diskRestores struct {
	client rest.Interface
	ns     string
}

// newDiskRestores returns a DiskRestores
func newDiskRestores(c *VirtualizationV1Client, namespace string) *diskRestores {
	return &diskRestores{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the diskRestore, and returns the corresponding diskRestore object, and an error if there is any.
func (c *diskRestores) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.DiskRestore, err error) {
	result = &v1.DiskRestore{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("diskrestores").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DiskRestores that match those selectors.
func (c *diskRestores) List(ctx context.Context, opts metav1.ListOptions) (result *v1.DiskRestoreList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.DiskRestoreList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("diskrestores").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested diskRestores.
func (c *diskRestores) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("diskrestores").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a diskRestore and creates it.  Returns the server's representation of the diskRestore, and an error, if there is any.
func (c *diskRestores) Create(ctx context.Context, diskRestore *v1.DiskRestore, opts metav1.CreateOptions) (result *v1.DiskRestore, err error) {
	result = &v1.DiskRestore{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("diskrestores").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(diskRestore).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a diskRestore and updates it. Returns the server's representation of the diskRestore, and an error, if there is any.
func (c *diskRestores) Update(ctx context.Context, diskRestore *v1.DiskRestore, opts metav1.UpdateOptions) (result *v1.DiskRestore, err error) {
	result = &v1.DiskRestore{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("diskrestores").
		Name(diskRestore.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(diskRestore).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *diskRestores) UpdateStatus(ctx context.Context, diskRestore *v1.DiskRestore, opts metav1.UpdateOptions) (result *v1.DiskRestore, err error) {
	result = &v1.DiskRestore{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("diskrestores").
		Name(diskRestore.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(diskRestore).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the diskRestore and deletes it. Returns an error if one occurs.
func (c *diskRestores) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("diskrestores").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *diskRestores) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("diskrestores").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched diskRestore.
func (c *diskRestores) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.DiskRestore, err error) {
	result = &v1.DiskRestore{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("diskrestores").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
