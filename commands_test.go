package main

import "testing"

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
	}{
		{"h", "hP09Ep7h"},
		{"m", "mjqGC18w"},
		{"mjt", "mjtGG77d"},
	}

	for _, test := range tests {
		got, err := getGameFullId(input, test.in)
		if err != nil {
			t.Errorf("Test failed due to error: '%s'", err)
		}
		if got != test.out {
			t.Errorf("Expected '%s' but got '%s", test.out, got)
		}
	}
}
