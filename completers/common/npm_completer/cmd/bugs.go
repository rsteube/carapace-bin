package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/npm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var bugsCmd = &cobra.Command{
	Use:   "bugs",
	Short: "Report bugs for a package in a web browser",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bugsCmd).Standalone()
	bugsCmd.Flags().String("browser", "", "browser called to open websites")

	rootCmd.AddCommand(bugsCmd)

	carapace.Gen(bugsCmd).FlagCompletion(carapace.ActionMap{
		"browser": carapace.ActionValuesDescribed(
			"open", "osx default",
			"start", "windows default",
			"xdg-open", "others default",
			"false", "print urls to terminal",
			"true", "use system default",
		),
	})

	carapace.Gen(bugsCmd).PositionalAnyCompletion(
		action.ActionPackages(bugsCmd),
	)
}
