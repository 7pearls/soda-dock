// Copyright (c) 2017 Huawei Technologies Co., Ltd. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
This module implements the common data structure.

*/

package model

type BaseModel struct {
	// The uuid of the object, it's unique in the context and generated by system
	// on successful creation of the object. It's not allowed to be modified by
	// the user.
	// +readOnly
	Id string `json:"id"`

	// CreateAt representing the server time when the object was created successfully.
	// Now, it's represented as a time string in RFC8601 format.
	// +readOnly
	CreatedAt string `json:"createdAt"`

	// UpdatedAt representing the server time when the object was updated successfully.
	// Now, it's represented as a time string in RFC8601 format.
	// +readOnly
	UpdatedAt string `json:"updatedAt"`
}

// DataStorageLoS can be used to describe a service option covering storage
// provisioning and availability.
type DataStorageLoS struct {
	// The enumeration literal specifies the time after a disaster that the
	// client shall regain conformant service level access to the primary
	// store.
	// The expectation is that the services required to implement this
	// capability are part of the advertising system.
	// +units:min
	RecoveryTimeObjective int64 `json:"recoveryTimeObjective,omitempty" yaml:"recoveryTimeObjective,omitempty"`

	// ProvisioningPolicy only supports "Fixed" and "Thin".
	ProvisioningPolicy string `json:"provisioningPolicy,omitempty" yaml:"provisioningPolicy,omitempty"`

	// IsSpaceEfficient indicates that the storage is compressed or deduplicated.
	// The default value for this prperty is false.
	IsSpaceEfficient bool `json:"isSpaceEfficient,omitempty" yaml:"isSpaceEfficient,omitempty"`
}

// IOConnectivityLoS can be used to specify the characteristics of storage
// connectivity.
type IOConnectivityLoS struct {
	// The Enumeration Literal shall specify the Access protocol for this
	// service option.
	AccessProtocol string `json:"accessProtocol,omitempty" yaml:"accessProtocol,omitempty"`

	// MaxIOPS shall be the maximum IOs per second that the connection shall
	// allow for the selected access protocol.
	// +units:[IO]/s
	MaxIOPS int64 `json:"maxIOPS,omitempty" yaml:"maxIOPS,omitempty"`

	// MaxBWS shall be the maximum amount of data that can be transmitted in a
	// fixed amount of time.
	// +units:[MB]/s
	MaxBWS int64 `json:"maxBWS,omitempty" yaml:"maxBWS,omitempty"`
}

// DataProtectionLoS describes a replica that protects data from loss. The
// requirements must be met collectively by the communication path and the
// replica.
// The expectation is that the services required to implement this
// capability are part of the advertising system.
type DataProtectionLoS struct {
	// IsIsolated shall indicate if the replica is in a separate fault domain.
	IsIsolated bool `json:"isIsolated,omitempty" yaml:"isIsolated,omitempty"`

	// MinLifetime shall be an ISO 8601 duration that specifies the minimum
	// required lifetime of the replica. For example, "P3Y6M4DT12H30M5S"
	// represents a duration of "3 years, 6 months, 4 days, 12 hours, 30 minutes
	// and 5 seconds".
	MinLifetime string `json:"minLifetime,omitempty" yaml:"minLifetime,omitempty"`

	// The enumeration literal specifies the geograhic scope of the failure
	// domain, and currently contains these options:
	// * Datacenter: A facility that provides communication, power, or cooling
	//   infrastructure to a co-located set of servers, networking and storage.
	// * Rack: A container within a datacenter that provides communication,
	//   power, or cooling to a set of components.
	// * RackGroup: A set of racks that may share common communication, power,
	//   or cooling.
	// * Region: A set of resources that are required to be either geographically
	//   or politically isolated from resources not in the resources.
	// * Row: A set of adjacent racks or rackgroups that may share common
	//   communication, power, or cooling.
	// * Server: Components of a CPU/memory complex that share the same
	//   infrastructure.
	RecoveryGeographicObject string `json:"recoveryGeographicObjective,omitempty" yaml:"recoveryGeographicObjective,omitempty"`

	// This value shall be an ISO 8601 duration that specifies the maximum time
	// over which source data may be lost on failure. For example,
	// "P3Y6M4DT12H30M5S" represents a duration of "3 years, 6 months, 4 days,
	// 12 hours, 30 minutes and 5 seconds".
	// In the case that IsIsolated = false, failure of the domain is not a
	// consideration.
	RecoveryPointObjectiveTime string `json:"recoveryPointObjectiveTime,omitempty" yaml:"recoveryPointObjectiveTime,omitempty"`

	// The enumeration literal specifies the time after a disaster that the
	// client shall regain conformant service level access to the primary
	// store. The possible values of this property could be:
	// * OnlineActive: Active access to synchronous replicas.
	// * OnlinePassive: Passive access to replicas via the same front-end
	//   interconnect.
	// * Nearline: Access to replica via a different front-end interconnect. A
	//   restore step is required before recovery can commence.
	// * Offline: No direct connection to the replica. (i.e. To a bunker
	//   containing backup media.)
	RecoveryTimeObjective string `json:"recoveryTimeObjective,omitempty" yaml:"recoveryTimeObjective,omitempty"`

	// The enumeration literals may be used to specify the intended outcome of
	// the replication. The possible values of this property could be:
	// * Clone: This enumeration literal shall indicate that replication shall
	//   create a point in time, full copy the source.
	// * Mirror: This enumeration literal shall indicate that replication shall
	//   create and maintain a copy of the source.
	// * Snapshot: This enumeration literal shall indicate that replication
	//   shall create a point in time, virtual copy of the source.
	// * TokenizedClone: This enumeration literal shall indicate that replication
	//   shall create a token based clone.
	ReplicaType string `json:"replicaType,omitempty" yaml:"replicaType,omitempty"`
}
