package cmd

import (
	"github.com/spf13/cobra"
)

var symbolic_refCmd = &cobra.Command{
	Use:   "symbolic-ref",
	Short: "Read, modify and delete symbolic refs",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	symbolic_refCmd.Flags().BoolP("delete", "d", false, "delete symbolic ref")
	symbolic_refCmd.Flags().StringS("m", "m", "", "reason of the update")
	symbolic_refCmd.Flags().BoolP("quiet", "q", false, "suppress error message for non-symbolic (detached) refs")
	symbolic_refCmd.Flags().Bool("short", false, "shorten ref output")
	rootCmd.AddCommand(symbolic_refCmd)
}
