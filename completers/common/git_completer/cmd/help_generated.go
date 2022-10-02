package cmd

import (
	"github.com/spf13/cobra"
)

var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Display help information about Git",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	helpCmd.Flags().BoolP("all", "a", false, "print all available commands")
	helpCmd.Flags().BoolP("config", "c", false, "print all configuration variable names")
	helpCmd.Flags().BoolP("guides", "g", false, "print list of useful guides")
	helpCmd.Flags().BoolP("info", "i", false, "show info page")
	helpCmd.Flags().BoolP("man", "m", false, "show man page")
	helpCmd.Flags().BoolP("verbose", "v", false, "print command description")
	helpCmd.Flags().BoolP("web", "w", false, "show manual in web browser")
	rootCmd.AddCommand(helpCmd)
}
