/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// sleepCmd represents the sleep command
var sleepCmd = &cobra.Command{
	Use: "sleep",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sleep called")
		for {
			fmt.Println("zzzzz......")
			time.Sleep(time.Second)
		}
	},
}

func init() {
	rootCmd.AddCommand(sleepCmd)
}
