package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/npm"
	"github.com/spf13/cobra"
)

var exploreCmd = &cobra.Command{
	Use:   "explore",
	Short: "Browse an installed package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(exploreCmd).Standalone()
	exploreCmd.Flags().String("shell", "", "shell to use")

	rootCmd.AddCommand(exploreCmd)

	carapace.Gen(exploreCmd).FlagCompletion(carapace.ActionMap{
		"shell": os.ActionShells(),
	})

	carapace.Gen(exploreCmd).PositionalCompletion(
		npm.ActionDependencyNames(),
	)
}
