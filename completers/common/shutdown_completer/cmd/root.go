package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/time"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "shutdown",
	Short: "Shut down the system",
	Long:  "https://linux.die.net/man/8/shutdown",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("c", "c", false, "Cancel a pending shutdown")
	rootCmd.Flags().BoolS("h", "h", false, "Equivalent to --poweroff, overridden by --halt")
	rootCmd.Flags().BoolP("halt", "H", false, "Halt the machine")
	rootCmd.Flags().Bool("help", false, "Show this help")
	rootCmd.Flags().BoolS("k", "k", false, "Don't halt/power-off/reboot, just send warnings")
	rootCmd.Flags().Bool("no-wall", false, "Don't send wall message before halt/power-off/reboot")
	rootCmd.Flags().BoolP("poweroff", "P", false, "Power-off the machine")
	rootCmd.Flags().BoolP("reboot", "r", false, "Reboot the machine")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.Batch(
			carapace.ActionValuesDescribed("now", "immediate shutdown"),
			time.ActionTime(),
			carapace.ActionMultiParts("+", func(c carapace.Context) carapace.Action {
				if len(c.Parts) == 0 {
					return carapace.ActionValuesDescribed("+", "shutdown in m minutes from now")
				}
				return carapace.ActionValues()
			}),
		).ToA(),
	)
}
