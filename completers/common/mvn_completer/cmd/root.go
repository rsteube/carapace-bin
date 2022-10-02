package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/mvn_completer/cmd/action"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mvn",
	Short: "Apache Maven is a software project management and comprehension tool",
	Long:  "https://maven.apache.org/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("activate-profiles", "P", "", "Comma-delimited list of profiles")
	rootCmd.Flags().Bool("also-make", false, "If project list is specified, also")
	rootCmd.Flags().Bool("also-make-dependents", false, "If project list is specified, also")
	rootCmd.Flags().BoolP("batch-mode", "B", false, "Run in non-interactive (batch)")
	rootCmd.Flags().StringP("builder", "b", "", "The id of the build strategy to use`")
	rootCmd.Flags().Bool("check-plugin-updates", false, "Ineffective, only kept for")
	rootCmd.Flags().BoolP("debug", "X", false, "Produce execution debug output")
	rootCmd.Flags().StringP("define", "D", "", "Define a system property")
	rootCmd.Flags().String("encrypt-master-password", "", "Encrypt master security password")
	rootCmd.Flags().String("encrypt-password", "", "Encrypt server password")
	rootCmd.Flags().BoolP("errors", "e", false, "Produce execution error messages")
	rootCmd.Flags().Bool("fail-at-end", false, "Only fail the build afterwards;")
	rootCmd.Flags().Bool("fail-fast", false, "Stop at first failure in")
	rootCmd.Flags().Bool("fail-never", false, "NEVER fail the build, regardless")
	rootCmd.Flags().StringP("file", "f", "", "Force the use of an alternate POM")
	rootCmd.Flags().String("global-settings", "", "Alternate path for the global")
	rootCmd.Flags().String("global-toolchains", "", "Alternate path for the global")
	rootCmd.Flags().BoolP("help", "h", false, "Display help information")
	rootCmd.Flags().BoolP("lax-checksums", "c", false, "Warn if checksums don't match")
	rootCmd.Flags().Bool("legacy-local-repository", false, "Use Maven 2 Legacy Local")
	rootCmd.Flags().StringP("log-file", "l", "", "Log file where all build output")
	rootCmd.Flags().Bool("no-plugin-registry", false, "Ineffective, only kept for")
	rootCmd.Flags().Bool("no-plugin-updates", false, "Ineffective, only kept for")
	rootCmd.Flags().Bool("no-snapshot-updates", false, "Suppress SNAPSHOT updates")
	rootCmd.Flags().Bool("no-transfer-progress", false, "Do not display transfer progress")
	rootCmd.Flags().BoolP("non-recursive", "N", false, "Do not recurse into sub-projects")
	rootCmd.Flags().BoolP("offline", "o", false, "Work offline")
	rootCmd.Flags().String("projects", "", "Comma-delimited list of specified")
	rootCmd.Flags().BoolP("quiet", "q", false, "Quiet output - only show errors")
	rootCmd.Flags().String("resume-from", "", "Resume reactor from specified")
	rootCmd.Flags().StringP("settings", "s", "", "Alternate path for the user")
	rootCmd.Flags().BoolP("show-version", "V", false, "Display version information")
	rootCmd.Flags().BoolP("strict-checksums", "C", false, "Fail the build if checksums don't")
	rootCmd.Flags().StringP("threads", "T", "", "Thread count, for instance 2.0C")
	rootCmd.Flags().StringP("toolchains", "t", "", "Alternate path for the user")
	rootCmd.Flags().Bool("update-plugins", false, "Ineffective, only kept for")
	rootCmd.Flags().BoolP("update-snapshots", "U", false, "Forces a check for missing")
	rootCmd.Flags().BoolP("version", "v", false, "Display version information")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"activate-profiles": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionProfiles(rootCmd.Flag("file").Value.String()).Invoke(c).Filter(c.Parts).ToA()
		}),
		"file":              carapace.ActionFiles(".xml"),
		"global-settings":   carapace.ActionFiles(".xml"),
		"global-toolchains": carapace.ActionFiles(".xml"),
		"log-file":          carapace.ActionFiles(),
		"settings":          carapace.ActionFiles(".xml"),
		"toolchains":        carapace.ActionFiles(".xml"),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionGoalsAndPhases(rootCmd.Flag("file").Value.String()).Invoke(c).Filter(c.Args).ToA()
		}),
	)
}
