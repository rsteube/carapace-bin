package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var cacheCmd = &cobra.Command{
	Use:    "cache",
	Short:  "Modify the syntax-definition and theme cache",
	Run:    func(cmd *cobra.Command, args []string) {},
	Hidden: true,
}

func init() {
	cacheCmd.Flags().Bool("blank", false, "Create completely new syntax and theme sets (instead of appending to the default sets).")
	cacheCmd.Flags().BoolP("build", "b", false, "Initialize or update the syntax/theme cache by loading from the source directory.")
	cacheCmd.Flags().BoolP("clear", "c", false, "Remove the cached syntax definitions and themes.")
	cacheCmd.Flags().BoolP("help", "h", false, "Prints help information")
	cacheCmd.Flags().String("source", "", "Use a different directory to load syntaxes and themes from.")
	cacheCmd.Flags().String("target", "", "Use a different directory to store the cached syntax and theme set.")
	rootCmd.AddCommand(cacheCmd)

	carapace.Gen(cacheCmd).FlagCompletion(carapace.ActionMap{
		"source": carapace.ActionDirectories(),
		"target": carapace.ActionDirectories(),
	})
}
