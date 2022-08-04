package tagcheck

import (
	"testing"

	"honnef.co/go/tools/analysis/lint/testutil"
)

func TestAll(t *testing.T) {
	checks := map[string][]testutil.Test{
		"STC1001": {{Dir: "gormtag"}},
	}

	testutil.Run(t, Analyzers, checks)
}
