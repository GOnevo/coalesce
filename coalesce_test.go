package coalesce_test

import (
	"fmt"
	"github.com/gonevo/coalesce"
	"testing"
)

func TestCoalesce(t *testing.T) {
	type Case struct {
		arguments []interface{}
		expected  interface{}
	}
	cases := []Case{
		{arguments: []interface{}{}, expected: nil},
		{arguments: []interface{}{nil}, expected: nil},
		{arguments: []interface{}{true}, expected: true},
		{arguments: []interface{}{nil, false, true}, expected: false},
		{arguments: []interface{}{nil, nil, true}, expected: true},
	}

	for _, c := range cases {
		t.Run(fmt.Sprint(c.expected), func(t *testing.T) {
			if coalesce.Coalesce(c.arguments...) != c.expected {
				t.Error(c)
			}
		})
	}
}
