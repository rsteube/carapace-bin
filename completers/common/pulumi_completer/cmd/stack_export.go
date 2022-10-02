package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var stack_exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export a stack's deployment to standard out",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stack_exportCmd).Standalone()
	stack_exportCmd.PersistentFlags().String("file", "", "A filename to write stack output to")
	stack_exportCmd.Flags().Bool("show-secrets", false, "Emit secrets in plaintext in exported stack. Defaults to `false`")
	stack_exportCmd.PersistentFlags().String("version", "", "Previous stack version to export. (If unset, will export the latest.)")
	stackCmd.AddCommand(stack_exportCmd)
}
