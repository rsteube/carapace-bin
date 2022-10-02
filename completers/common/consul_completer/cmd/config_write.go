package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var config_writeCmd = &cobra.Command{
	Use:   "write",
	Short: "Create or update a centralized config entry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_writeCmd).Standalone()
	addClientFlags(config_writeCmd)
	addServerFlags(config_writeCmd)

	config_writeCmd.Flags().Bool("cas", false, "Perform a Check-And-Set operation.")
	config_writeCmd.Flags().String("modify-index", "", "Unsigned integer representing the ModifyIndex of the config entry.")
	config_writeCmd.Flags().String("namespace", "", "Specifies the namespace to query.")
	configCmd.AddCommand(config_writeCmd)

	carapace.Gen(config_writeCmd).PositionalCompletion(
		carapace.ActionFiles(".json", ".hcl"),
	)
}
