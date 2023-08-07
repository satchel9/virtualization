// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	virtualizationmatrixiov1 "virtualization/pkg/apis/virtualization.matrix.io/v1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDiskRestores implements DiskRestoreInterface
type FakeDiskRestores struct {
	Fake *FakeVirtualizationV1
	ns   string
}

var diskrestoresResource = schema.GroupVersionResource{Group: "virtualization.matrix.io", Version: "v1", Resource: "diskrestores"}

var diskrestoresKind = schema.GroupVersionKind{Group: "virtualization.matrix.io", Version: "v1", Kind: "DiskRestore"}

// Get takes name of the diskRestore, and returns the corresponding diskRestore object, and an error if there is any.
func (c *FakeDiskRestores) Get(ctx context.Context, name string, options v1.GetOptions) (result *virtualizationmatrixiov1.DiskRestore, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(diskrestoresResource, c.ns, name), &virtualizationmatrixiov1.DiskRestore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*virtualizationmatrixiov1.DiskRestore), err
}

// List takes label and field selectors, and returns the list of DiskRestores that match those selectors.
func (c *FakeDiskRestores) List(ctx context.Context, opts v1.ListOptions) (result *virtualizationmatrixiov1.DiskRestoreList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(diskrestoresResource, diskrestoresKind, c.ns, opts), &virtualizationmatrixiov1.DiskRestoreList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &virtualizationmatrixiov1.DiskRestoreList{ListMeta: obj.(*virtualizationmatrixiov1.DiskRestoreList).ListMeta}
	for _, item := range obj.(*virtualizationmatrixiov1.DiskRestoreList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested diskRestores.
func (c *FakeDiskRestores) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(diskrestoresResource, c.ns, opts))

}

// Create takes the representation of a diskRestore and creates it.  Returns the server's representation of the diskRestore, and an error, if there is any.
func (c *FakeDiskRestores) Create(ctx context.Context, diskRestore *virtualizationmatrixiov1.DiskRestore, opts v1.CreateOptions) (result *virtualizationmatrixiov1.DiskRestore, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(diskrestoresResource, c.ns, diskRestore), &virtualizationmatrixiov1.DiskRestore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*virtualizationmatrixiov1.DiskRestore), err
}

// Update takes the representation of a diskRestore and updates it. Returns the server's representation of the diskRestore, and an error, if there is any.
func (c *FakeDiskRestores) Update(ctx context.Context, diskRestore *virtualizationmatrixiov1.DiskRestore, opts v1.UpdateOptions) (result *virtualizationmatrixiov1.DiskRestore, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(diskrestoresResource, c.ns, diskRestore), &virtualizationmatrixiov1.DiskRestore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*virtualizationmatrixiov1.DiskRestore), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDiskRestores) UpdateStatus(ctx context.Context, diskRestore *virtualizationmatrixiov1.DiskRestore, opts v1.UpdateOptions) (*virtualizationmatrixiov1.DiskRestore, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(diskrestoresResource, "status", c.ns, diskRestore), &virtualizationmatrixiov1.DiskRestore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*virtualizationmatrixiov1.DiskRestore), err
}

// Delete takes name of the diskRestore and deletes it. Returns an error if one occurs.
func (c *FakeDiskRestores) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(diskrestoresResource, c.ns, name, opts), &virtualizationmatrixiov1.DiskRestore{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDiskRestores) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(diskrestoresResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &virtualizationmatrixiov1.DiskRestoreList{})
	return err
}

// Patch applies the patch and returns the patched diskRestore.
func (c *FakeDiskRestores) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *virtualizationmatrixiov1.DiskRestore, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(diskrestoresResource, c.ns, name, pt, data, subresources...), &virtualizationmatrixiov1.DiskRestore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*virtualizationmatrixiov1.DiskRestore), err
}