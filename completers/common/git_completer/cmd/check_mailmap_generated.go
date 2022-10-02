package cmd

import (
	"github.com/spf13/cobra"
)

var check_mailmapCmd = &cobra.Command{
	Use:   "check-mailmap",
	Short: "Show canonical names and email addresses of contacts",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	check_mailmapCmd.Flags().Bool("stdin", false, "also read contacts from stdin")
	rootCmd.AddCommand(check_mailmapCmd)
}
