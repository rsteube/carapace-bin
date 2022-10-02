package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/micro_completer/cmd/action"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "micro",
	Short: "A modern and intuitive terminal-based text editor",
	Long:  "https://micro-editor.github.io/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("autoindent", "autoindent", "", "")
	rootCmd.Flags().StringS("autosave", "autosave", "", "")
	rootCmd.Flags().StringS("autosu", "autosu", "", "")
	rootCmd.Flags().StringS("backup", "backup", "", "")
	rootCmd.Flags().StringS("backupdir", "backupdir", "", "")
	rootCmd.Flags().StringS("basename", "basename", "", "")
	rootCmd.Flags().BoolS("clean", "clean", false, "Cleans the configuration directory")
	rootCmd.Flags().StringS("clipboard", "clipboard", "", "")
	rootCmd.Flags().StringS("colorcolumn", "colorcolumn", "", "")
	rootCmd.Flags().StringS("colorscheme", "colorscheme", "", "")
	rootCmd.Flags().StringS("config-dir", "config-dir", "", "Specify a custom location for the configuration directory")
	rootCmd.Flags().StringS("cursorline", "cursorline", "", "")
	rootCmd.Flags().BoolS("debug", "debug", false, "Enable debug mode")
	rootCmd.Flags().StringS("diffgutter", "diffgutter", "", "")
	rootCmd.Flags().StringS("divchars", "divchars", "", "")
	rootCmd.Flags().StringS("divreverse", "divreverse", "", "")
	rootCmd.Flags().StringS("encoding", "encoding", "", "")
	rootCmd.Flags().StringS("eofnewline", "eofnewline", "", "")
	rootCmd.Flags().StringS("fastdirty", "fastdirty", "", "")
	rootCmd.Flags().StringS("fileformat", "fileformat", "", "")
	rootCmd.Flags().StringS("filetype", "filetype", "", "")
	rootCmd.Flags().StringS("ignorecase", "ignorecase", "", "")
	rootCmd.Flags().StringS("incsearch", "incsearch", "", "")
	rootCmd.Flags().StringS("indentchar", "indentchar", "", "")
	rootCmd.Flags().StringS("infobar", "infobar", "", "")
	rootCmd.Flags().StringS("keepautoindent", "keepautoindent", "", "")
	rootCmd.Flags().StringS("keymenu", "keymenu", "", "")
	rootCmd.Flags().StringS("matchbrace", "matchbrace", "", "")
	rootCmd.Flags().StringS("mkparents", "mkparents", "", "")
	rootCmd.Flags().StringS("mouse", "mouse", "", "")
	rootCmd.Flags().BoolS("options", "options", false, "Show all option help")
	rootCmd.Flags().StringS("parsecursor", "parsecursor", "", "")
	rootCmd.Flags().StringS("paste", "paste", "", "")
	rootCmd.Flags().StringS("permbackup", "permbackup", "", "")
	rootCmd.Flags().StringS("plugin", "plugin", "", "Manage plugins")
	rootCmd.Flags().StringS("pluginchannels", "pluginchannels", "", "")
	rootCmd.Flags().StringS("pluginrepos", "pluginrepos", "", "")
	rootCmd.Flags().StringS("readonly", "readonly", "", "")
	rootCmd.Flags().StringS("relativeruler", "relativeruler", "", "")
	rootCmd.Flags().StringS("rmtrailingws", "rmtrailingws", "", "")
	rootCmd.Flags().StringS("ruler", "ruler", "", "")
	rootCmd.Flags().StringS("savecursor", "savecursor", "", "")
	rootCmd.Flags().StringS("savehistory", "savehistory", "", "")
	rootCmd.Flags().StringS("saveundo", "saveundo", "", "")
	rootCmd.Flags().StringS("scrollbar", "scrollbar", "", "")
	rootCmd.Flags().StringS("scrollmargin", "scrollmargin", "", "")
	rootCmd.Flags().StringS("scrollspeed", "scrollspeed", "", "")
	rootCmd.Flags().StringS("smartpaste", "smartpaste", "", "")
	rootCmd.Flags().StringS("softwrap", "softwrap", "", "")
	rootCmd.Flags().StringS("splitbottom", "splitbottom", "", "")
	rootCmd.Flags().StringS("splitright", "splitright", "", "")
	rootCmd.Flags().StringS("statusformatl", "statusformatl", "", "")
	rootCmd.Flags().StringS("statusformatr", "statusformatr", "", "")
	rootCmd.Flags().StringS("statusline", "statusline", "", "")
	rootCmd.Flags().StringS("sucmd", "sucmd", "", "")
	rootCmd.Flags().StringS("syntax", "syntax", "", "")
	rootCmd.Flags().StringS("tabmovement", "tabmovement", "", "")
	rootCmd.Flags().StringS("tabsize", "tabsize", "", "")
	rootCmd.Flags().StringS("tabstospaces", "tabstospaces", "", "")
	rootCmd.Flags().StringS("useprimary", "useprimary", "", "")
	rootCmd.Flags().BoolS("version", "version", false, "Show the version number and information")
	rootCmd.Flags().StringS("wordwrap", "wordwrap", "", "")
	rootCmd.Flags().StringS("xterm", "xterm", "", "")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"config-dir": carapace.ActionDirectories(),
		"plugin": carapace.ActionValuesDescribed(
			"available", "List available plugins",
			"install", "Install plugins",
			"list", "List installed plugins",
			"remove", "Remove plugins",
			"search", "Search for a plugin",
			"update", "Update plugins",
		),
		// TODO complete options
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if flag := rootCmd.Flag("plugin"); flag.Changed {
				switch flag.Value.String() {
				case "available":
					return carapace.ActionValues()
				case "install":
					return action.ActionAvailablePlugins().Invoke(c).Filter(c.Args).ToA()
				case "list":
					return carapace.ActionValues()
				case "remove":
					return action.ActionInstalledPlugins().Invoke(c).Filter(c.Args).ToA()
				case "search":
					return carapace.ActionValues()
				case "update":
					return action.ActionInstalledPlugins().Invoke(c).Filter(c.Args).ToA()
				default:
					return carapace.ActionValues()
				}
			}
			return carapace.ActionFiles()
		}),
	)
}
