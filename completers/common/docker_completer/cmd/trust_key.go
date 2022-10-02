package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var trust_keyCmd = &cobra.Command{
	Use:   "key",
	Short: "Manage keys for signing Docker images",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(trust_keyCmd).Standalone()

	trustCmd.AddCommand(trust_keyCmd)
}
