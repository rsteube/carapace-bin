package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var engine_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a local engine",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(engine_updateCmd).Standalone()

	engine_updateCmd.Flags().String("containerd", "", "override default location of containerd endpoint")
	engine_updateCmd.Flags().String("engine-image", "", "Specify engine image (default uses the same image as currently")
	engine_updateCmd.Flags().String("registry-prefix", "", "Override the current location where engine images are pulled")
	engine_updateCmd.Flags().String("version", "", "Specify engine version")
	engineCmd.AddCommand(engine_updateCmd)
}
