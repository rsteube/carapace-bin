package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var envCmd = &cobra.Command{
	Use:   "env",
	Short: "Manages the Environment Variables for your current Project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(envCmd).Standalone()

	rootCmd.AddCommand(envCmd)
}
