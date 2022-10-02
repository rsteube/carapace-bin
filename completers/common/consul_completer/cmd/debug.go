package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var debugCmd = &cobra.Command{
	Use:   "debug",
	Short: "Records a debugging archive for operators",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugCmd).Standalone()

	debugCmd.Flags().Bool("archive", false, "Boolean value for if the files should be archived and compressed.")
	debugCmd.Flags().StringArray("capture", []string{}, "One or more types of information to capture.")
	debugCmd.Flags().String("duration", "", "The total time to record information.")
	debugCmd.Flags().String("interval", "", "The interval in which to capture dynamic information.")
	debugCmd.Flags().String("output", "", "The path to the compressed archive that will be created with the information after collection.")
	rootCmd.AddCommand(debugCmd)

	carapace.Gen(debugCmd).FlagCompletion(carapace.ActionMap{
		"capture": carapace.ActionValuesDescribed(
			"agent", "Version and configuration information about the agent.",
			"cluster", "A list of all the WAN and LAN members in the cluster.",
			"host", "Information about resources on the host running the target agent such as CPU, memory, and disk.",
			"logs", "DEBUG level logs for the target agent, captured for the duration.",
			"metrics", "Metrics from the in-memory metrics endpoint in the target, captured at the interval.",
			"pprof", "Golang heap, CPU, goroutine, and trace profiling. ",
		),
		"output": carapace.ActionFiles(),
	})
}
