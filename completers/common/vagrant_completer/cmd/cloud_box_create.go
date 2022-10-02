package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var cloud_box_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates an empty box entry on Vagrant Cloud",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloud_box_createCmd).Standalone()

	cloud_box_createCmd.Flags().StringP("description", "d", "", "Full description of the box")
	cloud_box_createCmd.Flags().Bool("no-private", false, "Makes box public")
	cloud_box_createCmd.Flags().BoolP("private", "p", false, "Makes box private")
	cloud_boxCmd.AddCommand(cloud_box_createCmd)
}
