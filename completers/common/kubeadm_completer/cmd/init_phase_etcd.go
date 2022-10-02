package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var init_phase_etcdCmd = &cobra.Command{
	Use:   "etcd",
	Short: "Generate static Pod manifest file for local etcd",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(init_phase_etcdCmd).Standalone()
	init_phaseCmd.AddCommand(init_phase_etcdCmd)
}
