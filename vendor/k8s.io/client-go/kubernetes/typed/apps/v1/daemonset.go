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

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"

	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	appsv1 "k8s.io/client-go/applyconfigurations/apps/v1"
	gentype "k8s.io/client-go/gentype"
	scheme "k8s.io/client-go/kubernetes/scheme"
)

// DaemonSetsGetter has a method to return a DaemonSetInterface.
// A group's client should implement this interface.
type DaemonSetsGetter interface {
	DaemonSets(namespace string) DaemonSetInterface
}

// DaemonSetInterface has methods to work with DaemonSet resources.
type DaemonSetInterface interface {
	Create(ctx context.Context, daemonSet *v1.DaemonSet, opts metav1.CreateOptions) (*v1.DaemonSet, error)
	Update(ctx context.Context, daemonSet *v1.DaemonSet, opts metav1.UpdateOptions) (*v1.DaemonSet, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, daemonSet *v1.DaemonSet, opts metav1.UpdateOptions) (*v1.DaemonSet, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.DaemonSet, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.DaemonSetList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.DaemonSet, err error)
	Apply(ctx context.Context, daemonSet *appsv1.DaemonSetApplyConfiguration, opts metav1.ApplyOptions) (result *v1.DaemonSet, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, daemonSet *appsv1.DaemonSetApplyConfiguration, opts metav1.ApplyOptions) (result *v1.DaemonSet, err error)
	DaemonSetExpansion
}

// daemonSets implements DaemonSetInterface
type daemonSets struct {
	*gentype.ClientWithListAndApply[*v1.DaemonSet, *v1.DaemonSetList, *appsv1.DaemonSetApplyConfiguration]
}

// newDaemonSets returns a DaemonSets
func newDaemonSets(c *AppsV1Client, namespace string) *daemonSets {
	return &daemonSets{
		gentype.NewClientWithListAndApply[*v1.DaemonSet, *v1.DaemonSetList, *appsv1.DaemonSetApplyConfiguration](
			"daemonsets",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1.DaemonSet { return &v1.DaemonSet{} },
			func() *v1.DaemonSetList { return &v1.DaemonSetList{} }),
	}
}