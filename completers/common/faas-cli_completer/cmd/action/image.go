package action

import "github.com/rsteube/carapace"

func ActionImageArchitectures() carapace.Action {
	return carapace.ActionValuesDescribed(
		"amd64", "Linux x86-64",
		"arm32v5", "ARMv5 32-bit",
		"arm32v6", "ARMv6 32-bit",
		"arm32v7", "ARMv7 32-bit",
		"arm64v8", "ARMv8 64-bit",
		"i386", "x86/i686",
		"mips64le", "MIPS64 LE",
		"ppc64le", "IBM POWER8",
		"s390x", "IBM z Systems",
		"windows-amd64", "Windows x86-64",
		"x86_64", "Linux x86-64",
	)
}
