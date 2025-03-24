package main

import (
	"reflect"
	"strings"
	"testing"
)

const githubOrgURL = "https://github.com/github/"

var flagTestTable = []struct {
	field string
	in    []string
	out   string
}{
	{"listenAddress", []string{}, ":http"},
	{"listenAddress", []string{"--listen", ":2020"}, ":2020"},
	{"listenAddress", []string{"-l", ":2000"}, ":2000"},

	{"githubUserURL", []string{}, "https://github.com/docwhat/"},
	{"githubUserURL", []string{"--github-user-url", "https://github.com/github"}, githubOrgURL},
	{"githubUserURL", []string{"-g", "https://github.com/tpope"}, "https://github.com/tpope/"},
	{"githubUserURL", []string{"-g", githubOrgURL}, githubOrgURL},
}

func TestParseFlags(t *testing.T) {
	for _, tt := range flagTestTable {
		config, err := parseFlags(tt.in)
		if err != nil {
			t.Errorf("args %q parseFlags() returned error %v, wanted none", tt.in, err)
		}
		configValue := reflect.ValueOf(config)
		fieldValue := reflect.Indirect(configValue).FieldByName(tt.field).String()
		if !strings.EqualFold(fieldValue, tt.out) {
			t.Errorf("args %q config.%s => %q, want %q", tt.in, tt.field, fieldValue, tt.out)
		}
	}
}
