package goImportVarShadowLint_test

import (
	"github.com/Shion1305/goImportVarShadowLint"
	"golang.org/x/tools/go/analysis/analysistest"
	"testing"
)

func TestPkgShadow(t *testing.T) {
	testdata := analysistest.TestData()

	analysistest.RunWithSuggestedFixes(t, testdata, goImportVarShadowLint.Analyzer, "a")
	analysistest.RunWithSuggestedFixes(t, testdata, goImportVarShadowLint.Analyzer, "a")
}
