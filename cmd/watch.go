/*
Copyright © 2023 Gabriel Jeronimo dev.gjeronimo@gmail.com
*/

package cmd

import (
	"github.com/spf13/cobra"
)

var path, command string

var watchCmd = &cobra.Command{
	Use:   "watch",
	Short: "Watch program changes.",
	Long:  "Watch program changes and execute command to restart program execution.",
	Run: func(cmd *cobra.Command, args []string) {
		//...
	},
}

func init() {
	rootCmd.AddCommand(watchCmd)
	watchCmd.Flags().StringVarP(&path, "file", "f", "", "Define path to the program that will be executed")
	watchCmd.Flags().StringVarP(&command, "command", "c", "", "Command that will be executed if a change occurs")

	watchCmd.MarkFlagRequired("file")
	watchCmd.MarkFlagRequired("command")
}
