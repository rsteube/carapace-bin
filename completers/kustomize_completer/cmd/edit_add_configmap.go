package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var edit_add_configmapCmd = &cobra.Command{
	Use:   "configmap",
	Short: "Adds a configmap to the kustomization file.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(edit_add_configmapCmd).Standalone()
	edit_add_configmapCmd.Flags().String("behavior", "", "Specify the behavior for config map generation, i.e whether to create a new configmap (the default),  to merge with a previously defined one, or to replace an existing one. Merge and replace should be used only  when overriding an existing configmap defined in a base")
	edit_add_configmapCmd.Flags().Bool("disableNameSuffixHash", false, "Disable the name suffix for the configmap")
	edit_add_configmapCmd.Flags().String("from-env-file", "", "Specify the path to a file to read lines of key=val pairs to create a configmap (i.e. a Docker .env file).")
	edit_add_configmapCmd.Flags().StringSlice("from-file", []string{}, "Key file can be specified using its file path, in which case file basename will be used as configmap key, or optionally with a key and file path, in which case the given key will be used.  Specifying a directory will iterate each named file in the directory whose basename is a valid configmap key.")
	edit_add_configmapCmd.Flags().StringArray("from-literal", []string{}, "Specify a key and literal value to insert in configmap (i.e. mykey=somevalue)")
	edit_addCmd.AddCommand(edit_add_configmapCmd)
}
