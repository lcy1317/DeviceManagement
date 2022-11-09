package server

import (
	"bytes"
	"code.corp.bcollie.net/whistle/lcyutils/models"
	"code.corp.bcollie.net/whistle/lcyutils/restapi/operations/qq"
	"github.com/go-openapi/runtime/middleware"
	"io/ioutil"
	"net/http"
)

func postPrivatemsgError(message string, statusCode int64) middleware.Responder {
	response := &models.APIResponse{
		Code:    statusCode,
		Message: message,
	}
	return qq.NewPostSendPrivateMsgInternalServerError().WithPayload(response)
}

func (s *Server) handlePrivateMsg(params qq.PostSendPrivateMsgParams) middleware.Responder {
	targetUrl := s.baseURL + "send_private_msg"
	client := &http.Client{}

	post := "{ \"user_id\":\"" + params.UserID +
		"\",\"message\":\"" + params.Message +
		"\"}"

	var jsonStr = []byte(post)

	req, _ := http.NewRequest("POST", targetUrl, bytes.NewBuffer(jsonStr))
	//req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Content-Type", "application/json") // 发送请求
	resp, err := client.Do(req)
	defer resp.Body.Close()

	if err != nil { // 格式化返回错误
		postPrivatemsgError("请求错误", 400)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		postPrivatemsgError("解析错误", 400)
	}
	return qq.NewPostSendPrivateMsgOK().WithPayload("消息发送成功 " + string(body))
}
