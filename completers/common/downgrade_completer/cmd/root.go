package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/pacman"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "downgrade",
	Short: "Downgrade Arch Linux packages",
	Long:  "https://github.com/pbrisbin/downgrade",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("ala-only", false, "only use ALA server")
	rootCmd.Flags().String("ala-url", "", "location of ALA server")
	rootCmd.Flags().Bool("cached-only", false, "only use cached packages")
	rootCmd.Flags().BoolP("help", "h", false, "show help script")
	rootCmd.Flags().String("maxdepth", "", "maximum depth to search for cached packages")
	rootCmd.Flags().String("pacman", "", "pacman command to use")
	rootCmd.Flags().String("pacman-cache", "", "pacman cache directory")
	rootCmd.Flags().String("pacman-conf", "", "pacman configuration file")
	rootCmd.Flags().String("pacman-log", "", "pacman log file")
	rootCmd.Flags().Bool("version", false, "show downgrade version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"pacman-cache": carapace.ActionDirectories(),
		"pacman-conf":  carapace.ActionFiles(),
		"pacman-log":   carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return pacman.ActionPackages(pacman.PackageOption{Explicit: true}).Invoke(c).Filter(c.Args).ToA()
		}),
	)
}
