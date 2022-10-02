package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/faas-cli_completer/cmd/action"
	"github.com/spf13/cobra"
)

var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "Builds and pushes multi-arch OpenFaaS container images",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	publishCmd.Flags().StringArrayP("build-arg", "b", []string{}, "Add a build-arg for Docker (KEY=VALUE)")
	publishCmd.Flags().StringArray("build-label", []string{}, "Add a label for Docker image (LABEL=VALUE)")
	publishCmd.Flags().StringArrayP("build-option", "o", []string{}, "Set a build option, e.g. dev")
	publishCmd.Flags().StringArray("copy-extra", []string{}, "Extra paths that will be copied into the function build context")
	publishCmd.Flags().Bool("disable-stack-pull", false, "Disables the template configuration in the stack.yml")
	publishCmd.Flags().Bool("envsubst", true, "Substitute environment variables in stack.yml file")
	publishCmd.Flags().StringArray("extra-tag", []string{}, "Additional extra image tag")
	publishCmd.Flags().String("handler", "", "Directory with handler for function, e.g. handler.js")
	publishCmd.Flags().String("image", "", "Docker image name to build")
	publishCmd.Flags().String("lang", "", "Programming language template")
	publishCmd.Flags().String("name", "", "Name of the deployed function")
	publishCmd.Flags().Bool("no-cache", false, "Do not use Docker's build cache")
	publishCmd.Flags().Int("parallel", 1, "Build in parallel to depth specified.")
	publishCmd.Flags().String("platforms", "linux/amd64", "A set of platforms to publish")
	publishCmd.Flags().Bool("quiet", false, "Perform a quiet build, without showing output from Docker")
	publishCmd.Flags().Bool("shrinkwrap", false, "Just write files to ./build/ folder for shrink-wrapping")
	publishCmd.Flags().Bool("squash", false, "Use Docker's squash flag for smaller images [experimental] ")
	publishCmd.Flags().String("tag", "latest", "Override latest tag on function Docker image, accepts 'latest', 'sha', 'branch', or 'describe'")
	rootCmd.AddCommand(publishCmd)

	carapace.Gen(publishCmd).FlagCompletion(carapace.ActionMap{
		"lang":      action.ActionLanguageTemplates(),
		"name":      action.ActionFunctions(),
		"platforms": action.ActionImageArchitectures(), // TODO likely wrong as example is `linux/amd64`
		"tag":       carapace.ActionValues("latest", "sha", "branch", "describe"),
	})
}
