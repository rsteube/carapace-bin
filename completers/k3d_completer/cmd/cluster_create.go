package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/k3d"
	"github.com/spf13/cobra"
)

var cluster_createCmd = &cobra.Command{
	Use:   "create NAME",
	Short: "Create a new cluster",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cluster_createCmd).Standalone()

	cluster_createCmd.Flags().StringP("agents", "a", "", "Specify how many agents you want to create")
	cluster_createCmd.Flags().String("agents-memory", "", "Memory limit imposed on the agents nodes [From docker]")
	cluster_createCmd.Flags().String("api-port", "", "Specify the Kubernetes API server port exposed on the LoadBalancer (Format: `[HOST:]HOSTPORT`)")
	cluster_createCmd.Flags().StringP("config", "c", "", "Path of a config file to use")
	cluster_createCmd.Flags().StringSliceP("env", "e", []string{}, "Add environment variables to nodes (Format: `KEY[=VALUE][@NODEFILTER[;NODEFILTER...]]`")
	cluster_createCmd.Flags().String("gpus", "", "GPU devices to add to the cluster node containers ('all' to pass all GPUs) [From docker]")
	cluster_createCmd.Flags().StringSlice("host-alias", []string{}, "Add `ip:host[,host,...]` mappings")
	cluster_createCmd.Flags().Bool("host-pid-mode", false, "Enable host pid mode of server(s) and agent(s)")
	cluster_createCmd.Flags().StringP("image", "i", "", "Specify k3s image that you want to use for the nodes")
	cluster_createCmd.Flags().StringSlice("k3s-arg", []string{}, "Additional args passed to k3s command (Format: `ARG@NODEFILTER[;@NODEFILTER]`)")
	cluster_createCmd.Flags().StringSlice("k3s-node-label", []string{}, "Add label to k3s node (Format: `KEY[=VALUE][@NODEFILTER[;NODEFILTER...]]`")
	cluster_createCmd.Flags().Bool("kubeconfig-switch-context", false, "Directly switch the default kubeconfig's current-context to the new cluster's context (requires --kubeconfig-update-default)")
	cluster_createCmd.Flags().Bool("kubeconfig-update-default", false, "Directly update the default kubeconfig with the new cluster's context")
	cluster_createCmd.Flags().StringSlice("lb-config-override", []string{}, "Use dotted YAML path syntax to override nginx loadbalancer settings")
	cluster_createCmd.Flags().String("network", "", "Join an existing network")
	cluster_createCmd.Flags().Bool("no-image-volume", false, "Disable the creation of a volume for importing images")
	cluster_createCmd.Flags().Bool("no-lb", false, "Disable the creation of a LoadBalancer in front of the server nodes")
	cluster_createCmd.Flags().Bool("no-rollback", false, "Disable the automatic rollback actions, if anything goes wrong")
	cluster_createCmd.Flags().StringSliceP("port", "p", []string{}, "Map ports from the node containers (via the serverlb) to the host (Format: `[HOST:][HOSTPORT:]CONTAINERPORT[/PROTOCOL][@NODEFILTER]`)")
	cluster_createCmd.Flags().String("registry-config", "", "Specify path to an extra registries.yaml file")
	cluster_createCmd.Flags().String("registry-create", "", "Create a k3d-managed registry and connect it to the cluster (Format: `NAME[:HOST][:HOSTPORT]`")
	cluster_createCmd.Flags().StringSlice("registry-use", []string{}, "Connect to one or more k3d-managed registries running locally")
	cluster_createCmd.Flags().StringSlice("runtime-label", []string{}, "Add label to container runtime (Format: `KEY[=VALUE][@NODEFILTER[;NODEFILTER...]]`")
	cluster_createCmd.Flags().StringSlice("runtime-ulimit", []string{}, "Add ulimit to container runtime (Format: `NAME[=SOFT]:[HARD]`")
	cluster_createCmd.Flags().StringP("servers", "s", "", "Specify how many servers you want to create")
	cluster_createCmd.Flags().String("servers-memory", "", "Memory limit imposed on the server nodes [From docker]")
	cluster_createCmd.Flags().String("subnet", "", "[Experimental: IPAM] Define a subnet for the newly created container network (Example: `172.28.0.0/16`)")
	cluster_createCmd.Flags().String("timeout", "", "Rollback changes if cluster couldn't be created in specified duration.")
	cluster_createCmd.Flags().String("token", "", "Specify a cluster token. By default, we generate one.")
	cluster_createCmd.Flags().StringSliceP("volume", "v", []string{}, "Mount volumes into the nodes (Format: `[SOURCE:]DEST[@NODEFILTER[;NODEFILTER...]]`")
	cluster_createCmd.Flags().Bool("wait", false, "Wait for the server(s) to be ready before returning. Use '--timeout DURATION' to not wait forever.")
	clusterCmd.AddCommand(cluster_createCmd)

	carapace.Gen(cluster_createCmd).FlagCompletion(carapace.ActionMap{
		"api-port": carapace.ActionValues(), // TODO
		"config":   carapace.ActionFiles(),
		"env": carapace.ActionMultiPartsN("@", 2, func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionValues() // KEY=VALUE
			default:
				return k3d.ActionNodeFilter().List(";")
			}
		}),
		"gpus":               carapace.ActionValues(), // TODO
		"host-alias":         carapace.ActionValues(),
		"image":              carapace.ActionValues(),
		"k3s-arg":            carapace.ActionValues(),
		"k3s-node-label":     carapace.ActionValues(),
		"lb-config-override": carapace.ActionValues(),
		"network":            carapace.ActionValues(),
		"port":               carapace.ActionValues(),
		"registry-config":    carapace.ActionValues(),
		"registry-create":    carapace.ActionValues(),
		"registry-use":       carapace.ActionValues(),
		"runtime-label":      carapace.ActionValues(),
		"runtime-ulimit":     carapace.ActionValues(),
		"servers":            carapace.ActionValues(),
		"servers-memory":     carapace.ActionValues(),
		"subnet":             carapace.ActionValues(),
		"timeout":            carapace.ActionValues(),
		"token":              carapace.ActionValues(),
		"volume":             carapace.ActionValues(),
	})
}

func x() {
	// KEY[=VALUE][@NODEFILTER[;NODEFILTER...]]
	carapace.ActionMultiPartsN("@", 2, func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return carapace.ActionMultiPartsN("=", 2, func(c carapace.Context) carapace.Action {
				switch len(c.Parts) {
				case 0:
					return carapace.ActionValues("TODO:key")
				default:
					return carapace.ActionValues("TODO:value")
				}
			})
		default:
			return carapace.ActionValues("TODO:nodefilter").List(";")
		}
	}).List(",")
}
