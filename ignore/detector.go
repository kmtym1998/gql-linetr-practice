package ignore

import (
	"github.com/gqlgo/gqlanalysis"
	"github.com/vektah/gqlparser/v2/ast"
)

type Detector struct {
	Comments   []*gqlanalysis.Comment
	LinterName string
}

func (d *Detector) IsIgnoredDef(def *ast.Definition) bool {
	if def.BuiltIn {
		return true
	}

	for _, comment := range d.Comments {
		if def.Position.Line == comment.Pos.Line+1 {
			return true
		}
	}

	return false
}
