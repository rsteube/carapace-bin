package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/faas-cli_completer/cmd/action"
	"github.com/spf13/cobra"
)

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy OpenFaaS functions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	deployCmd.Flags().StringArray("annotation", []string{}, "Set one or more annotation (ANNOTATION=VALUE)")
	deployCmd.Flags().StringArray("constraint", []string{}, "Apply a constraint to the function")
	deployCmd.Flags().StringArrayP("env", "e", []string{}, "Set one or more environment variables (ENVVAR=VALUE)")
	deployCmd.Flags().Bool("envsubst", true, "Substitute environment variables in stack.yml file")
	deployCmd.Flags().String("fprocess", "", "fprocess value to be run as a serverless function by the watchdog")
	deployCmd.Flags().StringP("gateway", "g", "http://127.0.0.1:8080", "Gateway URL starting with http(s)://")
	deployCmd.Flags().String("handler", "", "Directory with handler for function, e.g. handler.js")
	deployCmd.Flags().String("image", "", "Docker image name to build")
	deployCmd.Flags().StringArrayP("label", "l", []string{}, "Set one or more label (LABEL=VALUE)")
	deployCmd.Flags().String("lang", "", "Programming language template")
	deployCmd.Flags().String("name", "", "Name of the deployed function")
	deployCmd.Flags().StringP("namespace", "n", "", "Namespace of the function")
	deployCmd.Flags().String("network", "", "Name of the network")
	deployCmd.Flags().Bool("read-template", true, "Read the function's template")
	deployCmd.Flags().Bool("readonly", false, "Force the root container filesystem to be read only")
	deployCmd.Flags().Bool("replace", false, "Remove and re-create existing function(s)")
	deployCmd.Flags().StringArray("secret", []string{}, "Give the function access to a secure secret")
	deployCmd.Flags().String("tag", "latest", "Override latest tag on function Docker image, accepts 'latest', 'sha', 'branch', or 'describe'")
	deployCmd.Flags().Bool("tls-no-verify", false, "Disable TLS validation")
	deployCmd.Flags().StringP("token", "k", "", "Pass a JWT token to use instead of basic auth")
	deployCmd.Flags().Bool("update", true, "Perform rolling update on existing function(s)")
	rootCmd.AddCommand(deployCmd)

	carapace.Gen(deployCmd).FlagCompletion(carapace.ActionMap{
		"handler":   carapace.ActionDirectories(),
		"lang":      action.ActionLanguageTemplates(),
		"namespace": action.ActionNamespaces(),
		"tag":       carapace.ActionValues("latest", "sha", "branch", "describe"),
	})
}
