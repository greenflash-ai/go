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

	"github.com/greenflash-ai/go/internal/apijson"
	"github.com/greenflash-ai/go/internal/apiquery"
	shimjson "github.com/greenflash-ai/go/internal/encoding/json"
	"github.com/greenflash-ai/go/internal/requestconfig"
	"github.com/greenflash-ai/go/option"
	"github.com/greenflash-ai/go/packages/param"
	"github.com/greenflash-ai/go/packages/respjson"
)

// Manage users
//
// UserService contains methods and other services that help with interacting with
// the Greenflash API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserService] method instead.
type UserService struct {
	options []option.RequestOption
}

// NewUserService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUserService(opts ...option.RequestOption) (r UserService) {
	r = UserService{}
	r.options = opts
	return
}

// Keep track of who's talking to your AI by creating user profiles with contact
// information and custom properties.
//
// Provide an `externalUserId` to identify the user—your ID from your own system.
// Don't worry about whether they already exist; we'll create them if they're new
// or update their profile if they already exist. This makes syncing user data
// effortless.
//
// You can then reference this user in other API calls using the same
// `externalUserId`.
//
// Optionally associate users with an organization by providing an
// `externalOrganizationId`. If the organization doesn't exist yet, we'll create it
// automatically.
func (r *UserService) New(ctx context.Context, body UserNewParams, opts ...option.RequestOption) (res *CreateUserResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "users"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Update specific fields of an existing user profile without changing everything.
//
// The `userId` in the URL path should be your `externalUserId`. Only the fields
// you include in your request will be updated—everything else stays the same.
// Perfect for targeted updates like changing an email address or adding new
// properties.
//
// Prefer a simpler approach? Use `POST /users` instead—it automatically creates or
// updates the user, so you don't need to know if they exist yet.
//
// Optionally associate the user with an organization by providing an
// `externalOrganizationId`. If the organization doesn't exist yet, we'll create it
// automatically.
func (r *UserService) Update(ctx context.Context, userID string, body UserUpdateParams, opts ...option.RequestOption) (res *UpdateUserResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if userID == "" {
		err = errors.New("missing required userId parameter")
		return nil, err
	}
	path := fmt.Sprintf("users/%s", url.PathEscape(userID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Browse through all the users in your workspace. Filter by organization to see
// who belongs to specific teams or companies. Results are paginated for easy
// navigation through large user bases.
func (r *UserService) List(ctx context.Context, query UserListParams, opts ...option.RequestOption) (res *ListUsersResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "users"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Understand how a specific user engages with your AI across all their
// conversations. Track their satisfaction, identify pain points, and spot
// opportunities to improve their experience.
//
// **⚠️ Requires Growth+ plan or higher**
//
// **Two modes available:**
//
//   - **simple mode**: Get aggregate metrics like average sentiment, frustration
//     levels, and conversation quality. Perfect for user dashboards. No rate
//     limiting.
//   - **insights mode** (default): Access detailed patterns, recurring topics, and
//     AI-generated recommendations specific to this user. Rate limited based on your
//     plan's `maxAnalysesPerHour`.
//
// Returns 404 if the user doesn't exist or has no conversations yet.
func (r *UserService) GetUserAnalytics(ctx context.Context, userID string, query UserGetUserAnalyticsParams, opts ...option.RequestOption) (res *GetUserAnalyticsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if userID == "" {
		err = errors.New("missing required userId parameter")
		return nil, err
	}
	path := fmt.Sprintf("users/%s/analytics", userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Request payload for creating a new user profile.
//
// The property ExternalUserID is required.
type CreateUserParams struct {
	// Your unique identifier for the user. Use this same ID in other API calls to
	// reference this user.
	ExternalUserID string `json:"externalUserId" api:"required"`
	// Whether to anonymize the user's personal information. Defaults to false.
	Anonymized param.Opt[bool] `json:"anonymized,omitzero"`
	// The user's email address.
	Email param.Opt[string] `json:"email,omitzero" format:"email"`
	// Your unique identifier for the organization this user belongs to. If provided,
	// the user will be associated with this organization.
	ExternalOrganizationID param.Opt[string] `json:"externalOrganizationId,omitzero"`
	// The user's full name.
	Name param.Opt[string] `json:"name,omitzero"`
	// The Greenflash organization ID that the user belongs to.
	OrganizationID param.Opt[string] `json:"organizationId,omitzero" format:"uuid"`
	// The user's phone number.
	Phone param.Opt[string] `json:"phone,omitzero"`
	// Additional data about the user (e.g., plan type, preferences).
	Properties map[string]any `json:"properties,omitzero"`
	paramObj
}

func (r CreateUserParams) MarshalJSON() (data []byte, err error) {
	type shadow CreateUserParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateUserParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Success response for user creation.
type CreateUserResponse struct {
	// The user profile.
	Participant Participant `json:"participant" api:"required"`
	// Whether the API call was successful.
	Success bool `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Participant respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreateUserResponse) RawJSON() string { return r.JSON.raw }
func (r *CreateUserResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GetUserAnalyticsResponse struct {
	// Distribution of sentiment changes.
	AverageChangeInUserSentiment GetUserAnalyticsResponseAverageChangeInUserSentiment `json:"averageChangeInUserSentiment" api:"required"`
	// Average commercial intent.
	AverageCommercialIntent GetUserAnalyticsResponseAverageCommercialIntent `json:"averageCommercialIntent" api:"required"`
	// Average conversation quality index.
	AverageConversationQualityIndex float64 `json:"averageConversationQualityIndex" api:"required"`
	// Average conversation rating.
	AverageConversationRating float64 `json:"averageConversationRating" api:"required"`
	// Average frustration level.
	AverageFrustration GetUserAnalyticsResponseAverageFrustration `json:"averageFrustration" api:"required"`
	// Average struggle level.
	AverageStruggle GetUserAnalyticsResponseAverageStruggle `json:"averageStruggle" api:"required"`
	// Average sentiment across all conversations.
	AverageUserSentiment GetUserAnalyticsResponseAverageUserSentiment `json:"averageUserSentiment" api:"required"`
	// Structured participant profile summary.
	Summary GetUserAnalyticsResponseSummary `json:"summary" api:"required"`
	// Total number of conversations analyzed.
	TotalConversations float64 `json:"totalConversations" api:"required"`
	// Keywords extracted (insights mode only).
	Keywords []GetUserAnalyticsResponseKeyword `json:"keywords"`
	// Topics discussed (insights mode only).
	Topics []GetUserAnalyticsResponseTopic `json:"topics"`
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
func (r GetUserAnalyticsResponse) RawJSON() string { return r.JSON.raw }
func (r *GetUserAnalyticsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Distribution of sentiment changes.
type GetUserAnalyticsResponseAverageChangeInUserSentiment struct {
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
func (r GetUserAnalyticsResponseAverageChangeInUserSentiment) RawJSON() string { return r.JSON.raw }
func (r *GetUserAnalyticsResponseAverageChangeInUserSentiment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Average commercial intent.
type GetUserAnalyticsResponseAverageCommercialIntent struct {
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
func (r GetUserAnalyticsResponseAverageCommercialIntent) RawJSON() string { return r.JSON.raw }
func (r *GetUserAnalyticsResponseAverageCommercialIntent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Average frustration level.
type GetUserAnalyticsResponseAverageFrustration struct {
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
func (r GetUserAnalyticsResponseAverageFrustration) RawJSON() string { return r.JSON.raw }
func (r *GetUserAnalyticsResponseAverageFrustration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Average struggle level.
type GetUserAnalyticsResponseAverageStruggle struct {
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
func (r GetUserAnalyticsResponseAverageStruggle) RawJSON() string { return r.JSON.raw }
func (r *GetUserAnalyticsResponseAverageStruggle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Average sentiment across all conversations.
type GetUserAnalyticsResponseAverageUserSentiment struct {
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
func (r GetUserAnalyticsResponseAverageUserSentiment) RawJSON() string { return r.JSON.raw }
func (r *GetUserAnalyticsResponseAverageUserSentiment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Structured participant profile summary.
type GetUserAnalyticsResponseSummary struct {
	// Behavioral patterns observed across conversations.
	BehavioralPatterns []GetUserAnalyticsResponseSummaryBehavioralPattern `json:"behavioralPatterns" api:"required"`
	// Engagement profile.
	Engagement GetUserAnalyticsResponseSummaryEngagement `json:"engagement" api:"required"`
	// Transparency about what data drove the analysis.
	Methodology string `json:"methodology" api:"required"`
	// Product-specific observations (when business context is available).
	ProductAlignment GetUserAnalyticsResponseSummaryProductAlignment `json:"productAlignment" api:"required"`
	// Executive summary of the participant.
	ProfileSummary string `json:"profileSummary" api:"required"`
	// Key signals the product owner should know about.
	Signals []GetUserAnalyticsResponseSummarySignal `json:"signals" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BehavioralPatterns respjson.Field
		Engagement         respjson.Field
		Methodology        respjson.Field
		ProductAlignment   respjson.Field
		ProfileSummary     respjson.Field
		Signals            respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GetUserAnalyticsResponseSummary) RawJSON() string { return r.JSON.raw }
func (r *GetUserAnalyticsResponseSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GetUserAnalyticsResponseSummaryBehavioralPattern struct {
	// Specific examples from conversations.
	Evidence string `json:"evidence" api:"required"`
	// How often this pattern appears.
	//
	// Any of "recurring", "occasional", "rare".
	Frequency string `json:"frequency" api:"required"`
	// What the participant consistently does.
	Pattern string `json:"pattern" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Evidence    respjson.Field
		Frequency   respjson.Field
		Pattern     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GetUserAnalyticsResponseSummaryBehavioralPattern) RawJSON() string { return r.JSON.raw }
func (r *GetUserAnalyticsResponseSummaryBehavioralPattern) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Engagement profile.
type GetUserAnalyticsResponseSummaryEngagement struct {
	// Explanation of the engagement assessment.
	Description string `json:"description" api:"required"`
	// Engagement level classification.
	//
	// Any of "power_user", "regular", "casual", "at_risk", "churning".
	Level string `json:"level" api:"required"`
	// Engagement trend direction.
	//
	// Any of "growing", "stable", "declining".
	Trajectory string `json:"trajectory" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		Level       respjson.Field
		Trajectory  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GetUserAnalyticsResponseSummaryEngagement) RawJSON() string { return r.JSON.raw }
func (r *GetUserAnalyticsResponseSummaryEngagement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Product-specific observations (when business context is available).
type GetUserAnalyticsResponseSummaryProductAlignment struct {
	// Where the product isn't serving them.
	Gaps []string `json:"gaps" api:"required"`
	// What's working well for this participant.
	Strengths []string `json:"strengths" api:"required"`
	// How the participant relates to product goals.
	Summary string `json:"summary" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Gaps        respjson.Field
		Strengths   respjson.Field
		Summary     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GetUserAnalyticsResponseSummaryProductAlignment) RawJSON() string { return r.JSON.raw }
func (r *GetUserAnalyticsResponseSummaryProductAlignment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GetUserAnalyticsResponseSummarySignal struct {
	// Evidence-based description.
	Description string `json:"description" api:"required"`
	// Signal priority.
	//
	// Any of "high", "medium", "low".
	Priority string `json:"priority" api:"required"`
	// Short headline.
	Title string `json:"title" api:"required"`
	// Signal category.
	//
	// Any of "opportunity", "risk", "insight".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		Priority    respjson.Field
		Title       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GetUserAnalyticsResponseSummarySignal) RawJSON() string { return r.JSON.raw }
func (r *GetUserAnalyticsResponseSummarySignal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GetUserAnalyticsResponseKeyword struct {
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
func (r GetUserAnalyticsResponseKeyword) RawJSON() string { return r.JSON.raw }
func (r *GetUserAnalyticsResponseKeyword) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GetUserAnalyticsResponseTopic struct {
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
func (r GetUserAnalyticsResponseTopic) RawJSON() string { return r.JSON.raw }
func (r *GetUserAnalyticsResponseTopic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListUsersResponse []Participant

// The user profile.
type Participant struct {
	// The Greenflash participant ID.
	ID string `json:"id" api:"required"`
	// Whether the participant's personal information is anonymized.
	Anonymized bool `json:"anonymized" api:"required"`
	// Your external user ID (matches the externalUserId from the request).
	ExternalID string `json:"externalId" api:"required"`
	// Your external identifier for the user's organization.
	ExternalOrganizationID string `json:"externalOrganizationId" api:"required"`
	// The internal organization ID that the user belongs to.
	OrganizationID string `json:"organizationId" api:"required" format:"uuid"`
	// Additional data about the participant.
	Properties map[string]any `json:"properties" api:"required"`
	// When the participant was first created.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// The participant's email address.
	Email string `json:"email"`
	// The participant's full name.
	Name string `json:"name"`
	// The participant's phone number.
	Phone string `json:"phone"`
	// When the participant was last updated.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		Anonymized             respjson.Field
		ExternalID             respjson.Field
		ExternalOrganizationID respjson.Field
		OrganizationID         respjson.Field
		Properties             respjson.Field
		CreatedAt              respjson.Field
		Email                  respjson.Field
		Name                   respjson.Field
		Phone                  respjson.Field
		UpdatedAt              respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Participant) RawJSON() string { return r.JSON.raw }
func (r *Participant) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request payload for updating an existing user profile.
type UpdateUserParams struct {
	// Whether to anonymize the user's personal information.
	Anonymized param.Opt[bool] `json:"anonymized,omitzero"`
	// The user's email address.
	Email param.Opt[string] `json:"email,omitzero" format:"email"`
	// Your unique identifier for the organization this user belongs to. If provided,
	// the user will be associated with this organization.
	ExternalOrganizationID param.Opt[string] `json:"externalOrganizationId,omitzero"`
	// The user's full name.
	Name param.Opt[string] `json:"name,omitzero"`
	// The Greenflash organization ID that the user belongs to.
	OrganizationID param.Opt[string] `json:"organizationId,omitzero" format:"uuid"`
	// The user's phone number.
	Phone param.Opt[string] `json:"phone,omitzero"`
	// Additional data about the user (e.g., plan type, preferences).
	Properties map[string]any `json:"properties,omitzero"`
	paramObj
}

func (r UpdateUserParams) MarshalJSON() (data []byte, err error) {
	type shadow UpdateUserParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateUserParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Success response for user update.
type UpdateUserResponse struct {
	// The user profile.
	Participant Participant `json:"participant" api:"required"`
	// Whether the API call was successful.
	Success bool `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Participant respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UpdateUserResponse) RawJSON() string { return r.JSON.raw }
func (r *UpdateUserResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserNewParams struct {
	// Request payload for creating a new user profile.
	CreateUserParams CreateUserParams
	paramObj
}

func (r UserNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateUserParams)
}
func (r *UserNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.CreateUserParams)
}

type UserUpdateParams struct {
	// Request payload for updating an existing user profile.
	UpdateUserParams UpdateUserParams
	paramObj
}

func (r UserUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateUserParams)
}
func (r *UserUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.UpdateUserParams)
}

type UserListParams struct {
	// Maximum number of results to return.
	Limit param.Opt[float64] `query:"limit,omitzero" json:"-"`
	// Offset for pagination.
	Offset param.Opt[float64] `query:"offset,omitzero" json:"-"`
	// Filter users by organization ID.
	OrganizationID param.Opt[string] `query:"organizationId,omitzero" format:"uuid" json:"-"`
	// Page number (used to derive offset = (page-1)\*limit).
	Page param.Opt[float64] `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [UserListParams]'s query parameters as `url.Values`.
func (r UserListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type UserGetUserAnalyticsParams struct {
	// Filter analytics by product ID.
	ProductID param.Opt[string] `query:"productId,omitzero" format:"uuid" json:"-"`
	// Filter analytics by version ID.
	VersionID param.Opt[string] `query:"versionId,omitzero" format:"uuid" json:"-"`
	// Analysis mode: "simple" returns only numeric aggregates (no rate limiting),
	// "insights" includes topics, keywords, and recommendations (rate limited per
	// tenant plan).
	//
	// Any of "simple", "insights".
	Mode UserGetUserAnalyticsParamsMode `query:"mode,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [UserGetUserAnalyticsParams]'s query parameters as
// `url.Values`.
func (r UserGetUserAnalyticsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Analysis mode: "simple" returns only numeric aggregates (no rate limiting),
// "insights" includes topics, keywords, and recommendations (rate limited per
// tenant plan).
type UserGetUserAnalyticsParamsMode string

const (
	UserGetUserAnalyticsParamsModeSimple   UserGetUserAnalyticsParamsMode = "simple"
	UserGetUserAnalyticsParamsModeInsights UserGetUserAnalyticsParamsMode = "insights"
)
