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

type DeploymentView struct{}

func (dv *DeploymentView) New(obj *types.Deployment, pl map[string]*types.Pod) *Deployment {
	d := Deployment{}
	d.ID = obj.Meta.Name
	d.Meta = d.ToMeta(obj.Meta)
	//d.Status = d.ToStatus(obj.Status)
	d.Spec = d.ToSpec(obj.Spec)
	d.Replicas = d.ToReplicas(obj.Replicas)
	d.Pods = d.ToPods(pl)
	return &d
}

func (di *Deployment) ToMeta(obj types.DeploymentMeta) DeploymentMeta {
	meta := DeploymentMeta{}
	meta.Name = obj.Name
	meta.Description = obj.Description
	meta.Version = obj.Version
	//meta.SelfLink = obj.SelfLink
	//meta.Namespace = obj.Namespace
	//meta.Service = obj.Service
	meta.Status = obj.Status
	meta.Endpoint = obj.Endpoint
	meta.Updated = obj.Updated
	meta.Created = obj.Created

	return meta
}

//func (di *Deployment) ToStatus(obj types.DeploymentStatus) DeploymentStatusInfo {
//	return DeploymentStatusInfo{
//Status:   obj.Status,
//Message: obj.Message,
//}
//}

func (di *Deployment) ToSpec(obj types.DeploymentSpec) DeploymentSpec {

	var spec = DeploymentSpec{
		Strategy: obj.Strategy,
		Triggers: obj.Triggers,
		Selector: obj.Selector,
		Replicas: obj.Replicas,
		Template: obj.Template,
	}

	return spec
}

func (di *Deployment) ToReplicas(obj types.DeploymentReplicas) DeploymentReplicasInfo {

	return DeploymentReplicasInfo{
		Total:     obj.Total,
		Provision: obj.Provision,
		Created:   obj.Created,
		Started:   obj.Started,
		Stopped:   obj.Stopped,
		Errored:   obj.Errored,
		Pulling:   obj.Pulling,
	}
}

func (di *Deployment) ToPods(obj map[string]*types.Pod) map[string]Pod {
	pods := make(map[string]Pod, 0)
	//for _, p := range obj {
	//	if p.Meta.Deployment == di.ID {
	//		pv := new(PodViewHelper)
	//		pods[p.Meta.Name] = pv.New(p)
	//	}
	//}
	return pods
}

func (di *Deployment) ToJson() ([]byte, error) {
	return json.Marshal(di)
}

func (dv *DeploymentView) NewList(obj map[string]*types.Deployment, pods map[string]*types.Pod) *DeploymentList {
	dl := make(DeploymentList, 0)
	for _, d := range obj {
		dv := new(DeploymentView)
		dp := dv.New(d, pods)
		dl = append(dl, dp)
	}
	return &dl
}

func (di *DeploymentList) ToJson() ([]byte, error) {
	return json.Marshal(di)
}
