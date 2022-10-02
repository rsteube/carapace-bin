package cmd

import (
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/fs"
	"github.com/rsteube/carapace-bin/pkg/actions/net/http"
	"github.com/rsteube/carapace-bin/pkg/actions/net/ssh"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "wget",
	Short: "a non-interactive network retriever",
	Long:  "https://www.gnu.org/software/wget/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("accept", "A", "", "comma-separated list of accepted extensions")
	rootCmd.Flags().String("accept-regex", "", "regex matching accepted URLs")
	rootCmd.Flags().BoolP("adjust-extension", "E", false, "save HTML/CSS documents with proper extensions")
	rootCmd.Flags().StringP("append-output", "a", "", "append messages to FILE")
	rootCmd.Flags().Bool("ask-password", false, "prompt for passwords")
	rootCmd.Flags().Bool("auth-no-challenge", false, "send Basic HTTP authentication information without first waiting for the server's challenge")
	rootCmd.Flags().BoolP("background", "b", false, "go to background after startup")
	rootCmd.Flags().BoolP("backup-converted", "K", false, "before converting file X, back up as X.orig")
	rootCmd.Flags().String("backups", "", "before writing file X, rotate up to N backup files")
	rootCmd.Flags().StringP("base", "B", "", "resolves HTML input-file links relative to URL")
	rootCmd.Flags().String("bind-address", "", "bind to ADDRESS (hostname or IP) on local host")
	rootCmd.Flags().String("body-data", "", "send STRING as data. --method MUST be set")
	rootCmd.Flags().String("body-file", "", "send contents of FILE. --method MUST be set")
	rootCmd.Flags().String("ca-certificate", "", "file with the bundle of CAs")
	rootCmd.Flags().String("ca-directory", "", "directory where hash list of CAs is stored")
	rootCmd.Flags().String("certificate", "", "client certificate file")
	rootCmd.Flags().String("certificate-type", "", "client certificate type, PEM or DER")
	rootCmd.Flags().String("ciphers", "", "Set the priority string (GnuTLS) or cipher list string (OpenSSL) directly.")
	rootCmd.Flags().String("compression", "", "choose compression")
	rootCmd.Flags().String("config", "", "specify config file to use")
	rootCmd.Flags().String("connect-timeout", "", "set the connect timeout to SECS")
	rootCmd.Flags().Bool("content-disposition", false, "honor the Content-Disposition header when choosing local file names")
	rootCmd.Flags().Bool("content-on-error", false, "output the received content on server errors")
	rootCmd.Flags().BoolP("continue", "c", false, "resume getting a partially-downloaded file")
	rootCmd.Flags().Bool("convert-file-only", false, "convert the file part of the URLs only (usually known as the basename)")
	rootCmd.Flags().BoolP("convert-links", "k", false, "make links in downloaded HTML or CSS point to local files")
	rootCmd.Flags().String("crl-file", "", "file with bundle of CRLs")
	rootCmd.Flags().String("cut-dirs", "", "ignore NUMBER remote directory components")
	rootCmd.Flags().BoolP("debug", "d", false, "print lots of debugging information")
	rootCmd.Flags().String("default-page", "", "change the default page name")
	rootCmd.Flags().Bool("delete-after", false, "delete files locally after downloading them")
	rootCmd.Flags().StringP("directory-prefix", "P", "", "save files to PREFIX/..")
	rootCmd.Flags().String("dns-timeout", "", "set the DNS lookup timeout to SECS")
	rootCmd.Flags().StringP("domains", "D", "", "comma-separated list of accepted domains")
	rootCmd.Flags().StringP("exclude-directories", "X", "", "list of excluded directories")
	rootCmd.Flags().String("exclude-domains", "", "comma-separated list of rejected domains")
	rootCmd.Flags().StringP("execute", "e", "", "execute a `.wgetrc'-style command")
	rootCmd.Flags().Bool("follow-ftp", false, "follow FTP links from HTML documents")
	rootCmd.Flags().String("follow-tags", "", "comma-separated list of followed HTML tags")
	rootCmd.Flags().BoolP("force-directories", "x", false, "force creation of directories")
	rootCmd.Flags().BoolP("force-html", "F", false, "treat input file as HTML")
	rootCmd.Flags().String("ftp-password", "", "set ftp password to PASS")
	rootCmd.Flags().String("ftp-user", "", "set ftp user to USER")
	rootCmd.Flags().Bool("ftps-clear-data-connection", false, "cipher the control channel only; all the data will be in plaintext")
	rootCmd.Flags().Bool("ftps-fallback-to-ftp", false, "fall back to FTP if FTPS is not supported in the target server")
	rootCmd.Flags().Bool("ftps-implicit", false, "use implicit FTPS (default port is 990)")
	rootCmd.Flags().Bool("ftps-resume-ssl", false, "resume the SSL/TLS session started in the control connection")
	rootCmd.Flags().StringArray("header", []string{}, "insert STRING among the headers")
	rootCmd.Flags().BoolP("help", "h", false, "print this help")
	rootCmd.Flags().Bool("hsts-file", false, "path of HSTS database (will override default)")
	rootCmd.Flags().String("http-password", "", "set http password to PASS")
	rootCmd.Flags().String("http-user", "", "set http user to USER")
	rootCmd.Flags().Bool("https-only", false, "only follow secure HTTPS links")
	rootCmd.Flags().Bool("ignore-case", false, "ignore case when matching files/directories")
	rootCmd.Flags().Bool("ignore-length", false, "ignore 'Content-Length' header field")
	rootCmd.Flags().String("ignore-tags", "", "comma-separated list of ignored HTML tags")
	rootCmd.Flags().StringP("include-directories", "I", "", "list of allowed directories")
	rootCmd.Flags().BoolP("inet4-only", "4", false, "connect only to IPv4 addresses")
	rootCmd.Flags().BoolP("inet6-only", "6", false, "connect only to IPv6 addresses")
	rootCmd.Flags().StringP("input-file", "i", "", "download URLs found in local or external FILE")
	rootCmd.Flags().Bool("keep-session-cookies", false, "load and save session (non-permanent) cookies")
	rootCmd.Flags().StringP("level", "l", "", "maximum recursion depth (inf or 0 for infinite)")
	rootCmd.Flags().String("limit-rate", "", "limit download rate to RATE")
	rootCmd.Flags().String("load-cookies", "", "load cookies from FILE before session")
	rootCmd.Flags().String("local-encoding", "", "use ENC as the local encoding for IRIs")
	rootCmd.Flags().Bool("max-redirect", false, "maximum redirections allowed per page")
	rootCmd.Flags().String("method", "", "use method \"HTTPMethod\" in the request")
	rootCmd.Flags().BoolP("mirror", "m", false, "shortcut for -N -r -l inf --no-remove-listing")
	rootCmd.Flags().Bool("no-cache", false, "disallow server-cached data")
	rootCmd.Flags().Bool("no-check-certificate", false, "don't validate the server's certificate")
	rootCmd.Flags().Bool("no-clobber", false, "skip downloads that would download to existing files")
	rootCmd.Flags().Bool("no-config", false, "do not read any config file")
	rootCmd.Flags().Bool("no-cookies", false, "don't use cookies")
	rootCmd.Flags().Bool("no-directories", false, "don't create directories")
	rootCmd.Flags().Bool("no-dns-cache", false, "disable caching DNS lookups")
	rootCmd.Flags().Bool("no-glob", false, "turn off FTP file name globbing")
	rootCmd.Flags().Bool("no-host-directories", false, "don't create host directories")
	rootCmd.Flags().Bool("no-hsts", false, "disable HSTS")
	rootCmd.Flags().Bool("no-http-keep-alive", false, "disable HTTP keep-alive (persistent connections)")
	rootCmd.Flags().Bool("no-if-modified-since", false, "don't use conditional if-modified-since get requests in timestamping mode")
	rootCmd.Flags().Bool("no-iri", false, "turn off IRI support")
	rootCmd.Flags().Bool("no-netrc", false, "don't try to obtain credentials from .netrc")
	rootCmd.Flags().Bool("no-parent", false, "don't ascend to the parent directory")
	rootCmd.Flags().Bool("no-passive-ftp", false, "disable the \"passive\" transfer mode")
	rootCmd.Flags().Bool("no-proxy", false, "explicitly turn off proxy")
	rootCmd.Flags().Bool("no-remove-listing", false, "don't remove '.listing' files")
	rootCmd.Flags().Bool("no-use-server-timestamps", false, "don't set the local file's timestamp by the one on the server")
	rootCmd.Flags().Bool("no-verbose", false, "turn off verboseness, without being quiet")
	rootCmd.Flags().Bool("no-warc-compression", false, "do not compress WARC files with GZIP")
	rootCmd.Flags().Bool("no-warc-digests", false, "do not calculate SHA1 digests")
	rootCmd.Flags().Bool("no-warc-keep-log", false, "do not store the log file in a WARC record")
	rootCmd.Flags().StringP("output-document", "O", "", "write documents to FILE")
	rootCmd.Flags().StringP("output-file", "o", "", "log messages to FILE")
	rootCmd.Flags().BoolP("page-requisites", "p", false, "get all images, etc. needed to display HTML page")
	rootCmd.Flags().String("password", "", "set both ftp and http password to PASS")
	rootCmd.Flags().String("pinnedpubkey", "", "Public key (PEM/DER) file, or any number of base64 encoded sha256 hashes")
	rootCmd.Flags().String("post-data", "", "use the POST method; send STRING as the data")
	rootCmd.Flags().String("post-file", "", "use the POST method; send contents of FILE")
	rootCmd.Flags().String("prefer-family", "", "connect first to addresses of specified family")
	rootCmd.Flags().Bool("preserve-permissions", false, "preserve remote file permissions")
	rootCmd.Flags().String("private-key", "", "private key file")
	rootCmd.Flags().String("private-key-type", "", "private key type, PEM or DER")
	rootCmd.Flags().String("progress", "", "select progress gauge type")
	rootCmd.Flags().Bool("protocol-directories", false, "use protocol name in directories")
	rootCmd.Flags().String("proxy-password", "", "set PASS as proxy password")
	rootCmd.Flags().String("proxy-user", "", "set USER as proxy username")
	rootCmd.Flags().BoolP("quiet", "q", false, "quiet (no output)")
	rootCmd.Flags().StringP("quota", "Q", "", "set retrieval quota to NUMBER")
	rootCmd.Flags().Bool("random-wait", false, "wait from 0.5*WAIT...1.5*WAIT secs between retrievals")
	rootCmd.Flags().String("read-timeout", "", "set the read timeout to SECS")
	rootCmd.Flags().BoolP("recursive", "r", false, "specify recursive download")
	rootCmd.Flags().String("referer", "", "include 'Referer: URL' header in HTTP request")
	rootCmd.Flags().String("regex-type", "", "regex type (posix|pcre)")
	rootCmd.Flags().StringP("reject", "R", "", "comma-separated list of rejected extensions")
	rootCmd.Flags().String("reject-regex", "", "regex matching rejected URLs")
	rootCmd.Flags().String("rejected-log", "", "log reasons for URL rejection to FILE")
	rootCmd.Flags().BoolP("relative", "L", false, "follow relative links only")
	rootCmd.Flags().String("remote-encoding", "", "use ENC as the default remote encoding")
	rootCmd.Flags().String("report-speed", "", "output bandwidth as TYPE.  TYPE can be bits")
	rootCmd.Flags().String("restrict-file-names", "", "restrict chars in file names to ones OS allows")
	rootCmd.Flags().Bool("retr-symlinks", false, "when recursing, get linked-to files (not dir)")
	rootCmd.Flags().Bool("retry-connrefused", false, "retry even if connection is refused")
	rootCmd.Flags().String("retry-on-http-error", "", "comma-separated list of HTTP errors to retry")
	rootCmd.Flags().String("save-cookies", "", "save cookies to FILE after session")
	rootCmd.Flags().Bool("save-headers", false, "save the HTTP headers to file")
	rootCmd.Flags().String("secure-protocol", "", "choose secure protocol, one of auto, SSLv2,")
	rootCmd.Flags().BoolP("server-response", "S", false, "print server response")
	rootCmd.Flags().Bool("show-progress", false, "display the progress bar in any verbosity mode")
	rootCmd.Flags().BoolP("span-hosts", "H", false, "go to foreign hosts when recursive")
	rootCmd.Flags().Bool("spider", false, "don't download anything")
	rootCmd.Flags().String("start-pos", "", "start downloading from zero-based position OFFSET")
	rootCmd.Flags().Bool("strict-comments", false, "turn on strict (SGML) handling of HTML comments")
	rootCmd.Flags().StringP("timeout", "T", "", "set all timeout values to SECONDS")
	rootCmd.Flags().BoolP("timestamping", "N", false, "don't re-retrieve files unless newer than local")
	rootCmd.Flags().StringP("tries", "t", "", "set number of retries to NUMBER (0 unlimits)")
	rootCmd.Flags().Bool("trust-server-names", false, "use the name specified by the redirection URL's last component")
	rootCmd.Flags().Bool("unlink", false, "remove file before clobber")
	rootCmd.Flags().String("use-askpass", "", "specify credential handler for requesting username and password")
	rootCmd.Flags().String("user", "", "set both ftp and http user to USER")
	rootCmd.Flags().StringP("user-agent", "U", "", "identify as AGENT instead of Wget/VERSION")
	rootCmd.Flags().BoolP("verbose", "v", false, "be verbose (this is the default)")
	rootCmd.Flags().BoolP("version", "V", false, "display the version of Wget and exit")
	rootCmd.Flags().StringP("wait", "w", "", "wait SECONDS between retrievals")
	rootCmd.Flags().String("waitretry", "", "wait 1..SECONDS between retries of a retrieval")
	rootCmd.Flags().Bool("warc-cdx", false, "write CDX index files")
	rootCmd.Flags().String("warc-dedup", "", "do not store records listed in this CDX file")
	rootCmd.Flags().String("warc-file", "", "save request/response data to a .warc.gz file")
	rootCmd.Flags().String("warc-header", "", "insert STRING into the warcinfo record")
	rootCmd.Flags().String("warc-max-size", "", "set maximum size of WARC files to NUMBER")
	rootCmd.Flags().String("warc-tempdir", "", "location for temporary files created by the")
	rootCmd.Flags().Bool("xattr", false, "turn on storage of metadata in extended file attributes")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"accept": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return fs.ActionFilenameExtensions().Invoke(c).Filter(c.Parts).ToA()
		}),
		"append-output":    carapace.ActionFiles(),
		"body-file":        carapace.ActionFiles(),
		"ca-certificate":   carapace.ActionFiles(),
		"ca-directory":     carapace.ActionDirectories(),
		"certificate":      carapace.ActionFiles(),
		"certificate-type": carapace.ActionValues("PEM", "DER"),
		"ciphers": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return ssh.ActionCiphers().Invoke(c).Filter(c.Parts).ToA()
		}),
		"compression":      carapace.ActionValues("auto", "gzip", "none"),
		"config":           carapace.ActionFiles(),
		"crl-file":         carapace.ActionFiles(),
		"directory-prefix": carapace.ActionDirectories(),
		"follow-tags": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return http.ActionTags().Invoke(c).Filter(c.Parts).ToA()
		}),
		"header": http.ActionRequestHeaders(),
		"ignore-tags": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return http.ActionTags().Invoke(c).Filter(c.Parts).ToA()
		}),
		"input-file":      carapace.ActionFiles(),
		"load-cookies":    carapace.ActionFiles(),
		"method":          http.ActionRequestMethods(),
		"output-document": carapace.ActionFiles(),
		"output-file":     carapace.ActionFiles(),
		"pinnedpubkey": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if !strings.HasPrefix(c.CallbackValue, "sha256//") {
				return carapace.Batch(
					carapace.ActionFiles(),
					carapace.ActionValues("sha256//"),
				).ToA()
			}
			return carapace.ActionMultiParts(";", func(c carapace.Context) carapace.Action {
				return carapace.ActionValues("sha256//")
			})
		}),
		"post-file":        carapace.ActionFiles(),
		"private-key":      carapace.ActionFiles(),
		"private-key-type": carapace.ActionValues("PEM", "DER"),
		"progress":         carapace.ActionValues("dot", "bar"),
		"regex-type":       carapace.ActionValues("posix", "pcre"),
		"reject": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return fs.ActionFilenameExtensions().Invoke(c).Filter(c.Parts).ToA()
		}),
		"rejected-log": carapace.ActionFiles(),
		"report-speed": carapace.ActionValues("bits"),
		"retry-on-http-error": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return http.ActionStatusCodes().Invoke(c).Filter(c.Parts).ToA()
		}),
		"save-cookies":    carapace.ActionFiles(),
		"secure-protocol": carapace.ActionValues("auto", "SSLv2"),
		"user-agent":      http.ActionUserAgents(),
		"warc-dedup":      carapace.ActionFiles(),
		"warc-file":       carapace.ActionFiles(),
		"warc-tempdir":    carapace.ActionDirectories(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		http.ActionUrls(),
	)
}
