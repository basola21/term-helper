/*
Copyright © 2024 NAME HERE basel21mahmoud@gmail.com
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

type messagesStruct struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type reqBody struct {
	Model    string          `json:"model"`
	Messages []messagesStruct `json:"messages"`
}

var rootCmd = &cobra.Command{
	Use:   "term-helper",
	Short: "This is an application to help you write term commands",
	Long: `This is an application to help you write term commands long description `,
	Run: func(cmd *cobra.Command, args []string) { 
    apiKey := os.Getenv("OPENAI_API_KEY")
		if apiKey == "" {
			fmt.Println("Please set the OPENAI_API_KEY environment variable.")
			return
		}

		body := reqBody{
			Model: "gpt-4o-mini",
			Messages: []messagesStruct{
				{Role: "system", Content: "You are a helpful assistant."},
				{Role: "user", Content: "Write a haiku that explains the concept of recursion."},
			},
		}

		jsonBody, err := json.Marshal(body)
		if err != nil {
			fmt.Println("Error creating JSON request body:", err)
			return
		}

		request, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(jsonBody))
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

		var result map[string]interface{}
		if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
			fmt.Println("Error reading response:", err)
			return
		}

		fmt.Printf("Response: %+v\n", result)

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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.term-helper.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}


