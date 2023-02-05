package debug

import (
	"fmt"
	"kmtym1998/gql-linter/ignore"
	"os"
	"strings"

	"github.com/gqlgo/gqlanalysis"
)

func NewAnalyzer(f *os.File) *gqlanalysis.Analyzer {
	return &gqlanalysis.Analyzer{
		Name: "debug",
		Doc:  `デバッグ用`,
		Run:  run(f),
	}
}

func run(f *os.File) func(pass *gqlanalysis.Pass) (interface{}, error) {
	return func(pass *gqlanalysis.Pass) (interface{}, error) {
		ignoreDetector := ignore.Detector{
			Comments: pass.Comments,
			RuleName: pass.Analyzer.Name,
		}

		fmt.Fprintln(f, "= comments ============================")
		for _, c := range pass.Comments {
			fmt.Fprintf(f, "%#v\n", c)
		}

		fmt.Fprintln(f, "= types ============================")
		for t, def := range pass.Schema.Types {
			fmt.Fprintf(f, "%#v\n", t)
			if ignoreDetector.IsIgnoredDef(def) {
				continue
			}

			if strings.Contains(t, "Gopher") {
				pass.Reportf(def.Position, "\"Gopher\" detected")
			}
		}

		return nil, nil
	}
}
