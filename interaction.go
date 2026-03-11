// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package greenflashpublicapi

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/stainless-sdks/greenflash-public-api-go/internal/apijson"
	"github.com/stainless-sdks/greenflash-public-api-go/internal/apiquery"
	"github.com/stainless-sdks/greenflash-public-api-go/internal/requestconfig"
	"github.com/stainless-sdks/greenflash-public-api-go/option"
	"github.com/stainless-sdks/greenflash-public-api-go/packages/param"
	"github.com/stainless-sdks/greenflash-public-api-go/packages/respjson"
)

// Capture interactions between users and AI
//
// InteractionService contains methods and other services that help with
// interacting with the Greenflash API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInteractionService] method instead.
type InteractionService struct {
	options []option.RequestOption
}

// NewInteractionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInteractionService(opts ...option.RequestOption) (r InteractionService) {
	r = InteractionService{}
	r.options = opts
	return
}

// Browse through all conversations in your workspace to understand how users are
// interacting with your AI. Filter by product or version to focus on specific
// areas of your application.
func (r *InteractionService) List(ctx context.Context, query InteractionListParams, opts ...option.RequestOption) (res *ListInteractionsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "interactions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Understand what happened in a specific conversation with AI-powered analysis.
// See sentiment shifts, detect frustration, identify commercial intent, and get
// actionable insights.
//
// **⚠️ Requires Growth+ plan or higher**
//
// **Two modes available:**
//
//   - **simple mode**: Get just the numbers—sentiment scores, frustration levels,
//     and key metrics. Perfect for dashboards and quick checks. No rate limiting.
//   - **insights mode** (default): Dive deeper with detailed keywords, insights, and
//     AI-generated suggestions for improvement. Rate limited based on your plan's
//     `maxAnalysesPerHour`.
//
// Returns 404 if the conversation doesn't exist or hasn't been analyzed yet.
func (r *InteractionService) GetInteractionAnalytics(ctx context.Context, interactionID string, query InteractionGetInteractionAnalyticsParams, opts ...option.RequestOption) (res *GetInteractionAnalyticsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if interactionID == "" {
		err = errors.New("missing required interactionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("interactions/%s/analytics", interactionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type GetInteractionAnalyticsResponse struct {
	// Average sentiment across user messages.
	AverageUserSentiment GetInteractionAnalyticsResponseAverageUserSentiment `json:"averageUserSentiment" api:"required"`
	// How sentiment changed during the interaction.
	ChangeInUserSentiment GetInteractionAnalyticsResponseChangeInUserSentiment `json:"changeInUserSentiment" api:"required"`
	// Commercial intent detected.
	CommercialIntent GetInteractionAnalyticsResponseCommercialIntent `json:"commercialIntent" api:"required"`
	// Quality index score for the interaction.
	ConversationQualityIndex float64 `json:"conversationQualityIndex" api:"required"`
	// Frustration level detected.
	Frustration GetInteractionAnalyticsResponseFrustration `json:"frustration" api:"required"`
	// Number of messages in the interaction.
	MessageCount float64 `json:"messageCount" api:"required"`
	// Most common emotion expressed by user.
	MostCommonUserEmotion string `json:"mostCommonUserEmotion" api:"required"`
	// User rating for this interaction.
	Rating float64 `json:"rating" api:"required"`
	// Struggle level detected.
	Struggle GetInteractionAnalyticsResponseStruggle `json:"struggle" api:"required"`
	// Summary of the interaction.
	Summary string `json:"summary" api:"required"`
	// Main topic discussed.
	Topic string `json:"topic" api:"required"`
	// Keywords extracted (insights mode only).
	Keywords []string `json:"keywords"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AverageUserSentiment     respjson.Field
		ChangeInUserSentiment    respjson.Field
		CommercialIntent         respjson.Field
		ConversationQualityIndex respjson.Field
		Frustration              respjson.Field
		MessageCount             respjson.Field
		MostCommonUserEmotion    respjson.Field
		Rating                   respjson.Field
		Struggle                 respjson.Field
		Summary                  respjson.Field
		Topic                    respjson.Field
		Keywords                 respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GetInteractionAnalyticsResponse) RawJSON() string { return r.JSON.raw }
func (r *GetInteractionAnalyticsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Average sentiment across user messages.
type GetInteractionAnalyticsResponseAverageUserSentiment struct {
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
func (r GetInteractionAnalyticsResponseAverageUserSentiment) RawJSON() string { return r.JSON.raw }
func (r *GetInteractionAnalyticsResponseAverageUserSentiment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How sentiment changed during the interaction.
type GetInteractionAnalyticsResponseChangeInUserSentiment struct {
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
func (r GetInteractionAnalyticsResponseChangeInUserSentiment) RawJSON() string { return r.JSON.raw }
func (r *GetInteractionAnalyticsResponseChangeInUserSentiment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Commercial intent detected.
type GetInteractionAnalyticsResponseCommercialIntent struct {
	PrimarySignal string  `json:"primarySignal" api:"required"`
	Score         float64 `json:"score" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PrimarySignal respjson.Field
		Score         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GetInteractionAnalyticsResponseCommercialIntent) RawJSON() string { return r.JSON.raw }
func (r *GetInteractionAnalyticsResponseCommercialIntent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Frustration level detected.
type GetInteractionAnalyticsResponseFrustration struct {
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
func (r GetInteractionAnalyticsResponseFrustration) RawJSON() string { return r.JSON.raw }
func (r *GetInteractionAnalyticsResponseFrustration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Struggle level detected.
type GetInteractionAnalyticsResponseStruggle struct {
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
func (r GetInteractionAnalyticsResponseStruggle) RawJSON() string { return r.JSON.raw }
func (r *GetInteractionAnalyticsResponseStruggle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListInteractionsResponse []ListInteractionsResponseItem

type ListInteractionsResponseItem struct {
	// The interaction ID.
	ID string `json:"id" api:"required" format:"uuid"`
	// When the interaction was created.
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Your external identifier for the interaction.
	ExternalID string `json:"externalId" api:"required"`
	// Your external identifier for the participant.
	ExternalParticipantID string `json:"externalParticipantId" api:"required"`
	// User feedback text.
	Feedback string `json:"feedback" api:"required"`
	// The AI model used.
	Model string `json:"model" api:"required"`
	// Your external identifier for the organization.
	OrganizationExternalID string `json:"organizationExternalId" api:"required"`
	// The organization ID.
	OrganizationID string `json:"organizationId" api:"required" format:"uuid"`
	// The user who participated in this interaction.
	ParticipantID string `json:"participantId" api:"required" format:"uuid"`
	// The product ID.
	ProductID string `json:"productId" api:"required" format:"uuid"`
	// User rating for this interaction.
	Rating float64 `json:"rating" api:"required"`
	// Maximum rating value.
	RatingMax float64 `json:"ratingMax" api:"required"`
	// Minimum rating value.
	RatingMin float64 `json:"ratingMin" api:"required"`
	// When the interaction was last updated.
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// The version ID.
	VersionID string `json:"versionId" api:"required" format:"uuid"`
	// Custom interaction properties.
	Properties map[string]any `json:"properties"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		CreatedAt              respjson.Field
		ExternalID             respjson.Field
		ExternalParticipantID  respjson.Field
		Feedback               respjson.Field
		Model                  respjson.Field
		OrganizationExternalID respjson.Field
		OrganizationID         respjson.Field
		ParticipantID          respjson.Field
		ProductID              respjson.Field
		Rating                 respjson.Field
		RatingMax              respjson.Field
		RatingMin              respjson.Field
		UpdatedAt              respjson.Field
		VersionID              respjson.Field
		Properties             respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListInteractionsResponseItem) RawJSON() string { return r.JSON.raw }
func (r *ListInteractionsResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InteractionListParams struct {
	// Maximum number of results to return.
	Limit param.Opt[float64] `query:"limit,omitzero" json:"-"`
	// Offset for pagination.
	Offset param.Opt[float64] `query:"offset,omitzero" json:"-"`
	// Page number
	Page param.Opt[float64] `query:"page,omitzero" json:"-"`
	// Filter interactions by product ID.
	ProductID param.Opt[string] `query:"productId,omitzero" format:"uuid" json:"-"`
	// Filter interactions by version ID.
	VersionID param.Opt[string] `query:"versionId,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [InteractionListParams]'s query parameters as `url.Values`.
func (r InteractionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type InteractionGetInteractionAnalyticsParams struct {
	// Analysis mode: "simple" returns only numeric aggregates (no rate limiting),
	// "insights" includes topics, keywords, and recommendations (rate limited per
	// tenant plan).
	//
	// Any of "simple", "insights".
	Mode InteractionGetInteractionAnalyticsParamsMode `query:"mode,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InteractionGetInteractionAnalyticsParams]'s query
// parameters as `url.Values`.
func (r InteractionGetInteractionAnalyticsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Analysis mode: "simple" returns only numeric aggregates (no rate limiting),
// "insights" includes topics, keywords, and recommendations (rate limited per
// tenant plan).
type InteractionGetInteractionAnalyticsParamsMode string

const (
	InteractionGetInteractionAnalyticsParamsModeSimple   InteractionGetInteractionAnalyticsParamsMode = "simple"
	InteractionGetInteractionAnalyticsParamsModeInsights InteractionGetInteractionAnalyticsParamsMode = "insights"
)
