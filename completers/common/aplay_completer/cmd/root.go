package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/aplay_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "aplay",
	Short: "command-line sound recorder and player for ALSA soundcard driver",
	Long:  "https://en.wikipedia.org/wiki/Aplay",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func ExecuteArecord() error {
	rootCmd.Use = "arecord"
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("avail-min", "A", "", "min available space for wakeup is # microseconds")
	rootCmd.Flags().String("buffer-size", "", "buffer duration is # frames")
	rootCmd.Flags().StringP("buffer-time", "B", "", "buffer duration is # microseconds")
	rootCmd.Flags().StringP("channels", "c", "", "channels")
	rootCmd.Flags().StringP("chmap", "m", "", "Give the channel map to override or follow")
	rootCmd.Flags().StringP("device", "D", "", "select PCM by name")
	rootCmd.Flags().Bool("disable-channels", false, "disable automatic channel conversions")
	rootCmd.Flags().Bool("disable-format", false, "disable automatic format conversions")
	rootCmd.Flags().Bool("disable-resample", false, "disable automatic rate resample")
	rootCmd.Flags().Bool("disable-softvol", false, "disable software volume control (softvol)")
	rootCmd.Flags().Bool("dump-hw-params", false, "dump hw_params of the device")
	rootCmd.Flags().StringP("duration", "d", "", "interrupt after # seconds")
	rootCmd.Flags().Bool("fatal-errors", false, "treat all errors as fatal")
	rootCmd.Flags().StringP("file-type", "t", "", "file type (voc, wav, raw or au)")
	rootCmd.Flags().StringP("format", "f", "", "sample format (case insensitive)")
	rootCmd.Flags().BoolP("help", "h", false, "help")
	rootCmd.Flags().BoolP("interactive", "i", false, "allow interactive operation from stdin")
	rootCmd.Flags().BoolP("list-devices", "l", false, "list all soundcards and digital audio devices")
	rootCmd.Flags().BoolP("list-pcms", "L", false, "list device names")
	rootCmd.Flags().String("max-file-time", "", "start another output file when the old file has recorded for this many seconds")
	rootCmd.Flags().BoolP("mmap", "M", false, "mmap stream")
	rootCmd.Flags().BoolP("nonblock", "N", false, "nonblocking mode")
	rootCmd.Flags().String("period-size", "", "distance between interrupts is # frames")
	rootCmd.Flags().StringP("period-time", "F", "", "distance between interrupts is # microseconds")
	rootCmd.Flags().Bool("process-id-file", false, "write the process ID here")
	rootCmd.Flags().BoolP("quiet", "q", false, "quiet mode")
	rootCmd.Flags().StringP("rate", "r", "", "sample rate")
	rootCmd.Flags().StringP("samples", "s", "", "interrupt after # samples per channel")
	rootCmd.Flags().StringP("start-delay", "R", "", "delay for automatic PCM start is # microseconds ")
	rootCmd.Flags().StringP("stop-delay", "T", "", "delay for automatic PCM stop is # microseconds from xrun")
	rootCmd.Flags().String("test-coef", "", "test coefficient for ring buffer position (default 8)")
	rootCmd.Flags().Bool("test-nowait", false, "do not wait for ring buffer - eats whole CPU")
	rootCmd.Flags().Bool("test-position", false, "test ring buffer position")
	rootCmd.Flags().Bool("use-strftime", false, "apply the strftime facility to the output file name")
	rootCmd.Flags().BoolP("verbose", "v", false, "show PCM structure and setup (accumulative)")
	rootCmd.Flags().Bool("version", false, "print current version")
	rootCmd.Flags().StringP("vumeter", "V", "", "enable VU meter (TYPE: mono or stereo)")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"device":    os.ActionSoundCards(),
		"file-type": carapace.ActionValues("voc", "wav", "raw", "au"),
		"format":    action.ActionSampleFormats(),
		"vumeter":   carapace.ActionValues("mono", "stereo"),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if rootCmd.Use == "arecord" && len(c.Args) > 0 {
				return carapace.ActionValues()
			}
			return carapace.ActionFiles()
		}),
	)
}
