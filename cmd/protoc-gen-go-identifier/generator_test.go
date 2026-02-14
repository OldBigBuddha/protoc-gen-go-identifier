package main

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProtocGenGoIdentifierGoldenTest(t *testing.T) {
	if _, err := exec.LookPath("buf"); err != nil {
		t.Skip("buf not found in PATH, skipping golden test")
	}
	if _, err := exec.LookPath("protoc-gen-go"); err != nil {
		t.Skip("protoc-gen-go not found in PATH, skipping golden test")
	}

	repoRoot := findRepoRoot(t)

	// Run buf generate on testdata
	cmd := exec.Command("buf", "generate", "--path", "cmd/protoc-gen-go-identifier/testdata/test.proto")
	cmd.Dir = repoRoot
	output, err := cmd.CombinedOutput()
	require.NoError(t, err, "buf generate failed: %s", output)

	// Read the generated file
	generatedPath := filepath.Join(repoRoot, "cmd/protoc-gen-go-identifier/testdata/test_identifier.pb.go")
	generated, err := os.ReadFile(generatedPath)
	require.NoError(t, err, "failed to read generated file")

	// Read the golden file
	goldenPath := filepath.Join(repoRoot, "cmd/protoc-gen-go-identifier/testdata/test_identifier.pb.go.golden")
	golden, err := os.ReadFile(goldenPath)
	require.NoError(t, err, "failed to read golden file")

	// Compare
	assert.True(t, bytes.Equal(generated, golden),
		"generated output does not match golden file.\n"+
			"To update golden file, run:\n"+
			"  cp %s %s\n"+
			"\nDiff:\n%s",
		generatedPath, goldenPath, diff(golden, generated))
}

// findRepoRoot finds the repository root by looking for go.mod.
func findRepoRoot(t *testing.T) string {
	t.Helper()

	dir, err := os.Getwd()
	require.NoError(t, err, "failed to get working directory")

	for i := 0; i < 10; i++ {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}

	t.Fatalf("could not find repository root (go.mod)")
	return ""
}

// diff returns a simple line-by-line diff between two byte slices.
func diff(expected, actual []byte) string {
	expectedLines := bytes.Split(expected, []byte("\n"))
	actualLines := bytes.Split(actual, []byte("\n"))

	var buf bytes.Buffer
	maxLines := len(expectedLines)
	if len(actualLines) > maxLines {
		maxLines = len(actualLines)
	}

	for i := 0; i < maxLines; i++ {
		var exp, act []byte
		if i < len(expectedLines) {
			exp = expectedLines[i]
		}
		if i < len(actualLines) {
			act = actualLines[i]
		}

		if !bytes.Equal(exp, act) {
			buf.WriteString("--- expected line ")
			buf.WriteString(string(rune('0' + i/100)))
			buf.WriteString(string(rune('0' + (i/10)%10)))
			buf.WriteString(string(rune('0' + i%10)))
			buf.WriteString(": ")
			buf.Write(exp)
			buf.WriteString("\n+++ actual line ")
			buf.WriteString(string(rune('0' + i/100)))
			buf.WriteString(string(rune('0' + (i/10)%10)))
			buf.WriteString(string(rune('0' + i%10)))
			buf.WriteString(": ")
			buf.Write(act)
			buf.WriteString("\n")
		}
	}

	return buf.String()
}

