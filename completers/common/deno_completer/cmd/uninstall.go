package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstalls an executable script in the installation root's bin directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(uninstallCmd).Standalone()

	uninstallCmd.Flags().BoolP("help", "h", false, "Print help information")
	uninstallCmd.Flags().BoolP("quiet", "q", false, "Suppress diagnostic output")
	uninstallCmd.Flags().String("root", "", "Installation root")
	uninstallCmd.Flags().Bool("unstable", false, "Enable unstable features and APIs")
	rootCmd.AddCommand(uninstallCmd)

	carapace.Gen(uninstallCmd).FlagCompletion(carapace.ActionMap{
		"root": carapace.ActionDirectories(),
	})

	// TODO positional completion
}
