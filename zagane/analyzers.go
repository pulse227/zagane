package zagane

import (
	"github.com/pulse227/zagane/passes/unclosetx"
	"github.com/pulse227/zagane/passes/unstopiter"
	"github.com/pulse227/zagane/passes/wraperr"
	"golang.org/x/tools/go/analysis"
)

// Analyzers returns analyzers of zagane.
func Analyzers() []*analysis.Analyzer {
	return []*analysis.Analyzer{
		unstopiter.Analyzer,
		unclosetx.Analyzer,
		wraperr.Analyzer,
	}
}
