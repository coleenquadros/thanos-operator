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

package controller

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-logr/logr"

	monitoringthanosiov1alpha1 "github.com/thanos-community/thanos-operator/api/v1alpha1"
	"github.com/thanos-community/thanos-operator/internal/pkg/handlers"
	"github.com/thanos-community/thanos-operator/internal/pkg/manifests"
	manifestreceive "github.com/thanos-community/thanos-operator/internal/pkg/manifests/receive"
	manifestsstore "github.com/thanos-community/thanos-operator/internal/pkg/manifests/store"
	controllermetrics "github.com/thanos-community/thanos-operator/internal/pkg/metrics"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	discoveryv1 "k8s.io/api/discovery/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/record"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

const (
	receiveFinalizer = "monitoring.thanos.io/receive-finalizer"
)

// ThanosReceiveReconciler reconciles a ThanosReceive object
type ThanosReceiveReconciler struct {
	client.Client
	Scheme *runtime.Scheme

	logger   logr.Logger
	metrics  controllermetrics.ThanosReceiveMetrics
	recorder record.EventRecorder

	handler *handlers.Handler
}

// NewThanosReceiveReconciler returns a reconciler for ThanosReceive resources.
func NewThanosReceiveReconciler(instrumentationConf InstrumentationConfig, client client.Client, scheme *runtime.Scheme) *ThanosReceiveReconciler {
	return &ThanosReceiveReconciler{
		Client:   client,
		Scheme:   scheme,
		logger:   instrumentationConf.Logger,
		metrics:  controllermetrics.NewThanosReceiveMetrics(instrumentationConf.MetricsRegistry, instrumentationConf.BaseMetrics),
		recorder: instrumentationConf.EventRecorder,
		handler:  handlers.NewHandler(client, scheme, instrumentationConf.Logger),
	}
}

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.17.3/pkg/reconcile
func (r *ThanosReceiveReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	r.metrics.ReconciliationsTotal.WithLabelValues(manifestreceive.Name).Inc()

	// Fetch the ThanosReceive instance to validate it is applied on the cluster.
	receiver := &monitoringthanosiov1alpha1.ThanosReceive{}
	err := r.Get(ctx, req.NamespacedName, receiver)
	if err != nil {
		r.metrics.ClientErrorsTotal.WithLabelValues(manifestreceive.Name).Inc()
		if apierrors.IsNotFound(err) {
			r.logger.Info("thanos receive resource not found. ignoring since object may be deleted")
			return ctrl.Result{}, nil
		}
		r.logger.Error(err, "failed to get ThanosReceive")
		r.metrics.ReconciliationsFailedTotal.WithLabelValues(manifestreceive.Name).Inc()
		r.recorder.Event(receiver, corev1.EventTypeWarning, "GetFailed", "Failed to get ThanosReceive resource")
		return ctrl.Result{}, err
	}

	// handle object being deleted - inferred from the existence of DeletionTimestamp
	if !receiver.GetDeletionTimestamp().IsZero() {
		return r.handleDeletionTimestamp(receiver)
	}

	err = r.syncResources(ctx, *receiver)
	if err != nil {
		r.metrics.ReconciliationsFailedTotal.WithLabelValues(manifestreceive.Name).Inc()
		r.recorder.Event(receiver, corev1.EventTypeWarning, "SyncFailed", fmt.Sprintf("Failed to sync resources: %v", err))
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

// +kubebuilder:rbac:groups=monitoring.thanos.io,resources=thanosreceives,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=monitoring.thanos.io,resources=thanosreceives/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=monitoring.thanos.io,resources=thanosreceives/finalizers,verbs=update
// +kubebuilder:rbac:groups=apps,resources=statefulsets;deployments,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups="",resources=services;configmaps;serviceaccounts,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups="discovery.k8s.io",resources=endpointslices,verbs=get;list;watch

// SetupWithManager sets up the controller with the Manager.
func (r *ThanosReceiveReconciler) SetupWithManager(mgr ctrl.Manager) error {
	bld := ctrl.NewControllerManagedBy(mgr)
	err := r.buildController(*bld)
	if err != nil {
		r.recorder.Event(&monitoringthanosiov1alpha1.ThanosReceive{}, corev1.EventTypeWarning, "SetupFailed", fmt.Sprintf("Failed to set up controller: %v", err))
		return err
	}

	return nil
}

// buildController sets up the controller with the Manager.
func (r *ThanosReceiveReconciler) buildController(bld builder.Builder) error {
	// add a selector to watch for the endpointslices that are owned by the ThanosReceive ingest Service(s).
	endpointSliceLS := metav1.LabelSelector{
		MatchLabels: map[string]string{manifests.ComponentLabel: manifestreceive.IngestComponentName},
	}
	endpointSlicePredicate, err := predicate.LabelSelectorPredicate(endpointSliceLS)
	if err != nil {
		return err
	}

	bld.
		For(&monitoringthanosiov1alpha1.ThanosReceive{}).
		Owns(&corev1.ConfigMap{}).
		Owns(&corev1.ServiceAccount{}).
		Owns(&corev1.Service{}).
		Owns(&appsv1.Deployment{}).
		Owns(&appsv1.StatefulSet{}).
		Watches(
			&discoveryv1.EndpointSlice{},
			r.enqueueForEndpointSlice(r.Client),
			builder.WithPredicates(predicate.GenerationChangedPredicate{}, endpointSlicePredicate),
		)

	return bld.Complete(r)
}

// syncResources syncs the resources for the ThanosReceive resource.
// It creates or updates the resources for the hashrings and the router.
func (r *ThanosReceiveReconciler) syncResources(ctx context.Context, receiver monitoringthanosiov1alpha1.ThanosReceive) error {
	var errCount int
	shardedObjects := r.buildIngesterHashrings(receiver)

	// todo - we need to prune any orphaned resources here at this point

	for _, shardObjs := range shardedObjects {
		errCount += r.handler.CreateOrUpdate(ctx, receiver.GetNamespace(), &receiver, shardObjs)
	}

	if errCount > 0 {
		r.metrics.ClientErrorsTotal.WithLabelValues(manifestsstore.Name).Add(float64(errCount))
		return fmt.Errorf("failed to create or update %d resources for receive hashring(s)", errCount)
	}

	var objs []client.Object
	hashringConf, err := r.buildHashringConfig(ctx, receiver)
	if err != nil {
		if !errors.Is(err, manifestreceive.ErrHashringsEmpty) {
			r.recorder.Event(&receiver, corev1.EventTypeWarning, "HashringConfigBuildFailed", fmt.Sprintf("Failed to build hashring configuration: %v", err))
			return fmt.Errorf("failed to build hashring configuration: %w", err)
		}
		// we can create the config map even if there are no hashrings
		objs = append(objs, hashringConf)
	} else {
		objs = append(objs, hashringConf)
		// bring up the router components only if there are ready hashrings to avoid crash looping the router
		objs = append(objs, r.buildRouter(receiver)...)
	}

	if errs := r.handler.CreateOrUpdate(ctx, receiver.GetNamespace(), &receiver, objs); errs > 0 {
		r.metrics.ClientErrorsTotal.WithLabelValues(manifestreceive.Name).Add(float64(errs))
		return fmt.Errorf("failed to create or update %d resources for the receive router", errs)
	}

	return nil
}

// buildIngesterHashrings builds the ingesters for the ThanosReceive resource.
// It returns a map of slices of client.Object that represents the desired state of the Thanos Ingester resources.
// Each key represents a named ingester and each slice value represents the resources for that ingester.
func (r *ThanosReceiveReconciler) buildIngesterHashrings(receiver monitoringthanosiov1alpha1.ThanosReceive) map[string][]client.Object {
	opts := r.specToIngestOptions(receiver)
	shardedObjects := make(map[string][]client.Object, len(opts))

	for hashring, opt := range opts {
		for _, ingester := range receiver.Spec.Ingester.Hashrings {
			if ingester.Paused != nil && *ingester.Paused && hashring == ingester.Name {
				r.logger.Info("ingester is paused", "ingester", hashring)
				r.recorder.Event(&receiver, corev1.EventTypeNormal, "Paused",
					"Reconciliation is paused for ThanosReceive resource - hashring "+hashring)
				continue
			}
		}

		storeObjs := manifestreceive.BuildIngester(opt)
		shardedObjects[hashring] = append(shardedObjects[hashring], storeObjs...)
	}
	return shardedObjects
}

func (r *ThanosReceiveReconciler) specToIngestOptions(receiver monitoringthanosiov1alpha1.ThanosReceive) map[string]manifestreceive.IngesterOptions {
	opts := make(map[string]manifestreceive.IngesterOptions, len(receiver.Spec.Ingester.Hashrings))
	for _, v := range receiver.Spec.Ingester.Hashrings {
		opts[v.Name] = receiverV1Alpha1ToIngesterOptions(receiver, v)
	}
	return opts
}

// buildHashringConfig builds the hashring configuration for the ThanosReceive resource.
func (r *ThanosReceiveReconciler) buildHashringConfig(ctx context.Context, receiver monitoringthanosiov1alpha1.ThanosReceive) (client.Object, error) {
	cm := &corev1.ConfigMap{}
	name := ReceiveRouterNameFromParent(receiver.GetName())
	err := r.Client.Get(ctx, client.ObjectKey{Namespace: receiver.GetNamespace(), Name: name}, cm)
	if err != nil {
		if !apierrors.IsNotFound(err) {
			return nil, fmt.Errorf("failed to get config map for resource %s: %w", name, err)
		}
	}

	opts := manifestreceive.HashringOptions{
		Options: manifests.Options{
			Name:      name,
			Namespace: receiver.GetNamespace(),
			Labels:    receiver.GetLabels(),
		},
		DesiredReplicationFactor: receiver.Spec.Router.ReplicationFactor,
		HashringSettings:         make(map[string]manifestreceive.HashringMeta, len(receiver.Spec.Ingester.Hashrings)),
	}

	totalHashrings := len(receiver.Spec.Ingester.Hashrings)
	for i, hashring := range receiver.Spec.Ingester.Hashrings {
		labelValue := ReceiveIngesterNameFromParent(receiver.GetName(), hashring.Name)
		// kubernetes sets this label on the endpoint slices - we want to match the generated name
		selectorListOpt := client.MatchingLabels{discoveryv1.LabelServiceName: labelValue}

		eps := discoveryv1.EndpointSliceList{}
		if err = r.Client.List(ctx, &eps, selectorListOpt, client.InNamespace(receiver.GetNamespace())); err != nil {
			return nil, fmt.Errorf("failed to list endpoint slices for resource %s: %w", receiver.GetName(), err)
		}

		opts.HashringSettings[labelValue] = manifestreceive.HashringMeta{
			DesiredReplicasReplicas:  hashring.Replicas,
			OriginalName:             hashring.Name,
			Tenants:                  hashring.Tenants,
			TenantMatcherType:        manifestreceive.TenantMatcher(hashring.TenantMatcherType),
			AssociatedEndpointSlices: eps,
			// set the priority by slice order for now
			Priority: totalHashrings - i,
		}
	}

	r.metrics.HashringsConfigured.WithLabelValues(receiver.GetName(), receiver.GetNamespace()).Set(float64(totalHashrings))
	r.recorder.Event(&receiver, corev1.EventTypeNormal, "HashringConfigBuilt", fmt.Sprintf("Built hashring configuration with %d hashrings", totalHashrings))
	return manifestreceive.BuildHashrings(r.logger, cm, opts)
}

// build hashring builds out the ingesters for the ThanosReceive resource.
func (r *ThanosReceiveReconciler) buildRouter(receiver monitoringthanosiov1alpha1.ThanosReceive) []client.Object {
	if receiver.Spec.Router.Paused != nil && *receiver.Spec.Router.Paused {
		r.logger.Info("router is paused")
		r.recorder.Event(&receiver, corev1.EventTypeNormal, "Paused",
			"Reconciliation is paused for ThanosReceive resource - router")
		return nil
	}
	opts := receiverV1Alpha1ToRouterOptions(receiver)
	return manifestreceive.BuildRouter(opts)
}

func (r *ThanosReceiveReconciler) handleDeletionTimestamp(receiveHashring *monitoringthanosiov1alpha1.ThanosReceive) (ctrl.Result, error) {
	if controllerutil.ContainsFinalizer(receiveHashring, receiveFinalizer) {
		r.logger.Info("performing Finalizer Operations for ThanosReceiveHashring before delete CR")

		r.recorder.Event(receiveHashring, "Warning", "Deleting",
			fmt.Sprintf("Custom Resource %s is being deleted from the namespace %s",
				receiveHashring.Name,
				receiveHashring.Namespace))
	}
	return ctrl.Result{}, nil
}

// enqueueForEndpointSlice enqueues requests for the ThanosReceive resource when an EndpointSlice event is triggered.
func (r *ThanosReceiveReconciler) enqueueForEndpointSlice(c client.Client) handler.EventHandler {
	return handler.EnqueueRequestsFromMapFunc(func(ctx context.Context, obj client.Object) []reconcile.Request {

		if len(obj.GetOwnerReferences()) != 1 || obj.GetOwnerReferences()[0].Kind != "Service" {
			return nil
		}

		svc := &corev1.Service{}
		if err := c.Get(ctx, types.NamespacedName{Namespace: obj.GetNamespace(), Name: obj.GetOwnerReferences()[0].Name}, svc); err != nil {
			return nil
		}

		if len(svc.GetOwnerReferences()) != 1 || svc.GetOwnerReferences()[0].Kind != "ThanosReceive" {
			return nil
		}

		r.metrics.EndpointWatchesReconciliationsTotal.Inc()
		return []reconcile.Request{
			{
				NamespacedName: types.NamespacedName{
					Namespace: obj.GetNamespace(),
					Name:      svc.GetOwnerReferences()[0].Name,
				},
			},
		}
	})
}
