package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/mount_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "umount",
	Short: "Unmount filesystems",
	Long:  "https://linux.die.net/man/8/umount",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("all", "a", false, "unmount all filesystems")
	rootCmd.Flags().BoolP("all-targets", "A", false, "unmount all mountpoints for the given device in the current namespace")
	rootCmd.Flags().BoolP("detach-loop", "d", false, "if mounted loop device, also free this loop device")
	rootCmd.Flags().Bool("fake", false, "dry run; skip the umount(2) syscall")
	rootCmd.Flags().BoolP("force", "f", false, "force unmount (in case of an unreachable NFS system)")
	rootCmd.Flags().BoolP("help", "h", false, "display this help")
	rootCmd.Flags().BoolP("internal-only", "i", false, "don't call the umount.<type> helpers")
	rootCmd.Flags().BoolP("lazy", "l", false, "detach the filesystem now, clean up things later")
	rootCmd.Flags().StringP("namespace", "N", "", "perform umount in another namespace")
	rootCmd.Flags().BoolP("no-canonicalize", "c", false, "don't canonicalize paths")
	rootCmd.Flags().BoolP("no-mtab", "n", false, "don't write to /etc/mtab")
	rootCmd.Flags().BoolP("quiet", "q", false, "suppress 'not mounted' error messages")
	rootCmd.Flags().BoolP("read-only", "r", false, "in case unmounting fails, try to remount read-only")
	rootCmd.Flags().BoolP("recursive", "R", false, "recursively unmount a target with all its children")
	rootCmd.Flags().StringP("test-opts", "O", "", "limit the set of filesystems (use with -a)")
	rootCmd.Flags().StringP("types", "t", "", "limit the set of filesystem types")
	rootCmd.Flags().BoolP("verbose", "v", false, "say what is being done")
	rootCmd.Flags().BoolP("version", "V", false, "display version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"test-opts": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionOptions().Invoke(c).Filter(c.Parts).ToA()
		}),
		"types": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return fs.ActionFilesystemTypes().Invoke(c).Filter(c.Parts).ToA()
		}),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if rootCmd.Flag("help").Changed ||
				rootCmd.Flag("version").Changed ||
				rootCmd.Flag("all").Changed {
				return carapace.ActionValues()
			}
			return carapace.Batch(
				action.ActionSources(),
				fs.ActionMounts(),
			).ToA()
		}),
	)
}
