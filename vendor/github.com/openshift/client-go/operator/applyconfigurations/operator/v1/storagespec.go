// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	operatorv1 "github.com/openshift/api/operator/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// StorageSpecApplyConfiguration represents an declarative configuration of the StorageSpec type for use
// with apply.
type StorageSpecApplyConfiguration struct {
	OperatorSpecApplyConfiguration `json:",inline"`
	VSphereStorageDriver           *operatorv1.StorageDriverType `json:"vsphereStorageDriver,omitempty"`
}

// StorageSpecApplyConfiguration constructs an declarative configuration of the StorageSpec type for use with
// apply.
func StorageSpec() *StorageSpecApplyConfiguration {
	return &StorageSpecApplyConfiguration{}
}

// WithManagementState sets the ManagementState field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ManagementState field is set to the value of the last call.
func (b *StorageSpecApplyConfiguration) WithManagementState(value operatorv1.ManagementState) *StorageSpecApplyConfiguration {
	b.ManagementState = &value
	return b
}

// WithLogLevel sets the LogLevel field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LogLevel field is set to the value of the last call.
func (b *StorageSpecApplyConfiguration) WithLogLevel(value operatorv1.LogLevel) *StorageSpecApplyConfiguration {
	b.LogLevel = &value
	return b
}

// WithOperatorLogLevel sets the OperatorLogLevel field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the OperatorLogLevel field is set to the value of the last call.
func (b *StorageSpecApplyConfiguration) WithOperatorLogLevel(value operatorv1.LogLevel) *StorageSpecApplyConfiguration {
	b.OperatorLogLevel = &value
	return b
}

// WithUnsupportedConfigOverrides sets the UnsupportedConfigOverrides field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UnsupportedConfigOverrides field is set to the value of the last call.
func (b *StorageSpecApplyConfiguration) WithUnsupportedConfigOverrides(value runtime.RawExtension) *StorageSpecApplyConfiguration {
	b.UnsupportedConfigOverrides = &value
	return b
}

// WithObservedConfig sets the ObservedConfig field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ObservedConfig field is set to the value of the last call.
func (b *StorageSpecApplyConfiguration) WithObservedConfig(value runtime.RawExtension) *StorageSpecApplyConfiguration {
	b.ObservedConfig = &value
	return b
}

// WithVSphereStorageDriver sets the VSphereStorageDriver field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the VSphereStorageDriver field is set to the value of the last call.
func (b *StorageSpecApplyConfiguration) WithVSphereStorageDriver(value operatorv1.StorageDriverType) *StorageSpecApplyConfiguration {
	b.VSphereStorageDriver = &value
	return b
}
