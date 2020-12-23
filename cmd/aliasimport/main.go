package main

import (
	"golang.org/x/tools/go/analysis/unitchecker"

	"github.com/iwata/aliasimport"
)

func main() { unitchecker.Main(aliasimport.Analyzer) }
