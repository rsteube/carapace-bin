package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "xonsh",
	Short: "Python-powered shell",
	Long:  "https://xon.sh/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("D", "D", "", "Define an environment variable")
	rootCmd.Flags().StringS("c", "c", "", "Run a single command and exit.")
	rootCmd.Flags().Bool("cache-everything", false, "Use a cache, even for interactive commands.")
	rootCmd.Flags().BoolP("help", "h", false, "Show help and exit.")
	rootCmd.Flags().BoolP("interactive", "i", false, "Force running in interactive mode.")
	rootCmd.Flags().BoolP("login", "l", false, "Run as a login shell.")
	rootCmd.Flags().Bool("no-rc", false, "Do not load the .xonshrc files.")
	rootCmd.Flags().Bool("no-script-cache", false, "Do not cache scripts as they are run.")
	rootCmd.Flags().String("rc", "", "The xonshrc files to load")
	rootCmd.Flags().String("shell-type", "", "What kind of shell should be used.")
	rootCmd.Flags().Bool("timings", false, "Prints timing information before the prompt is shown.")
	rootCmd.Flags().BoolP("version", "V", false, "Show version information and exit.")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"rc": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionFiles()
		}),
		"shell-type": carapace.ActionValues("b", "best", "d", "dumb", "ptk", "ptk1", "ptk2", "prompt-toolkit", "prompt_toolkit", "prompt-toolkit1", "prompt-toolkit2", "prompt-toolkit3", "prompt_toolkit3", "ptk3", "rand", "random", "rl", "readline"),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
