package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/fs"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "find",
	Short: "search for files in a directory hierarchy",
	Long:  "https://linux.die.net/man/1/find",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("D", "D", "", "Print diagnostic information")
	rootCmd.Flags().BoolS("H", "H", false, "Do not follow symbolic links")
	rootCmd.Flags().BoolS("L", "L", false, "Follow symbolic links")
	rootCmd.Flags().BoolS("O0", "O0", false, "Equivalent to optimisation level 1")
	rootCmd.Flags().BoolS("O1", "O1", false, "This  is  the default optimisation level and corresponds to the traditional behaviour.")
	rootCmd.Flags().BoolS("O2", "O2", false, "Any -type or -xtype tests are performed after any tests based only on the names of files.")
	rootCmd.Flags().BoolS("O3", "O3", false, "At  this  optimisation level, the full cost-based query optimiser is enabled.")
	rootCmd.Flags().BoolS("P", "P", false, "Never follow symbolic links")
	rootCmd.Flags().StringS("amin", "amin", "", "File was last accessed less than, more than or exactly n minutes ago.")
	rootCmd.Flags().StringS("anewer", "anewer", "", "Time of the last access of the current file is more recent than that  of  the  last  data modification of the reference file.")
	rootCmd.Flags().StringS("atime", "atime", "", "File was last accessed less than, more than or exactly n*24 hours ago.")
	rootCmd.Flags().StringS("cmin", "cmin", "", "File's status was last changed less than, more than or exactly n minutes ago.")
	rootCmd.Flags().StringS("cnewer", "cnewer", "", "Time  of  the last status change of the current file is more recent than that of the last data modification of the reference file.")
	rootCmd.Flags().StringS("context", "context", "", "Security context of the file matches glob pattern.")
	rootCmd.Flags().StringS("ctime", "ctime", "", "File's status was last changed less than, more than or exactly n*24 hours ago.")
	rootCmd.Flags().BoolS("daystart", "daystart", false, "Measure times from the beginning of today rather than from 24 hours ago.")
	rootCmd.Flags().BoolS("delete", "delete", false, "Delete  files.")
	rootCmd.Flags().BoolS("depth", "depth", false, "Process each directory's contents before the directory itself.")
	rootCmd.Flags().StringS("exec", "exec", "", "Execute command.")
	rootCmd.Flags().StringS("execdir", "execdir", "", "Like -exec, but the specified command is run from the subdirectory containing the matched file.")
	rootCmd.Flags().BoolS("executable", "executable", false, "Matches files which are executable and directories which are searchable by the current user.")
	rootCmd.Flags().BoolS("false", "false", false, "Always false.")
	rootCmd.Flags().BoolS("fls", "fls", false, "Unusual characters are always escaped.")
	rootCmd.Flags().BoolS("fprint", "fprint", false, "Quoting is handled in the same way as for -printf and -fprintf.")
	rootCmd.Flags().BoolS("fprint0", "fprint0", false, "Always print the exact filename, unchanged, even if the output is going to a terminal.")
	rootCmd.Flags().BoolS("fprintf", "fprintf", false, "If the output is not going to a terminal, it is printed as-is.")
	rootCmd.Flags().StringS("fstype", "fstype", "", "File  is  on  a filesystem of type type.")
	rootCmd.Flags().StringS("gid", "gid", "", "File's numeric group ID is less than, more than or exactly n.")
	rootCmd.Flags().StringS("group", "group", "", "File belongs to group gname (numeric group ID allowed).")
	rootCmd.Flags().BoolP("help", "help", false, "Print a summary of the command-line usage of find and exit.")
	rootCmd.Flags().BoolS("ignore_readdir_race", "ignore_readdir_race", false, "Do not  emit an error message when find fails to stat a file.")
	rootCmd.Flags().StringS("ilname", "ilname", "", "Like -lname, but the match is case insensitive.")
	rootCmd.Flags().StringS("iname", "iname", "", "Like -name, but the match is case insensitive.")
	rootCmd.Flags().StringS("inum", "inum", "", "File  has inode number smaller than, greater than or exactly n.")
	rootCmd.Flags().StringS("ipath", "ipath", "", "Like -path.  but the match is case insensitive.")
	rootCmd.Flags().StringS("iregex", "iregex", "", "Like -regex, but the match is case insensitive.")
	rootCmd.Flags().StringS("iwholename", "iwholename", "", "See -ipath.  This alternative is less portable than -ipath.")
	rootCmd.Flags().StringS("links", "links", "", "File has less than, more than or exactly n hard links.")
	rootCmd.Flags().StringS("lname", "lname", "", "File is a symbolic link whose contents match shell pattern pattern.")
	rootCmd.Flags().BoolS("ls", "ls", false, "List current file in ls -dils format on standard output.")
	rootCmd.Flags().StringS("maxdepth", "maxdepth", "", "Descend at most levels of directories below the startingpoints.")
	rootCmd.Flags().StringS("mindepth", "mindepth", "", "Do  not  apply  any tests or actions at levels less than levels.")
	rootCmd.Flags().StringS("mmin", "mmin", "", "File's data was last modified less than, more than or exactly n minutes ago.")
	rootCmd.Flags().BoolS("mount", "mount", false, "Don't descend directories on other filesystems.")
	rootCmd.Flags().StringS("mtime", "mtime", "", "File's data was last modified less than, more than or exactly n*24 hours  ago.")
	rootCmd.Flags().StringS("name", "name", "", "Base of file name matches  shell  pattern pattern.")
	rootCmd.Flags().StringS("newer", "newer", "", "Time of the last data modification of the current file is more recent than that of the last data modification of the reference file.")
	rootCmd.Flags().BoolP("version", "version", false, "Print the find version number and exit.")
	// TODO XY completion
	//rootCmd.Flags().StringS("newerXY", "newerXY", "", "Succeeds if timestamp X of the file being considered is newer than timestamp  Y  of  the file reference.")
	rootCmd.Flags().BoolS("nogroup", "nogroup", false, "No group corresponds to file's numeric group ID.")
	rootCmd.Flags().BoolS("noignore_readdir_race", "noignore_readdir_race", false, "Turns off the effect of -ignore_readdir_race.")
	rootCmd.Flags().BoolS("noleaf", "noleaf", false, "Do  not  optimize  by assuming that directories contain 2 fewer subdirectories than their hard link count.")
	rootCmd.Flags().BoolS("nouser", "nouser", false, "No user corresponds to file's numeric user ID.")
	rootCmd.Flags().BoolS("nowarn", "nowarn", false, "Turn  warning  messages  off.")
	rootCmd.Flags().StringS("ok", "ok", "", "Like  -exec but ask the user first.")
	rootCmd.Flags().StringS("okdir", "okdir", "", "Like -execdir but ask the user first in the same way as for -ok.")
	rootCmd.Flags().StringS("path", "path", "", "File name matches shell pattern pattern.")
	rootCmd.Flags().StringS("perm", "perm", "", "File's permission bits are exactly mode (octal or symbolic).")
	rootCmd.Flags().BoolS("print", "print", false, "Print the full file name on the standard output, followed by a newline.")
	rootCmd.Flags().BoolS("print0", "print0", false, "Print the full file name on the standard output, followed by a null character.")
	rootCmd.Flags().StringS("printf", "printf", "", "Print  format on the standard output, interpreting `\\' escapes and `%' directives.")
	rootCmd.Flags().BoolS("prune", "prune", false, "If the file is a directory, do not descend into it.")
	rootCmd.Flags().BoolS("quit", "quit", false, "Exit immediately.")
	rootCmd.Flags().BoolS("readable", "readable", false, "Matches  files  which  are  readable by the current user.")
	rootCmd.Flags().StringS("regex", "regex", "", "File  name  matches regular expression pattern.")
	rootCmd.Flags().StringS("regextype", "regextype", "", "Changes the regular expression syntax understood by -regex and -iregex tests which  occur later  on the command line.")
	rootCmd.Flags().StringS("samefile", "samefile", "", "File  refers  to the same inode as name.")
	rootCmd.Flags().StringS("size", "size", "", "File uses less than, more than or exactly n units of space, rounding up.")
	rootCmd.Flags().BoolS("true", "true", false, "Always true.")
	rootCmd.Flags().StringS("type", "type", "", "file type")
	rootCmd.Flags().StringS("used", "used", "", "File  was  last accessed less than, more than or exactly n days after its status was last changed.")
	rootCmd.Flags().StringS("user", "user", "", "File is owned by user uname.")
	rootCmd.Flags().BoolS("warn", "warn", false, "Turn  warning  messages  on.")
	rootCmd.Flags().StringS("wholename", "wholename", "", "See -path.  This alternative is less portable than -path.")
	rootCmd.Flags().BoolS("writable", "writable", false, "Matches files which are writable by the current user.")
	rootCmd.Flags().BoolS("xdev", "xdev", false, "Don't descend directories on other filesystems.")
	rootCmd.Flags().StringS("xtype", "xtype", "", "The same as -type unless the file is a symbolic link.")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"D": carapace.ActionValuesDescribed(
			"exec", "Show diagnostic information relating to -exec, -execdir, -ok and -okdir",
			"opt", "Prints diagnostic information relating to the optimisation of the expression tree",
			"rates", "Prints a summary indicating how often each predicate succeeded or failed.",
			"search", "Navigate the directory tree verbosely.",
			"stat", "Print  messages  as  files are examined with the stat and lstat system calls.",
			"tree", "Show the expression tree in its original and optimised form.",
			"all", "Enable all of the other debug options (but help).",
			"help", "Explain the debugging options.",
		),
		"fstype":    fs.ActionFilesystemTypes(),
		"group":     os.ActionGroups(),
		"perm":      fs.ActionFileModes(),
		"regextype": carapace.ActionValues("findutils-default", "ed", "emacs", "gnu-awk", "grep", "posix-awk", "awk", "posix-basic", "posix-egrep", "egrep", "posix-extended", "posix-minimal-basic", "sed"),
		"type": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return ActionFileTypes().Invoke(c).Filter(c.Parts).ToA()
		}),
		"user": os.ActionUsers(),
		"xtype": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return ActionFileTypes().Invoke(c).Filter(c.Parts).ToA()
		}),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionDirectories(),
	)
}

func ActionFileTypes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"b", "block (buffered) special",
		"c", "character (unbuffered) special",
		"d", "directory",
		"p", "named pipe (FIFO)",
		"f", "regular file",
		"l", "symbolic link",
		"s", "socket",
		"D", "door (Solaris)",
	)
}
