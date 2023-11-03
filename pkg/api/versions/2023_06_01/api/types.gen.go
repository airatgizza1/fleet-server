// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.0.0 DO NOT EDIT.
package api

import (
	"encoding/json"
	"fmt"
)

const (
	AgentApiKeyScopes = "agentApiKey.Scopes"
	ApiKeyScopes      = "apiKey.Scopes"
)

// Defines values for CheckinRequestStatus.
const (
	CheckinRequestStatusDegraded CheckinRequestStatus = "degraded"
	CheckinRequestStatusError    CheckinRequestStatus = "error"
	CheckinRequestStatusOnline   CheckinRequestStatus = "online"
	CheckinRequestStatusStarting CheckinRequestStatus = "starting"
)

// Defines values for EnrollRequestType.
const (
	PERMANENT EnrollRequestType = "PERMANENT"
)

// Defines values for EventSubtype.
const (
	ACKNOWLEDGED EventSubtype = "ACKNOWLEDGED"
	CONFIG       EventSubtype = "CONFIG"
	DATADUMP     EventSubtype = "DATA_DUMP"
	FAILED       EventSubtype = "FAILED"
	INPROGRESS   EventSubtype = "IN_PROGRESS"
	RUNNING      EventSubtype = "RUNNING"
	STARTING     EventSubtype = "STARTING"
	STOPPED      EventSubtype = "STOPPED"
	STOPPING     EventSubtype = "STOPPING"
	UNKNOWN      EventSubtype = "UNKNOWN"
)

// Defines values for EventType.
const (
	ACTION       EventType = "ACTION"
	ACTIONRESULT EventType = "ACTION_RESULT"
	ERROR        EventType = "ERROR"
	STATE        EventType = "STATE"
)

// Defines values for StatusResponseStatus.
const (
	StatusResponseStatusConfiguring StatusResponseStatus = "configuring"
	StatusResponseStatusDegraded    StatusResponseStatus = "degraded"
	StatusResponseStatusFailed      StatusResponseStatus = "failed"
	StatusResponseStatusHealthy     StatusResponseStatus = "healthy"
	StatusResponseStatusStarting    StatusResponseStatus = "starting"
	StatusResponseStatusStopped     StatusResponseStatus = "stopped"
	StatusResponseStatusStopping    StatusResponseStatus = "stopping"
	StatusResponseStatusUnknown     StatusResponseStatus = "unknown"
)

// Defines values for UploadBeginRequestSrc.
const (
	Agent    UploadBeginRequestSrc = "agent"
	Endpoint UploadBeginRequestSrc = "endpoint"
)

// AckRequest The request an elastic-agent sends to fleet-serve to acknowledge the execution of one or more actions.
type AckRequest struct {
	Events []Event `json:"events"`
}

// AckResponse Response to processing acknowledgement events.
type AckResponse struct {
	// Action The action result. Will have the value "acks".
	Action string `json:"action"`

	// Errors A flag to indicate if one or more errors occured when proccessing events.
	Errors bool `json:"errors,omitempty"`

	// Items The in-order list of results from processing events.
	Items []AckResponseItem `json:"items,omitempty"`
}

// AckResponseItem The results of processing an acknowledgement event.
type AckResponseItem struct {
	// Message HTTP status text.
	Message *string `json:"message,omitempty"`

	// Status An HTTP status code that indicates if the event was processed successfully or not.
	Status int `json:"status"`
}

// Action An action for an elastic-agent.
// The actions are defined in generic terms on the fleet-server.
// The elastic-agent will have additional details for what is expected when a specific action-type is recieved.
// Many attributes in this schema also contain yaml tags so the elastic-agent may serialize them.
// The structure of the `data` attribute will vary between action types.
//
// An additional consideration is Scheduled Actions. Scheduled actions are currently defined as actions that have non-empty values for both the `start_time` and `expiration` attributes.
type Action struct {
	// AgentId The agent ID.
	AgentId string `json:"agent_id"`

	// CreatedAt Time when the action was created.
	CreatedAt string `json:"created_at"`

	// Data An embedded action-specific object.
	Data interface{} `json:"data" yaml:"data"`

	// Expiration The latest start time for the action. Actions will be dropped by the agent if execution has not started by this time. Used for scheduled actions.
	Expiration *string `json:"expiration,omitempty" yaml:"expiration"`

	// Id The action ID.
	Id string `json:"id" yaml:"action_id"`

	// InputType The input type of the action for actions with type `INPUT_ACTION`.
	InputType string `json:"input_type" yaml:"input_type"`

	// Signed Optional action signing data.
	Signed *ActionSignature `json:"signed,omitempty" yaml:"signed"`

	// StartTime The earliest execution time for the action. Agent will not execute the action before this time. Used for scheduled actions.
	StartTime *string `json:"start_time,omitempty" yaml:"start_time"`

	// Timeout The timeout value (in seconds) for actions with type `INPUT_ACTION`.
	Timeout *int64 `json:"timeout,omitempty" yaml:"timeout"`

	// Traceparent APM traceparent for the action.
	Traceparent *string `json:"traceparent,omitempty" yaml:"traceparent"`

	// Type The action type.
	Type string `json:"type" yaml:"type"`
}

// ActionSignature Optional action signing data.
type ActionSignature struct {
	// Data The base64 encoded, UTF-8 JSON serialized action bytes that are signed.
	Data string `json:"data,omitempty" yaml:"data"`

	// Signature The base64 encoded signature.
	Signature string `json:"signature,omitempty" yaml:"signature"`
}

// CheckinRequest defines model for checkinRequest.
type CheckinRequest struct {
	// AckToken The ack_token form a previous response if the agent has checked in before.
	// Translated to a sequence number in fleet-server in order to retrieve any new actions for the agent from the last checkin.
	AckToken *string `json:"ack_token,omitempty"`

	// Components An embedded JSON object that holds component information that the agent is running.
	// Defined in fleet-server as a `json.RawMessage`, defined as an object in the elastic-agent.
	// fleet-server will update the components in an agent record if they differ from this object.
	Components *json.RawMessage `json:"components,omitempty"`

	// LocalMetadata An embedded JSON object that holds meta-data values.
	// Defined in fleet-server as a `json.RawMessage`, defined as an object in the elastic-agent.
	// elastic-agent will populate the object with information from the binary and host/system environment.
	// fleet-server will update the agent record if a checkin response contains different data from the record.
	LocalMetadata *json.RawMessage `json:"local_metadata,omitempty"`

	// Message State message, may be overridden or use the error message of a failing component.
	Message string `json:"message"`

	// PollTimeout An optional timeout value that informs fleet-server of when a client will time out on it's checkin request.
	// If not specified fleet-server will use the timeout values specified in the config (defaults to 5m polling and a 10m write timeout).
	// The value, if specified is expected to be a string that is parsable by [time.ParseDuration](https://pkg.go.dev/time#ParseDuration).
	// If specified fleet-server will set its poll timeout to `max(1m, poll_timeout-2m)` and its write timeout to `max(2m, poll_timout-1m)`.
	PollTimeout *string `json:"poll_timeout,omitempty"`

	// Status The agent state, inferred from agent control protocol states.
	Status CheckinRequestStatus `json:"status"`
}

// CheckinRequestStatus The agent state, inferred from agent control protocol states.
type CheckinRequestStatus string

// CheckinResponse defines model for checkinResponse.
type CheckinResponse struct {
	// AckToken The acknowlegment token used to indicate action delivery.
	AckToken *string `json:"ack_token,omitempty"`

	// Action The action result. Set to "checkin".
	Action string `json:"action"`

	// Actions A list of actions that the agent must execute.
	Actions *[]Action `json:"actions,omitempty"`
}

// EnrollMetadata Metadata associated with the agent that is enrolling to fleet.
type EnrollMetadata struct {
	// Local An embedded JSON object that holds meta-data values.
	// Defined in fleet-server as a `json.RawMessage`, defined as an object in the elastic-agent.
	// elastic-agent will populate the object with information from the binary and host/system environment.
	// If not empty fleet-server will update the value of `local["elastic"]["agent"]["id"]` to the agent ID (assuming the keys exist).
	// The (possibly updated) value is sent by fleet-server when creating the record for a new agent.
	Local json.RawMessage `json:"local"`

	// Tags User provided tags for the agent.
	// fleet-server will pass the tags to the agent record on enrollment.
	Tags []string `json:"tags"`

	// UserProvided An embedded JSON object that holds user-provided meta-data values.
	// Defined in fleet-server as a `json.RawMessage`.
	// fleet-server does not use these values on enrollment of an agent.
	//
	// Defined in the elastic-agent as a `map[string]interface{}` with no way to specify any values.
	// Deprecated:
	UserProvided json.RawMessage `json:"user_provided"`
}

// EnrollRequest A request to enroll a new agent into fleet.
type EnrollRequest struct {
	// Metadata Metadata associated with the agent that is enrolling to fleet.
	Metadata EnrollMetadata `json:"metadata"`

	// SharedId The shared ID of the agent.
	// To support pre-existing installs.
	//
	// Never implemented.
	// Deprecated:
	SharedId string `json:"shared_id"`

	// Type The enrollment type of the agent.
	// The agent only supports the PERMANENT value.
	// In the future the enrollment type may be used to indicate agents that use fleet for reporting and monitoring, but do not use policies.
	Type EnrollRequestType `json:"type"`
}

// EnrollRequestType The enrollment type of the agent.
// The agent only supports the PERMANENT value.
// In the future the enrollment type may be used to indicate agents that use fleet for reporting and monitoring, but do not use policies.
type EnrollRequestType string

// EnrollResponse The enrollment action response.
type EnrollResponse struct {
	// Action The action result. Will have the value "created".
	Action string `json:"action"`

	// Item Response to a successful enrollment of an agent into fleet.
	Item EnrollResponseItem `json:"item"`
}

// EnrollResponseItem Response to a successful enrollment of an agent into fleet.
type EnrollResponseItem struct {
	// AccessApiKey The ApiKey token that fleet-server has generated for the enrolling agent.
	AccessApiKey string `json:"access_api_key"`

	// AccessApiKeyId The id of the ApiKey that fleet-server has generated for the enrolling agent.
	AccessApiKeyId string `json:"access_api_key_id"`

	// Actions Defined in fleet-server and elastic-agent as `[]interface{}`.
	//
	// Never used by agent.
	// Deprecated:
	Actions []map[string]interface{} `json:"actions"`

	// Active If the agent is active in fleet.
	// Set to true upon enrollment.
	//
	// Handling of other values never implemented.
	// Deprecated:
	Active bool `json:"active"`

	// EnrolledAt The RFC3339 timestamp that the agent was enrolled at.
	EnrolledAt string `json:"enrolled_at"`

	// Id The agent ID
	Id string `json:"id"`

	// LocalMetadata A copy of the (updated) local metadata provided in the enrollment request.
	//
	// Never used by agent.
	// Deprecated:
	LocalMetadata json.RawMessage `json:"local_metadata"`

	// PolicyId The policy ID that the agent is enrolled with. Decoded from the API key used in the request.
	PolicyId string `json:"policy_id"`

	// Status Agent status from fleet-server.
	// fleet-ui may differ.
	//
	// Never used by agent.
	// Deprecated:
	Status string `json:"status"`

	// Tags A copy of the tags that were sent with the enrollment request.
	Tags []string `json:"tags"`

	// Type The enrollment request type.
	//
	// Handling of other values never implemented.
	// Deprecated:
	Type string `json:"type"`

	// UserProvidedMetadata A copy of the user provided metadata from the enrollment request.
	//
	// Currently will be empty.
	// Deprecated:
	UserProvidedMetadata json.RawMessage `json:"user_provided_metadata"`
}

// Error Error processing request.
type Error struct {
	// Error Error type.
	Error string `json:"error"`

	// Message (optional) Error message.
	Message *string `json:"message,omitempty"`

	// StatusCode The HTTP status code of the error.
	StatusCode int `json:"statusCode"`
}

// Event The ack for a specific action that the elastic-agent has executed.
type Event struct {
	// ActionData The action data for the input action being acknowledged.
	ActionData *json.RawMessage `json:"action_data,omitempty"`

	// ActionId The action ID.
	ActionId string `json:"action_id"`

	// ActionInputType The input_type of the action for input actions.
	ActionInputType string `json:"action_input_type"`

	// ActionResponse The action response for the input action being acknowledged.
	ActionResponse *json.RawMessage `json:"action_response,omitempty"`

	// AgentId The ID of the agent that executed the action.
	AgentId string `json:"agent_id"`

	// CompletedAt The time at which the action was completed. Used only when acknowledging input actions
	CompletedAt string `json:"completed_at"`

	// Data An embedded JSON object that has the data about the ack.
	//
	// Used by REQUEST_DIAGNOSTICS actions.
	// Contains a `upload_id` attribute used to communicate the successfullly uploaded diagnostics ID.
	Data *json.RawMessage `json:"data,omitempty"`

	// Error An error message.
	// If this is non-empty an error has occured when executing the action.
	// For some actions (such as UPGRADE actions) it may result in the action being marked as failed.
	Error *string `json:"error,omitempty"`

	// Message An acknowlegement message. The elastic-agent inserts the action ID and action type into this message.
	Message string `json:"message"`

	// Payload An embedded JSON object that contains additional information for the fleet-server to process.
	// Defined as a json.RawMessage in both the fleet-server and the elastic-agent.
	//
	// Is currently used by UPGRADE actions to signal retries.
	// If the error attribute is non empty payload is checked for `retry: bool` and `retry_attempt: int`.
	// If retry is true, fleet-serve will mark the agent as retrying, if it's false the upgrade will be marked as failed.
	//
	// Additional action status information can be provided in the data attribute.
	// Deprecated:
	Payload *json.RawMessage `json:"payload,omitempty"`

	// PolicyId Not used by the fleet-server.
	// Deprecated:
	PolicyId string `json:"policy_id"`

	// StartedAt The time at which the action was started. Used only when acknowledging input actions.
	StartedAt string `json:"started_at"`

	// StreamId Not used by the fleet-server.
	// Deprecated:
	StreamId string `json:"stream_id"`

	// Subtype The subtype of the ack event.
	// The elastic-agent will only generate ACKNOWLEDGED events.
	//
	// Not used by fleet-server.
	// Actions that have errored should use the error attribute to communicate an error status.
	// Additional action status information can be provided in the data attribute.
	// Deprecated:
	Subtype EventSubtype `json:"subtype"`

	// Timestamp The timestamp of the acknowledgement event. Has the format of "2006-01-02T15:04:05.99999-07:00"
	Timestamp string `json:"timestamp"`

	// Type The event type of the ack.
	// Currently the elastic-agent will only generate ACTION_RESULT events.
	//
	// Not used by fleet-server.
	// Actions that have errored should use the error attribute to communicate an error status.
	// Additional action status information can be provided in the data attribute.
	// Deprecated:
	Type EventType `json:"type"`
}

// EventSubtype The subtype of the ack event.
// The elastic-agent will only generate ACKNOWLEDGED events.
//
// Not used by fleet-server.
// Actions that have errored should use the error attribute to communicate an error status.
// Additional action status information can be provided in the data attribute.
type EventSubtype string

// EventType The event type of the ack.
// Currently the elastic-agent will only generate ACTION_RESULT events.
//
// Not used by fleet-server.
// Actions that have errored should use the error attribute to communicate an error status.
// Additional action status information can be provided in the data attribute.
type EventType string

// StatusAPIResponse Status response information.
type StatusAPIResponse struct {
	// Name Service name.
	Name string `json:"name"`

	// Status A Unit state that fleet-server may report.
	// Unit state is defined in the elastic-agent-client specification.
	Status StatusResponseStatus `json:"status"`

	// Version Version information included in the response to an authorized status request.
	Version *StatusResponseVersion `json:"version,omitempty"`
}

// StatusResponseStatus A Unit state that fleet-server may report.
// Unit state is defined in the elastic-agent-client specification.
type StatusResponseStatus string

// StatusResponseVersion Version information included in the response to an authorized status request.
type StatusResponseVersion struct {
	// BuildHash The commit that the fleet-server was built from.
	BuildHash *string `json:"build_hash,omitempty"`

	// BuildTime The date-time that the fleet-server binary was created.
	BuildTime *string `json:"build_time,omitempty"`

	// Number The fleet-server version.
	Number *string `json:"number,omitempty"`
}

// UploadBeginRequest defines model for uploadBeginRequest.
type UploadBeginRequest struct {
	// ActionId ID of the action that requested this file
	ActionId string `json:"action_id"`

	// AgentId Identifier of the agent uploading. Matches the ID usually found in agent.id
	AgentId string                  `json:"agent_id"`
	File    UploadBeginRequest_File `json:"file"`

	// Src The source integration sending this file
	Src                  UploadBeginRequestSrc  `json:"src"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// UploadBeginRequest_File defines model for UploadBeginRequest.File.
type UploadBeginRequest_File struct {
	// Compression The algorithm used to compress the file. Valid values: br,gzip,deflate,none
	Compression *string `json:"Compression,omitempty"`

	// Hash Checksums on the file contents
	Hash *struct {
		// Sha256 SHA256 of the contents
		Sha256 *string `json:"sha256,omitempty"`
	} `json:"hash,omitempty"`

	// MimeType MIME type of the file
	MimeType string `json:"mime_type"`

	// Name Name of the file including the extension, without the directory
	Name string `json:"name"`

	// Size Size of the file contents, in bytes
	Size                 int64                  `json:"size"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// UploadBeginRequestSrc The source integration sending this file
type UploadBeginRequestSrc string

// UploadBeginAPIResponse Response to initiating a file upload
type UploadBeginAPIResponse struct {
	// ChunkSize The required size (in bytes) that the file must be segmented into for each chunk
	ChunkSize int64 `json:"chunk_size"`

	// UploadId A unique identifier for the ensuing upload operation
	UploadId string `json:"upload_id"`
}

// UploadCompleteRequest Request to verify and finish an uploaded file
type UploadCompleteRequest struct {
	// Transithash the transithash (sha256 of the concatenation of each in-order chunk hash) of the entire file contents
	Transithash struct {
		// Sha256 SHA256 hash
		Sha256 string `json:"sha256"`
	} `json:"transithash"`
}

// ApiVersion defines model for apiVersion.
type ApiVersion = string

// RequestId defines model for requestId.
type RequestId = string

// UserAgent defines model for userAgent.
type UserAgent = string

// AgentNotFound Error processing request.
type AgentNotFound = Error

// BadRequest Error processing request.
type BadRequest = Error

// Deadline Error processing request.
type Deadline = Error

// InternalServerError Error processing request.
type InternalServerError = Error

// KeyNotEnabled Error processing request.
type KeyNotEnabled = Error

// Throttle Error processing request.
type Throttle = Error

// Unavailable Error processing request.
type Unavailable = Error

// AgentEnrollParams defines parameters for AgentEnroll.
type AgentEnrollParams struct {
	// UserAgent The user-agent header that is sent.
	// Must have the format "elastic agent X.Y.Z" where "X.Y.Z" indicates the agent version.
	// The agent version must not be greater than the version of the fleet-server.
	UserAgent UserAgent `json:"User-Agent"`

	// XRequestId The request tracking ID for APM.
	XRequestId *RequestId `json:"X-Request-Id,omitempty"`

	// ElasticApiVersion The API version to use, format should be "YYYY-MM-DD"
	ElasticApiVersion *ApiVersion `json:"Elastic-Api-Version,omitempty"`
}

// AgentAcksParams defines parameters for AgentAcks.
type AgentAcksParams struct {
	// XRequestId The request tracking ID for APM.
	XRequestId *RequestId `json:"X-Request-Id,omitempty"`

	// ElasticApiVersion The API version to use, format should be "YYYY-MM-DD"
	ElasticApiVersion *ApiVersion `json:"Elastic-Api-Version,omitempty"`
}

// AgentCheckinParams defines parameters for AgentCheckin.
type AgentCheckinParams struct {
	// AcceptEncoding If the agent is able to accept encoded responses.
	// Used to indicate if GZIP compression may be used by the server.
	// The elastic-agent does not use the accept-encoding header.
	AcceptEncoding *string `json:"Accept-Encoding,omitempty"`

	// UserAgent The user-agent header that is sent.
	// Must have the format "elastic agent X.Y.Z" where "X.Y.Z" indicates the agent version.
	// The agent version must not be greater than the version of the fleet-server.
	UserAgent UserAgent `json:"User-Agent"`

	// XRequestId The request tracking ID for APM.
	XRequestId *RequestId `json:"X-Request-Id,omitempty"`

	// ElasticApiVersion The API version to use, format should be "YYYY-MM-DD"
	ElasticApiVersion *ApiVersion `json:"Elastic-Api-Version,omitempty"`
}

// ArtifactParams defines parameters for Artifact.
type ArtifactParams struct {
	// XRequestId The request tracking ID for APM.
	XRequestId *RequestId `json:"X-Request-Id,omitempty"`

	// ElasticApiVersion The API version to use, format should be "YYYY-MM-DD"
	ElasticApiVersion *ApiVersion `json:"Elastic-Api-Version,omitempty"`
}

// GetFileParams defines parameters for GetFile.
type GetFileParams struct {
	// XRequestId The request tracking ID for APM.
	XRequestId *RequestId `json:"X-Request-Id,omitempty"`

	// ElasticApiVersion The API version to use, format should be "YYYY-MM-DD"
	ElasticApiVersion *ApiVersion `json:"Elastic-Api-Version,omitempty"`
}

// UploadBeginParams defines parameters for UploadBegin.
type UploadBeginParams struct {
	// XRequestId The request tracking ID for APM.
	XRequestId *RequestId `json:"X-Request-Id,omitempty"`

	// ElasticApiVersion The API version to use, format should be "YYYY-MM-DD"
	ElasticApiVersion *ApiVersion `json:"Elastic-Api-Version,omitempty"`
}

// UploadCompleteParams defines parameters for UploadComplete.
type UploadCompleteParams struct {
	// XRequestId The request tracking ID for APM.
	XRequestId *RequestId `json:"X-Request-Id,omitempty"`

	// ElasticApiVersion The API version to use, format should be "YYYY-MM-DD"
	ElasticApiVersion *ApiVersion `json:"Elastic-Api-Version,omitempty"`
}

// UploadChunkParams defines parameters for UploadChunk.
type UploadChunkParams struct {
	// XChunkSHA2 the SHA256 hash of the body contents for this request
	XChunkSHA2 string `json:"X-Chunk-SHA2"`

	// XRequestId The request tracking ID for APM.
	XRequestId *RequestId `json:"X-Request-Id,omitempty"`

	// ElasticApiVersion The API version to use, format should be "YYYY-MM-DD"
	ElasticApiVersion *ApiVersion `json:"Elastic-Api-Version,omitempty"`
}

// StatusParams defines parameters for Status.
type StatusParams struct {
	// XRequestId The request tracking ID for APM.
	XRequestId *RequestId `json:"X-Request-Id,omitempty"`

	// ElasticApiVersion The API version to use, format should be "YYYY-MM-DD"
	ElasticApiVersion *ApiVersion `json:"Elastic-Api-Version,omitempty"`
}

// AgentEnrollJSONRequestBody defines body for AgentEnroll for application/json ContentType.
type AgentEnrollJSONRequestBody = EnrollRequest

// AgentAcksJSONRequestBody defines body for AgentAcks for application/json ContentType.
type AgentAcksJSONRequestBody = AckRequest

// AgentCheckinJSONRequestBody defines body for AgentCheckin for application/json ContentType.
type AgentCheckinJSONRequestBody = CheckinRequest

// UploadBeginJSONRequestBody defines body for UploadBegin for application/json ContentType.
type UploadBeginJSONRequestBody = UploadBeginRequest

// UploadCompleteJSONRequestBody defines body for UploadComplete for application/json ContentType.
type UploadCompleteJSONRequestBody = UploadCompleteRequest

// Getter for additional properties for UploadBeginRequest. Returns the specified
// element and whether it was found
func (a UploadBeginRequest) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for UploadBeginRequest
func (a *UploadBeginRequest) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for UploadBeginRequest to handle AdditionalProperties
func (a *UploadBeginRequest) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["action_id"]; found {
		err = json.Unmarshal(raw, &a.ActionId)
		if err != nil {
			return fmt.Errorf("error reading 'action_id': %w", err)
		}
		delete(object, "action_id")
	}

	if raw, found := object["agent_id"]; found {
		err = json.Unmarshal(raw, &a.AgentId)
		if err != nil {
			return fmt.Errorf("error reading 'agent_id': %w", err)
		}
		delete(object, "agent_id")
	}

	if raw, found := object["file"]; found {
		err = json.Unmarshal(raw, &a.File)
		if err != nil {
			return fmt.Errorf("error reading 'file': %w", err)
		}
		delete(object, "file")
	}

	if raw, found := object["src"]; found {
		err = json.Unmarshal(raw, &a.Src)
		if err != nil {
			return fmt.Errorf("error reading 'src': %w", err)
		}
		delete(object, "src")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for UploadBeginRequest to handle AdditionalProperties
func (a UploadBeginRequest) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	object["action_id"], err = json.Marshal(a.ActionId)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'action_id': %w", err)
	}

	object["agent_id"], err = json.Marshal(a.AgentId)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'agent_id': %w", err)
	}

	object["file"], err = json.Marshal(a.File)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'file': %w", err)
	}

	object["src"], err = json.Marshal(a.Src)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'src': %w", err)
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for UploadBeginRequest_File. Returns the specified
// element and whether it was found
func (a UploadBeginRequest_File) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for UploadBeginRequest_File
func (a *UploadBeginRequest_File) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for UploadBeginRequest_File to handle AdditionalProperties
func (a *UploadBeginRequest_File) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["Compression"]; found {
		err = json.Unmarshal(raw, &a.Compression)
		if err != nil {
			return fmt.Errorf("error reading 'Compression': %w", err)
		}
		delete(object, "Compression")
	}

	if raw, found := object["hash"]; found {
		err = json.Unmarshal(raw, &a.Hash)
		if err != nil {
			return fmt.Errorf("error reading 'hash': %w", err)
		}
		delete(object, "hash")
	}

	if raw, found := object["mime_type"]; found {
		err = json.Unmarshal(raw, &a.MimeType)
		if err != nil {
			return fmt.Errorf("error reading 'mime_type': %w", err)
		}
		delete(object, "mime_type")
	}

	if raw, found := object["name"]; found {
		err = json.Unmarshal(raw, &a.Name)
		if err != nil {
			return fmt.Errorf("error reading 'name': %w", err)
		}
		delete(object, "name")
	}

	if raw, found := object["size"]; found {
		err = json.Unmarshal(raw, &a.Size)
		if err != nil {
			return fmt.Errorf("error reading 'size': %w", err)
		}
		delete(object, "size")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for UploadBeginRequest_File to handle AdditionalProperties
func (a UploadBeginRequest_File) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	if a.Compression != nil {
		object["Compression"], err = json.Marshal(a.Compression)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'Compression': %w", err)
		}
	}

	if a.Hash != nil {
		object["hash"], err = json.Marshal(a.Hash)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'hash': %w", err)
		}
	}

	object["mime_type"], err = json.Marshal(a.MimeType)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'mime_type': %w", err)
	}

	object["name"], err = json.Marshal(a.Name)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'name': %w", err)
	}

	object["size"], err = json.Marshal(a.Size)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'size': %w", err)
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}
