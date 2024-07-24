package accountpool

import (
	"errors"
	"strings"
	"sync"
)

type Account struct {
	Token string `json:"token"`
}

type IAccounts struct {
	Accounts []*Account `json:"accounts"`
	mx       sync.Mutex
}

func (a *IAccounts) Get() *Account {
	a.mx.Lock()
	defer a.mx.Unlock()
	if len(a.Accounts) == 0 {
		return nil
	}
	account := a.Accounts[0]
	a.Accounts = append(a.Accounts[1:], account)
	return account
}

func (a *IAccounts) GetList() []*Account {
	return a.Accounts
}

func (a *IAccounts) Add(tokens []string) error {
	a.mx.Lock()
	defer a.mx.Unlock()
	if len(tokens) == 0 {
		return errors.New("tokens is empty")
	}
	existingTokens := make(map[string]struct{})
	for _, token := range tokens {
		if _, exists := existingTokens[token]; !exists {
			a.Accounts = AddAccount(a.Accounts, token)
			existingTokens[token] = struct{}{}
		}
	}
	return nil
}

func AddAccount(Secrets []*Account, token string) []*Account {
	if !checkIsAPIKEY(token) {
		return Secrets
	}
	Secrets = append(Secrets, &Account{Token: token})
	return Secrets
}

func checkIsAPIKEY(token string) bool {
	return strings.HasPrefix(token, "gsk_")
}

func NewAccounts(accounts []*Account) *IAccounts {
	return &IAccounts{Accounts: accounts}
}
