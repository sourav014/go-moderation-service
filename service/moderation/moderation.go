package moderation

import (
	"context"
	"log"

	language "cloud.google.com/go/language/apiv1"
	"cloud.google.com/go/language/apiv1/languagepb"
	"google.golang.org/api/option"
)

func GetContentModerationDetails(content string) (map[string]float32, error) {
	ctx := context.Background()
	languageServiceClient, err := language.NewRESTClient(ctx, option.WithCredentialsFile("your-json-key.json"))
	if err != nil {
		log.Printf("Error creating language service client: %v", err)
		return nil, err
	}

	defer languageServiceClient.Close()

	req := &languagepb.ModerateTextRequest{
		Document: &languagepb.Document{
			Source: &languagepb.Document_Content{
				Content: content,
			},
			Type:     languagepb.Document_PLAIN_TEXT,
			Language: "en",
		},
	}
	resp, err := languageServiceClient.ModerateText(ctx, req)
	if err != nil {
		log.Printf("Error calling ModerateText: %v", err)
		return nil, err
	}

	moderationResults := make(map[string]float32)
	for _, category := range resp.GetModerationCategories() {
		moderationResults[category.GetName()] = category.GetConfidence()
	}

	return moderationResults, nil
}
