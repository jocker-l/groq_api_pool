package global

import (
	"github.com/jocker-l/groq_api_pool/pkg/accountpool"
	"github.com/jocker-l/groq_api_pool/pkg/proxypool"
)

var (
	Host           string                 // 服务器地址
	Port           string                 // 服务器端口
	ChinesePrompt  string                 // 中文提示
	Authorization  string                 // 授权信息
	AccountPool    *accountpool.IAccounts // API密钥
	Proxy          *proxypool.IProxy      // 代理
	GroqUrl        string
	SupportApikey  string
	ApiKeyPrefix   string
	IsVercel       string
	OpenAuthSecret string
	AuthSecret     string
)
