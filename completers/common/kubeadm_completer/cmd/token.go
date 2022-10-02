package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var tokenCmd = &cobra.Command{
	Use:   "token",
	Short: "Manage bootstrap tokens",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tokenCmd).Standalone()
	tokenCmd.PersistentFlags().Bool("dry-run", false, "Whether to enable dry-run mode or not")
	tokenCmd.PersistentFlags().String("kubeconfig", "", "The kubeconfig file to use when talking to the cluster. If the flag is not set, a set of standard locations can be searched for an existing kubeconfig file.")
	rootCmd.AddCommand(tokenCmd)

	carapace.Gen(tokenCmd).FlagCompletion(carapace.ActionMap{
		"kubeconfig": carapace.ActionFiles(),
	})
}
