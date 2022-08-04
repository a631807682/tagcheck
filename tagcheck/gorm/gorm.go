package gorm

import (
	"fmt"
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
	"honnef.co/go/tools/analysis/report"
)

func CheckTag(pass *analysis.Pass, field *ast.Field, tag string) {
	if len(tag) == 0 {
		return
	}

	options := strings.Split(tag, ";")

	optionNums := make(map[string]int)
	for _, o := range options {
		optKey := strings.TrimSpace(o)
		if len(optKey) == 0 { //is space
			continue
		}
		optVal := ""
		if i := strings.Index(o, ":"); i >= 0 {
			optKey = o[:i]
			optVal = o[i+1:]
		}

		optKey = strings.TrimSpace(optKey)
		optVal = strings.TrimSpace(optVal)

		gt, ok := gormTagMap[strings.ToUpper(optKey)]
		if !ok {
			report.Report(pass, field.Tag, fmt.Sprintf("not support Gorm option %q", optKey))
			continue
		}

		err := gt.checker.check(optVal)
		if err != nil {
			report.Report(pass, field.Tag, fmt.Sprintf("not support Gorm option %q value %q %s", optKey, optVal, err.Error()))
			continue
		}

		optionNums[optKey]++
	}

	for option, n := range optionNums {
		if n > 1 {
			report.Report(pass, field.Tag, fmt.Sprintf("duplicate Gorm option %q", option))
		}
	}
}
