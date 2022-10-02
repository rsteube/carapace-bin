package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var componentCmd = &cobra.Command{
	Use:   "component",
	Short: "Modify a toolchain's installed components",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(componentCmd).Standalone()

	componentCmd.Flags().BoolP("help", "h", false, "Prints help information")
	rootCmd.AddCommand(componentCmd)
}
