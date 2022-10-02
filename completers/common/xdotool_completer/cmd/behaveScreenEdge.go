package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var behaveScreenEdgeCmd = &cobra.Command{
	Use:   "behave_screen_edge",
	Short: "Bind an action to events when the mouse hits the screen edge or corner",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(behaveScreenEdgeCmd).Standalone()

	behaveScreenEdgeCmd.Flags().String("delay", "", "Delay in milliseconds before running the command")
	behaveScreenEdgeCmd.Flags().String("quiesce", "", "Delay in milliseconds before the next command will run")
	rootCmd.AddCommand(behaveScreenEdgeCmd)

	carapace.Gen(behaveScreenEdgeCmd).PositionalCompletion(
		carapace.ActionValues("left", "top-left", "top", "top-right", "right", "bottom-left", "bottom", "bottom-right"),
		// TODO command completion
	)
}
