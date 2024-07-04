package controller

import (
	"context"

	"github.com/ayushsatyam146/shipton/api"
	"github.com/go-logr/logr"
	tektonv1beta1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	tektonclient "github.com/tektoncd/pipeline/pkg/client/clientset/versioned"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

type ShiptonBuildReconciler struct {
	client.Client
	Log          logr.Logger
	Scheme       *runtime.Scheme
	TektonClient tektonclient.Interface
}

func (r *ShiptonBuildReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := r.Log.WithValues("shiptonbuild", req.NamespacedName)

	// Fetch the ShiptonBuild instance
	var shiptonBuild api.ShiptonBuild
	if err := r.Get(ctx, req.NamespacedName, &shiptonBuild); err != nil {
		if errors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		log.Error(err, "unable to fetch ShiptonBuild")
		return ctrl.Result{}, err
	}

	// Define Tekton PipelineRun
	pipelineRun := &tektonv1beta1.PipelineRun{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: "shipton-pipeline-run-",
			Namespace:    shiptonBuild.Namespace,
		},
		Spec: tektonv1beta1.PipelineRunSpec{
			PipelineRef: tektonv1beta1.PipelineRef{
				Name: "build-and-push-pipeline",
			},
			Params: []tektonv1beta1.Param{
				{
					Name:  "IMAGE",
					Value: tektonv1beta1.ArrayOrString{Type: tektonv1beta1.ParamTypeString, StringVal: shiptonBuild.Spec.Image},
				},
			},
		},
	}

	// Set ShiptonBuild instance as the owner and controller
	if err := controllerutil.SetControllerReference(&shiptonBuild, pipelineRun, r.Scheme); err != nil {
		return ctrl.Result{}, err
	}

	// Check if the PipelineRun already exists
	existingPipelineRun := &tektonv1beta1.PipelineRun{}
	err := r.TektonClient.TektonV1beta1().PipelineRuns(shiptonBuild.Namespace).Get(ctx, pipelineRun.Name, metav1.GetOptions{})
	if err != nil && errors.IsNotFound(err) {
		log.Info("Creating a new PipelineRun", "PipelineRun.Namespace", pipelineRun.Namespace, "PipelineRun.Name", pipelineRun.Name)
		err = r.TektonClient.TektonV1beta1().PipelineRuns(shiptonBuild.Namespace).Create(ctx, pipelineRun, metav1.CreateOptions{})
		if err != nil {
			return ctrl.Result{}, err
		}
	} else if err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *ShiptonBuildReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&api.ShiptonBuild{}).
		Complete(r)
}
