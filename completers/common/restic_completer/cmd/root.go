package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/restic_completer/cmd/action"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "restic",
	Short: "Backup and restore files",
	Long:  "https://restic.net/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.PersistentFlags().StringSlice("cacert", []string{}, "`file` to load root certificates from (default: use system certificates)")
	rootCmd.PersistentFlags().String("cache-dir", "", "set the cache `directory`. (default: use system default cache directory)")
	rootCmd.PersistentFlags().Bool("cleanup-cache", false, "auto remove old cache directories")
	rootCmd.Flags().BoolP("help", "h", false, "help for restic")
	rootCmd.PersistentFlags().Bool("insecure-tls", false, "skip TLS certificate verification when connecting to the repo (insecure)")
	rootCmd.PersistentFlags().Bool("json", false, "set output mode to JSON for commands that support it")
	rootCmd.PersistentFlags().String("key-hint", "", "`key` ID of key to try decrypting first (default: $RESTIC_KEY_HINT)")
	rootCmd.PersistentFlags().Int("limit-download", 0, "limits downloads to a maximum rate in KiB/s. (default: unlimited)")
	rootCmd.PersistentFlags().Int("limit-upload", 0, "limits uploads to a maximum rate in KiB/s. (default: unlimited)")
	rootCmd.PersistentFlags().Bool("no-cache", false, "do not use a local cache")
	rootCmd.PersistentFlags().Bool("no-lock", false, "do not lock the repository, this allows some operations on read-only repositories")
	rootCmd.PersistentFlags().StringSliceP("option", "o", []string{}, "set extended option (`key=value`, can be specified multiple times)")
	rootCmd.PersistentFlags().String("password-command", "", "shell `command` to obtain the repository password from (default: $RESTIC_PASSWORD_COMMAND)")
	rootCmd.PersistentFlags().StringP("password-file", "p", "", "`file` to read the repository password from (default: $RESTIC_PASSWORD_FILE)")
	rootCmd.PersistentFlags().BoolP("quiet", "q", false, "do not output comprehensive progress report")
	rootCmd.PersistentFlags().StringP("repo", "r", "", "`repository` to backup to or restore from (default: $RESTIC_REPOSITORY)")
	rootCmd.PersistentFlags().String("repository-file", "", "`file` to read the repository location from (default: $RESTIC_REPOSITORY_FILE)")
	rootCmd.PersistentFlags().String("tls-client-cert", "", "path to a `file` containing PEM encoded TLS client certificate and private key")
	rootCmd.PersistentFlags().StringArrayP("verbose", "v", []string{}, "be verbose (specify multiple times or a level using --verbose=`n`, max level/times is 3)")
	rootCmd.Flag("verbose").NoOptDefVal = "+1"

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"cacert": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionFiles()
		}),
		"cache-dir":        carapace.ActionDirectories(),
		"password-command": carapace.ActionFiles(),
		"password-file":    carapace.ActionFiles(),
		"repo":             action.ActionRepo(),
		"repository-file":  carapace.ActionFiles(),
		"tls-client-cert":  carapace.ActionFiles(),
		"verbose":          carapace.ActionValues("1", "2", "3"),
	})
}
