package cmd

import (
	"github.com/spf13/cobra"
)

var check_attrCmd = &cobra.Command{
	Use:   "check-attr",
	Short: "Display gitattributes information",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	check_attrCmd.Flags().BoolP("all", "a", false, "report all attributes set on file")
	check_attrCmd.Flags().Bool("cached", false, "use .gitattributes only from the index")
	check_attrCmd.Flags().Bool("stdin", false, "read file names from stdin")
	check_attrCmd.Flags().BoolS("z", "z", false, "terminate input and output records by a NUL character")
	rootCmd.AddCommand(check_attrCmd)
}
