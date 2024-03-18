package util

import (
	tunedv1 "github.com/openshift/cluster-node-tuning-operator/pkg/apis/tuned/v1"
)

func HasDeferredUpdateAnnotation(anns map[string]string) bool {
	if anns == nil {
		return false
	}
	_, ok := anns[tunedv1.TunedDeferredUpdate]
	return ok
}

func AddTunedDeferredUpdateAnnotation(tuned *tunedv1.Tuned) {
	if tuned.Annotations == nil {
		tuned.Annotations = make(map[string]string)
	}
	tuned.Annotations[tunedv1.TunedDeferredUpdate] = ""
}

func DeleteTunedDeferredUpdateAnnotation(tuned *tunedv1.Tuned) {
	if tuned.Annotations == nil {
		// nothing to do
		return
	}
	delete(tuned.Annotations, tunedv1.TunedDeferredUpdate)
}
