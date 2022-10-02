package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var show_profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "Show the current profile",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(show_profileCmd).Standalone()

	show_profileCmd.Flags().BoolP("help", "h", false, "Prints help information")
	show_profileCmd.Flags().BoolP("quiet", "q", false, "Disable progress output")
	show_profileCmd.Flags().BoolP("verbose", "v", false, "Enable verbose output")
	show_profileCmd.Flags().BoolP("version", "V", false, "Prints version information")
	showCmd.AddCommand(show_profileCmd)
}
