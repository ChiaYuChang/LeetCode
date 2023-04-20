package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestFindValidSplit(t *testing.T) {
	type testCase struct {
		nums []int
		ans  int
	}

	tcs := []testCase{
		{
			nums: []int{4, 7, 8, 15, 3, 5},
			ans:  2,
		},
		{
			nums: []int{4, 7, 15, 8, 3, 5},
			ans:  -1,
		},
		{
			nums: []int{
				13879, 383651, 65843, 293087, 558091, 645329, 333049, 753197,
				857341, 859181, 35831, 116471, 888163, 957331, 515777, 142787,
				762779, 266111, 169987, 89083, 451699, 376837, 856927, 517277,
				628423, 541231, 781423, 585877, 54647, 474809, 526387, 556817,
				337817, 655643, 431447, 245279, 873739, 357817, 13879, 713681,
				871901, 765169, 511211, 418007, 497111, 12517, 379343, 358811,
				351643, 674677, 688747, 826349, 805031, 529051, 923833, 488603,
				403703, 293087, 410087, 180137, 928637, 776453, 394187, 707159,
				184013, 126079, 57149, 683653, 652541, 156671, 55817, 89083,
				110569, 21851, 782329, 610733, 867151, 559067, 247381, 412147,
				475273, 554759, 825413, 181243, 262321, 930283, 846233, 760841,
				742619, 183473, 507641, 324179, 39887, 892709, 44293, 580691,
				101347, 524983, 501577, 808217, 916339, 389297, 880699, 346877,
				40841, 802667, 335693, 115879, 520703, 18773, 684889, 132911,
				173309, 984757, 821101, 514103, 110059, 634331, 277493, 645431,
				106487, 252157, 324949, 855419, 278497, 313409, 233921, 601021,
				940001, 474809, 608633, 139361, 659713, 20123, 2857, 886129, 126923,
				871303, 556817, 729607, 649631, 16921, 624997, 58543, 4391, 598877,
				909971, 693793, 339491, 255851, 911249, 227611, 196159, 451499,
				607337, 823717, 469141, 229529, 805019, 849131, 259943, 773681,
				588641, 920743, 960419, 390883, 69697, 908321, 321721, 548003,
				345773, 496897, 580691, 86629, 3529, 92401, 12689, 192343, 753197,
				736721, 624497, 842321, 415231, 228611, 489631, 902903, 388099,
				21767, 379013, 212207, 384509, 971759, 151, 200293, 132247, 216743,
				664891, 574367, 444623, 458207, 221471, 256541, 373, 286483, 607261,
				34267, 868873, 677459, 503077, 700561, 23497, 952117, 116849,
				381167, 84761, 17099, 563587, 421987, 258109, 939193, 290201, 454151,
				257069, 134153, 742541, 854599, 833717, 444967, 965773, 55219, 655211,
				370441, 68711, 334793, 897241, 308509, 516437, 731201, 887401, 746777,
				803939, 309599, 502669, 31177, 285749, 89057, 744371, 273709, 100297,
				87107, 423461, 93229, 445321, 129719, 998831, 933319, 795007, 970549,
				338717, 74771, 502883, 108541, 872173, 840223, 124489, 19421, 407317,
				134683, 64853, 143977, 567631, 626837, 835511, 133691, 476911, 904357,
				103979, 550427, 749843, 903257, 844321, 952207, 351047, 635969, 467903,
				617873, 648107, 380713, 418751, 112331, 550189, 384733, 58733, 416579,
				228929, 543187, 154493, 106877, 503441, 777041, 229769, 809201, 332009,
				24443, 305419, 660529, 103183, 456613, 233083, 76289, 345599, 269131,
				96553, 126823, 776879, 527489, 912859, 867121, 347707, 854569, 98953,
				879661, 253741, 578309, 885893, 581521, 762479, 321169, 298409, 717397,
				197359, 111103, 714563, 351347, 239509, 752681, 429223, 559633, 558401,
				485171, 636817, 239753, 855997, 897157, 545497, 131581, 831071, 664793,
				875393, 574621, 521527, 205553, 459317, 46867, 784109, 793187, 546893,
				85661, 969763, 621241, 642779, 713309, 198197, 135329,
			},
			ans: 178,
		},
		{
			nums: []int{1000000, 1000000, 1000000, 1000000, 1000000, 1000000, 1000000, 1000000, 1000000, 1000000, 1000000, 1000000, 1000000},
			ans:  -1,
		},
		{
			nums: []int{1, 1, 89},
			ans:  0,
		},
		{
			nums: []int{
				557663, 280817, 472963, 156253, 273349, 884803, 756869, 763183, 557663, 964357,
				821411, 887849, 891133, 453379, 464279, 574373, 852749, 15031, 156253, 360169,
				526159, 410203, 6101, 954851, 860599, 802573, 971693, 279173, 134243, 187367,
				896953, 665011, 277747, 439441, 225683, 555143, 496303, 290317, 652033, 713311,
				230107, 770047, 308323, 319607, 772907, 627217, 119311, 922463, 119311, 641131,
				922463, 404773, 728417, 948281, 612373, 857707, 990589, 12739, 9127, 857963,
				53113, 956003, 363397, 768613, 47981, 466027, 981569, 41597, 87149, 55021,
				600883, 111953, 119083, 471871, 125641, 922463, 562577, 269069, 806999, 981073,
				857707, 831587, 149351, 996461, 432457, 10903, 921091, 119083, 72671, 843289,
				567323, 317003, 802129, 612373,
			},
			ans: 18,
		},
		{
			nums: []int{3, 15, 5},
			ans:  -1,
		},
		{
			nums: []int{
				773494, 773554, 773618, 773678, 773702, 773774, 773782, 773842, 773854, 773926,
				773954, 773974, 773978, 773986, 774014, 774034, 774062, 774094, 774142, 774154,
				774166, 774178, 774218, 774274, 774302, 774322, 774338, 774346, 774374,
			},
			ans: -1,
		},
	}

	for i, tc := range tcs {
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				require.Equal(t, tc.ans, leetcode.FindValidSplit(tc.nums))
			},
		)
	}
}