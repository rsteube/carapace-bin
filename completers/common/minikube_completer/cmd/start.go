package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/minikube_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/docker"
	"github.com/rsteube/carapace-bin/pkg/util"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts a local Kubernetes cluster",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	startCmd.Flags().StringSlice("addons", []string{}, "Enable addons. see `minikube addons list` for a list of valid addon names.")
	startCmd.Flags().StringSlice("apiserver-ips", []string{}, "A set of apiserver IP Addresses which are used in the generated certificate for kubernetes.  This can be used if you want to make the apiserver available from outside the machine")
	startCmd.Flags().String("apiserver-name", "minikubeCA", "The authoritative apiserver hostname for apiserver certificates and connectivity. This can be used if you want to make the apiserver available from outside the machine")
	startCmd.Flags().StringSlice("apiserver-names", []string{}, "A set of apiserver names which are used in the generated certificate for kubernetes.  This can be used if you want to make the apiserver available from outside the machine")
	startCmd.Flags().Int("apiserver-port", 8443, "The apiserver listening port")
	startCmd.Flags().Bool("auto-update-drivers", true, "If set, automatically updates drivers to the latest version. Defaults to true.")
	startCmd.Flags().String("base-image", "gcr.io/k8s-minikube/kicbase:v0.0.24@sha256:ba324e0dc025040a8ea6b883d008ec4a43a47db106fb59ac7446982c20c2cdc5", "The base image to use for docker/podman drivers. Intended for local development.")
	startCmd.Flags().Bool("cache-images", true, "If true, cache docker images for the current bootstrapper and load them into the machine. Always false with --driver=none.")
	startCmd.Flags().String("cni", "", "CNI plug-in to use. Valid options: auto, bridge, calico, cilium, flannel, kindnet, or path to a CNI manifest (default: auto)")
	startCmd.Flags().String("container-runtime", "docker", "The container runtime to be used (docker, cri-o, containerd).")
	startCmd.Flags().String("cpus", "2", "Number of CPUs allocated to Kubernetes. Use \"max\" to use the maximum number of CPUs.")
	startCmd.Flags().String("cri-socket", "", "The cri socket path to be used.")
	startCmd.Flags().Bool("delete-on-failure", false, "If set, delete the current cluster if start fails and try again. Defaults to false.")
	startCmd.Flags().Bool("disable-driver-mounts", false, "Disables the filesystem mounts provided by the hypervisors")
	startCmd.Flags().String("disk-size", "20000mb", "Disk size allocated to the minikube VM (format: <number>[<unit>], where unit = b, k, m or g).")
	startCmd.Flags().String("dns-domain", "cluster.local", "The cluster dns domain name used in the Kubernetes cluster")
	startCmd.Flags().Bool("dns-proxy", false, "Enable proxy for NAT DNS requests (virtualbox driver only)")
	startCmd.Flags().StringArray("docker-env", []string{}, "Environment variables to pass to the Docker daemon. (format: key=value)")
	startCmd.Flags().StringArray("docker-opt", []string{}, "Specify arbitrary flags to pass to the Docker daemon. (format: key=value)")
	startCmd.Flags().Bool("download-only", false, "If true, only download and cache files for later use - don't install or start anything.")
	startCmd.Flags().String("driver", "", "Driver is one of: virtualbox, vmwarefusion, kvm2, vmware, none, docker, podman, ssh (defaults to auto-detect)")
	startCmd.Flags().Bool("dry-run", false, "dry-run mode. Validates configuration, but does not mutate system state")
	startCmd.Flags().Bool("embed-certs", false, "if true, will embed the certs in kubeconfig.")
	startCmd.Flags().Bool("enable-default-cni", false, "DEPRECATED: Replaced by --cni=bridge")
	startCmd.Flags().StringArray("extra-config", []string{}, "A set of key=value pairs that describe configuration that may be passed to different components.")
	startCmd.Flags().String("feature-gates", "", "A set of key=value pairs that describe feature gates for alpha/experimental features.")
	startCmd.Flags().Bool("force", false, "Force minikube to perform possibly dangerous operations")
	startCmd.Flags().Bool("force-systemd", false, "If set, force the container runtime to use systemd as cgroup manager. Defaults to false.")
	startCmd.Flags().Bool("host-dns-resolver", true, "Enable host resolver for NAT DNS requests (virtualbox driver only)")
	startCmd.Flags().String("host-only-cidr", "192.168.99.1/24", "The CIDR to be used for the minikube VM (virtualbox driver only)")
	startCmd.Flags().String("host-only-nic-type", "virtio", "NIC Type used for host only network. One of Am79C970A, Am79C973, 82540EM, 82543GC, 82545EM, or virtio (virtualbox driver only)")
	startCmd.Flags().String("hyperkit-vpnkit-sock", "", "Location of the VPNKit socket used for networking. If empty, disables Hyperkit VPNKitSock, if 'auto' uses Docker for Mac VPNKit connection, otherwise uses the specified VSock (hyperkit driver only)")
	startCmd.Flags().StringSlice("hyperkit-vsock-ports", []string{}, "List of guest VSock ports that should be exposed as sockets on the host (hyperkit driver only)")
	startCmd.Flags().String("hyperv-external-adapter", "", "External Adapter on which external switch will be created if no external switch is found. (hyperv driver only)")
	startCmd.Flags().Bool("hyperv-use-external-switch", false, "Whether to use external switch over Default Switch if virtual switch not explicitly specified. (hyperv driver only)")
	startCmd.Flags().String("hyperv-virtual-switch", "", "The hyperv virtual switch name. Defaults to first found. (hyperv driver only)")
	startCmd.Flags().String("image-mirror-country", "", "Country code of the image mirror to be used. Leave empty to use the global one. For Chinese mainland users, set it to cn.")
	startCmd.Flags().String("image-repository", "", "Alternative image repository to pull docker images from. This can be used when you have limited access to gcr.io. Set it to \"auto\" to let minikube decide one for you. For Chinese mainland users, you may use local gcr.io mirrors such as registry.cn-hangzhou.aliyuncs.com/google_containers")
	startCmd.Flags().StringSlice("insecure-registry", []string{}, "Insecure Docker registries to pass to the Docker daemon.  The default service CIDR range will automatically be added.")
	startCmd.Flags().Bool("install-addons", true, "If set, install addons. Defaults to true.")
	startCmd.Flags().Bool("interactive", true, "Allow user prompts for more information")
	startCmd.Flags().StringSlice("iso-url", []string{"https://storage.googleapis.com/minikube/iso/minikube-v0.0.0-unset.iso", "https://github.com/kubernetes/minikube/releases/download/v0.0.0-unset/minikube-v0.0.0-unset.iso", "https://kubernetes.oss-cn-hangzhou.aliyuncs.com/minikube/iso/minikube-v0.0.0-unset.iso"}, "Locations to fetch the minikube ISO from.")
	startCmd.Flags().Bool("keep-context", false, "This will keep the existing kubectl context and will create a minikube context.")
	startCmd.Flags().String("kubernetes-version", "", "The Kubernetes version that the minikube VM will use (ex: v1.2.3, 'stable' for v1.20.7, 'latest' for v1.22.0-alpha.2). Defaults to 'stable'.")
	startCmd.Flags().Bool("kvm-gpu", false, "Enable experimental NVIDIA GPU support in minikube")
	startCmd.Flags().Bool("kvm-hidden", false, "Hide the hypervisor signature from the guest in minikube (kvm2 driver only)")
	startCmd.Flags().String("kvm-network", "default", "The KVM default network name. (kvm2 driver only)")
	startCmd.Flags().Int("kvm-numa-count", 1, "Simulate numa node count in minikube, supported numa node count range is 1-8 (kvm2 driver only)")
	startCmd.Flags().String("kvm-qemu-uri", "qemu:///system", "The KVM QEMU connection URI. (kvm2 driver only)")
	startCmd.Flags().String("listen-address", "", "IP Address to use to expose ports (docker and podman driver only)")
	startCmd.Flags().String("memory", "", "Amount of RAM to allocate to Kubernetes (format: <number>[<unit>], where unit = b, k, m or g). Use \"max\" to use the maximum amount of memory.")
	startCmd.Flags().Bool("mount", false, "This will start the mount daemon and automatically mount files into minikube.")
	startCmd.Flags().String("mount-string", "/home/user:/minikube-host", "The argument to pass the minikube mount command on start.")
	startCmd.Flags().String("namespace", "default", "The named space to activate after start")
	startCmd.Flags().String("nat-nic-type", "virtio", "NIC Type used for nat network. One of Am79C970A, Am79C973, 82540EM, 82543GC, 82545EM, or virtio (virtualbox driver only)")
	startCmd.Flags().Bool("native-ssh", true, "Use native Golang SSH client (default true). Set to 'false' to use the command line 'ssh' command when accessing the docker machine. Useful for the machine drivers when they will not start with 'Waiting for SSH'.")
	startCmd.Flags().String("network", "", "network to run minikube with. Now it is used by docker/podman and KVM drivers. If left empty, minikube will create a new network.")
	startCmd.Flags().String("network-plugin", "", "Kubelet network plug-in to use (default: auto)")
	startCmd.Flags().StringSlice("nfs-share", []string{}, "Local folders to share with Guest via NFS mounts (hyperkit driver only)")
	startCmd.Flags().String("nfs-shares-root", "/nfsshares", "Where to root the NFS Shares, defaults to /nfsshares (hyperkit driver only)")
	startCmd.Flags().Bool("no-vtx-check", false, "Disable checking for the availability of hardware virtualization before the vm is started (virtualbox driver only)")
	startCmd.Flags().IntP("nodes", "n", 1, "The number of nodes to spin up. Defaults to 1.")
	startCmd.Flags().StringP("output", "o", "text", "Format to print stdout in. Options include: [text,json]")
	startCmd.Flags().StringSlice("ports", []string{}, "List of ports that should be exposed (docker and podman driver only)")
	startCmd.Flags().Bool("preload", true, "If set, download tarball of preloaded images if available to improve start time. Defaults to true.")
	startCmd.Flags().StringSlice("registry-mirror", []string{}, "Registry mirrors to pass to the Docker daemon")
	startCmd.Flags().String("service-cluster-ip-range", "10.96.0.0/12", "The CIDR to be used for service cluster IPs.")
	startCmd.Flags().String("ssh-ip-address", "", "IP address (ssh driver only)")
	startCmd.Flags().String("ssh-key", "", "SSH key (ssh driver only)")
	startCmd.Flags().Int("ssh-port", 22, "SSH port (ssh driver only)")
	startCmd.Flags().String("ssh-user", "root", "SSH user (ssh driver only)")
	startCmd.Flags().String("trace", "", "Send trace events. Options include: [gcp]")
	startCmd.Flags().String("uuid", "", "Provide VM UUID to restore MAC address (hyperkit driver only)")
	startCmd.Flags().Bool("vm", false, "Filter to use only VM Drivers")
	startCmd.Flags().String("vm-driver", "", "DEPRECATED, use `driver` instead.")
	startCmd.Flags().StringSlice("wait", []string{"apiserver", "system_pods"}, "comma separated list of Kubernetes components to verify and wait for after starting a cluster. defaults to \"apiserver,system_pods\", available options: \"apiserver,system_pods,default_sa,apps_running,node_ready,kubelet\" . other acceptable values are 'all' or 'none', 'true' and 'false'")
	startCmd.Flags().String("wait-timeout", "6m0s", "max time to wait per Kubernetes or host to be healthy.")
	rootCmd.AddCommand(startCmd)

	carapace.Gen(startCmd).FlagCompletion(carapace.ActionMap{
		"addons": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionAddons().Invoke(c).Filter(c.Parts).ToA()
		}),
		"base-image": docker.ActionRepositoryTags(),
		"cni": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if util.HasPathPrefix(c.CallbackValue) {
				return carapace.ActionFiles()
			}
			return carapace.ActionValues("auto", "bridge", "calico", "cilium", "flannel", "kindnet")
		}),
		"container-runtime":  carapace.ActionValues("docker", "cri-o", "containerd"),
		"cri-socket":         carapace.ActionFiles(),
		"driver":             carapace.ActionValues("virtualbox", "vmwarefusion", "kvm2", "vmware", "none", "docker", "podman", "ssh", "auto-detect"),
		"host-only-nic-type": carapace.ActionValues("Am79C970A", "Am79C973", "82540EM", "82543GC", "82545EM", "virtio"),
		"hyperkit-vpnkit-sock": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if util.HasPathPrefix(c.CallbackValue) {
				return carapace.ActionFiles()
			}
			return carapace.ActionValues("auto")
		}),
		"image-mirror-country": os.ActionLanguages(), // TODO country codes the same?
		"mount-string": carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionFiles()
			case 1:
				return carapace.ActionValues("/minikube-host")
			default:
				return carapace.ActionValues()
			}
		}),
		"nat-nic-type": carapace.ActionValues("Am79C970A", "Am79C973", "82540EM", "82543GC", "82545EM", "virtio"),
		"output":       carapace.ActionValues("text", "json"),
		"ssh-key":      carapace.ActionFiles(),
		"ssh-user":     os.ActionUsers(),
		"wait": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("apiserver", "system_pods", "default_sa", "apps_running", "node_ready", "kubelet", "all", "none").Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
