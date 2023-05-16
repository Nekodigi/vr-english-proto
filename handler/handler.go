package handler

import (
	"log"
	"net/http"

	"github.com/Nekodigi/vr-english-proto-backend/config"
	"github.com/Nekodigi/vr-english-proto-backend/handler/chatgpt"
	"github.com/Nekodigi/vr-english-proto-backend/handler/translate_api"
	infraFirestore "github.com/Nekodigi/vr-english-proto-backend/infrastructure/firestore"
	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
)

var (
	openaiClient *openai.Client
	fs           *infraFirestore.Firestore
)

func init() {
	var err error
	conf := config.Load()

	openaiClient = openai.NewClient(conf.ChatGPTToken)

	fs, err = infraFirestore.NewFirestore(conf.ProjectId)
	if err != nil {
		log.Fatalf("firestore.New: %+v", err)
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func Router(e *gin.Engine) {
	e.Use(CORSMiddleware())
	(&translate_api.Translate{}).Handle(e)
	(&chatgpt.Correct{OpenAI: openaiClient, Fs: fs}).Handle(e)
	e.GET("/ping", func(ctx *gin.Context) { ctx.String(http.StatusOK, "pong") })
}
