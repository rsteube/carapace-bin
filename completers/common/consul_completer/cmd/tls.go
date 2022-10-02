package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var tlsCmd = &cobra.Command{
	Use:   "tls",
	Short: "Builtin helpers for creating CAs and certificates",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tlsCmd).Standalone()

	rootCmd.AddCommand(tlsCmd)
}
