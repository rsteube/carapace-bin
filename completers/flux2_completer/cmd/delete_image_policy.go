package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var delete_image_policyCmd = &cobra.Command{
	Use:   "policy",
	Short: "Delete an ImagePolicy object",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(delete_image_policyCmd).Standalone()
	delete_imageCmd.AddCommand(delete_image_policyCmd)
}
