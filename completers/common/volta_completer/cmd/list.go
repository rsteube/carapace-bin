package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Displays the current toolchain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listCmd).Standalone()

	listCmd.Flags().BoolP("current", "c", false, "Show the currently-active tool(s)")
	listCmd.Flags().BoolP("default", "d", false, "Show your default tool(s).")
	listCmd.Flags().String("format", "", "Specify the output format [possible values: human, plain]")
	listCmd.Flags().BoolP("help", "h", false, "Prints help information")
	listCmd.Flags().Bool("quiet", false, "Prevents unnecessary output")
	listCmd.Flags().Bool("verbose", false, "Enables verbose diagnostics")
	rootCmd.AddCommand(listCmd)

	carapace.Gen(listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("human", "plain"),
	})
}
