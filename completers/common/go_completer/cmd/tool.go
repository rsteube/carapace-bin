package cmd

import (
	"strings"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var toolCmd = &cobra.Command{
	Use:   "tool",
	Short: "run specified go tool",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(toolCmd).Standalone()

	toolCmd.Flags().BoolS("n", "n", false, "only print the command that would be executed")
	rootCmd.AddCommand(toolCmd)

	carapace.Gen(toolCmd).PositionalCompletion(
		ActionTools(),
	)
}

func ActionTools() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("go", "tool")(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			return carapace.ActionValues(lines[:len(lines)-1]...)
		})
	})
}
