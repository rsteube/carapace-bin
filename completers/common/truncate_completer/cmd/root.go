package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "truncate",
	Short: "Shrink or extend the size of each FILE to the specified size",
	Long:  "https://linux.die.net/man/1/truncate",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().BoolP("io-blocks", "o", false, "treat SIZE as number of IO blocks instead of bytes")
	rootCmd.Flags().BoolP("no-create", "c", false, "do not create any files")
	rootCmd.Flags().StringP("reference", "r", "", "base size on RFILE")
	rootCmd.Flags().StringP("size", "s", "", "set or adjust the file size by SIZE bytes")
	rootCmd.Flags().Bool("version", false, "output version information and exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"reference": carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
