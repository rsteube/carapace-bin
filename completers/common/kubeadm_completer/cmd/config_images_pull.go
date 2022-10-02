package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var config_images_pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pull images used by kubeadm",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_images_pullCmd).Standalone()
	config_images_pullCmd.PersistentFlags().String("config", "", "Path to a kubeadm configuration file.")
	config_images_pullCmd.PersistentFlags().String("cri-socket", "", "Path to the CRI socket to connect. If empty kubeadm will try to auto-detect this value; use this option only if you have more than one CRI installed or if you have non-standard CRI socket.")
	config_images_pullCmd.PersistentFlags().String("feature-gates", "", "A set of key=value pairs that describe feature gates for various features.")
	config_images_pullCmd.PersistentFlags().String("image-repository", "k8s.gcr.io", "Choose a container registry to pull control plane images from")
	config_images_pullCmd.PersistentFlags().String("kubernetes-version", "stable-1", "Choose a specific Kubernetes version for the control plane.")
	config_imagesCmd.AddCommand(config_images_pullCmd)

	carapace.Gen(config_images_pullCmd).FlagCompletion(carapace.ActionMap{
		"config":     carapace.ActionFiles(),
		"cri-socket": carapace.ActionFiles(),
	})
}
