package cmd

import (
	"github.com/spf13/cobra"
)

var columnCmd = &cobra.Command{
	Use:   "column",
	Short: "Display data in columns",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	columnCmd.Flags().String("command", "", "lookup config vars")
	columnCmd.Flags().String("indent", "", "Padding space on left border")
	columnCmd.Flags().String("mode", "", "layout to use")
	columnCmd.Flags().String("nl", "", "Padding space on right border")
	columnCmd.Flags().String("padding", "", "Padding space between columns")
	columnCmd.Flags().String("raw-mode", "", "layout to use")
	columnCmd.Flags().String("width", "", "Maximum width")
	rootCmd.AddCommand(columnCmd)
}
