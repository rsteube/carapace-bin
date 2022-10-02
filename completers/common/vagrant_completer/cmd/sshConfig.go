package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/vagrant_completer/cmd/action"
	"github.com/spf13/cobra"
)

var sshConfigCmd = &cobra.Command{
	Use:   "ssh-config",
	Short: " outputs OpenSSH valid configuration to connect to the machine",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sshConfigCmd).Standalone()

	sshConfigCmd.Flags().String("host", "", "Name the host for the config")
	rootCmd.AddCommand(sshConfigCmd)

	carapace.Gen(sshConfigCmd).PositionalCompletion(
		action.ActionMachines(),
	)
}
