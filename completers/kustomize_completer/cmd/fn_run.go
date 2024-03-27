package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var fn_runCmd = &cobra.Command{
	Use:   "run",
	Short: "[Alpha] Reoncile config functions to Resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fn_runCmd).Standalone()
	fn_runCmd.Flags().Bool("as-current-user", false, "use the uid and gid of the command executor to run the function in the container")
	fn_runCmd.Flags().Bool("dry-run", false, "print results to stdout")
	fn_runCmd.Flags().Bool("enable-exec", false, "enable support for exec functions -- note: exec functions run arbitrary code -- do not use for untrusted configs!!! (Alpha)")
	fn_runCmd.Flags().Bool("enable-star", false, "enable support for starlark functions. (Alpha)")
	fn_runCmd.Flags().StringArrayP("env", "e", []string{}, "a list of environment variables to be used by functions")
	fn_runCmd.Flags().String("exec-path", "", "run an executable as a function. (Alpha)")
	fn_runCmd.Flags().StringSlice("fn-path", []string{}, "read functions from these directories instead of the configuration directory.")
	fn_runCmd.Flags().Bool("global-scope", false, "set global scope for functions.")
	fn_runCmd.Flags().String("image", "", "run this image as a function instead of discovering them.")
	fn_runCmd.Flags().Bool("include-subpackages", true, "also print resources from subpackages.")
	fn_runCmd.Flags().Bool("log-steps", false, "log steps to stderr")
	fn_runCmd.Flags().StringArray("mount", []string{}, "a list of storage options read from the filesystem")
	fn_runCmd.Flags().Bool("network", false, "enable network access for functions that declare it")
	fn_runCmd.Flags().String("results-dir", "", "write function results to this dir")
	fn_runCmd.Flags().String("star-name", "", "name of starlark program. (Alpha)")
	fn_runCmd.Flags().String("star-path", "", "run a starlark script as a function. (Alpha)")
	fn_runCmd.Flags().String("star-url", "", "run a starlark script as a function. (Alpha)")
	fnCmd.AddCommand(fn_runCmd)
}
