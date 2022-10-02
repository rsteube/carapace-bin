package cmd

import (
	"github.com/spf13/cobra"
)

var replaceCmd = &cobra.Command{
	Use:   "replace",
	Short: "Create, list, delete refs to replace objects",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	replaceCmd.Flags().Bool("convert-graft-file", false, "convert existing graft file")
	replaceCmd.Flags().BoolP("delete", "d", false, "delete replace refs")
	replaceCmd.Flags().BoolP("edit", "e", false, "edit existing object")
	replaceCmd.Flags().BoolP("force", "f", false, "replace the ref if it exists")
	replaceCmd.Flags().String("format", "", "use this format")
	replaceCmd.Flags().BoolP("graft", "g", false, "change a commit's parents")
	replaceCmd.Flags().BoolP("list", "l", false, "list replace refs")
	replaceCmd.Flags().Bool("raw", false, "do not pretty-print contents for --edit")
	rootCmd.AddCommand(replaceCmd)
}
