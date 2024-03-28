package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var get_images_policyCmd = &cobra.Command{
	Use:   "policy",
	Short: "Get ImagePolicy status",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(get_images_policyCmd).Standalone()
	get_imagesCmd.AddCommand(get_images_policyCmd)
}
