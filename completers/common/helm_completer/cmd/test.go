package cmd

import (
	"time"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/helm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "run tests for a release",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	testCmd.Flags().StringSlice("filter", []string{}, "specify tests by attribute (currently \"name\") using attribute=value syntax or '!attribute=value' to exclude a test (can specify multiple or separate values with commas: name=test1,name=test2)")
	testCmd.Flags().Bool("logs", false, "dump the logs from test pods (this runs after all tests are complete, but before any cleanup)")
	testCmd.Flags().Duration("timeout", 5*time.Minute, "time to wait for any individual Kubernetes operation (like Jobs for hooks)")
	rootCmd.AddCommand(testCmd)

	carapace.Gen(testCmd).PositionalCompletion(
		action.ActionReleases(),
	)
}
