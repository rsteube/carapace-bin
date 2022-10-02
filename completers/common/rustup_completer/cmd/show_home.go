package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var show_homeCmd = &cobra.Command{
	Use:   "home",
	Short: "Display the computed value of RUSTUP_HOME",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(show_homeCmd).Standalone()

	show_homeCmd.Flags().BoolP("help", "h", false, "Prints help information")
	show_homeCmd.Flags().BoolP("quiet", "q", false, "Disable progress output")
	show_homeCmd.Flags().BoolP("verbose", "v", false, "Enable verbose output")
	show_homeCmd.Flags().BoolP("version", "V", false, "Prints version information")
	showCmd.AddCommand(show_homeCmd)
}
