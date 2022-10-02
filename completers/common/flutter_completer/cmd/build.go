package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build an executable app or install bundle",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(buildCmd).Standalone()

	buildCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	rootCmd.AddCommand(buildCmd)
}
