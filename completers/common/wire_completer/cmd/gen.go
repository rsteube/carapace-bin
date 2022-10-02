package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "generate the wire_gen.go file for each package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(genCmd).Standalone()

	genCmd.Flags().StringS("header_file", "header_file", "", "path to file to insert as a header in wire_gen.go")
	genCmd.Flags().StringS("output_file_prefix", "output_file_prefix", "", "string to prepend to output file names.")
	genCmd.Flags().StringS("tags", "tags", "", "append build tags to the default wirebuild")
	rootCmd.AddCommand(genCmd)

	carapace.Gen(genCmd).FlagCompletion(carapace.ActionMap{
		"header_file": carapace.ActionFiles(),
	})
}
