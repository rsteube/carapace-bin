package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(searchCmd).Standalone()
	searchCmd.Flags().String("color", "always", "whether to show color in output")
	searchCmd.Flags().Bool("json", false, "output as json")
	searchCmd.Flags().BoolP("long", "l", false, "Search the registry for packages")
	searchCmd.Flags().Bool("no-description", false, "hide descripton")
	searchCmd.Flags().Bool("offline", false, "force offline mode")
	searchCmd.Flags().BoolP("parseable", "p", false, "output parseable results")
	searchCmd.Flags().Bool("prefer-offline", false, "bypass staleness checks for cached data")
	searchCmd.Flags().Bool("prefer-online", false, "force staleness checks for cached data")
	searchCmd.Flags().String("searchexclude", "", "space separated options that limit the results")
	searchCmd.Flags().String("searchopts", "", "space separated search options")

	rootCmd.AddCommand(searchCmd)

	carapace.Gen(searchCmd).FlagCompletion(carapace.ActionMap{
		"color": carapace.ActionValuesDescribed(
			"always", "always show colors",
			"true", "only for file descriptors",
			"false", "never show colors",
		).StyleF(style.ForKeyword),
	})
}
