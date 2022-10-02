package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/color"
	"github.com/spf13/cobra"
)

var label_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create labels for repository/project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(label_createCmd).Standalone()
	label_createCmd.Flags().StringP("color", "c", "#428BCA", "Color of label in plain or HEX code. (Default: #428BCA)")
	label_createCmd.Flags().StringP("description", "d", "", "Label description")
	label_createCmd.Flags().StringP("name", "n", "", "Name of label")
	labelCmd.AddCommand(label_createCmd)

	carapace.Gen(label_createCmd).FlagCompletion(carapace.ActionMap{
		"color": color.ActionHexColors(),
		"name":  action.ActionLabels(label_createCmd),
	})
}
