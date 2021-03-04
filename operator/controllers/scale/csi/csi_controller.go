/*
Copyright 2021.

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

package scale

import (
	"context"

	ctrl "sigs.k8s.io/controller-runtime"

	scalev1 "github.com/IBM/ibm-spectrum-scale-csi/apis/scale/v1"
	"github.com/IBM/ibm-spectrum-scale-csi/operator/internal/template"
	"github.com/IBM/ibm-spectrum-scale-csi/operator/pkg/services"
)

// CSIReconciler reconciles a CSI object
type CSIReconciler services.Reconcile

// +kubebuilder:rbac:groups=scale.spectrum.ibm.com,resources=csis,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=scale.spectrum.ibm.com,resources=csis/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=scale.spectrum.ibm.com,resources=csis/finalizers,verbs=update

func (r *CSIReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := r.Log.WithValues("csi", req.NamespacedName)

	// Gather the CSI resource we are reconciling on
	getCr := r.NewGroup(
		func() (ctrl.Result, error) {
			err := r.Get(ctx, req.NamespacedName, cr)
			if err != nil {
				//TODO? client.IgnoreNotFound should be in own if?
				log.Error(err, "Failed to get ScaleCluster")
				return ctrl.Result{}, client.IgnoreNotFound(err)
			}

			return ctrl.Result{}, nil
		},
	)

	template.Execute(cr, "driver", daemonset)

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *CSIReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&scalev1.CSI{}).
		Complete(r)
}
