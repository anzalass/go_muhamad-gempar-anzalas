package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sashabaranov/go-openai"
)

func main() {
	e := echo.New()

	e.POST("/", getRecomendation)

	e.Logger.Fatal(e.Start(":8080"))
}

func GetCompletionFromMessage(client *openai.Client, ctx context.Context, messages []openai.ChatCompletionMessage) (openai.ChatCompletionResponse, error) {

	model := openai.GPT3Dot5Turbo
	resp, err := client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model:    model,
			Messages: messages,
		},
	)

	return resp, err
}

func getRecomendation(c echo.Context) error {
	var apiKey string = "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"

	budget := c.FormValue("budget")
	purpose := c.FormValue("purpose")

	client := openai.NewClient(apiKey)
	ctx := context.Background()

	message := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleUser,
			Content: fmt.Sprintf("Berikan saya rekomendasi laptop dengan harga %s, untuk %s,", budget, purpose),
		},
	}

	resp, err := GetCompletionFromMessage(client, ctx, message)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, resp)
}
