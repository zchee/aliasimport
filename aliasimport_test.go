package aliasimport_test

import (
	"path/filepath"
	"testing"

	"github.com/gostaticanalysis/testutil"
	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/iwata/aliasimport"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	aliasimport.ExportSetRuleYAML(t, filepath.Join(analysistest.TestData(), "rules.yml"))
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)
	analysistest.Run(t, testdata, aliasimport.Analyzer, "a")
}

func TestParseRules_Error(t *testing.T) {
	tests := map[string]struct {
		ruleFile string
	}{
		"register duplicated package in aliases":                    {"duplicated_aliases.yml"},
		"register duplicated package in noaliases":                  {"duplicated_noaliases.yml"},
		"register duplicated package between aliases and noaliases": {"duplicated_aliases_and_noaliases.yml"},
	}
	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			err := aliasimport.ExportParseRules(filepath.Join(analysistest.TestData(), tt.ruleFile))
			if err == nil {
				t.Error("it should return an error, but got nil")
			}
			t.Log(err)
		})
	}
}
