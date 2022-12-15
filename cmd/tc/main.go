package main

import (
	"github.com/a631807682/tagcheck/tagcheck"
	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
	multichecker.Main(tagcheck.Analyzer)
}
