package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var powershellCmd = &cobra.Command{
	Use:   "powershell",
	Short: "connects to machine via powershell remoting",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(powershellCmd).Standalone()

	powershellCmd.Flags().StringP("command", "c", "", "Execute a powershell command directly")
	powershellCmd.Flags().BoolP("elevated", "e", false, "Execute a powershell command with elevated permissions")
	rootCmd.AddCommand(powershellCmd)
}
