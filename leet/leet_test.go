package leet_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"jack.com/leet/leet"
)

func TestTwoSum(t *testing.T) {
	nums := []int{2,7,11,15}
	target := 9 

	actual := leet.TwoSum(nums, target)

	assert.Equal(t, []int{1,0}, actual )
}

func TestTwoSum2(t *testing.T) {
	nums := []int{0, 4, 3, 0}
	target := 0 

	actual :=leet.TwoSum(nums, target)

	assert.Equal(t, []int{3,0}, actual )
}


func TestAddTwoNumbers( t *testing.T) {
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
		a := []int{ 3}
		b := []int{ -2, -1}

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

		expect :=  "AB"

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
		assert.Equal(t, 49, leet.MaxAreaPointers([]int{1,8,6,2,5,4,8,3,7}))
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
	t.Run("Sixes", func( t *testing.T) {
		assert.Equal(t, leet.RomanToInteger("VI"), 6)
	})

	t.Run("4", func( t *testing.T) {
		assert.Equal(t, leet.RomanToInteger("IV"), 4)
	})

	t.Run("Large", func( t *testing.T) {
		assert.Equal(t, leet.RomanToInteger("MMCXI"), 2111)
	})

	t.Run("THREEE", func( t *testing.T) {
		assert.Equal(t, leet.RomanToInteger("III"), 3)
	})

	t.Run("TEN", func( t *testing.T) {
		assert.Equal(t, leet.RomanToInteger("X"), 10)
	})

	t.Run("LV ", func( t *testing.T) {
		assert.Equal(t, leet.RomanToInteger("LVIII"), 58)
	})

	t.Run("Larger", func( t *testing.T) {
		assert.Equal(t, leet.RomanToInteger("MCMXCIV"), 1994)
	})

}

func TestLongestCommonPrefix(t *testing.T) {
	t.Run("No case", func(t *testing.T)  {
		assert.Equal(t, "", leet.LongestCommonPrefix([]string{"abc", "bcd"}))
	})

	t.Run("One", func(t *testing.T)  {
		assert.Equal(t, "a", leet.LongestCommonPrefix([]string{"abc", "acd"}))
	})

	t.Run("Many", func(t *testing.T)  {
		assert.Equal(t, "abc", leet.LongestCommonPrefix([]string{"abc", "abc"}))
	})

	t.Run("Leet1", func(t *testing.T)  {
		assert.Equal(t, "fl", leet.LongestCommonPrefix([]string{"flower", "flow", "flight"}))
	})

	t.Run("Leet2", func(t *testing.T)  {
		assert.Equal(t, "", leet.LongestCommonPrefix([]string{"dog","racecar","car"}))
	})

	t.Run("Empty", func(t *testing.T)  {
		assert.Equal(t, "", leet.LongestCommonPrefix([]string{""}))
	})
}


// func TestThreeSum(t *testing.T) {
// 	t.Run("No case", func(t *testing.T)  {
// 		out := [][]int{{1, -1, 0}}
// 		assert.Equal(t, out , leet.ThreeSum([]int{1, 1, -1 , 0}))
// 	})
// }

func TestIsValid( t *testing.T) {
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