// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package greenflashpublicapi

import (
	"context"
	"encoding/json"
	"net/http"
	"slices"
	"time"

	"github.com/stainless-sdks/greenflash-public-api-go/internal/apijson"
	shimjson "github.com/stainless-sdks/greenflash-public-api-go/internal/encoding/json"
	"github.com/stainless-sdks/greenflash-public-api-go/internal/requestconfig"
	"github.com/stainless-sdks/greenflash-public-api-go/option"
	"github.com/stainless-sdks/greenflash-public-api-go/packages/param"
	"github.com/stainless-sdks/greenflash-public-api-go/packages/respjson"
)

// Capture business events
//
// EventService contains methods and other services that help with interacting with
// the Greenflash API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEventService] method instead.
type EventService struct {
	options []option.RequestOption
}

// NewEventService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewEventService(opts ...option.RequestOption) (r EventService) {
	r = EventService{}
	r.options = opts
	return
}

// Track timestamped events representing user or organization actions. Events are
// used to track important business outcomes (signups, conversions, upgrades,
// cancellations, etc.) and integrate them into Greenflash's quality metrics. Each
// event can be optionally linked to a conversation, user, and organization.
func (r *EventService) New(ctx context.Context, body EventNewParams, opts ...option.RequestOption) (res *CreateEventResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "events"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Request payload for creating events.
//
// The properties EventType, ProductID, Value are required.
type CreateEventParams struct {
	// The specific name or category of the event being tracked (e.g., "trial_started",
	// "signup", "feature_usage"). This helps categorize events for analysis and often
	// pairs with "value" to define the outcome.
	EventType string `json:"eventType" api:"required"`
	// The unique identifier of the Greenflash product associated with this event. This
	// links the event to a specific product context.
	ProductID string `json:"productId" api:"required" format:"uuid"`
	// The specific value associated with the event (e.g., "99.00", "5",
	// "premium_plan"). This pairs with "valueType" and "eventType" to define the
	// magnitude or content of the event.
	Value string `json:"value" api:"required"`
	// The unique Greenflash identifier for the conversation. Links the event to a
	// specific chat session in Greenflash.
	ConversationID param.Opt[string] `json:"conversationId,omitzero" format:"uuid"`
	// The ISO 8601 timestamp of when the event actually occurred. Defaults to the
	// current time if not provided. Useful for backdating historical events.
	EventAt param.Opt[time.Time] `json:"eventAt,omitzero" format:"date-time"`
	// Your system's unique identifier for the conversation or thread where this event
	// occurred.
	ExternalConversationID param.Opt[string] `json:"externalConversationId,omitzero"`
	// Your system's unique identifier for the organization associated with this event.
	// Used to map events to your customer accounts.
	ExternalOrganizationID param.Opt[string] `json:"externalOrganizationId,omitzero"`
	// Your system's unique identifier for the user associated with this event. Used to
	// map Greenflash events back to your user records.
	ExternalUserID param.Opt[string] `json:"externalUserId,omitzero"`
	// When true, bypasses sampling and ensures this event is always ingested
	// regardless of sampleRate. Use for critical events that must be captured.
	ForceSample param.Opt[bool] `json:"forceSample,omitzero"`
	// A unique key for idempotency. If you retry a request with the same insertId, it
	// prevents creating a duplicate event record.
	InsertID param.Opt[string] `json:"insertId,omitzero"`
	// The unique Greenflash identifier for the organization. Provide this if you have
	// the Greenflash Organization ID.
	OrganizationID param.Opt[string] `json:"organizationId,omitzero" format:"uuid"`
	// A precise numeric score between -1.0 and 1.0 for direct control over the quality
	// impact. If omitted, it is automatically derived from the "influence" field.
	QualityImpactScore param.Opt[float64] `json:"qualityImpactScore,omitzero"`
	// Controls the percentage of requests that are ingested (0.0 to 1.0). For example,
	// 0.1 means 10% of events will be stored. Defaults to 1.0 (all events ingested).
	// Sampling is deterministic based on event type and organization.
	SampleRate param.Opt[float64] `json:"sampleRate,omitzero"`
	// The unique Greenflash identifier for the user. Provide this if you already have
	// the Greenflash User ID; otherwise, use "externalUserId".
	UserID param.Opt[string] `json:"userId,omitzero" format:"uuid"`
	// A high-level categorization of how this event generally "changed things" or
	// influenced quality (positive, negative, or neutral). Use this for broad
	// classification of outcomes.
	//
	// Any of "positive", "negative", "neutral".
	Influence CreateEventParamsInfluence `json:"influence,omitzero"`
	// A key-value object for storing additional, unstructured context about the event
	// (e.g., { source: "web_app", campaign_id: "123" }). Useful for custom filtering.
	Properties map[string]any `json:"properties,omitzero"`
	// Defines the format of the "value" field (currency, numeric, or text). This
	// ensures the value is interpreted and processed correctly.
	//
	// Any of "currency", "numeric", "text", "boolean".
	ValueType CreateEventParamsValueType `json:"valueType,omitzero"`
	paramObj
}

func (r CreateEventParams) MarshalJSON() (data []byte, err error) {
	type shadow CreateEventParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateEventParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A high-level categorization of how this event generally "changed things" or
// influenced quality (positive, negative, or neutral). Use this for broad
// classification of outcomes.
type CreateEventParamsInfluence string

const (
	CreateEventParamsInfluencePositive CreateEventParamsInfluence = "positive"
	CreateEventParamsInfluenceNegative CreateEventParamsInfluence = "negative"
	CreateEventParamsInfluenceNeutral  CreateEventParamsInfluence = "neutral"
)

// Defines the format of the "value" field (currency, numeric, or text). This
// ensures the value is interpreted and processed correctly.
type CreateEventParamsValueType string

const (
	CreateEventParamsValueTypeCurrency CreateEventParamsValueType = "currency"
	CreateEventParamsValueTypeNumeric  CreateEventParamsValueType = "numeric"
	CreateEventParamsValueTypeText     CreateEventParamsValueType = "text"
	CreateEventParamsValueTypeBoolean  CreateEventParamsValueType = "boolean"
)

// Success response for event creation.
type CreateEventResponse struct {
	// The unique Greenflash ID of the event record that was created.
	EventID string `json:"eventId" api:"required"`
	// Whether the API call was successful.
	Success bool `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EventID     respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreateEventResponse) RawJSON() string { return r.JSON.raw }
func (r *CreateEventResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventNewParams struct {
	// Request payload for creating events.
	CreateEventParams CreateEventParams
	paramObj
}

func (r EventNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateEventParams)
}
func (r *EventNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.CreateEventParams)
}
