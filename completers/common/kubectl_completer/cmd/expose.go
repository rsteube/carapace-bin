package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var exposeCmd = &cobra.Command{
	Use:   "expose",
	Short: "Expose a resource as a new Kubernetes Service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(exposeCmd).Standalone()

	exposeCmd.Flags().Bool("allow-missing-template-keys", false, "If true, ignore any errors in templates when a field or map key is missing in the template. Only app")
	exposeCmd.Flags().String("cluster-ip", "", "ClusterIP to be assigned to the service. Leave empty to auto-allocate, or set to 'None' to create a ")
	exposeCmd.Flags().String("container-port", "", "Synonym for --target-port")
	exposeCmd.Flags().String("dry-run", "", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent,")
	exposeCmd.Flags().String("external-ip", "", "Additional external IP address (not managed by Kubernetes) to accept for the service. If this IP is ")
	exposeCmd.Flags().String("field-manager", "", "Name of the manager used to track field ownership.")
	exposeCmd.Flags().StringP("filename", "f", "", "Filename, directory, or URL to files identifying the resource to expose a service")
	exposeCmd.Flags().String("generator", "", "The name of the API generator to use. There are 2 generators: 'service/v1' and 'service/v2'. The onl")
	exposeCmd.Flags().StringP("kustomize", "k", "", "Process the kustomization directory. This flag can't be used together with -f or -R.")
	exposeCmd.Flags().StringP("labels", "l", "", "Labels to apply to the service created by this call.")
	exposeCmd.Flags().String("load-balancer-ip", "", "IP to assign to the LoadBalancer. If empty, an ephemeral IP will be created and used (cloud-provider")
	exposeCmd.Flags().String("name", "", "The name for the newly created object.")
	exposeCmd.Flags().StringP("output", "o", "", "Output format. One of: json|yaml|name|go-template|go-template-file|template|templatefile|jsonpath|js")
	exposeCmd.Flags().String("overrides", "", "An inline JSON override for the generated object. If this is non-empty, it is used to override the g")
	exposeCmd.Flags().String("port", "", "The port that the service should serve on. Copied from the resource being exposed, if unspecified")
	exposeCmd.Flags().String("protocol", "", "The network protocol for the service to be created. Default is 'TCP'.")
	exposeCmd.Flags().Bool("record", false, "Record current kubectl command in the resource annotation. If set to false, do not record the comman")
	exposeCmd.Flags().BoolP("recursive", "R", false, "Process the directory used in -f, --filename recursively. Useful when you want to manage related man")
	exposeCmd.Flags().Bool("save-config", false, "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotat")
	exposeCmd.Flags().String("selector", "", "A label selector to use for this service. Only equality-based selector requirements are supported. I")
	exposeCmd.Flags().String("session-affinity", "", "If non-empty, set the session affinity for the service to this; legal values: 'None', 'ClientIP'")
	exposeCmd.Flags().String("target-port", "", "Name or number for the port on the container that the service should direct traffic to. Optional.")
	exposeCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The templa")
	exposeCmd.Flags().String("type", "", "Type for this service: ClusterIP, NodePort, LoadBalancer, or ExternalName. Default is 'ClusterIP'.")
	rootCmd.AddCommand(exposeCmd)

	carapace.Gen(exposeCmd).FlagCompletion(carapace.ActionMap{
		"dry-run":   action.ActionDryRunModes(),
		"filename":  carapace.ActionFiles(),
		"kustomize": carapace.ActionDirectories(),
		"output":    action.ActionOutputFormats(),
		"template":  carapace.ActionFiles(),
		"type":      carapace.ActionValues("ClusterIP", "NodePort", "LoadBalancer", "ExternalName"),
	})

	carapace.Gen(exposeCmd).PositionalCompletion(
		carapace.ActionValues("pod", "service", "replicationcontroller", "deployment", "replicaset"),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionResources("", c.Args[0])
		}),
	)
}
