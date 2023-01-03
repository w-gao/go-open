package goopen

import (
	"bufio"
	"os"
	"path"
	"testing"
)

// TODO: If the tests get more complex, consider pulling in a testing library.

func TestSmartOpenFile(t *testing.T) {
	tempDir := t.TempDir()
	file := path.Join(tempDir, "test.txt")

	data := "line1\nline2\n"
	os.WriteFile(file, []byte(data), 0644)

	f, err := SmartOpen(file)
	if err != nil {
		t.Errorf("%v\n", err)
		return
	}

	defer f.Close()
	sc := bufio.NewScanner(f)

	expected := []string{"line1", "line2"}
	actual := []string{}
	for sc.Scan() {
		actual = append(actual, sc.Text())
	}

	if len(actual) != len(expected) {
		t.Errorf("expected: %v, got: %v", expected, actual)
		return
	}

	for i, val := range expected {
		if actual[i] != val {
			t.Errorf("expected: %v, got: %v", expected, actual)
			break
		}
	}
}

func TestSmartOpenHTTP(t *testing.T) {
	url := "https://gist.githubusercontent.com/w-gao/821cb75216fbf4e8841a1791abd0e6c2/raw/a5f09cb5149c6a66a873986da9ae6b92005dc186/hexdump.py"

	f, err := SmartOpen(url)
	if err != nil {
		t.Errorf("%v\n", err)
		return
	}

	defer f.Close()

	// Just read the first 100 bytes
	N := 100
	bytes := make([]byte, N)
	n, err := f.Read(bytes)

	if n != N {
		t.Errorf("expected: %v, got: %v\n", N, n)
		return
	}

	if err != nil {
		t.Errorf("%v\n", err)
		return
	}

	expected := "# Copyright (c) 2020 w-gao\n# Created at 03-07-2020\n\n\ndef hex_dump(raw, length=16):\n    \"\"\"Simple hex"
	actual := string(bytes)

	if actual != expected {
		t.Errorf("expected: %v, got: %v\n", expected, actual)
	}
}
