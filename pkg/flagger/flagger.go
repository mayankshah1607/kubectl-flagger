package flagger

import (
	"context"
	"fmt"
	"os"

	"github.com/mayankshah1607/kubectl-flagger/pkg/k8s"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/remotecommand"
)

func getCurlCmdWithArgs(name, namespace, endpoint string) []string {
	return []string{
		"curl",
		"-d",
		fmt.Sprintf("'{\"name\": \"%s\", \"namespace\", \"%s\"}", name, namespace),
		fmt.Sprintf("http://localhost:8080%s", endpoint),
	}
}

func getLoadtesterPodName(loadtesterNs string) (string, error) {
	// select loadtester pod
	options := metav1.ListOptions{
		LabelSelector: "app.kubernetes.io/name=loadtester",
	}

	podList, err := k8s.Client.CoreV1().Pods(loadtesterNs).List(context.Background(), options)
	if err != nil {
		return "", err
	}

	return (*podList).Items[0].Name, nil
}

func execAndRunCurl(name, namespace, loadtesterNs, endpoint string) error {
	curlCmd := getCurlCmdWithArgs(name, namespace, endpoint)
	podName, err := getLoadtesterPodName(loadtesterNs)
	if err != nil {
		return fmt.Errorf("could not get flagger-loadtester pod name:%s", err)
	}
	req := k8s.Client.CoreV1().RESTClient().Post().Resource("pods").Name(podName).
		Namespace(loadtesterNs).SubResource("exec")

	option := &v1.PodExecOptions{
		Command: curlCmd,
		Stdin:   true,
		Stdout:  true,
		Stderr:  true,
		TTY:     true,
	}

	req.VersionedParams(option, scheme.ParameterCodec)
	exec, err := remotecommand.NewSPDYExecutor(k8s.RestConfig, "POST", req.URL())
	if err != nil {
		return fmt.Errorf("failed to create command executor: %s", err)
	}

	err = exec.Stream(remotecommand.StreamOptions{
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	})

	if err != nil {
		return fmt.Errorf("failed to run command: %s", err)
	}
	return nil
}

// Promote promotes a canary from flagger loadtester pod
func Promote(name, namespace, loadtesterNs string) error {
	err := execAndRunCurl(name, namespace, loadtesterNs, "/gate/open")
	if err != nil {
		return err
	}
	return nil
}

// Rollback aborts a canary
func Rollback(name, namespace, loadtesterNs string) error {
	err := execAndRunCurl(name, namespace, loadtesterNs, "/rollback/open")
	if err != nil {
		return err
	}
	return nil
}
