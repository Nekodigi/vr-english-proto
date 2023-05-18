package translate_api

import (
	"context"
	"fmt"
	"net/http"

	"cloud.google.com/go/translate"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/language"
)

type (
	Translate struct {
	}

	TranslateReq struct {
		Text   string `form:"text"`
		LangTo string `form:"langTo"`
	}
)

var client *translate.Client

func init() {

}

func (u *Translate) Handle(e *gin.Engine) {

	e.POST("/translate", func(c *gin.Context) {
		var translateReq TranslateReq
		c.Bind(&translateReq)
		fmt.Println(translateReq.Text)
		res, _ := TranslateText(translateReq.LangTo, translateReq.Text)
		c.JSON(http.StatusAccepted, res)
	})

}

func TranslateText(targetLanguage, text string) (string, error) {
	ctx := context.Background()

	client, err := translate.NewClient(ctx)
	if err != nil {
		//return "", err
	}
	defer client.Close()
	// text := "The Go Gopher is cute"
	//ctx := context.Background()
	//
	lang, err := language.Parse(targetLanguage) //zh
	if err != nil {
		return "", fmt.Errorf("language.Parse: %v", err)
	}
	resp, err := client.Translate(ctx, []string{text}, lang, nil)
	if err != nil {
		return "", fmt.Errorf("Translate: %v", err)
	}
	if len(resp) == 0 {
		return "", fmt.Errorf("Translate returned empty response to text: %s", text)
	}
	return resp[0].Text, nil
}
