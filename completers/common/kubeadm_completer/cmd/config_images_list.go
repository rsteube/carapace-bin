package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var config_images_listCmd = &cobra.Command{
	Use:   "list",
	Short: "Print a list of images kubeadm will use. The configuration file is used in case any images or image repositories are customized",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_images_listCmd).Standalone()
	config_images_listCmd.Flags().Bool("allow-missing-template-keys", true, "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.")
	config_images_listCmd.PersistentFlags().String("config", "", "Path to a kubeadm configuration file.")
	config_images_listCmd.Flags().StringP("experimental-output", "o", "text", "Output format.")
	config_images_listCmd.PersistentFlags().String("feature-gates", "", "A set of key=value pairs that describe feature gates for various features.")
	config_images_listCmd.PersistentFlags().String("image-repository", "k8s.gcr.io", "Choose a container registry to pull control plane images from")
	config_images_listCmd.PersistentFlags().String("kubernetes-version", "stable-1", "Choose a specific Kubernetes version for the control plane.")
	config_images_listCmd.Flags().Bool("show-managed-fields", false, "If true, keep the managedFields when printing objects in JSON or YAML format.")
	config_imagesCmd.AddCommand(config_images_listCmd)

	// TODO feature-gates
	carapace.Gen(config_images_listCmd).FlagCompletion(carapace.ActionMap{
		"config":              carapace.ActionFiles(),
		"experimental-output": carapace.ActionValues("text", "json", "yaml", "go-template", "go-template-file", "template", "templatefile", "jsonpath", "jsonpath-as-json", "jsonpath-file"),
	})
}
