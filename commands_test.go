package main

import (
	"testing"
)

func TestGetGameFullId(t *testing.T) {
	var input []nowPlaying
	input = []nowPlaying{
		{FullID: "hP09Ep7h"},
		{FullID: "mjqGC18w"},
		{FullID: "mjtGG77d"},
	}

	var tests = []struct {
		in  string
		out string
		err bool
	}{
		{"h", "hP09Ep7h", false},
		{"m", "Prefix 'm' matches multiple game IDs: [mjqGC18w mjtGG77d]", true},
		{"mjt", "mjtGG77d", false},
		{"j", "Unable to find game with ID prefixed with: 'j'", true},
	}

	for _, test := range tests {
		got, err := getGameFullId(input, test.in)

		if err != nil {
			if err.Error() != test.out {
				t.Errorf("Unexpected error: '%s'", err)
			}
		} else {
			if got != test.out {
				t.Errorf("Expected '%s' but got '%s", test.out, got)
			}
		}
	}
}
