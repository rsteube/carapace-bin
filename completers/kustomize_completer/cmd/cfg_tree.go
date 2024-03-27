package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var cfg_treeCmd = &cobra.Command{
	Use:   "tree",
	Short: "[Alpha] Display Resource structure from a directory or stdin.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cfg_treeCmd).Standalone()
	cfg_treeCmd.Flags().Bool("all", false, "print all field infos")
	cfg_treeCmd.Flags().Bool("args", false, "print args field")
	cfg_treeCmd.Flags().Bool("command", false, "print command field")
	cfg_treeCmd.Flags().Bool("env", false, "print env field")
	cfg_treeCmd.Flags().Bool("exclude-non-local", false, "if true, exclude non-local-config in the output.")
	cfg_treeCmd.Flags().StringSlice("field", []string{}, "print field")
	cfg_treeCmd.Flags().String("graph-structure", "", "Graph structure to use for printing the tree.  may be any of: owners,directory")
	cfg_treeCmd.Flags().Bool("image", false, "print image field")
	cfg_treeCmd.Flags().Bool("include-local", false, "if true, include local-config in the output.")
	cfg_treeCmd.Flags().Bool("name", false, "print name field")
	cfg_treeCmd.Flags().Bool("ports", false, "print ports field")
	cfg_treeCmd.Flags().Bool("replicas", false, "print replicas field")
	cfg_treeCmd.Flags().Bool("resources", false, "print resources field")
	cfgCmd.AddCommand(cfg_treeCmd)

	carapace.Gen(cfg_treeCmd).FlagCompletion(carapace.ActionMap{
		"graph-structure": carapace.ActionValues("owners", "directory"),
	})

	carapace.Gen(cfg_treeCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
