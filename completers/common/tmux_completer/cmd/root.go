package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tmux",
	Short: "terminal multiplexer",
	Long:  "https://github.com/tmux/tmux/wiki",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("2", "2", false, "Force tmux to assume the terminal supports 256 colours")
	rootCmd.Flags().BoolS("C", "C", false, "Start in control mode")
	rootCmd.Flags().BoolS("D", "D", false, "Do not start the tmux server as a daemon")
	rootCmd.Flags().StringS("L", "L", "", "specify socket name")
	rootCmd.Flags().BoolS("N", "N", false, "Do not start the server even if the command would normally do so")
	rootCmd.Flags().StringS("S", "S", "", "Specify a full alternative path to the server socket")
	rootCmd.Flags().StringS("T", "T", "", "Set terminal features for the client")
	rootCmd.Flags().BoolS("V", "V", false, "Report the tmux version.")
	rootCmd.Flags().StringS("c", "c", "", "Execute shell-command using the default shell")
	rootCmd.Flags().StringS("f", "f", "", "Specify an alternative configuration file")
	rootCmd.Flags().BoolS("l", "l", false, "Behave as a login shell")
	rootCmd.Flags().BoolS("u", "u", false, "Write UTF-8 output to the terminal")
	rootCmd.Flags().BoolS("v", "v", false, "Request verbose logging")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"S": carapace.ActionFiles(),
		"T": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return tmux.ActionFeatures().Invoke(c).Filter(c.Parts).ToA()
		}),
		"c": carapace.ActionFiles(),
		"f": carapace.ActionFiles(),
	})
}
