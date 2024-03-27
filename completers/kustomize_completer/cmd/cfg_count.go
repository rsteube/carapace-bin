package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var cfg_countCmd = &cobra.Command{
	Use:   "count",
	Short: "[Alpha] Count Resources Config from a local directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cfg_countCmd).Standalone()
	cfg_countCmd.Flags().Bool("kind", true, "count resources by kind.")
	cfg_countCmd.Flags().BoolP("recurse-subpackages", "R", true, "prints count of resources recursively in all the nested subpackages")
	cfgCmd.AddCommand(cfg_countCmd)

	carapace.Gen(cfg_countCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
