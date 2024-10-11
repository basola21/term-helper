/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/basola21/term-helper/db"
	"github.com/spf13/cobra"
)

// tokensCmd represents the tokens command
var tokensCmd = &cobra.Command{
	Use:   "tokens", // this is the command name called
	Short: "gets the number of tokens that the user used",
	Run:   getTokens,
}

func getTokens(cmd *cobra.Command, args []string) {
	result, err := db.GetUserTokens()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Tokens used: %+v\n", result)
}

func init() {
	rootCmd.AddCommand(tokensCmd)
	}
