/*
Copyright 2024 MatanMagen.

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

package controller

import (
	"context"

	api "github.com/MatanMagen/matanmagen-crd.git/api/v1alpha1"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	apiv1alpha1 "github.com/MatanMagen/matanmagen-crd.git/api/v1alpha1"
)

// MatanAppReconciler reconciles a MatanApp object
type MatanAppReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=api.core.matanmagen.io,resources=matanapps,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=api.core.matanmagen.io,resources=matanapps/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=api.core.matanmagen.io,resources=matanapps/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the MatanApp object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.18.2/pkg/reconcile
func (r *MatanAppReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {

	matanApp := &api.MatanApp{}
	err := r.Get(context.TODO(), req.NamespacedName, matanApp)
	if err != nil {
		// Handle error
	}

	imageName := matanApp.Spec.ImageName
	// Define a new Job object
	job := newJobForCR(matanApp, imageName)
	// Check if this Job already exists
	foundJob := &batchv1.Job{}
	err = r.Get(context.TODO(), types.NamespacedName{Name: job.Name, Namespace: job.Namespace}, foundJob)
	if err != nil {
		// If Job doesn't exist, create it
		err = r.Create(context.TODO(), job)
		if err != nil {
			// Handle error
		}
	}

	// Define a new Secret object
	secret := newSecretForCR(matanApp)
	// Check if this Secret already exists
	foundSecret := &corev1.Secret{}
	err = r.Get(context.TODO(), types.NamespacedName{Name: secret.Name, Namespace: secret.Namespace}, foundSecret)
	if err != nil {
		// If Secret doesn't exist, create it
		err = r.Create(context.TODO(), secret)
		if err != nil {
			// Handle error
		}
	}

	// Define a new ConfigMap object
	configMap := newConfigMapForCR(matanApp)
	// Check if this ConfigMap already exists
	foundConfigMap := &corev1.ConfigMap{}
	err = r.Get(context.TODO(), types.NamespacedName{Name: configMap.Name, Namespace: configMap.Namespace}, foundConfigMap)
	if err != nil {
		// If ConfigMap doesn't exist, create it
		err = r.Create(context.TODO(), configMap)
		if err != nil {
			// Handle error
		}
	}

	return ctrl.Result{}, nil
}

func newJobForCR(cr *apiv1alpha1.MatanApp, imageName string) *batchv1.Job {
	labels := map[string]string{
		"app": cr.Name,
	}
	return &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Name + "-job",
			Namespace: cr.Namespace,
			Labels:    labels,
		},
		Spec: batchv1.JobSpec{
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: corev1.PodSpec{
					RestartPolicy: corev1.RestartPolicyNever,
					Containers: []corev1.Container{
						{
							Name:            cr.Name + "-container",
							Image:           imageName,
							Command:         []string{"echo", "Hello, World!"},
							ImagePullPolicy: corev1.PullIfNotPresent,
							Resources: corev1.ResourceRequirements{
								Requests: corev1.ResourceList{
									corev1.ResourceCPU:    resource.MustParse("100m"),
									corev1.ResourceMemory: resource.MustParse("100Mi"),
								},
								Limits: corev1.ResourceList{
									corev1.ResourceCPU:    resource.MustParse("200m"),
									corev1.ResourceMemory: resource.MustParse("200Mi"),
								},
							},
							Env: []corev1.EnvVar{
								{
									Name:  "MY_ENV_VAR",
									Value: "my-value",
								},
							},
						},
					},
				},
			},
		},
	}
}

func newSecretForCR(cr *apiv1alpha1.MatanApp) *corev1.Secret {
	labels := map[string]string{
		"app": cr.Name,
	}
	return &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Name + "-secret",
			Namespace: cr.Namespace,
			Labels:    labels,
		},
		// Assuming that cr.Spec.Secret is of type corev1.Secret
		Type: cr.Spec.Secret.Type,
		Data: cr.Spec.Secret.Data,
	}
}

func newConfigMapForCR(cr *apiv1alpha1.MatanApp) *corev1.ConfigMap {
	labels := map[string]string{
		"app": cr.Name,
	}
	return &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Name + "-configmap",
			Namespace: cr.Namespace,
			Labels:    labels,
		},
		Data: cr.Spec.ConfigMap.Data,
	}
}

// SetupWithManager sets up the controller with the Manager.
func (r *MatanAppReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&apiv1alpha1.MatanApp{}).
		Complete(r)
}
