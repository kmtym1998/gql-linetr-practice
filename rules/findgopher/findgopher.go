package findgopher

import (
	"github.com/gqlgo/gqlanalysis"
)

var Analyzer = &gqlanalysis.Analyzer{
	Name: "findgopher",
	Doc:  `find a query which name begin "Gopher"`,
	Run:  run,
}

func run(pass *gqlanalysis.Pass) (interface{}, error) {
	for t, def := range pass.Schema.Types {
		if t == "Gopher" {
			pass.Reportf(def.Position, "NG")
		}
	}

	return nil, nil
}
