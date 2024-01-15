// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1 "open-cluster-management.io/api/work/v1"
)

// FakeManifestWorks implements ManifestWorkInterface
type FakeManifestWorks struct {
	Fake *FakeWorkV1
	ns   string
}

var manifestworksResource = v1.SchemeGroupVersion.WithResource("manifestworks")

var manifestworksKind = v1.SchemeGroupVersion.WithKind("ManifestWork")

// Get takes name of the manifestWork, and returns the corresponding manifestWork object, and an error if there is any.
func (c *FakeManifestWorks) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.ManifestWork, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(manifestworksResource, c.ns, name), &v1.ManifestWork{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ManifestWork), err
}

// List takes label and field selectors, and returns the list of ManifestWorks that match those selectors.
func (c *FakeManifestWorks) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ManifestWorkList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(manifestworksResource, manifestworksKind, c.ns, opts), &v1.ManifestWorkList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.ManifestWorkList{ListMeta: obj.(*v1.ManifestWorkList).ListMeta}
	for _, item := range obj.(*v1.ManifestWorkList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested manifestWorks.
func (c *FakeManifestWorks) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(manifestworksResource, c.ns, opts))

}

// Create takes the representation of a manifestWork and creates it.  Returns the server's representation of the manifestWork, and an error, if there is any.
func (c *FakeManifestWorks) Create(ctx context.Context, manifestWork *v1.ManifestWork, opts metav1.CreateOptions) (result *v1.ManifestWork, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(manifestworksResource, c.ns, manifestWork), &v1.ManifestWork{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ManifestWork), err
}

// Update takes the representation of a manifestWork and updates it. Returns the server's representation of the manifestWork, and an error, if there is any.
func (c *FakeManifestWorks) Update(ctx context.Context, manifestWork *v1.ManifestWork, opts metav1.UpdateOptions) (result *v1.ManifestWork, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(manifestworksResource, c.ns, manifestWork), &v1.ManifestWork{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ManifestWork), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeManifestWorks) UpdateStatus(ctx context.Context, manifestWork *v1.ManifestWork, opts metav1.UpdateOptions) (*v1.ManifestWork, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(manifestworksResource, "status", c.ns, manifestWork), &v1.ManifestWork{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ManifestWork), err
}

// Delete takes name of the manifestWork and deletes it. Returns an error if one occurs.
func (c *FakeManifestWorks) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(manifestworksResource, c.ns, name, opts), &v1.ManifestWork{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeManifestWorks) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(manifestworksResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1.ManifestWorkList{})
	return err
}

// Patch applies the patch and returns the patched manifestWork.
func (c *FakeManifestWorks) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ManifestWork, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(manifestworksResource, c.ns, name, pt, data, subresources...), &v1.ManifestWork{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ManifestWork), err
}
