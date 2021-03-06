/*
Copyright 2018 The Kubernetes Authors.

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

// This file was automatically generated by informer-gen

package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	cluster_v1alpha1 "k8s.io/kube-deploy/cluster-api/pkg/apis/cluster/v1alpha1"
	clientset "k8s.io/kube-deploy/cluster-api/pkg/client/clientset_generated/clientset"
	internalinterfaces "k8s.io/kube-deploy/cluster-api/pkg/client/informers_generated/externalversions/internalinterfaces"
	v1alpha1 "k8s.io/kube-deploy/cluster-api/pkg/client/listers_generated/cluster/v1alpha1"
	time "time"
)

// MachineInformer provides access to a shared informer and lister for
// Machines.
type MachineInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.MachineLister
}

type machineInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

// NewMachineInformer constructs a new informer for Machine type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewMachineInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				return client.ClusterV1alpha1().Machines(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				return client.ClusterV1alpha1().Machines(namespace).Watch(options)
			},
		},
		&cluster_v1alpha1.Machine{},
		resyncPeriod,
		indexers,
	)
}

func defaultMachineInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewMachineInformer(client, v1.NamespaceAll, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc})
}

func (f *machineInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&cluster_v1alpha1.Machine{}, defaultMachineInformer)
}

func (f *machineInformer) Lister() v1alpha1.MachineLister {
	return v1alpha1.NewMachineLister(f.Informer().GetIndexer())
}
