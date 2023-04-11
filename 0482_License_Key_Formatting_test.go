package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestLicenseKeyFormatting(t *testing.T) {
	type testCase struct {
		s string
		k int
		a string
	}

	tcs := []testCase{
		{
			s: "---?!!---",
			k: 3,
			a: "",
		},
		{
			s: "5F3Z-2e-9-w",
			k: 4,
			a: "5F3Z-2E9W",
		},
		{
			s: "2-5g-3-J",
			k: 2,
			a: "2-5G-3J",
		},
		{
			s: "zDa-1b4E0r-eA8eb3eR22------------",
			k: 1,
			a: "Z-D-A-1-B-4-E-0-R-E-A-8-E-B-3-E-R-2-2",
		},
		{
			s: "-----z---Da-1b4E0r-eA8eb3eR22",
			k: 3,
			a: "Z-DA1-B4E-0RE-A8E-B3E-R22",
		},
	}

	q := leetcode.Q0482{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d by Builder", i+1),
			func(t *testing.T) {
				t.Parallel()
				require.Equal(t, tc.a, q.LicenseKeyFormattingStringBuilder(tc.s, tc.k))
			},
		)
		t.Run(
			fmt.Sprintf("Case %d by Segments", i+1),
			func(t *testing.T) {
				t.Parallel()
				require.Equal(t, tc.a, q.LicenseKeyFormattingSegments(tc.s, tc.k))
			},
		)
	}
}
