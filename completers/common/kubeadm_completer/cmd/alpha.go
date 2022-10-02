package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var alphaCmd = &cobra.Command{
	Use:   "alpha",
	Short: "Kubeadm experimental sub-commands",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(alphaCmd).Standalone()
	rootCmd.AddCommand(alphaCmd)
}
