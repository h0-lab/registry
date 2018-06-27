//
// Last.Backend LLC CONFIDENTIAL
// __________________
//
// [2014] - [2018] Last.Backend LLC
// All Rights Reserved.
//
// NOTICE:  All information contained herein is, and remains
// the property of Last.Backend LLC and its suppliers,
// if any.  The intellectual and technical concepts contained
// herein are proprietary to Last.Backend LLC
// and its suppliers and may be covered by Russian Federation and Foreign Patents,
// patents in process, and are protected by trade secret or copyright law.
// Dissemination of this information or reproduction of this material
// is strictly forbidden unless prior written permission is obtained
// from Last.Backend LLC.
//

package views

import (
	"encoding/json"

	"github.com/lastbackend/registry/pkg/distribution/types"
)

type ClusterView struct{}

func (cv *ClusterView) New(obj *types.Cluster) *Cluster {
	c := Cluster{}
	c.Meta = cv.ToClusterMeta(obj.Meta)
	c.Status = cv.ToClusterStatus(obj.Status)
	return &c
}

func (cl *Cluster) ToJson() ([]byte, error) {
	return json.Marshal(cl)
}

func (cv *ClusterView) NewList(obj map[string]*types.Cluster) *ClusterList {
	if obj == nil {
		return nil
	}

	c := make(ClusterList, 0)
	for _, v := range obj {
		c = append(c, cv.New(v))
	}
	return &c
}

func (cl *ClusterList) ToJson() ([]byte, error) {
	if cl == nil {
		cl = &ClusterList{}
	}
	return json.Marshal(cl)
}

func (cv *ClusterView) ToClusterMeta(meta types.ClusterMeta) ClusterMeta {
	return ClusterMeta{
		Name:        meta.Name,
		Description: meta.Description,
		Region:      meta.Region,
		Provider:    meta.Provider,
		Labels:      meta.Labels,
		Created:     meta.Created,
		Updated:     meta.Updated,
	}
}

func (cv *ClusterView) ToClusterStatus(status types.ClusterStatus) ClusterStatus {
	return ClusterStatus{
		Nodes: status.Nodes,
		Capacity: ClusterResources{
			Containers: status.Capacity.Containers,
			Pods:       status.Capacity.Pods,
			Memory:     status.Capacity.Memory,
			Cpu:        status.Capacity.Cpu,
			Storage:    status.Capacity.Storage,
		},
		Allocated: ClusterResources{
			Containers: status.Allocated.Containers,
			Pods:       status.Allocated.Pods,
			Memory:     status.Allocated.Memory,
			Cpu:        status.Allocated.Cpu,
			Storage:    status.Allocated.Storage,
		},
		Deleted: status.Deleted,
	}
}
