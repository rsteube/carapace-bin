package cmd

import (
	"github.com/spf13/cobra"
)

var http_backendCmd = &cobra.Command{
	Use:   "http-backend",
	Short: "Server side implementation of Git over HTTP",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {

	rootCmd.AddCommand(http_backendCmd)
}
