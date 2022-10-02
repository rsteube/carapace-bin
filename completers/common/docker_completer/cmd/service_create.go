package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var service_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(service_createCmd).Standalone()

	service_createCmd.Flags().String("config", "", "Specify configurations to expose to the service")
	service_createCmd.Flags().String("constraint", "", "Placement constraints")
	service_createCmd.Flags().String("container-label", "", "Container labels")
	service_createCmd.Flags().String("credential-spec", "", "Credential spec for managed service account (Windows")
	service_createCmd.Flags().BoolP("detach", "d", false, "Exit immediately instead of waiting for the service")
	service_createCmd.Flags().String("dns", "", "Set custom DNS servers")
	service_createCmd.Flags().String("dns-option", "", "Set DNS options")
	service_createCmd.Flags().String("dns-search", "", "Set custom DNS search domains")
	service_createCmd.Flags().String("endpoint-mode", "", "Endpoint mode (vip or dnsrr) (default \"vip\")")
	service_createCmd.Flags().String("entrypoint", "", "Overwrite the default ENTRYPOINT of the image")
	service_createCmd.Flags().StringP("env", "e", "", "Set environment variables")
	service_createCmd.Flags().String("env-file", "", "Read in a file of environment variables")
	service_createCmd.Flags().String("generic-resource", "", "User defined resources")
	service_createCmd.Flags().String("group", "", "Set one or more supplementary user groups for the")
	service_createCmd.Flags().String("health-cmd", "", "Command to run to check health")
	service_createCmd.Flags().String("health-interval", "", "Time between running the check (ms|s|m|h)")
	service_createCmd.Flags().String("health-retries", "", "Consecutive failures needed to report unhealthy")
	service_createCmd.Flags().String("health-start-period", "", "Start period for the container to initialize before")
	service_createCmd.Flags().String("health-timeout", "", "Maximum time to allow one check to run (ms|s|m|h)")
	service_createCmd.Flags().String("host", "", "Set one or more custom host-to-IP mappings (host:ip)")
	service_createCmd.Flags().String("hostname", "", "Container hostname")
	service_createCmd.Flags().Bool("init", false, "Use an init inside each service container to forward")
	service_createCmd.Flags().String("isolation", "", "Service container isolation mode")
	service_createCmd.Flags().StringP("label", "l", "", "Service labels")
	service_createCmd.Flags().String("limit-cpu", "", "Limit CPUs")
	service_createCmd.Flags().String("limit-memory", "", "Limit Memory")
	service_createCmd.Flags().String("log-driver", "", "Logging driver for service")
	service_createCmd.Flags().String("log-opt", "", "Logging driver options")
	service_createCmd.Flags().String("mode", "", "Service mode (replicated or global) (default")
	service_createCmd.Flags().String("mount", "", "Attach a filesystem mount to the service")
	service_createCmd.Flags().String("name", "", "Service name")
	service_createCmd.Flags().String("network", "", "Network attachments")
	service_createCmd.Flags().Bool("no-healthcheck", false, "Disable any container-specified HEALTHCHECK")
	service_createCmd.Flags().Bool("no-resolve-image", false, "Do not query the registry to resolve image digest")
	service_createCmd.Flags().String("placement-pref", "", "Add a placement preference")
	service_createCmd.Flags().StringP("publish", "p", "", "Publish a port as a node port")
	service_createCmd.Flags().BoolP("quiet", "q", false, "Suppress progress output")
	service_createCmd.Flags().Bool("read-only", false, "Mount the container's root filesystem as read only")
	service_createCmd.Flags().String("replicas", "", "Number of tasks")
	service_createCmd.Flags().String("replicas-max-per-node", "", "Maximum number of tasks per node (default 0 = unlimited)")
	service_createCmd.Flags().String("reserve-cpu", "", "Reserve CPUs")
	service_createCmd.Flags().String("reserve-memory", "", "Reserve Memory")
	service_createCmd.Flags().String("restart-condition", "", "Restart when condition is met")
	service_createCmd.Flags().String("restart-delay", "", "Delay between restart attempts (ns|us|ms|s|m|h)")
	service_createCmd.Flags().String("restart-max-attempts", "", "Maximum number of restarts before giving up")
	service_createCmd.Flags().String("restart-window", "", "Window used to evaluate the restart policy")
	service_createCmd.Flags().String("rollback-delay", "", "Delay between task rollbacks (ns|us|ms|s|m|h)")
	service_createCmd.Flags().String("rollback-failure-action", "", "Action on rollback failure (\"pause\"|\"continue\")")
	service_createCmd.Flags().String("rollback-max-failure-ratio", "", "Failure rate to tolerate during a rollback (default 0)")
	service_createCmd.Flags().String("rollback-monitor", "", "Duration after each task rollback to monitor for")
	service_createCmd.Flags().String("rollback-order", "", "Rollback order (\"start-first\"|\"stop-first\") (default")
	service_createCmd.Flags().String("rollback-parallelism", "", "Maximum number of tasks rolled back simultaneously")
	service_createCmd.Flags().String("secret", "", "Specify secrets to expose to the service")
	service_createCmd.Flags().String("stop-grace-period", "", "Time to wait before force killing a container")
	service_createCmd.Flags().String("stop-signal", "", "Signal to stop the container")
	service_createCmd.Flags().String("sysctl", "", "Sysctl options")
	service_createCmd.Flags().BoolP("tty", "t", false, "Allocate a pseudo-TTY")
	service_createCmd.Flags().String("update-delay", "", "Delay between updates (ns|us|ms|s|m|h) (default 0s)")
	service_createCmd.Flags().String("update-failure-action", "", "Action on update failure")
	service_createCmd.Flags().String("update-max-failure-ratio", "", "Failure rate to tolerate during an update (default 0)")
	service_createCmd.Flags().String("update-monitor", "", "Duration after each task update to monitor for")
	service_createCmd.Flags().String("update-order", "", "Update order (\"start-first\"|\"stop-first\") (default")
	service_createCmd.Flags().String("update-parallelism", "", "Maximum number of tasks updated simultaneously (0 to")
	service_createCmd.Flags().StringP("user", "u", "", "Username or UID (format: <name|uid>[:<group|gid>])")
	service_createCmd.Flags().Bool("with-registry-auth", false, "Send registry authentication details to swarm agents")
	service_createCmd.Flags().StringP("workdir", "w", "", "Working directory inside the container")
	serviceCmd.AddCommand(service_createCmd)

	carapace.Gen(service_createCmd).PositionalCompletion(
		docker.ActionRepositoryTags(),
	)

	carapace.Gen(service_createCmd).FlagCompletion(carapace.ActionMap{
		"endpoint-mode": carapace.ActionValues("vip", "dnsrr"),
		"env-file":      carapace.ActionFiles(),
		"group":         os.ActionGroups(),
		"isolation":     carapace.ActionValues("default", "hyperv", "process"),
		"log-driver":    docker.ActionLogDrivers(),
		"mode":          carapace.ActionValues("replicated", "global"),
		"user":          os.ActionUsers(),
	})
}
