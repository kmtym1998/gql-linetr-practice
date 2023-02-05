package ignore

import (
	"strings"

	"github.com/gqlgo/gqlanalysis"
	"github.com/samber/lo"
	"github.com/vektah/gqlparser/v2/ast"
)

type Detector struct {
	Comments []*gqlanalysis.Comment
	RuleName string
}

func (d *Detector) IsIgnoredDef(def *ast.Definition) bool {
	if def.BuiltIn {
		return true
	}

	for _, c := range d.Comments {
		comment := c.Value
		if def.Position.Line != c.Pos.Line+1 {
			continue
		}

		if strings.HasPrefix(comment, "# nolint") {
			comment = strings.Replace(comment, "# nolint", "", 1)

			if strings.Contains(comment, ":") {
				comment = strings.Replace(comment, ":", "", 1)
			} else {
				return true
			}

			if lo.Contains(strings.Split(comment, ","), d.RuleName) {
				return true
			}
		}
	}

	return false
}
