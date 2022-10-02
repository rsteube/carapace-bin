package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var pkgnamesCmd = &cobra.Command{
	Use:   "pkgnames",
	Short: "print the name of each package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pkgnamesCmd).Standalone()

	rootCmd.AddCommand(pkgnamesCmd)
}
