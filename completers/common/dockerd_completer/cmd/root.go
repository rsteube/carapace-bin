package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dockerd",
	Short: "A self-sufficient runtime for containers",
	Long:  "https://docs.docker.com/engine/reference/commandline/dockerd/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("add-runtime", "", "Register an additional OCI compatible runtime (default [])")
	rootCmd.Flags().String("allow-nondistributable-artifacts", "", "Allow push of nondistributable artifacts to registry")
	rootCmd.Flags().String("api-cors-header", "", "Set CORS headers in the Engine API")
	rootCmd.Flags().String("authorization-plugin", "", "Authorization plugins to load")
	rootCmd.Flags().String("bip", "", "Specify network bridge IP")
	rootCmd.Flags().StringP("bridge", "b", "", "Attach containers to a network bridge")
	rootCmd.Flags().String("cgroup-parent", "", "Set parent cgroup for all containers")
	rootCmd.Flags().String("config-file", "", "Daemon configuration file (default \"/etc/docker/daemon.json\")")
	rootCmd.Flags().String("containerd", "", "containerd grpc address")
	rootCmd.Flags().String("containerd-namespace", "", "Containerd namespace to use (default \"moby\")")
	rootCmd.Flags().String("containerd-plugins-namespace", "", "Containerd namespace to use for plugins (default \"plugins.moby\")")
	rootCmd.Flags().String("cpu-rt-period", "", "Limit the CPU real-time period in microseconds for the parent cgroup for all containers")
	rootCmd.Flags().String("cpu-rt-runtime", "", "Limit the CPU real-time runtime in microseconds for the parent cgroup for all containers")
	rootCmd.Flags().Bool("cri-containerd", false, "start containerd with cri")
	rootCmd.Flags().String("data-root", "", "Root directory of persistent Docker state (default \"/var/lib/docker\")")
	rootCmd.Flags().BoolP("debug", "D", false, "Enable debug mode")
	rootCmd.Flags().String("default-address-pool", "", "Default address pools for node specific local networks")
	rootCmd.Flags().String("default-cgroupns-mode", "", "Default mode for containers cgroup namespace (\"host\" | \"private\") (default \"host\")")
	rootCmd.Flags().String("default-gateway", "", "Container default gateway IPv4 address")
	rootCmd.Flags().String("default-gateway-v6", "", "Container default gateway IPv6 address")
	rootCmd.Flags().String("default-ipc-mode", "", "Default mode for containers ipc (\"shareable\" | \"private\") (default \"private\")")
	rootCmd.Flags().String("default-runtime", "", "Default OCI runtime for containers (default \"runc\")")
	rootCmd.Flags().String("default-shm-size", "", "Default shm size for containers (default 64MiB)")
	rootCmd.Flags().String("default-ulimit", "", "Default ulimits for containers (default [])")
	rootCmd.Flags().String("dns", "", "DNS server to use")
	rootCmd.Flags().String("dns-opt", "", "DNS options to use")
	rootCmd.Flags().String("dns-search", "", "DNS search domains to use")
	rootCmd.Flags().String("exec-opt", "", "Runtime execution options")
	rootCmd.Flags().String("exec-root", "", "Root directory for execution state files (default \"/var/run/docker\")")
	rootCmd.Flags().Bool("experimental", false, "Enable experimental features")
	rootCmd.Flags().String("fixed-cidr", "", "IPv4 subnet for fixed IPs")
	rootCmd.Flags().String("fixed-cidr-v6", "", "IPv6 subnet for fixed IPs")
	rootCmd.Flags().StringP("group", "G", "", "Group for the unix socket (default \"docker\")")
	rootCmd.Flags().Bool("help", false, "Print usage")
	rootCmd.Flags().StringP("host", "H", "", "Daemon socket(s) to connect to")
	rootCmd.Flags().String("host-gateway-ip", "", "IP address that the special 'host-gateway' string in --add-host resolves to. Defaults to the IP address of the default bridge")
	rootCmd.Flags().Bool("icc", false, "Enable inter-container communication (default true)")
	rootCmd.Flags().Bool("init", false, "Run an init in the container to forward signals and reap processes")
	rootCmd.Flags().String("init-path", "", "Path to the docker-init binary")
	rootCmd.Flags().String("insecure-registry", "", "Enable insecure registry communication")
	rootCmd.Flags().String("ip", "", "Default IP when binding container ports (default 0.0.0.0)")
	rootCmd.Flags().Bool("ip-forward", false, "Enable net.ipv4.ip_forward (default true)")
	rootCmd.Flags().Bool("ip-masq", false, "Enable IP masquerading (default true)")
	rootCmd.Flags().Bool("ip6tables", false, "Enable addition of ip6tables rules")
	rootCmd.Flags().Bool("iptables", false, "Enable addition of iptables rules (default true)")
	rootCmd.Flags().Bool("ipv6", false, "Enable IPv6 networking")
	rootCmd.Flags().String("label", "", "Set key=value labels to the daemon")
	rootCmd.Flags().Bool("live-restore", false, "Enable live restore of docker when containers are still running")
	rootCmd.Flags().String("log-driver", "", "Default driver for container logs (default \"json-file\")")
	rootCmd.Flags().StringP("log-level", "l", "", "Set the logging level (\"debug\"|\"info\"|\"warn\"|\"error\"|\"fatal\") (default \"info\")")
	rootCmd.Flags().String("log-opt", "", "Default log driver options for containers (default map[])")
	rootCmd.Flags().String("max-concurrent-downloads", "", "Set the max concurrent downloads for each pull (default 3)")
	rootCmd.Flags().String("max-concurrent-uploads", "", "Set the max concurrent uploads for each push (default 5)")
	rootCmd.Flags().String("max-download-attempts", "", "Set the max download attempts for each pull (default 5)")
	rootCmd.Flags().String("metrics-addr", "", "Set default address and port to serve the metrics api on")
	rootCmd.Flags().String("mtu", "", "Set the containers network MTU")
	rootCmd.Flags().String("network-control-plane-mtu", "", "Network Control plane MTU (default 1500)")
	rootCmd.Flags().Bool("no-new-privileges", false, "Set no-new-privileges by default for new containers")
	rootCmd.Flags().String("node-generic-resource", "", "Advertise user-defined resource")
	rootCmd.Flags().String("oom-score-adjust", "", "Set the oom_score_adj for the daemon")
	rootCmd.Flags().StringP("pidfile", "p", "", "Path to use for daemon PID file (default \"/var/run/docker.pid\")")
	rootCmd.Flags().Bool("raw-logs", false, "Full timestamps without ANSI coloring")
	rootCmd.Flags().String("registry-mirror", "", "Preferred Docker registry mirror")
	rootCmd.Flags().Bool("rootless", false, "Enable rootless mode; typically used with RootlessKit")
	rootCmd.Flags().String("seccomp-profile", "", "Path to seccomp profile")
	rootCmd.Flags().Bool("selinux-enabled", false, "Enable selinux support")
	rootCmd.Flags().String("shutdown-timeout", "", "Set the default shutdown timeout (default 15)")
	rootCmd.Flags().StringP("storage-driver", "s", "", "Storage driver to use")
	rootCmd.Flags().String("storage-opt", "", "Storage driver options")
	rootCmd.Flags().String("swarm-default-advertise-addr", "", "Set default address or interface for swarm advertised address")
	rootCmd.Flags().Bool("tls", false, "Use TLS; implied by --tlsverify")
	rootCmd.Flags().String("tlscacert", "", "Trust certs signed only by this CA (default \"/home/user/.docker/ca.pem\")")
	rootCmd.Flags().String("tlscert", "", "Path to TLS certificate file (default \"/home/user/.docker/cert.pem\")")
	rootCmd.Flags().String("tlskey", "", "Path to TLS key file (default \"/home/user/.docker/key.pem\")")
	rootCmd.Flags().Bool("tlsverify", false, "Use TLS and verify the remote")
	rootCmd.Flags().Bool("userland-proxy", false, "Use userland proxy for loopback traffic (default true)")
	rootCmd.Flags().String("userland-proxy-path", "", "Path to the userland proxy binary")
	rootCmd.Flags().String("userns-remap", "", "User/Group setting for user namespaces")
	rootCmd.Flags().BoolP("version", "v", false, "Print version information and quit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"add-runtime": carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionValues("custom", "runc").Invoke(c).Suffix("=").ToA()
			case 1:
				return carapace.ActionFiles()
			default:
				return carapace.ActionValues()
			}
		}),
		"cgroup-parent":         os.ActionCgroups(),
		"config-file":           carapace.ActionFiles(".json"),
		"data-root":             carapace.ActionDirectories(),
		"default-cgroupns-mode": carapace.ActionValues("host", "private"),
		"default-ipc-mode":      carapace.ActionValues("shareable", "private"),
		"exec-root":             carapace.ActionDirectories(),
		"group":                 os.ActionGroups(),
		"init-path":             carapace.ActionDirectories(),
		"log-level":             carapace.ActionValues("debug", "info", "warn", "error", "fatal").StyleF(style.ForLogLevel),
		"pidfile":               carapace.ActionFiles(),
		"userland-proxy-path":   carapace.ActionFiles(),
	})
}
