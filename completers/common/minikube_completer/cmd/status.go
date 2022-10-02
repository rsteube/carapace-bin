package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/minikube_completer/cmd/action"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Gets the status of a local Kubernetes cluster",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	statusCmd.Flags().StringP("format", "f", `{{.Name}}
type: Control Plane
host: {{.Host}}
kubelet: {{.Kubelet}}
apiserver: {{.APIServer}}
kubeconfig: {{.Kubeconfig}}
{{- if .TimeToStop }}
timeToStop: {{.TimeToStop}}
{{- end }}
{{- if .DockerEnv }}
docker-env: {{.DockerEnv}}
{{- end }}
{{- if .PodManEnv }}
podman-env: {{.PodManEnv}}
{{- end }}

`, "Go template format string for the status output.")
	statusCmd.Flags().StringP("layout", "l", "nodes", "output layout (EXPERIMENTAL, JSON only): 'nodes' or 'cluster'")
	statusCmd.Flags().StringP("node", "n", "", "The node to check status for. Defaults to control plane. Leave blank with default format for status on all nodes.")
	statusCmd.Flags().StringP("output", "o", "text", "minikube status --output OUTPUT. json, text")
	statusCmd.Flags().StringP("watch", "w", "1s", "Continuously listing/getting the status with optional interval duration.")
	statusCmd.Flag("watch").NoOptDefVal = "1s"
	rootCmd.AddCommand(statusCmd)

	carapace.Gen(statusCmd).FlagCompletion(carapace.ActionMap{
		"node":   action.ActionNodes(),
		"output": carapace.ActionValues("json", "text"),
	})
}
