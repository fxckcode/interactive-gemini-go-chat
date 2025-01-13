package ai

import (
	"context"
	"fmt"

	"github.com/fxckcode/interactive-gemini-go-chat/ai/config"
	"github.com/fxckcode/interactive-gemini-go-chat/env"
	"github.com/fxckcode/interactive-gemini-go-chat/utils"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

var (
	ctx = context.Background()
)

func createClient(token string) *genai.Client {
	client, err := genai.NewClient(ctx, option.WithAPIKey(token))

	if err != nil {
		utils.Log.Fatalf("Error creating client: %v", err)
	}

	if client == nil {
		utils.Log.Fatal("Client is nil")
	}

	return client
}

func SearchBasicGemini(content string, modelAI string) string {
	token := env.ViperEnvVariable("AI_API_KEY")

	client := createClient(token)
	model := client.GenerativeModel(modelAI)

	model.SetTemperature(float32(config.TEMPERATURE))
    model.SystemInstruction = genai.NewUserContent(genai.Text(config.PROMPT))

	response, err := model.GenerateContent(ctx, genai.Text(content))

	if err != nil {
		utils.Log.Fatalf("Error generating content: %v", err)
	}

	output := make(chan string)
	go outputResponse(response, output)

	return <-output
}

func outputResponse(resp *genai.GenerateContentResponse, output chan string) {
	if resp != nil && len(resp.Candidates) > 0 {
		firstCandidate := resp.Candidates[0]
		if firstCandidate.Content != nil && len(firstCandidate.Content.Parts) > 0 {
			part := fmt.Sprint(firstCandidate.Content.Parts[0])
			output <- part
		} else {
			output <- "no content in response"
		}
	} else {
		output <- "response is empty"
	}
	close(output)
}



