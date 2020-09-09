/*
Copyright The Kubernetes Authors.

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
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	cloudgatewayv1 "k8s.io/kubernetes/pkg/apis/cloudgateway/v1"
)

// FakeEGateways implements EGatewayInterface
type FakeEGateways struct {
	Fake *FakeCloudgatewayV1
	ns   string
	te   string
}

var egatewaysResource = schema.GroupVersionResource{Group: "cloudgateway.arktos.io", Version: "v1", Resource: "egateways"}

var egatewaysKind = schema.GroupVersionKind{Group: "cloudgateway.arktos.io", Version: "v1", Kind: "EGateway"}

// Get takes name of the eGateway, and returns the corresponding eGateway object, and an error if there is any.
func (c *FakeEGateways) Get(name string, options v1.GetOptions) (result *cloudgatewayv1.EGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithMultiTenancy(egatewaysResource, c.ns, name, c.te), &cloudgatewayv1.EGateway{})

	if obj == nil {
		return nil, err
	}

	return obj.(*cloudgatewayv1.EGateway), err
}

// List takes label and field selectors, and returns the list of EGateways that match those selectors.
func (c *FakeEGateways) List(opts v1.ListOptions) (result *cloudgatewayv1.EGatewayList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithMultiTenancy(egatewaysResource, egatewaysKind, c.ns, opts, c.te), &cloudgatewayv1.EGatewayList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &cloudgatewayv1.EGatewayList{ListMeta: obj.(*cloudgatewayv1.EGatewayList).ListMeta}
	for _, item := range obj.(*cloudgatewayv1.EGatewayList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.AggregatedWatchInterface that watches the requested eGateways.
func (c *FakeEGateways) Watch(opts v1.ListOptions) watch.AggregatedWatchInterface {
	aggWatch := watch.NewAggregatedWatcher()
	watcher, err := c.Fake.
		InvokesWatch(testing.NewWatchActionWithMultiTenancy(egatewaysResource, c.ns, opts, c.te))

	aggWatch.AddWatchInterface(watcher, err)
	return aggWatch
}

// Create takes the representation of a eGateway and creates it.  Returns the server's representation of the eGateway, and an error, if there is any.
func (c *FakeEGateways) Create(eGateway *cloudgatewayv1.EGateway) (result *cloudgatewayv1.EGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithMultiTenancy(egatewaysResource, c.ns, eGateway, c.te), &cloudgatewayv1.EGateway{})

	if obj == nil {
		return nil, err
	}

	return obj.(*cloudgatewayv1.EGateway), err
}

// Update takes the representation of a eGateway and updates it. Returns the server's representation of the eGateway, and an error, if there is any.
func (c *FakeEGateways) Update(eGateway *cloudgatewayv1.EGateway) (result *cloudgatewayv1.EGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithMultiTenancy(egatewaysResource, c.ns, eGateway, c.te), &cloudgatewayv1.EGateway{})

	if obj == nil {
		return nil, err
	}

	return obj.(*cloudgatewayv1.EGateway), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeEGateways) UpdateStatus(eGateway *cloudgatewayv1.EGateway) (*cloudgatewayv1.EGateway, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithMultiTenancy(egatewaysResource, "status", c.ns, eGateway, c.te), &cloudgatewayv1.EGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cloudgatewayv1.EGateway), err
}

// Delete takes name of the eGateway and deletes it. Returns an error if one occurs.
func (c *FakeEGateways) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithMultiTenancy(egatewaysResource, c.ns, name, c.te), &cloudgatewayv1.EGateway{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEGateways) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithMultiTenancy(egatewaysResource, c.ns, listOptions, c.te)

	_, err := c.Fake.Invokes(action, &cloudgatewayv1.EGatewayList{})
	return err
}

// Patch applies the patch and returns the patched eGateway.
func (c *FakeEGateways) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *cloudgatewayv1.EGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithMultiTenancy(egatewaysResource, c.te, c.ns, name, pt, data, subresources...), &cloudgatewayv1.EGateway{})

	if obj == nil {
		return nil, err
	}

	return obj.(*cloudgatewayv1.EGateway), err
}