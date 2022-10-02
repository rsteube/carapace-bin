package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tmate",
	Short: "Instant terminal sharing",
	Long:  "https://tmate.io/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("F", "F", false, "set the foreground mode, useful for setting remote access")
	rootCmd.Flags().StringS("S", "S", "", "set the socket path, useful to issue commands to a running tmate instance")
	rootCmd.Flags().BoolS("V", "V", false, "print version")
	rootCmd.Flags().StringS("f", "f", "", "set the config file path")
	rootCmd.Flags().StringS("k", "k", "", "specify an api-key, necessary for using named sessions on tmate.io")
	rootCmd.Flags().StringS("n", "n", "", "specify the session token instead of getting a random one")
	rootCmd.Flags().StringS("r", "r", "", "same, but for the read-only token")
	rootCmd.Flags().BoolS("v", "v", false, "set verbosity (can be repeated)")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"S": carapace.ActionFiles(),
		"f": carapace.ActionFiles(),
	})

	// TODO tmux commnds
}
