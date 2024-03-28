package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var ioamCmd = &cobra.Command{
	Use:   "ioam",
	Short: "manage IOAM namespaces and IOAM schemas",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ioamCmd).Standalone()

	rootCmd.AddCommand(ioamCmd)
}
