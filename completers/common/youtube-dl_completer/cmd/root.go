package cmd

import (
	"github.com/rsteube/carapace"
	youtubedl "github.com/rsteube/carapace-bin/pkg/actions/tools/youtubedl"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "youtube-dl",
	Short: "download videos from youtube.com or other video platforms",
	Long:  "https://github.com/ytdl-org/youtube-dl",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("abort-on-error", false, "Abort downloading of further videos (in the")
	rootCmd.Flags().Bool("abort-on-unavailable-fragment", false, "Abort downloading when some fragment is not")
	rootCmd.Flags().String("add-header", "", "Specify a custom HTTP header and its value,")
	rootCmd.Flags().Bool("add-metadata", false, "Write metadata to the video file")
	rootCmd.Flags().String("age-limit", "", "Download only videos suitable for the given")
	rootCmd.Flags().Bool("all-formats", false, "Download all available video formats")
	rootCmd.Flags().Bool("all-subs", false, "Download all the available subtitles of the")
	rootCmd.Flags().Bool("ap-list-mso", false, "List all supported multiple-system")
	rootCmd.Flags().String("ap-mso", "", "Adobe Pass multiple-system operator (TV")
	rootCmd.Flags().String("ap-password", "", "Multiple-system operator account password.")
	rootCmd.Flags().String("ap-username", "", "Multiple-system operator account login")
	rootCmd.Flags().String("audio-format", "", "Specify audio format: \"best\", \"aac\",")
	rootCmd.Flags().String("audio-quality", "", "Specify ffmpeg/avconv audio quality, insert")
	rootCmd.Flags().String("autonumber-start", "", "Specify the start value for %(autonumber)s")
	rootCmd.Flags().StringP("batch-file", "a", "", "File containing URLs to download ('-' for")
	rootCmd.Flags().Bool("bidi-workaround", false, "Work around terminals that lack")
	rootCmd.Flags().String("buffer-size", "", "Size of download buffer (e.g. 1024 or 16K)")
	rootCmd.Flags().String("cache-dir", "", "Location in the filesystem where youtube-dl")
	rootCmd.Flags().BoolP("call-home", "C", false, "Contact the youtube-dl server for debugging")
	rootCmd.Flags().String("config-location", "", "Location of the configuration file; either")
	rootCmd.Flags().Bool("console-title", false, "Display progress in console titlebar")
	rootCmd.Flags().BoolP("continue", "c", false, "Force resume of partially downloaded files.")
	rootCmd.Flags().String("convert-subs", "", "Convert the subtitles to other format")
	rootCmd.Flags().String("cookies", "", "File to read cookies from and dump cookie")
	rootCmd.Flags().String("date", "", "Download only videos uploaded in this date")
	rootCmd.Flags().String("dateafter", "", "Download only videos uploaded on or after")
	rootCmd.Flags().String("datebefore", "", "Download only videos uploaded on or before")
	rootCmd.Flags().String("default-search", "", "Use this prefix for unqualified URLs. For")
	rootCmd.Flags().String("download-archive", "", "Download only videos not listed in the")
	rootCmd.Flags().BoolP("dump-json", "j", false, "Simulate, quiet but print JSON information.")
	rootCmd.Flags().Bool("dump-pages", false, "Print downloaded pages encoded using base64")
	rootCmd.Flags().BoolP("dump-single-json", "J", false, "Simulate, quiet but print JSON information")
	rootCmd.Flags().Bool("dump-user-agent", false, "Display the current browser identification")
	rootCmd.Flags().Bool("embed-subs", false, "Embed subtitles in the video (only for mp4,")
	rootCmd.Flags().Bool("embed-thumbnail", false, "Embed thumbnail in the audio as cover art")
	rootCmd.Flags().String("encoding", "", "Force the specified encoding (experimental)")
	rootCmd.Flags().String("exec", "", "Execute a command on the file after")
	rootCmd.Flags().String("external-downloader", "", "Use the specified external downloader.")
	rootCmd.Flags().String("external-downloader-args", "", "Give these arguments to the external")
	rootCmd.Flags().BoolP("extract-audio", "x", false, "Convert video files to audio-only files")
	rootCmd.Flags().Bool("extractor-descriptions", false, "Output descriptions of all supported")
	rootCmd.Flags().String("ffmpeg-location", "", "Location of the ffmpeg/avconv binary;")
	rootCmd.Flags().String("fixup", "", "Automatically correct known faults of the")
	rootCmd.Flags().Bool("flat-playlist", false, "Do not extract the videos of a playlist,")
	rootCmd.Flags().Bool("force-generic-extractor", false, "Force extraction to use the generic")
	rootCmd.Flags().BoolP("force-ipv4", "4", false, "Make all connections via IPv4")
	rootCmd.Flags().BoolP("force-ipv6", "6", false, "Make all connections via IPv6")
	rootCmd.Flags().StringP("format", "f", "", "Video format code, see the \"FORMAT")
	rootCmd.Flags().String("fragment-retries", "", "Number of retries for a fragment (default")
	rootCmd.Flags().Bool("geo-bypass", false, "Bypass geographic restriction via faking")
	rootCmd.Flags().String("geo-bypass-country", "", "Force bypass geographic restriction with")
	rootCmd.Flags().String("geo-bypass-ip-block", "", "Force bypass geographic restriction with")
	rootCmd.Flags().String("geo-verification-proxy", "", "Use this proxy to verify the IP address for")
	rootCmd.Flags().Bool("get-description", false, "Simulate, quiet but print video description")
	rootCmd.Flags().Bool("get-duration", false, "Simulate, quiet but print video length")
	rootCmd.Flags().Bool("get-filename", false, "Simulate, quiet but print output filename")
	rootCmd.Flags().Bool("get-format", false, "Simulate, quiet but print output format")
	rootCmd.Flags().Bool("get-id", false, "Simulate, quiet but print id")
	rootCmd.Flags().Bool("get-thumbnail", false, "Simulate, quiet but print thumbnail URL")
	rootCmd.Flags().BoolP("get-title", "e", false, "Simulate, quiet but print title")
	rootCmd.Flags().BoolP("get-url", "g", false, "Simulate, quiet but print URL")
	rootCmd.Flags().BoolP("help", "h", false, "Print this help text and exit")
	rootCmd.Flags().Bool("hls-prefer-ffmpeg", false, "Use ffmpeg instead of the native HLS")
	rootCmd.Flags().Bool("hls-prefer-native", false, "Use the native HLS downloader instead of")
	rootCmd.Flags().Bool("hls-use-mpegts", false, "Use the mpegts container for HLS videos,")
	rootCmd.Flags().String("http-chunk-size", "", "Size of a chunk for chunk-based HTTP")
	rootCmd.Flags().Bool("id", false, "Use only video ID in file name")
	rootCmd.Flags().Bool("ignore-config", false, "Do not read configuration files. When given")
	rootCmd.Flags().BoolP("ignore-errors", "i", false, "Continue on download errors, for example to")
	rootCmd.Flags().Bool("include-ads", false, "Download advertisements as well")
	rootCmd.Flags().Bool("keep-fragments", false, "Keep downloaded fragments on disk after")
	rootCmd.Flags().BoolP("keep-video", "k", false, "Keep the video file on disk after the post-")
	rootCmd.Flags().StringP("limit-rate", "r", "", "Maximum download rate in bytes per second")
	rootCmd.Flags().Bool("list-extractors", false, "List all supported extractors")
	rootCmd.Flags().BoolP("list-formats", "F", false, "List all available formats of requested")
	rootCmd.Flags().Bool("list-subs", false, "List all available subtitles for the video")
	rootCmd.Flags().Bool("list-thumbnails", false, "Simulate and list all available thumbnail")
	rootCmd.Flags().String("load-info-json", "", "JSON file containing the video information")
	rootCmd.Flags().Bool("mark-watched", false, "Mark videos watched (YouTube only)")
	rootCmd.Flags().String("match-filter", "", "Generic video filter. Specify any key (see")
	rootCmd.Flags().String("match-title", "", "Download only matching titles (regex or")
	rootCmd.Flags().String("max-downloads", "", "Abort after downloading NUMBER files")
	rootCmd.Flags().String("max-filesize", "", "Do not download any videos larger than SIZE")
	rootCmd.Flags().String("max-sleep-interval", "", "Upper bound of a range for randomized sleep")
	rootCmd.Flags().Bool("max-sleep-interval.", false, "")
	rootCmd.Flags().String("max-views", "", "Do not download any videos with more than")
	rootCmd.Flags().String("merge-output-format", "", "If a merge is required (e.g.")
	rootCmd.Flags().String("metadata-from-title", "", "Parse additional metadata like song title /")
	rootCmd.Flags().String("min-filesize", "", "Do not download any videos smaller than")
	rootCmd.Flags().String("min-views", "", "Do not download any videos with less than")
	rootCmd.Flags().BoolP("netrc", "n", false, "Use .netrc authentication data")
	rootCmd.Flags().Bool("newline", false, "Output progress bar as new lines")
	rootCmd.Flags().Bool("no-cache-dir", false, "Disable filesystem caching")
	rootCmd.Flags().Bool("no-call-home", false, "Do NOT contact the youtube-dl server for")
	rootCmd.Flags().Bool("no-check-certificate", false, "Suppress HTTPS certificate validation")
	rootCmd.Flags().Bool("no-color", false, "Do not emit color codes in output")
	rootCmd.Flags().Bool("no-continue", false, "Do not resume partially downloaded files")
	rootCmd.Flags().Bool("no-geo-bypass", false, "Do not bypass geographic restriction via")
	rootCmd.Flags().Bool("no-mark-watched", false, "Do not mark videos watched (YouTube only)")
	rootCmd.Flags().Bool("no-mtime", false, "Do not use the Last-modified header to set")
	rootCmd.Flags().BoolP("no-overwrites", "w", false, "Do not overwrite files")
	rootCmd.Flags().Bool("no-part", false, "Do not use .part files - write directly")
	rootCmd.Flags().Bool("no-playlist", false, "Download only the video, if the URL refers")
	rootCmd.Flags().Bool("no-post-overwrites", false, "Do not overwrite post-processed files; the")
	rootCmd.Flags().Bool("no-progress", false, "Do not print progress bar")
	rootCmd.Flags().Bool("no-resize-buffer", false, "Do not automatically adjust the buffer")
	rootCmd.Flags().Bool("no-warnings", false, "Ignore warnings")
	rootCmd.Flags().StringP("output", "o", "", "Output filename template, see the \"OUTPUT")
	rootCmd.Flags().StringP("password", "p", "", "Account password. If this option is left")
	rootCmd.Flags().String("playlist-end", "", "Playlist video to end at (default is last)")
	rootCmd.Flags().String("playlist-items", "", "Playlist video items to download. Specify")
	rootCmd.Flags().Bool("playlist-random", false, "Download playlist videos in random order")
	rootCmd.Flags().Bool("playlist-reverse", false, "Download playlist videos in reverse order")
	rootCmd.Flags().String("playlist-start", "", "Playlist video to start at (default is 1)")
	rootCmd.Flags().String("postprocessor-args", "", "Give these arguments to the postprocessor")
	rootCmd.Flags().Bool("prefer-avconv", false, "Prefer avconv over ffmpeg for running the")
	rootCmd.Flags().Bool("prefer-ffmpeg", false, "Prefer ffmpeg over avconv for running the")
	rootCmd.Flags().Bool("prefer-free-formats", false, "Prefer free video formats unless a specific")
	rootCmd.Flags().Bool("prefer-insecure", false, "Use an unencrypted connection to retrieve")
	rootCmd.Flags().Bool("print-json", false, "Be quiet and print the video information as")
	rootCmd.Flags().Bool("print-traffic", false, "Display sent and read HTTP traffic")
	rootCmd.Flags().String("proxy", "", "Use the specified HTTP/HTTPS/SOCKS proxy.")
	rootCmd.Flags().BoolP("quiet", "q", false, "Activate quiet mode")
	rootCmd.Flags().String("recode-video", "", "Encode the video to another format if")
	rootCmd.Flags().String("referer", "", "Specify a custom referer, use if the video")
	rootCmd.Flags().String("reject-title", "", "Skip download for matching titles (regex or")
	rootCmd.Flags().Bool("restrict-filenames", false, "Restrict filenames to only ASCII")
	rootCmd.Flags().StringP("retries", "R", "", "Number of retries (default is 10), or")
	rootCmd.Flags().Bool("rm-cache-dir", false, "Delete all filesystem cache files")
	rootCmd.Flags().BoolP("simulate", "s", false, "Do not download the video and do not write")
	rootCmd.Flags().Bool("skip-download", false, "Do not download the video")
	rootCmd.Flags().Bool("skip-unavailable-fragments", false, "Skip unavailable fragments (DASH, hlsnative")
	rootCmd.Flags().String("sleep-interval", "", "Number of seconds to sleep before each")
	rootCmd.Flags().String("socket-timeout", "", "Time to wait before giving up, in seconds")
	rootCmd.Flags().String("source-address", "", "Client-side IP address to bind to")
	rootCmd.Flags().String("sub-format", "", "Subtitle format, accepts formats")
	rootCmd.Flags().String("sub-lang", "", "Languages of the subtitles to download")
	rootCmd.Flags().StringP("twofactor", "2", "", "Two-factor authentication code")
	rootCmd.Flags().BoolP("update", "U", false, "Update this program to latest version. Make")
	rootCmd.Flags().String("user-agent", "", "Specify a custom user agent")
	rootCmd.Flags().StringP("username", "u", "", "Login with this account ID")
	rootCmd.Flags().BoolP("verbose", "v", false, "Print various debugging information")
	rootCmd.Flags().Bool("version", false, "Print program version and exit")
	rootCmd.Flags().String("video-password", "", "Video password (vimeo, smotri, youku)")
	rootCmd.Flags().Bool("write-all-thumbnails", false, "Write all thumbnail image formats to disk")
	rootCmd.Flags().Bool("write-annotations", false, "Write video annotations to a")
	rootCmd.Flags().Bool("write-auto-sub", false, "Write automatically generated subtitle file")
	rootCmd.Flags().Bool("write-description", false, "Write video description to a .description")
	rootCmd.Flags().Bool("write-info-json", false, "Write video metadata to a .info.json file")
	rootCmd.Flags().Bool("write-pages", false, "Write downloaded intermediary pages to")
	rootCmd.Flags().Bool("write-sub", false, "Write subtitle file")
	rootCmd.Flags().Bool("write-thumbnail", false, "Write thumbnail image to disk")
	rootCmd.Flags().Bool("xattr-set-filesize", false, "Set file xattribute ytdl.filesize with")
	rootCmd.Flags().Bool("xattrs", false, "Write metadata to the video file's xattrs")
	rootCmd.Flags().Bool("yes-playlist", false, "Download the playlist, if the URL refers to")
	rootCmd.Flags().Bool("youtube-skip-dash-manifest", false, "Do not download the DASH manifests and")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"audio-format":    carapace.ActionValues("best", "aac", "flac", "mp3", "m4a", "opus", "vorbis", "wav"),
		"audio-quality":   carapace.ActionValues("0", "1", "2", "3", "4", "5", "6", "7", "8", "9"),
		"batch-file":      carapace.ActionFiles(),
		"cache-dir":       carapace.ActionDirectories(),
		"config-location": carapace.ActionFiles(),
		"convert-subs":    carapace.ActionValues("srt", "ass", "vtt", "lrc"),
		"cookies":         carapace.ActionFiles(),
		"ffmpeg-location": carapace.ActionFiles(),
		"fixup": carapace.ActionValuesDescribed(
			"never", "do nothing",
			"warn", "only emit a warning",
			"detect_or_warn", "fix file if we can, warn otherwise",
		),
		"format": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) == 0 {
				return carapace.ActionMessage("missing url")
			}
			return youtubedl.ActionFormats(c.Args[0])
		}),
		"recode-video": carapace.ActionValues("mp4", "flv", "ogg", "webm", "mkv", "avi"),
		"sub-format": carapace.ActionMultiParts("/", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("ass", "srt", "best").Invoke(c).Filter(c.Parts).ToA()
		}),
		"sub-lang": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			if len(c.Args) == 0 {
				return carapace.ActionMessage("missing url")
			}
			return youtubedl.ActionSubLangs(c.Args[0]).Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
