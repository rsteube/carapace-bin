package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var token_createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a new authentication token",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(token_createCmd).Standalone()
	token_createCmd.Flags().StringArray("cidr", []string{}, "addresses used to limit access")
	token_createCmd.Flags().Bool("read-only", false, "mark token as unable to publish")

	tokenCmd.AddCommand(token_createCmd)
}
