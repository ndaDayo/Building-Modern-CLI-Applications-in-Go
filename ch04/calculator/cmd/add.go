package cmd

import (
	"calculator/storage"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add number",
	Short: "add value",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 {
			fmt.Println("only accepts a single argument")
			return
		}
		if len(args) == 0 {
			fmt.Println("command requires input value")
			return
		}
		floatVal, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
			fmt.Printf("unable to parse input[%s]: %v",
				args[0], err)
			return
		}
		value := storage.GetValue()
		value += floatVal
		storage.SetValue(value)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
