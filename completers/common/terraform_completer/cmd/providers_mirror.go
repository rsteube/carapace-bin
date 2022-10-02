package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var providers_mirrorCmd = &cobra.Command{
	Use:   "mirror",
	Short: "Save local copies of all required provider plugin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(providers_mirrorCmd).Standalone()

	providers_mirrorCmd.Flags().StringS("platform", "platform", "", "Choose which target platform to build a mirror for")
	providersCmd.AddCommand(providers_mirrorCmd)

	providers_mirrorCmd.Flag("platform").NoOptDefVal = " "

	// TODO platform completion

	carapace.Gen(providers_mirrorCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
