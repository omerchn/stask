package cmd

import (
	"fmt"
	"stask/internal"

	"github.com/spf13/cobra"
)

var compileCmd = &cobra.Command{
	Use:   "compile",
	Short: "Compile a .st file to a JSON file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		json := internal.StToJson(args[0])
		fmt.Printf("json: %v\n", json)
	},
}

func init() {
	rootCmd.AddCommand(compileCmd)
}
