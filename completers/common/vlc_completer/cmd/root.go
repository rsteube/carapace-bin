package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "vlc",
	Short: "the VLC media player",
	Long:  "https://www.videolan.org/vlc/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("advanced", false, "Show advanced options")
	rootCmd.Flags().String("aspect-ratio", "", "Source aspect ratio")
	rootCmd.Flags().Bool("audio", false, "Enable audio")
	rootCmd.Flags().String("audio-filter", "", "Audio filters")
	rootCmd.Flags().String("audio-language", "", "Audio language")
	rootCmd.Flags().String("audio-replay-gain-default", "", "Default replay gain")
	rootCmd.Flags().String("audio-replay-gain-mode", "", "Replay gain mode")
	rootCmd.Flags().Bool("audio-replay-gain-preamp", false, "VAL   Replay preamp")
	rootCmd.Flags().Bool("audio-time-stretch", false, "Enable time stretching audio")
	rootCmd.Flags().Bool("audio-visual", false, "Audio visualizations")
	rootCmd.Flags().Bool("auto-preparse", false, "Automatically preparse items")
	rootCmd.Flags().Bool("autoscale", false, "Video Auto Scaling")
	rootCmd.Flags().String("bookmark1", "", "Playlist bookmark 1")
	rootCmd.Flags().String("bookmark10", "", "Playlist bookmark 10")
	rootCmd.Flags().String("bookmark2", "", "Playlist bookmark 2")
	rootCmd.Flags().String("bookmark3", "", "Playlist bookmark 3")
	rootCmd.Flags().String("bookmark4", "", "Playlist bookmark 4")
	rootCmd.Flags().String("bookmark5", "", "Playlist bookmark 5")
	rootCmd.Flags().String("bookmark6", "", "Playlist bookmark 6")
	rootCmd.Flags().String("bookmark7", "", "Playlist bookmark 7")
	rootCmd.Flags().String("bookmark8", "", "Playlist bookmark 8")
	rootCmd.Flags().String("bookmark9", "", "Playlist bookmark 9")
	rootCmd.Flags().String("config", "", "use alternate config file")
	rootCmd.Flags().String("control", "", "Control interfaces")
	rootCmd.Flags().String("crop", "", "Video cropping")
	rootCmd.Flags().String("custom-aspect-ratios", "", "Custom aspect ratios list")
	rootCmd.Flags().String("custom-crop-ratios", "", "Custom crop ratios list")
	rootCmd.Flags().String("deinterlace", "", "Deinterlace")
	rootCmd.Flags().String("deinterlace-mode", "", "Deinterlace mode")
	rootCmd.Flags().String("dvd", "", "DVD device")
	rootCmd.Flags().String("extraintf", "", "Extra interface modules")
	rootCmd.Flags().String("extrashort-jump-size", "", "Very short jump length")
	rootCmd.Flags().String("force-dolby-surround", "", "Force detection of Dolby Surround")
	rootCmd.Flags().BoolP("full-help", "H", false, "Exhaustive help for VLC and its modules")
	rootCmd.Flags().BoolP("fullscreen", "f", false, "Fullscreen video output")
	rootCmd.Flags().String("global-key-aspect-ratio", "", "Cycle source aspect ratio")
	rootCmd.Flags().String("global-key-audio-track", "", "Cycle audio track")
	rootCmd.Flags().String("global-key-audiodevice-cycle", "", "Cycle through audio devices")
	rootCmd.Flags().String("global-key-crop", "", "Cycle video crop")
	rootCmd.Flags().String("global-key-decr-scalefactor", "", "Decrease scale factor")
	rootCmd.Flags().String("global-key-deinterlace", "", "Toggle deinterlacing")
	rootCmd.Flags().String("global-key-deinterlace-mode", "", "Cycle deinterlace modes")
	rootCmd.Flags().String("global-key-faster", "", "Faster")
	rootCmd.Flags().String("global-key-frame-next", "", "Next frame")
	rootCmd.Flags().String("global-key-incr-scalefactor", "", "Increase scale factor")
	rootCmd.Flags().String("global-key-intf-show", "", "Show controller in fullscreen")
	rootCmd.Flags().String("global-key-jump+extrashort", "", "Very short forward jump")
	rootCmd.Flags().String("global-key-jump+long", "", "Long forward jump")
	rootCmd.Flags().String("global-key-jump+medium", "", "Medium forward jump")
	rootCmd.Flags().String("global-key-jump+short", "", "Short forward jump")
	rootCmd.Flags().String("global-key-jump-extrashort", "", "Very short backwards jump")
	rootCmd.Flags().String("global-key-jump-long", "", "Long backwards jump")
	rootCmd.Flags().String("global-key-jump-medium", "", "Medium backwards jump")
	rootCmd.Flags().String("global-key-jump-short", "", "Short backwards jump")
	rootCmd.Flags().String("global-key-leave-fullscreen", "", "Exit fullscreen")
	rootCmd.Flags().String("global-key-loop", "", "Normal/Loop/Repeat")
	rootCmd.Flags().String("global-key-next", "", "Next")
	rootCmd.Flags().String("global-key-play-pause", "", "Play/Pause")
	rootCmd.Flags().String("global-key-program-sid-next", "", "Cycle next program Service ID")
	rootCmd.Flags().String("global-key-program-sid-prev", "", "Cycle previous program Service ID")
	rootCmd.Flags().String("global-key-quit", "", "Quit")
	rootCmd.Flags().String("global-key-random", "", "Random")
	rootCmd.Flags().String("global-key-rate-faster-fine", "", "Faster (fine)")
	rootCmd.Flags().String("global-key-rate-normal", "", "Normal rate")
	rootCmd.Flags().String("global-key-rate-slower-fine", "", "Slower (fine)")
	rootCmd.Flags().String("global-key-slower", "", "Slower")
	rootCmd.Flags().String("global-key-subtitle-revtrack", "", "Cycle subtitle track in reverse order")
	rootCmd.Flags().String("global-key-subtitle-toggle", "", "Toggle subtitles")
	rootCmd.Flags().String("global-key-subtitle-track", "", "Cycle subtitle track")
	rootCmd.Flags().String("global-key-toggle-autoscale", "", "Toggle autoscaling")
	rootCmd.Flags().String("global-key-toggle-fullscreen", "", "Fullscreen")
	rootCmd.Flags().String("global-key-vol-down", "", "Volume down")
	rootCmd.Flags().String("global-key-vol-mute", "", "Mute")
	rootCmd.Flags().String("global-key-vol-up", "", "Volume up")
	rootCmd.Flags().String("global-key-wallpaper", "", "Toggle wallpaper mode in video output")
	rootCmd.Flags().String("global-key-zoom-double", "", "2:1 Double")
	rootCmd.Flags().String("global-key-zoom-half", "", "1:2 Half")
	rootCmd.Flags().String("global-key-zoom-original", "", "1:1 Original")
	rootCmd.Flags().String("global-key-zoom-quarter", "", "1:4 Quarter")
	rootCmd.Flags().BoolP("help", "h", false, "print help for VLC")
	rootCmd.Flags().Bool("help-verbose", false, "ask for extra verbosity when displaying help")
	rootCmd.Flags().String("hotkeys-x-wheel-mode", "", "Mouse wheel horizontal axis control")
	rootCmd.Flags().String("hotkeys-y-wheel-mode", "", "Mouse wheel vertical axis control")
	rootCmd.Flags().Bool("ignore-config", false, "no configuration option will be loaded")
	rootCmd.Flags().String("ignore-filetypes", "", "Ignored extensions")
	rootCmd.Flags().Bool("input-fast-seek", false, "Fast seek")
	rootCmd.Flags().String("input-repeat", "", "Input repetitions")
	rootCmd.Flags().String("input-title-format", "", "Change title according to current media")
	rootCmd.Flags().Bool("interact", false, "Interface interaction")
	rootCmd.Flags().StringP("intf", "I", "", "Interface module")
	rootCmd.Flags().String("key-aspect-ratio", "", "Cycle source aspect ratio")
	rootCmd.Flags().String("key-audio-track", "", "Cycle audio track")
	rootCmd.Flags().String("key-audiodevice-cycle", "", "Cycle through audio devices")
	rootCmd.Flags().String("key-crop", "", "Cycle video crop")
	rootCmd.Flags().String("key-decr-scalefactor", "", "Decrease scale factor")
	rootCmd.Flags().String("key-deinterlace", "", "Toggle deinterlacing")
	rootCmd.Flags().String("key-deinterlace-mode", "", "Cycle deinterlace modes")
	rootCmd.Flags().String("key-faster", "", "Faster")
	rootCmd.Flags().String("key-frame-next", "", "Next frame")
	rootCmd.Flags().String("key-incr-scalefactor", "", "Increase scale factor")
	rootCmd.Flags().String("key-intf-show", "", "Show controller in fullscreen")
	rootCmd.Flags().String("key-jump+extrashort", "", "Very short forward jump")
	rootCmd.Flags().String("key-jump+long", "", "Long forward jump")
	rootCmd.Flags().String("key-jump+medium", "", "Medium forward jump")
	rootCmd.Flags().String("key-jump+short", "", "Short forward jump")
	rootCmd.Flags().String("key-jump-extrashort", "", "Very short backwards jump")
	rootCmd.Flags().String("key-jump-long", "", "Long backwards jump")
	rootCmd.Flags().String("key-jump-medium", "", "Medium backwards jump")
	rootCmd.Flags().String("key-jump-short", "", "Short backwards jump")
	rootCmd.Flags().String("key-leave-fullscreen", "", "Exit fullscreen")
	rootCmd.Flags().String("key-loop", "", "Normal/Loop/Repeat")
	rootCmd.Flags().String("key-next", "", "Next")
	rootCmd.Flags().String("key-play-pause", "", "Play/Pause")
	rootCmd.Flags().String("key-prev", "", "Previous")
	rootCmd.Flags().String("key-program-sid-next", "", "Cycle next program Service ID")
	rootCmd.Flags().String("key-program-sid-prev", "", "Cycle previous program Service ID")
	rootCmd.Flags().String("key-quit", "", "Quit")
	rootCmd.Flags().String("key-random", "", "Random")
	rootCmd.Flags().String("key-rate-faster-fine", "", "Faster (fine)")
	rootCmd.Flags().String("key-rate-normal", "", "Normal rate")
	rootCmd.Flags().String("key-rate-slower-fine", "", "Slower (fine)")
	rootCmd.Flags().String("key-slower", "", "Slower")
	rootCmd.Flags().String("key-stop", "", "Stop")
	rootCmd.Flags().String("key-subtitle-revtrack", "", "Cycle subtitle track in reverse order")
	rootCmd.Flags().String("key-subtitle-toggle", "", "Toggle subtitles")
	rootCmd.Flags().String("key-subtitle-track", "", "Cycle subtitle track")
	rootCmd.Flags().String("key-toggle-autoscale", "", "Toggle autoscaling")
	rootCmd.Flags().String("key-toggle-fullscreen", "", "Fullscreen")
	rootCmd.Flags().String("key-vol-down", "", "Volume down")
	rootCmd.Flags().String("key-vol-mute", "", "Mute")
	rootCmd.Flags().String("key-vol-up", "", "Volume up")
	rootCmd.Flags().String("key-wallpaper", "", "Toggle wallpaper mode in video output")
	rootCmd.Flags().String("key-zoom-double", "", "2:1 Double")
	rootCmd.Flags().String("key-zoom-half", "", "1:2 Half")
	rootCmd.Flags().String("key-zoom-original", "", "1:1 Original")
	rootCmd.Flags().String("key-zoom-quarter", "", "1:4 Quarter")
	rootCmd.Flags().BoolP("list", "l", false, "print a list of available modules")
	rootCmd.Flags().Bool("list-verbose", false, "print a list of available modules with extra")
	rootCmd.Flags().String("long-jump-size", "", "Long jump length")
	rootCmd.Flags().Bool("longhelp", false, "print help for VLC and all its modules")
	rootCmd.Flags().BoolP("loop", "L", false, "Repeat all")
	rootCmd.Flags().Bool("media-library", false, "Use media library")
	rootCmd.Flags().String("medium-jump-size", "", "Medium jump length")
	rootCmd.Flags().String("menu-language", "", "Menu language")
	rootCmd.Flags().Bool("metadata-network-access", false, "Allow metadata network access")
	rootCmd.Flags().StringP("module", "p", "", "print help on a specific module")
	rootCmd.Flags().String("mouse-hide-timeout", "", "Hide cursor and fullscreen controller after x milliseconds")
	rootCmd.Flags().Bool("no-advanced", false, "Dont show advanced options")
	rootCmd.Flags().Bool("no-audio", false, "Disable audio")
	rootCmd.Flags().Bool("no-audio-time-stretch", false, "Disable time stretching audio")
	rootCmd.Flags().Bool("no-auto-preparse", false, "Dont automatically preparse items")
	rootCmd.Flags().Bool("no-autoscale", false, "Video Auto Scaling")
	rootCmd.Flags().Bool("no-full-help", false, "Exhaustive help for VLC and its modules")
	rootCmd.Flags().Bool("no-fullscreen", false, "No Fullscreen video output")
	rootCmd.Flags().Bool("no-help", false, "print help for VLC")
	rootCmd.Flags().Bool("no-help-verbose", false, "ask for extra verbosity when displaying help")
	rootCmd.Flags().Bool("no-ignore-config", false, "no configuration option will be loaded")
	rootCmd.Flags().Bool("no-input-fast-seek", false, "Fast seek")
	rootCmd.Flags().Bool("no-interact", false, "No interface interaction")
	rootCmd.Flags().Bool("no-list", false, "print a list of available modules")
	rootCmd.Flags().Bool("no-list-verbose", false, "print a list of available modules with extra")
	rootCmd.Flags().Bool("no-longhelp", false, "print help for VLC and all its modules")
	rootCmd.Flags().Bool("no-loop", false, "Dont repeat all")
	rootCmd.Flags().Bool("no-media-library", false, "Use media library")
	rootCmd.Flags().Bool("no-metadata-network-access", false, "Dont allow metadata network access")
	rootCmd.Flags().Bool("no-osd", false, "No on Screen Display")
	rootCmd.Flags().Bool("no-play-and-exit", false, "Play and exit")
	rootCmd.Flags().Bool("no-play-and-stop", false, "Play and stop")
	rootCmd.Flags().Bool("no-playlist-autostart", false, "Auto start")
	rootCmd.Flags().Bool("no-playlist-cork", false, "Pause on audio communication")
	rootCmd.Flags().Bool("no-playlist-tree", false, "Display playlist tree")
	rootCmd.Flags().Bool("no-random", false, "Dont play files randomly forever")
	rootCmd.Flags().Bool("no-repeat", false, "Dont repeat current item")
	rootCmd.Flags().Bool("no-reset-config", false, "reset the current config to the default values")
	rootCmd.Flags().Bool("no-reset-plugins-cache", false, "resets the current plugins cache")
	rootCmd.Flags().Bool("no-show-hiddenfiles", false, "Dont show hidden files")
	rootCmd.Flags().Bool("no-snapshot-preview", false, "No Display video snapshot preview")
	rootCmd.Flags().Bool("no-snapshot-sequential", false, "Use sequential numbers instead of timestamps")
	rootCmd.Flags().Bool("no-spu", false, "Disable sub-pictures")
	rootCmd.Flags().Bool("no-start-paused", false, "Dont start paused")
	rootCmd.Flags().Bool("no-sub-autodetect-file", false, "Autodetect subtitle files")
	rootCmd.Flags().Bool("no-version", false, "print version information")
	rootCmd.Flags().Bool("no-video-on-top", false, "Not always on top")
	rootCmd.Flags().Bool("no-video-title-show", false, "Do not how media title on video")
	rootCmd.Flags().Bool("no-video-wallpaper", false, "Disable wallpaper mode")
	rootCmd.Flags().String("open", "", "Default stream")
	rootCmd.Flags().Bool("osd", false, "On Screen Display")
	rootCmd.Flags().String("pidfile", "", "Write process id to file")
	rootCmd.Flags().Bool("play-and-exit", false, "Play and exit")
	rootCmd.Flags().Bool("play-and-stop", false, "Play and stop")
	rootCmd.Flags().Bool("playlist-autostart", false, "Auto start")
	rootCmd.Flags().Bool("playlist-cork", false, "Pause on audio communication")
	rootCmd.Flags().Bool("playlist-tree", false, "Display playlist tree")
	rootCmd.Flags().String("preferred-resolution", "", "Preferred video resolution")
	rootCmd.Flags().String("preparse-timeout", "", "Preparsing timeout")
	rootCmd.Flags().BoolP("random", "Z", false, "Play files randomly forever")
	rootCmd.Flags().Bool("rate", false, "Playback speed")
	rootCmd.Flags().String("recursive", "", "Subdirectory behavior")
	rootCmd.Flags().BoolP("repeat", "R", false, "Repeat current item")
	rootCmd.Flags().Bool("reset-config", false, "reset the current config to the default values")
	rootCmd.Flags().Bool("reset-plugins-cache", false, "resets the current plugins cache")
	rootCmd.Flags().String("short-jump-size", "", "Short jump length")
	rootCmd.Flags().Bool("show-hiddenfiles", false, "Show hidden files")
	rootCmd.Flags().String("snapshot-format", "", "Video snapshot format")
	rootCmd.Flags().String("snapshot-path", "", "Video snapshot directory (or filename)")
	rootCmd.Flags().String("snapshot-prefix", "", "Video snapshot file prefix")
	rootCmd.Flags().Bool("snapshot-preview", false, "Display video snapshot preview")
	rootCmd.Flags().Bool("snapshot-sequential", false, "Use sequential numbers instead of timestamps")
	rootCmd.Flags().Bool("spu", false, "Enable sub-pictures")
	rootCmd.Flags().Bool("start-paused", false, "Start paused")
	rootCmd.Flags().String("stream-filter", "", "Stream filter module")
	rootCmd.Flags().Bool("sub-autodetect-file", false, "Autodetect subtitle files")
	rootCmd.Flags().String("sub-file", "", "Use subtitle file")
	rootCmd.Flags().String("sub-filter", "", "Subpictures filter module")
	rootCmd.Flags().String("sub-language", "", "Subtitle language")
	rootCmd.Flags().String("sub-source", "", "Subpictures source module")
	rootCmd.Flags().String("sub-text-scale", "", "Subtitles text scaling factor")
	rootCmd.Flags().String("vcd", "", "VCD device")
	rootCmd.Flags().StringP("verbose", "v", "", "Verbosity")
	rootCmd.Flags().Bool("version", false, "print version information")
	rootCmd.Flags().String("video-filter", "", "Video filter module")
	rootCmd.Flags().Bool("video-on-top", false, "Always on top")
	rootCmd.Flags().String("video-splitter", "", "Video splitter module")
	rootCmd.Flags().String("video-title-position", "", "Position of video title")
	rootCmd.Flags().Bool("video-title-show", false, "Show media title on video")
	rootCmd.Flags().String("video-title-timeout", "", "Show video title for x milliseconds")
	rootCmd.Flags().Bool("video-wallpaper", false, "Enable wallpaper mode")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"audio-replay-gain-mode": carapace.ActionValues("none", "track", "album"),
		"audio-visual":           carapace.ActionValues("any", "visual", "projectm", "goom", "glspectrum", "none"),
		"config":                 carapace.ActionFiles(),
		"deinterlace": carapace.ActionValuesDescribed(
			"0", "off",
			"-1,", "automatic",
			"1", "on",
		),
		"deinterlace-mode": carapace.ActionValues("auto", "discard", "blend", "mean", "bob", "linear", "x", "yadif", "yadif2x", "phosphor", "ivtc"),
		"force-dolby-surround": carapace.ActionValuesDescribed(
			"0", "auto",
			"1", "on",
			"2", "off",
		),
		"hotkeys-x-wheel-mode": carapace.ActionValuesDescribed(
			"-1", "Ignore",
			"0", "Volume control",
			"2", "Position control",
			"3", "Position control reversed",
		),
		"hotkeys-y-wheel-mode": carapace.ActionValuesDescribed(
			"-1", "Ignore",
			"0", "Volume control",
			"2", "Position control",
			"3", "Position control reversed",
		),
		"pidfile": carapace.ActionFiles(),
		"preferred-resolution": carapace.ActionValuesDescribed(
			"-1", "best available",
			"1080", "full hd (1080p)",
			"720", "hd (720p)",
			"576", "standard definition (576 or 480 lines)",
			"360", "low definition (360 lines)",
			"240", "very low definition (240 lines)",
		),
		"snapshot-format": carapace.ActionValues("png", "jpg", "tiff"),
		"snapshot-path":   carapace.ActionFiles(),
		"sub-file":        carapace.ActionFiles(),
		"verbose":         carapace.ActionValues("0", "1", "2"),
		"video-title-position": carapace.ActionValuesDescribed(
			"0", "Center",
			"1", "Left",
			"2", "Right",
			"4", "Top",
			"8", "Bottom",
			"5", "Top-Left",
			"6", "Top-Right",
			"9", "Bottom-Left",
			"10", "Bottom-Righ",
		),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
