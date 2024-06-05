package unclosetx_test

import (
	"testing"

	"github.com/gostaticanalysis/testutil"
	"github.com/pulse227/zagane/passes/unclosetx"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	vers := testutil.LatestVersion(t, "cloud.google.com/go/spanner", 2)
	testutil.RunWithVersions(t, analysistest.TestData(), unclosetx.Analyzer, vers, "a")
}
