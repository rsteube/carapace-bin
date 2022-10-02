package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/ng_completer/cmd/action"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Builds and serves your app, rebuilding on file changes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serveCmd).Standalone()

	serveCmd.Flags().String("allowed-hosts", "", "List of hosts that are allowed to access the dev server.")
	serveCmd.Flags().String("browser-target", "", "A browser builder target to serve")
	serveCmd.Flags().StringP("configuration", "c", "", "One or more named builder configurations")
	serveCmd.Flags().Bool("disable-host-check", false, "Don't verify connected clients are part of allowed hosts.")
	serveCmd.Flags().Bool("hmr", false, "Enable hot module replacement.")
	serveCmd.Flags().String("host", "", "Host to listen on.")
	serveCmd.Flags().Bool("live-reload", false, "Whether to reload the page on change, using live-reload.")
	serveCmd.Flags().BoolP("open", "o", false, "Opens the url in default browser.")
	serveCmd.Flags().String("poll", "", "Enable and define the file watching poll time period in milliseconds.")
	serveCmd.Flags().String("port", "", "Port to listen on.")
	serveCmd.Flags().String("proxy-config", "", "Proxy configuration file.")
	serveCmd.Flags().String("public-host", "", "The URL that the browser client should use to connect to the development server.")
	serveCmd.Flags().String("serve-path", "", "The pathname where the app will be served.")
	serveCmd.Flags().Bool("ssl", false, "Serve using HTTPS.")
	serveCmd.Flags().String("ssl-cert", "", "SSL certificate to use for serving HTTPS.")
	serveCmd.Flags().String("ssl-key", "", "SSL key to use for serving HTTPS.")
	serveCmd.Flags().Bool("verbose", false, "Adds more details to output logging.")
	serveCmd.Flags().Bool("watch", false, "Rebuild on change.")
	rootCmd.AddCommand(serveCmd)

	// TODO browser-target
	carapace.Gen(serveCmd).FlagCompletion(carapace.ActionMap{
		"configuration": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			if len(c.Args) < 1 {
				return carapace.ActionValues()
			}
			return action.ActionBuilderConfigurations(c.Args[0], "serve").Invoke(c).Filter(c.Parts).ToA()
		}),
		"proxy-config": carapace.ActionFiles(),
		"ssl-cert":     carapace.ActionFiles(),
		"ssl-key":      carapace.ActionFiles(),
	})

	carapace.Gen(serveCmd).PositionalCompletion(
		action.ActionProjects(),
	)
}
