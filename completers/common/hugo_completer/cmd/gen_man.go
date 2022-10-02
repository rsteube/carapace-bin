package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var gen_manCmd = &cobra.Command{
	Use:   "man",
	Short: "Generate man pages for the Hugo CLI",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	gen_manCmd.PersistentFlags().String("dir", "man/", "the directory to write the man pages.")
	genCmd.AddCommand(gen_manCmd)

	carapace.Gen(gen_manCmd).FlagCompletion(carapace.ActionMap{
		"dir": carapace.ActionDirectories(),
	})
}
