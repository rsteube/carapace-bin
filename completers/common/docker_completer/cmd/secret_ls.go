package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var secret_lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List secrets",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secret_lsCmd).Standalone()

	secret_lsCmd.Flags().StringP("filter", "f", "", "Filter output based on conditions provided")
	secret_lsCmd.Flags().String("format", "", "Pretty-print secrets using a Go template")
	secret_lsCmd.Flags().BoolP("quiet", "q", false, "Only display IDs")
	secretCmd.AddCommand(secret_lsCmd)
}
