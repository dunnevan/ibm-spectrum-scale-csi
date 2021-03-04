package k8s_helpers

import corev1 "k8s.io/api/core/v1"

func IsContainerReady(pod corev1.Pod, containerName string) bool {
	for _, containerStatus := range pod.Status.ContainerStatuses {
		if containerStatus.Name == containerName {
			return containerStatus.Ready
		}
	}
	return false
}
