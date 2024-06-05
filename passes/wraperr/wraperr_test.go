package wraperr_test

import (
	"testing"

	"github.com/gostaticanalysis/testutil"
	"github.com/pulse227/zagane/passes/wraperr"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	vers := testutil.LatestVersion(t, "cloud.google.com/go/spanner", 2)
	testutil.RunWithVersions(t, analysistest.TestData(), wraperr.Analyzer, vers, "a")
}
