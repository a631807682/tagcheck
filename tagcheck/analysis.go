package tagcheck

import (
	"fmt"
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"honnef.co/go/tools/analysis/code"
	"honnef.co/go/tools/analysis/lint"
	"honnef.co/go/tools/analysis/report"

	"github.com/a631807682/tagcheck/tagcheck/gorm"
)

var Analyzers = lint.InitializeAnalyzers(Docs, map[string]*analysis.Analyzer{
	"STC1001": {
		Run:      CheckStructTags,
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	},
})

func CheckStructTags(pass *analysis.Pass) (interface{}, error) {
	fn := func(node ast.Node) {
		structNode := node.(*ast.StructType)
		for _, field := range structNode.Fields.List {
			if field.Tag == nil {
				continue
			}
			tags, err := parseStructTag(field.Tag.Value[1 : len(field.Tag.Value)-1])
			if err != nil {
				report.Report(pass, field.Tag, fmt.Sprintf("unparseable struct tag: %s", err))
				continue
			}
			for k, v := range tags {
				switch k {
				case "gorm":
					gorm.CheckTag(pass, field, v[0])
				}
			}
		}
	}
	code.Preorder(pass, fn, (*ast.StructType)(nil))
	return nil, nil
}
