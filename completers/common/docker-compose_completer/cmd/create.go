package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/docker-compose_completer/cmd/action"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates containers for a service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(createCmd).Standalone()
	createCmd.Flags().Bool("build", false, "Build images before starting containers.")
	createCmd.Flags().Bool("force-recreate", false, "Recreate containers even if their configuration and image haven't changed.")
	createCmd.Flags().Bool("no-build", false, "Don't build an image, even if it's missing.")
	createCmd.Flags().Bool("no-recreate", false, "If containers already exist, don't recreate them. Incompatible with --force-recreate.")
	rootCmd.AddCommand(createCmd)

	carapace.Gen(createCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionServices(createCmd).Invoke(c).Filter(c.Args).ToA()
		}),
	)
}
