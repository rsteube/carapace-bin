package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kompose",
	Short: "A tool helping Docker Compose users move to Kubernetes",
	Long:  "https://kompose.io/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().Bool("error-on-warning", false, "Treat any warning as an error")
	rootCmd.PersistentFlags().StringArrayP("file", "f", []string{}, "Specify an alternative compose file")
	rootCmd.Flags().BoolP("help", "h", false, "help for kompose")
	rootCmd.PersistentFlags().String("provider", "kubernetes", "Specify a provider. Kubernetes or OpenShift.")
	rootCmd.PersistentFlags().Bool("suppress-warnings", false, "Suppress all warnings")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose output")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"file":     carapace.ActionFiles(),
		"provider": carapace.ActionValues("Kubernetes", "OpenShift"),
	})
}
