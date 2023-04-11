package leetcode_test

// func TestFind132Pattern(t *testing.T) {
// 	type testCase struct {
// 		nums []int
// 		ans  bool
// 	}

// 	tcs := []testCase{
// 		{
// 			nums: []int{1, 2, 3, 4},
// 			ans:  false,
// 		},
// 		{
// 			nums: []int{3, 1, 4, 2},
// 			ans:  true,
// 		},
// 		{
// 			nums: []int{-1, 3, 2, 0},
// 			ans:  true,
// 		},
// 		{
// 			nums: []int{-1, 3, 2, 0},
// 			ans:  true,
// 		},
// 		{
// 			nums: []int{1, 0, 1, -4, -3},
// 			ans:  false,
// 		},
// 		{
// 			nums: []int{1, 3, 2, 4, 5, 6, 7, 8, 9, 10},
// 			ans:  true,
// 		},
// 		{
// 			nums: []int{3, 5, 0, 3, 4},
// 			ans:  true,
// 		},
// 		{
// 			nums: []int{0, 1, 2, 0, 1, 2},
// 			ans:  true,
// 		},
// 		{
// 			nums: []int{0, 1, 2, 0, 1, 2, 0},
// 			ans:  true,
// 		},
// 	}

// 	for i := range tcs {
// 		tc := tcs[i]
// 		t.Run(fmt.Sprintf("Case %d", i+1),
// 			func(t *testing.T) {
// 				require.Equal(
// 					t, tc.ans,
// 					leetcode.Find132pattern(tc.nums))
// 			},
// 		)
// 	}
// }
