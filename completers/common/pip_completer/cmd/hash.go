package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var hashCmd = &cobra.Command{
	Use:   "hash",
	Short: "Compute a hash of a local package archive",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hashCmd).Standalone()

	hashCmd.Flags().StringP("algorithm", "a", "", "The hash algorithm to use")
	rootCmd.AddCommand(hashCmd)

	carapace.Gen(hashCmd).FlagCompletion(carapace.ActionMap{
		"algorithm": carapace.ActionValues("sha256", "sha384", "sha512"),
	})

	carapace.Gen(hashCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
