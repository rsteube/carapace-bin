package cmd

import (
	"time"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/helm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var rollbackCmd = &cobra.Command{
	Use:   "rollback",
	Short: "roll back a release to a previous revision",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	rollbackCmd.Flags().Bool("cleanup-on-fail", false, "allow deletion of new resources created in this rollback when rollback fails")
	rollbackCmd.Flags().Bool("dry-run", false, "simulate a rollback")
	rollbackCmd.Flags().Bool("force", false, "force resource update through delete/recreate if needed")
	rollbackCmd.Flags().Int("history-max", 10, "limit the maximum number of revisions saved per release. Use 0 for no limit")
	rollbackCmd.Flags().Bool("no-hooks", false, "prevent hooks from running during rollback")
	rollbackCmd.Flags().Bool("recreate-pods", false, "performs pods restart for the resource if applicable")
	rollbackCmd.Flags().Duration("timeout", 5*time.Minute, "time to wait for any individual Kubernetes operation (like Jobs for hooks)")
	rollbackCmd.Flags().Bool("wait", false, "if set, will wait until all Pods, PVCs, Services, and minimum number of Pods of a Deployment, StatefulSet, or ReplicaSet are in a ready state before marking the release as successful. It will wait for as long as --timeout")
	rollbackCmd.Flags().Bool("wait-for-jobs", false, "if set and --wait enabled, will wait until all Jobs have been completed before marking the release as successful. It will wait for as long as --timeout")
	rootCmd.AddCommand(rollbackCmd)

	carapace.Gen(rollbackCmd).PositionalCompletion(
		action.ActionReleases(),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionRevisions(c.Args[0])
		}),
	)
}
