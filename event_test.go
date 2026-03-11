// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package greenflashpublicapi_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stainless-sdks/greenflash-public-api-go"
	"github.com/stainless-sdks/greenflash-public-api-go/internal/testutil"
	"github.com/stainless-sdks/greenflash-public-api-go/option"
)

func TestEventNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Events.New(context.TODO(), greenflashpublicapi.EventNewParams{
		CreateEventParams: greenflashpublicapi.CreateEventParams{
			EventType:              "x",
			ProductID:              "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			Value:                  "value",
			ConversationID:         greenflashpublicapi.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			EventAt:                greenflashpublicapi.Time(time.Now()),
			ExternalConversationID: greenflashpublicapi.String("externalConversationId"),
			ExternalOrganizationID: greenflashpublicapi.String("externalOrganizationId"),
			ExternalUserID:         greenflashpublicapi.String("externalUserId"),
			ForceSample:            greenflashpublicapi.Bool(true),
			Influence:              greenflashpublicapi.CreateEventParamsInfluencePositive,
			InsertID:               greenflashpublicapi.String("insertId"),
			OrganizationID:         greenflashpublicapi.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Properties: map[string]any{
				"foo": "bar",
			},
			QualityImpactScore: greenflashpublicapi.Float(-1),
			SampleRate:         greenflashpublicapi.Float(0),
			UserID:             greenflashpublicapi.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			ValueType:          greenflashpublicapi.CreateEventParamsValueTypeCurrency,
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
