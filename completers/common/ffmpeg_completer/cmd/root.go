package cmd

import (
	"os"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/ffmpeg_completer/cmd/action"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ffmpeg",
	Short: "Hyper fast Audio and Video encoder",
	Long:  "https://ffmpeg.org/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	for _, prefix := range []string{
		"-c:a",
		"-c:v",
		"-b:a",
		"-b:v",
		"-disposition",
	} {
		for index, arg := range os.Args {
			if strings.HasPrefix(arg, prefix) {
				os.Args[index] = prefix // strip stream specifier
			}
		}
	}

	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringSliceS("ab", "ab", []string{}, "audio bitrate (please use -b:a)")
	rootCmd.Flags().StringSliceS("ac", "ac", []string{}, "set number of audio channels")
	rootCmd.Flags().StringSliceS("acodec", "acodec", []string{}, "force audio codec ('copy' to copy stream)")
	rootCmd.Flags().StringSliceS("af", "af", []string{}, "set audio filters")
	rootCmd.Flags().StringSliceS("aframes", "aframes", []string{}, "set the number of audio frames to output")
	rootCmd.Flags().BoolSliceS("an", "an", []bool{}, "disable audio")
	rootCmd.Flags().BoolSliceS("apad", "apad", []bool{}, "audio pad")
	rootCmd.Flags().StringSliceS("aq", "aq", []string{}, "set audio quality (codec-specific)")
	rootCmd.Flags().StringSliceS("ar", "ar", []string{}, "set audio sampling rate (in Hz)")
	rootCmd.Flags().StringSliceS("aspect", "aspect", []string{}, "set aspect ratio (4:3, 16:9 or 1.3333, 1.7777)")
	rootCmd.Flags().StringSliceS("b:a", "b:a", []string{}, "audio bitrate")
	rootCmd.Flags().StringSliceS("b:v", "b:v", []string{}, "video bitrate")
	rootCmd.Flags().StringSliceS("bits_per_raw_sample", "bits_per_raw_sample", []string{}, "set the number of bits per raw sample")
	rootCmd.Flags().BoolS("bsfs", "bsfs", false, "show available bit stream filters")
	rootCmd.Flags().BoolS("buildconf", "buildconf", false, "show build configuration")
	rootCmd.Flags().StringSliceS("c", "c", []string{}, "codec name")
	rootCmd.Flags().StringSliceS("c:a", "c:a", []string{}, "audio codec")
	rootCmd.Flags().StringSliceS("c:v", "c:v", []string{}, "video codec")
	rootCmd.Flags().StringSliceS("canvas_size", "canvas_size", []string{}, "set canvas size (WxH or abbreviation)")
	rootCmd.Flags().StringSliceS("codec", "codec", []string{}, "codec name")
	rootCmd.Flags().BoolS("codecs", "codecs", false, "show available codecs")
	rootCmd.Flags().BoolS("colors", "colors", false, "show available color names")
	rootCmd.Flags().BoolS("decoders", "decoders", false, "show available decoders")
	rootCmd.Flags().BoolS("demuxers", "demuxers", false, "show available demuxers")
	rootCmd.Flags().BoolS("devices", "devices", false, "show available devices")
	rootCmd.Flags().BoolSliceS("discard", "discard", []bool{}, "discard")
	rootCmd.Flags().BoolSliceS("disposition", "disposition", []bool{}, "disposition")
	rootCmd.Flags().BoolSliceS("dn", "dn", []bool{}, "disable data")
	rootCmd.Flags().BoolS("encoders", "encoders", false, "show available encoders")
	rootCmd.Flags().StringSliceS("f", "f", []string{}, "Force inputor output file format")
	rootCmd.Flags().StringSliceS("filter", "filter", []string{}, "set stream filtergraph")
	rootCmd.Flags().BoolS("filter_complex_threads", "filter_complex_threads", false, "number of threads for -filter_complex")
	rootCmd.Flags().StringSliceS("filter_script", "filter_script", []string{}, "read stream filtergraph description from a file")
	rootCmd.Flags().BoolS("filter_threads", "filter_threads", false, "number of non-complex filter threads")
	rootCmd.Flags().BoolS("filters", "filters", false, "show available filters")
	rootCmd.Flags().BoolSliceS("fix_sub_duration", "fix_sub_duration", []bool{}, "fix subtitles duration")
	rootCmd.Flags().BoolS("formats", "formats", false, "show available formats")
	rootCmd.Flags().StringSliceS("fpsmax", "fpsmax", []string{}, "set max frame rate (Hz value, fraction or abbreviation)")
	rootCmd.Flags().StringSliceS("frames", "frames", []string{}, "set the number of frames to output")
	rootCmd.Flags().StringSliceS("fs", "fs", []string{}, "set the limit file size in bytes")
	rootCmd.Flags().StringS("help", "help", "", "show help")
	rootCmd.Flags().BoolS("hide_banner", "hide_banner", false, "Suppress printing banner")
	rootCmd.Flags().StringSliceS("hwaccel", "hwaccel", []string{}, "use HW acceleration")
	rootCmd.Flags().BoolS("hwaccels", "hwaccels", false, "show available HW acceleration methods")
	rootCmd.Flags().StringSliceS("i", "i", []string{}, "input file")
	rootCmd.Flags().BoolS("ignore_unknown", "ignore_unknown", false, "Ignore unknown stream types")
	rootCmd.Flags().BoolS("layouts", "layouts", false, "show standard channel layouts")
	rootCmd.Flags().StringS("loglevel", "loglevel", "", "set logging level")
	rootCmd.Flags().StringSliceS("map_metadata", "map_metadata", []string{}, "set metadata information of outfile from infile")
	rootCmd.Flags().StringS("max_alloc", "max_alloc", "", "set maximum size of a single allocated block")
	rootCmd.Flags().StringSliceS("metadata", "metadata", []string{}, "add metadata")
	rootCmd.Flags().BoolS("muxers", "muxers", false, "show available muxers")
	rootCmd.Flags().BoolS("n", "n", false, "Do not overwrite output files")
	rootCmd.Flags().StringSliceS("pass", "pass", []string{}, "select the pass number (1 to 3)")
	rootCmd.Flags().BoolS("pix_fmts", "pix_fmts", false, "show available pixel formats")
	rootCmd.Flags().StringSliceS("pre", "pre", []string{}, "preset name")
	rootCmd.Flags().StringSliceS("program", "program", []string{}, "add program with specified streams")
	rootCmd.Flags().BoolS("protocols", "protocols", false, "show available protocols")
	rootCmd.Flags().BoolSliceS("reinit_filter", "reinit_filter", []bool{}, "reinit filtergraph on input parameter changes")
	rootCmd.Flags().BoolS("report", "report", false, "generate a report")
	rootCmd.Flags().BoolS("sample_fmts", "sample_fmts", false, "show available audio sample formats")
	rootCmd.Flags().StringSliceS("scodec", "scodec", []string{}, "force subtitle codec ('copy' to copy stream)")
	rootCmd.Flags().BoolSliceS("seek_timestamp", "seek_timestamp", []bool{}, "enable/disable seeking by timestamp with -ss")
	rootCmd.Flags().StringS("sinks", "sinks", "", "list sinks of the output device")
	rootCmd.Flags().BoolSliceS("sn", "sn", []bool{}, "disable subtitle")
	rootCmd.Flags().StringS("sources", "sources", "", "list sources of the input device")
	rootCmd.Flags().StringSliceS("spre", "spre", []string{}, "set the subtitle options to the indicated preset")
	rootCmd.Flags().StringSliceS("ss", "ss", []string{}, "set the start time offset")
	rootCmd.Flags().StringSliceS("sseof", "sseof", []string{}, "set the start time offset relative to EOF")
	rootCmd.Flags().StringSliceS("stag", "stag", []string{}, "force subtitle tag/fourcc")
	rootCmd.Flags().BoolS("stats", "stats", false, "print progress report during encoding")
	rootCmd.Flags().StringSliceS("target", "target", []string{}, "specify target file type (\"vcd\", \"svcd\", \"dvd\", \"dv\" or \"dv50\" with optional prefixes \"pal-\", \"ntsc-\" or \"film-\")")
	rootCmd.Flags().StringSliceS("timecode", "timecode", []string{}, "set initial TimeCode value.")
	rootCmd.Flags().StringSliceS("timestamp", "timestamp", []string{}, "set the recording timestamp ('now' to set the current time)")
	rootCmd.Flags().StringSliceS("to", "to", []string{}, "record or transcode stop time")
	rootCmd.Flags().StringSliceS("vcodec", "vcodec", []string{}, "force video codec ('copy' to copy stream)")
	rootCmd.Flags().BoolS("version", "version", false, "show version")
	rootCmd.Flags().StringSliceS("vf", "vf", []string{}, "set video filters")
	rootCmd.Flags().StringSliceS("vframes", "vframes", []string{}, "set the number of video frames to output")
	rootCmd.Flags().BoolSliceS("vn", "vn", []bool{}, "disable video")
	rootCmd.Flags().StringSliceS("vol", "vol", []string{}, "change audio volume (256=normal)")
	rootCmd.Flags().BoolS("y", "y", false, "Overwrite output files without asking")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"ab":     carapace.ActionValues("16", "32", "64", "128", "192", "256", "320"),
		"acodec": action.ActionCodecs(), // TODO only audio
		"af": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
				switch len(c.Parts) {
				case 0:
					return action.ActionFilters()
				default:
					return carapace.ActionValues()
				}
			})
		}),
		"ar":      carapace.ActionValues("22050", "44100", "48000"),
		"c":       action.ActionCodecs(),
		"c:a":     action.ActionCodecs(),
		"c:v":     action.ActionCodecs(),
		"codec":   action.ActionCodecs(),
		"f":       action.ActionFormats(),
		"help":    action.ActionHelpTopics(),
		"hwaccel": action.ActionHwAccelerations(),
		"i":       carapace.ActionFiles(),
		"loglevel": carapace.ActionValuesDescribed(
			"quiet", "Show nothing at all; be silent.",
			"panic", "Only show fatal errors which could lead the process to crash",
			"fatal", "Only show fatal errors.",
			"error", "Show all errors, including ones which can be recovered from.",
			"warning", "Show all warnings and errors.",
			"info", "Show informative messages during processing.",
			"verbose", "Same as \"info\", except more verbose.",
			"debug", "Show everything, including debugging information.",
			"trace", "",
		),
		"sinks": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return action.ActionDevices()
			default:
				return carapace.ActionValues()
			}
		}),
		"sources": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return action.ActionDevices()
			default:
				return carapace.ActionValues()
			}
		}),
		"vcodec": action.ActionCodecs(), // TODO only video
		"vf": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
				switch len(c.Parts) {
				case 0:
					return action.ActionFilters()
				default:
					return carapace.ActionValues()
				}
			})
		}),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
