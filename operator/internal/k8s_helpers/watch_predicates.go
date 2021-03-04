package k8s_helpers

import (
	appsv1 "k8s.io/api/apps/v1"
	"sigs.k8s.io/controller-runtime/pkg/event"
)

type DaemonsetNewNodePredicate struct {
}

func (p DaemonsetNewNodePredicate) Create(e event.CreateEvent) bool {
	return false
}

func (p DaemonsetNewNodePredicate) Delete(e event.DeleteEvent) bool {
	return false
}

// Update watches the daemonset desirednumberscheduled
func (p DaemonsetNewNodePredicate) Update(e event.UpdateEvent) bool {
	oldDs, ok := e.ObjectOld.(*appsv1.DaemonSet)
	// If not ok then this is not a daemonset event
	if !ok {
		return false
	}

	newDs, ok := e.ObjectNew.(*appsv1.DaemonSet)
	// If not ok then this is not a daemonset event
	if !ok {
		return false
	}

	if newDs.Status.DesiredNumberScheduled > oldDs.Status.DesiredNumberScheduled {
		return true
	}
	return false
}

func (p DaemonsetNewNodePredicate) Generic(e event.GenericEvent) bool {
	return false
}
