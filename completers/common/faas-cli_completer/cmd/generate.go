package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/faas-cli_completer/cmd/action"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate Kubernetes CRD YAML file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	generateCmd.Flags().StringArray("annotation", []string{}, "Any annotations you want to add (to store functions only)")
	generateCmd.Flags().String("api", "openfaas.com/v1", "CRD API version e.g openfaas.com/v1, serving.knative.dev/v1")
	generateCmd.Flags().String("arch", "x86_64", "Desired image arch. (Default x86_64)")
	generateCmd.Flags().Bool("envsubst", true, "Substitute environment variables in stack.yml file")
	generateCmd.Flags().String("from-store", "", "generate using a store image")
	generateCmd.Flags().StringP("namespace", "n", "", "Kubernetes namespace for functions")
	generateCmd.Flags().String("tag", "latest", "Override latest tag on function Docker image, accepts 'latest', 'sha', 'branch', 'describe'")
	rootCmd.AddCommand(generateCmd)

	carapace.Gen(generateCmd).FlagCompletion(carapace.ActionMap{
		"arch":      action.ActionImageArchitectures(),
		"namespace": action.ActionNamespaces(),
		"tag":       carapace.ActionValues("latest", "sha", "branch", "describe"),
	})
}
