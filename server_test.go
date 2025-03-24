package main

import "testing"
import "context"

func TestURLExists(t *testing.T) {
	urlTestTable := []struct {
		in  string
		out bool
	}{
		{"https://google.com/", true},
		{"https://doesnotexist.example.com/", false},
	}

	for _, tt := range urlTestTable {
		exists := urlExists(context.Background(), tt.in)
		if exists != tt.out {
			t.Errorf("urlExists(%v) => %v, want %v", tt.in, exists, tt.out)
		}
	}
}
