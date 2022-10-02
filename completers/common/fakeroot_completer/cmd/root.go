package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fakeroot",
	Short: "run a command in an environment faking root privileges for file manipulation",
	Long:  "https://linux.die.net/man/1/fakeroot-sysv",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("b", "b", "", "Specify  fd  base  (TCP  mode only).")
	rootCmd.Flags().String("faked", "", "Specify an alternative binary to use as faked.")
	rootCmd.Flags().BoolS("h", "h", false, "Display help.")
	rootCmd.Flags().StringS("i", "i", "", "Load a fakeroot environment previously saved using -s from load-file.")
	rootCmd.Flags().StringP("lib", "l", "", "Specify an alternative wrapper library.")
	rootCmd.Flags().StringS("s", "s", "", "Save the fakeroot environment to save-file on exit.")
	rootCmd.Flags().BoolP("unknown-is-real", "u", false, "Use the real ownership of files previously unknown to fakeroot.")
	rootCmd.Flags().BoolS("v", "v", false, "Display version.")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"faked": carapace.ActionFiles(),
		"i":     carapace.ActionFiles(),
		"lib":   carapace.ActionFiles(),
		"s":     carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
