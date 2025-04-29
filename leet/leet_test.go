package leet_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"jack.com/leet/leet"
)

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9

	actual := leet.TwoSum(nums, target)

	assert.Equal(t, []int{1, 0}, actual)
}

func TestTwoSum2(t *testing.T) {
	nums := []int{0, 4, 3, 0}
	target := 0

	actual := leet.TwoSum(nums, target)

	assert.Equal(t, []int{3, 0}, actual)
}

func TestAddTwoNumbers(t *testing.T) {
	l1 := leet.ListNode{Val: 5, Next: &leet.ListNode{Val: 2, Next: &leet.ListNode{Val: 4}}}
	l2 := leet.ListNode{Val: 4, Next: &leet.ListNode{Val: 0, Next: &leet.ListNode{Val: 1}}}

	expect := &leet.ListNode{Val: 9, Next: &leet.ListNode{Val: 2, Next: &leet.ListNode{Val: 5}}}

	assert.Equal(t, expect, leet.AddTwoNumbers(&l1, &l2))
}

func TestLengthOfLongestSubstring(t *testing.T) {

	s := "dvdf"
	e := 3

	a := leet.LengthOfLongestSubstring(s)

	assert.Equal(t, e, a)
}

func TestFindMedianSortedArrays(t *testing.T) {
	t.Run("Different length arrays", func(t *testing.T) {
		a := []int{3}
		b := []int{-2, -1}

		expect := float64(-1)

		assert.Equal(t, expect, leet.FindMedianSortedArrays(a, b))
	})
}

func TestLongestPalindrome(t *testing.T) {
	t.Run("standard", func(t *testing.T) {
		s := "babad"

		expect := "bab"

		assert.Equal(t, expect, leet.LongestPalindrome(s))
	})

	t.Run("standard", func(t *testing.T) {
		s := "cbbd"

		expect := "bb"

		assert.Equal(t, expect, leet.LongestPalindrome(s))
	})

	t.Run("standard", func(t *testing.T) {
		s := "a"

		expect := "a"

		assert.Equal(t, expect, leet.LongestPalindrome(s))
	})
}

func TestZigZagConvert(t *testing.T) {
	t.Run("Different length arrays", func(t *testing.T) {
		s := "AB"
		n := 1

		expect := "AB"

		assert.Equal(t, expect, leet.ZigZagConvert(s, n))
	})
}

func TestReverse(t *testing.T) {
	t.Run("Base case", func(t *testing.T) {
		in := int(-123)
		expected := int(-321)

		assert.Equal(t, expected, leet.Reverse32(in))
	})
}

func TestIsPalindrome(t *testing.T) {
	t.Run("Base case", func(t *testing.T) {
		in := int(56004065)
		out := false

		assert.Equal(t, out, leet.IsPalindrome(in))
	})
}

func TestMaxArea(t *testing.T) {
	t.Run("Base case", func(t *testing.T) {
		assert.Equal(t, 49, leet.MaxAreaPointers([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	})
}

func TestIntegerToRoman(t *testing.T) {
	t.Run("Base backward", func(t *testing.T) {
		assert.Equal(t, leet.IntegerToRoman(4), "IV")
	})

	t.Run("Base forward", func(t *testing.T) {
		assert.Equal(t, leet.IntegerToRoman(6), "VI")
	})

	t.Run("Base large", func(t *testing.T) {
		assert.Equal(t, leet.IntegerToRoman(3749), "MMMDCCXLIX")
	})

	t.Run("Base mid", func(t *testing.T) {
		assert.Equal(t, leet.IntegerToRoman(58), "LVIII")
	})

	t.Run("Base interesting", func(t *testing.T) {
		assert.Equal(t, leet.IntegerToRoman(1994), "MCMXCIV")
	})

	t.Run("Base 10", func(t *testing.T) {
		assert.Equal(t, leet.IntegerToRoman(10), "X")
	})
}

func TestRomanToInteger(t *testing.T) {
	// I want to test VI, IV, X, MCX
	t.Run("Sixes", func(t *testing.T) {
		assert.Equal(t, leet.RomanToInteger("VI"), 6)
	})

	t.Run("4", func(t *testing.T) {
		assert.Equal(t, leet.RomanToInteger("IV"), 4)
	})

	t.Run("Large", func(t *testing.T) {
		assert.Equal(t, leet.RomanToInteger("MMCXI"), 2111)
	})

	t.Run("THREEE", func(t *testing.T) {
		assert.Equal(t, leet.RomanToInteger("III"), 3)
	})

	t.Run("TEN", func(t *testing.T) {
		assert.Equal(t, leet.RomanToInteger("X"), 10)
	})

	t.Run("LV ", func(t *testing.T) {
		assert.Equal(t, leet.RomanToInteger("LVIII"), 58)
	})

	t.Run("Larger", func(t *testing.T) {
		assert.Equal(t, leet.RomanToInteger("MCMXCIV"), 1994)
	})

}

func TestLongestCommonPrefix(t *testing.T) {
	t.Run("No case", func(t *testing.T) {
		assert.Equal(t, "", leet.LongestCommonPrefix([]string{"abc", "bcd"}))
	})

	t.Run("One", func(t *testing.T) {
		assert.Equal(t, "a", leet.LongestCommonPrefix([]string{"abc", "acd"}))
	})

	t.Run("Many", func(t *testing.T) {
		assert.Equal(t, "abc", leet.LongestCommonPrefix([]string{"abc", "abc"}))
	})

	t.Run("Leet1", func(t *testing.T) {
		assert.Equal(t, "fl", leet.LongestCommonPrefix([]string{"flower", "flow", "flight"}))
	})

	t.Run("Leet2", func(t *testing.T) {
		assert.Equal(t, "", leet.LongestCommonPrefix([]string{"dog", "racecar", "car"}))
	})

	t.Run("Empty", func(t *testing.T) {
		assert.Equal(t, "", leet.LongestCommonPrefix([]string{""}))
	})
}

// func TestThreeSum(t *testing.T) {
// 	t.Run("No case", func(t *testing.T)  {
// 		out := [][]int{{1, -1, 0}}
// 		assert.Equal(t, out , leet.ThreeSum([]int{1, 1, -1 , 0}))
// 	})
// }

func TestIsValid(t *testing.T) {
	t.Run("base case", func(t *testing.T) {
		assert.Equal(t, true, leet.IsValid("()"))
	})

	t.Run("Independent brackets", func(t *testing.T) {
		assert.Equal(t, true, leet.IsValid("()[]{}"))
	})

	t.Run("Nested brackets", func(t *testing.T) {
		assert.Equal(t, true, leet.IsValid("([])"))
	})

	t.Run("Unclosed", func(t *testing.T) {
		assert.Equal(t, false, leet.IsValid("(]"))
	})

	t.Run("No opener", func(t *testing.T) {
		assert.Equal(t, false, leet.IsValid("]"))
	})

}

func TestThreeSum(t *testing.T) {
	t.Run("base case", func(t *testing.T) {
		assert.Equal(t, [][]int{{1, 2, -3}}, leet.ThreeSum([]int{1, 2, -3}))
	})

	t.Run("Zeroes", func(t *testing.T) {
		assert.Equal(t, [][]int{{0, 0, 0}}, leet.ThreeSum([]int{0, 0, 0}))
	})

	t.Run("Multiple Zeroes", func(t *testing.T) {
		assert.Equal(t, [][]int{{-1, 0, 1}, {-1, 2, -1}, {0, 1, -1}}, leet.ThreeSum([]int{-1, 0, 1, 2, -1, -4}))
	})
}

func TestLetterCombinations(t *testing.T) {
	t.Run("base case", func(t *testing.T) {
		assert.Equal(t, []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}, leet.LetterCombinations("23"))
	})

	t.Run("null case", func(t *testing.T) {
		assert.Equal(t, []string{}, leet.LetterCombinations(""))
	})

	t.Run("one letter", func(t *testing.T) {
		assert.Equal(t, []string{"a", "b", "c"}, leet.LetterCombinations("2"))
	})

	t.Run("4 numbered keys", func(t *testing.T) {
		assert.Equal(t, []string{"w", "x", "y", "z"}, leet.LetterCombinations("9"))
	})
}

func TestRemoveNthFromEnd(t *testing.T) {

	t.Run("base case", func(t *testing.T) {
		assert.Equal(t, &leet.ListNode{Val: 1}, leet.RemoveNthFromEnd(&leet.ListNode{Val: 1, Next: &leet.ListNode{Val: 2}}, 1))
	})
	t.Run("base case 2", func(t *testing.T) {
		assert.Equal(t, &leet.ListNode{Val: 2}, leet.RemoveNthFromEnd(&leet.ListNode{Val: 1, Next: &leet.ListNode{Val: 2}}, 2))
	})

	t.Run("base case long", func(t *testing.T) {
		assert.Equal(t, &leet.ListNode{Val: 1, Next: &leet.ListNode{Val: 2, Next: &leet.ListNode{Val: 3, Next: &leet.ListNode{Val: 5}}}}, leet.RemoveNthFromEnd(&leet.ListNode{Val: 1, Next: &leet.ListNode{Val: 2, Next: &leet.ListNode{Val: 3, Next: &leet.ListNode{Val: 4, Next: &leet.ListNode{Val: 5}}}}}, 2))
	})
	// t.Run("single", func (t *testing.T)  {
	// 	assert.Equal(t, nil ,leet.RemoveNthFromEnd(&leet.ListNode{Val: 1}, 1))
	// })
}

func TestGenerateParentheses(t *testing.T) {
	t.Run("n=1", func(t *testing.T) {
		assert.Equal(t, []string{"()"}, leet.GenerateParantheses(1))
	})

	t.Run("n=2", func(t *testing.T) {
		assert.Equal(t, []string{"(())", "()()"}, leet.GenerateParantheses(2))
	})

	t.Run("n=3", func(t *testing.T) {
		assert.Equal(t, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}, leet.GenerateParantheses(3))
	})
}

func TestSwapPairs(t *testing.T) {
	t.Run("base case", func(t *testing.T) {
		assert.Equal(t, &leet.ListNode{2, &leet.ListNode{Val: 1}}, leet.SwapPairs(&leet.ListNode{1, &leet.ListNode{Val: 2}}))
	})

	t.Run("base case - odd", func(t *testing.T) {
		assert.Equal(t, &leet.ListNode{2, &leet.ListNode{3, &leet.ListNode{Val: 1}}}, leet.SwapPairs(&leet.ListNode{3, &leet.ListNode{2, &leet.ListNode{Val: 1}}}))
	})

	t.Run("base case 1234", func(t *testing.T) {
		assert.Equal(t, &leet.ListNode{4, &leet.ListNode{3, &leet.ListNode{2, &leet.ListNode{Val: 1}}}}, leet.SwapPairs(&leet.ListNode{1, &leet.ListNode{2, &leet.ListNode{3, &leet.ListNode{Val: 4}}}}))
	})
}

// func TestReverseKGroup(t *testing.T) {
// 	t.Run("base case", func(t *testing.T) {
// 		assert.Equal(t, &leet.ListNode{1, &leet.ListNode{Val:2}}, leet.ReverseKGroup(&leet.ListNode{1, &leet.ListNode{Val:2}}, 1))
// 	})

// 	t.Run("base case", func(t *testing.T) {
// 		assert.Equal(t, &leet.ListNode{2, &leet.ListNode{Val:1}}, leet.ReverseKGroup(&leet.ListNode{1, &leet.ListNode{Val:2}}, 2))
// 	})

// 	t.Run("base case - odd", func(t *testing.T) {
// 		assert.Equal(t, &leet.ListNode{3, &leet.ListNode{2, &leet.ListNode{Val:1}}}, leet.ReverseKGroup( &leet.ListNode{3, &leet.ListNode{2, &leet.ListNode{Val:1}}}, 1))
// 	})

// 	t.Run("base case - odd", func(t *testing.T) {
// 		assert.Equal(t, &leet.ListNode{2, &leet.ListNode{3, &leet.ListNode{Val:1}}}, leet.ReverseKGroup( &leet.ListNode{3, &leet.ListNode{2, &leet.ListNode{Val:1}}}, 2))
// 	})

// 	t.Run("base case - odd", func(t *testing.T) {
// 		assert.Equal(t, &leet.ListNode{1, &leet.ListNode{2, &leet.ListNode{Val:3}}}, leet.ReverseKGroup( &leet.ListNode{3, &leet.ListNode{2, &leet.ListNode{Val:1}}}, 3))
// 	})

// }

func TestMergeSortedArray(t *testing.T) {
	t.Run("base case", func(t *testing.T) {
		nums1 := []int{1, 2, 3, 0, 0, 0}
		nums2 := []int{2, 5, 6}
		m, n := 3, 3
		expected := []int{1, 2, 2, 3, 5, 6}

		leet.MergeSortedArray(nums1, nums2, m, n)

		assert.Equal(t, expected, nums1)
	})

	t.Run("null case 1", func(t *testing.T) {
		nums1 := []int{1}
		nums2 := []int{}
		m, n := 1, 0
		expected := []int{1}

		leet.MergeSortedArray(nums1, nums2, m, n)

		assert.Equal(t, expected, nums1)
	})

	t.Run("null case 2", func(t *testing.T) {
		nums1 := []int{}
		nums2 := []int{}
		m, n := 0, 0
		expected := []int{}

		leet.MergeSortedArray(nums1, nums2, m, n)

		assert.Equal(t, expected, nums1)
	})

	t.Run("null case 3", func(t *testing.T) {
		nums1 := []int{0}
		nums2 := []int{1}
		m, n := 0, 1
		expected := []int{1}

		leet.MergeSortedArray(nums1, nums2, m, n)

		assert.Equal(t, expected, nums1)
	})

	t.Run("interesting", func(t *testing.T) {
		nums1 := []int{-1, 0, 0, 3, 3, 3, 0, 0, 0}
		nums2 := []int{1, 2, 2}
		m, n := 6, 3
		expected := []int{-1, 0, 0, 1, 2, 2, 3, 3, 3}

		leet.MergeSortedArray(nums1, nums2, m, n)

		assert.Equal(t, expected, nums1)
	})

}

func TestRemoveElement(t *testing.T) {
	t.Run("base case", func(t *testing.T) {
		nums := []int{3, 2, 2, 3}
		val := 3
		expected := []int{2, 2, 2, 3}
		count := leet.RemoveElement(nums, val)
		expectedCount := 2
		assert.Equal(t, expected, nums)
		assert.Equal(t, expectedCount, count)
	})

	t.Run("base case 2", func(t *testing.T) {
		nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
		val := 2
		expected := []int{0, 1, 3, 0, 4, 0, 4, 2}
		count := leet.RemoveElement(nums, val)
		expectedCount := 3
		assert.Equal(t, expected, nums)
		assert.Equal(t, expectedCount, count)
	})
	t.Run("null case", func(t *testing.T) {
		nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
		val := 5
		expected := []int{0, 1, 2, 2, 3, 0, 4, 2}
		count := leet.RemoveElement(nums, val)
		expectedCount := 0
		assert.Equal(t, expected, nums)
		assert.Equal(t, expectedCount, count)
	})
}

func TestRemoveDuplicates(t *testing.T) {
	t.Run("base case", func(t *testing.T) {
		in := []int{1, 2, 2}
		uniqueCount := 2
		expected := []int{1, 2, 2}

		actual := leet.RemoveDuplicates(in)
		assert.Equal(t, expected, in)
		assert.Equal(t, uniqueCount, actual)
	})

	t.Run("base case 2", func(t *testing.T) {
		in := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
		uniqueCount := 5
		expected := []int{0, 1, 2, 3, 4, 2, 2, 3, 3, 4}

		actual := leet.RemoveDuplicates(in)
		assert.Equal(t, expected, in)
		assert.Equal(t, uniqueCount, actual)
	})

	t.Run("null case", func(t *testing.T) {
		in := []int{0}
		uniqueCount := 1
		expected := []int{0}

		actual := leet.RemoveDuplicates(in)
		assert.Equal(t, expected, in)
		assert.Equal(t, uniqueCount, actual)
	})
}

func TestRemoveDuplicates2(t *testing.T) {
	t.Run("base case", func(t *testing.T) {
		in := []int{1, 2, 2}
		uniqueCount := 3
		expected := []int{1, 2, 2}

		actual := leet.RemoveDuplicates2(in)
		assert.Equal(t, expected, in)
		assert.Equal(t, uniqueCount, actual)
	})

	t.Run("base case 2", func(t *testing.T) {
		in := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
		uniqueCount := 9
		expected := []int{0, 0, 1, 1, 2, 2, 3, 3, 4, 4}

		actual := leet.RemoveDuplicates2(in)
		assert.Equal(t, expected, in)
		assert.Equal(t, uniqueCount, actual)
	})

	t.Run("null case", func(t *testing.T) {
		in := []int{0}
		uniqueCount := 1
		expected := []int{0}

		actual := leet.RemoveDuplicates2(in)
		assert.Equal(t, expected, in)
		assert.Equal(t, uniqueCount, actual)
	})
}

func TestMajorityElement(t *testing.T) {
	t.Run("base case", func(t *testing.T) {
		in := []int{3, 2, 3}
		expected := 3
		assert.Equal(t, expected, leet.MajorityElement(in))
	})

	t.Run("base case 2", func(t *testing.T) {
		in := []int{2, 2, 1, 1, 1, 2, 2}
		expected := 2
		assert.Equal(t, expected, leet.MajorityElement(in))
	})
}

func TestRotateArray(t *testing.T) {
	t.Run("base case", func(t *testing.T) {
		in := []int{1, 2, 3, 4, 5, 6, 7}
		k := 3
		expected := []int{5, 6, 7, 1, 2, 3, 4}
		leet.RotateArray(in, k)
		assert.Equal(t, expected, in)
	})

	t.Run("base case 2", func(t *testing.T) {
		in := []int{1, 2, 3, 4, 5, 6, 7}
		k := 4
		expected := []int{4, 5, 6, 7, 1, 2, 3}
		leet.RotateArray(in, k)
		assert.Equal(t, expected, in)
	})

	t.Run("base case k > l", func(t *testing.T) {
		in := []int{1, 2, 3, 4, 5, 6, 7}
		k := 10
		expected := []int{5, 6, 7, 1, 2, 3, 4}
		leet.RotateArray(in, k)
		assert.Equal(t, expected, in)
	})

	t.Run("base case divisble by 2", func(t *testing.T) {
		in := []int{-1, -100, 3, 99}
		k := 2
		expected := []int{3, 99, -1, -100}
		leet.RotateArray(in, k)
		assert.Equal(t, expected, in)
	})

	t.Run("failed case", func(t *testing.T) {
		in := []int{-1, 2, 3, 4, 5, 6}
		k := 2
		expected := []int{5, 6, -1, 2, 3, 4}
		leet.RotateArray(in, k)
		assert.Equal(t, expected, in)
	})

	t.Run("failed case 2", func(t *testing.T) {
		in := []int{1, 2, 3, 4, 5, 6}
		k := 4
		expected := []int{3, 4, 5, 6, 1, 2}
		leet.RotateArray(in, k)
		assert.Equal(t, expected, in)
	})
}

func TestMaxProfit(t *testing.T) {
	t.Run("base case", func(t *testing.T) {
		in := []int{7, 1, 5, 3, 6, 4}
		expected := 5
		assert.Equal(t, expected, leet.MaxProfit(in))
	})

	t.Run("base case 2", func(t *testing.T) {
		in := []int{7, 6, 4, 3, 1}
		expected := 0
		assert.Equal(t, expected, leet.MaxProfit(in))
	})
}

func TestMaxProfit2(t *testing.T) {
	t.Run("base case", func(t *testing.T) {
		in := []int{7, 1, 5, 3, 6, 4}
		expected := 7
		assert.Equal(t, expected, leet.MaxProfit2(in))
	})

	t.Run("base case 1.5", func(t *testing.T) {
		in := []int{1, 2, 3, 4, 5}
		expected := 4
		assert.Equal(t, expected, leet.MaxProfit2(in))
	})

	t.Run("base case 2", func(t *testing.T) {
		in := []int{7, 6, 4, 3, 1}
		expected := 0
		assert.Equal(t, expected, leet.MaxProfit2(in))
	})
}

func TestCanJump(t *testing.T) {
	t.Run("base case", func(t *testing.T) {
		in := []int{2, 3, 1, 1, 4}
		expected := true

		assert.Equal(t, expected, leet.CanJump(in))
	})

	t.Run("base case", func(t *testing.T) {
		in := []int{3, 2, 1, 0, 4}
		expected := false

		assert.Equal(t, expected, leet.CanJump(in))
	})

	t.Run("failed case", func(t *testing.T) {
		in := []int{2, 5, 0, 0}
		expected := true

		assert.Equal(t, expected, leet.CanJump(in))
	})
}

func TestCanConstruct(t *testing.T) {
	t.Run("base case", func(t *testing.T) {
		inMag := "abb"
		inNote := "ab"
		expected := true

		assert.Equal(t, expected, leet.CanConstruct(inNote, inMag))
	})

	t.Run("base case - fail", func(t *testing.T) {
		inMag := "aa"
		inNote := "ab"
		expected := false

		assert.Equal(t, expected, leet.CanConstruct(inNote, inMag))
	})

	t.Run("base case - fail - 2", func(t *testing.T) {
		inMag := "aab"
		inNote := "abc"
		expected := false

		assert.Equal(t, expected, leet.CanConstruct(inNote, inMag))
	})

	t.Run("null case", func(t *testing.T) {
		inMag := ""
		inNote := "abc"
		expected := false

		assert.Equal(t, expected, leet.CanConstruct(inNote, inMag))
	})

	t.Run("null case 2", func(t *testing.T) {
		inMag := "abc"
		inNote := ""
		expected := true

		assert.Equal(t, expected, leet.CanConstruct(inNote, inMag))
	})

}

func TestIsIsomorphic(t *testing.T) {
	t.Run("base case", func(t *testing.T) {
		s := "egg"
		u := "add"
		expected := true
		assert.Equal(t, expected, leet.IsIsomorphic(s, u))
		assert.Equal(t, expected, leet.IsIsomorphic(u, s))
	})

	t.Run("null case", func(t *testing.T) {
		s, u := "foo", "bar"
		expected := false
		assert.Equal(t, expected, leet.IsIsomorphic(s, u))
		assert.Equal(t, expected, leet.IsIsomorphic(u, s))
	})

	t.Run("additional case", func(t *testing.T) {
		s, u := "paper", "title"
		expected := true
		assert.Equal(t, expected, leet.IsIsomorphic(s, u))
		assert.Equal(t, expected, leet.IsIsomorphic(u, s))
	})
}

func TestWordPattern(t *testing.T) {
	t.Run("base case", func(t *testing.T) {
		assert.Equal(t, true, leet.WordPattern("abba", "dog cat cat dog"))
	})

	t.Run("null case", func(t *testing.T) {
		assert.Equal(t, false, leet.WordPattern("abba", "dog cat cat fish"))
	})

	t.Run("null case 2", func(t *testing.T) {
		assert.Equal(t, false, leet.WordPattern("aaaa", "dog cat cat dog"))
	})

	t.Run("null case 3", func(t *testing.T) {
		assert.Equal(t, false, leet.WordPattern("abba", "dog dog dog dog"))
	})

	t.Run("null case 5", func(t *testing.T) {
		assert.Equal(t, false, leet.WordPattern("aba", "dog dog dog cat"))
	})

	t.Run("null case 6", func(t *testing.T) {
		assert.Equal(t, false, leet.WordPattern("abc", "dog cat dog"))
	})

	t.Run("null case 7", func(t *testing.T) {
		assert.Equal(t, false, leet.WordPattern("aaa", "aa aa aa aa"))
	})

	t.Run("null case 7", func(t *testing.T) {
		assert.Equal(t, false, leet.WordPattern("jquery", "jquery"))
	})

}

func TestIsAnagram(t *testing.T) {
	t.Run("base case", func(t *testing.T) {
		assert.Equal(t, true, leet.IsAnagram("anagram", "nagaram"))
	})

	t.Run("base case 2", func(t *testing.T) {
		assert.Equal(t, false, leet.IsAnagram("rat", "car"))
	})
}

func TestGroupAnagrams(t *testing.T) {
	t.Run("base case 2", func(t *testing.T) {
		assert.Equal(t, [][]string{{"cat", "tac"}, {"dog"}}, leet.GroupAnagrams([]string{"cat", "tac", "dog"}))
	})
}

func TestHappyNumber(t *testing.T) {
	t.Run("base case", func(t *testing.T) {
		assert.Equal(t, true, leet.HappyNumber(19))
	})

	t.Run("second case", func(t *testing.T) {
		assert.Equal(t, false, leet.HappyNumber(2))
	})

	t.Run("third case", func(t *testing.T) {
		assert.Equal(t, true, leet.HappyNumber(1))
	})

	t.Run("failed case", func(t *testing.T) {
		assert.Equal(t, true, leet.HappyNumber(7))
	})
}

func TestContainsDuplicate(t *testing.T) {
	t.Run("base case", func(t *testing.T) {
		assert.Equal(t, true, leet.ContainsDuplicate([]int{1, 2, 3, 1}, 3))
	})

	t.Run("base case", func(t *testing.T) {
		assert.Equal(t, true, leet.ContainsDuplicate([]int{1, 0, 1, 1}, 1))
	})

	t.Run("base case", func(t *testing.T) {
		assert.Equal(t, false, leet.ContainsDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))
	})
}

func TestPlusOne(t *testing.T) {
	t.Run("base case", func(t *testing.T) {
		assert.Equal(t, []int{2, 0}, leet.PlusOne([]int{1, 9}))
	})

	t.Run("base case 2", func(t *testing.T) {
		assert.Equal(t, []int{1, 0, 0}, leet.PlusOne([]int{9, 9}))
	})

	t.Run("base case 3", func(t *testing.T) {
		assert.Equal(t, []int{1, 1}, leet.PlusOne([]int{1, 0}))
	})
}

func TestMySqrt(t *testing.T) {
	t.Run("base case", func(t *testing.T) {
		assert.Equal(t, 2, leet.MySqrt(4))
	})

	t.Run("base case", func(t *testing.T) {
		assert.Equal(t, 1, leet.MySqrt(1))
	})

	t.Run("base case", func(t *testing.T) {
		assert.Equal(t, 5, leet.MySqrt(27))
	})

	t.Run("base case", func(t *testing.T) {
		assert.Equal(t, 2, leet.MySqrt(8))
	})

	t.Run("base case", func(t *testing.T) {
		assert.Equal(t, 0, leet.MySqrt(0))
	})

	t.Run("base case", func(t *testing.T) {
		assert.Equal(t, 1, leet.MySqrt(2))
	})
}
