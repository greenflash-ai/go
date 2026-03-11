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

func TestMessageNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Messages.New(context.TODO(), greenflashpublicapi.MessageNewParams{
		CreateMessageParams: greenflashpublicapi.CreateMessageParams{
			ExternalUserID: "user-123",
			Messages: []greenflashpublicapi.MessageItemParam{{
				Content:           greenflashpublicapi.String("Hello!"),
				Context:           greenflashpublicapi.String("context"),
				CreatedAt:         greenflashpublicapi.Time(time.Now()),
				ExternalMessageID: greenflashpublicapi.String("user-msg-1"),
				Input: map[string]any{
					"foo": "bar",
				},
				MessageType: greenflashpublicapi.MessageItemMessageTypeUserMessage,
				Model:       greenflashpublicapi.String("model"),
				Output: map[string]any{
					"foo": "bar",
				},
				ParentExternalMessageID: greenflashpublicapi.String("parentExternalMessageId"),
				ParentMessageID:         greenflashpublicapi.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Properties: map[string]any{
					"foo": "bar",
				},
				Role:     greenflashpublicapi.MessageItemRoleUser,
				ToolName: greenflashpublicapi.String("toolName"),
			}, {
				Content:           greenflashpublicapi.String("Hi there! How can I help you?"),
				Context:           greenflashpublicapi.String("context"),
				CreatedAt:         greenflashpublicapi.Time(time.Now()),
				ExternalMessageID: greenflashpublicapi.String("assistant-msg-1"),
				Input: map[string]any{
					"foo": "bar",
				},
				MessageType: greenflashpublicapi.MessageItemMessageTypeUserMessage,
				Model:       greenflashpublicapi.String("model"),
				Output: map[string]any{
					"foo": "bar",
				},
				ParentExternalMessageID: greenflashpublicapi.String("parentExternalMessageId"),
				ParentMessageID:         greenflashpublicapi.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Properties: map[string]any{
					"foo": "bar",
				},
				Role:     greenflashpublicapi.MessageItemRoleAssistant,
				ToolName: greenflashpublicapi.String("toolName"),
			}, {
				Content:           greenflashpublicapi.String("Calling search tool"),
				Context:           greenflashpublicapi.String("context"),
				CreatedAt:         greenflashpublicapi.Time(time.Now()),
				ExternalMessageID: greenflashpublicapi.String("tool-call-1"),
				Input: map[string]any{
					"query": "bar",
				},
				MessageType: greenflashpublicapi.MessageItemMessageTypeToolCall,
				Model:       greenflashpublicapi.String("model"),
				Output: map[string]any{
					"foo": "bar",
				},
				ParentExternalMessageID: greenflashpublicapi.String("parentExternalMessageId"),
				ParentMessageID:         greenflashpublicapi.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Properties: map[string]any{
					"foo": "bar",
				},
				Role:     greenflashpublicapi.MessageItemRoleUser,
				ToolName: greenflashpublicapi.String("web_search"),
			}, {
				Content:           greenflashpublicapi.String("Search completed"),
				Context:           greenflashpublicapi.String("context"),
				CreatedAt:         greenflashpublicapi.Time(time.Now()),
				ExternalMessageID: greenflashpublicapi.String("tool-result-1"),
				Input: map[string]any{
					"foo": "bar",
				},
				MessageType: greenflashpublicapi.MessageItemMessageTypeObservation,
				Model:       greenflashpublicapi.String("model"),
				Output: map[string]any{
					"results": "bar",
				},
				ParentExternalMessageID: greenflashpublicapi.String("tool-call-1"),
				ParentMessageID:         greenflashpublicapi.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Properties: map[string]any{
					"foo": "bar",
				},
				Role:     greenflashpublicapi.MessageItemRoleUser,
				ToolName: greenflashpublicapi.String("toolName"),
			}, {
				Content:           greenflashpublicapi.String("Based on the search, today will be sunny with a high of 75°F."),
				Context:           greenflashpublicapi.String("context"),
				CreatedAt:         greenflashpublicapi.Time(time.Now()),
				ExternalMessageID: greenflashpublicapi.String("final-1"),
				Input: map[string]any{
					"foo": "bar",
				},
				MessageType: greenflashpublicapi.MessageItemMessageTypeFinalResponse,
				Model:       greenflashpublicapi.String("model"),
				Output: map[string]any{
					"foo": "bar",
				},
				ParentExternalMessageID: greenflashpublicapi.String("parentExternalMessageId"),
				ParentMessageID:         greenflashpublicapi.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Properties: map[string]any{
					"foo": "bar",
				},
				Role:     greenflashpublicapi.MessageItemRoleUser,
				ToolName: greenflashpublicapi.String("toolName"),
			}},
			ConversationID:         greenflashpublicapi.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			ExternalConversationID: greenflashpublicapi.String("conv-456"),
			ExternalOrganizationID: greenflashpublicapi.String("org-789"),
			ForceSample:            greenflashpublicapi.Bool(true),
			Model:                  greenflashpublicapi.String("gpt-greenflash-1"),
			ProductID:              greenflashpublicapi.String("123e4567-e89b-12d3-a456-426614174001"),
			Properties: map[string]any{
				"campaign": "bar",
			},
			SampleRate: greenflashpublicapi.Float(0),
			SystemPrompt: greenflashpublicapi.SystemPromptUnionParam{
				OfSystemPromptSystemPromptObject: &greenflashpublicapi.SystemPromptSystemPromptObjectParam{
					Components: []greenflashpublicapi.ComponentInputParam{{
						Content:             "You are a helpful assistant.",
						ComponentID:         greenflashpublicapi.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
						ExternalComponentID: greenflashpublicapi.String("externalComponentId"),
						IsDynamic:           greenflashpublicapi.Bool(true),
						Name:                greenflashpublicapi.String("name"),
						Source:              greenflashpublicapi.ComponentInputSourceCustomer,
						Type:                greenflashpublicapi.ComponentInputTypeSystem,
					}},
					Content:          greenflashpublicapi.String("x"),
					ExternalPromptID: greenflashpublicapi.String("externalPromptId"),
					PromptID:         greenflashpublicapi.String("123e4567-e89b-12d3-a456-426614174004"),
					Variables: map[string]string{
						"foo": "string",
					},
				},
			},
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
