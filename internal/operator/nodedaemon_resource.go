package operator

import (
	"context"
	"fmt"
	"os"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/utils/ptr"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

func (r *KSANStorageReconciler) createNodeDaemon(ctx context.Context, namespace string) error {
	if r.PodImage == "" {
		podImage, err := podImage(ctx, r.Client, namespace)
		if err != nil {
			return err
		}
		r.PodImage = podImage
	}

	resourceRequirements := corev1.ResourceRequirements{
		Requests: corev1.ResourceList{
			corev1.ResourceCPU:    resource.MustParse(nodeDaemonCPURequest),
			corev1.ResourceMemory: resource.MustParse(nodeDaemonMemRequest),
		},
	}

	labels := map[string]string{
		appKubernetesNameLabel:      nodeDaemonNameLabelVal,
		appKubernetesManagedByLabel: nodeDaemonManagedByLabelVal,
		appKubernetesPartOfLabel:    nodeDaemonPartOfLabelVal,
		appKubernetesComponentLabel: nodeDaemonNameLabelVal,
	}

	d := &appsv1.DaemonSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "ksan-storage-nodedaemon",
			Namespace: namespace,
			Labels:    labels,
		},
		Spec: appsv1.DaemonSetSpec{
			Selector: &metav1.LabelSelector{MatchLabels: labels},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: corev1.PodSpec{
					TerminationGracePeriodSeconds: ptr.To(int64(30)),
					HostPID:                       true,
					ServiceAccountName:            nodeDaemonServiceAccountName,
					PriorityClassName:             priorityClassNameUserCritical,
					Containers: []corev1.Container{
						{
							Name:    nodeDaemonContainerName,
							Image:   r.PodImage,
							Command: []string{"/ksan-storage", "nodedaemon"},
							SecurityContext: &corev1.SecurityContext{
								Privileged: ptr.To(true),
								RunAsUser:  ptr.To(int64(0)),
							},
							Ports: []corev1.ContainerPort{
								{Name: "healthz",
									ContainerPort: 8081,
									Protocol:      corev1.ProtocolTCP,
								},
							},
							StartupProbe: &corev1.Probe{
								ProbeHandler: corev1.ProbeHandler{
									HTTPGet: &corev1.HTTPGetAction{Path: "/healthz",
										Port: intstr.FromString("healthz")}},
								FailureThreshold:    60, // 60*10 = 600s / 10 min for long startup due to large volume group initialization
								InitialDelaySeconds: 2,
								TimeoutSeconds:      2,
								PeriodSeconds:       10},
							LivenessProbe: &corev1.Probe{
								ProbeHandler: corev1.ProbeHandler{
									HTTPGet: &corev1.HTTPGetAction{Path: "/healthz",
										Port: intstr.FromString("healthz")}},
								FailureThreshold:    3,
								InitialDelaySeconds: 1,
								TimeoutSeconds:      1,
								PeriodSeconds:       30},
							ReadinessProbe: &corev1.Probe{
								ProbeHandler: corev1.ProbeHandler{
									HTTPGet: &corev1.HTTPGetAction{Path: "/readyz",
										Port: intstr.FromString("healthz")}},
								FailureThreshold:    3,
								InitialDelaySeconds: 1,
								TimeoutSeconds:      1,
								PeriodSeconds:       60,
							},
							Resources: resourceRequirements,
							Env: []corev1.EnvVar{
								{
									Name: "NODE_NAME",
									ValueFrom: &corev1.EnvVarSource{
										FieldRef: &corev1.ObjectFieldSelector{
											FieldPath: "spec.nodeName",
										},
									},
								},
								{
									Name: "NAMESPACE",
									ValueFrom: &corev1.EnvVarSource{
										FieldRef: &corev1.ObjectFieldSelector{
											FieldPath: "metadata.namespace",
										},
									},
								},
								{
									Name: "NAME",
									ValueFrom: &corev1.EnvVarSource{
										FieldRef: &corev1.ObjectFieldSelector{
											FieldPath: "metadata.name",
										},
									},
								},
							},
							TerminationMessagePolicy: corev1.TerminationMessageFallbackToLogsOnError,
						},
					},
				},
			},
		},
	}
	_, err := controllerutil.CreateOrUpdate(ctx, r.Client, d, func() error {
		return nil
	})
	return err
}

func podImage(ctx context.Context, c client.Client, namespace string) (string, error) {
	podName := os.Getenv("NAME")
	if podName == "" {
		return podName, fmt.Errorf("failed to get pod name env variable, %s env variable is not set", "NAME")
	}

	pod := &corev1.Pod{}
	if err := c.Get(ctx, types.NamespacedName{Name: podName, Namespace: namespace}, pod); err != nil {
		return podName, fmt.Errorf("failed to get pod %s: %w", podName, err)
	}

	for _, c := range pod.Spec.Containers {
		if c.Name == "manager" {
			return c.Image, nil
		}
	}

	return podName, fmt.Errorf("failed to get container image for %s in pod %s", "operator", podName)
}
