// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package greenflashpublicapi_test

import (
	"context"
	"os"
	"testing"

	"github.com/greenflash-ai/go"
	"github.com/greenflash-ai/go/internal/testutil"
	"github.com/greenflash-ai/go/option"
)

func TestUsage(t *testing.T) {
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
	createMessageResponse, err := client.Messages.New(context.TODO(), greenflashpublicapi.MessageNewParams{
		CreateMessageParams: greenflashpublicapi.CreateMessageParams{
			ExternalUserID: "user-123",
			Messages:       []greenflashpublicapi.MessageItemParam{{}, {}, {}, {}, {}},
		},
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", createMessageResponse.ConversationID)
}
