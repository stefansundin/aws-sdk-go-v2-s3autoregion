package s3autoregion

import (
	"testing"
)

var regionShortNameTests = [][2]string{
	{"use1", "us-east-1"},
	{"usw2", "us-west-2"},
	{"afs1", "af-south-1"},
	{"ape1", "ap-east-1"},
	{"aps2", "ap-south-2"},
	{"apse5", "ap-southeast-5"},
	{"apne3", "ap-northeast-3"},
	{"cac1", "ca-central-1"},
	{"caw1", "ca-west-1"},
	{"euc1", "eu-central-1"},
	{"euw2", "eu-west-2"},
	{"eus1", "eu-south-1"},
	{"eun1", "eu-north-1"},
	{"ilc1", "il-central-1"},
	{"mes1", "me-south-1"},
	{"mec1", "me-central-1"},
	{"sae1", "sa-east-1"},
	{"usge1", "us-gov-east-1"},
	{"usgw1", "us-gov-west-1"},
}

func Test_expandRegionShortName(t *testing.T) {
	for _, test := range regionShortNameTests {
		shortName, expectedRegion := test[0], test[1]
		region, ok := expandRegionShortName(shortName)
		if !ok {
			t.Fatalf(`expandRegionShortName(%q) did not succeed, expected %q`, shortName, expectedRegion)
		} else if region != expectedRegion {
			t.Fatalf(`expandRegionShortName(%q) = %q, expected %q`, shortName, region, expectedRegion)
		}
	}
}
