package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push OpenFaaS functions to remote registry (Docker Hub)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	pushCmd.Flags().Bool("envsubst", true, "Substitute environment variables in stack.yml file")
	pushCmd.Flags().Int("parallel", 1, "Push images in parallel to depth specified.")
	pushCmd.Flags().String("tag", "latest", "Override latest tag on function Docker image, accepts 'latest', 'sha', 'branch', 'describe'")
	rootCmd.AddCommand(pushCmd)

	carapace.Gen(pushCmd).FlagCompletion(carapace.ActionMap{
		"tag": carapace.ActionValues("latest", "sha", "branch", "describe"),
	})
}
