package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var emuCmd = &cobra.Command{
	Use:   "emu",
	Short: "run emulator console command",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emuCmd).Standalone()

	rootCmd.AddCommand(emuCmd)
}
