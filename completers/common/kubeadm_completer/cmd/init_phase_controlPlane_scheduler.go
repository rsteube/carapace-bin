package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var init_phase_controlPlane_schedulerCmd = &cobra.Command{
	Use:   "scheduler",
	Short: "Generates the kube-scheduler static Pod manifest",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(init_phase_controlPlane_schedulerCmd).Standalone()
	init_phase_controlPlane_schedulerCmd.Flags().String("cert-dir", "/etc/kubernetes/pki", "The path where to save and store the certificates.")
	init_phase_controlPlane_schedulerCmd.Flags().String("config", "", "Path to a kubeadm configuration file.")
	init_phase_controlPlane_schedulerCmd.Flags().Bool("dry-run", false, "Don't apply any changes; just output what would be done.")
	init_phase_controlPlane_schedulerCmd.Flags().String("image-repository", "k8s.gcr.io", "Choose a container registry to pull control plane images from")
	init_phase_controlPlane_schedulerCmd.Flags().String("kubernetes-version", "stable-1", "Choose a specific Kubernetes version for the control plane.")
	init_phase_controlPlane_schedulerCmd.Flags().String("patches", "", "Path to a directory that contains files named \"target[suffix][+patchtype].extension\".")
	init_phase_controlPlane_schedulerCmd.Flags().String("scheduler-extra-args", "", "A set of extra flags to pass to the Scheduler or override default ones in form of <flagname>=<value>")
	init_phase_controlPlaneCmd.AddCommand(init_phase_controlPlane_schedulerCmd)

	carapace.Gen(init_phase_controlPlane_schedulerCmd).FlagCompletion(carapace.ActionMap{
		"cert-dir": carapace.ActionDirectories(),
		"config":   carapace.ActionFiles(),
		"patches":  carapace.ActionDirectories(),
	})
}
