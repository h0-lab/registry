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

import "time"

type Secret struct {
	Meta SecretMeta `json:"meta"`
	Data string     `json:"data"`
}

type SecretMeta struct {
	Name      string    `json:"name"`
	Namespace string    `json:"namespace"`
	SelfLink  string    `json:"self_link"`
	Updated   time.Time `json:"updated"`
	Created   time.Time `json:"created"`
}

type SecretMap map[string]*Secret
type SecretList []*Secret
