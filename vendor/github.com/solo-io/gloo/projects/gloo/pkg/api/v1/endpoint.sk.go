// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"sort"

	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/crd"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/errors"
	"github.com/solo-io/solo-kit/pkg/utils/hashutils"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func NewEndpoint(namespace, name string) *Endpoint {
	endpoint := &Endpoint{}
	endpoint.SetMetadata(core.Metadata{
		Name:      name,
		Namespace: namespace,
	})
	return endpoint
}

func (r *Endpoint) SetMetadata(meta core.Metadata) {
	r.Metadata = meta
}

func (r *Endpoint) Hash() uint64 {
	metaCopy := r.GetMetadata()
	metaCopy.ResourceVersion = ""
	return hashutils.HashAll(
		metaCopy,
		r.Upstreams,
		r.Address,
		r.Port,
	)
}

type EndpointList []*Endpoint
type EndpointsByNamespace map[string]EndpointList

// namespace is optional, if left empty, names can collide if the list contains more than one with the same name
func (list EndpointList) Find(namespace, name string) (*Endpoint, error) {
	for _, endpoint := range list {
		if endpoint.GetMetadata().Name == name {
			if namespace == "" || endpoint.GetMetadata().Namespace == namespace {
				return endpoint, nil
			}
		}
	}
	return nil, errors.Errorf("list did not find endpoint %v.%v", namespace, name)
}

func (list EndpointList) AsResources() resources.ResourceList {
	var ress resources.ResourceList
	for _, endpoint := range list {
		ress = append(ress, endpoint)
	}
	return ress
}

func (list EndpointList) Names() []string {
	var names []string
	for _, endpoint := range list {
		names = append(names, endpoint.GetMetadata().Name)
	}
	return names
}

func (list EndpointList) NamespacesDotNames() []string {
	var names []string
	for _, endpoint := range list {
		names = append(names, endpoint.GetMetadata().Namespace+"."+endpoint.GetMetadata().Name)
	}
	return names
}

func (list EndpointList) Sort() EndpointList {
	sort.SliceStable(list, func(i, j int) bool {
		return list[i].GetMetadata().Less(list[j].GetMetadata())
	})
	return list
}

func (list EndpointList) Clone() EndpointList {
	var endpointList EndpointList
	for _, endpoint := range list {
		endpointList = append(endpointList, resources.Clone(endpoint).(*Endpoint))
	}
	return endpointList
}

func (list EndpointList) Each(f func(element *Endpoint)) {
	for _, endpoint := range list {
		f(endpoint)
	}
}

func (list EndpointList) AsInterfaces() []interface{} {
	var asInterfaces []interface{}
	list.Each(func(element *Endpoint) {
		asInterfaces = append(asInterfaces, element)
	})
	return asInterfaces
}

func (byNamespace EndpointsByNamespace) Add(endpoint ...*Endpoint) {
	for _, item := range endpoint {
		byNamespace[item.GetMetadata().Namespace] = append(byNamespace[item.GetMetadata().Namespace], item)
	}
}

func (byNamespace EndpointsByNamespace) Clear(namespace string) {
	delete(byNamespace, namespace)
}

func (byNamespace EndpointsByNamespace) List() EndpointList {
	var list EndpointList
	for _, endpointList := range byNamespace {
		list = append(list, endpointList...)
	}
	return list.Sort()
}

func (byNamespace EndpointsByNamespace) Clone() EndpointsByNamespace {
	cloned := make(EndpointsByNamespace)
	for ns, list := range byNamespace {
		cloned[ns] = list.Clone()
	}
	return cloned
}

var _ resources.Resource = &Endpoint{}

// Kubernetes Adapter for Endpoint

func (o *Endpoint) GetObjectKind() schema.ObjectKind {
	t := EndpointCrd.TypeMeta()
	return &t
}

func (o *Endpoint) DeepCopyObject() runtime.Object {
	return resources.Clone(o).(*Endpoint)
}

var EndpointCrd = crd.NewCrd("gloo.solo.io",
	"endpoints",
	"gloo.solo.io",
	"v1",
	"Endpoint",
	"ep",
	false,
	&Endpoint{})
