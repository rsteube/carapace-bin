package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a command with custom Node, npm, and/or Yarn versions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runCmd).Standalone()

	runCmd.Flags().Bool("bundled-npm", false, "Forces npm to be the version bundled with Node")
	runCmd.Flags().String("env", "", "Set an environment variable (can be used multiple times)")
	runCmd.Flags().BoolP("help", "h", false, "Prints help information")
	runCmd.Flags().Bool("no-yarn", false, "Disables Yarn")
	runCmd.Flags().String("node", "", "Set the custom Node version")
	runCmd.Flags().String("npm", "", "Set the custom npm version")
	runCmd.Flags().Bool("quiet", false, "Prevents unnecessary output")
	runCmd.Flags().Bool("verbose", false, "Enables verbose diagnostics")
	runCmd.Flags().String("yarn", "", "Set the custom Yarn version")
	rootCmd.AddCommand(runCmd)
}
