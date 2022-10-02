package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/consul_completer/cmd/action"
	"github.com/spf13/cobra"
)

var maintCmd = &cobra.Command{
	Use:   "maint",
	Short: "Controls node or service maintenance mode",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(maintCmd).Standalone()
	addClientFlags(maintCmd)
	maintCmd.Flags().Bool("disable", false, "Disable maintenance mode.")
	maintCmd.Flags().Bool("enable", false, "Enable maintenance mode.")
	maintCmd.Flags().String("reason", "", "Text describing the maintenance reason.")
	maintCmd.Flags().String("service", "", "Control maintenance mode for a specific service ID.")
	rootCmd.AddCommand(maintCmd)

	carapace.Gen(maintCmd).FlagCompletion(carapace.ActionMap{
		"service": action.ActionServices(maintCmd),
	})
}
