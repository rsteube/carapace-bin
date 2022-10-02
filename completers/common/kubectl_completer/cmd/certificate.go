package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var certificateCmd = &cobra.Command{
	Use:   "certificate",
	Short: "Modify certificate resources",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(certificateCmd).Standalone()

	rootCmd.AddCommand(certificateCmd)
}
