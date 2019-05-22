package main

import (
	"github.com/tamayika/analyzer/pkg/mytool"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() {
	unitchecker.Main(mytool.Analyzer)
}
