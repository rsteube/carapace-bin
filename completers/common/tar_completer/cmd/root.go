package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/fs"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tar",
	Short: "tar - an archiving utility",
	Long:  "https://linux.die.net/man/1/tar",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("absolute-names", "P", false, "don't strip leading '/'s from file names")
	rootCmd.Flags().Bool("acls", false, "Enable the POSIX ACLs support")
	rootCmd.Flags().String("add-file", "", "add given FILE to the archive")
	rootCmd.Flags().String("after-date", "", "only store files older than DATE-OR-FILE")
	rootCmd.Flags().Bool("anchored", false, "patterns match file name start")
	rootCmd.Flags().BoolP("append", "r", false, "append files to the end of an archive")
	rootCmd.Flags().String("atime-preserve", "", "preserve access times on dumped files")
	rootCmd.Flags().BoolP("auto-compress", "a", false, "use archive suffix to determine the compression program")
	rootCmd.Flags().String("backup", "", "backup before removal")
	rootCmd.Flags().BoolP("block-number", "R", false, "show block number within archive with each message")
	rootCmd.Flags().StringP("blocking-factor", "b", "", "BLOCKS x 512 bytes per record")
	rootCmd.Flags().BoolP("bzip2", "j", false, "filter the archive through bzip2")
	rootCmd.Flags().BoolP("catenate", "A", false, "append tar files to an archive")
	rootCmd.Flags().Bool("check-device", false, "check device numbers when creating incremental archives")
	rootCmd.Flags().BoolP("check-links", "l", false, "print a message if not all links are dumped")
	rootCmd.Flags().String("checkpoint", "", "display progress messages every NUMBERth record")
	rootCmd.Flags().String("checkpoint-action", "", "execute ACTION on each checkpoint")
	rootCmd.Flags().Bool("clamp-mtime", false, "only set time when the file is more recent than what was given with --mtime")
	rootCmd.Flags().Bool("compare", false, "find differences between archive and file system")
	rootCmd.Flags().BoolP("compress", "Z", false, "filter the archive through compress")
	rootCmd.Flags().Bool("concatenate", false, "append tar files to an archive")
	rootCmd.Flags().Bool("confirmation", false, "ask for confirmation for every action")
	rootCmd.Flags().BoolP("create", "c", false, "create a new archive")
	rootCmd.Flags().Bool("delay-directory-restore", false, "delay setting modification times and permissions until the end of extraction")
	rootCmd.Flags().Bool("delete", false, "delete from the archive (not on mag tapes!)")
	rootCmd.Flags().BoolP("dereference", "h", false, "follow symlinks")
	rootCmd.Flags().BoolP("diff", "d", false, "find differences between archive and file system")
	rootCmd.Flags().StringP("directory", "C", "", "change to directory DIR")
	rootCmd.Flags().String("exclude", "", "exclude files, given as a PATTERN")
	rootCmd.Flags().Bool("exclude-backups", false, "exclude backup and lock files")
	rootCmd.Flags().Bool("exclude-caches", false, "exclude contents of directories containing CACHEDIR.TAG")
	rootCmd.Flags().Bool("exclude-caches-all", false, "exclude directories containing CACHEDIR.TAG")
	rootCmd.Flags().Bool("exclude-caches-under", false, "exclude everything under directories containing CACHEDIR.TAG")
	rootCmd.Flags().StringP("exclude-from", "X", "", "exclude patterns listed in FILE")
	rootCmd.Flags().String("exclude-ignore", "", "read exclude patterns for each directory from FILE")
	rootCmd.Flags().String("exclude-ignore-recursive", "", "read exclude patterns for each directory and its subdirectories from FILE")
	rootCmd.Flags().String("exclude-tag", "", "exclude contents of directories containing FILE")
	rootCmd.Flags().String("exclude-tag-all", "", "exclude directories containing FILE")
	rootCmd.Flags().String("exclude-tag-under", "", "exclude everything under directories containing FILE")
	rootCmd.Flags().Bool("exclude-vcs", false, "exclude version control system directories")
	rootCmd.Flags().Bool("exclude-vcs-ignores", false, "read exclude patterns from the VCS ignore files")
	rootCmd.Flags().BoolP("extract", "x", false, "extract files from an archive")
	rootCmd.Flags().StringP("file", "f", "", "use archive file or device ARCHIVE")
	rootCmd.Flags().StringP("files-from", "T", "", "get names to extract or create from FILE")
	rootCmd.Flags().Bool("force-local", false, "archive file is local even if it has a colon")
	rootCmd.Flags().StringP("format", "H", "", "create archive of the given format")
	rootCmd.Flags().Bool("full-time", false, "print file time to its full resolution")
	rootCmd.Flags().Bool("get", false, "extract files from an archive")
	rootCmd.Flags().String("group", "", "force NAME as group for added files")
	rootCmd.Flags().String("group-map", "", "use FILE to map file owner GIDs and names")
	rootCmd.Flags().Bool("gunzip", false, "filter the archive through gzip")
	rootCmd.Flags().BoolP("gzip", "z", false, "filter the archive through gzip")
	rootCmd.Flags().Bool("hard-dereference", false, "follow hard links")
	rootCmd.Flags().BoolP("help", "?", false, "give this help list")
	rootCmd.Flags().String("hole-detection", "", "technique to detect holes")
	rootCmd.Flags().Bool("ignore-case", false, "ignore case")
	rootCmd.Flags().Bool("ignore-command-error", false, "ignore exit codes of children")
	rootCmd.Flags().Bool("ignore-failed-read", false, "do not exit with nonzero on unreadable files")
	rootCmd.Flags().BoolP("ignore-zeros", "i", false, "ignore zeroed blocks in archive (means EOF)")
	rootCmd.Flags().BoolP("incremental", "G", false, "handle old GNU-format incremental backup")
	rootCmd.Flags().String("index-file", "", "send verbose output to FILE")
	rootCmd.Flags().StringP("info-script", "F", "", "run script at end of each tape")
	rootCmd.Flags().BoolP("interactive", "w", false, "ask for confirmation for every action")
	rootCmd.Flags().Bool("keep-directory-symlink", false, "preserve existing symlinks to directories when extracting")
	rootCmd.Flags().Bool("keep-newer-files", false, "don't replace existing files that are newer than their archive copies")
	rootCmd.Flags().BoolP("keep-old-files", "k", false, "don't replace existing files when extracting, treat them as errors")
	rootCmd.Flags().StringP("label", "V", "", "create archive with volume name TEXT")
	rootCmd.Flags().String("level", "", "dump level for created listed-incremental archive")
	rootCmd.Flags().BoolP("list", "t", false, "list the contents of an archive")
	rootCmd.Flags().StringP("listed-incremental", "g", "", "handle new GNU-format incremental backup")
	rootCmd.Flags().Bool("lzip", false, "filter the archive through lzip")
	rootCmd.Flags().Bool("lzma", false, "filter the archive through lzma")
	rootCmd.Flags().Bool("lzop", false, "filter the archive through lzop")
	rootCmd.Flags().String("mode", "", "force (symbolic) mode CHANGES for added files")
	rootCmd.Flags().String("mtime", "", "set mtime for added files from DATE-OR-FILE")
	rootCmd.Flags().BoolP("multi-volume", "M", false, "create/list/extract multi-volume archive")
	rootCmd.Flags().String("new-volume-script", "", "run script at end of each tape")
	rootCmd.Flags().StringP("newer", "N", "", "only store files newer than DATE-OR-FILE")
	rootCmd.Flags().String("newer-mtime", "", "compare date and time when data changed only")
	rootCmd.Flags().Bool("no-acls", false, "Disable the POSIX ACLs support")
	rootCmd.Flags().Bool("no-anchored", false, "patterns match after any '/'")
	rootCmd.Flags().Bool("no-auto-compress", false, "do not use archive suffix to determine the compression program")
	rootCmd.Flags().Bool("no-check-device", false, "do not check device numbers when creating")
	rootCmd.Flags().Bool("no-delay-directory-restore", false, "cancel the effect of --delay-directory-restore option")
	rootCmd.Flags().Bool("no-ignore-case", false, "case sensitive matching (default)")
	rootCmd.Flags().Bool("no-ignore-command-error", false, "treat non-zero exit codes of children as error")
	rootCmd.Flags().Bool("no-null", false, "disable the effect of the previous --null option")
	rootCmd.Flags().Bool("no-overwrite-dir", false, "preserve metadata of existing directories")
	rootCmd.Flags().String("no-quote-chars", "", "disable quoting for characters from STRING")
	rootCmd.Flags().Bool("no-recursion", false, "avoid descending automatically in directories")
	rootCmd.Flags().Bool("no-same-owner", false, "extract files as yourself")
	rootCmd.Flags().Bool("no-same-permissions", false, "apply the user's umask when extracting permissions from the archive")
	rootCmd.Flags().Bool("no-seek", false, "archive is not seekable")
	rootCmd.Flags().Bool("no-selinux", false, "Disable the SELinux context support")
	rootCmd.Flags().Bool("no-unquote", false, "do not unquote input file or member names")
	rootCmd.Flags().Bool("no-verbatim-files-from", false, "treats file names starting with dash as options")
	rootCmd.Flags().Bool("no-wildcards", false, "verbatim string matching")
	rootCmd.Flags().Bool("no-wildcards-match-slash", false, "wildcards do not match '/'")
	rootCmd.Flags().Bool("no-xattrs", false, "Disable extended attributes support")
	rootCmd.Flags().Bool("null", false, "reads null-terminated names; implies")
	rootCmd.Flags().Bool("numeric-owner", false, "always use numbers for user/group names")
	rootCmd.Flags().BoolS("o", "o", false, "when creating, same as --old-archive; when extracting, same as --no-same-owner")
	rootCmd.Flags().String("occurrence", "", "process only the NUMBERth occurrence of each file in the archive")
	rootCmd.Flags().Bool("old-archive", false, "same as --format=v7")
	rootCmd.Flags().Bool("one-file-system", false, "stay in local file system when creating archive")
	rootCmd.Flags().String("one-top-level", "", "create a subdirectory to avoid having loose files extracted")
	rootCmd.Flags().Bool("overwrite", false, "overwrite existing files when extracting")
	rootCmd.Flags().Bool("overwrite-dir", false, "overwrite metadata of existing directories when extracting")
	rootCmd.Flags().String("owner", "", "force NAME as owner for added files")
	rootCmd.Flags().String("owner-map", "", "use FILE to map file owner UIDs and names")
	rootCmd.Flags().String("pax-option", "", "control pax keywords")
	rootCmd.Flags().Bool("portability", false, "same as --format=v7")
	rootCmd.Flags().Bool("posix", false, "same as --format=posix")
	rootCmd.Flags().BoolP("preserve-order", "s", false, "member arguments are listed in the same order")
	rootCmd.Flags().BoolP("preserve-permissions", "p", false, "extract information about file permissions")
	rootCmd.Flags().String("quote-chars", "", "additionally quote characters from STRING")
	rootCmd.Flags().String("quoting-style", "", "set name quoting style")
	rootCmd.Flags().BoolP("read-full-records", "B", false, "reblock as we read (for 4.2BSD pipes)")
	rootCmd.Flags().String("record-size", "", "NUMBER of bytes per record, multiple of 512")
	rootCmd.Flags().Bool("recursion", false, "recurse into directories (default)")
	rootCmd.Flags().Bool("recursive-unlink", false, "empty hierarchies prior to extracting directory")
	rootCmd.Flags().Bool("remove-files", false, "remove files after adding them to the archive")
	rootCmd.Flags().Bool("restrict", false, "disable use of some potentially harmful options")
	rootCmd.Flags().String("rmt-command", "", "use given rmt COMMAND instead of rmt")
	rootCmd.Flags().String("rsh-command", "", "use remote COMMAND instead of rsh")
	rootCmd.Flags().Bool("same-order", false, "member arguments are listed in the same order")
	rootCmd.Flags().Bool("same-owner", false, "try extracting files with the same ownership")
	rootCmd.Flags().Bool("same-permissions", false, "extract information about file permissions")
	rootCmd.Flags().BoolP("seek", "n", false, "archive is seekable")
	rootCmd.Flags().Bool("selinux", false, "Enable the SELinux context support")
	rootCmd.Flags().Bool("show-defaults", false, "show tar defaults")
	rootCmd.Flags().Bool("show-omitted-dirs", false, "list each directory that does not match search criteria")
	rootCmd.Flags().Bool("show-snapshot-field-ranges", false, "show valid ranges for snapshot-file fields")
	rootCmd.Flags().Bool("show-stored-names", false, "show file or archive names after transformation")
	rootCmd.Flags().Bool("show-transformed-names", false, "show file or archive names after transformation")
	rootCmd.Flags().Bool("skip-old-files", false, "don't replace existing files when extracting")
	rootCmd.Flags().String("sort", "", "directory sorting order")
	rootCmd.Flags().BoolP("sparse", "S", false, "handle sparse files efficiently")
	rootCmd.Flags().String("sparse-version", "", "set version of the sparse format to use")
	rootCmd.Flags().StringP("starting-file", "K", "", "begin at member MEMBER-NAME when reading the archive")
	rootCmd.Flags().String("strip-components", "", "strip NUMBER leading components from file names on extraction")
	rootCmd.Flags().String("suffix", "", "backup before removal")
	rootCmd.Flags().StringP("tape-length", "L", "", "change tape after writing NUMBER x 1024 bytes")
	rootCmd.Flags().Bool("test-label", false, "test the archive volume label and exit")
	rootCmd.Flags().String("to-command", "", "pipe extracted files to another program")
	rootCmd.Flags().BoolP("to-stdout", "O", false, "extract files to standard output")
	rootCmd.Flags().String("totals", "", "print total bytes after processing the archive")
	rootCmd.Flags().BoolP("touch", "m", false, "don't extract file modified time")
	rootCmd.Flags().String("transform", "", "use sed replace EXPRESSION to transform file names")
	rootCmd.Flags().Bool("uncompress", false, "filter the archive through compress")
	rootCmd.Flags().Bool("ungzip", false, "filter the archive through gzip")
	rootCmd.Flags().BoolP("unlink-first", "U", false, "remove each file prior to extracting over it")
	rootCmd.Flags().Bool("unquote", false, "unquote input file or member names (default)")
	rootCmd.Flags().BoolP("update", "u", false, "only append files newer than copy in archive")
	rootCmd.Flags().Bool("usage", false, "give a short usage message")
	rootCmd.Flags().StringP("use-compress-program", "I", "", "filter through PROG (must accept -d)")
	rootCmd.Flags().Bool("utc", false, "print file modification times in UTC")
	rootCmd.Flags().Bool("verbatim-files-from", false, "reads file names verbatim")
	rootCmd.Flags().BoolP("verbose", "v", false, "verbosely list files processed")
	rootCmd.Flags().BoolP("verify", "W", false, "attempt to verify the archive after writing it")
	rootCmd.Flags().Bool("version", false, "print program version")
	rootCmd.Flags().String("volno-file", "", "use/update the volume number in FILE")
	rootCmd.Flags().String("warning", "", "warning control")
	rootCmd.Flags().Bool("wildcards", false, "use wildcards (default for exclusion)")
	rootCmd.Flags().Bool("wildcards-match-slash", false, "wildcards match '/' (default for exclusion)")
	rootCmd.Flags().Bool("xattrs", false, "Enable extended attributes support")
	rootCmd.Flags().String("xattrs-exclude", "", "specify the exclude pattern for xattr keys")
	rootCmd.Flags().String("xattrs-include", "", "specify the include pattern for xattr keys")
	rootCmd.Flags().String("xform", "", "use sed replace EXPRESSION to transform file names")
	rootCmd.Flags().BoolP("xz", "J", false, "filter the archive through xz")
	rootCmd.Flags().Bool("zstd", false, "filter the archive through zstd")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"add-file":                 carapace.ActionFiles(),
		"after-date":               carapace.ActionFiles(),
		"atime-preserve":           carapace.ActionValues("replace", "system"),
		"backup":                   carapace.ActionValues("existing", "never", "nil", "numbered", "simple", "t"),
		"directory":                carapace.ActionDirectories(),
		"exclude-from":             carapace.ActionFiles(),
		"exclude-ignore":           carapace.ActionFiles(),
		"exclude-ignore-recursive": carapace.ActionFiles(),
		"exclude-tag":              carapace.ActionFiles(),
		"exclude-tag-all":          carapace.ActionFiles(),
		"exclude-tag-under":        carapace.ActionFiles(),
		"file":                     carapace.ActionFiles(),
		"files-from":               carapace.ActionFiles(),
		"format":                   carapace.ActionValues("gnu", "oldgnu", "pax", "posix", "ustar", "v7"),
		"group":                    os.ActionGroups(),
		"group-map":                carapace.ActionFiles(),
		"index-file":               carapace.ActionFiles(),
		"info-script":              carapace.ActionFiles(),
		"mode":                     fs.ActionFileModesSymbolic(),
		"mtime":                    carapace.ActionFiles(),
		"new-volume-script":        carapace.ActionFiles(),
		"newer":                    carapace.ActionFiles(),
		"owner":                    os.ActionUsers(),
		"owner-map":                carapace.ActionFiles(),
		"quoting-style":            carapace.ActionValues("c", "clocale", "c-maybe", "escape", "literal", "locale", "shell", "shell-always"),
		"to-command":               carapace.ActionFiles(),
		"use-compress-program":     carapace.ActionFiles(),
		"volno-file":               carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if rootCmd.Flag("delete").Changed ||
				rootCmd.Flag("list").Changed ||
				rootCmd.Flag("extract").Changed {
				return fs.ActionTarFileContents(rootCmd.Flag("file").Value.String())
			}
			return carapace.ActionFiles()
		}),
	)
}
