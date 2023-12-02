/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

// surveyCmd represents the survey command
var surveyCmd = &cobra.Command{
	Use: "survey",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("survey called")

		var qs = []*survey.Question{
			{
				Name:      "name",
				Prompt:    &survey.Input{Message: "What is your first name?"},
				Validate:  survey.Required,
				Transform: survey.Title,
			},
		}

		answers := struct {
			Name string
		}{}

		err := survey.Ask(qs, &answers)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println("*********** SURVEY RESULTS ***********")
		fmt.Printf("First Name: %s\n", answers.Name)
	},
}

func init() {
	rootCmd.AddCommand(surveyCmd)
}
