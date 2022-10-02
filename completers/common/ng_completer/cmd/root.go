package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ng",
	Short: "The Angular CLI",
	Long:  "https://angular.io/cli",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.PersistentFlags().Bool("help", false, "Shows a help message for this command in the console")
	rootCmd.Flags().BoolS("h", "h", false, "Shows commands")
}
