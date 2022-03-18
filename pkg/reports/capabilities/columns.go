package capabilities

import (
	"awesomeProject/pkg/models"
)

type Bundle struct {
	BundleName      string   `json:"bundleName"`
	PackageName     string   `json:"packageName"`
	BundleImagePath string   `json:"bundleImagePath,omitempty"`
	DefaultChannel  string   `json:"defaultChannel,omitempty"`
	Channels        []string `json:"bundleChannel,omitempty"`
	AuditErrors     []string `json:"errors,omitempty"`
	Capabilities    bool     `json:"OperatorInstalled"`
	InstallLogs     []string `json:"InstallLogs"`
	CleanUpLogs     []string `json:"CleanUpLogs"`
}

func NewBundle(v models.AuditCapabilities) *Bundle {
	col := Bundle{}
	col.BundleName = v.OperatorBundleName
	col.PackageName = v.PackageName
	col.BundleImagePath = v.OperatorBundleImagePath
	col.DefaultChannel = v.DefaultChannel
	col.Channels = v.Channels
	col.AuditErrors = v.Errors
	col.Capabilities = v.Capabilities
	col.InstallLogs = v.InstallLogs
	col.CleanUpLogs = v.CleanUpLogs

	return &col
}
