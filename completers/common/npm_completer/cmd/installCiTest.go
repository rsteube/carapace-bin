package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var installCiTestCmd = &cobra.Command{
	Use:   "install-ci-test",
	Short: "Install a project with a clean slate and run tests",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installCiTestCmd).Standalone()
	installCiTestCmd.Flags().Bool("ignore-scripts", false, "do not run scripts specified in package.json")
	installCiTestCmd.Flags().String("script-shell", "", "shell to use")

	rootCmd.AddCommand(installCiTestCmd)

	carapace.Gen(installCiTestCmd).FlagCompletion(carapace.ActionMap{
		"script-shell": os.ActionShells(),
	})
}
