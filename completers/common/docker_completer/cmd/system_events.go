package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var system_eventsCmd = &cobra.Command{
	Use:   "events",
	Short: "Get real time events from the server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(system_eventsCmd).Standalone()

	system_eventsCmd.Flags().StringP("filter", "f", "", "Filter output based on conditions provided")
	system_eventsCmd.Flags().String("format", "", "Format the output using the given Go template")
	system_eventsCmd.Flags().String("since", "", "Show all events created since timestamp")
	system_eventsCmd.Flags().String("until", "", "Stream events until this timestamp")
	systemCmd.AddCommand(system_eventsCmd)

	rootAlias(system_eventsCmd, func(cmd *cobra.Command, isAlias bool) {})
}
