// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by schema-generate. DO NOT EDIT.

package model

import (
	"encoding/json"
)

// Root
type Root interface{}

type ESInitializer interface {
	ESInitialize(id string, seqno, version int64)
}

type ESDocument struct {
	Id      string `json:"-"`
	Version int64  `json:"-"`
	SeqNo   int64  `json:"-"`
}

func (d *ESDocument) ESInitialize(id string, seqno, version int64) {
	d.Id = id
	d.SeqNo = seqno
	d.Version = version
}

// Action An Elastic Agent action
type Action struct {
	ESDocument

	// The unique identifier for the Elastic Agent action. There could be multiple documents with the same action_id if the action is split into two separate documents.
	ActionID string `json:"action_id,omitempty"`

	// The Agent IDs the action is intended for. No support for json.RawMessage with the current generator. Could be useful to lazy parse the agent ids
	Agents []string `json:"agents,omitempty"`

	// The opaque payload.
	Data json.RawMessage `json:"data,omitempty"`

	// The action expiration date/time
	Expiration string `json:"expiration,omitempty"`

	// The input type the actions should be routed to.
	InputType string `json:"input_type,omitempty"`

	// The minimum time (in seconds) provided for an action execution when scheduled by fleet-server.
	MinimumExecutionDuration int64 `json:"minimum_execution_duration,omitempty"`

	// The rollout duration (in seconds) provided for an action execution when scheduled by fleet-server.
	RolloutDurationSeconds int64   `json:"rollout_duration_seconds,omitempty"`
	Signed                 *Signed `json:"signed,omitempty"`

	// The action start date/time
	StartTime string `json:"start_time,omitempty"`

	// The optional action timeout in seconds
	Timeout int64 `json:"timeout,omitempty"`

	// Date/time the action was created
	Timestamp string `json:"@timestamp,omitempty"`

	// APM traceparent for the action.
	Traceparent string `json:"traceparent,omitempty"`

	// The action type. INPUT_ACTION is the value for the actions that suppose to be routed to the endpoints/beats.
	Type string `json:"type,omitempty"`

	// The ID of the user who created the action.
	UserID string `json:"user_id,omitempty"`
}

// ActionResult An Elastic Agent action results
type ActionResult struct {
	ESDocument

	// The opaque payload.
	ActionData json.RawMessage `json:"action_data,omitempty"`

	// The action id.
	ActionID string `json:"action_id,omitempty"`

	// The input type of the original action.
	ActionInputType string `json:"action_input_type,omitempty"`

	// The custom action response payload.
	ActionResponse json.RawMessage `json:"action_response,omitempty"`

	// The agent id.
	AgentID string `json:"agent_id,omitempty"`

	// Date/time the action was completed
	CompletedAt string `json:"completed_at,omitempty"`

	// The opaque payload.
	Data json.RawMessage `json:"data,omitempty"`

	// The action error message.
	Error string `json:"error,omitempty"`

	// Date/time the action was started
	StartedAt string `json:"started_at,omitempty"`

	// Date/time the action was created
	Timestamp string `json:"@timestamp,omitempty"`
}

// Agent An Elastic Agent that has enrolled into Fleet
type Agent struct {
	ESDocument

	// ID of the API key the Elastic Agent must used to contact Fleet Server
	AccessAPIKeyID string `json:"access_api_key_id,omitempty"`

	// The last acknowledged action sequence number for the Elastic Agent
	ActionSeqNo []int64 `json:"action_seq_no,omitempty"`

	// Active flag
	Active bool           `json:"active"`
	Agent  *AgentMetadata `json:"agent,omitempty"`

	// Elastic Agent components detailed status information
	Components []ComponentsItems `json:"components,omitempty"`

	// Deprecated. Use Outputs instead. API key the Elastic Agent uses to authenticate with elasticsearch
	DefaultAPIKey string `json:"default_api_key,omitempty"`

	// Deprecated. Use Outputs instead. Default API Key History
	DefaultAPIKeyHistory []ToRetireAPIKeyIdsItems `json:"default_api_key_history,omitempty"`

	// Deprecated. Use Outputs instead. ID of the API key the Elastic Agent uses to authenticate with elasticsearch
	DefaultAPIKeyID string `json:"default_api_key_id,omitempty"`

	// Date/time the Elastic Agent enrolled
	EnrolledAt string `json:"enrolled_at"`

	// Enrollment ID
	EnrollmentID string `json:"enrollment_id,omitempty"`

	// Date/time the Elastic Agent checked in last time
	LastCheckin string `json:"last_checkin,omitempty"`

	// Last checkin message
	LastCheckinMessage string `json:"last_checkin_message,omitempty"`

	// Last checkin status
	LastCheckinStatus string `json:"last_checkin_status,omitempty"`

	// Date/time the Elastic Agent was last updated
	LastUpdated string `json:"last_updated,omitempty"`

	// Local metadata information for the Elastic Agent
	LocalMetadata json.RawMessage `json:"local_metadata,omitempty"`

	// Outputs is the policy output data, mapping the output name to its data
	Outputs map[string]*PolicyOutput `json:"outputs,omitempty"`

	// Packages array
	Packages []string `json:"packages,omitempty"`

	// The current policy coordinator for the Elastic Agent
	PolicyCoordinatorIdx int64 `json:"policy_coordinator_idx,omitempty"`

	// The policy ID for the Elastic Agent
	PolicyID string `json:"policy_id,omitempty"`

	// Deprecated. Use Outputs instead. The policy output permissions hash
	PolicyOutputPermissionsHash string `json:"policy_output_permissions_hash,omitempty"`

	// The current policy revision_idx for the Elastic Agent
	PolicyRevisionIdx int64 `json:"policy_revision_idx,omitempty"`

	// Shared ID
	SharedID string `json:"shared_id,omitempty"`

	// User provided tags for the Elastic Agent
	Tags []string `json:"tags,omitempty"`

	// Type
	Type string `json:"type"`

	// Date/time the Elastic Agent unenrolled
	UnenrolledAt string `json:"unenrolled_at,omitempty"`

	// Reason the Elastic Agent was unenrolled
	UnenrolledReason string `json:"unenrolled_reason,omitempty"`

	// Date/time the Elastic Agent unenrolled started
	UnenrollmentStartedAt string `json:"unenrollment_started_at,omitempty"`

	// Unhealthy reason: input/output/other
	UnhealthyReason []string `json:"unhealthy_reason,omitempty"`

	// Date/time the Elastic Agent was last updated
	UpdatedAt string `json:"updated_at,omitempty"`

	// Additional upgrade status details.
	UpgradeDetails *UpgradeDetails `json:"upgrade_details,omitempty"`

	// Date/time the Elastic Agent started the current upgrade
	UpgradeStartedAt string `json:"upgrade_started_at,omitempty"`

	// Upgrade status
	UpgradeStatus string `json:"upgrade_status,omitempty"`

	// Date/time the Elastic Agent was last upgraded
	UpgradedAt string `json:"upgraded_at,omitempty"`

	// User provided metadata information for the Elastic Agent
	UserProvidedMetadata json.RawMessage `json:"user_provided_metadata,omitempty"`
}

// AgentMetadata An Elastic Agent metadata
type AgentMetadata struct {

	// The unique identifier for the Elastic Agent
	ID string `json:"id"`

	// The version of the Elastic Agent
	Version string `json:"version"`
}

// Artifact An artifact served by Fleet
type Artifact struct {
	ESDocument

	// Encoded artifact data
	Body json.RawMessage `json:"body"`

	// Name of compression algorithm applied to artifact
	CompressionAlgorithm string `json:"compression_algorithm,omitempty"`

	// Timestamp artifact was created
	Created string `json:"created"`

	// SHA256 of artifact before encoding has been applied
	DecodedSha256 string `json:"decoded_sha256,omitempty"`

	// Size of artifact before encoding has been applied
	DecodedSize int64 `json:"decoded_size,omitempty"`

	// SHA256 of artifact after encoding has been applied
	EncodedSha256 string `json:"encoded_sha256,omitempty"`

	// Size of artifact after encoding has been applied
	EncodedSize int64 `json:"encoded_size,omitempty"`

	// Name of encryption algorithm applied to artifact
	EncryptionAlgorithm string `json:"encryption_algorithm,omitempty"`

	// Human readable artifact identifier
	Identifier string `json:"identifier"`

	// Name of the package that owns this artifact
	PackageName string `json:"package_name,omitempty"`
}

// Checkin An Elastic Agent checkin to Fleet
type Checkin struct {
	ESDocument
	Agent *AgentMetadata `json:"agent"`
	Host  *HostMetadata  `json:"host,omitempty"`

	// The current overall status message of the Elastic Agent
	Message string `json:"message"`

	// The current status of the applied policy
	Policy *CheckinPolicy  `json:"policy,omitempty"`
	Server *ServerMetadata `json:"server,omitempty"`

	// The current overall status of the Elastic Agent
	Status string `json:"status"`

	// Date/time the checkin was created
	Timestamp string `json:"@timestamp,omitempty"`
}

// CheckinPolicy The current status of the applied policy
type CheckinPolicy struct {

	// The ID for the policy
	ID string `json:"id"`

	// The current input status per policy
	Inputs []CheckinPolicyInputItems `json:"inputs"`

	// The revision of the policy
	Revision int64 `json:"revision"`
}

// CheckinPolicyInputItems
type CheckinPolicyInputItems struct {

	// The ID for the input
	ID string `json:"id"`

	// The current status message of the intput
	Message string `json:"message"`

	// The current status of the input
	Status string `json:"status"`

	// The template ID for the input
	TemplateID string `json:"template_id"`
}

// ComponentsItems
type ComponentsItems struct {
	ID      string       `json:"id,omitempty"`
	Message string       `json:"message,omitempty"`
	Status  string       `json:"status,omitempty"`
	Units   []UnitsItems `json:"units,omitempty"`
}

// DataStream
type DataStream struct {
	Dataset   string `json:"dataset,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	Type      string `json:"type,omitempty"`
}

// EnrollmentAPIKey An Elastic Agent enrollment API key
type EnrollmentAPIKey struct {
	ESDocument

	// Api key
	APIKey string `json:"api_key"`

	// The unique identifier for the enrollment key, currently xid
	APIKeyID string `json:"api_key_id"`

	// True when the key is active
	Active    bool   `json:"active,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	ExpireAt  string `json:"expire_at,omitempty"`

	// Enrollment key name
	Name      string `json:"name,omitempty"`
	PolicyID  string `json:"policy_id,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

// HostMetadata The host metadata for the Elastic Agent
type HostMetadata struct {

	// The architecture for the Elastic Agent
	Architecture string `json:"architecture"`

	// The ID of the host
	ID string `json:"id"`

	// The IP addresses of the Elastic Agent
	Ip []string `json:"ip,omitempty"`

	// The hostname of the Elastic Agent
	Name string `json:"name"`
}

// OutputHealth Output health represents a health state of an output
type OutputHealth struct {
	ESDocument
	DataStream *DataStream `json:"data_stream,omitempty"`

	// Long state message if unhealthy
	Message string `json:"message,omitempty"`

	// Output ID
	Output string `json:"output,omitempty"`

	// Health state, can be HEALTHY or DEGRADED
	State string `json:"state,omitempty"`

	// Timestamp of reported state
	Timestamp string `json:"@timestamp,omitempty"`
}

// Policy A policy that an Elastic Agent is attached to
type Policy struct {
	ESDocument

	// The coordinator index of the policy
	CoordinatorIdx int64       `json:"coordinator_idx"`
	Data           *PolicyData `json:"data"`

	// True when this policy is the default policy to start Fleet Server
	DefaultFleetServer bool `json:"default_fleet_server"`

	// The ID of the policy
	PolicyID string `json:"policy_id"`

	// The revision index of the policy
	RevisionIdx int64 `json:"revision_idx"`

	// Date/time the policy revision was created
	Timestamp string `json:"@timestamp,omitempty"`

	// Timeout (seconds) that an Elastic Agent should be un-enrolled.
	UnenrollTimeout int64 `json:"unenroll_timeout,omitempty"`
}

// PolicyData The policy data that an agent needs to run
type PolicyData struct {

	// The policy's agent configuration details
	Agent json.RawMessage `json:"agent,omitempty"`

	// The policy's fleet configuration details
	Fleet json.RawMessage `json:"fleet,omitempty"`

	// The policy's ID
	ID string `json:"id"`

	// A list of all inputs the agent should run
	Inputs []map[string]interface{} `json:"inputs,omitempty"`

	// The Elasticsearch permissions needed to run the policy
	OutputPermissions json.RawMessage `json:"output_permissions,omitempty"`

	// A map of all outputs that the agent running the policy can use to send data to.
	Outputs map[string]map[string]interface{} `json:"outputs"`

	// The policy revision number. Should match revision_idx
	Revision int64 `json:"revision"`

	// A list of all secrets fleet-server needs to inject into the policy before passing it to the agent. This attribute is removed when policy data is send to an agent.
	SecretReferences []SecretReferencesItems `json:"secret_references,omitempty"`
	Signed           *Signed                 `json:"signed,omitempty"`
}

// PolicyLeader The current leader Fleet Server for a policy
type PolicyLeader struct {
	ESDocument
	Server *ServerMetadata `json:"server"`

	// Date/time the leader was taken or held
	Timestamp string `json:"@timestamp,omitempty"`
}

// PolicyOutput holds the needed data to manage the output API keys
type PolicyOutput struct {
	ESDocument

	// API key the Elastic Agent uses to authenticate with elasticsearch
	APIKey string `json:"api_key"`

	// ID of the API key the Elastic Agent uses to authenticate with elasticsearch
	APIKeyID string `json:"api_key_id"`

	// The policy output permissions hash
	PermissionsHash string `json:"permissions_hash"`

	// API keys to be invalidated on next agent ack
	ToRetireAPIKeyIds []ToRetireAPIKeyIdsItems `json:"to_retire_api_key_ids,omitempty"`

	// Type is the output type. Currently only Elasticsearch is supported.
	Type string `json:"type"`
}

// SecretReferencesItems
type SecretReferencesItems struct {
	ID string `json:"id"`
}

// Server A Fleet Server
type Server struct {
	ESDocument
	Agent  *AgentMetadata  `json:"agent"`
	Host   *HostMetadata   `json:"host"`
	Server *ServerMetadata `json:"server"`

	// Date/time the server was updated
	Timestamp string `json:"@timestamp,omitempty"`
}

// ServerMetadata A Fleet Server metadata
type ServerMetadata struct {

	// The unique identifier for the Fleet Server
	ID string `json:"id"`

	// The version of the Fleet Server
	Version string `json:"version"`
}

// Signed The action signed data and signature.
type Signed struct {

	// The base64 encoded, UTF-8 JSON serialized action bytes that are signed.
	Data string `json:"data"`

	// The base64 encoded signature.
	Signature string `json:"signature"`
}

// ToRetireAPIKeyIdsItems the Output API Keys that were replaced and should be retired
type ToRetireAPIKeyIdsItems struct {

	// API Key identifier
	ID string `json:"id,omitempty"`

	// Output name where the API Key belongs
	Output string `json:"output,omitempty"`

	// Date/time the API key was retired
	RetiredAt string `json:"retired_at,omitempty"`
}

// UnitsItems
type UnitsItems struct {
	ID      string `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
	Status  string `json:"status,omitempty"`
	Type    string `json:"type,omitempty"`
}

// UpgradeDetails Additional upgrade status details.
type UpgradeDetails struct {
}
