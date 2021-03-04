package services

import (
	"bytes"
	"context"
	"fmt"

	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/remotecommand"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/apiutil"
)

type Reconciler struct {
	client.Client
	Log    logr.Logger
	Config *rest.Config
	Scheme *runtime.Scheme
}

func (r *Reconciler) NewGroup(does ...func() (ctrl.Result, error)) Group {
	g := Group{
		Ch: make(chan Outcome, 10),
	}
	return *g.Do(does...)
}

func (r *Reconciler) NewGroupOf(does ...[]func() (ctrl.Result, error)) Group {
	funcs := make([]func() (ctrl.Result, error), 0)
	for _, slice := range does {
		funcs = append(funcs, slice...)
	}
	return r.NewGroup(funcs...)
}

func (r *Reconciler) ApplyUnowned(ctx context.Context, manager string, into client.Object) (ctrl.Result, error) {
	intoNamed := client.ObjectKeyFromObject(into)
	intoGvk := into.GetObjectKind().GroupVersionKind()
	log := r.Log.WithValues(
		"apiVersion", intoGvk.GroupVersion(),
		"kind", intoGvk.Kind,
		"resource", intoNamed,
	)

	err := r.Patch(
		ctx, into, client.Apply,
		client.FieldOwner(manager),
		client.ForceOwnership,
	)
	if err != nil {
		return ctrl.Result{}, fmt.Errorf("failed to patch: %v", err)
	}
	log.V(1).Info("Applied")

	return ctrl.Result{}, nil
}

func (r *Reconciler) Apply(ctx context.Context, owner client.Object, into client.Object) (ctrl.Result, error) {
	ownerNamed := client.ObjectKeyFromObject(owner)
	intoNamed := client.ObjectKeyFromObject(into)
	intoGvk := into.GetObjectKind().GroupVersionKind()
	log := r.Log.WithValues(
		"apiVersion", intoGvk.GroupVersion(),
		"kind", intoGvk.Kind,
		"resource", intoNamed,
	)

	err := ctrl.SetControllerReference(owner.(metav1.Object), into.(metav1.Object), r.Scheme)
	if err != nil {
		return ctrl.Result{}, fmt.Errorf("cannot set owner: %v", err)
	}

	err = r.Patch(
		ctx, into, client.Apply,
		client.FieldOwner(ownerNamed.String()),
		client.ForceOwnership,
	)
	if err != nil {
		return ctrl.Result{}, fmt.Errorf("failed to patch: %v", err)
	}
	log.V(1).Info("Applied")

	return ctrl.Result{}, nil
}

func (r *Reconciler) Exec(container string, pod *corev1.Pod, commands ...string) (stdout string, stderr string, err error) {
	gvk := pod.GroupVersionKind()
	log := r.Log.WithValues(
		"apiVersion", gvk.GroupVersion(),
		"kind", gvk.Kind,
		"resource", types.NamespacedName{Name: pod.Name, Namespace: pod.Namespace},
		"container", container,
	)
	rest, err := apiutil.RESTClientForGVK(
		gvk, false,
		r.Config,
		serializer.NewCodecFactory(r.Scheme),
	)
	if err != nil {
		return stdout, stderr, fmt.Errorf("getting RESTClient: %v", err)
	}

	execReq := rest.Post().
		Resource("pods").
		Name(pod.Name).
		Namespace(pod.Namespace).
		SubResource("exec")
	execReq.VersionedParams(&corev1.PodExecOptions{
		Container: container,
		Command:   commands,
		Stdin:     false,
		Stdout:    true,
		Stderr:    true,
	}, scheme.ParameterCodec)

	exec, err := remotecommand.NewSPDYExecutor(r.Config, "POST", execReq.URL())
	if err != nil {
		return stdout, stderr, fmt.Errorf("while creating Executor: %v", err)
	}
	//TODO don't log passwords
	log.V(2).Info("Executing", "command", commands)
	stdoutBuf := &bytes.Buffer{}
	stderrBuf := &bytes.Buffer{}
	err = exec.Stream(remotecommand.StreamOptions{
		Stdout: stdoutBuf,
		Stderr: stderrBuf,
		Tty:    false,
	})
	stdout = stdoutBuf.String()
	stderr = stderrBuf.String()
	if err != nil {
		return stdout, stderr, fmt.Errorf("during exec.Stream: %w", err)
	}
	return stdout, stderr, nil
}
