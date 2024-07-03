/*
Copyright The Kubernetes Authors.

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

package v1

// BGPPolicyApplyModelApplyConfiguration represents an declarative configuration of the BGPPolicyApplyModel type for use
// with apply.
type BGPPolicyApplyModelApplyConfiguration struct {
	NeighIPAddress *string  `json:"ipAddress,omitempty"`
	PolicyType     *string  `json:"policyType,omitempty"`
	Policies       []string `json:"policies,omitempty"`
	RouteAction    *string  `json:"routeAction,omitempty"`
}

// BGPPolicyApplyModelApplyConfiguration constructs an declarative configuration of the BGPPolicyApplyModel type for use with
// apply.
func BGPPolicyApplyModel() *BGPPolicyApplyModelApplyConfiguration {
	return &BGPPolicyApplyModelApplyConfiguration{}
}

// WithNeighIPAddress sets the NeighIPAddress field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NeighIPAddress field is set to the value of the last call.
func (b *BGPPolicyApplyModelApplyConfiguration) WithNeighIPAddress(value string) *BGPPolicyApplyModelApplyConfiguration {
	b.NeighIPAddress = &value
	return b
}

// WithPolicyType sets the PolicyType field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PolicyType field is set to the value of the last call.
func (b *BGPPolicyApplyModelApplyConfiguration) WithPolicyType(value string) *BGPPolicyApplyModelApplyConfiguration {
	b.PolicyType = &value
	return b
}

// WithPolicies adds the given value to the Policies field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Policies field.
func (b *BGPPolicyApplyModelApplyConfiguration) WithPolicies(values ...string) *BGPPolicyApplyModelApplyConfiguration {
	for i := range values {
		b.Policies = append(b.Policies, values[i])
	}
	return b
}

// WithRouteAction sets the RouteAction field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RouteAction field is set to the value of the last call.
func (b *BGPPolicyApplyModelApplyConfiguration) WithRouteAction(value string) *BGPPolicyApplyModelApplyConfiguration {
	b.RouteAction = &value
	return b
}
