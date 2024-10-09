package models

type GroqResponse struct {
	ID                string   `json:"id"`
	Object            string   `json:"object"`
	Created           float64  `json:"created"`
	Model             string   `json:"model"`
	Choices           []Choice `json:"choices"`
	Usage             Usage    `json:"usage"`
	SystemFingerprint string   `json:"system_fingerprint"`
	XGroq             XGroq    `json:"x_groq"`
}

type Choice struct {
	FinishReason string  `json:"finish_reason"`
	Index        int     `json:"index"`
	Logprobs     *string `json:"logprobs"`
	Message      Message `json:"message"`
}

type Usage struct {
	CompletionTime   float64 `json:"completion_time"`
	CompletionTokens int     `json:"completion_tokens"`
	PromptTime       float64 `json:"prompt_time"`
	PromptTokens     int     `json:"prompt_tokens"`
	QueueTime        float64 `json:"queue_time"`
	TotalTime        float64 `json:"total_time"`
	TotalTokens      int     `json:"total_tokens"`
}

type XGroq struct {
	ID string `json:"id"`
}

type Message struct {
	Content string `json:"content"`
	Role    string `json:"role"`
}

type ReqBody struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}
