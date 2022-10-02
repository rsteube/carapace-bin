package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var certsCmd = &cobra.Command{
	Use:   "certs",
	Short: "Manages your SSL certificates",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(certsCmd).Standalone()

	certsCmd.Flags().String("ca", "", "CA certificate chain file")
	certsCmd.Flags().Bool("challenge-only", false, "Only show challenges needed to issue a cert")
	certsCmd.Flags().String("crt", "", "Certificate file")
	certsCmd.Flags().String("key", "", "Certificate key file")
	certsCmd.Flags().BoolP("next", "N", false, "Show next page of results")
	rootCmd.AddCommand(certsCmd)

	carapace.Gen(certsCmd).FlagCompletion(carapace.ActionMap{
		"ca":  carapace.ActionFiles(),
		"crt": carapace.ActionFiles(),
		"key": carapace.ActionFiles(),
	})
}
