package cmd

import (
	"github.com/jeronimo3875br/rikka/internal/file_handler"
	"github.com/spf13/cobra"
)

var ignore []string
var path, reflect string

var watchCmd = &cobra.Command{
	Use:   "watch",
	Short: "Watch program changes.",
	Long:  "Watch program changes and execute command to restart program execution.",
	Run: func(cmd *cobra.Command, args []string) {
		watchFile := file_handler.WatchFiles{
			path,
			reflect,
			ignore,
		}
		watchFile.Run()
	},
}

func init() {
	rootCmd.AddCommand(watchCmd)
	watchCmd.Flags().StringVarP(&path, "path", "p", "./", "Set path to program folders")
	watchCmd.Flags().StringSliceVarP(&ignore, "ignore", "i", []string{}, "List of folders and files to ignore")
	watchCmd.Flags().StringVarP(&reflect, "reflect", "r", "", "Command that will be executed if a change occurs")

	// Required flags
	watchCmd.MarkFlagRequired("reflect")
}
