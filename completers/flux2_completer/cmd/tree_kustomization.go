package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var tree_kustomizationCmd = &cobra.Command{
	Use:   "kustomization",
	Short: "Print the resource inventory of a Kustomization",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tree_kustomizationCmd).Standalone()
	tree_kustomizationCmd.Flags().Bool("compact", false, "list Flux resources only.")
	tree_kustomizationCmd.Flags().StringP("output", "o", "", "the format in which the tree should be printed. can be 'json' or 'yaml'")
	treeCmd.AddCommand(tree_kustomizationCmd)
}
