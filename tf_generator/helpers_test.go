package main

import "testing"

func TestToSnakeCase(t *testing.T) {
	testSet := []struct {
		in  string
		out string
	}{
		{in: "AaaaBbbb", out: "aaaa_bbbb"},
		{in: "aaaa_bbbb", out: "aaaa_bbbb"},
		{in: "AaaaBbbbCccc", out: "aaaa_bbbb_cccc"},
		{in: "aaaaBbbb", out: "aaaa_bbbb"},
		{in: "Aaaa", out: "aaaa"},
		{in: "aaaa", out: "aaaa"},
		{in: "", out: ""},
		{in: "_", out: "_"},
		{in: ".", out: "."},
	}
	for _, s := range testSet {
		if s.out != ToSnakeCase(s.in) {
			t.Errorf("%s expected %s, got %s", s.in, s.out, ToSnakeCase(s.in))
		}
	}
}
