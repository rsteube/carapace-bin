package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the kustomize version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(versionCmd).Standalone()
	versionCmd.Flags().Bool("short", false, "short form")
	rootCmd.AddCommand(versionCmd)
}
