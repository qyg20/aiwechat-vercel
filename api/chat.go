package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/pwh-pwh/aiwechat-vercel/chat"
	"golang.org/x/text/encoding/simplifiedchinese"
)

func Chat(rw http.ResponseWriter, req *http.Request) {
	accessCode := os.Getenv("accessCode")
	code := req.URL.Query().Get("code")
	if code != accessCode {
		fmt.Fprint(rw, "No valid query code provided.")
		return
	}

	msg := req.URL.Query().Get("msg")
	botType := req.URL.Query().Get("botType")
	if msg == "" {
		msg = "用10个字介绍你自己"
	}
	bot := chat.GetChatBot(botType)
	rpn := bot.Chat("admin", msg)
	s, err := simplifiedchinese.GBK.NewEncoder().String(rpn)
	if err != nil {
		fmt.Fprint(rw, "测试失败")
		fmt.Fprint(rw, err.Error())
		return
	}
	fmt.Fprint(rw, s)
}
