// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	kubernetesdaskorgv1 "github.com/dask/dask-kubernetes/v2023/dask_kubernetes/operator/go_client/pkg/apis/kubernetes.dask.org/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDaskWorkerGroups implements DaskWorkerGroupInterface
type FakeDaskWorkerGroups struct {
	Fake *FakeKubernetesV1
	ns   string
}

var daskworkergroupsResource = schema.GroupVersionResource{Group: "kubernetes.dask.org", Version: "v1", Resource: "daskworkergroups"}

var daskworkergroupsKind = schema.GroupVersionKind{Group: "kubernetes.dask.org", Version: "v1", Kind: "DaskWorkerGroup"}

// Get takes name of the daskWorkerGroup, and returns the corresponding daskWorkerGroup object, and an error if there is any.
func (c *FakeDaskWorkerGroups) Get(ctx context.Context, name string, options v1.GetOptions) (result *kubernetesdaskorgv1.DaskWorkerGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(daskworkergroupsResource, c.ns, name), &kubernetesdaskorgv1.DaskWorkerGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*kubernetesdaskorgv1.DaskWorkerGroup), err
}

// List takes label and field selectors, and returns the list of DaskWorkerGroups that match those selectors.
func (c *FakeDaskWorkerGroups) List(ctx context.Context, opts v1.ListOptions) (result *kubernetesdaskorgv1.DaskWorkerGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(daskworkergroupsResource, daskworkergroupsKind, c.ns, opts), &kubernetesdaskorgv1.DaskWorkerGroupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &kubernetesdaskorgv1.DaskWorkerGroupList{ListMeta: obj.(*kubernetesdaskorgv1.DaskWorkerGroupList).ListMeta}
	for _, item := range obj.(*kubernetesdaskorgv1.DaskWorkerGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested daskWorkerGroups.
func (c *FakeDaskWorkerGroups) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(daskworkergroupsResource, c.ns, opts))

}

// Create takes the representation of a daskWorkerGroup and creates it.  Returns the server's representation of the daskWorkerGroup, and an error, if there is any.
func (c *FakeDaskWorkerGroups) Create(ctx context.Context, daskWorkerGroup *kubernetesdaskorgv1.DaskWorkerGroup, opts v1.CreateOptions) (result *kubernetesdaskorgv1.DaskWorkerGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(daskworkergroupsResource, c.ns, daskWorkerGroup), &kubernetesdaskorgv1.DaskWorkerGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*kubernetesdaskorgv1.DaskWorkerGroup), err
}

// Update takes the representation of a daskWorkerGroup and updates it. Returns the server's representation of the daskWorkerGroup, and an error, if there is any.
func (c *FakeDaskWorkerGroups) Update(ctx context.Context, daskWorkerGroup *kubernetesdaskorgv1.DaskWorkerGroup, opts v1.UpdateOptions) (result *kubernetesdaskorgv1.DaskWorkerGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(daskworkergroupsResource, c.ns, daskWorkerGroup), &kubernetesdaskorgv1.DaskWorkerGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*kubernetesdaskorgv1.DaskWorkerGroup), err
}

// Delete takes name of the daskWorkerGroup and deletes it. Returns an error if one occurs.
func (c *FakeDaskWorkerGroups) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(daskworkergroupsResource, c.ns, name, opts), &kubernetesdaskorgv1.DaskWorkerGroup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDaskWorkerGroups) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(daskworkergroupsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &kubernetesdaskorgv1.DaskWorkerGroupList{})
	return err
}

// Patch applies the patch and returns the patched daskWorkerGroup.
func (c *FakeDaskWorkerGroups) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *kubernetesdaskorgv1.DaskWorkerGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(daskworkergroupsResource, c.ns, name, pt, data, subresources...), &kubernetesdaskorgv1.DaskWorkerGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*kubernetesdaskorgv1.DaskWorkerGroup), err
}
