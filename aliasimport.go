package aliasimport

import (
	"fmt"
	"go/ast"
	"io/ioutil"

	"github.com/goccy/go-yaml"
	"github.com/gostaticanalysis/analysisutil"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var ruleYAML string

//nolint:gochecknoinits // To set flag options
func init() {
	Analyzer.Flags.StringVar(&ruleYAML, "rule", "", "a file path for alias mapping rules")
}

const doc = "aliasimport can define alias name rules about import statement"

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "aliasimport",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

type rule struct {
	Aliases   map[string]string
	NoAliases map[string]struct{}
}

func parseRules(ruleFile string) (*rule, error) {
	yml, err := ioutil.ReadFile(ruleFile)
	if err != nil {
		return nil, fmt.Errorf("failed to open a rule YAML file: %w", err)
	}

	var r struct {
		Aliases   map[string]string `yaml:"aliases"`
		NoAliases []string          `yaml:"noaliases"`
	}
	if err := yaml.Unmarshal(yml, &r); err != nil {
		return nil, fmt.Errorf("failed to unmarshal a rule YAML file: %w", err)
	}

	res := &rule{
		Aliases:   map[string]string{},
		NoAliases: map[string]struct{}{},
	}

	// inverse mapping
	for alias, pkgPath := range r.Aliases {
		// double quoted
		p := fmt.Sprintf(`"%s"`, pkgPath)
		_, exist := res.Aliases[p]
		if exist {
			return nil, fmt.Errorf("duplicated aliases rule about %s", p)
		}
		res.Aliases[p] = alias
	}

	for _, pkgPath := range r.NoAliases {
		// double quoted
		p := fmt.Sprintf(`"%s"`, pkgPath)
		_, e1 := res.Aliases[p]
		if e1 {
			return nil, fmt.Errorf("conflict rules between aliases and noalises about %s", p)
		}
		_, e2 := res.NoAliases[p]
		if e2 {
			return nil, fmt.Errorf("duplicated noaliases rule about %s", p)
		}
		res.NoAliases[p] = struct{}{}
	}

	return res, nil
}

func run(pass *analysis.Pass) (interface{}, error) {
	rules, err := parseRules(ruleYAML)
	if err != nil {
		return nil, err
	}

	ins := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.ImportSpec)(nil),
	}

	// support false positive comments
	pass.Report = analysisutil.ReportWithoutIgnore(pass)

	ins.Preorder(nodeFilter, func(n ast.Node) {
		is, ok := n.(*ast.ImportSpec)
		if !ok {
			return
		}
		p := is.Path.Value
		validAlias, shouldAlias := rules.Aliases[p]
		_, shouldNoAlias := rules.NoAliases[p]

		// not matched any rules
		if !shouldAlias && !shouldNoAlias {
			return
		}

		// no aliases
		if is.Name == nil {
			if shouldAlias {
				pass.Reportf(is.Pos(), "the package %s should be imported with the alias name %s", p, validAlias)
			}
			return // shouldNoAlias
		}

		if shouldNoAlias {
			pass.Reportf(is.Pos(), "the package %s shouldn't be imported with any aliases, but with %s", p, is.Name)
			return
		}

		a := is.Name.Name
		if a != validAlias {
			pass.Reportf(is.Pos(), "the alias name of %s should be %s, not %s", p, validAlias, a)
		}
	})

	return nil, nil
}
