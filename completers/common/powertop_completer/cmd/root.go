package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "powertop",
	Short: "The Linux PowerTOP tool",
	Long:  "https://github.com/fenrus75/powertop",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("auto-tune", false, "sets all tunable options to their GOOD setting")
	rootCmd.Flags().BoolP("calibrate", "c", false, "runs powertop in calibration mode")
	rootCmd.Flags().StringP("csv", "C", "", "generate a csv report")
	rootCmd.Flags().Bool("debug", false, "run in \"debug\" mode")
	rootCmd.Flags().String("extech", "", "uses an Extech Power Analyzer for measurements")
	rootCmd.Flags().BoolP("help", "h", false, "print this help menu")
	rootCmd.Flags().StringP("html", "r", "", "generate a html report")
	rootCmd.Flags().StringP("iteration", "i", "", "number of times to run each test")
	rootCmd.Flags().BoolP("quiet", "q", false, "suppress stderr output")
	rootCmd.Flags().StringP("sample", "s", "", "interval for power consumption measurement")
	rootCmd.Flags().StringP("time", "t", "", "generate a report for 'x' seconds")
	rootCmd.Flags().BoolP("version", "V", false, "print version information")
	rootCmd.Flags().StringP("workload", "w", "", "file to execute for workload")

	rootCmd.Flag("csv").NoOptDefVal = " "
	rootCmd.Flag("extech").NoOptDefVal = " "
	rootCmd.Flag("html").NoOptDefVal = " "
	rootCmd.Flag("iteration").NoOptDefVal = " "
	rootCmd.Flag("sample").NoOptDefVal = " "
	rootCmd.Flag("time").NoOptDefVal = " "
	rootCmd.Flag("workload").NoOptDefVal = " "

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"csv":      carapace.ActionFiles(),
		"html":     carapace.ActionFiles(),
		"workload": carapace.ActionFiles(),
	})
}
