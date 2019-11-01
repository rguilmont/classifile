package main

import (
	"os"
	"path"
	"testing"
	"time"

	"github.com/otiai10/copy"
)

// This is the main test : It will test the whole pipeline and then
//  check if the result was the expected one.
func TestClassification(t *testing.T) {

	copy.Copy("test_assets/file-sample_100kB.doc", "test_assets/file_to_move.doc")

	sampleDir, resDir1, resDir2 := "./test_assets", "./result_test_assets1/", "./result_test_assets2/"

	for _, d := range []string{resDir1, resDir2} {
		err := os.RemoveAll(d)
		if err != nil {
			t.Logf("Can't remove directory %v : %v. Skipping.", d, err)
		}
	}

	run(Args{"./test_assets/test.yaml", false, false, true})
	time.Sleep(5 * time.Second)

	// Check presence of files given the rules
	for _, f := range []string{
		path.Join(resDir1, "file-sample_100kB.doc"),
		path.Join(resDir1, "file-sample_100kB.odt"),
		path.Join(sampleDir, "file-sample_100kB.doc"),
		path.Join(sampleDir, "file-sample_100kB.odt"),
		path.Join(sampleDir, "other_file", "file-sample_100kB.docx"),
		path.Join(sampleDir, "file-sample_150kB.pdf"),
		path.Join(resDir2, "file-sample_100kB.docx"),
	} {
		_, err := os.Stat(f)
		if err != nil {
			t.Errorf("Expecting presence of file %v, got %v", f, err)
		}
	}

	// Check absence of files given the rules
	for _, f := range []string{
		path.Join(resDir1, "file-sample_100kB.pdf"),
		path.Join(resDir1, "file-sample_100kB.docx"),
		path.Join(resDir2, "file-sample_100kB.pdf"),
		path.Join(resDir2, "file-sample_100kB.doc"),
		path.Join(resDir2, "file-sample_100kB.odt"),
		path.Join(resDir2, "invalid.exe"),
		path.Join(resDir1, "invalid.exe"),
		path.Join(sampleDir, "file_to_move.doc"),
	} {
		res, err := os.Stat(f)
		if err == nil {
			t.Errorf("Expecting absence of file %v, got %v", f, res)
		}
	}

}
