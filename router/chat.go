package router

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jocker-l/groq_api_pool/global"
	"github.com/jocker-l/groq_api_pool/middlewares"
	"github.com/jocker-l/groq_api_pool/pkg/custom_http"
	"github.com/jocker-l/groq_api_pool/pkg/groq_client"
	"github.com/jocker-l/groq_api_pool/pkg/net_http"
	"io"
	"strings"
)

var (
	SupportModels = map[string]string{
		"gemma2-9b-it":       "gemma2-9b-it",
		"gemma-7b-it":        "gemma-7b-it",
		"llama3-70b-8192":    "llama3-70b-8192",
		"llama3-8b-8192":     "llama3-8b-8192",
		"mixtral-8x7b-32768": "mixtral-8x7b-32768",
		"gpt-3.5-turbo":      "llama3-70b-8192",
		"gpt-4":              "llama3-70b-8192",
	}
)

func chat(c *gin.Context) {
	var apiReq groq_client.APIRequest
	if err := c.ShouldBindJSON(&apiReq); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	}

	if apiReq.Model != "" {
		if strings.HasPrefix(apiReq.Model, "gpt-3.5") {
			apiReq.Model = "gpt-3.5-turbo"
		}

		if strings.HasPrefix(apiReq.Model, "gpt-4") {
			apiReq.Model = "gpt-4"
		}
	} else {
		// 防呆，给默认值
		apiReq.Model = "llama3-70b-8192"
	}

	// 处理模型映射
	if _, ok := SupportModels[apiReq.Model]; ok {
		apiReq.Model = SupportModels[apiReq.Model]
	}

	// 默认插入中文prompt
	if global.ChinesePrompt == "true" {
		prompt := groq_client.APIMessage{
			Content: "使用中文回答，输出时不要带英文",
			Role:    "system",
		}
		apiReq.Messages = append([]groq_client.APIMessage{prompt}, apiReq.Messages...)
	}
	client := net_http.NewBasicClient()
	proxyUrl := global.Proxy.GetProxyIP()
	if proxyUrl != "" {
		client.SetProxy(proxyUrl)
	}
	apikey, err := processAPIKEY(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	bodyJson, _ := json.Marshal(apiReq)
	header := baseHeader()
	header.Set("Authorization", "Bearer "+apikey)
	response, err := client.Request("POST", global.GroqUrl+"/v1/chat/completions", header, nil, bytes.NewBuffer(bodyJson))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		body, _ := io.ReadAll(response.Body)
		var result map[string]interface{}
		err = json.Unmarshal(body, &result)
		c.JSON(response.StatusCode, result)
		c.Abort()
		return
	}
	readClient := groq_client.NewReadWriter(c.Writer, response)
	if global.IsVercel != "true" {
		readClient.StreamFlushHandler()
		return
	}
	readClient.StreamHandler()
}

func models(c *gin.Context) {
	client := net_http.NewBasicClient()
	proxyUrl := global.Proxy.GetProxyIP()
	if proxyUrl != "" {
		client.SetProxy(proxyUrl)
	}
	apikey, err := processAPIKEY(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	header := baseHeader()
	header.Set("Authorization", "Bearer "+apikey)
	response, err := client.Request("GET", global.GroqUrl+"/v1/models", header, nil, nil)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	c.JSON(response.StatusCode, result)
}

func baseHeader() custom_http.Headers {
	header := make(custom_http.Headers)
	header.Set("Content-Type", "application/json")
	header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	return header
}

func processAPIKEY(authorization string) (string, error) {
	if authorization != "" {
		customToken := strings.Replace(authorization, "Bearer ", "", 1)
		if customToken != "" {
			// 如果支持apikey调用，且以gsk_开头的字符串，说明传递的是apikey
			if global.SupportApikey == "true" && strings.HasPrefix(customToken, global.ApiKeyPrefix) {
				return customToken, nil
			}
		}
	}
	account := global.AccountPool.Get()
	if account != nil {
		return account.Token, errors.New("account not found")
	}
	return "", nil
}

func InitChat(Router *gin.RouterGroup) {
	Router.Use(middlewares.Authorization)
	Router.GET("models", models)
	ChatRouter := Router.Group("chat")
	{
		ChatRouter.OPTIONS("/completions", middlewares.Options)
		ChatRouter.POST("/completions", chat)
	}
}
