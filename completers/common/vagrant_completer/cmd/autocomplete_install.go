package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var autocomplete_installCmd = &cobra.Command{
	Use:   "install",
	Short: "install autocomplete",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autocomplete_installCmd).Standalone()

	autocomplete_installCmd.Flags().BoolP("bash", "b", false, "Install bash autocomplete")
	autocomplete_installCmd.Flags().BoolP("zsh", "z", false, "Install zsh autocomplete")
	autocompleteCmd.AddCommand(autocomplete_installCmd)
}
