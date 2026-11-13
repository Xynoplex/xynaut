package main_test

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// The program's files and every file of its libraries, at any depth: a
// message can be printed from any of them.
func TestNoMessageNamesATopLevelSecretVerb(t *testing.T) {
	files, err := filepath.Glob("*.go")
	if err != nil {
		t.Fatal(err)
	}
	err = filepath.WalkDir(libraries, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !entry.IsDir() && strings.HasSuffix(path, ".go") {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	for _, f := range files {
		if strings.HasSuffix(f, "_test.go") {
			continue
		}
		b, err := os.ReadFile(f)
		if err != nil {
			t.Fatal(err)
		}
		for i, line := range strings.Split(string(b), "\n") {
			if strings.Contains(line, "xynaut secret") {
				t.Errorf("%s:%d names `xynaut secret`: %s", f, i+1, strings.TrimSpace(line))
			}
		}
	}
}
