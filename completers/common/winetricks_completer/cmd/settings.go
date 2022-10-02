package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var settingsCmd = &cobra.Command{
	Use:   "settings",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(settingsCmd).Standalone()

	rootCmd.AddCommand(settingsCmd)
}
