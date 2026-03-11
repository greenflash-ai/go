// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package greenflashpublicapi_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/greenflash-ai/go"
	"github.com/greenflash-ai/go/internal/testutil"
	"github.com/greenflash-ai/go/option"
)

func TestRatingLogWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := greenflashpublicapi.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Ratings.Log(context.TODO(), greenflashpublicapi.RatingLogParams{
		LogRatingParams: greenflashpublicapi.LogRatingParams{
			ProductID:              "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			Rating:                 4,
			RatingMax:              5,
			RatingMin:              1,
			ConversationID:         greenflashpublicapi.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			ExternalConversationID: greenflashpublicapi.String("123e4567-e89b-12d3-a456-426614174000"),
			ExternalMessageID:      greenflashpublicapi.String("externalMessageId"),
			Feedback:               greenflashpublicapi.String("Helpful response!"),
			MessageID:              greenflashpublicapi.String("messageId"),
			RatedAt:                greenflashpublicapi.Time(time.Now()),
		},
	})
	if err != nil {
		var apierr *greenflashpublicapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
