package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pinky",
	Short: "lightweight finger",
	Long:  "https://linux.die.net/man/1/pinky",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("b", "b", false, "omit the user's home directory and shell in long format")
	rootCmd.Flags().BoolS("f", "f", false, "omit the line of column headings in short format")
	rootCmd.Flags().BoolS("h", "h", false, "omit the user's project file in long format")
	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().BoolS("i", "i", false, "omit the user's full name and remote host in short format")
	rootCmd.Flags().BoolS("l", "l", false, "produce long format output for the specified USERs")
	rootCmd.Flags().BoolS("p", "p", false, "omit the user's plan file in long format")
	rootCmd.Flags().BoolS("q", "q", false, "omit the user's full name, remote host and idle time")
	rootCmd.Flags().BoolS("s", "s", false, "do short format output, this is the default")
	rootCmd.Flags().Bool("version", false, "output version information and exit")
	rootCmd.Flags().BoolS("w", "w", false, "omit the user's full name in short format")

	carapace.Gen(rootCmd).PositionalAnyCompletion(os.ActionUsers())
}
