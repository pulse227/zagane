//go:build go1.12
// +build go1.12

package main

import (
	"github.com/pulse227/zagane/zagane"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() {
	if printVersion() {
		return
	}
	unitchecker.Main(zagane.Analyzers()...)
}
