package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/pacman"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build packages from AUR",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(buildCmd).Standalone()

	buildCmd.Flags().String("builddir", "", "build directory")
	buildCmd.Flags().BoolP("dry-run", "d", false, "only print what would be done")
	buildCmd.Flags().BoolP("keep", "k", false, "keep built packages in cache after installation")
	buildCmd.Flags().Bool("no-clone", false, "do not clone build files from AUR, only use local files")
	buildCmd.Flags().Bool("no-confirm", false, "bypass any and all confirmation messages")
	buildCmd.Flags().Bool("no-keep", false, "do not keep built packages in cache after installation")
	rootCmd.AddCommand(buildCmd)

	carapace.Gen(buildCmd).FlagCompletion(carapace.ActionMap{
		"builddir": carapace.ActionDirectories(),
	})

	carapace.Gen(buildCmd).PositionalAnyCompletion(carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return pacman.ActionPackages(pacman.PackageOption{}).Invoke(c).Filter(c.Args).ToA()
	}))
}
