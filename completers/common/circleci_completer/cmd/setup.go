package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Setup the CLI with your credentials",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setupCmd).Standalone()
	setupCmd.Flags().Bool("integration-testing", false, "Enable test mode to bypass interactive UI.")
	setupCmd.Flags().Bool("no-prompt", false, "Disable prompt to bypass interactive UI. (MUST supply --host and --token)")
	rootCmd.AddCommand(setupCmd)
}
