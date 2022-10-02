package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gofmt",
	Short: "format Go source code",
	Long:  "https://pkg.go.dev/cmd/gofmt",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("cpuprofile", "", "write cpu profile to this file")
	rootCmd.Flags().BoolS("d", "d", false, "display diffs instead of rewriting files")
	rootCmd.Flags().BoolS("e", "e", false, "report all errors (not just the first 10 on different lines)")
	rootCmd.Flags().BoolS("l", "l", false, "list files whose formatting differs from gofmt's")
	rootCmd.Flags().StringS("r", "r", "", "rewrite rule")
	rootCmd.Flags().BoolS("s", "s", false, "simplify code")
	rootCmd.Flags().BoolS("w", "w", false, "write result to (source) file instead of stdout")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"cpuprofile": carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
