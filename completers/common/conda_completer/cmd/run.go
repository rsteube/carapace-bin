package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/conda_completer/cmd/action"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run an executable in a conda environment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runCmd).Standalone()

	runCmd.Flags().String("cwd", "", "Current working directory for command to run in.")
	runCmd.Flags().Bool("debug-wrapper-scripts", false, "When this is set, where implemented, the shell wrapper scriptswill echo to stderr a lot of debugging information.")
	runCmd.Flags().Bool("dev", false, "Sets `CONDA_EXE` to `python -m conda`.")
	runCmd.Flags().BoolP("help", "h", false, "Show this help message and exit.")
	runCmd.Flags().Bool("live-stream", false, "Display the output for the subprocess stdout and stderr on real time.")
	runCmd.Flags().StringP("name", "n", "", "Name of environment.")
	runCmd.Flags().Bool("no-capture-output", false, "Don't capture stdout/stderr")
	runCmd.Flags().StringP("prefix", "p", "", "Full path to environment location (i.e. prefix).")
	runCmd.Flags().BoolP("verbose", "v", false, "Use once for info, twice for debug, three times for trace.")
	rootCmd.AddCommand(runCmd)

	carapace.Gen(runCmd).FlagCompletion(carapace.ActionMap{
		"name":   action.ActionEnvironments(),
		"prefix": carapace.ActionDirectories(),
	})

	carapace.Gen(runCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
