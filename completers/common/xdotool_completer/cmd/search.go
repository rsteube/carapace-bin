package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/xdotool_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/rsteube/carapace-bin/pkg/actions/ps"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for windows with titles, names, or classes with a regular expression pattern",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(searchCmd).Standalone()

	searchCmd.Flags().Bool("all", false, "Require that all conditions be met.")
	searchCmd.Flags().Bool("any", false, "Match windows that match any condition.")
	searchCmd.Flags().Bool("class", false, "Match against the window class.")
	searchCmd.Flags().Bool("classname", false, "Match against the window classname.")
	searchCmd.Flags().String("desktop", "", "Only match windows on a certain desktop.")
	searchCmd.Flags().String("limit", "", "Stop searching after finding N matching windows.")
	searchCmd.Flags().String("maxdepth", "", "Set recursion/child search depth.")
	searchCmd.Flags().Bool("name", false, "Match against the window name.")
	searchCmd.Flags().Bool("onlyvisible", false, "Show only visible windows in the results.")
	searchCmd.Flags().String("pid", "", "Match windows that belong to a specific process id.")
	searchCmd.Flags().Bool("role", false, "Match against the window role.")
	searchCmd.Flags().String("screen", "", "Select windows only on a specific screen.")
	searchCmd.Flags().Bool("sync", false, "Block until there are results.")
	rootCmd.AddCommand(searchCmd)

	carapace.Gen(searchCmd).FlagCompletion(carapace.ActionMap{
		"desktop": action.ActionDesktops(),
		"pid":     ps.ActionProcessIds(),
		"screen":  os.ActionScreens(true),
	})
}
