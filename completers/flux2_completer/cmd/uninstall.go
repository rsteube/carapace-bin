package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall Flux and its custom resource definitions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(uninstallCmd).Standalone()
	uninstallCmd.Flags().Bool("dry-run", false, "only print the objects that would be deleted")
	uninstallCmd.Flags().Bool("keep-namespace", false, "skip namespace deletion")
	uninstallCmd.Flags().BoolP("silent", "s", false, "delete components without asking for confirmation")
	rootCmd.AddCommand(uninstallCmd)
}
