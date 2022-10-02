package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/adb"
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Run Flutter unit tests for the current project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(testCmd).Standalone()

	testCmd.Flags().StringP("concurrency", "j", "", "The number of concurrent test processes to run.")
	testCmd.Flags().Bool("coverage", false, "Whether to collect coverage information.")
	testCmd.Flags().String("coverage-path", "", "Where to store coverage information (if coverage is enabled).")
	testCmd.Flags().StringArray("dart-define", []string{}, "Additional key-value pairs that will be available as constants.")
	testCmd.Flags().String("dds-port", "", "Dart Development Service (DDS) will be bound to the provided port.")
	testCmd.Flags().String("device-user", "", "Identifier number for a user or work profile on Android only.")
	testCmd.Flags().BoolP("exclude-tags", "x", false, "Run only tests that do not have the specified tags.")
	testCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	testCmd.Flags().Bool("merge-coverage", false, "Whether to merge coverage data with \"coverage/lcov.base.info\".")
	testCmd.Flags().String("name", "", "A regular expression matching substrings of the names of tests to run.")
	testCmd.Flags().Bool("no-null-assertions", false, "Do not perform additional null assertions on the boundaries of migrated and un-migrated code.")
	testCmd.Flags().Bool("no-pub", false, "Do not run \"flutter pub get\" before executing this command.")
	testCmd.Flags().Bool("no-run-skipped", false, "Do not run skipped tests instead of skipping them.")
	testCmd.Flags().Bool("no-test-assets", false, "Do not build the assets bundle for testing.")
	testCmd.Flags().Bool("no-track-widget-creation", false, "Do not track widget creation locations.")
	testCmd.Flags().Bool("null-assertions", false, "Perform additional null assertions on the boundaries of migrated and un-migrated code.")
	testCmd.Flags().String("plain-name", "", "A plain-text substring of the names of tests to run.")
	testCmd.Flags().Bool("pub", false, "Run \"flutter pub get\" before executing this command.")
	testCmd.Flags().StringP("reporter", "r", "", "Set how to print test results.")
	testCmd.Flags().Bool("run-skipped", false, "Run skipped tests instead of skipping them.")
	testCmd.Flags().String("shard-index", "", "Tests can be sharded with the \"--total-shards\" and \"--shard-index\" arguments.")
	testCmd.Flags().Bool("start-paused", false, "Start in a paused mode and wait for a debugger to connect.")
	testCmd.Flags().BoolP("tags", "t", false, "Run only tests associated with the specified tags.")
	testCmd.Flags().Bool("test-assets", false, "Build the assets bundle for testing.")
	testCmd.Flags().Bool("test-randomize-ordering-seed", false, "The seed to randomize the execution order of test cases within test files.")
	testCmd.Flags().String("timeout", "", "The default test timeout, specified either in seconds.")
	testCmd.Flags().String("total-shards", "", "Tests can be sharded with the \"--total-shards\" and \"--shard-index\" arguments.")
	testCmd.Flags().Bool("track-widget-creation", false, "Track widget creation locations.")
	testCmd.Flags().Bool("update-goldens", false, "Whether \"matchesGoldenFile()\" calls within your test methods should update the golden files rather than test for an existing match.")
	testCmd.Flags().Bool("web-renderer", false, "The renderer implementation to use when building for the web.")
	rootCmd.AddCommand(testCmd)

	carapace.Gen(testCmd).FlagCompletion(carapace.ActionMap{
		// TODO plain-name
		"coverage-path": carapace.ActionFiles(),
		"device-user":   adb.ActionUsers(),
		"reporter": carapace.ActionValuesDescribed(
			"compact", "A single line that updates dynamically.",
			"expanded", "A separate line for each update.",
			"json", "A machine-readable format.",
		),
	})
}
