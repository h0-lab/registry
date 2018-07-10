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

package v1

import (
	"context"

	"github.com/lastbackend/registry/pkg/distribution/errors"
	"github.com/lastbackend/registry/pkg/api/client/http/request"
	vv1 "github.com/lastbackend/registry/pkg/api/types/v1/views"
)

type RegistryClient struct {
	client *request.RESTClient
}

func (rc RegistryClient) Get(ctx context.Context) (*vv1.Registry, error) {

	var s *vv1.Registry
	var e *errors.Http

	err := rc.client.Get("/registry").
		AddHeader("Content-Type", "application/json").
		JSON(&s, &e)

	if err != nil {
		return nil, err
	}
	if e != nil {
		return nil, errors.New(e.Message)
	}

	return s, nil
}

func newRegistryClient(req *request.RESTClient) RegistryClient {
	return RegistryClient{client: req}
}
