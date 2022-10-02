package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ion",
	Short: "The Ion Shell",
	Long:  "https://gitlab.redox-os.org/redox-os/ion/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("c", "c", "", "Evaluate given commands instead of reading from the commandline")
	rootCmd.Flags().BoolP("fake-interactive", "f", false, "Use a fake interactive mode, where errors don't exit the shell")
	rootCmd.Flags().BoolP("help", "h", false, "Prints help information")
	rootCmd.Flags().BoolP("interactive", "i", false, "Force interactive mode")
	rootCmd.Flags().BoolP("no-execute", "n", false, "Do not execute any commands, perform only syntax checking")
	rootCmd.Flags().StringS("o", "o", "", "Shortcut layout.")
	rootCmd.Flags().BoolP("version", "v", false, "Print the version, platform and revision of Ion then exit")
	rootCmd.Flags().BoolS("x", "x", false, "Print commands before execution")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"o": carapace.ActionValues("vi", "emacs"),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if !rootCmd.Flag("c").Changed {
				return carapace.ActionFiles()
			} else {
				return carapace.ActionValues()
			}
		}),
	)
}
