package mytool

import (
	"golang.org/x/tools/go/analysis"
)

// Analyzer is test analyzer
var Analyzer = &analysis.Analyzer{
	Name: "mytool",
	Doc:  "mytool doc",
	Run: func(pass *analysis.Pass) (interface{}, error) {
		pass.Reportf(pass.Files[0].Package, "mytool warning")
		return nil, nil
	},
}
