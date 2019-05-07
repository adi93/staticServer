package main

import (
	"reflect"
	"regexp"
	"strings"
	"testing"
)

func TestGetRoleAndPaths(t *testing.T) {
	tests := []struct {
		config      string
		expectedMap map[Role][]*regexp.Regexp
	}{
		{"MINIMAL\ntest.*", map[Role][]*regexp.Regexp{
			Role("MINIMAL"): regex("test.*"),
		}},
		{"MINIMAL\ntest.*\n\nADMIN\nhello\nworld", map[Role][]*regexp.Regexp{
			Role("MINIMAL"): regex("test.*"),
			Role("ADMIN"):   regex("hello", "world"),
		}},
	}

	for i, tt := range tests {

		reader := strings.NewReader(tt.config)
		m, e := getRoleAndPaths(reader)
		if e != nil {
			t.Fatalf("tests[%d] - got unexpected error %v", i, e)
		}

		if !reflect.DeepEqual(m, tt.expectedMap) {
			t.Fatalf("tests[%d] - expected %v, got %v", i, tt.expectedMap, m)
		}
	}

}

func regex(s ...string) []*regexp.Regexp {
	r := []*regexp.Regexp{}
	for _, v := range s {
		re, _ := regexp.Compile(v)
		r = append(r, re)
	}
	return r
}
