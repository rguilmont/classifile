package main

import (
	"testing"
)

func TestInvalidRuleValidation(t *testing.T) {

	r := Rule{
		[]Match{
			Match{
				"test",
				"(.?)romain",
				true,
			},
		},
		[]Action{},
	}

	status := validateRule(r)
	//t.Log(err)
	if status == true {
		t.Errorf("Rule is not valid, but got : %v - %v ", status, r)
	}
}

func TestMatchingRules(t *testing.T) {

	m := make(map[string]string)
	m["Author"] = "Romain Guilmont"
	a := AnalysedFile{
		"This is a valid file, and this is the content.",
		m,
		"application/pdf",
		"/path/to/file.pdf",
		[]Rule{},
	}

	rule := Rule{
		[]Match{
			Match{
				"content",
				"valid file",
				true,
			},
		},
		[]Action{},
	}
	if !processRule(rule, a) {
		t.Errorf("Rule is not matching, but should match. Rule : %v", rule)
	}

	rule = Rule{
		[]Match{
			Match{
				"content",
				"this is not matching",
				true,
			},
		},
		[]Action{},
	}
	if processRule(rule, a) {
		t.Errorf("Rule is matching, but should not match. Rule : %v", rule)
	}

	rule = Rule{
		[]Match{
			Match{
				"metadata",
				"(?i)ro.ain",
				true,
			},
		},
		[]Action{},
	}
	if !processRule(rule, a) {
		t.Errorf("Rule is not matching, but should match. Rule : %v", rule)
	}

	rule = Rule{
		[]Match{
			Match{
				"This is not",
				"a valid rule",
				true,
			},
		},
		[]Action{},
	}
	if processRule(rule, a) {
		t.Errorf("Rule is matching, but should be invalid and not match. Rule : %v", rule)
	}
}
