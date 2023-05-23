package utils

import (
	"context"
	"fmt"
	"net/http"

	corev1 "k8s.io/api/core/v1"
	apiErrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// GetFenceAgentsRemediationPod fetches the FAR pod based on FAR's label and namespace
func GetFenceAgentsRemediationPod(r client.Reader) (*corev1.Pod, error) {
	var podNamespace string
	pods := &corev1.PodList{}
	selector := labels.NewSelector()
	requirement, _ := labels.NewRequirement("app", selection.Equals, []string{"fence-agents-remediation-operator"})
	selector = selector.Add(*requirement)
	podNamespace, err := GetDeploymentNamespace()
	if err != nil {
		return nil, fmt.Errorf("failed fetching FAR namespace - %w", err)
	}
	err = r.List(context.Background(), pods, &client.ListOptions{LabelSelector: selector, Namespace: podNamespace})
	if err != nil {
		return nil, fmt.Errorf("failed fetching FAR pod - %w", err)
	}
	if len(pods.Items) == 0 {
		podNotFoundErr := &apiErrors.StatusError{ErrStatus: metav1.Status{
			Status: metav1.StatusFailure,
			Code:   http.StatusNotFound,
			Reason: metav1.StatusReasonNotFound,
		}}
		return nil, fmt.Errorf("no Fence Agent pods were found - %w", podNotFoundErr)
	}
	return &pods.Items[0], nil
}
