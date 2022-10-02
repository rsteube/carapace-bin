package cmd

import (
	"github.com/spf13/cobra"
)

var fsckCmd = &cobra.Command{
	Use:   "fsck",
	Short: "Verifies the connectivity and validity of the objects in the database",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	fsckCmd.Flags().Bool("cache", false, "make index objects head nodes")
	fsckCmd.Flags().Bool("connectivity-only", false, "check only connectivity")
	fsckCmd.Flags().Bool("dangling", false, "show dangling objects")
	fsckCmd.Flags().Bool("full", false, "also consider packs and alternate objects")
	fsckCmd.Flags().Bool("lost-found", false, "write dangling objects in .git/lost-found")
	fsckCmd.Flags().Bool("name-objects", false, "show verbose names for reachable objects")
	fsckCmd.Flags().Bool("progress", false, "show progress")
	fsckCmd.Flags().Bool("reflogs", false, "make reflogs head nodes (default)")
	fsckCmd.Flags().Bool("root", false, "report root nodes")
	fsckCmd.Flags().Bool("strict", false, "enable more strict checking")
	fsckCmd.Flags().Bool("tags", false, "report tags")
	fsckCmd.Flags().Bool("unreachable", false, "show unreachable objects")
	fsckCmd.Flags().BoolP("verbose", "v", false, "be verbose")
	rootCmd.AddCommand(fsckCmd)
}
