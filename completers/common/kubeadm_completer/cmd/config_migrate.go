package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var config_migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Read an older version of the kubeadm configuration API types from a file, and output the similar config object for the newer version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_migrateCmd).Standalone()
	config_migrateCmd.Flags().String("new-config", "", "Path to the resulting equivalent kubeadm config file using the new API version. Optional, if not specified output will be sent to STDOUT.")
	config_migrateCmd.Flags().String("old-config", "", "Path to the kubeadm config file that is using an old API version and should be converted. This flag is mandatory.")
	configCmd.AddCommand(config_migrateCmd)

	carapace.Gen(config_migrateCmd).FlagCompletion(carapace.ActionMap{
		"new-config": carapace.ActionFiles(),
		"old-config": carapace.ActionFiles(),
	})
}
