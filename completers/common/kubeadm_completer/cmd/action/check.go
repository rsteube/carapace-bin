package action

import "github.com/rsteube/carapace"

func ActionChecks() carapace.Action {
	return carapace.ActionValues(
		"ContainerRuntime",
		"DirAvailable",
		"ExternalEtcdVersion",
		"FileAvailable",
		"FileContent",
		"FileExisting",
		"Firewalld",
		"HTTPProxy",
		"HTTPProxyCIDR",
		"Hostname",
		"ImagePull",
		"InPath",
		"IsPrivilegedUser",
		"KubeletVersion",
		"KubernetesVersion",
		"Mem",
		"NumCPU",
		"PortOpen",
		"Service",
		"Swap",
		"SystemVerification",
		"all",
	)
}
