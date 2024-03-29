/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "app",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	SetupSignalHandler()
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func SetupSignalHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTSTP)
	go func() {
		sig := <-c
		switch sig {
		case os.Interrupt, syscall.SIGINT:
			fmt.Println("\r- Wakeup! Sleep has been interrupted")
			os.Exit(0)
		case os.Interrupt, syscall.SIGTSTP:
			fmt.Println("\r- Wakeup stop sleep")
			os.Exit(0)
		}
	}()
}
