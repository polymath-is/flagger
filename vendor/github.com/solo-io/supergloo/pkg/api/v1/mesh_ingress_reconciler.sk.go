// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/reconcile"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
)

// Option to copy anything from the original to the desired before writing. Return value of false means don't update
type TransitionMeshIngressFunc func(original, desired *MeshIngress) (bool, error)

type MeshIngressReconciler interface {
	Reconcile(namespace string, desiredResources MeshIngressList, transition TransitionMeshIngressFunc, opts clients.ListOpts) error
}

func meshIngresssToResources(list MeshIngressList) resources.ResourceList {
	var resourceList resources.ResourceList
	for _, meshIngress := range list {
		resourceList = append(resourceList, meshIngress)
	}
	return resourceList
}

func NewMeshIngressReconciler(client MeshIngressClient) MeshIngressReconciler {
	return &meshIngressReconciler{
		base: reconcile.NewReconciler(client.BaseClient()),
	}
}

type meshIngressReconciler struct {
	base reconcile.Reconciler
}

func (r *meshIngressReconciler) Reconcile(namespace string, desiredResources MeshIngressList, transition TransitionMeshIngressFunc, opts clients.ListOpts) error {
	opts = opts.WithDefaults()
	opts.Ctx = contextutils.WithLogger(opts.Ctx, "meshIngress_reconciler")
	var transitionResources reconcile.TransitionResourcesFunc
	if transition != nil {
		transitionResources = func(original, desired resources.Resource) (bool, error) {
			return transition(original.(*MeshIngress), desired.(*MeshIngress))
		}
	}
	return r.base.Reconcile(namespace, meshIngresssToResources(desiredResources), transitionResources, opts)
}
