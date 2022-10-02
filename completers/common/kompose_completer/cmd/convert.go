package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert a Docker Compose file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	convertCmd.Flags().String("build", "none", "Set the type of build (\"local\"|\"build-config\"(OpenShift only)|\"none\")")
	convertCmd.Flags().String("build-branch", "", "Specify repository branch to use for buildconfig (default master)")
	convertCmd.Flags().String("build-repo", "", "Specify source repository for buildconfig (default remote origin)")
	convertCmd.Flags().BoolP("chart", "c", false, "Create a Helm chart for converted objects")
	convertCmd.Flags().String("controller", "", "Set the output controller (\"deployment\"|\"daemonSet\"|\"replicationController\")")
	convertCmd.Flags().Bool("daemon-set", false, "Generate a Kubernetes daemonset object (deprecated, use --controller instead)")
	convertCmd.Flags().BoolP("deployment", "d", false, "Generate a Kubernetes deployment object (deprecated, use --controller instead)")
	convertCmd.Flags().Bool("deployment-config", true, "Generate an OpenShift deploymentconfig object")
	convertCmd.Flags().Bool("emptyvols", false, "Use Empty Volumes. Do not generate PVCs")
	convertCmd.Flags().Int("indent", 2, "Spaces length to indent generated yaml files")
	convertCmd.Flags().Bool("insecure-repository", false, "Use an insecure Docker repository for OpenShift ImageStream")
	convertCmd.Flags().BoolP("json", "j", false, "Generate resource files into JSON format")
	convertCmd.Flags().Bool("multiple-container-mode", false, "Create multiple containers grouped by 'kompose.service.group' label")
	convertCmd.Flags().StringP("out", "o", "", "Specify a file name or directory to save objects to (if path does not exist, a file will be created)")
	convertCmd.Flags().Bool("push-image", false, "If we should push the docker image we built")
	convertCmd.Flags().String("push-image-registry", "", "Specify registry for pushing image, which will override registry from image name.")
	convertCmd.Flags().String("pvc-request-size", "", "Specify the size of pvc storage requests in the generated resource spec")
	convertCmd.Flags().Int("replicas", 1, "Specify the number of replicas in the generated resource spec")
	convertCmd.Flags().Bool("replication-controller", false, "Generate a Kubernetes replication controller object (deprecated, use --controller instead)")
	convertCmd.Flags().String("service-group-mode", "", "Group multiple service to create single workload by `label`(`kompose.service.group`) or `volume`(shared volumes)")
	convertCmd.Flags().String("service-group-name", "", "Using with --service-group-mode=volume to specific a final service name for the group")
	convertCmd.Flags().Bool("stdout", false, "Print converted objects to stdout")
	convertCmd.Flags().String("volumes", "persistentVolumeClaim", "Volumes to be generated (\"persistentVolumeClaim\"|\"emptyDir\"|\"hostPath\" | \"configMap\")")
	convertCmd.Flags().Bool("with-kompose-annotation", true, "Add kompose annotations to generated resource")
	convertCmd.Flags().BoolP("yaml", "y", false, "Generate resource files into YAML format")
	rootCmd.AddCommand(convertCmd)

	carapace.Gen(convertCmd).FlagCompletion(carapace.ActionMap{
		"build": carapace.ActionValues("local", "build-config", "none"),
		"build-branch": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if flag := convertCmd.Flag("build-repo"); flag.Changed {
				return git.ActionLsRemoteRefs(git.LsRemoteRefOption{Url: flag.Value.String(), Branches: true})
			}
			return carapace.ActionValues()
		}),
		"build-repo": git.ActionRepositorySearch(),
		"controller": carapace.ActionValues("deployment", "daemonSet", "replicationController"),
		"out":        carapace.ActionFiles(),
		"service-group-mode": carapace.ActionValuesDescribed(
			"label", "kompose.service.group",
			"volume", "shared volumes",
		),
		"volumes": carapace.ActionValues("persistentVolumeClaim", "emptyDir", "hostPath", "configMap"),
	})

	carapace.Gen(convertCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
