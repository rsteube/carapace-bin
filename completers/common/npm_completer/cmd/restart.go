package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restart a package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(restartCmd).Standalone()
	restartCmd.Flags().Bool("ignore-scripts", false, "do not run scripts specified in package.json")
	restartCmd.Flags().String("script-shell", "", "shell to use")

	carapace.Gen(restartCmd).FlagCompletion(carapace.ActionMap{
		"script-shell": os.ActionShells(),
	})

	rootCmd.AddCommand(restartCmd)
}
