package controllers

import (
	"context"
	"testing"

	ihpav1beta2 "github.com/cyberagent-oss/intelligent-hpa/ihpa-controller/api/v1beta2"
	logtesting "github.com/go-logr/logr/testing"
	appsv1 "k8s.io/api/apps/v1"
	autoscalingv2beta2 "k8s.io/api/autoscaling/v2beta2"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestIHPAReconcileScaleBackendHPA(t *testing.T) {
	ctx := context.Background()
	scheme := testReconcileScheme(t)
	ihpa := testReconcileIHPA(ihpav1beta2.ScaleBackendSpec{})
	existingScaledObject := newScaledObjectUnstructured("default", "ihpa-nginx")
	existingScaledObject.Object["spec"] = map[string]interface{}{
		"scaleTargetRef": map[string]interface{}{"name": "nginx"},
		"triggers":       []interface{}{map[string]interface{}{"type": "cpu"}},
	}

	reconciler := &IntelligentHorizontalPodAutoscalerReconciler{
		Client:        fake.NewFakeClientWithScheme(scheme, testReconcileKubeSystemNamespace(), testReconcileDeployment(), ihpa, existingScaledObject),
		Log:           logtesting.NullLogger{},
		Scheme:        scheme,
		fittingJobMap: make(map[string]map[string]struct{}),
	}

	if _, err := reconciler.Reconcile(ctrl.Request{NamespacedName: types.NamespacedName{Namespace: "default", Name: "nginx"}}); err != nil {
		t.Fatalf("reconcile failed: %v", err)
	}

	hpa := newHorizontalPodAutoscalerUnstructured("default", "ihpa-nginx")
	if err := reconciler.Get(ctx, types.NamespacedName{Namespace: "default", Name: "ihpa-nginx"}, hpa); err != nil {
		t.Fatalf("expected hpa to exist: %v", err)
	}
	scaledObject := newScaledObjectUnstructured("default", "ihpa-nginx")
	if err := reconciler.Get(ctx, types.NamespacedName{Namespace: "default", Name: "ihpa-nginx"}, scaledObject); !apierrors.IsNotFound(err) {
		t.Fatalf("expected scaledobject to be deleted, got err=%v object=%#v", err, scaledObject)
	}
}

func TestIHPAReconcileScaleBackendKEDA(t *testing.T) {
	ctx := context.Background()
	scheme := testReconcileScheme(t)
	ihpa := testReconcileIHPA(ihpav1beta2.ScaleBackendSpec{
		Type: ihpav1beta2.ScaleBackendTypeKEDA,
		KEDA: &ihpav1beta2.KEDAScaleBackendSpec{
			Triggers: []ihpav1beta2.KEDATriggerSpec{
				{
					Type:       "cpu",
					MetricType: "Utilization",
					Metadata: map[string]string{
						"value": "50",
					},
				},
			},
		},
	})
	existingHPA := newHorizontalPodAutoscalerUnstructured("default", "ihpa-nginx")
	existingHPA.Object["spec"] = map[string]interface{}{
		"scaleTargetRef": map[string]interface{}{
			"apiVersion": "apps/v1",
			"kind":       "Deployment",
			"name":       "nginx",
		},
	}

	reconciler := &IntelligentHorizontalPodAutoscalerReconciler{
		Client:        fake.NewFakeClientWithScheme(scheme, testReconcileKubeSystemNamespace(), testReconcileDeployment(), ihpa, existingHPA),
		Log:           logtesting.NullLogger{},
		Scheme:        scheme,
		fittingJobMap: make(map[string]map[string]struct{}),
	}

	if _, err := reconciler.Reconcile(ctrl.Request{NamespacedName: types.NamespacedName{Namespace: "default", Name: "nginx"}}); err != nil {
		t.Fatalf("reconcile failed: %v", err)
	}

	scaledObject := newScaledObjectUnstructured("default", "ihpa-nginx")
	if err := reconciler.Get(ctx, types.NamespacedName{Namespace: "default", Name: "ihpa-nginx"}, scaledObject); err != nil {
		t.Fatalf("expected scaledobject to exist: %v", err)
	}
	spec := scaledObject.Object["spec"].(map[string]interface{})
	if _, ok := spec["triggers"]; !ok {
		t.Fatalf("expected scaledobject triggers: %#v", spec)
	}

	hpa := newHorizontalPodAutoscalerUnstructured("default", "ihpa-nginx")
	if err := reconciler.Get(ctx, types.NamespacedName{Namespace: "default", Name: "ihpa-nginx"}, hpa); !apierrors.IsNotFound(err) {
		t.Fatalf("expected hpa to be deleted, got err=%v object=%#v", err, hpa)
	}
}

func testReconcileScheme(t *testing.T) *runtime.Scheme {
	t.Helper()
	scheme := runtime.NewScheme()
	for _, addToScheme := range []func(*runtime.Scheme) error{
		corev1.AddToScheme,
		appsv1.AddToScheme,
		autoscalingv2beta2.AddToScheme,
		rbacv1.AddToScheme,
		ihpav1beta2.AddToScheme,
	} {
		if err := addToScheme(scheme); err != nil {
			t.Fatal(err)
		}
	}
	return scheme
}

func testReconcileIHPA(scaleBackend ihpav1beta2.ScaleBackendSpec) *ihpav1beta2.IntelligentHorizontalPodAutoscaler {
	return &ihpav1beta2.IntelligentHorizontalPodAutoscaler{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "ihpa.ake.cyberagent.co.jp/v1beta2",
			Kind:       "IntelligentHorizontalPodAutoscaler",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "nginx",
			Namespace: "default",
			UID:       "9fc642f3-bb9d-404d-afd5-d04f1d6149ad",
		},
		Spec: ihpav1beta2.IntelligentHorizontalPodAutoscalerSpec{
			ScaleBackend: scaleBackend,
			HorizontalPodAutoscalerTemplate: ihpav1beta2.ExtendedHorizontalPodAutoscalerTemplateSpec{
				Spec: ihpav1beta2.ExtendedHorizontalPodAutoscalerSpec{
					ScaleTargetRef: autoscalingv2beta2.CrossVersionObjectReference{
						APIVersion: "apps/v1",
						Kind:       "Deployment",
						Name:       "nginx",
					},
					MinReplicas: func(i int32) *int32 { return &i }(1),
					MaxReplicas: 5,
					Metrics: []ihpav1beta2.ExtendedMetricSpec{
						{
							Type: "Resource",
							Resource: &autoscalingv2beta2.ResourceMetricSource{
								Name: "cpu",
								Target: autoscalingv2beta2.MetricTarget{
									Type:               "Utilization",
									AverageUtilization: func(i int32) *int32 { return &i }(50),
								},
							},
							FittingJobPatchSpec: ihpav1beta2.FittingJobPatchSpec{
								Seasonality: "daily",
								ExecuteOn:   4,
							},
						},
					},
				},
			},
			EstimatorPatchSpec: ihpav1beta2.EstimatorPatchSpec{
				Mode:       "adjust",
				GapMinutes: 10,
			},
			MetricProvider: ihpav1beta2.MetricProvider{
				Name: "datadog",
				ProviderSource: ihpav1beta2.ProviderSource{
					Datadog: &ihpav1beta2.DatadogProviderSource{
						APIKey: "xxx",
						APPKey: "yyy",
					},
				},
			},
		},
	}
}

func testReconcileKubeSystemNamespace() *corev1.Namespace {
	return &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: "kube-system",
			UID:  "46f9e396-d3c4-4103-a807-49054f47bbfb",
		},
	}
}

func testReconcileDeployment() *appsv1.Deployment {
	return &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "nginx",
			Namespace: "default",
		},
		Spec: appsv1.DeploymentSpec{
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name: "nginx",
							Resources: corev1.ResourceRequirements{
								Requests: corev1.ResourceList{
									corev1.ResourceCPU: *resource.NewScaledQuantity(500, resource.Milli),
								},
							},
						},
					},
				},
			},
		},
	}
}
