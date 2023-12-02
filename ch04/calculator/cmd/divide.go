/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// divideCmd represents the divide command
var divideCmd = &cobra.Command{
	Use: "divide",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("divide called")
	},
}

func init() {
	rootCmd.AddCommand(divideCmd)

}
