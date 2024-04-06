package cmd

import (
	"fmt"
	"log"
	"stask/internal"

	"github.com/spf13/cobra"
)

var compileCmd = &cobra.Command{
	Use:   "compile",
	Short: "Compile a tasks markdown file to JSON",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		file := args[0]
		json, err := internal.TasksToJson(file)
		if err != nil {
			log.SetFlags(0)
			log.Print(err)
			return
		}
		fmt.Printf("%v\n", json)
	},
}

func init() {
	rootCmd.AddCommand(compileCmd)
}
