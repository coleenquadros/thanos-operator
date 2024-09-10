package manifests

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/utils/ptr"
)

const (
	DefaultThanosImage   = "quay.io/thanos/thanos"
	DefaultThanosVersion = "v0.35.1"

	defaultLogLevel  = "info"
	defaultLogFormat = "logfmt"
)

// Options is a struct that holds the options for the common manifests
type Options struct {
	// Name is the name of the object
	Name string
	// Namespace is the namespace of the object
	Namespace string
	// Replicas is the number of replicas for the object
	Replicas int32
	// Labels is the labels for the object
	Labels map[string]string
	// Image is the image to use for the component
	Image *string
	// Version is the version of Thanos
	Version *string
	// ResourceRequirements for the component
	ResourceRequirements *corev1.ResourceRequirements
	// LogLevel is the log level for the component
	LogLevel *string
	// LogFormat is the log format for the component
	LogFormat      *string
	containerImage string
	// EnableServiceMonitor is a flag to enable ServiceMonitor
	EnableServiceMonitor bool
}

// ApplyDefaults applies the default values to the options
func (o Options) ApplyDefaults() Options {
	if o.Image == nil || *o.Image == "" {
		o.Image = ptr.To(DefaultThanosImage)
	}

	if o.Version == nil || *o.Version == "" {
		o.Version = ptr.To(DefaultThanosVersion)
	}
	o.containerImage = fmt.Sprintf("%s:%s", *o.Image, *o.Version)

	if o.LogLevel == nil || *o.LogLevel == "" {
		o.LogLevel = ptr.To(defaultLogLevel)
	}

	if o.LogFormat == nil || *o.LogFormat == "" {
		o.LogFormat = ptr.To(defaultLogFormat)
	}

	return o
}

// GetContainerImage for the Options
func (o Options) GetContainerImage() string {
	return o.containerImage
}
