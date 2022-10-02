package cmd

import (
	"github.com/spf13/cobra"
)

var pack_refsCmd = &cobra.Command{
	Use:   "pack-refs",
	Short: "Pack heads and tags for efficient repository access",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	pack_refsCmd.Flags().Bool("all", false, "pack everything")
	pack_refsCmd.Flags().Bool("prune", false, "prune loose refs (default)")
	rootCmd.AddCommand(pack_refsCmd)
}
