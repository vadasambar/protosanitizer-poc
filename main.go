package main

import (
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-csi/csi-lib-utils/protosanitizer"
	"k8s.io/klog/v2"
)

func main() {
	req := &csi.NodePublishVolumeRequest{
		Secrets: map[string]string{
			"password": "foobar",
		},
		VolumeContext: map[string]string{
			"mount": "true",
		},
	}

	valuesLogger := klog.LoggerWithValues(klog.NewKlogr(), "pod-name", req.VolumeContext["pod-name"], "namespace", req.VolumeContext["namespace"], "uid", req.VolumeContext["uid"])

	valuesLogger.Info("Incoming NodePublishVolume request", "unstripped", req)
	valuesLogger.Info("Incoming NodePublishVolume request", "stripped", protosanitizer.StripSecrets(req))
}
