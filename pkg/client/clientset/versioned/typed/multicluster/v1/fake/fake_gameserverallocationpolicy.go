// Copyright 2024 Google LLC All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was autogenerated. Do not edit directly.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1 "agones.dev/agones/pkg/apis/multicluster/v1"
	multiclusterv1 "agones.dev/agones/pkg/client/applyconfiguration/multicluster/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeGameServerAllocationPolicies implements GameServerAllocationPolicyInterface
type FakeGameServerAllocationPolicies struct {
	Fake *FakeMulticlusterV1
	ns   string
}

var gameserverallocationpoliciesResource = v1.SchemeGroupVersion.WithResource("gameserverallocationpolicies")

var gameserverallocationpoliciesKind = v1.SchemeGroupVersion.WithKind("GameServerAllocationPolicy")

// Get takes name of the gameServerAllocationPolicy, and returns the corresponding gameServerAllocationPolicy object, and an error if there is any.
func (c *FakeGameServerAllocationPolicies) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.GameServerAllocationPolicy, err error) {
	emptyResult := &v1.GameServerAllocationPolicy{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(gameserverallocationpoliciesResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.GameServerAllocationPolicy), err
}

// List takes label and field selectors, and returns the list of GameServerAllocationPolicies that match those selectors.
func (c *FakeGameServerAllocationPolicies) List(ctx context.Context, opts metav1.ListOptions) (result *v1.GameServerAllocationPolicyList, err error) {
	emptyResult := &v1.GameServerAllocationPolicyList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(gameserverallocationpoliciesResource, gameserverallocationpoliciesKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.GameServerAllocationPolicyList{ListMeta: obj.(*v1.GameServerAllocationPolicyList).ListMeta}
	for _, item := range obj.(*v1.GameServerAllocationPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested gameServerAllocationPolicies.
func (c *FakeGameServerAllocationPolicies) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(gameserverallocationpoliciesResource, c.ns, opts))

}

// Create takes the representation of a gameServerAllocationPolicy and creates it.  Returns the server's representation of the gameServerAllocationPolicy, and an error, if there is any.
func (c *FakeGameServerAllocationPolicies) Create(ctx context.Context, gameServerAllocationPolicy *v1.GameServerAllocationPolicy, opts metav1.CreateOptions) (result *v1.GameServerAllocationPolicy, err error) {
	emptyResult := &v1.GameServerAllocationPolicy{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(gameserverallocationpoliciesResource, c.ns, gameServerAllocationPolicy, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.GameServerAllocationPolicy), err
}

// Update takes the representation of a gameServerAllocationPolicy and updates it. Returns the server's representation of the gameServerAllocationPolicy, and an error, if there is any.
func (c *FakeGameServerAllocationPolicies) Update(ctx context.Context, gameServerAllocationPolicy *v1.GameServerAllocationPolicy, opts metav1.UpdateOptions) (result *v1.GameServerAllocationPolicy, err error) {
	emptyResult := &v1.GameServerAllocationPolicy{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(gameserverallocationpoliciesResource, c.ns, gameServerAllocationPolicy, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.GameServerAllocationPolicy), err
}

// Delete takes name of the gameServerAllocationPolicy and deletes it. Returns an error if one occurs.
func (c *FakeGameServerAllocationPolicies) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(gameserverallocationpoliciesResource, c.ns, name, opts), &v1.GameServerAllocationPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGameServerAllocationPolicies) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(gameserverallocationpoliciesResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1.GameServerAllocationPolicyList{})
	return err
}

// Patch applies the patch and returns the patched gameServerAllocationPolicy.
func (c *FakeGameServerAllocationPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.GameServerAllocationPolicy, err error) {
	emptyResult := &v1.GameServerAllocationPolicy{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(gameserverallocationpoliciesResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.GameServerAllocationPolicy), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied gameServerAllocationPolicy.
func (c *FakeGameServerAllocationPolicies) Apply(ctx context.Context, gameServerAllocationPolicy *multiclusterv1.GameServerAllocationPolicyApplyConfiguration, opts metav1.ApplyOptions) (result *v1.GameServerAllocationPolicy, err error) {
	if gameServerAllocationPolicy == nil {
		return nil, fmt.Errorf("gameServerAllocationPolicy provided to Apply must not be nil")
	}
	data, err := json.Marshal(gameServerAllocationPolicy)
	if err != nil {
		return nil, err
	}
	name := gameServerAllocationPolicy.Name
	if name == nil {
		return nil, fmt.Errorf("gameServerAllocationPolicy.Name must be provided to Apply")
	}
	emptyResult := &v1.GameServerAllocationPolicy{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(gameserverallocationpoliciesResource, c.ns, *name, types.ApplyPatchType, data, opts.ToPatchOptions()), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.GameServerAllocationPolicy), err
}
