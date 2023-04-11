package leetcode_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestFoobar(t *testing.T) {
	ns := []int{0, 1, 2, 20}

	q := leetcode.Q1115{}
	for _, n := range ns {
		require.Equal(t, q.FooBar(n), strings.Repeat("foobar", n))
	}
}
