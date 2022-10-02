package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var listClientsCmd = &cobra.Command{
	Use:   "list-clients",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listClientsCmd).Standalone()

	listClientsCmd.Flags().StringS("F", "F", "", "format")
	listClientsCmd.Flags().StringS("t", "t", "", "target-session")
	rootCmd.AddCommand(listClientsCmd)
}
