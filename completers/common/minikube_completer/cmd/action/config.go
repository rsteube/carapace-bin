package action

import (
	"encoding/json"

	"github.com/rsteube/carapace"
)

func ActionConfigNames() carapace.Action {
	return carapace.ActionValues(
		"driver",
		"vm-driver",
		"container-runtime",
		"feature-gates",
		"v",
		"cpus",
		"disk-size",
		"host-only-cidr",
		"memory",
		"log_dir",
		"kubernetes-version",
		"iso-url",
		"WantUpdateNotification",
		"WantBetaUpdateNotification",
		"ReminderWaitPeriodInHours",
		"WantReportError",
		"WantReportErrorPrompt",
		"WantKubectlDownloadMsg",
		"WantNoneDriverWarning",
		"profile",
		"bootstrapper",
		"ShowDriverDeprecationNotification",
		"ShowBootstrapperDeprecationNotification",
		"insecure-registry",
		"hyperv-virtual-switch",
		"disable-driver-mounts",
		"cache",
		"embed-certs",
	)
}

func ActionConfigDefaults(config string) carapace.Action {
	return carapace.ActionExecCommand("minikube", "config", "defaults", config, "--output", "json")(func(output []byte) carapace.Action {
		var vals []string
		if err := json.Unmarshal(output, &vals); err != nil {
			return carapace.ActionMessage(err.Error())
		}
		return carapace.ActionValues(vals...)
	})
}
