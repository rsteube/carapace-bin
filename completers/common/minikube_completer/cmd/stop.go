package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stops a running local Kubernetes cluster",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	stopCmd.Flags().Bool("all", false, "Set flag to stop all profiles (clusters)")
	stopCmd.Flags().Bool("cancel-scheduled", false, "cancel any existing scheduled stop requests")
	stopCmd.Flags().Bool("keep-context-active", false, "keep the kube-context active after cluster is stopped. Defaults to false.")
	stopCmd.Flags().StringP("output", "o", "text", "Format to print stdout in. Options include: [text,json]")
	stopCmd.Flags().String("schedule", "0s", "Set flag to stop cluster after a set amount of time (e.g. --schedule=5m)")
	rootCmd.AddCommand(stopCmd)

	carapace.Gen(stopCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionValues("text", "json"),
	})
}
