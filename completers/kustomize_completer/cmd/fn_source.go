package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var fn_sourceCmd = &cobra.Command{
	Use:   "source",
	Short: "[Alpha] Implement a Source by reading a local directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fn_sourceCmd).Standalone()
	fn_sourceCmd.Flags().String("function-config", "", "path to function config.")
	fn_sourceCmd.Flags().String("wrap-kind", "ResourceList", "output using this format.")
	fn_sourceCmd.Flags().String("wrap-version", "config.kubernetes.io/v1", "output using this format.")
	fnCmd.AddCommand(fn_sourceCmd)
}
