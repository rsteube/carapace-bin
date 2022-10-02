package cmd

import (
	"time"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/helm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "uninstall a release",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	uninstallCmd.Flags().String("description", "", "add a custom description")
	uninstallCmd.Flags().Bool("dry-run", false, "simulate a uninstall")
	uninstallCmd.Flags().Bool("keep-history", false, "remove all associated resources and mark the release as deleted, but retain the release history")
	uninstallCmd.Flags().Bool("no-hooks", false, "prevent hooks from running during uninstallation")
	uninstallCmd.Flags().Duration("timeout", 5*time.Minute, "time to wait for any individual Kubernetes operation (like Jobs for hooks)")
	rootCmd.AddCommand(uninstallCmd)

	carapace.Gen(uninstallCmd).PositionalCompletion(
		action.ActionReleases(),
	)
}
