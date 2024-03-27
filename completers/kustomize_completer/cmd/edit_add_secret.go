package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var edit_add_secretCmd = &cobra.Command{
	Use:   "secret",
	Short: "Adds a secret to the kustomization file.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(edit_add_secretCmd).Standalone()
	edit_add_secretCmd.Flags().Bool("disableNameSuffixHash", false, "Disable the name suffix for the secret")
	edit_add_secretCmd.Flags().String("from-env-file", "", "Specify the path to a file to read lines of key=val pairs to create a secret (i.e. a Docker .env file).")
	edit_add_secretCmd.Flags().StringSlice("from-file", []string{}, "Key file can be specified using its file path, in which case file basename will be used as secret key, or optionally with a key and file path, in which case the given key will be used.  Specifying a directory will iterate each named file in the directory whose basename is a valid secret key.")
	edit_add_secretCmd.Flags().StringArray("from-literal", []string{}, "Specify a key and literal value to insert in secret (i.e. mykey=somevalue)")
	edit_add_secretCmd.Flags().String("namespace", "", "Specify the namespace of the secret")
	edit_add_secretCmd.Flags().String("type", "Opaque", "Specify the secret type this can be 'Opaque' (default), or 'kubernetes.io/tls'")
	edit_addCmd.AddCommand(edit_add_secretCmd)
}
