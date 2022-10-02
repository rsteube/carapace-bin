package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var engineCmd = &cobra.Command{
	Use:   "engine",
	Short: "Manage the docker engine",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(engineCmd).Standalone()

	rootCmd.AddCommand(engineCmd)
}
