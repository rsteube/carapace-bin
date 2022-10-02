package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var service_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(service_updateCmd).Standalone()

	service_updateCmd.Flags().String("args", "", "Service command args")
	service_updateCmd.Flags().String("config-add", "", "Add or update a config file on a service")
	service_updateCmd.Flags().String("config-rm", "", "Remove a configuration file")
	service_updateCmd.Flags().String("constraint-add", "", "Add or update a placement constraint")
	service_updateCmd.Flags().String("constraint-rm", "", "Remove a constraint")
	service_updateCmd.Flags().String("container-label-add", "", "Add or update a container label")
	service_updateCmd.Flags().String("container-label-rm", "", "Remove a container label by its key")
	service_updateCmd.Flags().String("credential-spec", "", "Credential spec for managed service account (Windows only)")
	service_updateCmd.Flags().BoolP("detach", "d", false, "Exit immediately instead of waiting for the service to converge")
	service_updateCmd.Flags().String("dns-add", "", "Add or update a custom DNS server")
	service_updateCmd.Flags().String("dns-option-add", "", "Add or update a DNS option")
	service_updateCmd.Flags().String("dns-option-rm", "", "Remove a DNS option")
	service_updateCmd.Flags().String("dns-rm", "", "Remove a custom DNS server")
	service_updateCmd.Flags().String("dns-search-add", "", "Add or update a custom DNS search domain")
	service_updateCmd.Flags().String("dns-search-rm", "", "Remove a DNS search domain")
	service_updateCmd.Flags().String("endpoint-mode", "", "Endpoint mode (vip or dnsrr)")
	service_updateCmd.Flags().String("entrypoint", "", "Overwrite the default ENTRYPOINT of the image")
	service_updateCmd.Flags().String("env-add", "", "Add or update an environment variable")
	service_updateCmd.Flags().String("env-rm", "", "Remove an environment variable")
	service_updateCmd.Flags().Bool("force", false, "Force update even if no changes require it")
	service_updateCmd.Flags().String("generic-resource-add", "", "Add a Generic resource")
	service_updateCmd.Flags().String("generic-resource-rm", "", "Remove a Generic resource")
	service_updateCmd.Flags().String("group-add", "", "Add an additional supplementary user group to the container")
	service_updateCmd.Flags().String("group-rm", "", "Remove a previously added supplementary user group from the container")
	service_updateCmd.Flags().String("health-cmd", "", "Command to run to check health")
	service_updateCmd.Flags().String("health-interval", "", "Time between running the check (ms|s|m|h)")
	service_updateCmd.Flags().String("health-retries", "", "Consecutive failures needed to report unhealthy")
	service_updateCmd.Flags().String("health-start-period", "", "Start period for the container to initialize before counting retries towards unstable (ms|s|m|h)")
	service_updateCmd.Flags().String("health-timeout", "", "Maximum time to allow one check to run (ms|s|m|h)")
	service_updateCmd.Flags().String("host-add", "", "Add a custom host-to-IP mapping (host:ip)")
	service_updateCmd.Flags().String("host-rm", "", "Remove a custom host-to-IP mapping (host:ip)")
	service_updateCmd.Flags().String("hostname", "", "Container hostname")
	service_updateCmd.Flags().String("image", "", "Service image tag")
	service_updateCmd.Flags().Bool("init", false, "Use an init inside each service container to forward signals and reap processes")
	service_updateCmd.Flags().String("isolation", "", "Service container isolation mode")
	service_updateCmd.Flags().String("label-add", "", "Add or update a service label")
	service_updateCmd.Flags().String("label-rm", "", "Remove a label by its key")
	service_updateCmd.Flags().String("limit-cpu", "", "Limit CPUs")
	service_updateCmd.Flags().String("limit-memory", "", "Limit Memory")
	service_updateCmd.Flags().String("log-driver", "", "Logging driver for service")
	service_updateCmd.Flags().String("log-opt", "", "Logging driver options")
	service_updateCmd.Flags().String("mount-add", "", "Add or update a mount on a service")
	service_updateCmd.Flags().String("mount-rm", "", "Remove a mount by its target path")
	service_updateCmd.Flags().String("network-add", "", "Add a network")
	service_updateCmd.Flags().String("network-rm", "", "Remove a network")
	service_updateCmd.Flags().Bool("no-healthcheck", false, "Disable any container-specified HEALTHCHECK")
	service_updateCmd.Flags().Bool("no-resolve-image", false, "Do not query the registry to resolve image digest and supported platforms")
	service_updateCmd.Flags().String("placement-pref-add", "", "Add a placement preference")
	service_updateCmd.Flags().String("placement-pref-rm", "", "Remove a placement preference")
	service_updateCmd.Flags().String("publish-add", "", "Add or update a published port")
	service_updateCmd.Flags().String("publish-rm", "", "Remove a published port by its target port")
	service_updateCmd.Flags().BoolP("quiet", "q", false, "Suppress progress output")
	service_updateCmd.Flags().Bool("read-only", false, "Mount the container's root filesystem as read only")
	service_updateCmd.Flags().String("replicas", "", "Number of tasks")
	service_updateCmd.Flags().String("replicas-max-per-node", "", "Maximum number of tasks per node (default 0 = unlimited)")
	service_updateCmd.Flags().String("reserve-cpu", "", "Reserve CPUs")
	service_updateCmd.Flags().String("reserve-memory", "", "Reserve Memory")
	service_updateCmd.Flags().String("restart-condition", "", "Restart when condition is met (\"none\"|\"on-failure\"|\"any\")")
	service_updateCmd.Flags().String("restart-delay", "", "Delay between restart attempts (ns|us|ms|s|m|h)")
	service_updateCmd.Flags().String("restart-max-attempts", "", "Maximum number of restarts before giving up")
	service_updateCmd.Flags().String("restart-window", "", "Window used to evaluate the restart policy (ns|us|ms|s|m|h)")
	service_updateCmd.Flags().Bool("rollback", false, "Rollback to previous specification")
	service_updateCmd.Flags().String("rollback-delay", "", "Delay between task rollbacks (ns|us|ms|s|m|h)")
	service_updateCmd.Flags().String("rollback-failure-action", "", "Action on rollback failure (\"pause\"|\"continue\")")
	service_updateCmd.Flags().String("rollback-max-failure-ratio", "", "Failure rate to tolerate during a rollback")
	service_updateCmd.Flags().String("rollback-monitor", "", "Duration after each task rollback to monitor for failure (ns|us|ms|s|m|h)")
	service_updateCmd.Flags().String("rollback-order", "", "Rollback order (\"start-first\"|\"stop-first\")")
	service_updateCmd.Flags().String("rollback-parallelism", "", "Maximum number of tasks rolled back simultaneously (0 to roll back all at once)")
	service_updateCmd.Flags().String("secret-add", "", "Add or update a secret on a service")
	service_updateCmd.Flags().String("secret-rm", "", "Remove a secret")
	service_updateCmd.Flags().String("stop-grace-period", "", "Time to wait before force killing a container (ns|us|ms|s|m|h)")
	service_updateCmd.Flags().String("stop-signal", "", "Signal to stop the container")
	service_updateCmd.Flags().String("sysctl-add", "", "Add or update a Sysctl option")
	service_updateCmd.Flags().String("sysctl-rm", "", "Remove a Sysctl option")
	service_updateCmd.Flags().BoolP("tty", "t", false, "Allocate a pseudo-TTY")
	service_updateCmd.Flags().String("update-delay", "", "Delay between updates (ns|us|ms|s|m|h)")
	service_updateCmd.Flags().String("update-failure-action", "", "Action on update failure (\"pause\"|\"continue\"|\"rollback\")")
	service_updateCmd.Flags().String("update-max-failure-ratio", "", "Failure rate to tolerate during an update")
	service_updateCmd.Flags().String("update-monitor", "", "Duration after each task update to monitor for failure (ns|us|ms|s|m|h)")
	service_updateCmd.Flags().String("update-order", "", "Update order (\"start-first\"|\"stop-first\")")
	service_updateCmd.Flags().String("update-parallelism", "", "Maximum number of tasks updated simultaneously (0 to update all at once)")
	service_updateCmd.Flags().StringP("user", "u", "", "Username or UID (format: <name|uid>[:<group|gid>])")
	service_updateCmd.Flags().Bool("with-registry-auth", false, "Send registry authentication details to swarm agents")
	service_updateCmd.Flags().StringP("workdir", "w", "", "Working directory inside the container")
	serviceCmd.AddCommand(service_updateCmd)

	carapace.Gen(service_updateCmd).FlagCompletion(carapace.ActionMap{
		"group-add":  os.ActionGroups(),
		"log-driver": docker.ActionLogDrivers(),
		"user":       os.ActionUsers(),
	})

	carapace.Gen(service_updateCmd).PositionalCompletion(
		docker.ActionServices(),
	)
}
