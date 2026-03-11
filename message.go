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

// Capture interactions between users and AI
//
// MessageService contains methods and other services that help with interacting
// with the Greenflash API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessageService] method instead.
type MessageService struct {
	options []option.RequestOption
}

// NewMessageService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMessageService(opts ...option.RequestOption) (r MessageService) {
	r = MessageService{}
	r.options = opts
	return
}

// Send us your AI conversations so we can analyze them for you. Works with
// everything from simple chatbots to complex agentic systems.
//
// **Getting Started (Simple Chat):** Just provide the `role` ("user", "assistant",
// or "system") and `content` for each message, along with an
// `externalConversationId` and your `productId`. That's it!
//
// **Advanced Usage (Agentic Workflows):** Capture the full execution trace of your
// AI agents using `messageType` for tool calls, thoughts, observations, and more.
// Include structured data via `input`/`output` fields to track what your agents
// are doing.
//
// **Key Features:**
//
//   - **Automatic Ordering:** Messages are stored with sequential timestamps, or
//     provide your own `createdAt` timestamps for historical data.
//   - **Threading:** Create nested conversations by referencing parent messages
//     using `parentMessageId` or `parentExternalMessageId`.
//   - **Organization Tracking:** Associate users with organizations via
//     `externalOrganizationId`. We'll create the organization automatically if it
//     doesn't exist.
//   - **Automatic De-duplication:** Messages with an `externalMessageId` that
//     already exists in the conversation are automatically skipped. This allows you
//     to safely resend a batch of messages with new messages appended — previously
//     ingested messages will be deduplicated and only new messages will be inserted.
//     Each message in the response includes a `status` field ("created" or
//     "deduplicated") so you know what happened.
//
// Perfect for understanding how your AI is performing in production and
// identifying areas for improvement.
func (r *MessageService) New(ctx context.Context, body MessageNewParams, opts ...option.RequestOption) (res *CreateMessageResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "messages"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Request payload for logging conversations and messages.
//
// The properties ExternalUserID, Messages are required.
type CreateMessageParams struct {
	// Your external user ID that will be mapped to a user in our system.
	ExternalUserID string `json:"externalUserId" api:"required"`
	// Array of conversation messages.
	Messages []MessageItemParam `json:"messages,omitzero" api:"required"`
	// The Greenflash conversation ID. When provided, updates an existing conversation
	// instead of creating a new one. Either conversationId, externalConversationId,
	// productId must be provided.
	ConversationID param.Opt[string] `json:"conversationId,omitzero" format:"uuid"`
	// Your external identifier for the conversation. Either conversationId,
	// externalConversationId, productId must be provided.
	ExternalConversationID param.Opt[string] `json:"externalConversationId,omitzero"`
	// Your unique identifier for the organization this user belongs to. If provided,
	// the user will be associated with this organization.
	ExternalOrganizationID param.Opt[string] `json:"externalOrganizationId,omitzero"`
	// When true, bypasses sampling and ensures this request is always ingested
	// regardless of sampleRate. Use for critical conversations that must be captured.
	ForceSample param.Opt[bool] `json:"forceSample,omitzero"`
	// The AI model used for the conversation.
	Model param.Opt[string] `json:"model,omitzero"`
	// The Greenflash product this conversation belongs to. Either conversationId,
	// externalConversationId, productId must be provided.
	ProductID param.Opt[string] `json:"productId,omitzero" format:"uuid"`
	// Controls the percentage of requests that are ingested (0.0 to 1.0). For example,
	// 0.1 means 10% of requests will be stored. Defaults to 1.0 (all requests
	// ingested). Sampling is deterministic based on conversation ID.
	SampleRate param.Opt[float64] `json:"sampleRate,omitzero"`
	// Additional data about the conversation.
	Properties map[string]any `json:"properties,omitzero"`
	// System prompt for the conversation. Can be a simple string or a prompt object
	// with components.
	SystemPrompt SystemPromptUnionParam `json:"systemPrompt,omitzero"`
	paramObj
}

func (r CreateMessageParams) MarshalJSON() (data []byte, err error) {
	type shadow CreateMessageParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateMessageParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Success response for message logging.
type CreateMessageResponse struct {
	// The ID of the conversation that was created or updated.
	ConversationID string `json:"conversationId" api:"required" format:"uuid"`
	// The messages that were processed.
	Messages []CreateMessageResponseMessage `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success bool `json:"success" api:"required"`
	// The component IDs used internally to track the system prompt components.
	SystemPromptComponentIDs []string `json:"systemPromptComponentIds" api:"required" format:"uuid"`
	// The prompt ID used internally to track the system prompt.
	SystemPromptPromptID string `json:"systemPromptPromptId" api:"required" format:"uuid"`
	// Template variables used or detected for this conversation.
	PromptVariables map[string]string `json:"promptVariables"`
	// Template match info when content was auto-matched against an existing template.
	TemplateMatch CreateMessageResponseTemplateMatch `json:"templateMatch"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConversationID           respjson.Field
		Messages                 respjson.Field
		Success                  respjson.Field
		SystemPromptComponentIDs respjson.Field
		SystemPromptPromptID     respjson.Field
		PromptVariables          respjson.Field
		TemplateMatch            respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreateMessageResponse) RawJSON() string { return r.JSON.raw }
func (r *CreateMessageResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CreateMessageResponseMessage struct {
	// The internal Greenflash message ID.
	MessageID string `json:"messageId" api:"required"`
	// The type of the message that was created.
	MessageType string `json:"messageType" api:"required"`
	// Whether the message was newly created or deduplicated. Messages with an
	// externalMessageId that already exists in the conversation are automatically
	// skipped and returned with status "deduplicated".
	//
	// Any of "created", "deduplicated".
	Status string `json:"status" api:"required"`
	// Your external identifier for the message, if provided.
	ExternalMessageID string `json:"externalMessageId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MessageID         respjson.Field
		MessageType       respjson.Field
		Status            respjson.Field
		ExternalMessageID respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreateMessageResponseMessage) RawJSON() string { return r.JSON.raw }
func (r *CreateMessageResponseMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Template match info when content was auto-matched against an existing template.
type CreateMessageResponseTemplateMatch struct {
	Matched bool `json:"matched" api:"required"`
	// Any of "exact", "high", "medium".
	Confidence        string            `json:"confidence"`
	DetectedVariables map[string]string `json:"detectedVariables"`
	PromptID          string            `json:"promptId" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Matched           respjson.Field
		Confidence        respjson.Field
		DetectedVariables respjson.Field
		PromptID          respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreateMessageResponseTemplateMatch) RawJSON() string { return r.JSON.raw }
func (r *CreateMessageResponseTemplateMatch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageItemParam struct {
	// Additional context (e.g., RAG data) used to generate the message.
	Context param.Opt[string] `json:"context,omitzero"`
	// The message content. Required for language-based analyses.
	Content param.Opt[string] `json:"content,omitzero"`
	// When this message was created. If not provided, messages get sequential
	// timestamps. Use for importing historical data.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Your external identifier for this message. Used to reference the message in
	// other API calls.
	ExternalMessageID param.Opt[string] `json:"externalMessageId,omitzero"`
	// The AI model used for this specific message. Use for multi-agent scenarios where
	// different messages use different models. Overrides the conversation-level model
	// for this message.
	Model param.Opt[string] `json:"model,omitzero"`
	// The external ID of the parent message for threading. Cannot be used with
	// parentMessageId.
	ParentExternalMessageID param.Opt[string] `json:"parentExternalMessageId,omitzero"`
	// The internal ID of the parent message for threading. Cannot be used with
	// parentExternalMessageId.
	ParentMessageID param.Opt[string] `json:"parentMessageId,omitzero" format:"uuid"`
	// Name of the tool being called. Required for tool_call messages.
	ToolName param.Opt[string] `json:"toolName,omitzero"`
	// Structured input data for tool calls, retrievals, or other operations.
	Input map[string]any `json:"input,omitzero"`
	// Detailed message type for agentic workflows. Cannot be used with role. Available
	// types: user_message, assistant_message, system_message, thought, tool_call,
	// observation, final_response, retrieval, memory_read, memory_write, chain_start,
	// chain_end, embedding, tool_error, callback, llm, task, workflow
	//
	// Any of "user_message", "assistant_message", "system_message", "thought",
	// "tool_call", "observation", "final_response", "retrieval", "memory_read",
	// "memory_write", "chain_start", "chain_end", "embedding", "tool_error",
	// "callback", "llm", "task", "workflow".
	MessageType MessageItemMessageType `json:"messageType,omitzero"`
	// Structured output data from tool calls, retrievals, or other operations.
	Output map[string]any `json:"output,omitzero"`
	// Custom message properties.
	Properties map[string]any `json:"properties,omitzero"`
	// Simple message role for basic chat: user, assistant, or system. Cannot be used
	// with messageType.
	//
	// Any of "user", "assistant", "system".
	Role MessageItemRole `json:"role,omitzero"`
	paramObj
}

func (r MessageItemParam) MarshalJSON() (data []byte, err error) {
	type shadow MessageItemParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageItemParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Detailed message type for agentic workflows. Cannot be used with role. Available
// types: user_message, assistant_message, system_message, thought, tool_call,
// observation, final_response, retrieval, memory_read, memory_write, chain_start,
// chain_end, embedding, tool_error, callback, llm, task, workflow
type MessageItemMessageType string

const (
	MessageItemMessageTypeUserMessage      MessageItemMessageType = "user_message"
	MessageItemMessageTypeAssistantMessage MessageItemMessageType = "assistant_message"
	MessageItemMessageTypeSystemMessage    MessageItemMessageType = "system_message"
	MessageItemMessageTypeThought          MessageItemMessageType = "thought"
	MessageItemMessageTypeToolCall         MessageItemMessageType = "tool_call"
	MessageItemMessageTypeObservation      MessageItemMessageType = "observation"
	MessageItemMessageTypeFinalResponse    MessageItemMessageType = "final_response"
	MessageItemMessageTypeRetrieval        MessageItemMessageType = "retrieval"
	MessageItemMessageTypeMemoryRead       MessageItemMessageType = "memory_read"
	MessageItemMessageTypeMemoryWrite      MessageItemMessageType = "memory_write"
	MessageItemMessageTypeChainStart       MessageItemMessageType = "chain_start"
	MessageItemMessageTypeChainEnd         MessageItemMessageType = "chain_end"
	MessageItemMessageTypeEmbedding        MessageItemMessageType = "embedding"
	MessageItemMessageTypeToolError        MessageItemMessageType = "tool_error"
	MessageItemMessageTypeCallback         MessageItemMessageType = "callback"
	MessageItemMessageTypeLlm              MessageItemMessageType = "llm"
	MessageItemMessageTypeTask             MessageItemMessageType = "task"
	MessageItemMessageTypeWorkflow         MessageItemMessageType = "workflow"
)

// Simple message role for basic chat: user, assistant, or system. Cannot be used
// with messageType.
type MessageItemRole string

const (
	MessageItemRoleUser      MessageItemRole = "user"
	MessageItemRoleAssistant MessageItemRole = "assistant"
	MessageItemRoleSystem    MessageItemRole = "system"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SystemPromptUnionParam struct {
	OfString                         param.Opt[string]                    `json:",omitzero,inline"`
	OfSystemPromptSystemPromptObject *SystemPromptSystemPromptObjectParam `json:",omitzero,inline"`
	paramUnion
}

func (u SystemPromptUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfSystemPromptSystemPromptObject)
}
func (u *SystemPromptUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// System prompt as a prompt object. Can reference an existing prompt by ID or
// define new components inline.
type SystemPromptSystemPromptObjectParam struct {
	// Simple string content (shorthand for a single system component). Mutually
	// exclusive with components.
	Content param.Opt[string] `json:"content,omitzero"`
	// Your external identifier for the prompt. Can be used to reference an existing
	// prompt created via system prompt APIs.
	ExternalPromptID param.Opt[string] `json:"externalPromptId,omitzero"`
	// Greenflash's internal prompt ID. Can be used to reference an existing prompt
	// created via system prompt APIs.
	PromptID param.Opt[string] `json:"promptId,omitzero" format:"uuid"`
	// Array of component objects. When provided with promptId/externalPromptId, will
	// upsert the prompt. When omitted with promptId/externalPromptId, will reference
	// an existing prompt.
	Components []ComponentInputParam `json:"components,omitzero"`
	// Template variables for {{placeholder}} interpolation in component content.
	Variables map[string]string `json:"variables,omitzero"`
	paramObj
}

func (r SystemPromptSystemPromptObjectParam) MarshalJSON() (data []byte, err error) {
	type shadow SystemPromptSystemPromptObjectParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SystemPromptSystemPromptObjectParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageNewParams struct {
	// Request payload for logging conversations and messages.
	CreateMessageParams CreateMessageParams
	paramObj
}

func (r MessageNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateMessageParams)
}
func (r *MessageNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.CreateMessageParams)
}
