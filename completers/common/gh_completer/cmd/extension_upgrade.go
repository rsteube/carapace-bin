package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var extension_upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade installed extensions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(extension_upgradeCmd).Standalone()
	extension_upgradeCmd.Flags().Bool("all", false, "Upgrade all extensions")
	extension_upgradeCmd.Flags().Bool("dry-run", false, "Only display upgrades")
	extension_upgradeCmd.Flags().Bool("force", false, "Force upgrade extension")
	extensionCmd.AddCommand(extension_upgradeCmd)

	carapace.Gen(extension_upgradeCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if extension_upgradeCmd.Flag("all").Changed {
				return carapace.ActionValues()
			}
			return action.ActionExtensions()
		}),
	)
}
