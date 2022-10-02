package cmd

import (
	"github.com/spf13/cobra"
)

var fast_importCmd = &cobra.Command{
	Use:   "fast-import",
	Short: "Backend for fast Git data importers",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {

	rootCmd.AddCommand(fast_importCmd)
}
