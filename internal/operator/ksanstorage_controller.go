/*
Copyright 2024.

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

package operator

import (
	"context"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	ksanv1alpha1 "openshift/ksan-operator/api/v1alpha1"
)

// KSANStorageReconciler reconciles a KSANStorage object
type KSANStorageReconciler struct {
	client.Client
	Scheme            *runtime.Scheme
	OperatorNamespace string
	PodImage          string
}

// +kubebuilder:rbac:groups=ksan.openshift.io,resources=ksanstorages,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=ksan.openshift.io,resources=ksanstorages/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=ksan.openshift.io,resources=ksanstorages/finalizers,verbs=update
// +kubebuilder:rbac:groups=ksan.openshift.io,resources=ksannodes,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=nodes,verbs=get;list;watch;patch;update
// +kubebuilder:rbac:groups=core,resources=pods,verbs=get;list;watch;delete
// +kubebuilder:rbac:groups=apps,resources=daemonsets,verbs=get;list;watch;create;update;patch;delete

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the KSANStorage object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.18.4/pkg/reconcile
func (r *KSANStorageReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	storage := &ksanv1alpha1.KSANStorage{}
	err := r.Client.Get(ctx, req.NamespacedName, storage)
	if err != nil {
		return ctrl.Result{}, err
	}

	nodes, err := r.nodesMatchingSelector(ctx, storage.Spec.NodeSelector)
	if err != nil {
		return ctrl.Result{}, err
	}

	err = r.ensureKSANNodes(ctx, storage.Spec, nodes, r.OperatorNamespace)
	if err != nil {
		return ctrl.Result{}, err
	}

	err = r.createNodeDaemon(ctx, r.OperatorNamespace)
	if err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *KSANStorageReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&ksanv1alpha1.KSANStorage{}).
		Complete(r)
}
