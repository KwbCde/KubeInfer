package controller

import (
	"context"
	"time"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"

	mlv1 "github.com/KwbCde/KubeInfer/api/v1"
)

type RealClock struct{}

func (RealClock) Now() time.Time { return time.Now() } // nolint:staticcheck

// Clock know how to get current time
// Used to fake out timing for testing
type Clock interface {
	Now() time.Time
}

// InferenceJobReconciler reconciles a InferenceJob object
type InferenceJobReconciler struct {
	client.Client
	Scheme *runtime.Scheme
	Clock
}

// +kubebuilder:rbac:groups=ml.kubeinfer.io,resources=inferencejobs,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=ml.kubeinfer.io,resources=inferencejobs/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=ml.kubeinfer.io,resources=inferencejobs/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the InferenceJob object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.22.4/pkg/reconcile
func (r *InferenceJobReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := logf.FromContext(ctx)

	job := &mlv1.InferenceJob{}

	if err := r.Get(ctx, req.NamespacedName, job); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// If this is the first time reconciling the job set phase = pending
	if job.Status.Phase == "" {
		job.Status.Phase = "Pending"
		logger.Info("Setting initial phase to Pending", "job", job.Name)

		if err := r.Status().Update(ctx, job); err != nil {
			logger.Error(err, "Failed to update initial job status")
			return ctrl.Result{}, err
		}

		return ctrl.Result{RequeueAfter: 1 * time.Second}, nil
	}
	switch job.Status.Phase {
	case "Pending":
		logger.Info("Job is pending, next step is pod creation")
		// todo: create inference pod here
	case "Running":
		logger.Info("Job is running. Check pod status")
	case "Succeeded", "Failed":
		logger.Info("Job completed. Nothing to do")
	}
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *InferenceJobReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&mlv1.InferenceJob{}).
		Named("inferencejob").
		Complete(r)
}
