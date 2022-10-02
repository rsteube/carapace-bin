package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var connect_ca_setConfigCmd = &cobra.Command{
	Use:   "set-config",
	Short: "Modify the current Connect CA configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_ca_setConfigCmd).Standalone()
	addClientFlags(connect_ca_setConfigCmd)
	addServerFlags(connect_ca_setConfigCmd)

	connect_ca_setConfigCmd.Flags().String("config-file", "", "The path to the config file to use.")
	connect_ca_setConfigCmd.Flags().Bool("force-without-cross-signing", false, "Indicates that the CA reconfiguration should go ahead even if the current CA is unable to cross sign certificates.")
	connect_caCmd.AddCommand(connect_ca_setConfigCmd)

	carapace.Gen(connect_ca_setConfigCmd).FlagCompletion(carapace.ActionMap{
		"config-file": carapace.ActionFiles(),
	})
}
