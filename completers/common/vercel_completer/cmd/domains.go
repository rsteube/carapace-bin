package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var domainsCmd = &cobra.Command{
	Use:   "domains",
	Short: "Manages your domain names",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(domainsCmd).Standalone()

	rootCmd.AddCommand(domainsCmd)
}
