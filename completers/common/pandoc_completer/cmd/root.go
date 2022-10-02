package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/pandoc_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/net/http"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pandoc",
	Short: "general markup converter",
	Long:  "https://pandoc.org/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("abbreviations", "", "Specifies a custom abbreviations file")
	rootCmd.Flags().Bool("ascii", false, "Use only ASCII characters in output")
	rootCmd.Flags().Bool("bash-completion", false, "Generate a bash completion script")
	rootCmd.Flags().Bool("biblatex", false, "Use biblatex for citations in LaTeX output")
	rootCmd.Flags().String("bibliography", "", "Set the bibliography field in the document’s metadata")
	rootCmd.Flags().String("citation-abbreviations", "", "Set the citation-abbreviations field in the document’s metadata")
	rootCmd.Flags().String("citeproc", "", "Process the citations in the file")
	rootCmd.Flags().String("columns", "", "Specify length of lines in characters")
	rootCmd.Flags().String("csl", "", "Set the csl field in the document’s metadat")
	rootCmd.Flags().String("data-dir", "", "Specify the user data directory")
	rootCmd.Flags().String("default-image-extension", "", "Specify a default extension to use when image paths/URLs have no extension")
	rootCmd.Flags().StringP("defaults", "d", "", "Specify a set of default option settings")
	rootCmd.Flags().String("dpi", "", "Specify the default dpi value for conversion")
	rootCmd.Flags().Bool("dump-args", false, "Print information about command-line arguments to stdout")
	rootCmd.Flags().String("email-obfuscation", "", "Specify a method for obfuscating mailto")
	rootCmd.Flags().String("eol", "", "Manually specify line endings")
	rootCmd.Flags().String("epub-chapter-level", "", "Specify the heading level at which to split the EPUB into separate files")
	rootCmd.Flags().String("epub-cover-image", "", "Use the specified image as the EPUB cover")
	rootCmd.Flags().String("epub-embed-font", "", "Embed the specified font in the EPUB")
	rootCmd.Flags().String("epub-metadata", "", "Look in the specified XML file for metadata for the EPUB")
	rootCmd.Flags().String("epub-subdirectory", "", "Specify the subdirectory that is to hold the EPUB-specific  contents")
	rootCmd.Flags().String("extract-media", "", "Extract images and other media contained in or linked from the source document")
	rootCmd.Flags().Bool("fail-if-warnings", false, "Exit with error status if there are any warnings")
	rootCmd.Flags().Bool("file-scope", false, "Parse each file individually before combining for multifile documents")
	rootCmd.Flags().StringP("filter", "F", "", "Specify  an  executable  to  be used as a filter transforming the pandoc AST")
	rootCmd.Flags().StringP("from", "f", "", "Specify input format")
	rootCmd.Flags().Bool("gladtex", false, "Enclose TeX math in <eq> tags in HTML output")
	rootCmd.Flags().BoolP("help", "h", false, "Show usage message")
	rootCmd.Flags().String("highlight-style", "", "Specifies the coloring style to be used in highlighted source code")
	rootCmd.Flags().Bool("html-q-tags", false, "Use <q> tags for quotes in HTML")
	rootCmd.Flags().String("id-prefix", "", "Specify a prefix to be added to all identifiers and internal links")
	rootCmd.Flags().Bool("ignore-args", false, "Ignore command-line arguments")
	rootCmd.Flags().StringP("include-after-body", "A", "", "Include contents of FILE after body")
	rootCmd.Flags().StringP("include-before-body", "B", "", "Include contents of FILE before body")
	rootCmd.Flags().StringP("include-in-header", "H", "", "Include contents of FILE in header")
	rootCmd.Flags().String("incremental", "", "Make list items in slide shows display incrementally")
	rootCmd.Flags().String("indented-code-classes", "", "Specify classes to use for indented code blocks")
	rootCmd.Flags().String("ipynb-output", "", "Determines how ipynb output cells are treated")
	rootCmd.Flags().String("katex", "", "Use  KaTeX to display embedded TeX math in HTML output")
	rootCmd.Flags().String("list-extensions", "", "List  supported  extensions for FORMAT")
	rootCmd.Flags().Bool("list-highlight-languages", false, "List supported languages for syntax highlighting")
	rootCmd.Flags().Bool("list-highlight-styles", false, "List supported styles for syntax highlighting")
	rootCmd.Flags().Bool("list-input-formats", false, "List supported input formats")
	rootCmd.Flags().Bool("list-output-formats", false, "List supported output formats")
	rootCmd.Flags().Bool("listings", false, "Use the listings package for LaTeX code blocks")
	rootCmd.Flags().String("log", "", "Write log messages in machine-readable JSON format to FILE")
	rootCmd.Flags().StringP("lua-filter", "L", "", "Transform the document in a similar fashion as JSON filters")
	rootCmd.Flags().String("markdown-headings", "", "Specify heading style")
	rootCmd.Flags().String("mathjax", "", "Use  MathJax to display embedded TeX math in HTML output")
	rootCmd.Flags().Bool("mathml", false, "Convert TeX math to MathML")
	rootCmd.Flags().String("metadata-file", "", "Read metadata from the supplied YAML (or JSON) file")
	rootCmd.Flags().Bool("natbib", false, "Use natbib for citations in LaTeX output")
	rootCmd.Flags().Bool("no-check-certificate", false, "Disable the certificate verification to allow access to unsecure HTTP resources")
	rootCmd.Flags().Bool("no-highlight", false, "Disables syntax highlighting for code blocks and inlines")
	rootCmd.Flags().String("number-offset", "", "Offset for section headings in HTML output")
	rootCmd.Flags().String("number-sections", "", "Number section headings")
	rootCmd.Flags().StringP("output", "o", "", "Write output to FILE instead of stdout")
	rootCmd.Flags().String("pdf-engine", "", "Use the specified engine when producing PDF output")
	rootCmd.Flags().String("pdf-engine-opt", "", "Use the given string as a command-line argument to the pdf-engine")
	rootCmd.Flags().BoolP("preserve-tabs", "p", false, "Preserve  tabs  instead  of  converting them to spaces")
	rootCmd.Flags().String("print-default-data-file", "", "Print a system default data file")
	rootCmd.Flags().StringP("print-default-template", "D", "", "Print the system default template for an output FORMAT")
	rootCmd.Flags().String("print-highlight-style", "", "")
	rootCmd.Flags().Bool("quiet", false, "Suppress warning messages.")
	rootCmd.Flags().StringP("read", "r", "", "Specify input format")
	rootCmd.Flags().String("reference-doc", "", "Use the specified file as a style reference")
	rootCmd.Flags().Bool("reference-links", false, "Use reference-style links")
	rootCmd.Flags().String("reference-location", "", "Specify whether footnotes are placed")
	rootCmd.Flags().String("request-header", "", "Set the request header NAME to the value VAL when making HTTP requests")
	rootCmd.Flags().String("resource-path", "", "List of paths to search for images and other resources")
	rootCmd.Flags().Bool("section-divs", false, "Wrap sections in <section> tags")
	rootCmd.Flags().Bool("self-contained", false, "Produce a standalone HTML file with no external dependencies")
	rootCmd.Flags().String("shift-heading-level-by", "", "Shift heading levels by a positive or negative integer")
	rootCmd.Flags().String("slide-level", "", "Specifies that headings with the specified level create slides")
	rootCmd.Flags().String("standalone", "", "Produce  output with an appropriate header and footer")
	rootCmd.Flags().Bool("strip-comments", false, "Strip out HTML comments in the Markdown or Textile source")
	rootCmd.Flags().String("syntax-definition", "", "Instructs pandoc to load a KDE XML syntax definition file")
	rootCmd.Flags().String("tab-stop", "", "Specify the number of spaces per tab (default is 4)")
	rootCmd.Flags().Bool("table-of-contents", false, "Include  an automatically generated table of contents")
	rootCmd.Flags().String("template", "", "Use the specified file as a custom template for the  generated  document")
	rootCmd.Flags().StringP("to", "t", "", "Specify output format")
	rootCmd.Flags().Bool("toc", false, "Include  an automatically generated table of contents")
	rootCmd.Flags().String("toc-depth", "", "Specify the number of section levels to include in the table of  contents")
	rootCmd.Flags().String("top-level-division", "", "Treat top-level headings as the given division type")
	rootCmd.Flags().String("track-changes", "", "Specifies what to do with insertions, deletions, and comments")
	rootCmd.Flags().Bool("verbose", false, "Give verbose debugging output.")
	rootCmd.Flags().BoolP("version", "v", false, "Print version")
	rootCmd.Flags().String("webtex", "", "Convert TeX formulas to <img> tags")
	rootCmd.Flags().String("wrap", "", "Determine how text is wrapped in the output")
	rootCmd.Flags().StringP("write", "w", "", "Specify output format")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		// TODO intended-code-classes
		"abbreviations":           carapace.ActionFiles(),
		"bibliography":            carapace.ActionFiles(),
		"citation-abbreviations":  carapace.ActionFiles(),
		"csl":                     carapace.ActionFiles(),
		"data-dir":                carapace.ActionDirectories(),
		"defaults":                carapace.ActionFiles(),
		"email-obfuscation":       carapace.ActionValues("none", "javascript", "references"),
		"eol":                     carapace.ActionValues("crlf", "lf", "native"),
		"epub-cover-image":        carapace.ActionFiles(),
		"epub-embed-font":         carapace.ActionFiles(),
		"epub-metadata":           carapace.ActionFiles(),
		"epub-subdirectory":       carapace.ActionDirectories(),
		"extract-media":           carapace.ActionDirectories(),
		"filter":                  carapace.ActionFiles(),
		"from":                    action.ActionInputFormats(),
		"highlight-style":         action.ActionHighlightStyles(),
		"include-after-body":      carapace.ActionFiles(),
		"include-before-body":     carapace.ActionFiles(),
		"include-in-header":       carapace.ActionFiles(),
		"ipynb-output":            carapace.ActionValues("all", "none", "best"),
		"list-extensions":         action.ActionFormats(),
		"log":                     carapace.ActionFiles(),
		"lua-filter":              carapace.ActionFiles(".lua"),
		"markdown-headings":       carapace.ActionValues("setext", "atx"),
		"metadata-file":           carapace.ActionFiles(),
		"output":                  carapace.ActionFiles(),
		"pdf-engine":              carapace.ActionFiles(),
		"print-default-data-file": carapace.ActionFiles(),
		"print-default-template":  action.ActionFormats(),
		"print-highlight-style":   action.ActionHighlightStyles(),
		"read":                    action.ActionInputFormats(),
		"reference-doc":           carapace.ActionFiles(),
		"reference-location":      carapace.ActionValues("block", "section", "document"),
		"request-header": carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return http.ActionRequestHeaderNames().Invoke(c).Suffix(":").ToA()
			case 1:
				return http.ActionRequestHeaderValues(c.Parts[0])
			default:
				return carapace.ActionValues()
			}
		}),
		"resource-path":      carapace.ActionDirectories(),
		"syntax-definition":  carapace.ActionFiles(),
		"template":           carapace.ActionFiles(),
		"to":                 action.ActionOutputFormats(),
		"top-level-division": carapace.ActionValues("default", "section", "chapter", "part"),
		"track-changes":      carapace.ActionValues("accept", "reject", "all"),
		"wrap":               carapace.ActionValues("auto", "none", "preserve"),
		"write":              action.ActionOutputFormats(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
