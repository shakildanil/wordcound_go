package main

import "testing"

type testdata struct {
	words         []string
	words_count   int
	lines_count   int
	symbols_count int
}

var tests = []testdata{
	{[]string{"cat"}, 1, 1, 3},
	{[]string{"cat", "cat", "cat", "cat"}, 4, 4, 12},
	{[]string{"cat", "dog", "cow", "bum"}, 4, 4, 12},
	{[]string{"cat", "cat", "dog", "cow", "dog", "cat"}, 6, 6, 18},
	{[]string{}, 0, 0, 0},
	{[]string{"cat dog", "cat cat", "dog cow", "dog dog cat cow"}, 10, 4, 36},
}

func TestCountLines(t *testing.T) {
	for _, test := range tests {
		if countLines(test.words) != test.lines_count {
			t.Error("For", test.words, "expected", test.lines_count, "got", countLines(test.words))
		}
	}
}

func TestCountSymbols(t *testing.T) {
	for _, test := range tests {
		if countSymbols(test.words) != test.symbols_count {
			t.Error("For", test.words, "expected", test.symbols_count, "got", countSymbols(test.words))
		}
	}
}

func TestCountWords(t *testing.T) {
	for _, test := range tests {
		if countWords(test.words) != test.words_count {
			t.Error("For", test.words, "expected", test.words_count, "got", countWords(test.words))
		}
	}
}
