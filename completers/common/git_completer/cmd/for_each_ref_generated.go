package cmd

import (
	"github.com/spf13/cobra"
)

var for_each_refCmd = &cobra.Command{
	Use:   "for-each-ref",
	Short: "Output information on each ref",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	for_each_refCmd.Flags().String("color", "", "respect format colors")
	for_each_refCmd.Flags().String("contains", "", "print only refs which contain the commit")
	for_each_refCmd.Flags().String("count", "", "show only <n> matched refs")
	for_each_refCmd.Flags().String("format", "", "format to use for the output")
	for_each_refCmd.Flags().Bool("ignore-case", false, "sorting and filtering are case insensitive")
	for_each_refCmd.Flags().String("merged", "", "print only refs that are merged")
	for_each_refCmd.Flags().String("no-contains", "", "print only refs which don't contain the commit")
	for_each_refCmd.Flags().String("no-merged", "", "print only refs that are not merged")
	for_each_refCmd.Flags().BoolP("perl", "p", false, "quote placeholders suitably for perl")
	for_each_refCmd.Flags().String("points-at", "", "print only refs which points at the given object")
	for_each_refCmd.Flags().Bool("python", false, "quote placeholders suitably for python")
	for_each_refCmd.Flags().BoolP("shell", "s", false, "quote placeholders suitably for shells")
	for_each_refCmd.Flags().String("sort", "", "field name to sort on")
	for_each_refCmd.Flags().Bool("tcl", false, "quote placeholders suitably for Tcl")
	rootCmd.AddCommand(for_each_refCmd)
}
