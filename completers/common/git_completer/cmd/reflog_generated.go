package cmd

import (
	"github.com/spf13/cobra"
)

var reflogCmd = &cobra.Command{
	Use:   "reflog",
	Short: "Manage reflog information",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {

	rootCmd.AddCommand(reflogCmd)
}
