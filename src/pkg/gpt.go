package lib

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

type Request struct {
	APIKey       string      `json:"-"`
	GPTModel     *string     `json:"gpt_model"`
	SystemPrompt *string     `json:"system_prompt"`
	MaxTokens    *int        `json:"max_tokens"`
	Language     *string     `json:"language"`
	Length       *int        `json:"length"`
	Model        interface{} `json:"model"`
}

type gptRequest struct {
	Model     string       `json:"model"`
	MaxTokens *int         `json:"max_tokens"`
	Messages  []gptMessage `json:"messages"`
}

type gptMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func GetGPTResponse(request Request) (string, error) {
	if len(request.APIKey) == 0 {
		return "", errors.New("API key is required")
	}
	if request.Model == nil {
		return "", errors.New("model is required")
	}
	gptReq, err := makeGPTRequest(request)
	if err != nil {
		return "", err
	}
	gptJson, err := json.Marshal(gptReq)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", strings.NewReader(string(gptJson)))
	if err != nil {
		return "", err
	}
	req.Header["Authorization"] = []string{"Bearer " + request.APIKey}
	req.Header["Content-Type"] = []string{"application/json"}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return "", errors.New("OpenAI API returned status code " + fmt.Sprint(resp.StatusCode))
	}
	var gptResp map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&gptResp)
	if err != nil {
		return "", err
	}
	r := gptResp["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)
	rgx := regexp.MustCompile(`/\\"/g`)
	r = rgx.ReplaceAllString(r, `"`)
	return r, nil
}

func makeGPTRequest(request Request) (*gptRequest, error) {
	gptReq := gptRequest{
		MaxTokens: request.MaxTokens,
	}
	if request.GPTModel != nil && len(*request.GPTModel) > 0 {
		gptReq.Model = *request.GPTModel
	} else {
		gptReq.Model = "gpt-3.5-turbo"
	}
	sysMsg := makeSystemMessage(request)
	gptReq.Messages = append(gptReq.Messages, sysMsg)
	userMsg, err := makeUserMessage(request)
	if err != nil {
		return nil, err
	}
	gptReq.Messages = append(gptReq.Messages, *userMsg)
	return &gptReq, nil
}

func makeSystemMessage(request Request) gptMessage {
	sysMsg := gptMessage{
		Role: "system",
	}
	if request.SystemPrompt != nil && len(*request.SystemPrompt) > 0 {
		sysMsg.Content = *request.SystemPrompt
	} else {
		sysMsg.Content = "You're system to generate placeholder json that's intended to mock an actual api for developers. Users will send you the json they want. The key is actual key they want, and its value is a description for you on the data to genenrate. if the object is in an array, respond with an array of the described object with 10 objects if the user didn't sepcify. If the length is one input is an object, reply with an object otherwise an array. The user may output sepcify language too if not, use English. Your reply should be json only and nothing else"
	}
	return sysMsg
}

func makeUserMessage(request Request) (*gptMessage, error) {
	userMsg := gptMessage{
		Role:    "user",
		Content: "language: ",
	}
	if request.Language != nil && len(*request.Language) > 0 {
		userMsg.Content += *request.Language
	} else {
		userMsg.Content += "english. "
	}
	if request.Length != nil && *request.Length > 0 {
		userMsg.Content += "length: " + fmt.Sprint(*request.Length) + ". "
	}
	requestedJson, err := json.Marshal(request.Model)
	if err != nil {
		return nil, err
	}
	rgx := regexp.MustCompile(`/"/g`)
	userMsg.Content += rgx.ReplaceAllString(string(requestedJson), "\\\"")
	return &userMsg, nil
}
