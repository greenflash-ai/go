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
	"time"

	"github.com/stainless-sdks/greenflash-public-api-go/internal/apijson"
	"github.com/stainless-sdks/greenflash-public-api-go/internal/apiquery"
	shimjson "github.com/stainless-sdks/greenflash-public-api-go/internal/encoding/json"
	"github.com/stainless-sdks/greenflash-public-api-go/internal/requestconfig"
	"github.com/stainless-sdks/greenflash-public-api-go/option"
	"github.com/stainless-sdks/greenflash-public-api-go/packages/param"
	"github.com/stainless-sdks/greenflash-public-api-go/packages/respjson"
)

// Manage users
//
// OrganizationService contains methods and other services that help with
// interacting with the Greenflash API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrganizationService] method instead.
type OrganizationService struct {
	options []option.RequestOption
}

// NewOrganizationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOrganizationService(opts ...option.RequestOption) (r OrganizationService) {
	r = OrganizationService{}
	r.options = opts
	return
}

// Group your users by company, team, or any organizational structure that makes
// sense for your business.
//
// Provide an `externalOrganizationId` to identify the organization—your ID from
// your own system. Don't worry about whether it already exists; we'll create it if
// it's new or update it if it already exists. This makes syncing organization data
// effortless.
//
// Reference this organization when creating users (via `/users`) or logging
// messages (via `/messages`) using the same `externalOrganizationId`. Perfect for
// B2B products where you need to track which company each user belongs to.
func (r *OrganizationService) New(ctx context.Context, body OrganizationNewParams, opts ...option.RequestOption) (res *CreateOrganizationResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "organizations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Update specific fields of an existing organization without changing everything.
//
// The `organizationId` in the URL path should be your `externalOrganizationId`.
// Only the fields you include in your request will be updated—everything else
// stays the same. Perfect for targeted updates like renaming a company or updating
// properties.
//
// Prefer a simpler approach? Use `POST /organizations` instead—it automatically
// creates or updates the organization, so you don't need to know if it exists yet.
func (r *OrganizationService) Update(ctx context.Context, organizationID string, body OrganizationUpdateParams, opts ...option.RequestOption) (res *UpdateOrganizationResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if organizationID == "" {
		err = errors.New("missing required organizationId parameter")
		return nil, err
	}
	path := fmt.Sprintf("organizations/%s", url.PathEscape(organizationID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Browse through all the organizations (companies, teams, etc.) in your workspace.
// Search for specific organizations or paginate through the full list. Perfect for
// building admin dashboards or organization management interfaces.
//
// The response includes a `Link` header with URLs for next/previous pages, making
// pagination straightforward.
func (r *OrganizationService) List(ctx context.Context, query OrganizationListParams, opts ...option.RequestOption) (res *ListOrganizationsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "organizations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// See how an entire organization (company, team, etc.) engages with your AI across
// all their users and conversations. Spot trends, measure satisfaction, and
// identify opportunities to improve the experience for your biggest customers.
//
// **⚠️ Requires Growth+ plan or higher**
//
// **Two modes available:**
//
//   - **simple mode**: Get organization-wide metrics like average sentiment,
//     frustration levels, commercial intent, and quality scores. Perfect for
//     executive dashboards and health monitoring. No rate limiting.
//   - **insights mode** (default): Dive into detailed patterns, common topics, and
//     AI-generated recommendations for improving this organization's experience.
//     Rate limited based on your plan's `maxAnalysesPerHour`.
//
// If analytics don't exist yet, they'll be generated in real-time from the
// organization's conversations (this may take a few seconds). Returns 404 if the
// organization doesn't exist or has no conversations.
func (r *OrganizationService) GetOrganizationAnalytics(ctx context.Context, organizationID string, query OrganizationGetOrganizationAnalyticsParams, opts ...option.RequestOption) (res *GetOrganizationAnalyticsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if organizationID == "" {
		err = errors.New("missing required organizationId parameter")
		return nil, err
	}
	path := fmt.Sprintf("organizations/%s/analytics", organizationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Request payload for creating a new organization.
//
// The property ExternalOrganizationID is required.
type CreateOrganizationParams struct {
	// Your unique identifier for the organization. Use this same ID in other API calls
	// to reference this organization.
	ExternalOrganizationID string `json:"externalOrganizationId" api:"required"`
	// The organization's name.
	Name param.Opt[string] `json:"name,omitzero"`
	// Custom organization properties.
	Properties map[string]any `json:"properties,omitzero"`
	paramObj
}

func (r CreateOrganizationParams) MarshalJSON() (data []byte, err error) {
	type shadow CreateOrganizationParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateOrganizationParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Success response for organization creation.
type CreateOrganizationResponse struct {
	// The organization that was created or updated.
	Organization TenantOrganization `json:"organization" api:"required"`
	// Whether the API call was successful.
	Success bool `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Organization respjson.Field
		Success      respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreateOrganizationResponse) RawJSON() string { return r.JSON.raw }
func (r *CreateOrganizationResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GetOrganizationAnalyticsResponse struct {
	// Distribution of sentiment changes.
	AverageChangeInUserSentiment GetOrganizationAnalyticsResponseAverageChangeInUserSentiment `json:"averageChangeInUserSentiment" api:"required"`
	// Average commercial intent.
	AverageCommercialIntent GetOrganizationAnalyticsResponseAverageCommercialIntent `json:"averageCommercialIntent" api:"required"`
	// Average conversation quality index.
	AverageConversationQualityIndex float64 `json:"averageConversationQualityIndex" api:"required"`
	// Average conversation rating.
	AverageConversationRating float64 `json:"averageConversationRating" api:"required"`
	// Average frustration level.
	AverageFrustration GetOrganizationAnalyticsResponseAverageFrustration `json:"averageFrustration" api:"required"`
	// Average struggle level.
	AverageStruggle GetOrganizationAnalyticsResponseAverageStruggle `json:"averageStruggle" api:"required"`
	// Average sentiment across all conversations.
	AverageUserSentiment GetOrganizationAnalyticsResponseAverageUserSentiment `json:"averageUserSentiment" api:"required"`
	// Summary of the organization analytics.
	Summary GetOrganizationAnalyticsResponseSummary `json:"summary" api:"required"`
	// Total number of conversations analyzed.
	TotalConversations float64 `json:"totalConversations" api:"required"`
	// Keywords extracted (insights mode only).
	Keywords []GetOrganizationAnalyticsResponseKeyword `json:"keywords"`
	// Topics discussed (insights mode only).
	Topics []GetOrganizationAnalyticsResponseTopic `json:"topics"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AverageChangeInUserSentiment    respjson.Field
		AverageCommercialIntent         respjson.Field
		AverageConversationQualityIndex respjson.Field
		AverageConversationRating       respjson.Field
		AverageFrustration              respjson.Field
		AverageStruggle                 respjson.Field
		AverageUserSentiment            respjson.Field
		Summary                         respjson.Field
		TotalConversations              respjson.Field
		Keywords                        respjson.Field
		Topics                          respjson.Field
		ExtraFields                     map[string]respjson.Field
		raw                             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GetOrganizationAnalyticsResponse) RawJSON() string { return r.JSON.raw }
func (r *GetOrganizationAnalyticsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Distribution of sentiment changes.
type GetOrganizationAnalyticsResponseAverageChangeInUserSentiment struct {
	Label string  `json:"label" api:"required"`
	Score float64 `json:"score" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Label       respjson.Field
		Score       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GetOrganizationAnalyticsResponseAverageChangeInUserSentiment) RawJSON() string {
	return r.JSON.raw
}
func (r *GetOrganizationAnalyticsResponseAverageChangeInUserSentiment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Average commercial intent.
type GetOrganizationAnalyticsResponseAverageCommercialIntent struct {
	Label string  `json:"label" api:"required"`
	Score float64 `json:"score" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Label       respjson.Field
		Score       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GetOrganizationAnalyticsResponseAverageCommercialIntent) RawJSON() string { return r.JSON.raw }
func (r *GetOrganizationAnalyticsResponseAverageCommercialIntent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Average frustration level.
type GetOrganizationAnalyticsResponseAverageFrustration struct {
	Label string  `json:"label" api:"required"`
	Score float64 `json:"score" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Label       respjson.Field
		Score       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GetOrganizationAnalyticsResponseAverageFrustration) RawJSON() string { return r.JSON.raw }
func (r *GetOrganizationAnalyticsResponseAverageFrustration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Average struggle level.
type GetOrganizationAnalyticsResponseAverageStruggle struct {
	Label string  `json:"label" api:"required"`
	Score float64 `json:"score" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Label       respjson.Field
		Score       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GetOrganizationAnalyticsResponseAverageStruggle) RawJSON() string { return r.JSON.raw }
func (r *GetOrganizationAnalyticsResponseAverageStruggle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Average sentiment across all conversations.
type GetOrganizationAnalyticsResponseAverageUserSentiment struct {
	Label string  `json:"label" api:"required"`
	Score float64 `json:"score" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Label       respjson.Field
		Score       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GetOrganizationAnalyticsResponseAverageUserSentiment) RawJSON() string { return r.JSON.raw }
func (r *GetOrganizationAnalyticsResponseAverageUserSentiment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Summary of the organization analytics.
type GetOrganizationAnalyticsResponseSummary struct {
	Analysis string `json:"analysis" api:"required"`
	Reason   string `json:"reason" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Analysis    respjson.Field
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GetOrganizationAnalyticsResponseSummary) RawJSON() string { return r.JSON.raw }
func (r *GetOrganizationAnalyticsResponseSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GetOrganizationAnalyticsResponseKeyword struct {
	Count float64 `json:"count" api:"required"`
	Name  string  `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GetOrganizationAnalyticsResponseKeyword) RawJSON() string { return r.JSON.raw }
func (r *GetOrganizationAnalyticsResponseKeyword) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GetOrganizationAnalyticsResponseTopic struct {
	Count float64 `json:"count" api:"required"`
	Name  string  `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GetOrganizationAnalyticsResponseTopic) RawJSON() string { return r.JSON.raw }
func (r *GetOrganizationAnalyticsResponseTopic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListOrganizationsResponse []TenantOrganization

// The organization that was created or updated.
type TenantOrganization struct {
	// The Greenflash organization ID.
	ID string `json:"id" api:"required"`
	// Custom organization properties.
	Properties map[string]any `json:"properties" api:"required"`
	// When the organization was first created.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Your external organization ID.
	ExternalID string `json:"externalId"`
	// The organization name.
	Name string `json:"name"`
	// When the organization was last updated.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Properties  respjson.Field
		CreatedAt   respjson.Field
		ExternalID  respjson.Field
		Name        respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TenantOrganization) RawJSON() string { return r.JSON.raw }
func (r *TenantOrganization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request payload for updating an organization.
type UpdateOrganizationParams struct {
	// The organization's name.
	Name param.Opt[string] `json:"name,omitzero"`
	// Custom organization properties.
	Properties map[string]any `json:"properties,omitzero"`
	paramObj
}

func (r UpdateOrganizationParams) MarshalJSON() (data []byte, err error) {
	type shadow UpdateOrganizationParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateOrganizationParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Success response for organization update.
type UpdateOrganizationResponse struct {
	// The organization that was created or updated.
	Organization TenantOrganization `json:"organization" api:"required"`
	// Whether the API call was successful.
	Success bool `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Organization respjson.Field
		Success      respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UpdateOrganizationResponse) RawJSON() string { return r.JSON.raw }
func (r *UpdateOrganizationResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationNewParams struct {
	// Request payload for creating a new organization.
	CreateOrganizationParams CreateOrganizationParams
	paramObj
}

func (r OrganizationNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateOrganizationParams)
}
func (r *OrganizationNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.CreateOrganizationParams)
}

type OrganizationUpdateParams struct {
	// Request payload for updating an organization.
	UpdateOrganizationParams UpdateOrganizationParams
	paramObj
}

func (r OrganizationUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateOrganizationParams)
}
func (r *OrganizationUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.UpdateOrganizationParams)
}

type OrganizationListParams struct {
	// Maximum number of results to return.
	Limit param.Opt[float64] `query:"limit,omitzero" json:"-"`
	// Offset for pagination.
	Offset param.Opt[float64] `query:"offset,omitzero" json:"-"`
	// Page number (used to derive offset = (page-1)\*limit).
	Page param.Opt[float64] `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OrganizationListParams]'s query parameters as `url.Values`.
func (r OrganizationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OrganizationGetOrganizationAnalyticsParams struct {
	// Filter analytics by product ID.
	ProductID param.Opt[string] `query:"productId,omitzero" format:"uuid" json:"-"`
	// Filter analytics by version ID.
	VersionID param.Opt[string] `query:"versionId,omitzero" format:"uuid" json:"-"`
	// Analysis mode: "simple" returns only numeric aggregates (no rate limiting),
	// "insights" includes topics, keywords, and recommendations (rate limited per
	// tenant plan).
	//
	// Any of "simple", "insights".
	Mode OrganizationGetOrganizationAnalyticsParamsMode `query:"mode,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OrganizationGetOrganizationAnalyticsParams]'s query
// parameters as `url.Values`.
func (r OrganizationGetOrganizationAnalyticsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Analysis mode: "simple" returns only numeric aggregates (no rate limiting),
// "insights" includes topics, keywords, and recommendations (rate limited per
// tenant plan).
type OrganizationGetOrganizationAnalyticsParamsMode string

const (
	OrganizationGetOrganizationAnalyticsParamsModeSimple   OrganizationGetOrganizationAnalyticsParamsMode = "simple"
	OrganizationGetOrganizationAnalyticsParamsModeInsights OrganizationGetOrganizationAnalyticsParamsMode = "insights"
)
