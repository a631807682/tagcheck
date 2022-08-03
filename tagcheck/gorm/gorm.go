package gorm

import (
	"fmt"
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
	"honnef.co/go/tools/analysis/report"
)

func init() {
	initGormTagChecker()
}

func CheckTag(pass *analysis.Pass, field *ast.Field, tag string) {
	if len(tag) == 0 {
		return
	}

	fields := strings.Split(tag, ";")

	options := make(map[string]int)
	for _, f := range fields {
		fkey := strings.TrimSpace(f)
		if len(fkey) == 0 { //is space
			continue
		}
		fvalue := ""
		if i := strings.Index(f, ":"); i >= 0 {
			fkey = f[:i]
			fvalue = f[i+1:]
		}

		fkey = strings.TrimSpace(fkey)
		fvalue = strings.TrimSpace(fvalue)

		gt, ok := gormTagMap[strings.ToUpper(fkey)]
		if !ok {
			report.Report(pass, field.Tag, fmt.Sprintf("not support Gorm option %q", fkey))
			continue
		}

		err := gt.checker.check(fvalue)
		if err != nil {
			report.Report(pass, field.Tag, fmt.Sprintf("not support Gorm value:%q err:%q", fvalue, err.Error()))
			continue
		}

		options[fkey]++
	}

	for option, n := range options {
		if n > 1 {
			report.Report(pass, field.Tag, fmt.Sprintf("duplicate Gorm option %q", option))
		}
	}
}
