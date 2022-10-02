package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validates a Vagrantfile config",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(validateCmd).Standalone()

	validateCmd.Flags().BoolP("ignore-provider", "p", false, "Ignores provider config options")
	rootCmd.AddCommand(validateCmd)
}
