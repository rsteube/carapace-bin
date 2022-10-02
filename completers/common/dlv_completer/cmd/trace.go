package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/ps"
	"github.com/spf13/cobra"
)

var traceCmd = &cobra.Command{
	Use:   "trace",
	Short: "Compile and begin tracing program.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(traceCmd).Standalone()
	traceCmd.Flags().Bool("ebpf", false, "Trace using eBPF (experimental).")
	traceCmd.Flags().StringP("exec", "e", "", "Binary file to exec and trace.")
	traceCmd.Flags().String("output", "debug", "Output path for the binary.")
	traceCmd.Flags().IntP("pid", "p", 0, "Pid to attach to.")
	traceCmd.Flags().IntP("stack", "s", 0, "Show stack trace with given depth. (Ignored with -ebpf)")
	traceCmd.Flags().BoolP("test", "t", false, "Trace a test binary.")
	rootCmd.AddCommand(traceCmd)

	carapace.Gen(traceCmd).FlagCompletion(carapace.ActionMap{
		"exec":   carapace.ActionFiles(),
		"output": carapace.ActionFiles(),
		"pid":    ps.ActionProcessIds(),
	})
}
