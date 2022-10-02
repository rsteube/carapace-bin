package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fish",
	Short: "the friendly interactive shell",
	Long:  "https://fishshell.com/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("command", "c", "", "evaluate the specified commands")
	rootCmd.Flags().StringP("debug", "d", "", "enables debug output")
	rootCmd.Flags().StringP("debug-output", "o", "", "Specify a file path to receive the debug output")
	rootCmd.Flags().StringP("debug-stack-frames", "D", "", "specify  how  many  stack frames to display")
	rootCmd.Flags().StringP("features", "f", "", "enables one or more feature flags")
	rootCmd.Flags().StringP("init-command", "C", "", "evaluate the specified commands after reading  the configuration")
	rootCmd.Flags().BoolP("interactive", "i", false, "specify that fish is to run in interactive mode")
	rootCmd.Flags().BoolP("login", "l", false, "specify that fish is to run as a login shell")
	rootCmd.Flags().BoolP("no-execute", "n", false, "do not execute any commands")
	rootCmd.Flags().Bool("print-debug-categories", false, "outputs the list of debug categories")
	rootCmd.Flags().Bool("print-rusage-self", false, "output stats from getrusage")
	rootCmd.Flags().BoolP("private", "P", false, "enables private mode")
	rootCmd.Flags().StringP("profile", "p", "", "output timing  information to the specified file")
	rootCmd.Flags().BoolP("version", "v", false, "display version and exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"debug-output": carapace.ActionFiles(),
		"features": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValuesDescribed(
				"stderr-nocaret", "^ no longer redirects stderr",
				"qmark-noglob", "? no longer globs",
				"regex-easyesc", `string replace -r needs fewer \'s`,
			).Invoke(c).Filter(c.Parts).ToA()
		}),
		"profile": carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(carapace.ActionFiles())
}
