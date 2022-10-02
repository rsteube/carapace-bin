package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var allowCmd = &cobra.Command{
	Use:     "allow",
	Aliases: []string{"permit", "grant"},
	Short:   "Grants direnv permission to load the given .envrc or .env file",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(allowCmd).Standalone()

	rootCmd.AddCommand(allowCmd)

	carapace.Gen(allowCmd).PositionalCompletion(
		carapace.ActionFiles(".envrc", ".env"),
	)
}
