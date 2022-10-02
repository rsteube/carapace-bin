package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var operatorCmd = &cobra.Command{
	Use:   "operator",
	Short: "Provides cluster-level tools for Consul operator",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(operatorCmd).Standalone()

	rootCmd.AddCommand(operatorCmd)
}
