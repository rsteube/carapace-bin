package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var cfg_grepCmd = &cobra.Command{
	Use:   "grep",
	Short: "[Alpha] Search for matching Resources in a directory or from stdin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cfg_grepCmd).Standalone()
	cfg_grepCmd.Flags().Bool("annotate", true, "annotate resources with their file origins.")
	cfg_grepCmd.Flags().Bool("invert-match", false, "Selected Resources are those not matching any of the specified patterns..")
	cfg_grepCmd.Flags().BoolP("recurse-subpackages", "R", true, "also print resources recursively in all the nested subpackages")
	cfgCmd.AddCommand(cfg_grepCmd)

	carapace.Gen(cfg_grepCmd).PositionalCompletion(
		carapace.ActionValues(),
		carapace.ActionDirectories(),
	)
}
