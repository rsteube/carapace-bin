package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var gcCmd = &cobra.Command{
	Use:   "gc",
	Short: "Cleanup unnecessary files and optimize the local repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gcCmd).Standalone()

	gcCmd.Flags().Bool("aggressive", false, "be more thorough (increased runtime)")
	gcCmd.Flags().Bool("auto", false, "enable auto-gc mode")
	gcCmd.Flags().Bool("force", false, "force running gc even if there may be another gc running")
	gcCmd.Flags().Bool("keep-largest-pack", false, "repack all other packs except the largest pack")
	gcCmd.Flags().String("prune", "", "prune unreferenced objects")
	gcCmd.Flags().BoolP("quiet", "q", false, "suppress progress reporting")
	rootCmd.AddCommand(gcCmd)

	gcCmd.Flag("prune").NoOptDefVal = " "
}
