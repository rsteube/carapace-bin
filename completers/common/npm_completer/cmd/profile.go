package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "Change settings on your registry profile",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(profileCmd).Standalone()
	profileCmd.PersistentFlags().Bool("json", false, "output as json")
	profileCmd.PersistentFlags().BoolP("parseable", "p", false, "output parseable results")
	profileCmd.PersistentFlags().String("otp", "", "one-time password")

	rootCmd.AddCommand(profileCmd)
}
