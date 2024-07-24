package initialize

import (
	"bufio"
	"github.com/jocker-l/groq_api_pool/global"
	"github.com/jocker-l/groq_api_pool/pkg/accountpool"
	"os"
)

func InitAuth() {
	var Secrets []*accountpool.Account
	// Read accounts.txt and create a list of accounts
	if _, err := os.Stat("session_tokens.txt"); err == nil {
		// Each line is a proxy, put in proxies array
		file, _ := os.Open("session_tokens.txt")
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			// Split by :
			token := scanner.Text()
			if len(token) == 0 {
				continue
			}
			// Append to accounts
			Secrets = accountpool.AddAccount(Secrets, token)
		}
	}

	// 支持从环境变量中设置单独账号
	if os.Getenv("SESSION_TOKEN") != "" {
		Secrets = accountpool.AddAccount(Secrets, os.Getenv("SESSION_TOKEN"))
	}
	global.AccountPool = accountpool.NewAccounts(Secrets)
}
