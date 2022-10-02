package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var trust_key_generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate and load a signing key-pair",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(trust_key_generateCmd).Standalone()

	trust_key_generateCmd.Flags().String("dir", "", "Directory to generate key in, defaults to current directory")
	trust_keyCmd.AddCommand(trust_key_generateCmd)

	carapace.Gen(trust_key_generateCmd).FlagCompletion(carapace.ActionMap{
		"dir": carapace.ActionDirectories(),
	})
}
