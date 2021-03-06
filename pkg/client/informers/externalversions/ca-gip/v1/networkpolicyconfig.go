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

package v1

import (
	time "time"

	cagipv1 "github.com/ca-gip/kubi/pkg/apis/ca-gip/v1"
	versioned "github.com/ca-gip/kubi/pkg/client/clientset/versioned"
	internalinterfaces "github.com/ca-gip/kubi/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/ca-gip/kubi/pkg/client/listers/ca-gip/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// NetworkPolicyConfigInformer provides access to a shared informer and lister for
// NetworkPolicyConfigs.
type NetworkPolicyConfigInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.NetworkPolicyConfigLister
}

type networkPolicyConfigInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewNetworkPolicyConfigInformer constructs a new informer for NetworkPolicyConfig type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewNetworkPolicyConfigInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredNetworkPolicyConfigInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredNetworkPolicyConfigInformer constructs a new informer for NetworkPolicyConfig type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredNetworkPolicyConfigInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CagipV1().NetworkPolicyConfigs().List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CagipV1().NetworkPolicyConfigs().Watch(options)
			},
		},
		&cagipv1.NetworkPolicyConfig{},
		resyncPeriod,
		indexers,
	)
}

func (f *networkPolicyConfigInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredNetworkPolicyConfigInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *networkPolicyConfigInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&cagipv1.NetworkPolicyConfig{}, f.defaultInformer)
}

func (f *networkPolicyConfigInformer) Lister() v1.NetworkPolicyConfigLister {
	return v1.NewNetworkPolicyConfigLister(f.Informer().GetIndexer())
}
