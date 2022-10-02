package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var show_activeToolchainCmd = &cobra.Command{
	Use:   "active-toolchain",
	Short: "Show the active toolchain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(show_activeToolchainCmd).Standalone()

	show_activeToolchainCmd.Flags().BoolP("help", "h", false, "Prints help information")
	show_activeToolchainCmd.Flags().BoolP("quiet", "q", false, "Disable progress output")
	show_activeToolchainCmd.Flags().BoolP("verbose", "v", false, "Enable verbose output")
	show_activeToolchainCmd.Flags().BoolP("version", "V", false, "Prints version information")
	showCmd.AddCommand(show_activeToolchainCmd)
}
