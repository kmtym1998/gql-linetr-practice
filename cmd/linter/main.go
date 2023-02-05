package main

import (
	"kmtym1998/gql-linter/rules/debug"
	"os"

	"github.com/gqlgo/gqlanalysis/multichecker"
)

func main() {
	f, err := os.Create(".vscode/debug.txt")
	if err != nil {
		panic(err)
	}

	multichecker.Main(
		debug.NewAnalyzer(f),
	)
}
