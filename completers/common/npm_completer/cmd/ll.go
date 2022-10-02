package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/npm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var llCmd = &cobra.Command{
	Use:   "ll",
	Short: "List installed packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(llCmd).Standalone()
	llCmd.Flags().BoolP("all", "a", false, "show all outdated or installed packages")
	llCmd.Flags().String("depth", "", "depth to ge when recursing packages")
	llCmd.Flags().BoolP("global", "g", false, "operate in global mode")
	llCmd.Flags().Bool("json", false, "output as json")
	llCmd.Flags().Bool("link", false, "output only packages that are linked")
	llCmd.Flags().BoolP("long", "l", false, "show extended information")
	llCmd.Flags().StringArray("omit", []string{}, "omit dependency types")
	llCmd.Flags().Bool("package-lock-only", false, "only use the package-lock.json")
	llCmd.Flags().BoolP("parseable", "p", false, "output parseable results")
	llCmd.Flags().Bool("unicode", false, "output unicode characters")
	addWorkspaceFlags(llCmd)

	rootCmd.AddCommand(llCmd)

	carapace.Gen(llCmd).FlagCompletion(carapace.ActionMap{
		"omit": carapace.ActionValues("dev", "optional", "peer"),
	})

	carapace.Gen(llCmd).PositionalAnyCompletion(
		action.ActionPackages(llCmd),
	)
}
