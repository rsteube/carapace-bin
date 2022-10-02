package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "batdiff",
	Short: "Diff a file against the current git index, or display the diff between two files",
	Long:  "https://github.com/eth-p/bat-extras/blob/master/doc/batdiff.md",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("color", false, "Force color output.")
	rootCmd.Flags().StringP("context", "C", "", "The number of lines to show before and after the differing lines.")
	rootCmd.Flags().Bool("delta", false, "Display diffs using delta.")
	rootCmd.Flags().Bool("no-color", false, "Force disable color output.")
	rootCmd.Flags().String("pager", "", "Specify the pager to use.")
	rootCmd.Flags().String("paging", "", "Enable/disable paging.")
	rootCmd.Flags().String("terminal-width", "", "Generate output for the specified terminal width.")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"pager": carapace.Batch(
			os.ActionPathExecutables(),
			carapace.ActionFiles(),
		).ToA(),
		"paging": carapace.ActionValues("always", "never").StyleF(style.ForKeyword),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
		carapace.ActionFiles(),
	)
}
