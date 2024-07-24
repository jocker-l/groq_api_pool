package initialize

import (
	"github.com/jocker-l/groq_api_pool/global"
	"github.com/joho/godotenv"
	"os"
)

func InitConfig() {
	_ = godotenv.Load(".env")
	global.Host = os.Getenv("SERVER_HOST")
	if global.Host == "" {
		global.Host = "0.0.0.0"
	}
	global.Port = os.Getenv("SERVER_PORT")
	if global.Port == "" {
		global.Port = "8080"
	}
	global.GroqUrl = os.Getenv("BASE_URL")
	if global.GroqUrl == "" {
		global.GroqUrl = "https://api.groq.com/openai"
	}
	global.Authorization = os.Getenv("AUTHORIZATION")
	global.ChinesePrompt = os.Getenv("CHINESE_PROMPT")
	global.SupportApikey = os.Getenv("SUPPORT_APIKEY")
	if global.SupportApikey == "" {
		global.SupportApikey = "true"
	}
	global.IsVercel = os.Getenv("IS_VERCEL")
	global.ApiKeyPrefix = os.Getenv("API_KEY_PREFIX")
	if global.ApiKeyPrefix == "" {
		global.ApiKeyPrefix = "gsk_"
	}
	global.Authorization = os.Getenv("Authorization")
	global.OpenAuthSecret = os.Getenv("OpenAuthSecret")
	global.AuthSecret = os.Getenv("AuthSecret")
	if global.AuthSecret == "" {
		if global.Authorization == "" {
			global.AuthSecret = "root"
		} else {
			global.AuthSecret = global.Authorization
		}
	}
}
