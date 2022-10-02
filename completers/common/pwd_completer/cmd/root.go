package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pwd",
	Short: "Print the full filename of the current working directory",
	Long:  "https://linux.die.net/man/1/pwd",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().BoolP("logical", "L", false, "use PWD from environment, even if it contains symlinks")
	rootCmd.Flags().BoolP("physical", "P", false, "avoid all symlinks")
	rootCmd.Flags().Bool("version", false, "output version information and exit")
}
