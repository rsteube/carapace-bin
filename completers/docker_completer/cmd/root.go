package cmd

import (
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bridge/pkg/actions/bridge"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "docker [OPTIONS] COMMAND [ARG...]",
	Short: "A self-sufficient runtime for containers",
	Long:  "https://docs.docker.com/compose/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.AddGroup(
		&cobra.Group{ID: "common", Title: "Common Commands"},
		&cobra.Group{ID: "management", Title: "Management Commands"},
		&cobra.Group{ID: "swarm", Title: "Swarm Commands"},
	)

	rootCmd.Flags().String("config", "/home/rsteube/.docker", "Location of client config files")
	rootCmd.Flags().StringP("context", "c", "", "Name of the context to use to connect to the daemon (overrides DOCKER_HOST env var and default context set with \"docker context use\")")
	rootCmd.Flags().BoolP("debug", "D", false, "Enable debug mode")
	rootCmd.PersistentFlags().BoolP("help", "h", false, "Print usage")
	rootCmd.Flags().StringP("host", "H", "", "Daemon socket(s) to connect to")
	rootCmd.Flags().StringP("log-level", "l", "info", "Set the logging level (\"debug\", \"info\", \"warn\", \"error\", \"fatal\")")
	rootCmd.Flags().Bool("tls", false, "Use TLS; implied by --tlsverify")
	rootCmd.Flags().String("tlscacert", "/home/rsteube/.docker/ca.pem", "Trust certs signed only by this CA")
	rootCmd.Flags().String("tlscert", "/home/rsteube/.docker/cert.pem", "Path to TLS certificate file")
	rootCmd.Flags().String("tlskey", "/home/rsteube/.docker/key.pem", "Path to TLS key file")
	rootCmd.Flags().Bool("tlsverify", false, "Use TLS and verify the remote")
	rootCmd.Flags().BoolP("version", "v", false, "Print version information and quit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"config":    carapace.ActionFiles(),
		"log-level": carapace.ActionValues("debug", "info", "warn", "error", "fatal").StyleF(style.ForLogLevel),
		"tlscacert": carapace.ActionFiles(),
		"tlscert":   carapace.ActionFiles(),
		"tlskey":    carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PreRun(func(cmd *cobra.Command, args []string) {
		r := regexp.MustCompile(`^  (?P<plugin>[^ ]+)[*] +(?P<description>.*)$`)
		if output, err := (carapace.Context{}).Command("docker", "--help").Output(); err == nil {
			lines := strings.Split(string(output), "\n")
			for _, line := range lines {
				if matches := r.FindStringSubmatch(line); len(matches) > 2 {
					name := matches[1]
					description := matches[2]
					pluginCmd := &cobra.Command{
						Use:                name,
						Short:              description,
						GroupID:            "management",
						Run:                func(cmd *cobra.Command, args []string) {},
						DisableFlagParsing: true,
					}

					// TODO
					// if d := completers.Description("docker-" + name); d != "" {
					// 	pluginCmd.Short = d
					// }

					carapace.Gen(pluginCmd).PositionalAnyCompletion(
						bridge.ActionCarapaceBin("docker-" + name),
					)

					rootCmd.AddCommand(pluginCmd)
				}
			}
		}
	})
}
