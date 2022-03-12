package main

import "testing"

var tests = []struct {
	name     string
	dividend float32
	divisor  float32
	expected float32
	isErr    bool
}{
	{"valid_case", 100.0, 10.0, 10.0, false},
	{"invalid_case", 100.0, 0.0, 0.0, true},
}

func TestDivisionCases(t *testing.T) {

	for _, tt := range tests {
		got, err := divide(tt.dividend, tt.divisor)

		if tt.isErr {
			if err == nil {
				t.Error("expected an error, but didn't get one")
			}
		} else {
			if err != nil {
				t.Error("didn't expect error, but got one")

			}
		}

		if got != tt.expected {
			t.Errorf("expected %f but got %f", tt.expected, got)
		}
	}
}
