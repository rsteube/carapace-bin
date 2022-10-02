package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "batgrep",
	Short: "Quickly search through and highlight files using ripgrep",
	Long:  "https://github.com/eth-p/bat-extras/blob/master/doc/batgrep.md",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("after-context", "A", "", "Display the next n lines after a matched line.")
	rootCmd.Flags().StringP("before-context", "B", "", "Display the previous n lines before a matched line.")
	rootCmd.Flags().BoolP("case-sensitive", "s", false, "Use case sensitive searching.")
	rootCmd.Flags().Bool("color", false, "Force color output.")
	rootCmd.Flags().StringP("context", "C", "", "A combination of --after-context and --before-context.")
	rootCmd.Flags().BoolP("ignore-case", "i", false, "Use case insensitive searching.")
	rootCmd.Flags().Bool("no-color", false, "Force disable color output.")
	rootCmd.Flags().Bool("no-follow", false, "Do not follow symlinks.")
	rootCmd.Flags().Bool("no-highlight", false, "Do not highlight matching lines.")
	rootCmd.Flags().Bool("no-separator", false, "Disable printing separator between files")
	rootCmd.Flags().Bool("no-snip", false, "Do not show the snip decoration.")
	rootCmd.Flags().String("pager", "", "Specify the pager to use.")
	rootCmd.Flags().String("paging", "", "Enable/disable paging.")
	rootCmd.Flags().BoolP("search-pattern", "p", false, "Tell pager to search for PATTERN. Currently supported pagers: less.")
	rootCmd.Flags().BoolP("smart-case", "S", false, "Use smart case searching.")
	rootCmd.Flags().String("terminal-width", "", "Generate output for the specified terminal width.")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"pager": carapace.Batch(
			os.ActionPathExecutables(),
			carapace.ActionFiles(),
		).ToA(),
		"paging": carapace.ActionValues("always", "never").StyleF(style.ForKeyword),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValues(),
	)

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
