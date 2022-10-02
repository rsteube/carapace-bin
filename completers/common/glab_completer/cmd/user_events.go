package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var user_eventsCmd = &cobra.Command{
	Use:   "events",
	Short: "View user events",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(user_eventsCmd).Standalone()
	user_eventsCmd.Flags().BoolP("all", "a", false, "Get events from all projects")
	userCmd.AddCommand(user_eventsCmd)
}
