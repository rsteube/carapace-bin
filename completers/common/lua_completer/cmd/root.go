package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "lua",
	Short: "Lua interpreter",
	Long:  "https://www.lua.org/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("E", "E", false, "ignore environment variables")
	rootCmd.Flags().BoolS("W", "W", false, "turn warnings on")
	rootCmd.Flags().StringS("e", "e", "", "execute string 'stat'")
	rootCmd.Flags().BoolS("i", "i", false, "enter interactive mode after executing 'script'")
	rootCmd.Flags().StringS("l", "l", "", "require library 'name' into global 'name'")
	rootCmd.Flags().BoolS("v", "v", false, "show version information")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(".lua"),
	)
}
