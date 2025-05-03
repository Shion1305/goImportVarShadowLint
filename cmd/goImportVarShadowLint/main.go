package main

import (
	"github.com/Shion1305/goImportVarShadowLint"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(goImportVarShadowLint.Analyzer)
}
