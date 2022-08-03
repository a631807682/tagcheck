package tagcheck

import "honnef.co/go/tools/analysis/lint"

var Docs = lint.Markdownify(map[string]*lint.RawDocumentation{
	"TAGFMT1000": {
		Title:    `Invalid tag`,
		Since:    "2022.8",
		Severity: lint.SeverityError,
		MergeIf:  lint.MergeIfAny,
	},
})
