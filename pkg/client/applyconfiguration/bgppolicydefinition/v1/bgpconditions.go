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

// BGPConditionsApplyConfiguration represents an declarative configuration of the BGPConditions type for use
// with apply.
type BGPConditionsApplyConfiguration struct {
	AfiSafiIn         []string                           `json:"afiSafiIn,omitempty"`
	AsPathSet         *BGPAsPathSetApplyConfiguration    `json:"matchAsPathSet,omitempty"`
	AsPathLength      *BGPAsPathLengthApplyConfiguration `json:"asPathLength,omitempty"`
	CommunitySet      *BGPCommunitySetApplyConfiguration `json:"matchCommunitySet,omitempty"`
	ExtCommunitySet   *BGPCommunitySetApplyConfiguration `json:"matchExtCommunitySet,omitempty"`
	LargeCommunitySet *BGPCommunitySetApplyConfiguration `json:"largeCommunitySet,omitempty"`
	RouteType         *string                            `json:"routeType,omitempty"`
	NextHopInList     []string                           `json:"nextHopInList,omitempty"`
	Rpki              *string                            `json:"rpki,omitempty"`
}

// BGPConditionsApplyConfiguration constructs an declarative configuration of the BGPConditions type for use with
// apply.
func BGPConditions() *BGPConditionsApplyConfiguration {
	return &BGPConditionsApplyConfiguration{}
}

// WithAfiSafiIn adds the given value to the AfiSafiIn field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the AfiSafiIn field.
func (b *BGPConditionsApplyConfiguration) WithAfiSafiIn(values ...string) *BGPConditionsApplyConfiguration {
	for i := range values {
		b.AfiSafiIn = append(b.AfiSafiIn, values[i])
	}
	return b
}

// WithAsPathSet sets the AsPathSet field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AsPathSet field is set to the value of the last call.
func (b *BGPConditionsApplyConfiguration) WithAsPathSet(value *BGPAsPathSetApplyConfiguration) *BGPConditionsApplyConfiguration {
	b.AsPathSet = value
	return b
}

// WithAsPathLength sets the AsPathLength field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AsPathLength field is set to the value of the last call.
func (b *BGPConditionsApplyConfiguration) WithAsPathLength(value *BGPAsPathLengthApplyConfiguration) *BGPConditionsApplyConfiguration {
	b.AsPathLength = value
	return b
}

// WithCommunitySet sets the CommunitySet field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CommunitySet field is set to the value of the last call.
func (b *BGPConditionsApplyConfiguration) WithCommunitySet(value *BGPCommunitySetApplyConfiguration) *BGPConditionsApplyConfiguration {
	b.CommunitySet = value
	return b
}

// WithExtCommunitySet sets the ExtCommunitySet field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ExtCommunitySet field is set to the value of the last call.
func (b *BGPConditionsApplyConfiguration) WithExtCommunitySet(value *BGPCommunitySetApplyConfiguration) *BGPConditionsApplyConfiguration {
	b.ExtCommunitySet = value
	return b
}

// WithLargeCommunitySet sets the LargeCommunitySet field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LargeCommunitySet field is set to the value of the last call.
func (b *BGPConditionsApplyConfiguration) WithLargeCommunitySet(value *BGPCommunitySetApplyConfiguration) *BGPConditionsApplyConfiguration {
	b.LargeCommunitySet = value
	return b
}

// WithRouteType sets the RouteType field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RouteType field is set to the value of the last call.
func (b *BGPConditionsApplyConfiguration) WithRouteType(value string) *BGPConditionsApplyConfiguration {
	b.RouteType = &value
	return b
}

// WithNextHopInList adds the given value to the NextHopInList field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the NextHopInList field.
func (b *BGPConditionsApplyConfiguration) WithNextHopInList(values ...string) *BGPConditionsApplyConfiguration {
	for i := range values {
		b.NextHopInList = append(b.NextHopInList, values[i])
	}
	return b
}

// WithRpki sets the Rpki field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Rpki field is set to the value of the last call.
func (b *BGPConditionsApplyConfiguration) WithRpki(value string) *BGPConditionsApplyConfiguration {
	b.Rpki = &value
	return b
}
