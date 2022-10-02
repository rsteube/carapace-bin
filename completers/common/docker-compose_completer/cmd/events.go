package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/docker-compose_completer/cmd/action"
	"github.com/spf13/cobra"
)

var eventsCmd = &cobra.Command{
	Use:   "events",
	Short: "Receive real time events from containers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eventsCmd).Standalone()
	eventsCmd.Flags().Bool("json", false, "Output events as a stream of json objects")
	rootCmd.AddCommand(eventsCmd)

	carapace.Gen(eventsCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionServices(eventsCmd).Invoke(c).Filter(c.Args).ToA()
		}),
	)
}
