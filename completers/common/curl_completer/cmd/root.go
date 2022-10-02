package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/curl_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/net"
	"github.com/rsteube/carapace-bin/pkg/actions/net/http"
	"github.com/rsteube/carapace-bin/pkg/actions/net/ssh"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "curl",
	Short: "transfer a URL",
	Long:  "https://github.com/curl/curl",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("abstract-unix-socket", "", "Connect via abstract Unix domain socket")
	rootCmd.Flags().String("alt-svc", "", "Enable alt-svc with this cache file")
	rootCmd.Flags().Bool("anyauth", false, "Pick any authentication method")
	rootCmd.Flags().BoolP("append", "a", false, "Append to target file when uploading")
	rootCmd.Flags().Bool("basic", false, "Use HTTP Basic Authentication")
	rootCmd.Flags().String("cacert", "", "CA certificate to verify peer against")
	rootCmd.Flags().String("capath", "", "CA directory to verify peer against")
	rootCmd.Flags().StringP("cert", "E", "", "Client certificate file and password")
	rootCmd.Flags().Bool("cert-status", false, "Verify the status of the server certificate")
	rootCmd.Flags().String("cert-type", "", "Certificate file type (DER/PEM/ENG)")
	rootCmd.Flags().String("ciphers", "", "of ciphers> SSL ciphers to use")
	rootCmd.Flags().Bool("compressed", false, "Request compressed response")
	rootCmd.Flags().String("compressed-ssh", "", "SSH compression")
	rootCmd.Flags().StringP("config", "K", "", "Read config from a file")
	rootCmd.Flags().String("connect-timeout", "", "Maximum time allowed for connection")
	rootCmd.Flags().String("connect-to", "", "Connect to host")
	rootCmd.Flags().StringP("continue-at", "C", "", "Resumed transfer offset")
	rootCmd.Flags().StringP("cookie", "b", "", "Send cookies from string/file")
	rootCmd.Flags().StringP("cookie-jar", "c", "", "Write cookies to <filename> after operation")
	rootCmd.Flags().Bool("create-dirs", false, "Create necessary local directory hierarchy")
	rootCmd.Flags().Bool("crlf", false, "Convert LF to CRLF in upload")
	rootCmd.Flags().String("crlfile", "", "Get a CRL list in PEM format from the given file")
	rootCmd.Flags().StringP("data", "d", "", "HTTP POST data")
	rootCmd.Flags().String("data-ascii", "", "HTTP POST ASCII data")
	rootCmd.Flags().String("data-binary", "", "HTTP POST binary data")
	rootCmd.Flags().String("data-raw", "", "HTTP POST data, '@' allowed")
	rootCmd.Flags().String("data-urlencode", "", "HTTP POST data url encoded")
	rootCmd.Flags().String("delegation", "", "GSS-API delegation permission")
	rootCmd.Flags().Bool("digest", false, "Use HTTP Digest Authentication")
	rootCmd.Flags().BoolP("disable", "q", false, "Disable .curlrc")
	rootCmd.Flags().Bool("disable-eprt", false, "Inhibit using EPRT or LPRT")
	rootCmd.Flags().Bool("disable-epsv", false, "Inhibit using EPSV")
	rootCmd.Flags().String("disallow-username-in-url", "", "username in url")
	rootCmd.Flags().String("dns-interface", "", "Interface to use for DNS requests")
	rootCmd.Flags().String("dns-ipv4-addr", "", "IPv4 address to use for DNS requests")
	rootCmd.Flags().String("dns-ipv6-addr", "", "IPv6 address to use for DNS requests")
	rootCmd.Flags().String("dns-servers", "", "DNS server addrs to use")
	rootCmd.Flags().String("doh-url", "", "Resolve host names over DOH")
	rootCmd.Flags().StringP("dump-header", "D", "", "Write the received headers to <filename>")
	rootCmd.Flags().String("egd-file", "", "EGD socket path for random data")
	rootCmd.Flags().String("engine", "", "Crypto engine to use")
	rootCmd.Flags().String("etag-compare", "", "Get an ETag from a file and send a conditional request")
	rootCmd.Flags().String("etag-save", "", "Get an ETag from response header and save it to a FILE")
	rootCmd.Flags().String("expect100-timeout", "", "How long to wait for 100-continue")
	rootCmd.Flags().BoolP("fail", "f", false, "Fail silently (no output at all) on HTTP errors")
	rootCmd.Flags().Bool("fail-early", false, "Fail on first transfer error, do not continue")
	rootCmd.Flags().Bool("false-start", false, "Enable TLS False Start")
	rootCmd.Flags().StringP("form", "F", "", "Specify multipart MIME data")
	rootCmd.Flags().String("form-string", "", "Specify multipart MIME data")
	rootCmd.Flags().String("ftp-account", "", "Account data string")
	rootCmd.Flags().String("ftp-alternative-to-user", "", "String to replace USER [name]")
	rootCmd.Flags().String("ftp-create-dirs", "", "the remote dirs if not present")
	rootCmd.Flags().String("ftp-method", "", "Control CWD usage")
	rootCmd.Flags().Bool("ftp-pasv", false, "Use PASV/EPSV instead of PORT")
	rootCmd.Flags().StringP("ftp-port", "P", "", "Use PORT instead of PASV")
	rootCmd.Flags().Bool("ftp-pret", false, "Send PRET before PASV")
	rootCmd.Flags().String("ftp-skip-pasv-ip", "", "the IP address for PASV")
	rootCmd.Flags().Bool("ftp-ssl-ccc", false, "Send CCC after authenticating")
	rootCmd.Flags().String("ftp-ssl-ccc-mode", "", "Set CCC mode")
	rootCmd.Flags().String("ftp-ssl-control", "", "SSL/TLS for FTP login, clear for transfer")
	rootCmd.Flags().BoolP("get", "G", false, "Put the post data in the URL and use GET")
	rootCmd.Flags().BoolP("globoff", "g", false, "Disable URL sequences and ranges using {} and []")
	rootCmd.Flags().String("happy-eyeballs-timeout-ms", "", "How long to wait in milliseconds for IPv6 before trying IPv4")
	rootCmd.Flags().String("haproxy-protocol", "", "HAProxy PROXY protocol v1 header")
	rootCmd.Flags().BoolP("head", "I", false, "Show document info only")
	rootCmd.Flags().StringArrayP("header", "H", []string{}, "Pass custom header(s) to server")
	rootCmd.Flags().BoolP("help", "h", false, "This help text")
	rootCmd.Flags().String("hostpubmd5", "", "Acceptable MD5 hash of the host public key")
	rootCmd.Flags().Bool("http0.9", false, "Allow HTTP 0.9 responses")
	rootCmd.Flags().BoolP("http1.0", "0", false, "Use HTTP 1.0")
	rootCmd.Flags().Bool("http1.1", false, "Use HTTP 1.1")
	rootCmd.Flags().Bool("http2", false, "Use HTTP 2")
	rootCmd.Flags().String("http2-prior-knowledge", "", "HTTP 2 without HTTP/1.1 Upgrade")
	rootCmd.Flags().Bool("http3", false, "Use HTTP v3")
	rootCmd.Flags().String("ignore-content-length", "", "the size of the remote resource")
	rootCmd.Flags().BoolP("include", "i", false, "Include protocol response headers in the output")
	rootCmd.Flags().BoolP("insecure", "k", false, "Allow insecure server connections when using SSL")
	rootCmd.Flags().String("interface", "", "Use network INTERFACE (or address)")
	rootCmd.Flags().BoolP("ipv4", "4", false, "Resolve names to IPv4 addresses")
	rootCmd.Flags().BoolP("ipv6", "6", false, "Resolve names to IPv6 addresses")
	rootCmd.Flags().StringP("junk-session-cookies", "j", "", "session cookies read from file")
	rootCmd.Flags().String("keepalive-time", "", "Interval time for keepalive probes")
	rootCmd.Flags().String("key", "", "Private key file name")
	rootCmd.Flags().String("key-type", "", "Private key file type (DER/PEM/ENG)")
	rootCmd.Flags().String("krb", "", "Enable Kerberos with security <level>")
	rootCmd.Flags().String("libcurl", "", "Dump libcurl equivalent code of this command line")
	rootCmd.Flags().String("limit-rate", "", "Limit transfer speed to RATE")
	rootCmd.Flags().BoolP("list-only", "l", false, "List only mode")
	rootCmd.Flags().String("local-port", "", "Force use of RANGE for local port numbers")
	rootCmd.Flags().BoolP("location", "L", false, "Follow redirects")
	rootCmd.Flags().String("location-trusted", "", "--location, and send auth to other hosts")
	rootCmd.Flags().String("login-options", "", "Server login options")
	rootCmd.Flags().String("mail-auth", "", "Originator address of the original email")
	rootCmd.Flags().String("mail-from", "", "Mail from this address")
	rootCmd.Flags().String("mail-rcpt", "", "Mail to this address")
	rootCmd.Flags().String("mail-rcpt-allowfails", "", "RCPT TO command to fail for some recipients")
	rootCmd.Flags().BoolP("manual", "M", false, "Display the full manual")
	rootCmd.Flags().String("max-filesize", "", "Maximum file size to download")
	rootCmd.Flags().String("max-redirs", "", "Maximum number of redirects allowed")
	rootCmd.Flags().StringP("max-time", "m", "", "Maximum time allowed for the transfer")
	rootCmd.Flags().Bool("metalink", false, "Process given URLs as metalink XML file")
	rootCmd.Flags().Bool("negotiate", false, "Use HTTP Negotiate (SPNEGO) authentication")
	rootCmd.Flags().BoolP("netrc", "n", false, "Must read .netrc for user name and password")
	rootCmd.Flags().String("netrc-file", "", "Specify FILE for netrc")
	rootCmd.Flags().String("netrc-optional", "", "either .netrc or URL")
	rootCmd.Flags().BoolP("next", ":", false, "Make next URL use its separate set of options")
	rootCmd.Flags().Bool("no-alpn", false, "Disable the ALPN TLS extension")
	rootCmd.Flags().BoolP("no-buffer", "N", false, "Disable buffering of the output stream")
	rootCmd.Flags().Bool("no-keepalive", false, "Disable TCP keepalive on the connection")
	rootCmd.Flags().Bool("no-npn", false, "Disable the NPN TLS extension")
	rootCmd.Flags().String("no-progress-meter", "", "not show the progress meter")
	rootCmd.Flags().Bool("no-sessionid", false, "Disable SSL session-ID reusing")
	rootCmd.Flags().String("noproxy", "", "List of hosts which do not use proxy")
	rootCmd.Flags().Bool("ntlm", false, "Use HTTP NTLM authentication")
	rootCmd.Flags().Bool("ntlm-wb", false, "Use HTTP NTLM authentication with winbind")
	rootCmd.Flags().String("oauth2-bearer", "", "OAuth 2 Bearer Token")
	rootCmd.Flags().StringP("output", "o", "", "Write to file instead of stdout")
	rootCmd.Flags().String("output-dir", "", "specifies the directory in which files should be stored")
	rootCmd.Flags().BoolP("parallel", "Z", false, "Perform transfers in parallel")
	rootCmd.Flags().String("parallel-immediate", "", "not wait for multiplexing (with --parallel)")
	rootCmd.Flags().Bool("parallel-max", false, "Maximum concurrency for parallel transfers")
	rootCmd.Flags().String("pass", "", "Pass phrase for the private key")
	rootCmd.Flags().Bool("path-as-is", false, "Do not squash .. sequences in URL path")
	rootCmd.Flags().String("pinnedpubkey", "", "FILE/HASHES Public key to verify peer against")
	rootCmd.Flags().Bool("post301", false, "Do not switch to GET after following a 301")
	rootCmd.Flags().Bool("post302", false, "Do not switch to GET after following a 302")
	rootCmd.Flags().Bool("post303", false, "Do not switch to GET after following a 303")
	rootCmd.Flags().String("preproxy", "", "Use this proxy first")
	rootCmd.Flags().BoolP("progress-bar", "#", false, "Display transfer progress as a bar")
	rootCmd.Flags().String("proto", "", "Enable/disable PROTOCOLS")
	rootCmd.Flags().String("proto-default", "", "Use PROTOCOL for any URL missing a scheme")
	rootCmd.Flags().String("proto-redir", "", "Enable/disable PROTOCOLS on redirect")
	rootCmd.Flags().StringP("proxy", "x", "", "Use this proxy")
	rootCmd.Flags().String("proxy-anyauth", "", "any proxy authentication method")
	rootCmd.Flags().Bool("proxy-basic", false, "Use Basic authentication on the proxy")
	rootCmd.Flags().String("proxy-cacert", "", "CA certificate to verify peer against for proxy")
	rootCmd.Flags().String("proxy-capath", "", "CA directory to verify peer against for proxy")
	rootCmd.Flags().String("proxy-cert", "", "Set client certificate for proxy")
	rootCmd.Flags().String("proxy-cert-type", "", "Client certificate type for HTTPS proxy")
	rootCmd.Flags().String("proxy-ciphers", "", "SSL ciphers to use for proxy")
	rootCmd.Flags().String("proxy-crlfile", "", "Set a CRL list for proxy")
	rootCmd.Flags().Bool("proxy-digest", false, "Use Digest authentication on the proxy")
	rootCmd.Flags().String("proxy-header", "", "Pass custom header(s) to proxy")
	rootCmd.Flags().String("proxy-insecure", "", "HTTPS proxy connections without verifying the proxy")
	rootCmd.Flags().String("proxy-key", "", "Private key for HTTPS proxy")
	rootCmd.Flags().String("proxy-key-type", "", "Private key file type for proxy")
	rootCmd.Flags().String("proxy-negotiate", "", "HTTP Negotiate (SPNEGO) authentication on the proxy")
	rootCmd.Flags().Bool("proxy-ntlm", false, "Use NTLM authentication on the proxy")
	rootCmd.Flags().String("proxy-pass", "", "Pass phrase for the private key for HTTPS proxy")
	rootCmd.Flags().String("proxy-pinnedpubkey", "", "FILE/HASHES public key to verify proxy with")
	rootCmd.Flags().String("proxy-service-name", "", "SPNEGO proxy service name")
	rootCmd.Flags().String("proxy-ssl-allow-beast", "", "security flaw for interop for HTTPS proxy")
	rootCmd.Flags().String("proxy-tls13-ciphers", "", "TLS 1.3 ciphersuites for proxy (OpenSSL)")
	rootCmd.Flags().String("proxy-tlsauthtype", "", "TLS authentication type for HTTPS proxy")
	rootCmd.Flags().String("proxy-tlspassword", "", "TLS password for HTTPS proxy")
	rootCmd.Flags().String("proxy-tlsuser", "", "TLS username for HTTPS proxy")
	rootCmd.Flags().Bool("proxy-tlsv1", false, "Use TLSv1 for HTTPS proxy")
	rootCmd.Flags().StringP("proxy-user", "U", "", "Proxy user and password")
	rootCmd.Flags().String("proxy1.0", "", "Use HTTP/1.0 proxy on given port")
	rootCmd.Flags().BoolP("proxytunnel", "p", false, "Operate through an HTTP proxy tunnel (using CONNECT)")
	rootCmd.Flags().String("pubkey", "", "SSH Public key file name")
	rootCmd.Flags().BoolP("quote", "Q", false, "Send command(s) to server before transfer")
	rootCmd.Flags().String("random-file", "", "File for reading random data from")
	rootCmd.Flags().StringP("range", "r", "", "Retrieve only the bytes within RANGE")
	rootCmd.Flags().Bool("raw", false, "Do HTTP \"raw\"; no transfer decoding")
	rootCmd.Flags().StringP("referer", "e", "", "Referrer URL")
	rootCmd.Flags().StringP("remote-header-name", "J", "", "the header-provided filename")
	rootCmd.Flags().BoolP("remote-name", "O", false, "Write output to a file named as the remote file")
	rootCmd.Flags().String("remote-name-all", "", "the remote file name for all URLs")
	rootCmd.Flags().BoolP("remote-time", "R", false, "Set the remote file's time on the local output")
	rootCmd.Flags().StringP("request", "X", "", "Specify request command to use")
	rootCmd.Flags().String("request-target", "", "the target for this request")
	rootCmd.Flags().String("resolve", "", "Resolve the host+port to this address")
	rootCmd.Flags().String("retry", "", "Retry request if transient problems occur")
	rootCmd.Flags().String("retry-connrefused", "", "on connection refused (use with --retry)")
	rootCmd.Flags().String("retry-delay", "", "Wait time between retries")
	rootCmd.Flags().String("retry-max-time", "", "Retry only within this period")
	rootCmd.Flags().String("sasl-authzid", "", "Use this identity to act as during SASL PLAIN authentication")
	rootCmd.Flags().Bool("sasl-ir", false, "Enable initial response in SASL authentication")
	rootCmd.Flags().String("service-name", "", "SPNEGO service name")
	rootCmd.Flags().BoolP("show-error", "S", false, "Show error even when -s is used")
	rootCmd.Flags().BoolP("silent", "s", false, "Silent mode")
	rootCmd.Flags().String("socks4", "", "SOCKS4 proxy on given host + port")
	rootCmd.Flags().String("socks4a", "", "SOCKS4a proxy on given host + port")
	rootCmd.Flags().String("socks5", "", "SOCKS5 proxy on given host + port")
	rootCmd.Flags().Bool("socks5-basic", false, "Enable username/password auth for SOCKS5 proxies")
	rootCmd.Flags().String("socks5-gssapi", "", "GSS-API auth for SOCKS5 proxies")
	rootCmd.Flags().String("socks5-gssapi-nec", "", "with NEC SOCKS5 server")
	rootCmd.Flags().String("socks5-gssapi-service", "", "SOCKS5 proxy service name for GSS-API")
	rootCmd.Flags().String("socks5-hostname", "", "SOCKS5 proxy, pass host name to proxy")
	rootCmd.Flags().StringP("speed-limit", "Y", "", "Stop transfers slower than this")
	rootCmd.Flags().StringP("speed-time", "y", "", "Trigger 'speed-limit' abort after this time")
	rootCmd.Flags().Bool("ssl", false, "Try SSL/TLS")
	rootCmd.Flags().String("ssl-allow-beast", "", "security flaw to improve interop")
	rootCmd.Flags().String("ssl-no-revoke", "", "cert revocation checks (Schannel)")
	rootCmd.Flags().Bool("ssl-reqd", false, "Require SSL/TLS")
	rootCmd.Flags().String("ssl-revoke-best-effort", "", "revocation offline or missing revocation list errors (Schannel)")
	rootCmd.Flags().BoolP("sslv2", "2", false, "Use SSLv2")
	rootCmd.Flags().BoolP("sslv3", "3", false, "Use SSLv3")
	rootCmd.Flags().Bool("stderr", false, "Where to redirect stderr")
	rootCmd.Flags().String("styled-output", "", "styled output for HTTP headers")
	rootCmd.Flags().String("suppress-connect-headers", "", "proxy CONNECT response headers")
	rootCmd.Flags().Bool("tcp-fastopen", false, "Use TCP Fast Open")
	rootCmd.Flags().Bool("tcp-nodelay", false, "Use the TCP_NODELAY option")
	rootCmd.Flags().StringP("telnet-option", "t", "", "Set telnet option")
	rootCmd.Flags().String("tftp-blksize", "", "Set TFTP BLKSIZE option")
	rootCmd.Flags().String("tftp-no-options", "", "not send any TFTP options")
	rootCmd.Flags().StringP("time-cond", "z", "", "Transfer based on a time condition")
	rootCmd.Flags().String("tls-max", "", "Set maximum allowed TLS version")
	rootCmd.Flags().String("tls13-ciphers", "", "TLS 1.3 ciphersuites (OpenSSL)")
	rootCmd.Flags().String("tlsauthtype", "", "TLS authentication type")
	rootCmd.Flags().Bool("tlspassword", false, "TLS password")
	rootCmd.Flags().String("tlsuser", "", "TLS user name")
	rootCmd.Flags().BoolP("tlsv1", "1", false, "Use TLSv1.0 or greater")
	rootCmd.Flags().Bool("tlsv1.0", false, "Use TLSv1.0 or greater")
	rootCmd.Flags().Bool("tlsv1.1", false, "Use TLSv1.1 or greater")
	rootCmd.Flags().Bool("tlsv1.2", false, "Use TLSv1.2 or greater")
	rootCmd.Flags().Bool("tlsv1.3", false, "Use TLSv1.3 or greater")
	rootCmd.Flags().Bool("tr-encoding", false, "Request compressed transfer encoding")
	rootCmd.Flags().String("trace", "", "Write a debug trace to FILE")
	rootCmd.Flags().String("trace-ascii", "", "Like --trace, but without hex output")
	rootCmd.Flags().Bool("trace-time", false, "Add time stamps to trace/verbose output")
	rootCmd.Flags().String("unix-socket", "", "Connect through this Unix domain socket")
	rootCmd.Flags().StringP("upload-file", "T", "", "Transfer local FILE to destination")
	rootCmd.Flags().String("url", "", "URL to work with")
	rootCmd.Flags().BoolP("use-ascii", "B", false, "Use ASCII/text transfer")
	rootCmd.Flags().StringP("user", "u", "", "Server user and password")
	rootCmd.Flags().StringP("user-agent", "A", "", "Send User-Agent <name> to server")
	rootCmd.Flags().BoolP("verbose", "v", false, "Make the operation more talkative")
	rootCmd.Flags().BoolP("version", "V", false, "Show version number and quit")
	rootCmd.Flags().StringP("write-out", "w", "", "Use output FORMAT after completion")
	rootCmd.Flags().Bool("xattr", false, "Store metadata in extended file attributes")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"alt-svc":   carapace.ActionFiles(),
		"cacert":    carapace.ActionFiles(),
		"capath":    carapace.ActionDirectories(),
		"cert":      carapace.ActionFiles(),
		"cert-type": carapace.ActionValues("DER", "PEM", "ENG"),
		"ciphers": carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			return ssh.ActionCiphers().Invoke(c).Filter(c.Parts).ToA()
		}),
		"config": carapace.ActionFiles(),
		"connect-to": carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return net.ActionHosts().Invoke(c).Suffix(":").ToA()
			case 1:
				return net.ActionPorts().Invoke(c).Suffix(":").ToA()
			case 2:
				return net.ActionHosts().Invoke(c).Suffix(":").ToA()
			case 3:
				return net.ActionPorts()
			default:
				return carapace.ActionValues()
			}
		}),
		"cookie":       carapace.ActionFiles(),
		"cookie-jar":   carapace.ActionFiles(),
		"crlfile":      carapace.ActionFiles(),
		"egd-file":     carapace.ActionFiles(),
		"engine":       action.ActionEngines(),
		"etag-compare": carapace.ActionFiles(),
		"etag-save":    carapace.ActionFiles(),
		"ftp-method": carapace.ActionValuesDescribed(
			"multicwd", "curl does a single CWD operation for each path part in the given URL",
			"nocwd", "curl does no CWD at all",
			"singlecwd", "curl does one CWD with the full target directory and then operates on the file \"normally\"",
		),
		"header": carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return http.ActionRequestHeaderNames().Invoke(c).Suffix(":").ToA()
			case 1:
				return http.ActionRequestHeaderValues(c.Parts[0])
			default:
				return carapace.ActionValues()
			}
		}),
		"key":             carapace.ActionFiles(),
		"key-type":        carapace.ActionValues("DER", "PEM", "ENG"),
		"krb":             carapace.ActionValues("clear", "safe", "confidential", "private"),
		"output":          carapace.ActionFiles(),
		"output-dir":      carapace.ActionDirectories(),
		"proxy-cacert":    carapace.ActionFiles(),
		"proxy-capath":    carapace.ActionDirectories(),
		"proxy-cert":      carapace.ActionFiles(),
		"proxy-cert-type": carapace.ActionValues("DER", "PEM", "ENG"),
		"proxy-ciphers": carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			return ssh.ActionCiphers().Invoke(c).Filter(c.Parts).ToA()
		}),
		"proxy-key":   carapace.ActionFiles(),
		"random-file": carapace.ActionFiles(),
		"trace":       carapace.ActionFiles(),
		"unix-socket": carapace.ActionFiles(),
		"upload-file": carapace.ActionFiles(),
		"user-agent":  http.ActionUserAgents(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		http.ActionUrls(),
	)
}
