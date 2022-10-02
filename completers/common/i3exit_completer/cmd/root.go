package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "i3exit",
	Short: "exit-script for i3",
	Long:  "https://aur.archlinux.org/packages/i3exit/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValues("lock", "logout", "switch_user", "suspend", "hibernate", "reboot", "shutdown"),
	)
}
