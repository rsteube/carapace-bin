package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var fetchurlCmd = &cobra.Command{
	Use:   "fetchurl",
	Short: "Fetches a given URL into direnv's CAS",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fetchurlCmd).Standalone()

	rootCmd.AddCommand(fetchurlCmd)
}
