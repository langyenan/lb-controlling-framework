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

// Code generated by informer-gen. DO NOT EDIT.

package v1beta1

import (
	time "time"

	lbcftkecloudtencentcomv1beta1 "tkestack.io/lb-controlling-framework/pkg/apis/lbcf.tke.cloud.tencent.com/v1beta1"
	versioned "tkestack.io/lb-controlling-framework/pkg/client-go/clientset/versioned"
	internalinterfaces "tkestack.io/lb-controlling-framework/pkg/client-go/informers/externalversions/internalinterfaces"
	v1beta1 "tkestack.io/lb-controlling-framework/pkg/client-go/listers/lbcf.tke.cloud.tencent.com/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// BackendGroupInformer provides access to a shared informer and lister for
// BackendGroups.
type BackendGroupInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.BackendGroupLister
}

type backendGroupInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewBackendGroupInformer constructs a new informer for BackendGroup type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewBackendGroupInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredBackendGroupInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredBackendGroupInformer constructs a new informer for BackendGroup type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredBackendGroupInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.LbcfV1beta1().BackendGroups(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.LbcfV1beta1().BackendGroups(namespace).Watch(options)
			},
		},
		&lbcftkecloudtencentcomv1beta1.BackendGroup{},
		resyncPeriod,
		indexers,
	)
}

func (f *backendGroupInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredBackendGroupInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *backendGroupInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&lbcftkecloudtencentcomv1beta1.BackendGroup{}, f.defaultInformer)
}

func (f *backendGroupInformer) Lister() v1beta1.BackendGroupLister {
	return v1beta1.NewBackendGroupLister(f.Informer().GetIndexer())
}
