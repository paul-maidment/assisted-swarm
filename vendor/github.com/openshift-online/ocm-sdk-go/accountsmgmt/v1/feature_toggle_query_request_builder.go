/*
Copyright (c) 2020 Red Hat, Inc.

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

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1

// FeatureToggleQueryRequestBuilder contains the data and logic needed to build 'feature_toggle_query_request' objects.
//
//
type FeatureToggleQueryRequestBuilder struct {
	bitmap_        uint32
	organizationID string
}

// NewFeatureToggleQueryRequest creates a new builder of 'feature_toggle_query_request' objects.
func NewFeatureToggleQueryRequest() *FeatureToggleQueryRequestBuilder {
	return &FeatureToggleQueryRequestBuilder{}
}

// OrganizationID sets the value of the 'organization_ID' attribute to the given value.
//
//
func (b *FeatureToggleQueryRequestBuilder) OrganizationID(value string) *FeatureToggleQueryRequestBuilder {
	b.organizationID = value
	b.bitmap_ |= 1
	return b
}

// Copy copies the attributes of the given object into this builder, discarding any previous values.
func (b *FeatureToggleQueryRequestBuilder) Copy(object *FeatureToggleQueryRequest) *FeatureToggleQueryRequestBuilder {
	if object == nil {
		return b
	}
	b.bitmap_ = object.bitmap_
	b.organizationID = object.organizationID
	return b
}

// Build creates a 'feature_toggle_query_request' object using the configuration stored in the builder.
func (b *FeatureToggleQueryRequestBuilder) Build() (object *FeatureToggleQueryRequest, err error) {
	object = new(FeatureToggleQueryRequest)
	object.bitmap_ = b.bitmap_
	object.organizationID = b.organizationID
	return
}
