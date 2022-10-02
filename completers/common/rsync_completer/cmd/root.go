package cmd

import (
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/rsync_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/fs"
	"github.com/rsteube/carapace-bin/pkg/actions/net"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/rsteube/carapace-bin/pkg/util"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rsync",
	Short: "a fast, versatile, remote (and local) file-copying tool",
	Long:  "https://github.com/WayneD/rsync/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("8-bit-output", "8", false, "leave high-bit chars unescaped in output")
	rootCmd.Flags().BoolS("D", "D", false, "same as --devices --specials")
	rootCmd.Flags().BoolS("F", "F", false, "same as --filter='dir-merge /.rsync-filter'")
	rootCmd.Flags().BoolS("P", "P", false, "same as --partial --progress")
	rootCmd.Flags().BoolP("acls", "A", false, "preserve ACLs (implies --perms)")
	rootCmd.Flags().String("address", "", "bind address for outgoing socket to daemon")
	rootCmd.Flags().Bool("append", false, "append data onto shorter files")
	rootCmd.Flags().Bool("append-verify", false, "--append w/old data in file checksum")
	rootCmd.Flags().BoolP("archive", "a", false, "archive mode; equals -rlptgoD (no -H,-A,-X)")
	rootCmd.Flags().BoolP("atimes", "U", false, "preserve access (use) times")
	rootCmd.Flags().BoolP("backup", "b", false, "make backups (see --suffix & --backup-dir)")
	rootCmd.Flags().String("backup-dir", "", "make backups into hierarchy based in DIR")
	rootCmd.Flags().StringP("block-size", "B", "", "force a fixed checksum block-size")
	rootCmd.Flags().Bool("blocking-io", false, "use blocking I/O for the remote shell")
	rootCmd.Flags().String("bwlimit", "", "limit socket I/O bandwidth")
	rootCmd.Flags().BoolP("checksum", "c", false, "skip based on checksum, not mod-time & size")
	rootCmd.Flags().String("checksum-choice", "", "choose the checksum algorithm (aka --cc)")
	rootCmd.Flags().String("checksum-seed", "", "set block/file checksum seed (advanced)")
	rootCmd.Flags().String("chmod", "", "affect file and/or directory permissions")
	rootCmd.Flags().String("chown", "", "simple username/groupname mapping")
	rootCmd.Flags().String("compare-dest", "", "also compare destination files relative to DIR")
	rootCmd.Flags().BoolP("compress", "z", false, "compress file data during the transfer")
	rootCmd.Flags().String("compress-choice", "", "choose the compression algorithm (aka --zc)")
	rootCmd.Flags().String("compress-level", "", "explicitly set compression level (aka --zl)")
	rootCmd.Flags().String("contimeout", "", "set daemon connection timeout in seconds")
	rootCmd.Flags().String("copy-as", "", "specify user & optional group for the copy")
	rootCmd.Flags().String("copy-dest", "", "... and include copies of unchanged files")
	rootCmd.Flags().BoolP("copy-dirlinks", "k", false, "transform symlink to dir into referent dir")
	rootCmd.Flags().BoolP("copy-links", "L", false, "transform symlink into referent file/dir")
	rootCmd.Flags().Bool("copy-unsafe-links", false, "only \"unsafe\" symlinks are transformed")
	rootCmd.Flags().BoolP("crtimes", "N", false, "preserve create times (newness)")
	rootCmd.Flags().BoolP("cvs-exclude", "C", false, "auto-ignore files in the same way CVS does")
	rootCmd.Flags().String("debug", "", "fine-grained debug verbosity")
	rootCmd.Flags().Bool("del", false, "an alias for --delete-during")
	rootCmd.Flags().Bool("delay-updates", false, "put all updated files into place at end")
	rootCmd.Flags().Bool("delete", false, "delete extraneous files from dest dirs")
	rootCmd.Flags().Bool("delete-after", false, "receiver deletes after transfer, not during")
	rootCmd.Flags().Bool("delete-before", false, "receiver deletes before xfer, not during")
	rootCmd.Flags().Bool("delete-delay", false, "find deletions during, delete after")
	rootCmd.Flags().Bool("delete-during", false, "receiver deletes during the transfer")
	rootCmd.Flags().Bool("delete-excluded", false, "also delete excluded files from dest dirs")
	rootCmd.Flags().Bool("delete-missing-args", false, "delete missing source args from destination")
	rootCmd.Flags().Bool("devices", false, "preserve device files (super-user only)")
	rootCmd.Flags().BoolP("dirs", "d", false, "transfer directories without recursing")
	rootCmd.Flags().BoolP("dry-run", "n", false, "perform a trial run with no changes made")
	rootCmd.Flags().String("early-input", "", "use FILE for daemon's early exec input")
	rootCmd.Flags().String("exclude", "", "exclude files matching PATTERN")
	rootCmd.Flags().String("exclude-from", "", "read exclude patterns from FILE")
	rootCmd.Flags().BoolP("executability", "E", false, "preserve executability")
	rootCmd.Flags().Bool("existing", false, "skip creating new files on receiver")
	rootCmd.Flags().Bool("fake-super", false, "store/recover privileged attrs using xattrs")
	rootCmd.Flags().String("files-from", "", "read list of source-file names from FILE")
	rootCmd.Flags().StringP("filter", "f", "", "add a file-filtering RULE")
	rootCmd.Flags().Bool("force", false, "force deletion of dirs even if not empty")
	rootCmd.Flags().BoolP("from0", "0", false, "all *-from/filter files are delimited by 0s")
	rootCmd.Flags().BoolP("fuzzy", "y", false, "find similar file for basis if no dest file")
	rootCmd.Flags().BoolP("group", "g", false, "preserve group")
	rootCmd.Flags().String("groupmap", "", "custom groupname mapping")
	rootCmd.Flags().BoolP("hard-links", "H", false, "preserve hard links")
	rootCmd.Flags().BoolP("help", "h", false, "show this help (* -h is help only on its own)")
	rootCmd.Flags().Bool("human-readable", false, "output numbers in a human-readable format")
	rootCmd.Flags().String("iconv", "", "request charset conversion of filenames")
	rootCmd.Flags().Bool("ignore-errors", false, "delete even if there are I/O errors")
	rootCmd.Flags().Bool("ignore-existing", false, "skip updating files that exist on receiver")
	rootCmd.Flags().Bool("ignore-missing-args", false, "ignore missing source args without error")
	rootCmd.Flags().BoolP("ignore-times", "I", false, "don't skip files that match size and time")
	rootCmd.Flags().String("include", "", "don't exclude files matching PATTERN")
	rootCmd.Flags().String("include-from", "", "read include patterns from FILE")
	rootCmd.Flags().String("info", "", "fine-grained informational verbosity")
	rootCmd.Flags().Bool("inplace", false, "update destination files in-place")
	rootCmd.Flags().BoolP("ipv4", "4", false, "prefer IPv4")
	rootCmd.Flags().BoolP("ipv6", "6", false, "prefer IPv6")
	rootCmd.Flags().BoolP("itemize-changes", "i", false, "output a change-summary for all updates")
	rootCmd.Flags().BoolP("keep-dirlinks", "K", false, "treat symlinked dir on receiver as dir")
	rootCmd.Flags().String("link-dest", "", "hardlink to files in DIR when unchanged")
	rootCmd.Flags().BoolP("links", "l", false, "copy symlinks as symlinks")
	rootCmd.Flags().Bool("list-only", false, "list the files instead of copying them")
	rootCmd.Flags().String("log-file", "", "log what we're doing to the specified FILE")
	rootCmd.Flags().String("log-file-format", "", "log updates using the specified FMT")
	rootCmd.Flags().String("max-alloc", "", "change a limit relating to memory alloc")
	rootCmd.Flags().String("max-delete", "", "don't delete more than NUM files")
	rootCmd.Flags().String("max-size", "", "don't transfer any file larger than SIZE")
	rootCmd.Flags().String("min-size", "", "don't transfer any file smaller than SIZE")
	rootCmd.Flags().Bool("mkpath", false, "create the destination's path component")
	rootCmd.Flags().StringP("modify-window", "@", "", "set the accuracy for mod-time comparisons")
	rootCmd.Flags().Bool("munge-links", false, "munge symlinks to make them safe & unusable")
	rootCmd.Flags().Bool("no-OPTION", false, "turn off an implied OPTION (e.g. --no-D)")
	rootCmd.Flags().Bool("no-implied-dirs", false, "don't send implied dirs with --relative")
	rootCmd.Flags().Bool("no-motd", false, "suppress daemon-mode MOTD")
	rootCmd.Flags().Bool("numeric-ids", false, "don't map uid/gid values by user/group name")
	rootCmd.Flags().BoolP("omit-dir-times", "O", false, "omit directories from --times")
	rootCmd.Flags().BoolP("omit-link-times", "J", false, "omit symlinks from --times")
	rootCmd.Flags().BoolP("one-file-system", "x", false, "don't cross filesystem boundaries")
	rootCmd.Flags().String("only-write-batch", "", "like --write-batch but w/o updating dest")
	rootCmd.Flags().Bool("open-noatime", false, "avoid changing the atime on opened files")
	rootCmd.Flags().String("out-format", "", "output updates using the specified FORMAT")
	rootCmd.Flags().String("outbuf", "", "set out buffering to None, Line, or Block")
	rootCmd.Flags().BoolP("owner", "o", false, "preserve owner (super-user only)")
	rootCmd.Flags().Bool("partial", false, "keep partially transferred files")
	rootCmd.Flags().String("partial-dir", "", "put a partially transferred file into DIR")
	rootCmd.Flags().String("password-file", "", "read daemon-access password from FILE")
	rootCmd.Flags().BoolP("perms", "p", false, "preserve permissions")
	rootCmd.Flags().String("port", "", "specify double-colon alternate port number")
	rootCmd.Flags().Bool("preallocate", false, "allocate dest files before writing them")
	rootCmd.Flags().Bool("progress", false, "show progress during transfer")
	rootCmd.Flags().BoolP("protect-args", "s", false, "no space-splitting; wildcard chars only")
	rootCmd.Flags().String("protocol", "", "force an older protocol version to be used")
	rootCmd.Flags().BoolP("prune-empty-dirs", "m", false, "prune empty directory chains from file-list")
	rootCmd.Flags().BoolP("quiet", "q", false, "suppress non-error messages")
	rootCmd.Flags().String("read-batch", "", "read a batched update from FILE")
	rootCmd.Flags().BoolP("recursive", "r", false, "recurse into directories")
	rootCmd.Flags().BoolP("relative", "R", false, "use relative path names")
	rootCmd.Flags().StringP("remote-option", "M", "", "send OPTION to the remote side only")
	rootCmd.Flags().Bool("remove-source-files", false, "sender removes synchronized files (non-dir)")
	rootCmd.Flags().StringP("rsh", "e", "", "specify the remote shell to use")
	rootCmd.Flags().String("rsync-path", "", "specify the rsync to run on remote machine")
	rootCmd.Flags().Bool("safe-links", false, "ignore symlinks that point outside the tree")
	rootCmd.Flags().Bool("size-only", false, "skip files that match in size")
	rootCmd.Flags().String("skip-compress", "", "skip compressing files with suffix in LIST")
	rootCmd.Flags().String("sockopts", "", "specify custom TCP options")
	rootCmd.Flags().BoolP("sparse", "S", false, "turn sequences of nulls into sparse blocks")
	rootCmd.Flags().Bool("specials", false, "preserve special files")
	rootCmd.Flags().Bool("stats", false, "give some file-transfer stats")
	rootCmd.Flags().String("stderr", "", "change stderr output mode (default: errors)")
	rootCmd.Flags().String("stop-after", "", "Stop rsync after MINS minutes have elapsed")
	rootCmd.Flags().String("stop-at", "", "Stop rsync at the specified point in time")
	rootCmd.Flags().String("suffix", "", "backup suffix (default ~ w/o --backup-dir)")
	rootCmd.Flags().Bool("super", false, "receiver attempts super-user activities")
	rootCmd.Flags().StringP("temp-dir", "T", "", "create temporary files in directory DIR")
	rootCmd.Flags().String("timeout", "", "set I/O timeout in seconds")
	rootCmd.Flags().BoolP("times", "t", false, "preserve modification times")
	rootCmd.Flags().BoolP("update", "u", false, "skip files that are newer on the receiver")
	rootCmd.Flags().String("usermap", "", "custom username mapping")
	rootCmd.Flags().BoolP("verbose", "v", false, "increase verbosity")
	rootCmd.Flags().BoolP("version", "V", false, "print the version + other info and exit")
	rootCmd.Flags().BoolP("whole-file", "W", false, "copy files whole (w/o delta-xfer algorithm)")
	rootCmd.Flags().String("write-batch", "", "write a batched update to FILE")
	rootCmd.Flags().Bool("write-devices", false, "write to devices as files (implies --inplace)")
	rootCmd.Flags().BoolP("xattrs", "X", false, "preserve extended attributes")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"backup-dir":      carapace.ActionDirectories(),
		"checksum-choice": carapace.ActionValues("auto", "xxh128", "xxh3", "xxh64", "md5", "md4", "none"),
		"chmod": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			if c.CallbackValue == "" {
				return carapace.Batch(
					fs.ActionFileModes(),
					carapace.ActionValuesDescribed(
						"D", "only apply to directory",
						"F", "only apply to file",
					),
				).ToA()
			}

			if strings.HasPrefix(c.CallbackValue, "D") || strings.HasPrefix(c.CallbackValue, "F") {
				prefix := string(c.CallbackValue[0])
				c.CallbackValue = c.CallbackValue[1:]
				return fs.ActionFileModes().Invoke(c).Prefix(prefix).ToA()
			}
			return fs.ActionFileModes()
		}),
		"compress-choice": carapace.ActionValues("zstd", "lz4", "zlibx", "zlib", "none"),
		"copy-as":         os.ActionUserGroup(),
		"debug": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			r := regexp.MustCompile(`\d`)
			for index, part := range c.Parts {
				c.Parts[index] = r.ReplaceAllString(part, "")
			}
			return action.ActionFlags().Invoke(c).Filter(c.Parts).ToA()
		}),
		"early-input":  carapace.ActionFiles(),
		"exclude-from": carapace.ActionFiles(),
		"files-from":   carapace.ActionFiles(),
		"groupmap": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
				switch len(c.Parts) {
				case 0:
					return os.ActionGroups().Invoke(c).Suffix(":").ToA()
				default:
					return carapace.ActionValues()
				}
			})
		}),
		"include-from": carapace.ActionFiles(),
		"info": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			r := regexp.MustCompile(`\d`)
			for index, part := range c.Parts {
				c.Parts[index] = r.ReplaceAllString(part, "")
			}
			return action.ActionFlags().Invoke(c).Filter(c.Parts).ToA()
		}),
		"log-file":         carapace.ActionFiles(),
		"log-file-format":  action.ActionFormats(),
		"only-write-batch": carapace.ActionFiles(),
		"out-format":       action.ActionFormats(),
		"outbuf":           carapace.ActionValues("None", "Line", "Block"),
		"partial-dir":      carapace.ActionDirectories(),
		"password-file":    carapace.ActionFiles(),
		"port":             net.ActionPorts(),
		"read-batch":       carapace.ActionFiles(),
		"skip-compress": carapace.ActionMultiParts("/", func(c carapace.Context) carapace.Action {
			return fs.ActionFilenameExtensions().Invoke(c).Filter(c.Parts).ToA()
		}),
		"stderr": carapace.ActionValuesDescribed(
			"errors", "causes all the rsync processes to send an error directly to stderr",
			"all", "causes all rsync messages to get written directly to stderr from all processes",
			"client", "causes all rsync messages to be sent to the client side via the  protocol stream",
		),
		"temp-dir": carapace.ActionDirectories(),
		"usermap": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
				switch len(c.Parts) {
				case 0:
					return os.ActionUsers().Invoke(c).Suffix(":").ToA()
				default:
					return carapace.ActionValues()
				}
			})
		}),
		"write-batch": carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if util.HasPathPrefix(c.CallbackValue) {
				return carapace.ActionFiles()
			}

			return carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
				switch len(c.Parts) {
				case 0:
					return carapace.ActionMultiParts("@", func(c carapace.Context) carapace.Action {
						switch len(c.Parts) {
						case 0:
							return carapace.Batch(
								os.ActionUsers().Invoke(c).Suffix("@").ToA(),
								net.ActionHosts().Invoke(c).Suffix(":").ToA(),
							).ToA()
						case 1:
							return net.ActionHosts().Invoke(c).Suffix(":").ToA()
						default:
							return carapace.ActionValues()
						}
					})
				default:
					return carapace.ActionValues()
				}
			})
		}),
	)
}
