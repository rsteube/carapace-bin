package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var uiCmd = &cobra.Command{
	Use:   "ui",
	Short: "Open the web UI",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(uiCmd).Standalone()

	uiCmd.Flags().Bool("authenticate", false, "Creates a new invite token and passes it to the UI for authorization.")

	addGlobalOptions(uiCmd)

	rootCmd.AddCommand(uiCmd)
}
