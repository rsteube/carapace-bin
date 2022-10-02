package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop a package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stopCmd).Standalone()
	stopCmd.Flags().Bool("ignore-scripts", false, "do not run scripts specified in package.json")
	stopCmd.Flags().String("script-shell", "", "shell to use")

	rootCmd.AddCommand(stopCmd)

	carapace.Gen(stopCmd).FlagCompletion(carapace.ActionMap{
		"script-shell": os.ActionShells(),
	})
}
