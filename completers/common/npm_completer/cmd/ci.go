package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var ciCmd = &cobra.Command{
	Use:   "ci",
	Short: "Install a project with a clean slate",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ciCmd).Standalone()
	ciCmd.Flags().Bool("ignore-scripts", false, "Disable scripts")
	ciCmd.Flags().String("script-shell", "", "shell to use for scripts")

	rootCmd.AddCommand(ciCmd)

	carapace.Gen(ciCmd).FlagCompletion(carapace.ActionMap{
		"script-shell": os.ActionShells(),
	})
}
