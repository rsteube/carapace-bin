package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var fn_wrapCmd = &cobra.Command{
	Use:   "wrap",
	Short: "Wrap an executable so it implements the config fn interface",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fn_wrapCmd).Standalone()
	fn_wrapCmd.Flags().Bool("env-only", true, "only set env vars, not arguments.")
	fn_wrapCmd.Flags().String("wrap-kind", "List", "wrap the input xargs give to the command in this type.")
	fn_wrapCmd.Flags().String("wrap-version", "v1", "wrap the input xargs give to the command in this type.")
	fnCmd.AddCommand(fn_wrapCmd)
}
