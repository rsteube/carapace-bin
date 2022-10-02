package action

import (
	"strings"

	"github.com/rsteube/carapace"
)

func ActionApiResources() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("kubectl", "api-resources", "--output=name", "--cached")(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			for index, line := range lines {
				lines[index] = strings.SplitN(line, ".", 2)[0]
			}
			return carapace.ActionValues(lines[:len(lines)-1]...)
		})
	})
}

func ActionResources(namespace string, types string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("kubectl", "--namespace", namespace, "get", "-o", "go-template={{range .items}}{{.metadata.name}}\n{{.kind}}\n{{end}}", types)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			return carapace.ActionValuesDescribed(lines[:len(lines)-1]...)
		})
	})
}

func ActionDryRunModes() carapace.Action {
	return carapace.ActionValues("none", "server", "client")
}

func ActionOutputFormats() carapace.Action {
	return carapace.ActionValues("json", "yaml", "name", "go-template", "go-template-file", "template", "templatefile", "jsonpath", "jsonpath-as-json", "jsonpath-file")
}

func ActionResourceVerbs() carapace.Action {
	return carapace.ActionValues("get", "list", "create", "update", "patch", "watch", "delete", "deletecollection")
}

func ActionNamespaceServiceAccounts() carapace.Action {
	return carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return ActionResources("", "namespaces").Invoke(c).Suffix(":").ToA()
		case 1:
			return ActionResources(c.Parts[0], "serviceaccounts")
		default:
			return carapace.ActionValues()
		}
	})
}

func ActionApiResourceResources() carapace.Action {
	return carapace.ActionMultiParts("/", func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return ActionApiResources().Invoke(c).Suffix("/").ToA()
		case 1:
			return ActionResources("", c.Parts[0])
		default:
			return carapace.ActionValues()
		}
	})
}

func ActionContainers(namespace string, resource string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("kubectl", "--namespace", namespace, "get", "-o", "go-template={{range .spec.containers}}{{.name}}\n{{end}}", resource)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			return carapace.ActionValues(lines[:len(lines)-1]...)
		})
	})
}

func ActionLabels(namespace string, resource string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("kubectl", "--namespace", namespace, "get", "-o", "go-template={{range $key, $value := .metadata.labels}}{{$key}}\n{{$value}}\n{{end}}", resource)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			return carapace.ActionValuesDescribed(lines[:len(lines)-1]...)
		})
	})
}

func ActionAnnotations(namespace string, resource string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("kubectl", "--namespace", namespace, "get", "-o", "go-template={{range $key, $value := .metadata.annotations}}{{$key}}\n{{$value}}\n{{end}}", resource)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			return carapace.ActionValuesDescribed(lines[:len(lines)-1]...)
		})
	})
}

func ActionApiGroups() carapace.Action {
	return carapace.ActionValues(
		"admissionregistration.k8s.io",
		"apiextensions.k8s.io",
		"apiregistration.k8s.io",
		"apps",
		"authentication.k8s.io",
		"authorization.k8s.io",
		"autoscaling",
		"batch",
		"certificates.k8s.io",
		"coordination.k8s.io",
		"core",
		"discovery.k8s.io",
		"events.k8s.io",
		"extensions",
		"flowcontrol.apiserver.k8s.io",
		"networking.k8s.io",
		"node.k8s.io",
		"policy",
		"rbac.authorization.k8s.io",
		"scheduling.k8s.io",
		"settings.k8s.io",
		"storage.k8s.io",
	)
}

func ActionClusters() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("kubectl", "config", "get-clusters")(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			return carapace.ActionValues(lines[1 : len(lines)-1]...)
		})
	})
}

func ActionContexts() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("kubectl", "config", "get-contexts", "-o", "name")(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			return carapace.ActionValues(lines[:len(lines)-1]...)
		})
	})
}
