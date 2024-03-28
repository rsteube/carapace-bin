package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var suspend_imageCmd = &cobra.Command{
	Use:   "image",
	Short: "Suspend image automation objects",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(suspend_imageCmd).Standalone()
	suspendCmd.AddCommand(suspend_imageCmd)
}
