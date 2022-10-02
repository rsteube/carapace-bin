package cmd

import (
	"github.com/spf13/cobra"
)

var update_server_infoCmd = &cobra.Command{
	Use:   "update-server-info",
	Short: "Update auxiliary info file to help dumb servers",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	update_server_infoCmd.Flags().BoolP("force", "f", false, "update the info files from scratch")
	rootCmd.AddCommand(update_server_infoCmd)
}
