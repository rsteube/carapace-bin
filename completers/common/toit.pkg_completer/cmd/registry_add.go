package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var registry_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(registry_addCmd).Standalone()
	registry_addCmd.Flags().Bool("local", false, "Registry is local")
	registryCmd.AddCommand(registry_addCmd)

	carapace.Gen(registry_addCmd).PositionalCompletion(
		carapace.ActionValues(),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if registry_addCmd.Flag("local").Changed {
				return carapace.ActionDirectories()
			}
			return git.ActionRepositorySearch()
		}),
	)
}
