package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var hostnameCmd = &cobra.Command{
	Use:   "hostname",
	Short: "Application URLs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hostnameCmd).Standalone()

	rootCmd.AddCommand(hostnameCmd)
}
