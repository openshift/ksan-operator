package controller

import (
	"context"
	v1 "k8s.io/api/core/v1"
	corev1helper "k8s.io/component-helpers/scheduling/corev1"
)

func (r *KSANStorageReconciler) nodesMatchingSelector(ctx context.Context, nodeSelector *v1.NodeSelector) ([]v1.Node, error) {
	nodes := &v1.NodeList{}
	err := r.Client.List(ctx, nodes)
	if err != nil {
		return nil, err
	}

	if nodeSelector == nil {
		return nodes.Items, nil
	}

	res := make([]v1.Node, 0)
	for _, node := range nodes.Items {
		ok, err := corev1helper.MatchNodeSelectorTerms(&node, nodeSelector)
		if err != nil {
			return nil, err
		}
		if ok {
			res = append(res, node)
		}
	}
	return res, nil
}
