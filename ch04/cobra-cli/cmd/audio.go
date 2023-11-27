/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

// audioCmd represents the audio command
var audioCmd = &cobra.Command{
	Use:   "audio",
	Short: "A brief description of your command",
	RunE: func(cmd *cobra.Command, args []string) error {
		filename, err := cmd.Flags().GetString("filename")
		if err != nil {
			fmt.Printf("error retrieving filename: %s\n", err.Error())
			return err
		}
		if filename == "" {
			return errors.New("missing filename")
		}
		fmt.Println("uploading audio file", filename)
		return nil
	},
}

func init() {
	audioCmd.Flags().StringP("filename", "f", "", "audio file")
	uploadCmd.AddCommand(audioCmd)
}
