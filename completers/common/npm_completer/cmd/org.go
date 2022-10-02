package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var orgCmd = &cobra.Command{
	Use:   "org",
	Short: "Manage orgs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(orgCmd).Standalone()
	orgCmd.PersistentFlags().String("otp", "", "one-time password")
	orgCmd.PersistentFlags().Bool("json", false, "output as json")
	orgCmd.PersistentFlags().BoolP("parseable", "p", false, "output as json")

	rootCmd.AddCommand(orgCmd)
}
