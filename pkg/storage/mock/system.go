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

package mock

import (
	"context"

	"github.com/lastbackend/lastbackend/pkg/log"
	"github.com/lastbackend/registry/pkg/distribution/types"
	"github.com/lastbackend/registry/pkg/storage/storage"
)

type SystemStorage struct {
	storage.System
}

func (s *SystemStorage) Get(ctx context.Context) (*types.System, error) {
	log.V(logLevel).Debugf("%s:system:get> get system", logPrefix)
	return nil, nil
}

func (s *SystemStorage) Update(ctx context.Context, system *types.System) error {
	log.V(logLevel).Debugf("%s:system:update> update system %#v", logPrefix, system)
	return nil
}

func newSystemStorage() *SystemStorage {
	s := new(SystemStorage)
	return s
}
