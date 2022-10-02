package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "helmsman",
	Short: "Helmsman is a Helm Charts as Code tool",
	Long:  "https://github.com/Praqma/Helmsman",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringArrayS("e", "e", []string{}, "file(s) to load environment variables from (default .env), may be supplied more than once")
	rootCmd.Flags().StringArrayS("f", "f", []string{}, "desired state file name(s), may be supplied more than once to merge state files")
	rootCmd.Flags().String("group", "", "limit execution to specific group of apps.")
	rootCmd.Flags().String("target", "", "limit execution to specific app.")
	rootCmd.Flags().Int("diff-context", -1, "number of lines of context to show around changes in helm diff output")
	rootCmd.Flags().IntS("p", "p", 1, "max number of concurrent helm releases to run")
	rootCmd.Flags().Bool("always-upgrade", false, "upgrade release even if no changes are found")
	rootCmd.Flags().Bool("apply", false, "apply the plan directly")
	rootCmd.Flags().String("context-override", "", "override releases context defined in release state with this one")
	rootCmd.Flags().Bool("debug", false, "show the debug execution logs and actual helm/kubectl commands. This can log secrets and should only be used for debugging purposes.")
	rootCmd.Flags().Bool("destroy", false, "delete all deployed releases.")
	rootCmd.Flags().Bool("dry-run", false, "apply the dry-run option for helm commands.")
	rootCmd.Flags().Bool("force-upgrades", false, "use --force when upgrading helm releases. May cause resources to be recreated.")
	rootCmd.Flags().Bool("keep-untracked-releases", false, "keep releases that are managed by Helmsman from the used DSFs in the command, and are no longer tracked in your desired state.")
	rootCmd.Flags().String("kubeconfig", "", "path to the kubeconfig file to use for CLI requests")
	rootCmd.Flags().Bool("kubectl-diff", false, "use kubectl diff instead of helm diff. Defalts to false if the helm diff plugin is installed.")
	rootCmd.Flags().Bool("migrate-context", false, "updates the context name for all apps defined in the DSF and applies Helmsman labels. Using this flag is required if you want to change context name after it has been set.")
	rootCmd.Flags().Bool("no-banner", false, "don't show the banner")
	rootCmd.Flags().Bool("no-cleanup", false, "keeps any credentials files that has been downloaded on the host where helmsman runs.")
	rootCmd.Flags().Bool("no-color", false, "don't use colors")
	rootCmd.Flags().Bool("no-env-subst", false, "turn off environment substitution globally")
	rootCmd.Flags().Bool("no-fancy", false, "don't display the banner and don't use colors")
	rootCmd.Flags().Bool("no-ns", false, "don't create namespaces")
	rootCmd.Flags().Bool("no-ssm-subst", false, "turn off SSM parameter substitution globally")
	rootCmd.Flags().Bool("no-update", false, "skip updating helm repos")
	rootCmd.Flags().String("ns-override", "", "override defined namespaces with this one")
	rootCmd.Flags().Bool("replace-on-rename", false, "Uninstall the existing release when a chart with a different name is used.")
	rootCmd.Flags().Bool("show-diff", false, "show helm diff results. Can expose sensitive information.")
	rootCmd.Flags().Bool("skip-validation", false, "skip desired state validation")
	rootCmd.Flags().Bool("subst-env-values", false, "turn on environment substitution in values files.")
	rootCmd.Flags().Bool("subst-ssm-values", false, "turn on SSM parameter substitution in values files.")
	rootCmd.Flags().Bool("update-deps", false, "run 'helm dep up' for local charts")
	rootCmd.Flags().BoolS("v", "v", false, "show the version")
	rootCmd.Flags().Bool("verbose", false, "show verbose execution logs.")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"f":          carapace.ActionFiles(),
		"kubeconfig": carapace.ActionFiles(),
	})
}
