package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var getmouselocationCmd = &cobra.Command{
	Use:   "getmouselocation",
	Short: "Outputs the x, y, screen, and window id of the mouse cursor",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(getmouselocationCmd).Standalone()
	getmouselocationCmd.Flags().Bool("shell", false, "This makes getmouselocation output shell data you can eval")

	rootCmd.AddCommand(getmouselocationCmd)
}
