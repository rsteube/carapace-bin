package cmd

import (
	"strings"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var moduleCmd = &cobra.Command{
	Use:   "module",
	Short: "Prints a specific prompt module",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(moduleCmd).Standalone()

	moduleCmd.Flags().StringP("cmd-duration", "d", "", "The execution duration of the last command, in milliseconds")
	moduleCmd.Flags().BoolP("help", "h", false, "Prints help information")
	moduleCmd.Flags().StringP("jobs", "j", "", "The number of currently running jobs")
	moduleCmd.Flags().StringP("keymap", "k", "", "The keymap of fish/zsh")
	moduleCmd.Flags().BoolP("list", "l", false, "List out all supported modules")
	moduleCmd.Flags().StringP("path", "p", "", "The path that the prompt should render for")
	moduleCmd.Flags().StringP("status", "s", "", "The status code of the previously run command")
	moduleCmd.Flags().BoolP("version", "V", false, "Prints version information")
	rootCmd.AddCommand(moduleCmd)

	carapace.Gen(moduleCmd).FlagCompletion(carapace.ActionMap{
		"path": carapace.ActionDirectories(),
	})

	carapace.Gen(moduleCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return carapace.ActionExecCommand("starship", "module", "--list")(func(output []byte) carapace.Action {
				lines := strings.Split(string(output), "\n")
				return carapace.ActionValues(lines[2 : len(lines)-1]...)
			})
		}),
	)
}
