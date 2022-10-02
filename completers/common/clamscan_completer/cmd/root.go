package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "clamscan",
	Short: "scan files and directories for viruses",
	Long:  "http://www.clamav.net/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("alert-broken", false, "Alert on broken executable files (PE & ELF)")
	rootCmd.Flags().Bool("alert-broken-media", false, "Alert on broken graphics files (JPEG, TIFF, PNG, GIF)")
	rootCmd.Flags().Bool("alert-encrypted", false, "Alert on encrypted archives and documents")
	rootCmd.Flags().Bool("alert-encrypted-archive", false, "Alert on encrypted archives")
	rootCmd.Flags().Bool("alert-encrypted-doc", false, "Alert on encrypted documents")
	rootCmd.Flags().Bool("alert-exceeds-max", false, "Alert on files that exceed max file size, max scan size, or max recursion limit")
	rootCmd.Flags().Bool("alert-macros", false, "Alert on OLE2 files containing VBA macros")
	rootCmd.Flags().Bool("alert-partition-intersection", false, "Alert on raw DMG image files containing partition intersections")
	rootCmd.Flags().Bool("alert-phishing-cloak", false, "Alert on emails containing cloaked URLs")
	rootCmd.Flags().Bool("alert-phishing-ssl", false, "Alert on emails containing SSL mismatches in URLs")
	rootCmd.Flags().BoolP("allmatch", "z", false, "Continue scanning within file after finding a match")
	rootCmd.Flags().BoolP("archive-verbose", "a", false, "Show filenames inside scanned archives")
	rootCmd.Flags().Bool("bell", false, "Sound bell on virus detection")
	rootCmd.Flags().Bool("bytecode", false, "Load bytecode from the database")
	rootCmd.Flags().String("bytecode-timeout", "", "Set bytecode timeout (in milliseconds)")
	rootCmd.Flags().Bool("bytecode-unsigned", false, "Load unsigned bytecode")
	rootCmd.Flags().String("copy", "", "Copy infected files into DIRECTORY")
	rootCmd.Flags().Bool("cross-fs", false, "Scan files and directories on other filesystems")
	rootCmd.Flags().StringP("database", "d", "", "Load virus database from FILE or load all supported db files from DIR")
	rootCmd.Flags().Bool("debug", false, "Enable libclamav's debug messages")
	rootCmd.Flags().Bool("detect-pua", false, "Detect Possibly Unwanted Applications")
	rootCmd.Flags().Bool("detect-structured", false, "Detect structured data (SSN, Credit Card)")
	rootCmd.Flags().Bool("disable-cache", false, "Disable caching and cache checks for hash sums of scanned files.")
	rootCmd.Flags().Bool("dumpcerts", false, "Dump authenticode certificate chain in PE files")
	rootCmd.Flags().String("exclude", "", "Don't scan file names matching REGEX")
	rootCmd.Flags().String("exclude-dir", "", "Don't scan directories matching REGEX")
	rootCmd.Flags().String("exclude-pua", "", "Skip PUA sigs of category CAT")
	rootCmd.Flags().StringP("file-list", "f", "", "Scan files from FILE")
	rootCmd.Flags().Bool("follow-dir-symlinks", false, "Follow directory symlinks (0 = never, 1 = direct, 2 = always)")
	rootCmd.Flags().Bool("follow-file-symlinks", false, "Follow file symlinks (0 = never, 1 = direct, 2 = always)")
	rootCmd.Flags().Bool("gen-json", false, "Generate JSON metadata for the scanned file(s). For testing & development use ONLY.")
	rootCmd.Flags().BoolP("help", "h", false, "Show this help")
	rootCmd.Flags().Bool("heuristic-alerts", false, "Heuristic alerts")
	rootCmd.Flags().Bool("heuristic-scan-precedence", false, "Stop scanning as soon as a heuristic match is found")
	rootCmd.Flags().String("include", "", "Only scan file names matching REGEX")
	rootCmd.Flags().String("include-dir", "", "Only scan directories matching REGEX")
	rootCmd.Flags().String("include-pua", "", "Load PUA sigs of category CAT")
	rootCmd.Flags().BoolP("infected", "i", false, "Only print infected files")
	rootCmd.Flags().Bool("leave-temps", false, "Do not remove temporary files")
	rootCmd.Flags().StringP("log", "l", "", "Save scan report to FILE")
	rootCmd.Flags().String("max-dir-recursion", "", "Maximum directory recursion level")
	rootCmd.Flags().String("max-embeddedpe", "", "Maximum size file to check for embedded PE")
	rootCmd.Flags().String("max-files", "", "The maximum number of files to scan for each container file (**)")
	rootCmd.Flags().String("max-filesize", "", "Files larger than this will be skipped and assumed clean")
	rootCmd.Flags().String("max-htmlnormalize", "", "Maximum size of HTML file to normalize")
	rootCmd.Flags().String("max-htmlnotags", "", "Maximum size of normalized HTML file to scan")
	rootCmd.Flags().String("max-iconspe", "", "Maximum number of icons in PE file to be scanned")
	rootCmd.Flags().String("max-partitions", "", "Maximum number of partitions in disk image to be scanned")
	rootCmd.Flags().String("max-rechwp3", "", "Maximum recursive calls to HWP3 parsing function")
	rootCmd.Flags().String("max-recursion", "", "Maximum archive recursion level for container file (**)")
	rootCmd.Flags().String("max-scansize", "", "The maximum amount of data to scan for each container file (**)")
	rootCmd.Flags().String("max-scantime", "", "Scan time longer than this will be skipped and assumed clean (milliseconds)")
	rootCmd.Flags().String("max-scriptnormalize", "", "Maximum size of script file to normalize")
	rootCmd.Flags().String("max-ziptypercg", "", "Maximum size zip to type reanalyze")
	rootCmd.Flags().String("move", "", "Move infected files into DIRECTORY")
	rootCmd.Flags().Bool("no-summary", false, "Disable summary at end of scanning")
	rootCmd.Flags().Bool("nocerts", false, "Disable authenticode certificate chain verification in PE files")
	rootCmd.Flags().Bool("normalize", false, "Normalize html, script, and text files. Use normalize=no for yara compatibility")
	rootCmd.Flags().Bool("official-db-only", false, "Only load official signatures")
	rootCmd.Flags().String("pcre-match-limit", "", "Maximum calls to the PCRE match function.")
	rootCmd.Flags().String("pcre-max-filesize", "", "Maximum size file to perform PCRE subsig matching.")
	rootCmd.Flags().String("pcre-recmatch-limit", "", "Maximum recursive calls to the PCRE match function.")
	rootCmd.Flags().Bool("phishing-scan-urls", false, "Enable URL signature-based phishing detection")
	rootCmd.Flags().Bool("phishing-sigs", false, "Enable email signature-based phishing detection")
	rootCmd.Flags().Bool("quiet", false, "Only output error messages")
	rootCmd.Flags().BoolP("recursive", "r", false, "Scan subdirectories recursively")
	rootCmd.Flags().Bool("remove", false, "Remove infected files. Be careful!")
	rootCmd.Flags().Bool("scan-archive", false, "Scan archive files (supported by libclamav)")
	rootCmd.Flags().Bool("scan-elf", false, "Scan ELF files")
	rootCmd.Flags().Bool("scan-html", false, "Scan HTML files")
	rootCmd.Flags().Bool("scan-hwp3", false, "Scan HWP3 files")
	rootCmd.Flags().Bool("scan-mail", false, "Scan mail files")
	rootCmd.Flags().Bool("scan-ole2", false, "Scan OLE2 containers")
	rootCmd.Flags().Bool("scan-pdf", false, "Scan PDF files")
	rootCmd.Flags().Bool("scan-pe", false, "Scan PE files")
	rootCmd.Flags().Bool("scan-swf", false, "Scan SWF files")
	rootCmd.Flags().Bool("scan-xmldocs", false, "Scan xml-based document files")
	rootCmd.Flags().Bool("statistics", false, "Collect and print execution statistics")
	rootCmd.Flags().Bool("stdout", false, "Write to stdout instead of stderr. Does not affect 'debug' messages.")
	rootCmd.Flags().String("structured-cc-count", "", "Min CC count to generate a detect")
	rootCmd.Flags().String("structured-cc-mode", "", "CC mode (0=credit debit and private label, 1=credit cards only")
	rootCmd.Flags().String("structured-ssn-count", "", "Min SSN count to generate a detect")
	rootCmd.Flags().String("structured-ssn-format", "", "SSN format (0=normal,1=stripped,2=both)")
	rootCmd.Flags().BoolP("suppress-ok-results", "o", false, "Skip printing OK files")
	rootCmd.Flags().String("tempdir", "", "Create temporary files in DIRECTORY")
	rootCmd.Flags().BoolP("verbose", "v", false, "Be verbose")
	rootCmd.Flags().BoolP("version", "V", false, "Print version number")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"copy":      carapace.ActionDirectories(),
		"database":  carapace.ActionFiles(),
		"file-list": carapace.ActionFiles(),
		"log":       carapace.ActionFiles(),
		"move":      carapace.ActionDirectories(),
		"tempdir":   carapace.ActionDirectories(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
