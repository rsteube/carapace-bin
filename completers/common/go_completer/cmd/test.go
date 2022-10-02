package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/go_completer/cmd/action"
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "test packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(testCmd).Standalone()

	testCmd.Flags().BoolS("args", "args", false, "pass the remainder of the command line to the test binary")
	testCmd.Flags().StringS("bench", "bench", "", "run only those benchmarks matching a regular expression")
	testCmd.Flags().StringS("benchtime", "benchtime", "", "Run enough iterations of each benchmark to take given duration")
	testCmd.Flags().BoolS("c", "c", false, "compile the test binary to pkg.test but do not run it")
	testCmd.Flags().StringS("count", "count", "", "run each test and benchmark n times")
	testCmd.Flags().BoolS("cover", "cover", false, "enable coverage analysis")
	testCmd.Flags().StringS("covermode", "covermode", "", "set the mode for coverage analysis for the package")
	testCmd.Flags().StringS("coverpkg", "coverpkg", "", "apply coverage analysis in each test to packages matching the patterns")
	testCmd.Flags().BoolS("cpu", "cpu", false, "specify a list of GOMAXPROCS values for which the tests or benchmarks should be executed")
	testCmd.Flags().StringS("exec", "exec", "", "run the test binary using xprog")
	testCmd.Flags().BoolS("failfast", "failfast", false, "Do not start new tests after the first test failure")
	testCmd.Flags().BoolS("i", "i", false, "install packages that are dependencies of the test")
	testCmd.Flags().BoolS("json", "json", false, "convert test output to JSON")
	testCmd.Flags().StringS("list", "list", "", "list tests, benchmarks, or examples matching the regular expression")
	testCmd.Flags().StringS("o", "o", "", "compile the test binary to the named file")
	testCmd.Flags().StringS("parallel", "parallel", "", "Allow parallel execution of test functions that call t.Paralle")
	testCmd.Flags().StringS("run", "run", "", "run only those tests and examples matching the regular expression")
	testCmd.Flags().BoolS("short", "short", false, "tell long-running tests to shorten their run time")
	testCmd.Flags().StringS("timeout", "timeout", "", "if a test binary runs longer than duration d, panic")
	testCmd.Flags().BoolS("v", "v", false, "verbose output")
	testCmd.Flags().StringS("vet", "vet", "", "configure the invocation of \"go vet\" during \"go test\" to use the comma-separated list of vet check")
	rootCmd.AddCommand(testCmd)

	carapace.Gen(testCmd).FlagCompletion(carapace.ActionMap{
		"bench": carapace.ActionMultiParts("|", func(c carapace.Context) carapace.Action {
			return action.ActionTests(c.Args, action.TestOpts{Benchmark: true}).Invoke(c).Filter(c.Parts).ToA()
		}),
		"covermode": carapace.ActionValues("set", "count,atomic"),
		"run": carapace.ActionMultiParts("|", func(c carapace.Context) carapace.Action {
			return action.ActionTests(c.Args, action.TestOpts{Example: true, Test: true}).Invoke(c).Filter(c.Parts).ToA()
		}),
	})

}
