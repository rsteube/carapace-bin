package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var artifactCmd = &cobra.Command{
	Use:   "artifact",
	Short: "Artifact and build management",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(artifactCmd).Standalone()

	rootCmd.AddCommand(artifactCmd)
}
