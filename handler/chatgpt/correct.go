package chatgpt

import (
	"fmt"
	"net/http"

	infraFirestore "github.com/Nekodigi/vr-english-proto-backend/infrastructure/firestore"
	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
)

type (
	Correct struct {
		OpenAI *openai.Client
		Fs     *infraFirestore.Firestore
	}

	CorrectReq struct {
		Text string `form:"text"`
	}
)

func (c *Correct) Handle(e *gin.Engine) {
	e.POST("/correct", func(ctx *gin.Context) {
		var correctReq CorrectReq
		ctx.Bind(&correctReq)
		fmt.Println(correctReq.Text)
		resp, err := c.OpenAI.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo, Messages: GetCorrectPrompt(correctReq.Text),
		})
		//TODO CHAT GPT request
		if err != nil {
			fmt.Printf("ChatCompletion error: %v\n", err)
		}
		ctx.JSON(http.StatusAccepted, resp.Choices[0].Message.Content)
	})
}

func GetCorrectPrompt(target string) []openai.ChatCompletionMessage {
	prompt := []openai.ChatCompletionMessage{
		{Role: "user", Content: "Fix mistake of following sentence:Have you eat dinner."},
		{Role: "assistant", Content: "Have you eaten dinner?"},
		{Role: "user", Content: "Fix mistake of following sentence:Do you sushi"},
		{Role: "assistant", Content: "Do you like sushi?"},
		{Role: "user", Content: target},
	}
	return prompt
}
