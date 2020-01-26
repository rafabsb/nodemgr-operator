package nodemgr

import (
	"context"
	rafabsbv1alpha1 "github.com/rafabsb/nodemgr-operator/pkg/apis/rafabsb/v1alpha1"
	"os"
	"reflect"

	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var (
	log              = logf.Log.WithName("controller_nodemgr")
	keyPreffixBase   = getEnv("KEY_PREFIX_BASE", "rafabsb.com")
	nodeKeyPreffix   = "node." + keyPreffixBase
	nodemgrFinalizer = "finalizer." + keyPreffixBase
)

// const nodemgrFinalizer = "finalizer.rafabsb.com"

// Add creates a new Nodemgr Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileNodemgr{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("nodemgr-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	src := &source.Kind{Type: &rafabsbv1alpha1.Nodemgr{}}
	h := &handler.EnqueueRequestsFromMapFunc{
		ToRequests: handler.ToRequestsFunc(func(a handler.MapObject) []reconcile.Request {
			return []reconcile.Request{
				{NamespacedName: types.NamespacedName{
					Name:      a.Meta.GetName(),
					Namespace: a.Meta.GetNamespace(),
				}},
				// {NamespacedName: types.NamespacedName{
				// 	Name:      a.Meta.GetName() + "-2",
				// 	Namespace: a.Object.GetObjectKind().GroupVersionKind().GroupKind().String(),
				// }},
			}
		}),
	}
	// Watch for changes to primary resource Nodemgr
	err = c.Watch(src, h)
	if err != nil {
		return err
	}

	// TODO(user): Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Pods and requeue the owner Nodemgr
	src = &source.Kind{Type: &corev1.Node{}}
	h = &handler.EnqueueRequestsFromMapFunc{
		ToRequests: handler.ToRequestsFunc(func(a handler.MapObject) []reconcile.Request {
			return []reconcile.Request{
				{NamespacedName: types.NamespacedName{
					Name:      a.Meta.GetName(),
					Namespace: a.Meta.GetNamespace(),
				}},
			}
		}),
	}

	pred := predicate.Funcs{
		UpdateFunc: func(e event.UpdateEvent) bool {
			return e.MetaOld.GetResourceVersion() != e.MetaNew.GetResourceVersion()
		},
		DeleteFunc: func(e event.DeleteEvent) bool {
			// Evaluates to false if the object has been confirmed deleted.
			return !e.DeleteStateUnknown
		},
	}

	err = c.Watch(src, h, pred)
	if err != nil {
		return err
	}

	return nil
}

// blank assignment to verify that ReconcileNodemgr implements reconcile.Reconciler
var _ reconcile.Reconciler = &ReconcileNodemgr{}

// ReconcileNodemgr reconciles a Nodemgr object
type ReconcileNodemgr struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a Nodemgr object and makes changes based on the state read
// and what is in the Nodemgr.Spec

// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileNodemgr) Reconcile(request reconcile.Request) (reconcile.Result, error) {

	reqLogger := log.WithValues("Request.Name", request.Name)
	// reqLogger.Info("Reconciling Nodemgr")

	// Fetch the Nodemgr instance
	nodemgr := &rafabsbv1alpha1.Nodemgr{}

	err := r.client.Get(context.TODO(), request.NamespacedName, nodemgr)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			reqLogger.Info("Nodemgr not found.")

			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}
	//create a instance to simplify the calculation
	nodemgrMetadata := rafabsbv1alpha1.NodemgrMetadata{
		Labels:      nodemgr.Spec.Labels,
		Annotations: nodemgr.Spec.Annotations,
		Taints:      nodemgr.Spec.Taints,
		KeyPrefix:   nodeKeyPreffix,
		Node:        nodemgr,
	}

	// Check if this Node exists
	node := &corev1.Node{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: nodemgr.Name}, node)
	if err != nil && errors.IsNotFound(err) {
		reqLogger.Info("Node not found. Check if the nodemgr object name matchs some node name.")
		nodemgr.Status.Phase = "not found"
		err := r.client.Status().Update(context.TODO(), nodemgr)
		if err != nil {
			return reconcile.Result{}, err
		}
		return reconcile.Result{}, nil
	} else if err != nil {
		return reconcile.Result{}, err
	}

	nodeMetadata := rafabsbv1alpha1.NodeMetadata{
		Labels:      node.GetLabels(),
		Annotations: node.GetAnnotations(),
		Taints:      node.Spec.Taints,
		KeyPrefix:   nodeKeyPreffix,
		Node:        node,
	}

	// Check if the Nodemgr instance is marked to be deleted, which is
	// indicated by the deletion timestamp being set.
	isNodemgrMarkedToBeDeleted := nodemgr.GetDeletionTimestamp() != nil
	if isNodemgrMarkedToBeDeleted {
		if contains(nodemgr.GetFinalizers(), nodemgrFinalizer) {
			// Run finalization logic for nodemgrFinalizer. If the
			// finalization logic fails, don't remove the finalizer so
			// that we can retry during the next reconciliation.
			if err := r.finalizeNodemgr(reqLogger, nodeMetadata); err != nil {
				return reconcile.Result{}, err
			}

			// Remove nodemgrFinalizer. Once all finalizers have been
			// removed, the object will be deleted.
			nodemgr.SetFinalizers(remove(nodemgr.GetFinalizers(), nodemgrFinalizer))
			err := r.client.Update(context.TODO(), nodemgr)
			if err != nil {
				return reconcile.Result{}, err
			}
		}
		return reconcile.Result{}, nil
	}

	// Add finalizer for this CR
	if !contains(nodemgr.GetFinalizers(), nodemgrFinalizer) {
		if err := r.addFinalizer(reqLogger, nodemgr); err != nil {
			return reconcile.Result{}, err
		}
	}

	// newLabels + systemLabels that node has represent the desired state of node labels.
	newLabels := labels(nodemgrMetadata, nodeMetadata)

	// Check if node should be updated.
	if !reflect.DeepEqual(newLabels, node.GetLabels()) {
		reqLogger.Info("Reconciling Node", "Node.Name", node.Name)
		node.SetLabels(newLabels)
		err = r.client.Update(context.TODO(), node)
		if err != nil {
			reqLogger.Error(err, "Error on update labels")
			return reconcile.Result{}, err
		}
		reqLogger.Info("Node labels updated!")
	}

	//prepare the taints that should be applied to node
	newTaints := taints(nodeMetadata, nodemgrMetadata)

	if len(newTaints) > 0 || len(node.Spec.Taints) > 0 {
		// Check if node taints are updated.
		if !reflect.DeepEqual(newTaints, node.Spec.Taints) {
			reqLogger.Info("Reconciling Node", "Node.Name", node.Name)
			node.Spec.Taints = newTaints
			err = r.client.Update(context.TODO(), node)
			if err != nil {
				reqLogger.Error(err, "Error on update Taints")
				return reconcile.Result{}, err
			}
			reqLogger.Info("Node Taints updated!")
		}
	}

	newAnnotations := annotations(nodemgrMetadata, nodeMetadata)

	if !reflect.DeepEqual(newAnnotations, node.GetAnnotations()) {
		node.SetAnnotations(newAnnotations)
		err = r.client.Update(context.TODO(), node)
		if err != nil {
			reqLogger.Error(err, "Error on update Annotations")
			return reconcile.Result{}, err
		}
		reqLogger.Info("Node Annotations updated!")
	}

	// Node already Ok - don't requeue
	// reqLogger.Info("Skip reconcile: node already ok")

	// update node status
	if !reflect.DeepEqual("synced", nodemgr.Status.Phase) {
		nodemgr.Status.Phase = "synced"
		err := r.client.Status().Update(context.TODO(), nodemgr)
		if err != nil {
			return reconcile.Result{}, err
		}
	}

	return reconcile.Result{}, nil
}

// labels merge valid system labels from nodes with node labels.
func labels(n rafabsbv1alpha1.Nodes, m rafabsbv1alpha1.Nodes) map[string]string {
	newMeta := n.GetValidLabels()
	for k, v := range m.GetValidLabels() {
		newMeta[k] = v
	}
	return newMeta
}

// annotations return the labels that must exists in node.
func annotations(n rafabsbv1alpha1.Nodes, m rafabsbv1alpha1.Nodes) map[string]string {
	newMeta := n.GetValidAnnotations()
	for k, v := range m.GetValidAnnotations() {
		newMeta[k] = v
	}
	return newMeta
}

// Taints return the labels that must exists in node.
func taints(n rafabsbv1alpha1.Nodes, m rafabsbv1alpha1.Nodes) []corev1.Taint {
	newMeta := n.GetValidTaints()
	for _, v := range m.GetValidTaints() {
		newMeta = append(newMeta, v)
	}
	return newMeta
}

// getEnv returns the value of the specific ENV var
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func remove(list []string, s string) []string {
	for i, v := range list {
		if v == s {
			list = append(list[:i], list[i+1:]...)
		}
	}
	return list
}

func contains(list []string, s string) bool {
	for _, v := range list {
		if v == s {
			return true
		}
	}
	return false
}

func (r *ReconcileNodemgr) finalizeNodemgr(reqLogger logr.Logger, node rafabsbv1alpha1.NodeMetadata) error {
	// TODO(user): Add the cleanup steps that the operator
	// needs to do before the CR can be deleted. Examples
	// of finalizers include performing backups and deleting
	// resources that are not owned by this CR, like a PVC.

	node.Node.SetLabels(node.GetValidLabels())
	node.Node.SetAnnotations(node.GetValidAnnotations())
	node.Node.Spec.Taints = node.GetValidTaints()

	err := r.client.Update(context.TODO(), node.Node)
	if err != nil {
		reqLogger.Info("Failed to clean Node metadata before deletion of Nodemgr.")
		return err
	}
	reqLogger.Info("Successfully cleaned Node metadata and finalized Nodemgr")
	return nil
}

func (r *ReconcileNodemgr) addFinalizer(reqLogger logr.Logger, m *rafabsbv1alpha1.Nodemgr) error {
	reqLogger.Info("Adding Finalizer for the Nodemgr")
	m.SetFinalizers(append(m.GetFinalizers(), nodemgrFinalizer))

	// Update CR
	err := r.client.Update(context.TODO(), m)
	if err != nil {
		reqLogger.Error(err, "Failed to update Nodemgr with finalizer")
		return err
	}
	return nil
}
