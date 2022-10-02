package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var drainCmd = &cobra.Command{
	Use:   "drain",
	Short: "Drain node in preparation for maintenance",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(drainCmd).Standalone()

	drainCmd.Flags().Bool("delete-emptydir-data", false, "Continue even if there are pods using emptyDir (local data that will be deleted when the node is dra")
	drainCmd.Flags().Bool("delete-local-data", false, "Continue even if there are pods using emptyDir (local data that will be deleted when the node is dra")
	drainCmd.Flags().Bool("disable-eviction", false, "Force drain to use delete, even if eviction is supported. This will bypass checking PodDisruptionBud")
	drainCmd.Flags().String("dry-run", "", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent,")
	drainCmd.Flags().Bool("force", false, "Continue even if there are pods not managed by a ReplicationController, ReplicaSet, Job, DaemonSet o")
	drainCmd.Flags().String("grace-period", "", "Period of time in seconds given to each pod to terminate gracefully. If negative, the default value ")
	drainCmd.Flags().Bool("ignore-daemonsets", false, "Ignore DaemonSet-managed pods.")
	drainCmd.Flags().String("pod-selector", "", "Label selector to filter pods on the node")
	drainCmd.Flags().StringP("selector", "l", "", "Selector (label query) to filter on")
	drainCmd.Flags().String("skip-wait-for-delete-timeout", "", "If pod DeletionTimestamp older than N seconds, skip waiting for the pod.  Seconds must be greater th")
	drainCmd.Flags().String("timeout", "", "The length of time to wait before giving up, zero means infinite")
	rootCmd.AddCommand(drainCmd)

	carapace.Gen(drainCmd).FlagCompletion(carapace.ActionMap{
		"dry-run": action.ActionDryRunModes(),
	})

	carapace.Gen(drainCmd).PositionalCompletion(
		action.ActionResources("", "nodes"),
	)
}
