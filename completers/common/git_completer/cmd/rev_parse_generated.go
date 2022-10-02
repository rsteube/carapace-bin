package cmd

import (
	"github.com/spf13/cobra"
)

var rev_parseCmd = &cobra.Command{
	Use:   "rev-parse",
	Short: "Pick out and massage parameters",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {

	rootCmd.AddCommand(rev_parseCmd)
}
