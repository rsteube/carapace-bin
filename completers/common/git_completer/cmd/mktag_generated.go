package cmd

import (
	"github.com/spf13/cobra"
)

var mktagCmd = &cobra.Command{
	Use:   "mktag",
	Short: "Creates a tag object",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {

	rootCmd.AddCommand(mktagCmd)
}
