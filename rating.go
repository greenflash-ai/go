// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package greenflashpublicapi

import (
	"context"
	"encoding/json"
	"net/http"
	"slices"
	"time"

	"github.com/greenflash-ai/go/internal/apijson"
	shimjson "github.com/greenflash-ai/go/internal/encoding/json"
	"github.com/greenflash-ai/go/internal/requestconfig"
	"github.com/greenflash-ai/go/option"
	"github.com/greenflash-ai/go/packages/param"
	"github.com/greenflash-ai/go/packages/respjson"
)

// Capture interactions between users and AI
//
// RatingService contains methods and other services that help with interacting
// with the Greenflash API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRatingService] method instead.
type RatingService struct {
	options []option.RequestOption
}

// NewRatingService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRatingService(opts ...option.RequestOption) (r RatingService) {
	r = RatingService{}
	r.options = opts
	return
}

// Record user feedback and ratings for conversations or individual messages.
//
// Use this endpoint to collect feedback about response quality or overall
// conversation experiences. You can rate either a specific message (using
// `messageId` or `externalMessageId`) or an entire conversation (using
// `conversationId` or `externalConversationId`).
func (r *RatingService) Log(ctx context.Context, body RatingLogParams, opts ...option.RequestOption) (res *LogRatingResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "ratings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Request payload for logging ratings.
//
// The properties ProductID, Rating, RatingMax, RatingMin are required.
type LogRatingParams struct {
	// The Greenflash product ID to rate.
	ProductID string `json:"productId" api:"required" format:"uuid"`
	// The rating value. Must be between ratingMin and ratingMax (inclusive).
	Rating float64 `json:"rating" api:"required"`
	// The maximum possible rating value (e.g., 5 for a 1-5 scale).
	RatingMax float64 `json:"ratingMax" api:"required"`
	// The minimum possible rating value (e.g., 1 for a 1-5 scale).
	RatingMin float64 `json:"ratingMin" api:"required"`
	// The Greenflash conversation ID to rate. Either conversationId,
	// externalConversationId, messageId, or externalMessageId must be provided.
	ConversationID param.Opt[string] `json:"conversationId,omitzero" format:"uuid"`
	// Your external conversation identifier to rate. Either conversationId,
	// externalConversationId, messageId, or externalMessageId must be provided.
	ExternalConversationID param.Opt[string] `json:"externalConversationId,omitzero"`
	// Your external message identifier to rate. Either conversationId,
	// externalConversationId, messageId, or externalMessageId must be provided.
	ExternalMessageID param.Opt[string] `json:"externalMessageId,omitzero"`
	// Optional text feedback accompanying the rating.
	Feedback param.Opt[string] `json:"feedback,omitzero"`
	// The Greenflash message ID to rate. Either conversationId,
	// externalConversationId, messageId, or externalMessageId must be provided.
	MessageID param.Opt[string] `json:"messageId,omitzero"`
	// When the rating was given. Defaults to current time if not provided.
	RatedAt param.Opt[time.Time] `json:"ratedAt,omitzero" format:"date-time"`
	paramObj
}

func (r LogRatingParams) MarshalJSON() (data []byte, err error) {
	type shadow LogRatingParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogRatingParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Success response for rating logging.
type LogRatingResponse struct {
	// Whether the API call was successful.
	Success bool `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogRatingResponse) RawJSON() string { return r.JSON.raw }
func (r *LogRatingResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RatingLogParams struct {
	// Request payload for logging ratings.
	LogRatingParams LogRatingParams
	paramObj
}

func (r RatingLogParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.LogRatingParams)
}
func (r *RatingLogParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.LogRatingParams)
}
