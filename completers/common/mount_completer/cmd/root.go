package cmd

import (
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/mount_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mount",
	Short: "mount a filesystem",
	Long:  "https://linux.die.net/man/8/mount",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("all", "a", false, "mount all filesystems mentioned in fstab")
	rootCmd.Flags().BoolP("bind", "B", false, "mount a subtree somewhere else (same as -o bind)")
	rootCmd.Flags().BoolP("fake", "f", false, "dry run; skip the mount(2) syscall")
	rootCmd.Flags().BoolP("fork", "F", false, "fork off for each device (use with -a)")
	rootCmd.Flags().StringP("fstab", "T", "", "alternative file to /etc/fstab")
	rootCmd.Flags().BoolP("help", "h", false, "display this help")
	rootCmd.Flags().BoolP("internal-only", "i", false, "don't call the mount.<type> helpers")
	rootCmd.Flags().StringP("label", "L", "", "synonym for LABEL=<label>")
	rootCmd.Flags().Bool("make-private", false, "mark a subtree as private")
	rootCmd.Flags().Bool("make-rprivate", false, "recursively mark a whole subtree as private")
	rootCmd.Flags().Bool("make-rshared", false, "recursively mark a whole subtree as shared")
	rootCmd.Flags().Bool("make-rslave", false, "recursively mark a whole subtree as slave")
	rootCmd.Flags().Bool("make-runbindable", false, "recursively mark a whole subtree as unbindable")
	rootCmd.Flags().Bool("make-shared", false, "mark a subtree as shared")
	rootCmd.Flags().Bool("make-slave", false, "mark a subtree as slave")
	rootCmd.Flags().Bool("make-unbindable", false, "mark a subtree as unbindable")
	rootCmd.Flags().BoolP("move", "M", false, "move a subtree to some other place")
	rootCmd.Flags().StringP("namespace", "N", "", "perform mount in another namespace")
	rootCmd.Flags().BoolP("no-canonicalize", "c", false, "don't canonicalize paths")
	rootCmd.Flags().BoolP("no-mtab", "n", false, "don't write to /etc/mtab")
	rootCmd.Flags().StringP("options", "o", "", "comma-separated list of mount options")
	rootCmd.Flags().String("options-mode", "", "what to do with options loaded from fstab")
	rootCmd.Flags().String("options-source", "", "mount options source")
	rootCmd.Flags().Bool("options-source-force", false, "force use of options from fstab/mtab")
	rootCmd.Flags().BoolP("rbind", "R", false, "mount a subtree and all submounts somewhere else")
	rootCmd.Flags().BoolP("read-only", "r", false, "mount the filesystem read-only (same as -o ro)")
	rootCmd.Flags().StringP("rw", "w", "", "mount the filesystem read-write (default)")
	rootCmd.Flags().BoolP("show-labels", "l", false, "show also filesystem labels")
	rootCmd.Flags().String("source", "", "explicitly specifies source (path, label, uuid)")
	rootCmd.Flags().String("target", "", "explicitly specifies mountpoint")
	rootCmd.Flags().String("target-prefix", "", "specifies path used for all mountpoints")
	rootCmd.Flags().StringP("test-opts", "O", "", "limit the set of filesystems (use with -a)")
	rootCmd.Flags().StringP("types", "t", "", "limit the set of filesystem types")
	rootCmd.Flags().StringP("uuid", "U", "", "synonym for UUID=<uuid>")
	rootCmd.Flags().BoolP("verbose", "v", false, "say what is being done")
	rootCmd.Flags().BoolP("version", "V", false, "display version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"fstab": carapace.ActionFiles(),
		"label": fs.ActionLabels(),
		"options": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			keys := make([]string, 0)
			for _, part := range c.Parts {
				keys = append(keys, strings.Split(part, "=")[0])
			}
			return action.ActionOptions().Invoke(c).Filter(keys).ToA()
		}),
		"options-mode": carapace.ActionValues("ignore", "append", "prepend", "replace"),
		"options-source": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("fstab", "mtab", "disable").Invoke(c).Filter(c.Parts).ToA()
		}),
		"source":        action.ActionSources(),
		"target":        carapace.ActionDirectories(),
		"target-prefix": carapace.ActionDirectories(),
		"types": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return fs.ActionFilesystemTypes().Invoke(c).Filter(c.Parts).ToA()
		}),
		"uuid": fs.ActionUuids(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if rootCmd.Flag("all").Changed {
				return carapace.ActionValues()
			}
			if isOperation(rootCmd) {
				return fs.ActionMounts()
			}
			if rootCmd.Flag("source").Changed {
				if rootCmd.Flag("target").Changed {
					return carapace.ActionValues()
				}
				return carapace.ActionDirectories()
			}
			return action.ActionSources()
		}),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if rootCmd.Flag("all").Changed {
				return carapace.ActionValues()
			}
			if isOperation(rootCmd) {
				if rootCmd.Flag("bind").Changed ||
					rootCmd.Flag("rbind").Changed ||
					rootCmd.Flag("move").Changed {
					return carapace.ActionDirectories()
				}
			}

			if !rootCmd.Flag("source").Changed &&
				!rootCmd.Flag("target").Changed {
				return carapace.ActionDirectories()
			}
			return carapace.ActionValues()
		}),
	)
}

func isOperation(cmd *cobra.Command) bool {
	return cmd.Flag("bind").Changed ||
		cmd.Flag("move").Changed ||
		cmd.Flag("rbind").Changed ||
		cmd.Flag("make-shared").Changed ||
		cmd.Flag("make-slave").Changed ||
		cmd.Flag("make-private").Changed ||
		cmd.Flag("make-unbindable").Changed ||
		cmd.Flag("make-rshared").Changed ||
		cmd.Flag("make-rslave").Changed ||
		cmd.Flag("make-rprivate").Changed
}
