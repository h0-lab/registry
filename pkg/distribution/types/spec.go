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

package types

import (
	"fmt"
	"io"
)

const ContainerRolePrimary = "primary"
const ContainerRoleSlave = "slave"

// SpecState is a state of the spec
// swagger:model types_spec_state
type SpecState struct {
	Destroy     bool `json:"destroy"`
	Maintenance bool `json:"maintenance"`
}

// SpecTemplate is a template of the spec
// swagger:model types_spec_template
type SpecTemplate struct {
	// Network spec for service
	Network SpecTemplateNetwork `json:"network"`
	// Template spec for volume
	Volumes SpecTemplateVolumeList `json:"volumes"`
	// Template main container
	Containers SpecTemplateContainers `json:"container"`
	// Termination period
	Termination int `json:"termination"`
}

// swagger:ignore
// SpecTemplateNetwork is a map of spec template for network
// swagger:model types_spec_template_network
type SpecTemplateNetwork struct {
	IP     string            `json:"ip"`
	Ports  map[uint16]string `json:"ports"`
	Policy string            `json:"policy"`
}

// swagger:ignore
// SpecTemplateVolumeMap is a map of spec template volumes
// swagger:model types_spec_template_volume_map
type SpecTemplateVolumeMap map[string]*SpecTemplateVolume

// SpecTemplateVolumeList is a list of spec template volumes
// swagger:model types_spec_template_volume_list
type SpecTemplateVolumeList []*SpecTemplateVolume

// swagger:model types_spec_template_volume
type SpecTemplateVolume struct {
	// Template volume name
	Name string `json:"name"`
}

// swagger:ignore
// swagger:model types_spec_template_volume_mounts
type SpecTemplateVolumeMounts struct {
	// Template volume mounts name
	Name string `json:"name"`
}

// SpecTemplateContainers is a list of spec template containers
// swagger:model types_spec_template_container_list
type SpecTemplateContainers []SpecTemplateContainer

// swagger:model types_spec_template_container
type SpecTemplateContainer struct {
	// Template container id
	ID string `json:"id"`
	// Template container name
	Name string `json:"name"`
	// Template container role
	Role string `json:"role"`
	// Automatically remove container when it exits
	AutoRemove bool `json:"autoremove"`
	// Labels list
	Labels map[string]string `json:"labels"`
	// Template container image
	Image SpecTemplateContainerImage `json:"image"`
	// Template container ports binding
	Ports SpecTemplateContainerPorts `json:"ports"`
	// Template container envs
	EnvVars SpecTemplateContainerEnvs `json:"env"`
	// Template container resources
	Resources SpecTemplateContainerResources `json:"resources"`
	// Template container exec options
	Exec SpecTemplateContainerExec `json:"exec"`
	// Template container volumes
	Volumes SpecTemplateContainerVolumes `json:"volumes"`
	// Template container probes
	Probes SpecTemplateContainerProbes `json:"probes"`
	// Template container security
	Security SpecTemplateContainerSecurity `json:"security"`
	// Network container settings
	Network SpecTemplateContainerNetwork `json:"network"`
	// Container DNS configuration
	DNS SpecTemplateContainerDNS `json:"dns"`
	// List of extra hosts
	ExtraHosts []string `json:"extra_hosts"`
	// Should docker publish all exposed port for the container
	PublishAllPorts bool `json:"publish"`
	// Links to another containers
	Links []SpecTemplateContainerLink `json:"links"`
	// Restart Policy
	RestartPolicy SpecTemplateRestartPolicy `json:"restart"`
}

// swagger:model types_spec_template_container_image
type SpecTemplateContainerImage struct {
	Name   string `json:"name"`
	Auth   string `json:"auth"`
	Policy string `json:"policy"`
}

// swagger:ignore
// SpecBuildImage is an image of the spec build
// swagger:model types_spec_build_image
type SpecBuildImage struct {
	Tags           []string
	NoCache        bool
	Memory         int64
	Dockerfile     string
	SuppressOutput bool
	AuthConfigs    map[string]AuthConfig
	Context        io.Reader
	ExtraHosts     []string // List of extra hosts
}

// swagger:ignore
// AuthConfig contains authorization information for connecting to a Registry
// swagger:model types_authConfig
type AuthConfig struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Auth     string `json:"auth,omitempty"`
	// Email is an optional value associated with the username.
	// This field is deprecated and will be removed in a later
	// version of docker.
	Email         string `json:"email,omitempty"`
	ServerAddress string `json:"serveraddress,omitempty"`
	// IdentityToken is used to authenticate the user and get
	// an access token for the registry.
	IdentityToken string `json:"identitytoken,omitempty"`
	// RegistryToken is a bearer token to be sent to a registry
	RegistryToken string `json:"registrytoken,omitempty"`
}

// SpecTemplateContainerPorts is a list of spec template container ports
// swagger:model types_spec_template_container_port_list
type SpecTemplateContainerPorts []SpecTemplateContainerPort

// SpecTemplateContainerPort is a port of the spec template container
// swagger:model types_spec_template_container_port
type SpecTemplateContainerPort struct {
	// Container port
	ContainerPort uint16 `json:"container_port"`
	// Binding protocol
	Protocol string `json:"protocol"`
}

// SpecTemplateContainerPorts is a list of spec template container env vars
// swagger:model types_spec_template_container_env_list
type SpecTemplateContainerEnvs []SpecTemplateContainerEnv

// swagger:model types_spec_template_container_env
type SpecTemplateContainerEnv struct {
	Name  string                         `json:"name"`
	Value string                         `json:"value"`
	From  SpecTemplateContainerEnvSecret `json:"from"`
}

// swagger:model types_spec_template_container_env_secret
type SpecTemplateContainerEnvSecret struct {
	Name string `json:"name"`
	Key  string `json:"key"`
}

// swagger:model types_spec_template_container_resources
type SpecTemplateContainerResources struct {
	// Limit resources
	Limits SpecTemplateContainerResource `json:"limits"`
	// Request resources
	Request SpecTemplateContainerResource `json:"quota"`
}

// swagger:model types_spec_template_container_exec
type SpecTemplateContainerExec struct {
	Command []string `json:"command"`
	// Container enrtypoint
	Entrypoint []string `json:"entrypoint"`
	// Container run workdir option
	Workdir string `json:"workdir"`
	// Container run command arguments
	Args []string `json:"args"`
}

// swagger:model types_spec_template_container_resource
type SpecTemplateContainerResource struct {
	// CPU resource option
	CPU int64 `json:"cpu"`
	// RAM resource option
	RAM int64 `json:"ram"`
}

// SpecTemplateContainerVolumes is a list of spec template container volumes
// swagger:model types_spec_template_container_volume_list
type SpecTemplateContainerVolumes []SpecTemplateContainerVolume

// swagger:model types_spec_template_container_volume
type SpecTemplateContainerVolume struct {
	// Volume name
	Name string `json:"name"`
	// Volume mount mode
	Mode string `json:"mode"`
	// Volume mount path
	Path string `json:"path"`
}

// swagger:model types_spec_template_container_probes
type SpecTemplateContainerProbes struct {
	LiveProbe SpecTemplateContainerProbe `json:"live_probe"`
	ReadProbe SpecTemplateContainerProbe `json:"read_probe"`
}

// swagger:model types_spec_template_container_probe
type SpecTemplateContainerProbe struct {
	// Exec command to check container liveness
	Exec struct {
		Command []string `json:"command"`
	} `json:"exec"`

	Socket struct {
		Protocol string `json:"protocol"`
		Port     int    `json:"port"`
	} `json:"socket"`

	InitialDelaySeconds int `json:"initial_delay"`
	TimeoutSeconds      int `json:"timeout_seconds"`
	PeriodSeconds       int `json:"period_seconds"`
	ThresholdSuccess    int `json:"threshold_success"`
	ThresholdFailure    int `json:"threshold_failure"`
}

// swagger:model types_spec_template_container_security
type SpecTemplateContainerSecurity struct {
	// Start container in priveleged mode
	Privileged bool `json:"privileged"`
	// Add linux security options
	LinuxOptions SpecTemplateContainerSecurityLinuxOptions `json:"linux_options"`
	// Run container as particular user
	User int `json:"user"`
}

// swagger:model types_spec_template_container_security_linux
type SpecTemplateContainerSecurityLinuxOptions struct {
	Level string `json:"level"`
}

// swagger:model types_spec_template_container_network
type SpecTemplateContainerNetwork struct {
	// Container hostname
	Hostname string `json:"hostname"`
	// Container host domain
	Domain string `json:"domain"`
	// Network Hash to use
	Network string `json:"network"`
	// Network Mode to use
	Mode string `json:"mode"`
}

// swagger:model types_spec_template_container_dns
type SpecTemplateContainerDNS struct {
	// List of DNS servers
	Server []string `json:"server"`
	// DNS server search options
	Search []string `json:"search"`
	// DNS server other options
	Options []string `json:"options"`
}

// swagger:model types_spec_template_container_link
type SpecTemplateContainerLink struct {
	// Link name
	Link string `json:"link"`
	// Container alias
	Alias string `json:"alias"`
}

// swagger:model types_spec_template_policy
type SpecTemplateRestartPolicy struct {
	// Restart policy name
	Policy string `json:"policy"`
	// Attempt period
	Attempt int `json:"attempt"`
}

// swagger:model types_spec_strategy
type SpecStrategy struct {
	Type           string                     `json:"type"` // Rolling
	RollingOptions SpecStrategyRollingOptions `json:"rollingOptions"`
	Resources      SpecStrategyResources      `json:"resources"`
	Deadline       int                        `json:"deadline"`
}

// swagger:model types_spec_strategy_resources
type SpecStrategyResources struct {
}

// swagger:model types_spec_strategy_rolling
type SpecStrategyRollingOptions struct {
	PeriodUpdate   int `json:"period_update"`
	Interval       int `json:"interval"`
	Timeout        int `json:"timeout"`
	MaxUnavailable int `json:"max_unavailable"`
	MaxSurge       int `json:"max_surge"`
}

// SpecTriggers is a list of spec triggers
// swagger:model types_spec_trigger_list
type SpecTriggers []SpecTrigger

// swagger:model types_spec_trigger
type SpecTrigger struct {
}

// swagger:model types_spec_selector
type SpecSelector struct {
}

func (s *SpecTemplateContainerEnvs) ToLinuxFormat() []string {
	env := make([]string, 0)

	for _, e := range *s {
		env = append(env, fmt.Sprintf("%s=%s", e.Name, e.Value))
	}

	return env
}

func (s *SpecTemplate) SetDefault() {
	// Set default configurations

	s.Containers = make(SpecTemplateContainers, 1)
	s.Volumes = make(SpecTemplateVolumeList, 0)
}

func (s *SpecTemplateContainer) SetDefault() {
	s.Labels = make(map[string]string, 0)
	s.Resources.Limits.RAM = int64(128)
	s.Resources.Request.RAM = int64(128)
	s.EnvVars = make(SpecTemplateContainerEnvs, 0)
	s.Ports = make(SpecTemplateContainerPorts, 0)
	s.Volumes = make(SpecTemplateContainerVolumes, 0)
	s.Exec.Command = make([]string, 0)
	s.Exec.Entrypoint = make([]string, 0)
	s.Probes.LiveProbe.Exec.Command = make([]string, 0)
	s.Probes.ReadProbe.Exec.Command = make([]string, 0)
	s.RestartPolicy = SpecTemplateRestartPolicy{
		Policy: "always",
	}
}
