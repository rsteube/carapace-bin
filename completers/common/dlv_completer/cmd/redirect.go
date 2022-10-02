package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var redirectCmd = &cobra.Command{
	Use:   "redirect",
	Short: "Help about file redirection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redirectCmd).Standalone()
	rootCmd.AddCommand(redirectCmd)
}
