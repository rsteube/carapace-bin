package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var describeCmd = &cobra.Command{
	Use:   "describe",
	Short: "Give an object a human readable name based on an available ref",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(describeCmd).Standalone()

	describeCmd.Flags().String("abbrev", "", "use <n> digits to display object names")
	describeCmd.Flags().Bool("all", false, "use any ref")
	describeCmd.Flags().Bool("always", false, "show abbreviated commit object as fallback")
	describeCmd.Flags().String("broken", "", "append <mark> on broken working tree (default: \"-broken\")")
	describeCmd.Flags().String("candidates", "", "consider <n> most recent tags (default: 10)")
	describeCmd.Flags().Bool("contains", false, "find the tag that comes after the commit")
	describeCmd.Flags().Bool("debug", false, "debug search strategy on stderr")
	describeCmd.Flags().String("dirty", "", "append <mark> on dirty working tree (default: \"-dirty\")")
	describeCmd.Flags().Bool("exact-match", false, "only output exact matches")
	describeCmd.Flags().String("exclude", "", "do not consider tags matching <pattern>")
	describeCmd.Flags().Bool("first-parent", false, "only follow first parent")
	describeCmd.Flags().Bool("long", false, "always use long format")
	describeCmd.Flags().String("match", "", "only consider tags matching <pattern>")
	describeCmd.Flags().Bool("tags", false, "use any tag, even unannotated")
	rootCmd.AddCommand(describeCmd)

	describeCmd.Flag("abbrev").NoOptDefVal = "8"
	describeCmd.Flag("dirty").NoOptDefVal = "-dirty"
	describeCmd.Flag("broken").NoOptDefVal = "-broken"

	carapace.Gen(describeCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if describeCmd.Flag("dirty").Changed || describeCmd.Flag("broken").Changed {
				return carapace.ActionValues()
			} else {
				return git.ActionRefs(git.RefOption{}.Default())
			}
		}),
	)
}
