package action

import "github.com/rsteube/carapace"

func ActionSecretsProvider() carapace.Action {
	return carapace.ActionValues("default", "passphrase", "awskms", "azurekeyvault", "gcpkms", "hashivault")
}
