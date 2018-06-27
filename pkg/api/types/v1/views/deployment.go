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
	"time"

	"encoding/json"
	"github.com/lastbackend/registry/pkg/distribution/types"
	"github.com/lastbackend/registry/pkg/log"
	"io"
	"io/ioutil"
)

type Deployment struct {
	ID       string                 `json:"id"`
	Meta     DeploymentMeta         `json:"meta"`
	Status   DeploymentStatusInfo   `json:"status"`
	Spec     DeploymentSpec         `json:"spec"`
	Sources  DeploymentSources      `json:"sources"`
	Replicas DeploymentReplicasInfo `json:"replicas"`
	Pods     map[string]Pod         `json:"pods"`
}

type DeploymentMap map[string]*Deployment
type DeploymentList []*Deployment

type DeploymentMeta struct {
	Name        string `json:"name"`
	Description string `json:"description"`

	Version   int    `json:"version"`
	Namespace string `json:"namespace"`
	Service   string `json:"service"`
	Endpoint  string `json:"endpoint"`
	SelfLink  string `json:"self_link"`
	Status    string `json:"status"`

	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

type DeploymentSources struct {
	// Image namespace name
	Namespace string `json:"namespace"`
	// Image tag
	Tag string `json:"tag"`
	// Hash
	Hash string `json:"hash"`
}

type DeploymentSpec struct {
	Strategy types.SpecStrategy `json:"strategy"`
	Triggers types.SpecTriggers `json:"triggers"`
	Replicas int                `json:"replicas"`
	Selector types.SpecSelector `json:"selector"`
	Template types.SpecTemplate `json:"template"`
}

type DeploymentStatusInfo struct {
	State   string `json:"state"`
	Message string `json:"message"`
}

type DeploymentReplicasInfo struct {
	Total     int `json:"total"`
	Provision int `json:"provision"`
	Pulling   int `json:"pulling"`
	Created   int `json:"created"`
	Started   int `json:"started"`
	Stopped   int `json:"stopped"`
	Errored   int `json:"errored"`
}

type RequestDeploymentScaleOptions struct {
	Replicas *int `json:"replicas"`
}

func (s *RequestDeploymentScaleOptions) DecodeAndValidate(reader io.Reader) (types.DeploymentOptions, error) {

	opts := types.DeploymentOptions{}

	log.V(logLevel).Debug("Request: Deployment: decode and validate data for creating")

	body, err := ioutil.ReadAll(reader)
	if err != nil {
		log.V(logLevel).Errorf("Request: Deployment: decode and validate data for creating err: %s", err)
		return opts, err
	}

	err = json.Unmarshal(body, s)
	if err != nil {
		log.V(logLevel).Errorf("Request: Deployment: convert struct from json err: %s", err)
		return opts, err
	}

	//opts.Replicas = *s.Replicas

	//if s.Replicas != nil && *s.Replicas < 1 {
	//	opts.Replicas = 1
	//}

	return opts, nil
}
