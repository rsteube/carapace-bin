package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/vagrant_completer/cmd/action"
	"github.com/spf13/cobra"
)

var winrmCmd = &cobra.Command{
	Use:   "winrm",
	Short: "executes commands on a machine via WinRM",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(winrmCmd).Standalone()

	winrmCmd.Flags().StringP("command", "c", "", "Execute a WinRM command directly")
	winrmCmd.Flags().BoolP("elevated", "e", false, "Run with elevated credentials")
	winrmCmd.Flags().StringP("shell", "s", "", "Use specified shell (powershell, cmd)")
	rootCmd.AddCommand(winrmCmd)

	carapace.Gen(winrmCmd).FlagCompletion(carapace.ActionMap{
		"shell": carapace.ActionValues("powershell", "cmd"),
	})

	carapace.Gen(winrmCmd).PositionalCompletion(
		action.ActionMachines(),
	)
}
