package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/vagrant_completer/cmd/action"
	"github.com/spf13/cobra"
)

var winrmConfigCmd = &cobra.Command{
	Use:   "winrm-config",
	Short: "outputs WinRM configuration to connect to the machine",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(winrmConfigCmd).Standalone()

	winrmConfigCmd.Flags().String("host", "", "Name the host for the config")
	rootCmd.AddCommand(winrmConfigCmd)

	carapace.Gen(winrmConfigCmd).PositionalCompletion(
		action.ActionMachines(),
	)
}
