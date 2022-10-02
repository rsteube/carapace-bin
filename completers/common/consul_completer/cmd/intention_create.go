package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var intention_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create intentions for service connections.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(intention_createCmd).Standalone()
	addClientFlags(intention_createCmd)
	addServerFlags(intention_createCmd)

	intention_createCmd.Flags().Bool("allow", false, "Create an intention that allows when matched.")
	intention_createCmd.Flags().Bool("deny", false, "Create an intention that denies when matched.")
	intention_createCmd.Flags().Bool("file", false, "Read intention data from one or more files.")
	intention_createCmd.Flags().StringArray("meta", []string{}, "Metadata to set on the intention, formatted as key=value.")
	intention_createCmd.Flags().String("namespace", "", "Specifies the namespace to query.")
	intention_createCmd.Flags().Bool("replace", false, "Replace matching intentions.")
	intentionCmd.AddCommand(intention_createCmd)

	// TODO namespace completion

	carapace.Gen(intention_createCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if intention_createCmd.Flag("file").Changed {
				return carapace.ActionFiles()
			}
			// TODO SRC DST completion
			return carapace.ActionValues()
		}),
	)
}
