package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/docker-compose_completer/cmd/action"
	"github.com/spf13/cobra"
)

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Converts the compose file to platform's canonical format",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(convertCmd).Standalone()
	convertCmd.Flags().String("format", "yaml", "Format the output. Values: [yaml | json]")
	convertCmd.Flags().String("hash", "", "Print the service config hash, one per line.")
	convertCmd.Flags().Bool("images", false, "Print the image names, one per line.")
	convertCmd.Flags().Bool("no-interpolate", false, "Don't interpolate environment variables.")
	convertCmd.Flags().Bool("no-normalize", false, "Don't normalize compose model.")
	convertCmd.Flags().StringP("output", "o", "", "Save to file (default to stdout)")
	convertCmd.Flags().Bool("profiles", false, "Print the profile names, one per line.")
	convertCmd.Flags().BoolP("quiet", "q", false, "Only validate the configuration, don't print anything.")
	convertCmd.Flags().Bool("resolve-image-digests", false, "Pin image tags to digests.")
	convertCmd.Flags().Bool("services", false, "Print the service names, one per line.")
	convertCmd.Flags().Bool("volumes", false, "Print the volume names, one per line.")
	rootCmd.AddCommand(convertCmd)

	carapace.Gen(convertCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("yaml", "json"),
		"hash":   action.ActionServices(convertCmd),
		"output": carapace.ActionFiles(),
	})

	carapace.Gen(convertCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionServices(convertCmd).Invoke(c).Filter(c.Args).ToA()
		}),
	)
}
