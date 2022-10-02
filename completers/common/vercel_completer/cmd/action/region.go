package action

import "github.com/rsteube/carapace"

func ActionRegions() carapace.Action {
	return carapace.ActionValuesDescribed(
		"arn1", "Stockholm, Sweden AWS eu-north-1",
		"bom1", "Mumbai, India AWS ap-south-1",
		"cdg1", "Paris, France AWS eu-west-3",
		"cle1", "Cleveland, USA AWS us-east-2",
		"dub1", "Dublin, Ireland AWS eu-west-1",
		"fra1", "Frankfurt, Germany AWS eu-central-1",
		"gru1", "São Paulo, Brazil AWS sa-east-1",
		"hkg1", "Hong Kong AWS ap-east-1",
		"hnd1", "Tokyo, Japan AWS ap-northeast-1",
		"iad1", "Washington, D.C., USA AWS us-east-1",
		"icn1", "Seoul, South Korea AWS ap-northeast-2",
		"lhr1", "London, United Kingdom AWS eu-west-2",
		"pdx1", "Portland, USA AWS us-west-2",
		"sfo1", "San Francisco, USA AWS us-west-1",
		"sin1", "Singapore AWS ap-southeast-1",
		"syd1", "Sydney, Australia AWS ap-southeast-2",
	)
}
