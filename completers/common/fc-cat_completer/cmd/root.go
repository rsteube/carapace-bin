package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fc-cat",
	Short: "read font information cache files",
	Long:  "https://linux.die.net/man/1/fc-cat",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("help", "h", false, "display this help and exit")
	rootCmd.Flags().BoolP("recurse", "r", false, "recurse into subdirectories")
	rootCmd.Flags().BoolP("verbose", "v", false, "be verbose")
	rootCmd.Flags().BoolP("version", "V", false, "display font config version and exit")

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
