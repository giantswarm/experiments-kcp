/*
Copyright The KCP Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"

	v1alpha1 "github.com/kcp-dev/kcp/contrib/mounts-vw/apis/mounts/v1alpha1"
)

// VClusterSpecApplyConfiguration represents a declarative configuration of the VClusterSpec type for use
// with apply.
type VClusterSpecApplyConfiguration struct {
	Version      *string                   `json:"version,omitempty"`
	Mode         *v1alpha1.KubeClusterMode `json:"mode,omitempty"`
	SecretString *string                   `json:"secretString,omitempty"`
	SecretRef    *v1.ObjectReference       `json:"secretRef,omitempty"`
}

// VClusterSpecApplyConfiguration constructs a declarative configuration of the VClusterSpec type for use with
// apply.
func VClusterSpec() *VClusterSpecApplyConfiguration {
	return &VClusterSpecApplyConfiguration{}
}

// WithVersion sets the Version field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Version field is set to the value of the last call.
func (b *VClusterSpecApplyConfiguration) WithVersion(value string) *VClusterSpecApplyConfiguration {
	b.Version = &value
	return b
}

// WithMode sets the Mode field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Mode field is set to the value of the last call.
func (b *VClusterSpecApplyConfiguration) WithMode(value v1alpha1.KubeClusterMode) *VClusterSpecApplyConfiguration {
	b.Mode = &value
	return b
}

// WithSecretString sets the SecretString field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SecretString field is set to the value of the last call.
func (b *VClusterSpecApplyConfiguration) WithSecretString(value string) *VClusterSpecApplyConfiguration {
	b.SecretString = &value
	return b
}

// WithSecretRef sets the SecretRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SecretRef field is set to the value of the last call.
func (b *VClusterSpecApplyConfiguration) WithSecretRef(value v1.ObjectReference) *VClusterSpecApplyConfiguration {
	b.SecretRef = &value
	return b
}