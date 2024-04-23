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

func AddTunedDeferredUpdateAnnotation(prof *tunedv1.Profile) {
	if prof.Annotations == nil {
		prof.Annotations = make(map[string]string)
	}
	prof.Annotations[tunedv1.TunedDeferredUpdate] = ""
}

func DeleteTunedDeferredUpdateAnnotation(prof *tunedv1.Profile) {
	if prof.Annotations == nil {
		// nothing to do
		return
	}
	delete(prof.Annotations, tunedv1.TunedDeferredUpdate)
}
