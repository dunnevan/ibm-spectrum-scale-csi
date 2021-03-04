package v1

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"
)

var (
	nodeAffinityIsOsLinux = corev1.NodeSelectorRequirement{
		Key:      corev1.LabelOSStable,
		Operator: corev1.NodeSelectorOpIn,
		Values:   []string{"linux"},
	}
	nodeAffinityInArch = corev1.NodeSelectorRequirement{
		Key:      corev1.LabelArchStable,
		Operator: corev1.NodeSelectorOpIn,
		Values:   []string{"amd64", "ppc64le", "s390x"},
	}
	NodeAffinityIsSupportedPlatform = []corev1.NodeSelectorRequirement{
		nodeAffinityIsOsLinux,
		nodeAffinityInArch,
	}
)

var (
	ListRequirementIsOsLinux        labels.Requirement
	ListRequirementInArch           labels.Requirement
	ListSelectorIsSupportedPlatform labels.Selector
)

func init() {
	isOsLinux, err := labels.NewRequirement(
		corev1.LabelOSStable, selection.Equals, []string{"linux"},
	)
	if err != nil {
		panic(fmt.Errorf("supported os label malformed: %v", err))
	}
	ListRequirementIsOsLinux = *isOsLinux

	inArch, err := labels.NewRequirement(
		corev1.LabelArchStable, selection.In, []string{"amd64", "ppc64le", "s390x"},
	)
	if err != nil {
		panic(fmt.Errorf("supported arch label malformed: %v", err))
	}
	ListRequirementInArch = *inArch

	ListSelectorIsSupportedPlatform = labels.NewSelector().Add(
		ListRequirementInArch,
		ListRequirementIsOsLinux,
	)
}
