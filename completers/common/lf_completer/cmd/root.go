package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "lf",
	Short: "terminal file manager",
	Long:  "https://github.com/gokcehan/lf",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("command", "", "command to execute on client initialization")
	rootCmd.Flags().String("config", "", "path to the config file")
	rootCmd.Flags().String("cpuprofile", "", "path to the file to write the CPU profile")
	rootCmd.Flags().Bool("doc", false, "show documentation")
	rootCmd.Flags().String("last-dir-path", "", "path to the file to write the last dir on exit (to use for cd)")
	rootCmd.Flags().String("memprofile", "", "path to the file to write the memory profile")
	rootCmd.Flags().String("remote", "", "send remote command to server")
	rootCmd.Flags().String("selection-path", "", "path to the file to write selected files on open")
	rootCmd.Flags().Bool("server", false, "start server (automatic)")
	rootCmd.Flags().Bool("single", false, "start a client without server")
	rootCmd.Flags().Bool("version", false, "show version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"config":         carapace.ActionFiles(),
		"cpuprofile":     carapace.ActionFiles(),
		"last-dir-path":  carapace.ActionFiles(),
		"memprofile":     carapace.ActionFiles(),
		"selection-path": carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
