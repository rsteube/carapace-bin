package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var show_keysCmd = &cobra.Command{
	Use:   "keys",
	Short: "Display the known PGP keys",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(show_keysCmd).Standalone()

	show_keysCmd.Flags().BoolP("help", "h", false, "Prints help information")
	show_keysCmd.Flags().BoolP("quiet", "q", false, "Disable progress output")
	show_keysCmd.Flags().BoolP("verbose", "v", false, "Enable verbose output")
	show_keysCmd.Flags().BoolP("version", "V", false, "Prints version information")
	showCmd.AddCommand(show_keysCmd)
}
