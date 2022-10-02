package cmd

import (
	"github.com/spf13/cobra"
)

var repo_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update information of available charts locally from chart repositories",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	repoCmd.AddCommand(repo_updateCmd)
}
