/*
Copyright 2022 The fornax-serverless Authors.

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

	corev1 "centaurusinfra.io/fornax-serverless/pkg/apis/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeApplicationInstances implements ApplicationInstanceInterface
type FakeApplicationInstances struct {
	Fake *FakeCoreV1
	ns   string
}

var applicationinstancesResource = schema.GroupVersionResource{Group: "core.fornax-serverless.centaurusinfra.io", Version: "v1", Resource: "applicationinstances"}

var applicationinstancesKind = schema.GroupVersionKind{Group: "core.fornax-serverless.centaurusinfra.io", Version: "v1", Kind: "ApplicationInstance"}

// Get takes name of the applicationInstance, and returns the corresponding applicationInstance object, and an error if there is any.
func (c *FakeApplicationInstances) Get(ctx context.Context, name string, options v1.GetOptions) (result *corev1.ApplicationInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(applicationinstancesResource, c.ns, name), &corev1.ApplicationInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.ApplicationInstance), err
}

// List takes label and field selectors, and returns the list of ApplicationInstances that match those selectors.
func (c *FakeApplicationInstances) List(ctx context.Context, opts v1.ListOptions) (result *corev1.ApplicationInstanceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(applicationinstancesResource, applicationinstancesKind, c.ns, opts), &corev1.ApplicationInstanceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &corev1.ApplicationInstanceList{ListMeta: obj.(*corev1.ApplicationInstanceList).ListMeta}
	for _, item := range obj.(*corev1.ApplicationInstanceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested applicationInstances.
func (c *FakeApplicationInstances) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(applicationinstancesResource, c.ns, opts))

}

// Create takes the representation of a applicationInstance and creates it.  Returns the server's representation of the applicationInstance, and an error, if there is any.
func (c *FakeApplicationInstances) Create(ctx context.Context, applicationInstance *corev1.ApplicationInstance, opts v1.CreateOptions) (result *corev1.ApplicationInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(applicationinstancesResource, c.ns, applicationInstance), &corev1.ApplicationInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.ApplicationInstance), err
}

// Update takes the representation of a applicationInstance and updates it. Returns the server's representation of the applicationInstance, and an error, if there is any.
func (c *FakeApplicationInstances) Update(ctx context.Context, applicationInstance *corev1.ApplicationInstance, opts v1.UpdateOptions) (result *corev1.ApplicationInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(applicationinstancesResource, c.ns, applicationInstance), &corev1.ApplicationInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.ApplicationInstance), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeApplicationInstances) UpdateStatus(ctx context.Context, applicationInstance *corev1.ApplicationInstance, opts v1.UpdateOptions) (*corev1.ApplicationInstance, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(applicationinstancesResource, "status", c.ns, applicationInstance), &corev1.ApplicationInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.ApplicationInstance), err
}

// Delete takes name of the applicationInstance and deletes it. Returns an error if one occurs.
func (c *FakeApplicationInstances) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(applicationinstancesResource, c.ns, name, opts), &corev1.ApplicationInstance{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeApplicationInstances) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(applicationinstancesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &corev1.ApplicationInstanceList{})
	return err
}

// Patch applies the patch and returns the patched applicationInstance.
func (c *FakeApplicationInstances) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *corev1.ApplicationInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(applicationinstancesResource, c.ns, name, pt, data, subresources...), &corev1.ApplicationInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.ApplicationInstance), err
}