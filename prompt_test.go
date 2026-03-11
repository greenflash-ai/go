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

func TestPromptNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Prompts.New(context.TODO(), greenflashpublicapi.PromptNewParams{
		CreatePromptParams: greenflashpublicapi.CreatePromptParams{
			Components: []greenflashpublicapi.ComponentInputParam{{
				Content:             "You are a helpful assistant for {{productName}}. Greet {{userName}} warmly.",
				ComponentID:         greenflashpublicapi.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ExternalComponentID: greenflashpublicapi.String("externalComponentId"),
				IsDynamic:           greenflashpublicapi.Bool(false),
				Name:                greenflashpublicapi.String("Base Instructions"),
				Source:              greenflashpublicapi.ComponentInputSourceCustomer,
				Type:                greenflashpublicapi.ComponentInputTypeSystem,
			}},
			Name:             "Customer Support Prompt",
			ProductID:        "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			Role:             "x",
			Description:      greenflashpublicapi.String("Standard customer support  prompt"),
			ExternalPromptID: greenflashpublicapi.String("support-v1"),
			Source:           greenflashpublicapi.CreatePromptParamsSourceCustomer,
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

func TestPromptUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Prompts.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		greenflashpublicapi.PromptUpdateParams{
			UpdatePromptParams: greenflashpublicapi.UpdatePromptParams{
				Components: []greenflashpublicapi.ComponentUpdateParam{{
					Content:             "You are a helpful assistant for {{productName}}. Always be polite to {{userName}}.",
					ComponentID:         greenflashpublicapi.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					ExternalComponentID: greenflashpublicapi.String("externalComponentId"),
					IsDynamic:           greenflashpublicapi.Bool(true),
					Name:                greenflashpublicapi.String("Base Instructions V2"),
					Source:              greenflashpublicapi.ComponentUpdateSourceCustomer,
					Type:                greenflashpublicapi.ComponentUpdateTypeSystem,
				}},
				Description: greenflashpublicapi.String("Updated description"),
				Name:        greenflashpublicapi.String("Updated Customer Support Prompt"),
				Role:        greenflashpublicapi.String("role"),
				Source:      greenflashpublicapi.UpdatePromptParamsSourceCustomer,
			},
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

func TestPromptListWithOptionalParams(t *testing.T) {
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
	_, err := client.Prompts.List(context.TODO(), greenflashpublicapi.PromptListParams{
		ActiveOnly:      greenflashpublicapi.Bool(true),
		IncludeArchived: greenflashpublicapi.Bool(true),
		Limit:           greenflashpublicapi.Float(100),
		Page:            greenflashpublicapi.Float(1),
		PageSize:        greenflashpublicapi.Float(1),
		ProductID:       greenflashpublicapi.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		VersionID:       greenflashpublicapi.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
	})
	if err != nil {
		var apierr *greenflashpublicapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPromptDelete(t *testing.T) {
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
	_, err := client.Prompts.Delete(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *greenflashpublicapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPromptGet(t *testing.T) {
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
	_, err := client.Prompts.Get(context.TODO(), "id")
	if err != nil {
		var apierr *greenflashpublicapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
