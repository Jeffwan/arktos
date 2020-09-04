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

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1 "k8s.io/kubernetes/pkg/apis/cloudgateway/v1"
)

// EPolicyLister helps list EPolicies.
type EPolicyLister interface {
	// List lists all EPolicies in the indexer.
	List(selector labels.Selector) (ret []*v1.EPolicy, err error)
	// EPolicies returns an object that can list and get EPolicies.
	EPolicies(namespace string) EPolicyNamespaceLister
	EPoliciesWithMultiTenancy(namespace string, tenant string) EPolicyNamespaceLister
	EPolicyListerExpansion
}

// ePolicyLister implements the EPolicyLister interface.
type ePolicyLister struct {
	indexer cache.Indexer
}

// NewEPolicyLister returns a new EPolicyLister.
func NewEPolicyLister(indexer cache.Indexer) EPolicyLister {
	return &ePolicyLister{indexer: indexer}
}

// List lists all EPolicies in the indexer.
func (s *ePolicyLister) List(selector labels.Selector) (ret []*v1.EPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.EPolicy))
	})
	return ret, err
}

// EPolicies returns an object that can list and get EPolicies.
func (s *ePolicyLister) EPolicies(namespace string) EPolicyNamespaceLister {
	return ePolicyNamespaceLister{indexer: s.indexer, namespace: namespace, tenant: "system"}
}

func (s *ePolicyLister) EPoliciesWithMultiTenancy(namespace string, tenant string) EPolicyNamespaceLister {
	return ePolicyNamespaceLister{indexer: s.indexer, namespace: namespace, tenant: tenant}
}

// EPolicyNamespaceLister helps list and get EPolicies.
type EPolicyNamespaceLister interface {
	// List lists all EPolicies in the indexer for a given tenant/namespace.
	List(selector labels.Selector) (ret []*v1.EPolicy, err error)
	// Get retrieves the EPolicy from the indexer for a given tenant/namespace and name.
	Get(name string) (*v1.EPolicy, error)
	EPolicyNamespaceListerExpansion
}

// ePolicyNamespaceLister implements the EPolicyNamespaceLister
// interface.
type ePolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
	tenant    string
}

// List lists all EPolicies in the indexer for a given namespace.
func (s ePolicyNamespaceLister) List(selector labels.Selector) (ret []*v1.EPolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.tenant, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.EPolicy))
	})
	return ret, err
}

// Get retrieves the EPolicy from the indexer for a given namespace and name.
func (s ePolicyNamespaceLister) Get(name string) (*v1.EPolicy, error) {
	key := s.tenant + "/" + s.namespace + "/" + name
	if s.tenant == "system" {
		key = s.namespace + "/" + name
	}
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("epolicy"), name)
	}
	return obj.(*v1.EPolicy), nil
}