package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ts",
	Short: "timestamp input",
	Long:  "https://linux.die.net/man/1/ts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("i", "i", false, "print incremental timestamps since ast timestamp")
	rootCmd.Flags().BoolS("m", "m", false, "use monotonic clock")
	rootCmd.Flags().BoolS("r", "r", false, "convert existing timestamps to relative times")
	rootCmd.Flags().BoolS("s", "s", false, "print incremental timestamps since start of program")
}
