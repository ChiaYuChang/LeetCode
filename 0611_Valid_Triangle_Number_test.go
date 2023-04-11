package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestSearchLeftMost(t *testing.T) {
	type testCase struct {
		nums   []int
		target int
		ans    int
	}

	tcs := []testCase{
		{
			nums:   []int{2, 3, 4, 7},
			target: 5,
			ans:    7,
		},
		{
			nums:   []int{2, 3, 4, 7},
			target: 3,
			ans:    3,
		},
		{
			nums:   []int{2, 3, 4, 7},
			target: 10,
			ans:    -1,
		},
		{
			nums:   []int{5},
			target: 5,
			ans:    5,
		},
		{
			nums:   []int{5},
			target: 6,
			ans:    -1,
		},
	}

	q := leetcode.Q0611{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				i := q.SearchLeftMost(tc.nums, tc.target)
				if tc.ans < 0 {
					require.Equal(t, tc.ans, i)
				} else {
					require.Equal(t, tc.ans, tc.nums[i])
				}
			},
		)
	}
}

func TestSearchRightMost(t *testing.T) {
	type testCase struct {
		nums   []int
		target int
		ans    int
	}

	tcs := []testCase{
		{
			nums:   []int{2, 3, 4, 7},
			target: 5,
			ans:    4,
		},
		{
			nums:   []int{2, 3, 4, 7},
			target: 3,
			ans:    3,
		},
		{
			nums:   []int{2, 3, 4, 7},
			target: 10,
			ans:    7,
		},
		{
			nums:   []int{5},
			target: 5,
			ans:    5,
		},
		{
			nums:   []int{5},
			target: 4,
			ans:    -1,
		},
	}

	q := leetcode.Q0611{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				i := q.SearchRightMost(tc.nums, tc.target)
				if tc.ans < 0 {
					require.Equal(t, tc.ans, i)
				} else {
					require.Equal(t, tc.ans, tc.nums[i])
				}
			},
		)
	}
}

func TestClosedIntervalSearch(t *testing.T) {
	type testCase struct {
		name     string
		nums     []int
		min, max int
		from, to int
		ok       bool
	}

	tcs := []testCase{
		{
			name: "Includes lb and ub 1",
			nums: []int{1, 2, 3, 4, 5, 6},
			min:  2,
			max:  5,
			from: 1,
			to:   4,
			ok:   true,
		},
		{
			name: "Includes lb and ub 2",
			nums: []int{1, 3, 5, 7, 9, 11},
			min:  3,
			max:  9,
			from: 1,
			to:   4,
			ok:   true,
		},
		{
			name: "only includes ub",
			nums: []int{1, 3, 5, 7, 9, 11},
			min:  2,
			max:  9,
			from: 1,
			to:   4,
			ok:   true,
		},
		{
			name: "only includes lb",
			nums: []int{1, 3, 5, 7, 9, 11},
			min:  3,
			max:  10,
			from: 1,
			to:   4,
			ok:   true,
		},
		{
			name: "neithor includes ub nor lb",
			nums: []int{1, 3, 5, 7, 9, 11},
			min:  2,
			max:  10,
			from: 1,
			to:   4,
			ok:   true,
		},
		{
			name: "out of range (left)",
			nums: []int{3, 5},
			min:  7,
			max:  10,
			ok:   false,
		},
		{
			name: "out of range (right)",
			nums: []int{3, 5},
			min:  1,
			max:  2,
			ok:   false,
		},
	}

	q := leetcode.Q0611{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d-%s", i+1, tc.name),
			func(t *testing.T) {
				from, to, ok := q.ClosedIntervalSearch(tc.nums, tc.min, tc.max)
				require.Equal(t, tc.ok, ok)
				if ok {
					require.Equal(t, tc.to, to)
					require.Equal(t, tc.from, from)
				}
			},
		)
	}
}

func TestTriangleNumber(t *testing.T) {
	type testCase struct {
		nums []int
		ans  int
	}

	tcs := []testCase{
		{
			nums: []int{2, 2, 3, 4},
			ans:  3,
		},
		{
			nums: []int{4, 2, 3, 4},
			ans:  4,
		},
		{
			nums: []int{2, 2, 3, 4, 2, 5, 7, 7, 19, 23, 14, 22, 12, 17, 13},
			ans:  116,
		},
		{
			nums: []int{0, 0, 0},
			ans:  0,
		},
		{
			nums: []int{
				16, 70, 16, 36, 17, 39, 44, 72, 25, 88, 18, 22, 20,
				84, 18, 66, 71, 74, 87, 59, 48, 91, 52, 15, 92, 29,
				85, 21, 10, 34, 86, 68, 10, 27, 3, 42, 49, 48, 18,
				15, 6, 39, 61, 28, 19, 52, 60, 8, 37, 32, 71, 82, 54,
				38, 47, 33, 10, 64, 52, 71, 39, 63, 64, 79, 86, 47, 16, 7},
			ans: 24109,
		},
		{
			nums: []int{
				149, 383, 34, 859, 983, 44, 173, 612, 155, 730, 924, 393, 836,
				383, 660, 393, 453, 141, 113, 240, 353, 171, 237, 568, 145, 787,
				441, 350, 819, 642, 551, 862, 131, 710, 756, 375, 552, 621, 194,
				280, 140, 257, 706, 141, 490, 55, 913, 317, 447, 257, 600, 436,
				48, 592, 38, 991, 581, 906, 208, 620, 390, 341, 972, 122, 495, 22,
				521, 944, 729, 115, 227, 777, 362, 519, 973, 263, 567, 704, 910, 458,
				883, 90, 425, 448, 974, 218, 838, 842, 576, 255, 856, 924, 942,
				177, 335, 858, 534, 148, 597, 206, 428, 135, 530, 75, 351, 704,
				332, 909, 483, 268, 683, 281, 774, 953, 291, 441, 491, 70, 688, 555,
				585, 175, 156, 274, 242, 859, 604, 624, 819, 341, 738, 38, 876, 538,
				928, 691, 942, 707, 338, 934, 583, 316, 60, 317, 955, 901, 515, 554,
				205, 800, 577, 477, 761, 165, 701, 766, 657, 636, 923, 909, 481, 88,
				694, 357, 162, 104, 138, 755, 84, 718, 753, 659, 976, 464, 988, 323,
				304, 924, 818, 892, 32, 940, 557, 114, 989, 254, 629, 471, 503, 261,
				650, 261, 437, 38, 911, 970, 491, 186, 770, 767, 349, 598, 277, 932,
				720, 645, 411, 369, 322, 408, 160, 945, 929, 543, 42, 255, 649, 971,
				329, 125, 803, 6, 796, 135, 567, 38, 55, 561, 937, 764, 528, 924, 710,
				986, 333, 366, 827, 867, 712, 94, 832, 599, 582, 341, 550, 657, 949,
				1000, 112, 302, 520, 566, 336, 673, 918, 891, 558, 17, 314, 537, 154,
				998, 935, 830, 391, 465, 11, 217, 658, 533, 633, 543, 168, 400, 66,
				131, 963, 63, 514, 922, 90, 582, 107, 664, 585, 962, 922, 585, 791,
				67, 20, 551, 974, 382, 604, 177, 293, 145, 563, 713, 139, 644, 478,
				423, 124, 714, 614, 473, 577, 79, 722, 472, 559, 276, 274, 285, 563,
				965, 632, 574, 370, 914, 720, 431, 789, 756, 855, 600, 50, 816, 260,
				329, 341, 980, 595, 154, 193, 386, 299, 704, 26, 403, 130, 134, 658,
				534, 724, 346, 719, 224, 915, 134, 212, 40, 709, 745, 891, 30, 923,
				823, 439, 592, 973, 803, 330, 493, 39, 34, 565, 304, 491, 307, 851,
				58, 961, 369, 692, 796, 319, 700, 972, 789, 192, 402, 93, 444, 641,
				83, 2, 934, 25, 887, 832, 113, 63, 230, 816, 724, 524, 774, 638, 413,
				911, 350, 48, 105, 446, 608, 30, 14, 853, 242, 628, 767, 67, 783, 932,
				36, 383, 677, 111, 768, 252, 668, 672, 800, 811, 541, 41, 783, 169,
				694, 382, 336, 765, 903, 969, 726, 113, 467, 234, 934, 17, 607, 438, 684,
				460, 109, 157, 455, 172, 734, 804, 363, 967, 120, 125, 847, 535, 353, 916,
				899, 531, 51, 286, 433, 532, 783, 437, 253, 870, 906, 340, 95, 327, 892,
				897, 990, 305, 608, 25, 275, 908, 513, 848, 954, 287, 775, 691, 803, 950,
				850, 91, 350, 673, 828, 821, 791, 18, 387, 353, 754, 42, 974, 502, 548,
				894, 722, 632, 436, 180, 650, 74, 434, 654, 688, 203, 495, 212, 29, 893,
				175, 465, 735, 300, 266, 161, 526, 898, 803, 32, 475, 492, 624, 644, 500,
				240, 74, 476, 365, 57, 812, 993, 625, 362, 653, 24, 524, 495, 957, 526,
				986, 180, 477, 170, 835, 670, 630, 660, 303, 452, 561, 270, 850, 953, 301,
				981, 214, 29, 42, 799, 663, 974, 369, 217, 199, 148, 98, 155, 935, 616,
				240, 543, 158, 671, 60, 467, 135, 767, 593, 207, 709, 856, 158, 345,
				311, 510, 486, 789, 10, 204, 607, 333, 963, 182, 367, 434, 76, 270,
				554, 297, 49, 710, 945, 686, 18, 947, 313, 698, 507, 76, 73, 872, 582,
				846, 996, 717, 552, 661, 720, 158, 80, 220, 705, 437, 153, 533, 710,
				682, 253, 547, 570, 440, 885, 538, 382, 701, 695, 48, 689, 967, 21, 179,
				544, 351, 241, 995, 249, 205, 330, 769, 131, 109, 92, 654, 896, 510, 686,
				442, 611, 488, 305, 180, 908, 145, 53, 946, 755, 616, 338, 582, 724, 524,
				924, 527, 194, 364, 184, 257, 791, 945, 138, 217, 574, 585, 870, 514, 887,
				591, 629, 220, 745, 609, 692, 820, 786, 318, 502, 915, 840, 202, 358, 336,
				191, 109, 423, 857, 1, 28, 648, 201, 943, 134, 463, 106, 906, 471, 604, 358,
				117, 55, 246, 237, 980, 893, 565, 630, 288, 326, 951, 800, 884, 925, 448,
				432, 859, 896, 900, 129, 351, 976, 812, 79, 261, 756, 686, 290, 891, 829,
				693, 336, 381, 710, 80, 231, 622, 19, 218, 534, 592, 579, 299, 448, 746,
				684, 777, 53, 105, 797, 731, 991, 972, 772, 76, 358, 470, 687, 860, 19,
				229, 141, 772, 624, 833, 635, 849, 369, 829, 998, 2, 468, 190, 541, 915,
				111, 412, 924, 797, 127, 637, 827, 754, 449, 985, 253, 214, 428, 224, 445,
				780, 643, 981, 964, 194, 809, 809, 864, 839, 466, 851, 865, 75, 32, 594,
				403, 695, 575, 743, 205, 872, 305, 979, 196, 401, 275, 710, 891, 381, 578,
				499, 231, 611, 776, 403, 607, 887, 905, 335, 140, 875, 999, 233, 236, 802,
				305, 593, 474, 113, 443, 488, 615, 594, 118, 991, 470, 854, 788, 473, 825,
				442, 892, 462, 988, 34, 933, 802, 44, 94, 20, 921, 717, 980, 831, 968, 55,
				576, 142, 452, 104, 343, 363, 702, 19, 737, 992, 292, 773, 460, 172, 504,
				347, 424, 114, 769, 179, 906, 236, 375, 322, 208, 20, 301, 342, 467, 520,
				217, 676, 95, 838, 312, 550, 311, 474, 195, 12, 763, 387, 801, 142, 693, 427,
				515, 820, 435, 413, 550, 878, 41, 646, 314, 522, 783, 828, 246, 877, 642, 661,
				175, 180, 684, 222, 762, 295, 751, 454, 652, 146, 77, 435, 626, 563, 708, 440,
				797, 100, 209, 568, 631, 502, 433, 268, 500, 688, 378, 371, 329, 600, 242,
			},
			ans: 80112256,
		},
	}

	q := leetcode.Q0611{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Fast Case %d", i+1),
			func(t *testing.T) {
				require.Equal(t, tc.ans, q.TriangleNumber(tc.nums))
			},
		)
	}

	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Slow Case %d", i+1),
			func(t *testing.T) {
				require.Equal(t, tc.ans, q.SlowTriangleNumber(tc.nums))
			},
		)
	}
}