package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var box_outdatedCmd = &cobra.Command{
	Use:   "outdated",
	Short: "check if there is a new version for the box",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(box_outdatedCmd).Standalone()

	box_outdatedCmd.Flags().String("cacert", "", "CA certificate for SSL download")
	box_outdatedCmd.Flags().String("capath", "", "CA certificate directory for SSL download")
	box_outdatedCmd.Flags().String("cert", "", "A client SSL cert, if needed")
	box_outdatedCmd.Flags().BoolP("force", "f", false, "Force checks for latest box updates")
	box_outdatedCmd.Flags().Bool("global", false, "Check all boxes installed")
	box_outdatedCmd.Flags().Bool("insecure", false, "Do not validate SSL certificates")
	boxCmd.AddCommand(box_outdatedCmd)

	carapace.Gen(box_outdatedCmd).FlagCompletion(carapace.ActionMap{
		"cacert": carapace.ActionFiles(),
		"capath": carapace.ActionDirectories(),
		"cert":   carapace.ActionFiles(),
	})
}
