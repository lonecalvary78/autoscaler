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
	"fmt"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/oci/vendor-internal/github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Vnic A virtual network interface card. Each VNIC resides in a subnet in a VCN.
// An instance attaches to a VNIC to obtain a network connection into the VCN
// through that subnet. Each instance has a *primary VNIC* that is automatically
// created and attached during launch. You can add *secondary VNICs* to an
// instance after it's launched. For more information, see
// Virtual Network Interface Cards (VNICs) (https://docs.oracle.com/iaas/Content/Network/Tasks/managingVNICs.htm).
// Each VNIC has a *primary private IP* that is automatically assigned during launch.
// You can add *secondary private IPs* to a VNIC after it's created. For more
// information, see CreatePrivateIp and
// IP Addresses (https://docs.oracle.com/iaas/Content/Network/Tasks/managingIPaddresses.htm).
//
// If you are an Oracle Cloud VMware Solution customer, you will have secondary VNICs
// that reside in a VLAN instead of a subnet. These VNICs have other differences, which
// are called out in the descriptions of the relevant attributes in the `Vnic` object.
// Also see Vlan.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// Getting Started with Policies (https://docs.oracle.com/iaas/Content/Identity/Concepts/policygetstarted.htm).
type Vnic struct {

	// The VNIC's availability domain.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the VNIC.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VNIC.
	Id *string `mandatory:"true" json:"id"`

	// The current state of the VNIC.
	LifecycleState VnicLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the VNIC was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Security attributes (https://docs.oracle.com/iaas/Content/zero-trust-packet-routing/zpr-artifacts.htm#security-attributes) are labels
	// for a resource that can be referenced in a Zero Trust Packet Routing (https://docs.oracle.com/iaas/Content/zero-trust-packet-routing/overview.htm)
	// (ZPR) policy to control access to ZPR-supported resources.
	// Example: `{"Oracle-DataSecurity-ZPR": {"MaxEgressCount": {"value":"42","mode":"audit"}}}`
	SecurityAttributes map[string]map[string]interface{} `mandatory:"false" json:"securityAttributes"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The hostname for the VNIC's primary private IP. Used for DNS. The value is the hostname
	// portion of the primary private IP's fully qualified domain name (FQDN)
	// (for example, `bminstance1` in FQDN `bminstance1.subnet123.vcn1.oraclevcn.com`).
	// Must be unique across all VNICs in the subnet and comply with
	// RFC 952 (https://tools.ietf.org/html/rfc952) and
	// RFC 1123 (https://tools.ietf.org/html/rfc1123).
	// For more information, see
	// DNS in Your Virtual Cloud Network (https://docs.oracle.com/iaas/Content/Network/Concepts/dns.htm).
	// Example: `bminstance1`
	HostnameLabel *string `mandatory:"false" json:"hostnameLabel"`

	// Whether the VNIC is the primary VNIC (the VNIC that is automatically created
	// and attached during instance launch).
	IsPrimary *bool `mandatory:"false" json:"isPrimary"`

	// The MAC address of the VNIC.
	// If the VNIC belongs to a VLAN as part of the Oracle Cloud VMware Solution,
	// the MAC address is learned. If the VNIC belongs to a subnet, the
	// MAC address is a static, Oracle-provided value.
	// Example: `00:00:00:00:00:01`
	MacAddress *string `mandatory:"false" json:"macAddress"`

	// A list of the OCIDs of the network security groups that the VNIC belongs to.
	// If the VNIC belongs to a VLAN as part of the Oracle Cloud VMware Solution (instead of
	// belonging to a subnet), the value of the `nsgIds` attribute is ignored. Instead, the
	// VNIC belongs to the NSGs that are associated with the VLAN itself. See Vlan.
	// For more information about NSGs, see
	// NetworkSecurityGroup.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// If the VNIC belongs to a VLAN as part of the Oracle Cloud VMware Solution (instead of
	// belonging to a subnet), the `vlanId` is the OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VLAN the VNIC is in. See
	// Vlan. If the VNIC is instead in a subnet, `subnetId` has a value.
	VlanId *string `mandatory:"false" json:"vlanId"`

	// The private IP address of the primary `privateIp` object on the VNIC.
	// The address is within the CIDR of the VNIC's subnet.
	// Example: `10.0.3.3`
	PrivateIp *string `mandatory:"false" json:"privateIp"`

	// The public IP address of the VNIC, if one is assigned.
	PublicIp *string `mandatory:"false" json:"publicIp"`

	// Whether the source/destination check is disabled on the VNIC.
	// Defaults to `false`, which means the check is performed. For information
	// about why you would skip the source/destination check, see
	// Using a Private IP as a Route Target (https://docs.oracle.com/iaas/Content/Network/Tasks/managingroutetables.htm#privateip).
	//
	// If the VNIC belongs to a VLAN as part of the Oracle Cloud VMware Solution (instead of
	// belonging to a subnet), the `skipSourceDestCheck` attribute is `true`.
	// This is because the source/destination check is always disabled for VNICs in a VLAN.
	// Example: `true`
	SkipSourceDestCheck *bool `mandatory:"false" json:"skipSourceDestCheck"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet the VNIC is in.
	SubnetId *string `mandatory:"false" json:"subnetId"`

	// List of IPv6 addresses assigned to the VNIC.
	// Example: `2001:DB8::`
	Ipv6Addresses []string `mandatory:"false" json:"ipv6Addresses"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the route table the IP address or VNIC will use. For more information, see
	// Source Based Routing (https://docs.oracle.com/iaas/Content/Network/Tasks/managingroutetables.htm#Overview_of_Routing_for_Your_VCN__source_routing).
	RouteTableId *string `mandatory:"false" json:"routeTableId"`
}

func (m Vnic) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Vnic) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingVnicLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetVnicLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// VnicLifecycleStateEnum Enum with underlying type: string
type VnicLifecycleStateEnum string

// Set of constants representing the allowable values for VnicLifecycleStateEnum
const (
	VnicLifecycleStateProvisioning VnicLifecycleStateEnum = "PROVISIONING"
	VnicLifecycleStateAvailable    VnicLifecycleStateEnum = "AVAILABLE"
	VnicLifecycleStateTerminating  VnicLifecycleStateEnum = "TERMINATING"
	VnicLifecycleStateTerminated   VnicLifecycleStateEnum = "TERMINATED"
)

var mappingVnicLifecycleStateEnum = map[string]VnicLifecycleStateEnum{
	"PROVISIONING": VnicLifecycleStateProvisioning,
	"AVAILABLE":    VnicLifecycleStateAvailable,
	"TERMINATING":  VnicLifecycleStateTerminating,
	"TERMINATED":   VnicLifecycleStateTerminated,
}

var mappingVnicLifecycleStateEnumLowerCase = map[string]VnicLifecycleStateEnum{
	"provisioning": VnicLifecycleStateProvisioning,
	"available":    VnicLifecycleStateAvailable,
	"terminating":  VnicLifecycleStateTerminating,
	"terminated":   VnicLifecycleStateTerminated,
}

// GetVnicLifecycleStateEnumValues Enumerates the set of values for VnicLifecycleStateEnum
func GetVnicLifecycleStateEnumValues() []VnicLifecycleStateEnum {
	values := make([]VnicLifecycleStateEnum, 0)
	for _, v := range mappingVnicLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetVnicLifecycleStateEnumStringValues Enumerates the set of values in String for VnicLifecycleStateEnum
func GetVnicLifecycleStateEnumStringValues() []string {
	return []string{
		"PROVISIONING",
		"AVAILABLE",
		"TERMINATING",
		"TERMINATED",
	}
}

// GetMappingVnicLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVnicLifecycleStateEnum(val string) (VnicLifecycleStateEnum, bool) {
	enum, ok := mappingVnicLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
