package cmd

import (
	"errors"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/net/http"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "http",
	Short: "command-line HTTP client for the API era",
	Long:  "https://httpie.io/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

// ExecuteHttps replace Use with `https` before execution (https alias command for httpie)
func ExecuteHttps() error {
	rootCmd.Use = "https" // TODO works for now
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("all", false, "Show any intermediary requests/responses as well.")
	rootCmd.Flags().StringP("auth", "a", "", "If only the username is provided (-a username), HTTPie will prompt for the password.")
	rootCmd.Flags().StringP("auth-type", "A", "", "The authentication mechanism to be used.")
	rootCmd.Flags().BoolP("body", "b", false, "Print only the response body.")
	rootCmd.Flags().String("boundary", "", "Specify a custom boundary string for multipart/form-data requests.")
	rootCmd.Flags().String("cert", "", "You can specify a local cert to use as client side SSL certificate.")
	rootCmd.Flags().String("cert-key", "", "The private key to use with SSL.")
	rootCmd.Flags().Bool("check-status", false, "By default, HTTPie exits with 0 when no network or other fatal errors occur.")
	rootCmd.Flags().Bool("chunked", false, "")
	rootCmd.Flags().String("ciphers", "", "A string in the OpenSSL cipher list format.")
	rootCmd.Flags().BoolP("compress", "x", false, "Content compressed (encoded) with Deflate algorithm.")
	rootCmd.Flags().BoolP("continue", "c", false, "Resume an interrupted download.")
	rootCmd.Flags().Bool("debug", false, "Prints the exception traceback should one occur.")
	rootCmd.Flags().String("default-scheme", "", "The default scheme to use if not specified in the URL.")
	rootCmd.Flags().BoolP("download", "d", false, "Do not print the response body to stdout.")
	rootCmd.Flags().BoolP("follow", "F", false, "Follow 30x Location redirects.")
	rootCmd.Flags().BoolP("form", "f", false, "Data items from the command line are serialized as form fields.")
	rootCmd.Flags().String("format-options", "", "Controls output formatting.")
	rootCmd.Flags().BoolP("headers", "h", false, "Print only the response headers.")
	rootCmd.Flags().Bool("help", false, "Show this help message and exit.")
	rootCmd.Flags().StringP("history-print", "P", "", "The same as --print, -p but applies only to intermediary requests/responses.")
	rootCmd.Flags().Bool("ignore-netrc", false, "Ignore credentials from .netrc.")
	rootCmd.Flags().BoolP("ignore-stdin", "I", false, "Do not attempt to read stdin.")
	rootCmd.Flags().BoolP("json", "j", false, "Data items from the command line are serialized as a JSON object.")
	rootCmd.Flags().String("max-headers", "", "The maximum number of response headers to be read before giving up.")
	rootCmd.Flags().String("max-redirects", "", "By default, requests have a limit of 30 redirects (works with --follow).")
	rootCmd.Flags().Bool("multipart", false, "Similar to --form, but always sends a multipart/form-data request.")
	rootCmd.Flags().Bool("offline", false, "Build the request and print it but don’t actually send it.")
	rootCmd.Flags().StringP("output", "o", "", "Save output to FILE instead of stdout.")
	rootCmd.Flags().Bool("path-as-is", false, "Bypass dot segment (/../ or /./) URL squashing.")
	rootCmd.Flags().String("pretty", "", "Controls output processing.")
	rootCmd.Flags().StringP("print", "p", "", "String specifying what the output should contain.")
	rootCmd.Flags().String("proxy", "", "String mapping protocol to the URL of the proxy.")
	rootCmd.Flags().BoolP("quiet", "q", false, "Do not print to stdout or stderr.")
	rootCmd.Flags().String("session", "", "Create, or reuse and update a session.")
	rootCmd.Flags().String("session-read-only", "", "Create or read a session without updating it form the request/response exchange.")
	rootCmd.Flags().Bool("sorted", false, "Re-enables all sorting options while formatting output.")
	rootCmd.Flags().String("ssl", "", "The desired protocol version to use.")
	rootCmd.Flags().BoolP("stream", "S", false, "Always stream the response body by line, i.e., behave like `tail -f'.")
	rootCmd.Flags().StringP("style", "s", "", "")
	rootCmd.Flags().String("timeout", "", "The connection timeout of the request in seconds.")
	rootCmd.Flags().Bool("traceback", false, "Prints the exception traceback should one occur.")
	rootCmd.Flags().Bool("unsorted", false, "Disables all sorting while formatting output.")
	rootCmd.Flags().BoolP("verbose", "v", false, "Verbose output.")
	rootCmd.Flags().String("verify", "", "Set to \"no\" (or \"false\") to skip checking the host's SSL certificate.")
	rootCmd.Flags().Bool("version", false, "Show version and exit.")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"auth-type": carapace.ActionValues("basic", "digest"),
		"cert":      carapace.ActionFiles(),
		"cert-key":  carapace.ActionFiles(),
		"ciphers": carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues() // TODO complete ciphers
		}),
		"default-scheme":    carapace.ActionValues("http://", "https://"),
		"format-options":    ActionFormatOptions(),
		"history-print":     ActionPrintOptions(),
		"output":            carapace.ActionFiles(),
		"pretty":            carapace.ActionValues("all", "colors", "format", "none"),
		"print":             ActionPrintOptions(),
		"session":           carapace.ActionFiles(), // TODO complete names
		"session-read-only": carapace.ActionFiles(), // TODO complete names
		"ssl":               carapace.ActionValues("ssl2.3", "tls1", "tls1.1", "tls1.2"),
		"style":             ActionStyles(),
		"verify":            carapace.ActionValues("yes", "no").StyleF(style.ForKeyword),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.Batch(
			http.ActionRequestMethods(),
			http.ActionUrls(),
		).ToA(),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			requestMethods := map[string]bool{
				"GET":     true,
				"HEAD":    true,
				"POST":    true,
				"PUT":     true,
				"DELETE":  true,
				"CONNECT": true,
				"OPTIONS": true,
				"TRACE":   true,
				"PATCH":   true,
			}
			if _, ok := requestMethods[c.Args[0]]; ok { // arg[1] is url
				return http.ActionUrls()
			}
			return ActionRequestItem()
		}),
	)

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		ActionRequestItem(),
	)
}

func ActionPrintOptions() carapace.Action {
	return carapace.ActionMultiParts("", func(c carapace.Context) carapace.Action {
		return carapace.ActionValuesDescribed(
			"H", "request headers",
			"B", "request body",
			"h", "response headers",
			"b", "response body",
		).Invoke(c).Filter(c.Parts).ToA()
	})
}

func ActionStyles() carapace.Action {
	return carapace.ActionValues(
		"abap", "algol", "algol_nu", "arduino", "auto", "autumn", "borland", "bw",
		"colorful", "default", "emacs", "friendly", "fruity", "igor", "inkpot",
		"lovelace", "manni", "material", "monokai", "murphy", "native", "paraiso-dark",
		"paraiso-light", "pastie", "perldoc", "rainbow_dash", "rrt", "sas", "solarized",
		"solarized-dark", "solarized-light", "stata", "stata-dark", "stata-light", "tango",
		"trac", "vim", "vs", "xcode", "zenburn",
	)
}

func ActionFormatOptions() carapace.Action {
	return carapace.ActionMultiParts(",", func(cEntries carapace.Context) carapace.Action {
		return carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			options := map[string]carapace.Action{
				"headers.sort":   carapace.ActionValues("true", "false").StyleF(style.ForKeyword),
				"json.format":    carapace.ActionValues("true", "false").StyleF(style.ForKeyword),
				"json.indent":    carapace.ActionValues(),
				"json.sort_keys": carapace.ActionValues("true", "false").StyleF(style.ForKeyword),
			}
			switch len(c.Parts) {
			case 0:
				keys := make([]string, 0)
				for key := range options {
					keys = append(keys, key)
				}

				entryKeys := make([]string, len(cEntries.Parts))
				for index, entry := range cEntries.Parts {
					entryKeys[index] = strings.Split(entry, ":")[0]
				}

				return carapace.ActionValues(keys...).Invoke(c).Filter(entryKeys).Suffix(":").ToA()
			case 1:
				if a, ok := options[c.Parts[0]]; ok {
					return a
				}
				return carapace.ActionValues()
			default:
				return carapace.ActionValues()
			}
		})
	})
}

func determineSeparator(callbackValue string) (result string, err error) {
	resultIndex := len(callbackValue)
	for _, separator := range []string{"==", "=@", "@", "=", ":=@", ":=", ":"} {
		if index := strings.Index(callbackValue, separator); index != -1 && index < resultIndex {
			resultIndex = index
			result = separator
		}
	}

	if result == "" {
		err = errors.New("contains no separator")
	}
	return
}

func ActionRequestItem() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if separator, err := determineSeparator(c.CallbackValue); err == nil {
			return carapace.ActionMultiParts(separator, func(c carapace.Context) carapace.Action {
				switch len(c.Parts) {
				case 1:
					switch separator {
					case "=@":
						return carapace.ActionFiles()
					case ":=@":
						return carapace.ActionFiles()
					case ":":
						return http.ActionRequestHeaderValues(c.Parts[0])
					default:
						return carapace.ActionValues()
					}
				default:
					return carapace.ActionValues()
				}
			})
		}

		context := c
		if index := strings.IndexAny(c.CallbackValue, ":=@"); index != -1 {
			context.CallbackValue = context.CallbackValue[:index]
		}

		headers := http.ActionRequestHeaderNames().Invoke(c).Suffix(":") // use full context
		return carapace.ActionMultiParts("", func(c carapace.Context) carapace.Action {
			return ActionSeparators()
		}).Invoke(context).Merge(headers).ToA()
	})
}

func ActionSeparators() carapace.Action {
	return carapace.ActionValuesDescribed(
		":", "HTTP headers",
		"==", "URL parameters to be appended to the request URI",
		"=", "Data fields to be serialized into a JSON object",
		":=", "Non-string JSON data fields (only with --json, -j)",
		"@", "Form file fields (only with --form or --multipart)",
		"=@", "A data field like '=', but takes a file path and embeds its content",
		":=@", "A raw JSON field like ':=', but takes a file path and embeds its content",
	)
}
