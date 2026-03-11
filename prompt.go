// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package greenflashpublicapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/greenflash-public-api-go/internal/apijson"
	"github.com/stainless-sdks/greenflash-public-api-go/internal/apiquery"
	shimjson "github.com/stainless-sdks/greenflash-public-api-go/internal/encoding/json"
	"github.com/stainless-sdks/greenflash-public-api-go/internal/requestconfig"
	"github.com/stainless-sdks/greenflash-public-api-go/option"
	"github.com/stainless-sdks/greenflash-public-api-go/packages/param"
	"github.com/stainless-sdks/greenflash-public-api-go/packages/respjson"
)

// Manage prompts
//
// PromptService contains methods and other services that help with interacting
// with the Greenflash API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPromptService] method instead.
type PromptService struct {
	options []option.RequestOption
}

// NewPromptService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPromptService(opts ...option.RequestOption) (r PromptService) {
	r = PromptService{}
	r.options = opts
	return
}

// Create a new prompt that you can use across your AI applications. Build prompts
// from one or more components, and use handlebars-style variables like
// `{{userName}}` for personalization.
//
// **Safe by Default:** Creating a prompt creates a new version but doesn't
// activate it. Your production prompts stay unchanged until you explicitly
// activate the new version (via the UI or when you reference it in the Messages
// API). This lets you test and prepare new prompts without risk.
//
// **Versioning:** Every prompt is immutable and versioned with fingerprinting, so
// you can safely iterate and track changes over time.
func (r *PromptService) New(ctx context.Context, body PromptNewParams, opts ...option.RequestOption) (res *CreatePromptResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "prompts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Update a prompt with new content or properties. Your production prompt stays
// safe—updates create new versions without affecting what's currently active.
//
// **How it Works:**
//
//   - **Updating components:** Creates a new immutable version with fingerprinting.
//     The new version is created but NOT activated, so you can test before going
//     live.
//   - **Updating only properties (name/description):** Updates the prompt in-place
//     without creating a new version.
//
// **Version Safety:** Old versions always point to their original prompts,
// preserving your message history and allowing you to roll back if needed.
func (r *PromptService) Update(ctx context.Context, id string, body PromptUpdateParams, opts ...option.RequestOption) (res *UpdatePromptResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("prompts/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Browse through all your prompts to see what you're using across your AI
// applications. Returns all prompts by default (both active and inactive
// versions), or filter by product or version to narrow down the results.
//
// **Filtering & Pagination:**
//
//   - Filter by `productId` to see prompts for a specific product
//   - Filter by `versionId` to see a specific version
//   - Choose your pagination style: cursor-based (`limit` + `cursor`) or page-based
//     (`page` + `pageSize`)
//   - Check the `Link` header for easy pagination navigation
//
// **Note:** This returns lightweight data with just component IDs. Use
// `GET /prompts/:id` to fetch the full prompt content.
func (r *PromptService) List(ctx context.Context, query PromptListParams, opts ...option.RequestOption) (res *ListPromptsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "prompts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Archive a prompt you no longer need. Archived prompts are soft-deleted (we set
// an `archived_at` timestamp) so you can still access them for historical data.
//
// **Safety First:**
//
//   - Can't archive a prompt that's currently active. You must activate a different
//     version first.
//   - Historical data is preserved—old conversations continue to reference archived
//     prompts so your message history stays intact.
//   - Archived prompts remain accessible for reporting and analysis.
func (r *PromptService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *DeletePromptResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("prompts/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Retrieve a prompt and optionally personalize it with dynamic variables. Perfect
// for fetching the prompt you want to use right before sending it to your AI.
//
// **Dynamic Variables:** Use handlebars-style placeholders like `{{userName}}` in
// your prompt, then pass query parameters to fill them in.
//
// **Example:** Calling `/prompts/abc-123?userName=Alice&productName=Premium` will
// replace `{{userName}}` with "Alice" and `{{productName}}` with "Premium" in the
// returned prompt.
func (r *PromptService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *GetPromptResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("prompts/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// The property Content is required.
type ComponentInputParam struct {
	// The content of the component.
	Content string `json:"content" api:"required"`
	// The Greenflash component ID.
	ComponentID param.Opt[string] `json:"componentId,omitzero" format:"uuid"`
	// Your external identifier for the component.
	ExternalComponentID param.Opt[string] `json:"externalComponentId,omitzero"`
	// Whether the component content changes dynamically.
	IsDynamic param.Opt[bool] `json:"isDynamic,omitzero"`
	// Component name.
	Name param.Opt[string] `json:"name,omitzero"`
	// Component source: customer, participant, greenflash, or agent.
	//
	// Any of "customer", "participant", "greenflash", "agent".
	Source ComponentInputSource `json:"source,omitzero"`
	// Component type: system, user, tool, guardrail, rag, agent, or a custom type
	// (other).
	//
	// Any of "system", "user", "tool", "guardrail", "rag", "agent", "other".
	Type ComponentInputType `json:"type,omitzero"`
	paramObj
}

func (r ComponentInputParam) MarshalJSON() (data []byte, err error) {
	type shadow ComponentInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ComponentInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Component source: customer, participant, greenflash, or agent.
type ComponentInputSource string

const (
	ComponentInputSourceCustomer    ComponentInputSource = "customer"
	ComponentInputSourceParticipant ComponentInputSource = "participant"
	ComponentInputSourceGreenflash  ComponentInputSource = "greenflash"
	ComponentInputSourceAgent       ComponentInputSource = "agent"
)

// Component type: system, user, tool, guardrail, rag, agent, or a custom type
// (other).
type ComponentInputType string

const (
	ComponentInputTypeSystem    ComponentInputType = "system"
	ComponentInputTypeUser      ComponentInputType = "user"
	ComponentInputTypeTool      ComponentInputType = "tool"
	ComponentInputTypeGuardrail ComponentInputType = "guardrail"
	ComponentInputTypeRag       ComponentInputType = "rag"
	ComponentInputTypeAgent     ComponentInputType = "agent"
	ComponentInputTypeOther     ComponentInputType = "other"
)

// The property Content is required.
type ComponentUpdateParam struct {
	// Updated component content.
	Content string `json:"content" api:"required"`
	// The Greenflash component ID.
	ComponentID param.Opt[string] `json:"componentId,omitzero" format:"uuid"`
	// External component identifier.
	ExternalComponentID param.Opt[string] `json:"externalComponentId,omitzero"`
	// Dynamic flag.
	IsDynamic param.Opt[bool] `json:"isDynamic,omitzero"`
	// Component name.
	Name param.Opt[string] `json:"name,omitzero"`
	// Component source.
	//
	// Any of "customer", "participant", "greenflash", "agent".
	Source ComponentUpdateSource `json:"source,omitzero"`
	// Component type: system, user, tool, guardrail, rag, agent, or a custom type
	// (other).
	//
	// Any of "system", "user", "tool", "guardrail", "rag", "agent", "other".
	Type ComponentUpdateType `json:"type,omitzero"`
	paramObj
}

func (r ComponentUpdateParam) MarshalJSON() (data []byte, err error) {
	type shadow ComponentUpdateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ComponentUpdateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Component source.
type ComponentUpdateSource string

const (
	ComponentUpdateSourceCustomer    ComponentUpdateSource = "customer"
	ComponentUpdateSourceParticipant ComponentUpdateSource = "participant"
	ComponentUpdateSourceGreenflash  ComponentUpdateSource = "greenflash"
	ComponentUpdateSourceAgent       ComponentUpdateSource = "agent"
)

// Component type: system, user, tool, guardrail, rag, agent, or a custom type
// (other).
type ComponentUpdateType string

const (
	ComponentUpdateTypeSystem    ComponentUpdateType = "system"
	ComponentUpdateTypeUser      ComponentUpdateType = "user"
	ComponentUpdateTypeTool      ComponentUpdateType = "tool"
	ComponentUpdateTypeGuardrail ComponentUpdateType = "guardrail"
	ComponentUpdateTypeRag       ComponentUpdateType = "rag"
	ComponentUpdateTypeAgent     ComponentUpdateType = "agent"
	ComponentUpdateTypeOther     ComponentUpdateType = "other"
)

// The properties Components, Name, ProductID, Role are required.
type CreatePromptParams struct {
	// Array of component objects.
	Components []ComponentInputParam `json:"components,omitzero" api:"required"`
	// Prompt name.
	Name string `json:"name" api:"required"`
	// Product this prompt will map to.
	ProductID string `json:"productId" api:"required" format:"uuid"`
	// Role key in the product mapping (e.g. "agent tool").
	Role string `json:"role" api:"required"`
	// Prompt description.
	Description param.Opt[string] `json:"description,omitzero"`
	// Your external identifier for the prompt.
	ExternalPromptID param.Opt[string] `json:"externalPromptId,omitzero"`
	// Prompt source.
	//
	// Any of "customer", "participant", "greenflash", "agent".
	Source CreatePromptParamsSource `json:"source,omitzero"`
	paramObj
}

func (r CreatePromptParams) MarshalJSON() (data []byte, err error) {
	type shadow CreatePromptParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreatePromptParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Prompt source.
type CreatePromptParamsSource string

const (
	CreatePromptParamsSourceCustomer    CreatePromptParamsSource = "customer"
	CreatePromptParamsSourceParticipant CreatePromptParamsSource = "participant"
	CreatePromptParamsSourceGreenflash  CreatePromptParamsSource = "greenflash"
	CreatePromptParamsSourceAgent       CreatePromptParamsSource = "agent"
)

type CreatePromptResponse struct {
	// The IDs of the created prompt components.
	ComponentIDs []string `json:"componentIds" api:"required" format:"uuid"`
	// The created prompt ID.
	PromptID string `json:"promptId" api:"required" format:"uuid"`
	// The created version ID. Version is created but not activated (activation happens
	// via UI or Messages API).
	VersionID string `json:"versionId" api:"required" format:"uuid"`
	// The external prompt ID.
	ExternalPromptID string `json:"externalPromptId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ComponentIDs     respjson.Field
		PromptID         respjson.Field
		VersionID        respjson.Field
		ExternalPromptID respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreatePromptResponse) RawJSON() string { return r.JSON.raw }
func (r *CreatePromptResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeletePromptResponse struct {
	// ISO 8601 timestamp when archived.
	ArchivedAt string `json:"archivedAt" api:"required"`
	// The archived prompt ID.
	PromptID string `json:"promptId" api:"required" format:"uuid"`
	// The external prompt ID.
	ExternalPromptID string `json:"externalPromptId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ArchivedAt       respjson.Field
		PromptID         respjson.Field
		ExternalPromptID respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeletePromptResponse) RawJSON() string { return r.JSON.raw }
func (r *DeletePromptResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GetPromptResponse struct {
	// The prompt with variables interpolated from query parameters.
	ComposedPrompt string `json:"composedPrompt" api:"required"`
	// The full prompt object with components.
	Prompt Prompt `json:"prompt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ComposedPrompt respjson.Field
		Prompt         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GetPromptResponse) RawJSON() string { return r.JSON.raw }
func (r *GetPromptResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListPromptsResponse []SlimPrompt

// The full prompt object with components.
type Prompt struct {
	// The Greenflash prompt ID.
	ID string `json:"id" api:"required" format:"uuid"`
	// ISO 8601 timestamp when archived, or null if active.
	ArchivedAt string `json:"archivedAt" api:"required"`
	// Array of prompt components that make up this prompt.
	Components []PromptComponent `json:"components" api:"required"`
	// ISO 8601 timestamp when created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Prompt description.
	Description string `json:"description" api:"required"`
	// Prompt name.
	Name string `json:"name" api:"required"`
	// The product ID this prompt is associated with.
	ProductID string `json:"productId" api:"required" format:"uuid"`
	// Prompt source.
	Source string `json:"source" api:"required"`
	// ISO 8601 timestamp when last updated.
	UpdatedAt string `json:"updatedAt" api:"required"`
	// Your external identifier for the prompt.
	ExternalPromptID string `json:"externalPromptId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		ArchivedAt       respjson.Field
		Components       respjson.Field
		CreatedAt        respjson.Field
		Description      respjson.Field
		Name             respjson.Field
		ProductID        respjson.Field
		Source           respjson.Field
		UpdatedAt        respjson.Field
		ExternalPromptID respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Prompt) RawJSON() string { return r.JSON.raw }
func (r *Prompt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PromptComponent struct {
	// The Greenflash component ID.
	ID string `json:"id" api:"required" format:"uuid"`
	// The content of the component.
	Content string `json:"content" api:"required"`
	// ISO 8601 timestamp when created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Whether the component content changes dynamically.
	IsDynamic bool `json:"isDynamic" api:"required"`
	// Component name.
	Name string `json:"name" api:"required"`
	// Component source (e.g., customer, participant, greenflash).
	Source string `json:"source" api:"required"`
	// ISO 8601 timestamp when last updated.
	UpdatedAt string `json:"updatedAt" api:"required"`
	// Your external identifier for the component.
	ExternalComponentID string `json:"externalComponentId"`
	// Component type: system, user, tool, guardrail, rag, agent, or a custom type
	// (other).
	//
	// Any of "system", "user", "tool", "guardrail", "rag", "agent", "other".
	Type PromptComponentType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		Content             respjson.Field
		CreatedAt           respjson.Field
		IsDynamic           respjson.Field
		Name                respjson.Field
		Source              respjson.Field
		UpdatedAt           respjson.Field
		ExternalComponentID respjson.Field
		Type                respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PromptComponent) RawJSON() string { return r.JSON.raw }
func (r *PromptComponent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Component type: system, user, tool, guardrail, rag, agent, or a custom type
// (other).
type PromptComponentType string

const (
	PromptComponentTypeSystem    PromptComponentType = "system"
	PromptComponentTypeUser      PromptComponentType = "user"
	PromptComponentTypeTool      PromptComponentType = "tool"
	PromptComponentTypeGuardrail PromptComponentType = "guardrail"
	PromptComponentTypeRag       PromptComponentType = "rag"
	PromptComponentTypeAgent     PromptComponentType = "agent"
	PromptComponentTypeOther     PromptComponentType = "other"
)

type SlimPrompt struct {
	// The Greenflash prompt ID.
	ID string `json:"id" api:"required" format:"uuid"`
	// ISO 8601 timestamp when archived, or null if active.
	ArchivedAt string `json:"archivedAt" api:"required"`
	// Array of prompt component IDs that make up this prompt.
	Components []SlimPromptComponent `json:"components" api:"required"`
	// ISO 8601 timestamp when created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Your external identifier for the prompt.
	ExternalPromptID string `json:"externalPromptId" api:"required"`
	// Prompt name.
	Name string `json:"name" api:"required"`
	// The product ID this prompt is associated with.
	ProductID string `json:"productId" api:"required" format:"uuid"`
	// ISO 8601 timestamp when last updated.
	UpdatedAt string `json:"updatedAt" api:"required"`
	// The version ID this prompt is associated with, or null if the prompt is not part
	// of any version.
	VersionID string `json:"versionId" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		ArchivedAt       respjson.Field
		Components       respjson.Field
		CreatedAt        respjson.Field
		ExternalPromptID respjson.Field
		Name             respjson.Field
		ProductID        respjson.Field
		UpdatedAt        respjson.Field
		VersionID        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SlimPrompt) RawJSON() string { return r.JSON.raw }
func (r *SlimPrompt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SlimPromptComponent struct {
	// The Greenflash component ID.
	ID string `json:"id" api:"required" format:"uuid"`
	// Your external identifier for the component.
	ExternalComponentID string `json:"externalComponentId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		ExternalComponentID respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SlimPromptComponent) RawJSON() string { return r.JSON.raw }
func (r *SlimPromptComponent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UpdatePromptParams struct {
	// Updated prompt description.
	Description param.Opt[string] `json:"description,omitzero"`
	// Updated prompt name.
	Name param.Opt[string] `json:"name,omitzero"`
	// Role key in the product mapping.
	Role param.Opt[string] `json:"role,omitzero"`
	// Updated components (if provided, creates new immutable prompt and version).
	Components []ComponentUpdateParam `json:"components,omitzero"`
	// Prompt source.
	//
	// Any of "customer", "participant", "greenflash", "agent".
	Source UpdatePromptParamsSource `json:"source,omitzero"`
	paramObj
}

func (r UpdatePromptParams) MarshalJSON() (data []byte, err error) {
	type shadow UpdatePromptParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdatePromptParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Prompt source.
type UpdatePromptParamsSource string

const (
	UpdatePromptParamsSourceCustomer    UpdatePromptParamsSource = "customer"
	UpdatePromptParamsSourceParticipant UpdatePromptParamsSource = "participant"
	UpdatePromptParamsSourceGreenflash  UpdatePromptParamsSource = "greenflash"
	UpdatePromptParamsSourceAgent       UpdatePromptParamsSource = "agent"
)

type UpdatePromptResponse struct {
	// The updated prompt ID.
	PromptID string `json:"promptId" api:"required" format:"uuid"`
	// The version ID. Version is created/updated but not activated (activation happens
	// via UI). Null if only prompt metadata was updated without components.
	VersionID string `json:"versionId" api:"required" format:"uuid"`
	// The external prompt ID.
	ExternalPromptID string `json:"externalPromptId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PromptID         respjson.Field
		VersionID        respjson.Field
		ExternalPromptID respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UpdatePromptResponse) RawJSON() string { return r.JSON.raw }
func (r *UpdatePromptResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PromptNewParams struct {
	CreatePromptParams CreatePromptParams
	paramObj
}

func (r PromptNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreatePromptParams)
}
func (r *PromptNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.CreatePromptParams)
}

type PromptUpdateParams struct {
	UpdatePromptParams UpdatePromptParams
	paramObj
}

func (r PromptUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdatePromptParams)
}
func (r *PromptUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.UpdatePromptParams)
}

type PromptListParams struct {
	// Filter to only show prompts that are part of active versions. When true, only
	// returns prompts associated with versions where isActive=true.
	ActiveOnly param.Opt[bool] `query:"activeOnly,omitzero" json:"-"`
	// Include archived prompts.
	IncludeArchived param.Opt[bool] `query:"includeArchived,omitzero" json:"-"`
	// Page size limit (cursor-based pagination). Use either limit/cursor OR
	// page/pageSize, not both.
	Limit param.Opt[float64] `query:"limit,omitzero" json:"-"`
	// Page number (page-based pagination). Use either page/pageSize OR limit/cursor,
	// not both.
	Page param.Opt[float64] `query:"page,omitzero" json:"-"`
	// Number of items per page (page-based pagination). Use either page/pageSize OR
	// limit/cursor, not both.
	PageSize param.Opt[float64] `query:"pageSize,omitzero" json:"-"`
	// Filter prompts by product ID.
	ProductID param.Opt[string] `query:"productId,omitzero" format:"uuid" json:"-"`
	// Filter prompts by specific version ID.
	VersionID param.Opt[string] `query:"versionId,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [PromptListParams]'s query parameters as `url.Values`.
func (r PromptListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
