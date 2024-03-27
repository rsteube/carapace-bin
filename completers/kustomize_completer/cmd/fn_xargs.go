package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var fn_xargsCmd = &cobra.Command{
	Use:   "xargs",
	Short: "Convert functionConfig to commandline flags and envs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fn_xargsCmd).Standalone()
	fn_xargsCmd.Flags().Bool("env-only", false, "only add env vars, not flags")
	fn_xargsCmd.Flags().String("wrap-kind", "List", "wrap the input xargs give to the command in this type.")
	fn_xargsCmd.Flags().String("wrap-version", "v1", "wrap the input xargs give to the command in this type.")
	fnCmd.AddCommand(fn_xargsCmd)
}
