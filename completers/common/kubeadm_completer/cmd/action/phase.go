package action

import "github.com/rsteube/carapace"

func ActionPhases() carapace.Action {
	return carapace.ActionValuesDescribed(
		"preflight", "Run pre-flight checks",
		"certs", "Certificate generation",
		"certs/ca", "Generate the self-signed Kubernetes CA to provision identities for other Kubernetes components",
		"certs/apiserver", "Generate the certificate for serving the Kubernetes API",
		"certs/apiserver-kubelet-client", "Generate the certificate for the API server to connect to kubelet",
		"certs/front-proxy-ca", "Generate the self-signed CA to provision identities for front proxy",
		"certs/front-proxy-client", "Generate the certificate for the front proxy client",
		"certs/etcd-ca", "Generate the self-signed CA to provision identities for etcd",
		"certs/etcd-server", "Generate the certificate for serving etcd",
		"certs/etcd-peer", "Generate the certificate for etcd nodes to communicate with each other",
		"certs/etcd-healthcheck-client", "Generate the certificate for liveness probes to healthcheck etcd",
		"certs/apiserver-etcd-client", "Generate the certificate the apiserver uses to access etcd",
		"certs/sa", "Generate a private key for signing service account tokens along with its public key",
		"kubeconfig", "Generate all kubeconfig files necessary to establish the control plane and the admin kubeconfig file",
		"kubeconfig/admin", "Generate a kubeconfig file for the admin to use and for kubeadm itself",
		"kubeconfig/kubelet", "Generate a kubeconfig file for the kubelet to use *only* for cluster bootstrapping purposes",
		"kubeconfig/controller-manager", "Generate a kubeconfig file for the controller manager to use",
		"kubeconfig/scheduler", "Generate a kubeconfig file for the scheduler to use",
		"kubelet-start", "Write kubelet settings and (re)start the kubelet",
		"control-plane", "Generate all static Pod manifest files necessary to establish the control plane",
		"control-plane/apiserver", "Generates the kube-apiserver static Pod manifest",
		"control-plane/controller-manager", "Generates the kube-controller-manager static Pod manifest",
		"control-plane/scheduler", "Generates the kube-scheduler static Pod manifest",
		"etcd", "Generate static Pod manifest file for local etcd",
		"etcd/local", "Generate the static Pod manifest file for a local, single-node local etcd instance",
		"upload-config", "Upload the kubeadm and kubelet configuration to a ConfigMap",
		"upload-config/kubeadm", "Upload the kubeadm ClusterConfiguration to a ConfigMap",
		"upload-config/kubelet", "Upload the kubelet component config to a ConfigMap",
		"upload-certs", "Upload certificates to kubeadm-certs",
		"mark-control-plane", "Mark a node as a control-plane",
		"bootstrap-token", "Generates bootstrap tokens used to join a node to a cluster",
		"kubelet-finalize", "Updates settings relevant to the kubelet after TLS bootstrap",
		"kubelet-finalize/experimental-cert-rotation", "Enable kubelet client certificate rotation",
		"addon", "Install required addons for passing conformance tests",
		"addon/coredns", "Install the CoreDNS addon to a Kubernetes cluster",
		"addon/kube-proxy", "Install the kube-proxy addon to a Kubernetes cluster",
	)
}
