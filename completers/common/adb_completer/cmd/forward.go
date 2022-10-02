package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var forwardCmd = &cobra.Command{
	Use:   "forward",
	Short: "forward socket connection",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forwardCmd).Standalone()
	forwardCmd.Flags().Bool("list", false, "list all forward socket connections")
	forwardCmd.Flags().Bool("no-rebind", false, "don't replace existing connection")
	forwardCmd.Flags().Bool("remove", false, "remove specific forward socket connections")
	forwardCmd.Flags().Bool("remove-all", false, "remove all forward socket connections")

	rootCmd.AddCommand(forwardCmd)

	// TODO LOCAL/REMOTE positional completion
}
