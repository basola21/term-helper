package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/basola21/term-helper/db"
	"github.com/basola21/term-helper/models"
	"github.com/spf13/cobra"
)

const (
	apiURL       = "https://api.groq.com/openai/v1/chat/completions"
	contentType  = "application/json"
	authHeader   = "Authorization"
	defaultModel = "llama3-8b-8192"
)

var prompt string

var rootCmd = &cobra.Command{
	Use:   "term-helper",
	Short: "An application to help you write terminal commands",
	Long:  `Term Helper assists users in generating Linux terminal commands using the Groq API.`,
	Run:   runPrompt,
}

func runPrompt(cmd *cobra.Command, args []string) {
	if prompt == "" && len(args) > 0 {
		prompt = args[0]
	}

	if prompt == "" {
		log.Println("No prompt provided")
		return
	}

	apiKey := getAPIKey()
	if apiKey == "" {
		log.Println("Please set the GROQ_API_KEY environment variable.")
		return
	}

	body := buildRequestBody(prompt)
	response, err := sendRequest(apiKey, body)
	if err != nil {
		log.Printf("Error making API request: %v\n", err)
		return
	}

	printResponse(response)
	err = db.SaveResponse(
		response.ID,
		response.Model,
		response.Choices[0].Message.Content,
		response.Created,
		response.Usage.TotalTokens,
		response.Usage.TotalTime,
	)
	if err != nil {
		log.Printf("Error saving response to database: %v\n", err)
	}
}

func getAPIKey() string {
	return os.Getenv("GROQ_API_KEY")
}

func buildRequestBody(prompt string) models.ReqBody {
	return models.ReqBody{
		Model: defaultModel,
		Messages: []models.Message{
			{
				Role:    "system",
				Content: "You are a helpful assistant that helps the user write Linux terminal commands.",
			},
			{Role: "user", Content: prompt},
		},
	}
}

func sendRequest(apiKey string, body models.ReqBody) (models.GroqResponse, error) {
	var groqResponse models.GroqResponse

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return groqResponse, fmt.Errorf("error creating JSON body: %v", err)
	}

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return groqResponse, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Content-Type", contentType)
	req.Header.Set(authHeader, "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return groqResponse, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&groqResponse); err != nil {
		return groqResponse, fmt.Errorf("error decoding response: %v", err)
	}

	return groqResponse, nil
}

func printResponse(response models.GroqResponse) {
	fmt.Printf("Response from Groq:\nID: %s\nModel: %s\n", response.ID, response.Model)
	for _, choice := range response.Choices {
		fmt.Printf("Message: %s\n", choice.Message.Content)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error executing command: %v", err)
	}
}

func init() {
	db.InitDB() 
	defer db.CloseDB() 
	rootCmd.Flags().StringVarP(&prompt, "prompt", "p", "", "Prompt to process (optional)")
}
