package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func fakeAsyncFunction(s string) string {
	time.Sleep(1 * time.Second)
	return s
}

func TestAsync(t *testing.T) {
	type testCase struct {
		input    string
		expected string
	}

	tests := []testCase{
		{input: "abc", expected: "abc"},
		{input: "abcd", expected: "abcd"},
		{input: "abcde", expected: "abcde"},
		{input: "abcdef", expected: "abcdef"},
		{input: "abcdefg", expected: "abcdefg"},
		{input: "abcdefgh", expected: "abcdefgh"},
	}

	c := make(chan string, len(tests))

	for i := range tests {
		tc := tests[i]
		go func(tc testCase) {
			t.Run(tc.input, func(t *testing.T) {
				got := fakeAsyncFunction(tc.input)
				assert.Equal(t, tc.expected, got)
				c <- ""
			})
		}(tc)
	}

	for range tests {
		<-c
	}
}
