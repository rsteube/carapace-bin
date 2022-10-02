package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sync",
	Short: "Synchronize cached writes to persistent storage",
	Long:  "https://linux.die.net/man/8/sync",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("data", "d", false, "sync only file data, no unneeded metadata")
	rootCmd.Flags().BoolP("file-system", "f", false, "sync the file systems that contain the files")
	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().Bool("version", false, "output version information and exit")
}
