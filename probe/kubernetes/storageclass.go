package kubernetes

import (
	"github.com/weaveworks/scope/report"
	storagev1 "k8s.io/api/storage/v1"
)

// StorageClass represent kubernetes StorageClass interface
type StorageClass interface {
	Meta
	GetNode(probeID string) report.Node
	GetProvisioner() string
}

// storageClass represents kubernetes storage classes
type storageClass struct {
	*storagev1.StorageClass
	Meta
}

// NewStorageClass returns new Storage Class type
func NewStorageClass(p *storagev1.StorageClass) StorageClass {
	return &storageClass{StorageClass: p, Meta: meta{p.ObjectMeta}}
}

// GetProvisioner returns the provisioner name
func (p *storageClass) GetProvisioner() string {
	return p.Provisioner
}

// GetNode returns StorageClass as Node
func (p *storageClass) GetNode(probeID string) report.Node {
	return p.MetaNode(report.MakeStorageClassNodeID(p.UID())).WithLatests(map[string]string{
		NodeType:              "Storage Class",
		Name:                  p.GetName(),
		Provisioner:           p.GetProvisioner(),
		report.ControlProbeID: probeID,
	}).WithLatestActiveControls(Describe)
}
