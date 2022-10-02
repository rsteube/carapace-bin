package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var setWindowCmd = &cobra.Command{
	Use:   "set_window",
	Short: "Set properties about a window",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setWindowCmd).Standalone()

	setWindowCmd.Flags().String("class", "", "Set window class")
	setWindowCmd.Flags().String("classname", "", "Set window class name")
	setWindowCmd.Flags().String("icon-name", "", "Set window WM_ICON_NAME")
	setWindowCmd.Flags().String("name", "", "Set window WM_NAME")
	setWindowCmd.Flags().String("overrideredirect", "", "Set window's override_redirect value")
	setWindowCmd.Flags().String("role", "", "Set window WM_WINDOW_ROLE")
	setWindowCmd.Flags().String("urgency", "", "Set window urgency hint")
	rootCmd.AddCommand(setWindowCmd)
}
