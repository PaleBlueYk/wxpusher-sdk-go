package wxpusher

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/PaleBlueYk/wxpusher-sdk-go/msg"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// URLBase 接口域名
const URLBase = "http://wxpusher.zjiecode.com"

// URLSendMessage 发送消息
const URLSendMessage = URLBase + "/api/send/message"

// URLMessageStatus 查询发送状态
const URLMessageStatus = URLBase + "/api/send/query/${messageID}"

// URLCreateQrcode 创建参数二维码
const URLCreateQrcode = URLBase + "/api/fun/create/qrcode"

// URLQueryWxUser 查询App的关注用户
const URLQueryWxUser = URLBase + "/api/fun/wxuser"

// SendMessage 发送消息
func SendMessage(message *msg.Message) ([]msg.SendMsgResult, error) {
	msgResults := make([]msg.SendMsgResult, 0)
	// 校验消息内容
	err := message.Check()
	if err != nil {
		return msgResults, err
	}
	// 参数转json
	requestBody, _ := json.Marshal(message)
	resp, err := http.Post(URLSendMessage, "application/json", bytes.NewReader(requestBody))
	if err != nil {
		return msgResults, msg.NewSDKError(err)
	}
	defer func() { _ = resp.Body.Close() }()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return msgResults, msg.NewSDKError(err)
	}
	result := msg.Result{}
	err = json.Unmarshal(respBody, &result)
	if err != nil {
		return msgResults, msg.NewSDKError(err)
	}
	if !result.Success() {
		return msgResults, msg.NewError(result.Code, errors.New(result.Msg))
	}
	// result.Data 转为[]model.SendMsgResult
	byteData, err := json.Marshal(result.Data)
	if err != nil {
		return msgResults, msg.NewSDKError(err)
	}
	err = json.Unmarshal(byteData, &msgResults)
	if err != nil {
		return msgResults, msg.NewSDKError(err)
	}
	return msgResults, nil
}

// QueryMessageStatus 查询消息发送状态
func QueryMessageStatus(messageID int) (msg.Result, error) {
	var result msg.Result
	url := strings.ReplaceAll(URLMessageStatus, "${messageID}", strconv.Itoa(messageID))
	resp, err := http.Get(url)
	if err != nil {
		return result, msg.NewSDKError(err)
	}
	defer func() { _ = resp.Body.Close() }()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, msg.NewSDKError(err)
	}
	err = json.Unmarshal(respBody, &result)
	if err != nil {
		return result, msg.NewSDKError(err)
	}
	if !result.Success() {
		return result, msg.NewError(result.Code, errors.New(result.Msg))
	}
	return result, nil
}

// CreateQrcode 创建参数二维码
func CreateQrcode(qrcode *msg.Qrcode) (msg.CreateQrcodeResult, error) {
	var qrResult msg.CreateQrcodeResult
	err := qrcode.Check()
	if err != nil {
		return qrResult, err
	}
	requestBody, _ := json.Marshal(qrcode)
	resp, err := http.Post(URLCreateQrcode, "application/json", bytes.NewReader(requestBody))
	if err != nil {
		return qrResult, msg.NewSDKError(err)
	}
	defer func() { _ = resp.Body.Close() }()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return qrResult, msg.NewSDKError(err)
	}
	result := msg.Result{}
	err = json.Unmarshal(respBody, &result)
	if err != nil {
		return qrResult, msg.NewSDKError(err)
	}
	if !result.Success() {
		return qrResult, msg.NewError(result.Code, errors.New(result.Msg))
	}
	// result.Data 转为model.CreateQrcodeResult
	byteData, err := json.Marshal(result.Data)
	if err != nil {
		return qrResult, msg.NewSDKError(err)
	}
	err = json.Unmarshal(byteData, &qrResult)
	if err != nil {
		return qrResult, msg.NewSDKError(err)
	}
	return qrResult, nil
}

// QueryWxUser 查询App的关注用户
func QueryWxUser(appToken string, page, pageSize int) (msg.QueryWxUserResult, error) {
	var queryResult msg.QueryWxUserResult
	req, _ := http.NewRequest("GET", URLQueryWxUser, nil)
	q := req.URL.Query()
	q.Add("appToken", appToken)
	q.Add("page", strconv.Itoa(page))
	q.Add("pageSize", strconv.Itoa(pageSize))
	req.URL.RawQuery = q.Encode()
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return queryResult, msg.NewSDKError(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return queryResult, msg.NewSDKError(err)
	}
	defer func() { _ = resp.Body.Close() }()
	result := msg.Result{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return queryResult, msg.NewSDKError(err)
	}
	// result.Data 转为model.QueryWxUserResult
	byteData, err := json.Marshal(result.Data)
	if err != nil {
		return queryResult, msg.NewSDKError(err)
	}
	err = json.Unmarshal(byteData, &queryResult)
	if err != nil {
		return queryResult, msg.NewSDKError(err)
	}
	return queryResult, nil
}
