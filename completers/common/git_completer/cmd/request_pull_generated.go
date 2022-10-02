package cmd

import (
	"github.com/spf13/cobra"
)

var request_pullCmd = &cobra.Command{
	Use:   "request-pull",
	Short: "Generates a summary of pending changes",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	request_pullCmd.Flags().BoolS("p", "p", false, "show patch text as well")
	rootCmd.AddCommand(request_pullCmd)
}
