package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/docker-compose_completer/cmd/action"
	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Removes stopped service containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rmCmd).Standalone()
	rmCmd.Flags().BoolP("all", "a", false, "Deprecated - no effect")
	rmCmd.Flags().BoolP("force", "f", false, "Don't ask to confirm removal")
	rmCmd.Flags().BoolP("stop", "s", false, "Stop the containers, if required, before removing")
	rmCmd.Flags().BoolP("volumes", "v", false, "Remove any anonymous volumes attached to containers")
	rootCmd.AddCommand(rmCmd)

	carapace.Gen(rmCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionServices(rmCmd).Invoke(c).Filter(c.Args).ToA()
		}),
	)
}
