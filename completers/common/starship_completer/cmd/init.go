package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Prints the shell function used to execute starship",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(initCmd).Standalone()

	initCmd.Flags().BoolP("help", "h", false, "Prints help information")
	initCmd.Flags().Bool("print-full-init", false, "Print the main initialization script (as opposed to the init stub)")
	initCmd.Flags().BoolP("version", "V", false, "Prints version information")
	rootCmd.AddCommand(initCmd)

	carapace.Gen(initCmd).PositionalCompletion(
		carapace.ActionValues("zsh", "bash", "fish", "powershell", "elvish"),
	)
}
