package operator

import (
	"context"
	"math/rand"
	ksanv1alpha1 "openshift/ksan-operator/api/v1alpha1"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	controllerruntime "sigs.k8s.io/controller-runtime"
)

func (r *KSANStorageReconciler) ensureKSANNodes(ctx context.Context, storageSpec ksanv1alpha1.KSANStorageSpec, nodes []v1.Node, namespace string) error {
	for _, node := range nodes {
		ksanNode := &ksanv1alpha1.KSANNode{
			ObjectMeta: metav1.ObjectMeta{
				Name:      node.Name,
				Namespace: namespace,
			},
		}

		_, err := controllerruntime.CreateOrUpdate(ctx, r.Client, ksanNode, func() error {
			if ksanNode.Spec.Storage == nil {
				ksanNode.Spec.Storage = make(map[string]ksanv1alpha1.KSANNodeStorage)
			}

			if ksanNode.Spec.HostID == 0 {
				//TODO: replace with management
				ksanNode.Spec.HostID = 1 + rand.Intn(1999)
			}

			ksanNode.Spec.Storage[storageSpec.VolumeGroupName] = ksanv1alpha1.KSANNodeStorage{
				VolumeGroupName: storageSpec.VolumeGroupName,
				Devices:         storageSpec.DevicesConfig.Devices,
			}

			return nil
		})
		if err != nil {
			return err
		}
	}

	return nil
}
