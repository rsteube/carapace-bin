package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a particular image on the cluster",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runCmd).Standalone()

	runCmd.Flags().Bool("allow-missing-template-keys", false, "If true, ignore any errors in templates when a field or map key is missing in the template. Only app")
	runCmd.Flags().String("annotations", "", "Annotations to apply to the pod.")
	runCmd.Flags().Bool("attach", false, "If true, wait for the Pod to start running, and then attach to the Pod as if 'kubectl attach ...' we")
	runCmd.Flags().String("cascade", "", "Must be \"background\", \"orphan\", or \"foreground\". Selects the deletion cascading strategy for the dep")
	runCmd.Flags().Bool("command", false, "If true and extra arguments are present, use them as the 'command' field in the container, rather th")
	runCmd.Flags().String("dry-run", "", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent,")
	runCmd.Flags().String("env", "", "Environment variables to set in the container.")
	runCmd.Flags().Bool("expose", false, "If true, service is created for the container(s) which are run")
	runCmd.Flags().String("field-manager", "", "Name of the manager used to track field ownership.")
	runCmd.Flags().StringP("filename", "f", "", "to use to replace the resource.")
	runCmd.Flags().Bool("force", false, "If true, immediately remove resources from API and bypass graceful deletion. Note that immediate del")
	runCmd.Flags().String("generator", "", "The name of the API generator to use, see http://kubernetes.io/docs/user-guide/kubectl-conventions/#")
	runCmd.Flags().String("grace-period", "", "Period of time in seconds given to the resource to terminate gracefully. Ignored if negative. Set to")
	runCmd.Flags().String("hostport", "", "The host port mapping for the container port. To demonstrate a single-machine container.")
	runCmd.Flags().String("image", "", "The image for the container to run.")
	runCmd.Flags().String("image-pull-policy", "", "The image pull policy for the container. If left empty, this value will not be specified by the clie")
	runCmd.Flags().StringP("kustomize", "k", "", "Process a kustomization directory. This flag can't be used together with -f or -R.")
	runCmd.Flags().StringP("labels", "l", "", "Comma separated labels to apply to the pod(s). Will override previous values.")
	runCmd.Flags().Bool("leave-stdin-open", false, "If the pod is started in interactive mode or with stdin, leave stdin open after the first attach com")
	runCmd.Flags().String("limits", "", "The resource requirement limits for this container.  For example, 'cpu=200m,memory=512Mi'.  Note tha")
	runCmd.Flags().StringP("output", "o", "", "Output format. One of: json|yaml|name|go-template|go-template-file|template|templatefile|jsonpath|js")
	runCmd.Flags().String("overrides", "", "An inline JSON override for the generated object. If this is non-empty, it is used to override the g")
	runCmd.Flags().String("pod-running-timeout", "", "The length of time (like 5s, 2m, or 3h, higher than zero) to wait until at least one pod is running")
	runCmd.Flags().String("port", "", "The port that this container exposes.")
	runCmd.Flags().Bool("privileged", false, "If true, run the container in privileged mode.")
	runCmd.Flags().Bool("quiet", false, "If true, suppress prompt messages.")
	runCmd.Flags().Bool("record", false, "Record current kubectl command in the resource annotation. If set to false, do not record the comman")
	runCmd.Flags().BoolP("recursive", "R", false, "Process the directory used in -f, --filename recursively. Useful when you want to manage related man")
	runCmd.Flags().StringP("replicas", "r", "", "Number of replicas to create for this container. Default is 1.")
	runCmd.Flags().String("requests", "", "The resource requirement requests for this container.  For example, 'cpu=100m,memory=256Mi'.  Note t")
	runCmd.Flags().String("restart", "", "The restart policy for this Pod.  Legal values [Always, OnFailure, Never].")
	runCmd.Flags().Bool("rm", false, "If true, delete resources created in this command for attached containers.")
	runCmd.Flags().Bool("save-config", false, "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotat")
	runCmd.Flags().String("schedule", "", "A schedule in the Cron format the job should be run with.")
	runCmd.Flags().String("service-generator", "", "The name of the generator to use for creating a service.  Only used if --expose is true")
	runCmd.Flags().String("service-overrides", "", "An inline JSON override for the generated service object. If this is non-empty, it is used to overri")
	runCmd.Flags().String("serviceaccount", "", "Service account to set in the pod spec.")
	runCmd.Flags().BoolP("stdin", "i", false, "Keep stdin open on the container(s) in the pod, even if nothing is attached.")
	runCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The templa")
	runCmd.Flags().String("timeout", "", "The length of time to wait before giving up on a delete, zero means determine a timeout from the siz")
	runCmd.Flags().BoolP("tty", "t", false, "Allocated a TTY for each container in the pod.")
	runCmd.Flags().Bool("wait", false, "If true, wait for resources to be gone before returning. This waits for finalizers.")
	rootCmd.AddCommand(runCmd)

	carapace.Gen(runCmd).FlagCompletion(carapace.ActionMap{
		"dry-run":           action.ActionDryRunModes(),
		"image-pull-policy": carapace.ActionValues("Never", "Always", "IfNotPresent"),
		"kustomize":         carapace.ActionDirectories(),
		"output":            action.ActionOutputFormats(),
		"restart":           carapace.ActionValues("Always", "OnFailure", "Never"),
		"template":          carapace.ActionFiles(),
	})
}
