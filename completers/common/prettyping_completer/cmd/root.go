package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "prettyping",
	Short: "This script is a wrapper around the system's \"ping\" tool",
	Long:  "https://github.com/denilsonsa/prettyping",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("6", "6", false, "Shortcut for: --pingbin ping6")
	rootCmd.Flags().BoolS("R", "R", false, "Record route mode is not allowed in prettyping.")
	rootCmd.Flags().BoolS("a", "a", false, "Audible ping is not implemented yet.")
	rootCmd.Flags().String("awkbin", "", "Override the awk interpreter.")
	rootCmd.Flags().Bool("color", false, "Enable color output.")
	rootCmd.Flags().String("columns", "", "Override auto-detection of terminal dimensions.")
	rootCmd.Flags().BoolS("f", "f", false, "Flood mode is not allowed in prettyping.")
	rootCmd.Flags().String("last", "", "Use the last \"n\" pings at the statistics line.")
	rootCmd.Flags().Bool("legend", false, "Enable/disable the latency legend.")
	rootCmd.Flags().String("lines", "", "Override auto-detection of terminal dimensions.")
	rootCmd.Flags().Bool("multicolor", false, "Enable multi-color unicode output.")
	rootCmd.Flags().Bool("nocolor", false, "Disable color output.")
	rootCmd.Flags().Bool("nolegend", false, "Enable/disable the latency legend.")
	rootCmd.Flags().Bool("nomulticolor", false, "Disable multi-color unicode output.")
	rootCmd.Flags().Bool("noterminal", false, "Force the output designed to a terminal.")
	rootCmd.Flags().Bool("nounicode", false, "Disable unicode characters.")
	rootCmd.Flags().String("pingbin", "", "Override the ping tool.")
	rootCmd.Flags().BoolS("q", "q", false, "Quiet output is not allowed in prettyping.")
	rootCmd.Flags().String("rttmax", "", "Maximum RTT represented in the unicode graph.")
	rootCmd.Flags().String("rttmin", "", "Minimum RTT represented in the unicode graph.")
	rootCmd.Flags().Bool("terminal", false, "Force the output designed to a terminal.")
	rootCmd.Flags().Bool("unicode", false, "Enable unicode characters.")
	rootCmd.Flags().BoolS("v", "v", false, "Verbose output seems to be the default mode in ping.")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"awkbin":  carapace.ActionFiles(),
		"pingbin": carapace.ActionFiles(),
	})
}
