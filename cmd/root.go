/*
Copyright © 2024 NAME HERE basel21mahmoud@gmail.com
*/package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

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

type reqBody struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

var prompt string

var rootCmd = &cobra.Command{
	Use:   "term-helper",
	Short: "This is an application to help you write term commands",
	Long:  `This is an application to help you write term commands long description `,
	Run: func(cmd *cobra.Command, args []string) {
		if prompt == "" && len(args) > 0 {
			prompt = args[0]
		}

		if prompt == "" {
			fmt.Println("No prompt provided")
			return
		}

		apiKey := os.Getenv("GROQ_API_KEY")
		if apiKey == "" {
			fmt.Println("Please set the GROQ_API_KEY environment variable.")
			return
		}

		body := reqBody{
			Model: "llama3-8b-8192",
			Messages: []Message{
				{
					Role:    "system",
					Content: "You are a helpful assistant. that helps the user write linux terminal commands",
				},
				{
          Role: "user", 
          Content: prompt,
        },
			},
		}

		jsonBody, err := json.Marshal(body)
		if err != nil {
			fmt.Println("Error creating JSON request body:", err)
			return
		}

		request, err := http.NewRequest(
			"POST",
			"https://api.groq.com/openai/v1/chat/completions",
			bytes.NewBuffer(jsonBody),
		)
		if err != nil {
			fmt.Println("Error creating request:", err)
			return
		}

		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("Authorization", "Bearer "+apiKey)

		client := &http.Client{}
		response, err := client.Do(request)
		if err != nil {
			fmt.Println("Error making request:", err)
			return
		}

		defer response.Body.Close()

		var groqResponse GroqResponse
		if err := json.NewDecoder(response.Body).Decode(&groqResponse); err != nil {
			fmt.Println("Error reading response:", err)
			return
		}

		fmt.Printf("Response from Groq:\nID: %s\nModel: %s\n", groqResponse.ID, groqResponse.Model)
		for _, choice := range groqResponse.Choices {
			fmt.Printf("Message: %s\n", choice.Message.Content)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().StringVarP(&prompt, "prompt", "p", "", "Prompt to process (optional)")
}
