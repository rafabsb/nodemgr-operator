package v1alpha1

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	corev1 "k8s.io/api/core/v1"
)

// Nodes interface to simply calculations on controller.
type Nodes interface {
	GetValidLabels() map[string]string
	GetValidAnnotations() map[string]string
	GetValidTaints() []corev1.Taint
}

// NodemgrMetadata its the struct that resume all Nodemgr metadata.
type NodemgrMetadata struct {
	Labels      map[string]string
	Annotations NodemgrAnnotations
	Taints      []NodemgrTaint
	KeyPrefix   string
	Node        *Nodemgr
}

// NodeMetadata its the struct that resume all NodeMetadata that can be update in a Node.
type NodeMetadata struct {
	Labels      map[string]string
	Annotations map[string]string
	Taints      []corev1.Taint
	KeyPrefix   string
	Node        *corev1.Node
}

// GetValidLabels return only system NodeMetadata that shoud be preserverd.
func (n NodemgrMetadata) GetValidLabels() map[string]string {
	var nodeMetadata map[string]string
	nodeMetadata = make(map[string]string)

	for k, v := range n.Labels {
		if contains(n.KeyPrefix, k) {
			nodeMetadata[k] = v
		}
	}
	return nodeMetadata
}

// GetValidLabels return only system NodeMetadata that shoud be preserverd.
func (n NodeMetadata) GetValidLabels() map[string]string {
	var nodeMetadata map[string]string
	nodeMetadata = make(map[string]string)

	for k, v := range n.Labels {
		if !contains(n.KeyPrefix, k) {
			nodeMetadata[k] = v
		}
	}
	return nodeMetadata
}

// GetValidAnnotations return only system NodeMetadata that shoud be preserverd.
func (n NodemgrMetadata) GetValidAnnotations() map[string]string {
	var newAnnotations map[string]string
	newAnnotations = make(map[string]string)
	v := reflect.ValueOf(n.Annotations)
	typeOfV := v.Type()

	for i := 0; i < v.NumField(); i++ {
		structFieldName := strings.ToLower(string(typeOfV.Field(i).Name))
		structFieldValue := v.Field(i).Interface()

		annotationJSON, err := json.Marshal(structFieldValue)
		if err != nil {
			fmt.Println("ERROR CONVERTING ANNOTATION TO JSON")
			return nil
		}
		newAnnotations[n.KeyPrefix+"/"+structFieldName] = string(annotationJSON)
	}
	return newAnnotations
}

// GetValidAnnotations return only system NodeMetadata that shoud be preserverd.
func (n NodeMetadata) GetValidAnnotations() map[string]string {
	var nodeMetadata map[string]string
	nodeMetadata = make(map[string]string)

	for k, v := range n.Annotations {
		if !contains(n.KeyPrefix, k) {
			nodeMetadata[k] = v
		}
	}
	return nodeMetadata
}

// GetValidTaints return only system NodeMetadata that shoud be preserverd.
func (n NodemgrMetadata) GetValidTaints() []corev1.Taint {
	var returnTaints []corev1.Taint
	var taintEffect corev1.TaintEffect
	for _, v := range n.Taints {
		switch effect := v.Effect; effect {
		case "NoExecute":
			taintEffect = corev1.TaintEffectNoExecute
		case "NoSchedule":
			taintEffect = corev1.TaintEffectNoSchedule
		case "PreferNoSchedule":
			taintEffect = corev1.TaintEffectPreferNoSchedule
		default:
			continue
		}
		if contains(n.KeyPrefix, v.Key) {
			returnTaints = append(returnTaints, corev1.Taint{
				Key:    v.Key,
				Value:  v.Value,
				Effect: taintEffect,
			})
		}
	}
	return returnTaints
}

// GetValidTaints return only system NodeMetadata that shoud be preserverd.
func (n NodeMetadata) GetValidTaints() []corev1.Taint {
	var nodeMetadata = []corev1.Taint{}

	for _, v := range n.Taints {
		if !contains(n.KeyPrefix, v.Key) {
			nodeMetadata = append(nodeMetadata, v)
		}
	}
	return nodeMetadata
}

func contains(KeyPrefix string, s string) bool {

	if strings.Contains(s, KeyPrefix) {
		return true
	}

	return false
}
