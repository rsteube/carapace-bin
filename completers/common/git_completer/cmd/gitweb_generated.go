package cmd

import (
	"github.com/spf13/cobra"
)

var gitwebCmd = &cobra.Command{
	Use:   "gitweb",
	Short: "Git web interface (web frontend to Git repositories)",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {

	rootCmd.AddCommand(gitwebCmd)
}
