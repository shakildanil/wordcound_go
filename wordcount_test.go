package main

import "testing"

func TestCountLines(t *testing.T) {
	path := "./file_test.txt"
	r, _ := countLines(path)
	if r != 4 {
		t.Error("Expected 4, got", r)
	}
}

func TestCountSymbols(t *testing.T) {
	path := "./file_test.txt"
	r, _ := countSymbols(path)
	if r != 36 {
		t.Error("Expected 36, got", r)
	}
}

func TestScanWords(t *testing.T) {
	ans := [10]string{"cat", "dog", "cat", "cat", "dog", "cow", "dog", "dog", "cat", "cow"}
	path := "./file_test.txt"
	r, _ := scanWords(path)
	if len(r) != len(ans) {
		t.Error("Result should be an array with length of 10, but it has length of", len(r))
	} else {
		for i := 0; i < len(r); i++ {
			if ans[i] != r[i] {
				t.Error("Expected", ans[i], ", got", r[i])
			}
		}
	}
}
