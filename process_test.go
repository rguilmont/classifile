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

type testExecutor1 struct {
	t *testing.T
}

func (testExecutor1) copy(src string, dest string) error {
	return nil
}
func (e testExecutor1) move(src string, dest string) error {
	e.t.Error("Was not expecting move")
	// This won't be called anyway.
	return nil
}

type testExecutor2 struct {
	t *testing.T
}

func (e testExecutor2) copy(src string, dest string) error {
	e.t.Error("Was not expecting copy")
	// This won't be called anyway.
	return nil
}
func (testExecutor2) move(src string, dest string) error {
	return nil
}
func TestProcessFileMove(t *testing.T) {
	meta := make(map[string]string)

	// This should be moved, not copied. Will fail if copied.
	f := AnalysedFile{
		"content",
		meta,
		"application/pdf",
		"/src/path/to/file",
		[]Rule{
			Rule{
				[]Match{
					Match{
						"content",
						"boeuf",
						true,
					},
				},
				[]Action{
					Action{
						"copy",
						"/dst/path/to/file",
					},
				},
			},
			Rule{
				[]Match{
					Match{
						"content",
						"content",
						true,
					},
				},
				[]Action{
					Action{
						"move",
						"/dst/path/to/file",
					},
				},
			},
		},
	}

	process := make(chan AnalysedFile)
	defer close(process)
	go processFile(testExecutor2{t}, process)
	process <- f
}

func TestProcessFileCopy(t *testing.T) {
	meta := make(map[string]string)

	// This should be copied, not moved. Will fail if moved.
	f := AnalysedFile{
		"content",
		meta,
		"application/pdf",
		"/src/path/to/file",
		[]Rule{
			Rule{
				[]Match{
					Match{
						"content",
						"boeuf",
						true,
					},
				},
				[]Action{
					Action{
						"move",
						"/dst/path/to/file",
					},
				},
			},
			Rule{
				[]Match{
					Match{
						"content",
						"content",
						true,
					},
				},
				[]Action{
					Action{
						"copy",
						"/dst/path/to/file",
					},
				},
			},
		},
	}

	process := make(chan AnalysedFile)
	defer close(process)
	go processFile(testExecutor1{t}, process)
	process <- f
}
