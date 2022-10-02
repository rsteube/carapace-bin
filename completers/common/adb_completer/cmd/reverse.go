package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var reverseCmd = &cobra.Command{
	Use:   "reverse",
	Short: "reverse socket connection",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reverseCmd).Standalone()

	reverseCmd.Flags().Bool("list", false, "list all reverse socket connections from device")
	reverseCmd.Flags().Bool("no-rebind", false, "don't replace existing connection")
	reverseCmd.Flags().Bool("remove", false, "remove specific reverse socket connection")
	reverseCmd.Flags().Bool("remove-all", false, "remove all reverse socket connection")

	rootCmd.AddCommand(reverseCmd)
}
