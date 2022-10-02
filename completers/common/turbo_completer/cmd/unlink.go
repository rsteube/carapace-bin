package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var unlinkCmd = &cobra.Command{
	Use:   "unlink",
	Short: "Unlink the current directory from your Vercel organization and disable Remote Caching",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unlinkCmd).Standalone()

	rootCmd.AddCommand(unlinkCmd)
}
