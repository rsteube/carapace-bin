package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var getDevpathCmd = &cobra.Command{
	Use:   "get-devpath",
	Short: "print <device-path>",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(getDevpathCmd).Standalone()

	rootCmd.AddCommand(getDevpathCmd)
}
