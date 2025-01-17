package nodedaemon

import (
	"context"
	"openshift/ksan-operator/api/v1alpha1"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

type Reconciler struct {
	Client   client.Client
	Scheme   *runtime.Scheme
	NodeName string
}

func (r *Reconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha1.KSANNode{}).
		WithEventFilter(predicate.NewPredicateFuncs(func(object client.Object) bool {
			// each daemonset processes only node related configuration
			return object.GetName() == r.NodeName
		})).
		Complete(r)
}

func (r *Reconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	return ctrl.Result{}, nil
}
