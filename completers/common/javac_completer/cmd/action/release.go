package action

import "github.com/rsteube/carapace"

func ActionReleases() carapace.Action {
	return carapace.ActionValuesDescribed(
		"1.3", "The compiler does not support features introduced after Java SE 1.3.",
		"1.4", "The compiler accepts code containing assertions, which were introduced in Java SE 1.4.",
		"1.5", "The compiler accepts code containing generics and other language features introduced in Java SE 5.",
		"5", "Synonym for 1.5.",
		"1.6", "No language changes were introduced in Java SE 6.",
		"6", "Synonym for 1.6.",
		"1.7", "The compiler accepts code with features introduced in Java SE 7.",
		"7", "Synonym for 1.7.",
		"1.8", "The compiler accepts code with features introduced in Java SE 8.",
		"8", "Synonym for 1.8.",
	)
}
