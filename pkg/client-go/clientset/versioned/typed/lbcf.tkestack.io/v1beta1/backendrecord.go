/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2020 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1beta1 "tkestack.io/lb-controlling-framework/pkg/apis/lbcf.tkestack.io/v1beta1"
	scheme "tkestack.io/lb-controlling-framework/pkg/client-go/clientset/versioned/scheme"
)

// BackendRecordsGetter has a method to return a BackendRecordInterface.
// A group's client should implement this interface.
type BackendRecordsGetter interface {
	BackendRecords(namespace string) BackendRecordInterface
}

// BackendRecordInterface has methods to work with BackendRecord resources.
type BackendRecordInterface interface {
	Create(ctx context.Context, backendRecord *v1beta1.BackendRecord, opts v1.CreateOptions) (*v1beta1.BackendRecord, error)
	Update(ctx context.Context, backendRecord *v1beta1.BackendRecord, opts v1.UpdateOptions) (*v1beta1.BackendRecord, error)
	UpdateStatus(ctx context.Context, backendRecord *v1beta1.BackendRecord, opts v1.UpdateOptions) (*v1beta1.BackendRecord, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.BackendRecord, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.BackendRecordList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.BackendRecord, err error)
	BackendRecordExpansion
}

// backendRecords implements BackendRecordInterface
type backendRecords struct {
	client rest.Interface
	ns     string
}

// newBackendRecords returns a BackendRecords
func newBackendRecords(c *LbcfV1beta1Client, namespace string) *backendRecords {
	return &backendRecords{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the backendRecord, and returns the corresponding backendRecord object, and an error if there is any.
func (c *backendRecords) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.BackendRecord, err error) {
	result = &v1beta1.BackendRecord{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("backendrecords").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of BackendRecords that match those selectors.
func (c *backendRecords) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.BackendRecordList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.BackendRecordList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("backendrecords").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested backendRecords.
func (c *backendRecords) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("backendrecords").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a backendRecord and creates it.  Returns the server's representation of the backendRecord, and an error, if there is any.
func (c *backendRecords) Create(ctx context.Context, backendRecord *v1beta1.BackendRecord, opts v1.CreateOptions) (result *v1beta1.BackendRecord, err error) {
	result = &v1beta1.BackendRecord{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("backendrecords").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(backendRecord).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a backendRecord and updates it. Returns the server's representation of the backendRecord, and an error, if there is any.
func (c *backendRecords) Update(ctx context.Context, backendRecord *v1beta1.BackendRecord, opts v1.UpdateOptions) (result *v1beta1.BackendRecord, err error) {
	result = &v1beta1.BackendRecord{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("backendrecords").
		Name(backendRecord.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(backendRecord).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *backendRecords) UpdateStatus(ctx context.Context, backendRecord *v1beta1.BackendRecord, opts v1.UpdateOptions) (result *v1beta1.BackendRecord, err error) {
	result = &v1beta1.BackendRecord{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("backendrecords").
		Name(backendRecord.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(backendRecord).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the backendRecord and deletes it. Returns an error if one occurs.
func (c *backendRecords) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("backendrecords").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *backendRecords) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("backendrecords").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched backendRecord.
func (c *backendRecords) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.BackendRecord, err error) {
	result = &v1beta1.BackendRecord{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("backendrecords").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
