package main

import (
	"reflect"
	"strings"
	"testing"
)

var flagTestTable = []struct {
	field string
	in    []string
	out   string
}{
	{"listenAddress", []string{}, ":http"},
	{"listenAddress", []string{"--listen", ":2020"}, ":2020"},
	{"listenAddress", []string{"-l", ":2000"}, ":2000"},

	{"githubUserURL", []string{}, "https://github.com/docwhat/"},
	{"githubUserURL", []string{"--github-user-url", "https://github.com/github"}, "https://github.com/github/"},
	{"githubUserURL", []string{"-g", "https://github.com/tpope"}, "https://github.com/tpope/"},
	{"githubUserURL", []string{"-g", "https://github.com/github/"}, "https://github.com/github/"},
}

func TestParseFlags(t *testing.T) {
	for _, tt := range flagTestTable {
		config := parseFlags(tt.in)
		configValue := reflect.ValueOf(config)
		fieldValue := reflect.Indirect(configValue).FieldByName(tt.field).String()
		if !strings.EqualFold(fieldValue, tt.out) {
			t.Errorf("args %q config.%s => %q, want %q", tt.in, tt.field, fieldValue, tt.out)
		}
	}
}
