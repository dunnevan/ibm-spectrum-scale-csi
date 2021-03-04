package services

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/selection"
)

func SelectionOperator_ListAPI_ToCoreV1(op selection.Operator) (corev1.NodeSelectorOperator, error) {
	switch op {
	case selection.In, selection.Equals, selection.DoubleEquals:
		return corev1.NodeSelectorOpIn, nil
	case selection.NotIn, selection.NotEquals:
		return corev1.NodeSelectorOpNotIn, nil
	case selection.Exists:
		return corev1.NodeSelectorOpExists, nil
	case selection.DoesNotExist:
		return corev1.NodeSelectorOpDoesNotExist, nil
	case selection.GreaterThan:
		return corev1.NodeSelectorOpGt, nil
	case selection.LessThan:
		return corev1.NodeSelectorOpLt, nil
	default:
		return "", fmt.Errorf("unknown selection.Operator: %s", op)
	}
}

func NodeSelectionOperator_ListAPI_ToCoreV1(op corev1.NodeSelectorOperator) (selection.Operator, error) {
	switch op {
	case corev1.NodeSelectorOpIn:
		return selection.In, nil
	case corev1.NodeSelectorOpNotIn:
		return selection.NotIn, nil
	case corev1.NodeSelectorOpExists:
		return selection.Exists, nil
	case corev1.NodeSelectorOpDoesNotExist:
		return selection.DoesNotExist, nil
	case corev1.NodeSelectorOpGt:
		return selection.GreaterThan, nil
	case corev1.NodeSelectorOpLt:
		return selection.LessThan, nil
	default:
		return "", fmt.Errorf("unknown corev1.NodeSelectorOperator: %s", op)
	}
}
