package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var bisect_visualizeCmd = &cobra.Command{
	Use:     "visualize",
	Aliases: []string{"view"},
	Short:   "show bisect status in gitk",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bisect_visualizeCmd).Standalone()
	bisect_visualizeCmd.Flags().Bool("p", false, "")
	bisect_visualizeCmd.Flags().Bool("stat", false, "")

	bisectCmd.AddCommand(bisect_visualizeCmd)
}
