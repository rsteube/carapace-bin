package cmd

import (
	"github.com/spf13/cobra"
)

var updateCheckCmd = &cobra.Command{
	Use:   "update-check",
	Short: "Print current and latest version number",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	rootCmd.AddCommand(updateCheckCmd)
}
