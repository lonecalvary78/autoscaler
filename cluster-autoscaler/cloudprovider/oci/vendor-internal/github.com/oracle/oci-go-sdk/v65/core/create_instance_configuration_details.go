// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"encoding/json"
	"fmt"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/oci/vendor-internal/github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateInstanceConfigurationDetails Details for creating an instance configuration by providing a list of configuration settings.
type CreateInstanceConfigurationDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment
	// containing the instance configuration.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	InstanceDetails InstanceConfigurationInstanceDetails `mandatory:"true" json:"instanceDetails"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`
}

// GetCompartmentId returns CompartmentId
func (m CreateInstanceConfigurationDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetDefinedTags returns DefinedTags
func (m CreateInstanceConfigurationDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetDisplayName returns DisplayName
func (m CreateInstanceConfigurationDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetFreeformTags returns FreeformTags
func (m CreateInstanceConfigurationDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

func (m CreateInstanceConfigurationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateInstanceConfigurationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateInstanceConfigurationDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateInstanceConfigurationDetails CreateInstanceConfigurationDetails
	s := struct {
		DiscriminatorParam string `json:"source"`
		MarshalTypeCreateInstanceConfigurationDetails
	}{
		"NONE",
		(MarshalTypeCreateInstanceConfigurationDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *CreateInstanceConfigurationDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DefinedTags     map[string]map[string]interface{}    `json:"definedTags"`
		DisplayName     *string                              `json:"displayName"`
		FreeformTags    map[string]string                    `json:"freeformTags"`
		CompartmentId   *string                              `json:"compartmentId"`
		InstanceDetails instanceconfigurationinstancedetails `json:"instanceDetails"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DefinedTags = model.DefinedTags

	m.DisplayName = model.DisplayName

	m.FreeformTags = model.FreeformTags

	m.CompartmentId = model.CompartmentId

	nn, e = model.InstanceDetails.UnmarshalPolymorphicJSON(model.InstanceDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.InstanceDetails = nn.(InstanceConfigurationInstanceDetails)
	} else {
		m.InstanceDetails = nil
	}

	return
}
