package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var vendorCmd = &cobra.Command{
	Use:   "vendor",
	Short: "Vendor remote modules into a local directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vendorCmd).Standalone()

	vendorCmd.Flags().String("cert", "", "Load certificate authority from PEM encoded file")
	vendorCmd.Flags().StringP("config", "c", "", "Specify the configuration file")
	vendorCmd.Flags().BoolP("force", "f", false, "Forcefully overwrite conflicting files in existing output directory")
	vendorCmd.Flags().BoolP("help", "h", false, "Print help information")
	vendorCmd.Flags().String("import-map", "", "Load import map file")
	vendorCmd.Flags().String("lock", "", "Check the specified lock file")
	vendorCmd.Flags().Bool("no-config", false, "Disable automatic loading of the configuration file.")
	vendorCmd.Flags().String("output", "", "The directory to output the vendored modules to")
	vendorCmd.Flags().BoolP("quiet", "q", false, "Suppress diagnostic output")
	vendorCmd.Flags().StringP("reload", "r", "", "Reload source code cache (recompile TypeScript)")
	vendorCmd.Flags().Bool("unstable", false, "Enable unstable features and APIs")
	rootCmd.AddCommand(vendorCmd)

	vendorCmd.Flag("reload").NoOptDefVal = " "

	carapace.Gen(vendorCmd).FlagCompletion(carapace.ActionMap{
		"cert":       carapace.ActionFiles(),
		"config":     carapace.ActionFiles(),
		"import-map": carapace.ActionFiles(),
		"lock":       carapace.ActionFiles(),
		"output":     carapace.ActionDirectories(),
	})

	// TODO positional completion
}
