package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var mod_verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "verify dependencies have expected content",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mod_verifyCmd).Standalone()

	modCmd.AddCommand(mod_verifyCmd)
}
