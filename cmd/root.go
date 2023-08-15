package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "rikka",
	Short: "Automatic reload for any program.",
	Long:  `"Rikka" is a program designed for detecting changes in folders and files within your application, and it also facilitates automatic restarting.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {}
