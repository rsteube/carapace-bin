package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/color"
	"github.com/spf13/cobra"
)

var label_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new label",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(label_createCmd).Standalone()
	label_createCmd.Flags().StringP("color", "c", "", "Color of the label")
	label_createCmd.Flags().StringP("description", "d", "", "Description of the label")
	label_createCmd.Flags().BoolP("force", "f", false, "Update the label color and description if label already exists")
	labelCmd.AddCommand(label_createCmd)

	carapace.Gen(label_createCmd).FlagCompletion(carapace.ActionMap{
		"color": color.ActionHexColors(),
	})
}
