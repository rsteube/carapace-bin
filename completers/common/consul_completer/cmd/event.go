package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/consul_completer/cmd/action"
	"github.com/spf13/cobra"
)

var eventCmd = &cobra.Command{
	Use:   "event",
	Short: "Fire a new event",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eventCmd).Standalone()
	addClientFlags(eventCmd)

	eventCmd.Flags().String("name", "", "Name of the event.")
	eventCmd.Flags().String("node", "", "Regular expression to filter on node names.")
	eventCmd.Flags().String("service", "", "Regular expression to filter on service instances.")
	eventCmd.Flags().String("tag", "", "Regular expression to filter on service tags.")
	rootCmd.AddCommand(eventCmd)

	carapace.Gen(eventCmd).FlagCompletion(carapace.ActionMap{
		"node":    action.ActionNodes(eventCmd),
		"service": action.ActionServices(eventCmd),
	})
}
