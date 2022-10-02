package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var timingsCmd = &cobra.Command{
	Use:   "timings",
	Short: "Prints timings of all active modules",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(timingsCmd).Standalone()

	timingsCmd.Flags().BoolP("help", "h", false, "Prints help information")
	timingsCmd.Flags().BoolP("version", "V", false, "Prints version information")
	rootCmd.AddCommand(timingsCmd)
}
