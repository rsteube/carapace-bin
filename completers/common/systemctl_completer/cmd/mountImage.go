package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var mountImageCmd = &cobra.Command{
	Use:   "mount-image",
	Short: "Mount an image from the host into a unit's namespace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mountImageCmd).Standalone()

	rootCmd.AddCommand(mountImageCmd)
}
