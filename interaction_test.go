// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package greenflashpublicapi_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/greenflash-public-api-go"
	"github.com/stainless-sdks/greenflash-public-api-go/internal/testutil"
	"github.com/stainless-sdks/greenflash-public-api-go/option"
)

func TestInteractionListWithOptionalParams(t *testing.T) {
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
	_, err := client.Interactions.List(context.TODO(), greenflashpublicapi.InteractionListParams{
		Limit:     greenflashpublicapi.Float(1),
		Offset:    greenflashpublicapi.Float(0),
		Page:      greenflashpublicapi.Float(1),
		ProductID: greenflashpublicapi.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		VersionID: greenflashpublicapi.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
	})
	if err != nil {
		var apierr *greenflashpublicapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInteractionGetInteractionAnalyticsWithOptionalParams(t *testing.T) {
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
	_, err := client.Interactions.GetInteractionAnalytics(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		greenflashpublicapi.InteractionGetInteractionAnalyticsParams{
			Mode: greenflashpublicapi.InteractionGetInteractionAnalyticsParamsModeSimple,
		},
	)
	if err != nil {
		var apierr *greenflashpublicapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
