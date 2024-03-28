package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var export_image_policyCmd = &cobra.Command{
	Use:   "policy",
	Short: "Export ImagePolicy resources in YAML format",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(export_image_policyCmd).Standalone()
	export_imageCmd.AddCommand(export_image_policyCmd)
}
