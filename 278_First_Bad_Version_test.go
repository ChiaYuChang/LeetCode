package leetcode_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

type q278API int

func (api q278API) isBadVersion(n int) bool {
	if int(api) > n {
		return false
	}
	return true
}
func TestFirstBadVersion(t *testing.T) {
	type testCase struct {
		name          string
		length        int
		lastNodBadVer q278API
	}

	tcs := make([]testCase, 10)
	for i := range tcs {
		lbv := rand.Intn(10000)
		length := rand.Intn(100000)
		if lbv > length {
			length, lbv = lbv, length
		}
		tcs[i].name = fmt.Sprintf("Case %2d", i+1)
		tcs[i].lastNodBadVer = q278API(lbv)
		tcs[i].length = length
	}

	for i := range tcs {
		tc := tcs[i]
		t.Run(
			tc.name,
			func(t *testing.T) {
				vc := leetcode.VersionCheckerAPI{
					IsBadVersion: func(n int) bool {
						return tc.lastNodBadVer.isBadVersion(n)
					},
				}
				require.Equal(
					t,
					int(tc.lastNodBadVer),
					vc.FirstBadVersion(tc.length),
				)
			},
		)
	}
}
