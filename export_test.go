package aliasimport

import "testing"

func ExportSetRuleYAML(t *testing.T, yamlPath string) {
	tmp := ruleYAML
	ruleYAML = yamlPath
	t.Cleanup(func() {
		ruleYAML = tmp
	})
}

func ExportParseRules(ruleFile string) error {
	_, err := parseRules(ruleFile)
	return err
}
