package groq_client

type APIMessage struct {
	Content string `json:"content"`
	Role    string `json:"role"`
}

type APIRequest struct {
	MaxTokens   int          `json:"max_tokens"`
	Messages    []APIMessage `json:"messages"`
	Model       string       `json:"model"`
	Stream      bool         `json:"stream"`
	Temperature float64      `json:"temperature"`
	TopP        float64      `json:"top_p"`
	Stop        any          `json:"stop"`
}
