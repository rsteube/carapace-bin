package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "halt",
	Short: "halt the machine",
	Long:  "https://linux.die.net/man/8/halt",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func ExecutePoweroff() error {
	rootCmd.Use = "poweroff"
	rootCmd.Short = "poweroff the machine"
	return Execute()
}

func ExecuteReboot() error {
	rootCmd.Use = "reboot"
	rootCmd.Short = "reboot the machine"
	return Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("force", "f", false, "Force immediate halt/power-off/reboot")
	rootCmd.Flags().Bool("halt", false, "Halt the machine")
	rootCmd.Flags().Bool("help", false, "Show this help")
	rootCmd.Flags().Bool("no-wall", false, "Don't send wall message before halt/power-off/reboot")
	rootCmd.Flags().BoolP("no-wtmp", "d", false, "Don't write wtmp record")
	rootCmd.Flags().BoolP("poweroff", "p", false, "Switch off the machine")
	rootCmd.Flags().Bool("reboot", false, "Reboot the machine")
}
