package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a new chart with the given name",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	createCmd.Flags().StringP("starter", "p", "", "the name or absolute path to Helm starter scaffold")
	rootCmd.AddCommand(createCmd)

	carapace.Gen(createCmd).FlagCompletion(carapace.ActionMap{
		"starter": carapace.ActionFiles(),
	})
}
