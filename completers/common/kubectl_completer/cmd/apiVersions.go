package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var apiVersionsCmd = &cobra.Command{
	Use:   "api-versions",
	Short: "Print the supported API versions on the server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apiVersionsCmd).Standalone()

	rootCmd.AddCommand(apiVersionsCmd)
}
