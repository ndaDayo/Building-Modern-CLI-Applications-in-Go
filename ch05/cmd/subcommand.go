/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// subcommandCmd represents the subcommand command
var subcommandCmd = &cobra.Command{
	Use: "subcommand",
	Run: func(cmd *cobra.Command, args []string) {
		persistentFlag, _ := cmd.Flags().GetBool("persistentFlag")
		fmt.Printf("persistentFlag is set to %v\n", persistentFlag)
	},
}

func init() {
	commandCmd.AddCommand(subcommandCmd)
}
