package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/npm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var lsCmd = &cobra.Command{
	Use:   "ll",
	Short: "List installed packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lsCmd).Standalone()
	lsCmd.Flags().BoolP("all", "a", false, "show all outdated or installed packages")
	lsCmd.Flags().String("depth", "", "depth to ge when recursing packages")
	lsCmd.Flags().BoolP("global", "g", false, "operate in global mode")
	lsCmd.Flags().Bool("json", false, "output as json")
	lsCmd.Flags().Bool("link", false, "output only packages that are linked")
	lsCmd.Flags().BoolP("long", "l", false, "show extended information")
	lsCmd.Flags().StringArray("omit", []string{}, "omit dependency types")
	lsCmd.Flags().Bool("package-lock-only", false, "only use the package-lock.json")
	lsCmd.Flags().BoolP("parseable", "p", false, "output parseable results")
	lsCmd.Flags().Bool("unicode", false, "output unicode characters")
	addWorkspaceFlags(lsCmd)

	rootCmd.AddCommand(lsCmd)

	carapace.Gen(lsCmd).FlagCompletion(carapace.ActionMap{
		"omit": carapace.ActionValues("dev", "optional", "peer"),
	})

	carapace.Gen(lsCmd).PositionalAnyCompletion(
		action.ActionPackages(lsCmd),
	)
}
