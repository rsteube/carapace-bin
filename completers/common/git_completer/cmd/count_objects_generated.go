package cmd

import (
	"github.com/spf13/cobra"
)

var count_objectsCmd = &cobra.Command{
	Use:   "count-objects",
	Short: "Count unpacked number of objects and their disk consumption",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	count_objectsCmd.Flags().BoolP("human-readable", "H", false, "print sizes in human readable format")
	count_objectsCmd.Flags().BoolP("verbose", "v", false, "be verbose")
	rootCmd.AddCommand(count_objectsCmd)
}
