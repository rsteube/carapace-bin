package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var providers_lockCmd = &cobra.Command{
	Use:   "lock",
	Short: "Write out dependency locks for the configured provide",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(providers_lockCmd).Standalone()

	providers_lockCmd.Flags().StringS("fs-mirror", "fs-mirror", "", "Consult the given filesystem mirror directory")
	providers_lockCmd.Flags().StringS("net-mirror", "net-mirror", "", "Consult the given network mirror")
	providers_lockCmd.Flags().StringS("platform", "platform", "", "Choose a target platform to request package checksums for")
	providersCmd.AddCommand(providers_lockCmd)

	providers_lockCmd.Flag("fs-mirror").NoOptDefVal = " "
	providers_lockCmd.Flag("net-mirror").NoOptDefVal = " "
	providers_lockCmd.Flag("platform").NoOptDefVal = " "

	carapace.Gen(providers_lockCmd).FlagCompletion(carapace.ActionMap{
		"fs-mirror": carapace.ActionDirectories(),
	})
}
