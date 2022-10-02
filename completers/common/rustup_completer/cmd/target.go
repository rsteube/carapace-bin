package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var targetCmd = &cobra.Command{
	Use:   "target",
	Short: "Modify a toolchain's supported targets",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(targetCmd).Standalone()

	targetCmd.Flags().BoolP("help", "h", false, "Prints help information")
	rootCmd.AddCommand(targetCmd)
}
