package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var config_imagesCmd = &cobra.Command{
	Use:   "images",
	Short: "Interact with container images used by kubeadm",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_imagesCmd).Standalone()
	configCmd.AddCommand(config_imagesCmd)
}
