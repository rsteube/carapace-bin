package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/rsteube/carapace-bin/pkg/util"
	"github.com/spf13/cobra"
)

var whatchangedCmd = &cobra.Command{
	Use:   "whatchanged",
	Short: "Show logs with difference each commit introduces",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	carapace.Gen(whatchangedCmd).Standalone()

	whatchangedCmd.Flags().StringS("L", "L", "", "trace the evolution of line range")
	whatchangedCmd.Flags().String("decorate", "", "decorate options")
	whatchangedCmd.Flags().String("decorate-refs", "", "only decorate refs that match <pattern>")
	whatchangedCmd.Flags().String("decorate-refs-exclude", "", "do not decorate refs that match <pattern>")
	whatchangedCmd.Flags().Bool("mailmap", false, "alias of --use-mailmap")
	whatchangedCmd.Flags().BoolP("quiet", "q", false, "suppress diff output")
	whatchangedCmd.Flags().Bool("source", false, "show source")
	whatchangedCmd.Flags().Bool("use-mailmap", false, "use mail map file")
	rootCmd.AddCommand(whatchangedCmd)

	whatchangedCmd.Flag("decorate").NoOptDefVal = "short"

	carapace.Gen(whatchangedCmd).FlagCompletion(carapace.ActionMap{
		"decorate": carapace.ActionValuesDescribed(
			"short", "do not print ref prefixes",
			"full", "print ref prefixes",
			"auto", "short format when output to terminal",
			"no", "no decoration",
		),
	})

	carapace.Gen(whatchangedCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if util.HasPathPrefix(c.CallbackValue) {
				return carapace.ActionFiles()
			} else {
				return carapace.ActionMultiParts("...", func(c carapace.Context) carapace.Action {
					if len(c.Parts) < 2 {
						return git.ActionRefs(git.RefOption{}.Default())
					} else {
						return carapace.ActionValues()
					}
				})
			}
		}),
	)
}
