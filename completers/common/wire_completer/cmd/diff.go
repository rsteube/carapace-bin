package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var diffCmd = &cobra.Command{
	Use:   "diff",
	Short: "output a diff between existing wire_gen.go files and what gen would generate",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(diffCmd).Standalone()

	diffCmd.Flags().StringS("header_file", "header_file", "", "path to file to insert as a header in wire_gen.go")
	diffCmd.Flags().StringS("tags", "tags", "", "append build tags to the default wirebuild")
	rootCmd.AddCommand(diffCmd)

	carapace.Gen(diffCmd).FlagCompletion(carapace.ActionMap{
		"header_file": carapace.ActionFiles(),
	})
}
