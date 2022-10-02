package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "show branches and their commits",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(showCmd).Standalone()

	showCmd.Flags().StringS("L", "L", "", "Process line range n,m in file, counting from 1")
	showCmd.Flags().String("decorate", "", "decorate options")
	showCmd.Flags().String("decorate-refs", "", "only decorate refs that match <pattern>")
	showCmd.Flags().String("decorate-refs-exclude", "", "do not decorate refs that match <pattern>")
	showCmd.Flags().Bool("mailmap", false, "alias of --use-mailmap")
	showCmd.Flags().BoolP("quiet", "q", false, "suppress diff output")
	showCmd.Flags().Bool("source", false, "show source")
	showCmd.Flags().Bool("use-mailmap", false, "Use mail map file")
	rootCmd.AddCommand(showCmd)

	showCmd.Flag("decorate").NoOptDefVal = "short"

	carapace.Gen(showCmd).FlagCompletion(carapace.ActionMap{
		"decorate": carapace.ActionValuesDescribed(
			"short", "do not print ref prefixes",
			"full", "print ref prefixes",
			"auto", "short format when output to terminal",
			"no", "no decoration",
		),
	})

	carapace.Gen(showCmd).PositionalAnyCompletion(
		carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return git.ActionRefs(git.RefOption{}.Default())
			case 1:
				return git.ActionRefFiles(c.Parts[0])
			default:
				return carapace.ActionFiles()
			}
		}),
	)
}
