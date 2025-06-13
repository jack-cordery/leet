package leet

import (
	"fmt"
	"log"
	"maps"
	"math"
	"strconv"
)

func TwoSum(nums []int, target int) []int {
	deltas := make(map[int]int)
	for i, num := range nums {
		delta := target - num
		if j, found := deltas[delta]; found {
			return []int{i, j}
		}
		deltas[num] = i
	}
	return []int{}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// List node represents an actual integer i.e. 425, except they are reversed and structures
//ListNode(Val: 5, Next: ListNode(Val: 2, Next: ListNode(Val: 4, Next: nil)))
//ListNode(Val: 5, Next: ListNode(Val: 2, Next: ListNode(Val: 4, Next: nil)))

func AddTwoNumbersL(l1 *ListNode, l2 *ListNode) *ListNode {
	currentl1 := l1
	currentl2 := l2
	a := currentl1.Val
	b := currentl2.Val
	lPrev := &ListNode{Val: (a + b) % 10}
	carry := (a + b) / 10

	for {
		var l1break bool
		var l2break bool
		if currentl1.Next != nil {
			currentl1 = currentl1.Next
		} else {
			currentl1 = &ListNode{Val: 0}
			l1break = true
		}
		if currentl2.Next != nil {
			currentl2 = currentl2.Next
		} else {
			currentl2 = &ListNode{Val: 0}
			l2break = true
		}
		if l1break && l2break {
			if carry != 0 {
				return &ListNode{Val: carry, Next: lPrev}
			} else {
				return lPrev
			}
		}
		a = currentl1.Val
		b = currentl2.Val
		digit := (a + b + carry) % 10
		carry = (a + b + carry) / 10
		lPrev = &ListNode{Val: digit, Next: lPrev}
	}
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy, sum := new(ListNode), 0
	for curr := dummy; l1 != nil || l2 != nil || sum != 0; curr = curr.Next {
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		curr.Next = &ListNode{Val: sum % 10}
		sum /= 10
	}
	return dummy.Next
}

func LengthOfLongestSubstring(s string) int {
	store, left, right, result := make(map[uint8]int), 0, 0, 0

	for right < len(s) {
		var r = s[right] // r is the rune
		store[r] += 1
		for store[r] > 1 {
			l := s[left]
			store[l] -= 1
			left += 1
		}
		result = max(result, right-left+1)

		right += 1
	}
	return result
}

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	var ans float64
	l1, l2 := 0, 0
	r1, r2 := len(nums1)-1, len(nums2)-1
	for {
		if r1-l1 == 0 && r2-l2 == 0 {
			ans = float64(nums1[r1]+nums2[r2]) / 2
			break
		}
		// Past left
		if r1-l1 < 0 {
			if r2-l2 < 2 {
				if r2-l2 < 1 {
					ans = float64(nums2[r2])
				} else {
					ans = float64(nums2[r2]+nums2[l2]) / 2
				}
				break
			}
			l2++
			r2--
			continue
		}
		// Past right
		if r2-l2 < 0 {
			if r1-l1 < 2 {
				if r1-l1 < 1 {
					ans = float64(nums1[r1])
				} else {
					ans = float64(nums1[r1]+nums1[l1]) / 2
				}
				break
			}
			l1++
			r1--
			continue
		}

		if nums1[l1] <= nums2[l2] {
			l1++
		} else {
			l2++
		}

		if nums1[r1] >= nums2[r2] {
			r1--
		} else {
			r2--
		}

	}
	return ans
}

func LongestPalindrome(s string) string {
	// Find the longest string that is itslef forward and backwards
	// Brute force would be check all combos to see if they are a palindrome, pick largest
	// Pretty sure you can get this is < O(n^2)
	// Similar to longest unqique substring
	// Can i go through storing the previous valid palindrome chars, resetting when no longer valid
	// This doenst work as a substring of a palindeom is not neccessarily a palindrone. INTERERESTINT
	// Sliding window from largest and stop when i find my first
	arr := []byte(s)
	for r := len(arr); r > 0; r-- {
	OuterLoop:
		for i := 0; i <= len(arr)-r; i++ {
			window := arr[i : i+r]
			for j := 0; j <= len(window)/2; j++ {
				if window[j] != window[len(window)-1-j] {
					continue OuterLoop
				}
			}
			return string(window)
		}
	}
	return ""
}

func LongestPalindromeQ(s string) string {
	// What we can actually do is use Manacher's algo which gets its motif from the fact that every character
	// can be the centre of a palindrome and so you can go O(N) through the list and find the longest pal from that
	// character and then just return the max
	return ""
}

func ZigZagConvert(s string, numRows int) string {

	// First row and last row only have one step which are i + 2(n-1) apart
	b := []byte(s)
	l := len(s)
	var o []byte
	if len(s) == 0 {
		return ""
	}
	if len(s) == 1 || numRows == 1 {
		return s
	}
	for i := 0; i < numRows; i++ {
		fmt.Println("i", i)
		for j := 0; float32(j) < float32(l-i)/float32((2*(numRows-1))); j++ {
			step := i + (2 * j * (numRows - 1))
			fmt.Println("j", j)
			o = append(o, b[step])
			if step+2*(numRows-(i+1)) < l && i != 0 && i != numRows-1 {
				o = append(o, b[step+2*(numRows-(i+1))])
			}
			fmt.Println("cond", float32(l-i)/float32((2*(numRows-1))))
		}
	}
	return string(o)
}

func Reverse(x int) int {
	// Takes in int and returns the reversed number. If that number goes outside of uint32
	// then returns 0
	// 32 bits is [-2^31, 2^31 -1 ]
	// Im thinking deal with it as a string and reverse with some checks to turn 0
	// or as just an array with the sign taken at the start?
	// O(N)
	isPositive := x >= 0
	if float64(x) >= math.Pow(2, 31) || float64(x) < -math.Pow(2, 31) {
		return 0
	}

	var o, xArray []byte
	if isPositive {
		xArray = []byte(strconv.Itoa(x))
		o = append(o, '+')
	} else {
		xArray = []byte(strconv.Itoa(x)[1:])
		o = append(o, '-')
	}

	fmt.Println(strconv.Itoa(x), len(xArray))
	for j := len(xArray); j > 0; j-- {
		o = append(o, xArray[j-1])
	}
	r64, err := strconv.ParseInt(string(o), 10, 32)
	if err != nil {
		return 0
	}

	return int(r64)
}

func Reverse32(x int) int {
	// This time we are going to do it by going digit by digit so avoiding string conversion
	// which feels both faster, safer and more elegant
	x32 := int32(x)
	base := int32(10)
	var result, newResult int32
	result = 0
	for x32 != 0 {
		digit := (x32 % base)
		newResult = result + digit
		if (newResult - result) != digit {
			return 0
		}
		x32 -= digit
		x32 /= base
		if x32 != 0 {
			if ((newResult * base) / base) != newResult {
				return 0
			}
			result = newResult * base
		} else {
			result = newResult
		}
	}
	return int(result)
}

func IsPalindrome(x int) bool {
	// So we can read in any int and then see if its a palindorme
	// Any negative value automatically is as -121 is 121-
	if x < 0 {
		return false
	}
	// Start in the centre and step outwards checking if they are the same
	// but we dont really want to convert to strings, or get byte arrays
	// maybe there is a numeric trick?
	// Maybe just create the palindrome and then check qequiv // much slower
	// We can use log_10 to work out the largest digit too and so we can work from outward in!
	bigBase := int(math.Log10(float64(x)))

	for i := 0; i < bigBase; i++ {
		// So this loop goes up to half way
		lowerBase := int(math.Pow(10, float64(i+1)))
		upperBase := int(math.Pow(10, float64(bigBase-i)))
		lDigit := (x / upperBase) % 10
		rDigit := (x % lowerBase) / (lowerBase / 10)
		if lDigit != rDigit {
			return false
		}
	}

	return true
}

func MaxArea(height []int) int {
	// I cant think of a trick that would allow me to do anythong more
	// efficeint than checking
	// Im going to implement it by looping through window sizes and going
	// through all the permutations
	maxArea := 0
	for windowSize := len(height) - 1; windowSize > 0; windowSize-- {
		for i := 0; i <= (len(height)-1)-windowSize; i++ {
			j := i + windowSize
			fmt.Println("loop", windowSize, i, j)
			area := min(height[j], height[i]) * (j - i)
			fmt.Println("inside", area, maxArea)
			maxArea = max(maxArea, area)
		}
	}
	return maxArea
}

func MaxAreaPointers(height []int) int {
	// We are going to use pointers to check between two values and then slide the
	// left one across
	maxArea := 0
	i, j := 0, len(height)-1
	for i < j {
		maxArea = max(maxArea, (j-i)*min(height[i], height[j]))
		if height[i] <= height[j] {
			i++
		} else {
			j--
		}
	}
	return maxArea
}

func IntegerToRoman(x int) string {
	// So we just need a function that for an integer, less than 4000,
	// it outputs the roman numeral version
	outs := ""
	romans := make(map[int]string)
	romans[1] = "I"
	romans[5] = "V"
	romans[10] = "X"
	romans[50] = "L"
	romans[100] = "C"
	romans[500] = "D"
	romans[1000] = "M"

	base := 10
	for x > 0 {
		d := x % base
		leadingDigit := d * 10 / base
		digitString := ""

		fmt.Println("Start", x, base, d, leadingDigit, digitString, outs)

		switch leadingDigit {
		case 0:
			digitString = ""
		case 1:
			digitString = romans[base/10]
		case 2:
			digitString = romans[base/10] + romans[base/10]
		case 3:
			digitString = romans[base/10] + romans[base/10] + romans[base/10]
		case 4:
			digitString = romans[base/10] + romans[5*base/10]
		case 5:
			digitString = romans[5*base/10]
		case 6:
			digitString = romans[5*base/10] + romans[base/10]
		case 7:
			digitString = romans[5*base/10] + romans[base/10] + romans[base/10]
		case 8:
			digitString = romans[5*base/10] + romans[base/10] + romans[base/10] + romans[base/10]
		case 9:
			digitString = romans[base/10] + romans[base]
		default:
			log.Fatal("Unknown digit")
		}

		x -= d
		base *= 10

		outs = digitString + outs

		fmt.Println("Ends", x, base, d, leadingDigit, digitString, outs)

	}

	return outs
}

func RomanToInteger(roman string) int {
	// We are going to convert some roman numeral strings to int

	// Create map for strings to ints
	romans := make(map[string]int)
	romans["I"] = 1
	romans["IV"] = 4
	romans["V"] = 5
	romans["IX"] = 9
	romans["X"] = 10
	romans["XL"] = 40
	romans["L"] = 50
	romans["XC"] = 90
	romans["C"] = 100
	romans["CD"] = 400
	romans["D"] = 500
	romans["CM"] = 900
	romans["M"] = 1000

	rommanArray := []byte(roman)
	l := len(rommanArray)
	total := 0

	for l > 1 {
		fmt.Println("LoopS", l)
		chars := string(rommanArray[l-2]) + string(rommanArray[l-1])
		value := romans[string(chars)]
		if value == 0 {
			// its invalid therefore its a single value
			total += romans[string(rommanArray[l-1])]
			l--
		} else {
			total += value
			l -= 2
		}
		fmt.Println("LoopE", l, chars, value, total)
	}

	if l == 1 {
		return total + romans[string(rommanArray[0])]
	} else {
		return total
	}
}

func LongestCommonPrefix(strs []string) string {
	// Loop through the list of strings on each character checking whether there is a common prefix and continuining until there isnt
	working := true
	i := 0
	output := ""
	smLength := 400
	for _, str := range strs {
		smLength = min(smLength, len(str))
		if str == "" {
			return output
		}
	}

	for working {
		char := strs[0][i]
		for _, str := range strs {
			newChar := str[i]
			if newChar != char {
				working = false
				break
			}
			char = newChar
		}
		if working {
			i++
			output += string(char)
		}
		if len(output) == smLength {
			return output
		}
	}

	return output
}

// func ThreeSum(ints []int) [][]int {
//     // We are going to implement the
//     // for each value is there a pair that sums to the negative
//     out := [][]int{}
//     for i:=0; i < len(ints); i++ {
//         var restOfSlice []int
//         if i == 0  {
//             restOfSlice = ints[1:]
//         } else {
//             restOfSlice = append(ints[:i - 1], ints[i + 1:]...)
//         }
//         // I want to loop through the rest of the slices looking for two combos, stroing to make sure i dont bother with duplicates

//         out = append(out, append(twos, ints[i]))
//     }
//     return out
// }

func IsValid(s string) bool {

	// store state with an array of runes i.e. ["("] then can essentially check
	// that they are opened and closed in the right order
	// Think will try and use pointers to check for valid closes and track opens
	var state []byte
	bracketMap := make(map[byte]byte)
	bracketMap[byte(')')] = byte('(')
	bracketMap[byte(']')] = byte('[')
	bracketMap[byte('}')] = byte('{')

	for i := 0; i < len(s); i++ {
		// if j is an open bracket then set that to active
		// if j is a close bracket check that it was opened and then set to not active
		// if not opened then fail
		currentChar := string(s[i])
		if currentChar == "[" || currentChar == "(" || currentChar == "{" {
			// open
			state = append(state, s[i])
		} else {
			//close
			if len(state) == 0 {
				return false
			}
			if state[len(state)-1] != bracketMap[s[i]] { // not sure if s[i] is the byte representation
				// not closed corectectly so return false
				return false
			}
			state = state[:len(state)-1]
		}
	}
	return len(state) == 0
}

// TODO: So this works apart from the duplication condition, i thought i could get around it with maps to make distinct but alas that didnt work
// Im sure with a slightly different approach it would resolve
func ThreeSum(nums []int) [][]int {
	// So as discussed in the notes I am going to take the approach of looping through
	// taking each number a_i and then loop through the list to the right checking values that - sum to that value
	var output [][]int
	for i := 0; i < len(nums)-2; i++ {
		// so now we need to get the value at i and get the slice to the right of it that we will loop through
		val, rest := nums[i], nums[i+1:]
		fmt.Println("Start of loop", i, val, rest)
		deltas := make(map[int]int)
		for j := 0; j < len(rest); j++ {
			delta := -val - rest[j]
			fmt.Println("Inside of loop", j, delta)
			if j, found := deltas[rest[j]]; found {
				output = append(output, []int{val, rest[j], -(val + rest[j])})
				fmt.Println("Found loop", j, found, delta, val, rest[j])
			}
			deltas[-val-rest[j]] = j
			fmt.Println("Exit loop", deltas)
		}
	}
	return output
}

func ThreeSumClosest(nums []int, target int) int {
	// So this is the same as threeSum apart from that we just want to return the closest
	// possible sum to a given value
	// The idea then is to three sum as before just tracking the closest deltas
	// Early break if we hit zero

	return 0
}

func LetterCombinations(digits string) []string {
	// So i want to take an input of a string of digits that represent a mapping of a
	// old phone keyboard i.e. 2 -> a,b,c, etc. Return all combos
	numberMap := make(map[rune][]string)
	numberMap['2'] = []string{"a", "b", "c"}
	numberMap['3'] = []string{"d", "e", "f"}
	numberMap['4'] = []string{"g", "h", "i"}
	numberMap['5'] = []string{"j", "k", "l"}
	numberMap['6'] = []string{"m", "n", "o"}
	numberMap['7'] = []string{"p", "q", "r", "s"}
	numberMap['8'] = []string{"t", "u", "v"}
	numberMap['9'] = []string{"w", "x", "y", "z"}

	if digits == "" {
		return []string{}
	}

	// brute force
	// loop through each number and for each number
	result := numberMap[rune(digits[0])]
	for i := 1; i < len(digits); i++ {
		digit := digits[i]
		loopCharSet := numberMap[rune(digit)]
		newResult := []string{}
		for _, r := range result {
			// We want to iterate over the current reuslt and add all permutation for next digit
			// "a" -> "ad", "ae", "af"
			for _, c := range loopCharSet {
				newResult = append(newResult, r+c)
			}
		}
		result = newResult
	}
	return result
}

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	// So the interesting part of this is how do i know its the nth from end. Simplest way would be to go all the way to the end esentially unpacking it all and then repacking
	// Is there a cooler/smarter way?
	nextNode := head

	valueStore := []int{}
	count := 0

	for nextNode != nil {
		// Store values in an array and count the number to be repacked after
		count++
		valueStore = append(valueStore, nextNode.Val)
		nextNode = nextNode.Next

	}

	if count == 1 {
		return nil
	}

	result := new(ListNode)
	firstTracker := 0
	for i := count - 1; i >= 0; i-- {
		curr := new(ListNode)
		fmt.Println("Loop", i, curr)
		if i == count-n {
			// this is the one to remove so just skip
			continue
		} else {
			curr.Val = valueStore[i]
			if firstTracker == 0 {
				curr.Next = nil
				firstTracker++
			} else {
				curr.Next = result
			}
			result = curr
		}
	}

	return result
}

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// We are given two lists sorted in asceding order. We need to return a list containing all the elements of each lists, in ascending order
	// Immedialty feels like a two pointer problem. i.e. have a pointer at a positiion on each list then incremenet the pointer of the smaller value and add it to the result
	//
	result := new(ListNode)
	current := result
	for list1 != nil || list2 != nil {
		if (list2 != nil && list1 != nil && list1.Val <= list2.Val) || (list2 == nil) {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	return result.Next
}

func GenerateParantheses(n int) []string {
	parantheses := []string{"()"}

	for i := 1; i < n; i++ {
		result := []string{}
		for _, p := range parantheses {
			p1 := "(" + p + ")"
			p2 := "()" + p
			if p2 != p+"()" {
				p3 := p + "()"
				result = append(result, p1, p2, p3)
			} else {
				result = append(result, p1, p2)
			}
			parantheses = result
		}
	}
	return parantheses
}

func SwapPairs(head *ListNode) *ListNode {
	// Swap alternating pairs, if its even then we swap all pairs
	// IF odd swap all until the last node
	// Are supposed to swap nores not values
	// Iterate through list and just replace the current list node with the next
	dummy := &ListNode{Next: head}
	prev, first := dummy, head

	for first != nil && first.Next != nil {
		nextPair := first.Next.Next
		second := first.Next

		second.Next = first
		first.Next = nextPair
		prev.Next = second

		prev = first
		first = nextPair
	}

	return dummy.Next
}

// TODO: Finish this
// func ReverseKGroup(head *ListNode, k int) *ListNode {
//     // So i think the idea is that I have to take a hit somehwere. Either i would need to know in advance at what point to stop. Or i would need to backtrack
//     // The former gives me a bit better simplicity in terms of code, and is only worse marginally
//     // Idea is that by knowing the length of the linked list in advance i can know in advance when to stop for the last

//     if k == 1 {
//         return head
//     }

//     lengthDummy, length := head, 0
//     for lengthDummy != nil {
//         length++
//         lengthDummy = lengthDummy.Next
//     }
//     fmt.Printf("head is %d long \n", length)

//     nGroups := length / k
//     remainder := length - length % k

//     fmt.Printf("number of groups in linked list is %d and remainder is %d \n", nGroups, remainder)

//     // Now I can reverse nodes up to the k-th element, reset and do that nGroup times and then stop leaving the remainders

//     groupTrack := 0
//     globalTrack := 0
//     var prev, next *ListNode
//     curr := head

//     for curr != nil{
//         fmt.Printf("Loop: groupTrack %d, globalTrack %d \n", groupTrack, globalTrack)
//         fmt.Println("Curr: ", curr)
//         fmt.Println("prev:", prev)
//         fmt.Println("next: ", next)
//         if groupTrack < k  && globalTrack < remainder  {
//             fmt.Printf("Reverse block: ")
//             next = curr.Next
//             curr.Next = prev
//             fmt.Println("Curr: ", curr)
//             fmt.Println("Curr.Next: ", curr.Next)
//             fmt.Println("prev:", prev)
//             fmt.Println("next: ", next)

//             prev = curr
//             curr = next
//             fmt.Println("Curr: ", curr)
//             fmt.Println("prev:", prev)
//             fmt.Println("next: ", next)
//             groupTrack++
//         } else {
//             fmt.Printf("Skip block: ")
//             // Here i need to set the
//             //curr.Next = prev
//             // I need a dummy to point the first element at rather than nil

//             prev = curr
//             curr = curr.Next
//             fmt.Println("Curr: ", curr)
//             fmt.Println("prev:", prev)
//             fmt.Println("next: ", next)
//             groupTrack = 0
//         }
//         globalTrack++
//     }

//     // So full reverse works! thats good
//     // All the others fail, not so good.

//     // temp, count := prev, 0
//     // for temp != nil {
//     //     count++
//     //     fmt.Println(count,  temp)
//     //     temp = temp.Next
//     // }
//     return prev
// }

func MergeSortedArray(nums1, nums2 []int, m, n int) {
	for n > 0 || m > 0 {
		fmt.Println("pos: ", m, n)
		var candidate1, candidate2 int
		if n == 0 {
			break
		}
		if m == 0 {
			nums1[n-1] = nums2[n-1]
			n--
			continue
		}
		candidate1, candidate2 = nums1[m-1], nums2[n-1]

		if candidate1 > candidate2 {
			nums1[n+m-1] = candidate1
			m--
		} else {
			nums1[n+m-1] = candidate2
			n--
		}
	}
}

func RemoveElement(nums []int, val int) int {
	count := 0
	for i, tracker := 1, 0; i <= len(nums); i++ {
		fmt.Println("Loop: ", i, tracker, count, nums)
		if nums[i-1] == val {
			count++
			continue
		}
		nums[tracker] = nums[i-1]
		tracker++
	}
	return count
}

func RemoveDuplicates(nums []int) int {
	// remove duplicates in places, preserving rest of the order, and retyurning the count of unique elements
	count := 0
	for i, tracker, prev := 0, 0, 0; i < len(nums); i++ {
		// Given its in order we just need to check the previous value is the same
		current := nums[i]
		if current != prev || i == 0 {
			// its unique so replace value and move tracker
			nums[tracker] = current
			prev = current
			tracker++
			count++
		}
	}
	return count
}

func RemoveDuplicates2(nums []int) int {
	// remove duplicates in places, preserving rest of the order, and retyurning the count of unique elements
	count := 0
	duplicates := make(map[int]int)
	for i, tracker := 0, 0; i < len(nums); i++ {
		// Given its in order we just need to check the previous value is the same
		current := nums[i]
		if c, ok := duplicates[current]; c < 2 || i == 0 {
			// its unique so replace value and move tracker
			nums[tracker] = current
			if ok {
				duplicates[current]++
			} else {
				duplicates[current] = 1
			}
			tracker++
			count++
		}
	}
	return count
}

func MajorityElement(nums []int) int {
	count := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		count[nums[i]]++
	}

	maxElement := 0
	maxCount := 0
	for num, c := range count {
		if c > maxCount {
			maxElement = num
			maxCount = c
		}
	}
	return maxElement
}

func RotateArray(nums []int, k int) {
	// rotate array to the right by k steps where k >= 0
	// It means move everything along the line.
	// Needs to be done in place - which makes it slightly interesting
	// Imagine the array is a circle. It means make the starting postion 3 to the right
	// Think you can do it with some sort of mod
	// Input: nums = [1,2,3,4,5,6,7], k = 3
	// Output: [5,6,7,1,2,3,4]
	// [n-k, n-k+1, n-k+2, 0, 1, 2, 3] until n-k+j = len(nums)
	// [4,5,6,0,1,2,3]
	// [i + 1 + 3 mod(7)] -> [4, 5, 6, 0, 1,2,3]
	// if k = 4
	// [i + 1 + (l - k ) mod (l)]
	// k = 3
	// [i+1 + 4  mod(l)]
	// [4 - k mod(l)]
	// [ i + (l-k) mod(l)]
	// k = 3 -> [i + (4) mod(l)]
	// [4, 5, 6, 0, 1 , 2 ,3]
	// k = 4 -> [i + 3 mod(l)]
	// [3,4,5,6,7,0,1,2]
	// Now just need to figure out the order to update so that i dont overwrite
	l := len(nums)
	if k == 0 {
		return
	}
	// fmt.Println("divisibility of ",k, l, " is ", (l % k) == 0, l / k == 2 )
	// if  (l % k) == 0 && l / k == 2 {
	//     // if even swap paris
	//     fmt.Println("even case")
	//     for i:= 0; i < l/2; i++ {
	//         dummy:= nums[(i + k) % l]
	//         fmt.Println("dummy set to: ", dummy)
	//         nums[(i + k) % l] = nums[i]
	//         fmt.Println("pos: ", (i +k) % l, "set to ", nums[(i+k)%l])
	//         nums[i] = dummy
	//         fmt.Println("pos: ", i, "set to ", nums[i])
	//     }
	//     return
	// }
	// Is it divisible?
	// We need to do swaps for each of the pairs
	aa, mm := k, l
	for mm != 0 {
		aa, mm = mm, aa%mm
	}
	g := aa // gcd(a, m)

	// Compute result
	cycleLength := l / g

	numCycles := l / cycleLength
	fmt.Println("divisible cycleLength: ", cycleLength, "numCycles: ", numCycles)
	for i := 0; i < numCycles; i++ {
		prev := nums[(i)%l]
		fmt.Println("prev set to: ", prev)
		for j := 0; j < cycleLength; j++ {
			fmt.Println("i: ", i, "j: ", j)
			dummy := nums[(i+(j+1)*k)%l]
			fmt.Println("dummy set to: ", dummy)
			nums[(i+(j+1)*k)%l] = prev
			fmt.Println("pos: ", (i+(j+1)*k)%l, "set to ", nums[(i+(j+1)*k)%l])

			prev = dummy
			fmt.Println("prev set to: ", prev)
		}
	}
} // We need to do swaps for each of the pair

func MaxProfit(prices []int) int {
	// Return the largest delta between sequential numbers i.e. [1,2,3,0,5,6,7] => 7
	maxDelta := 0
	min := prices[0]
	for i := 1; i < len(prices); i++ {
		p := prices[i]
		if p < min {
			min = p
		} else {
			delta := p - min
			if delta > maxDelta {
				maxDelta = delta
			}
		}
	}
	return maxDelta
}

func MaxProfit2(prices []int) int {
	maxProfit := 0
	for i := 0; i < len(prices)-1; i++ {
		curr := prices[i]
		next := prices[i+1]

		if next > curr {
			maxProfit += next - curr
		}
	}
	return maxProfit
}

func CanJump(nums []int) bool {
	// loop through sotring the max value. When we get to a zero if ifs not big enough to get passt then false, otherwise true
	maxJump := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == 0 {
			if maxJump <= i {
				return false
			}
		} else {
			currJump := nums[i] + i
			if currJump > maxJump {
				maxJump = currJump
			}
		}

	}
	return true
}

func CanConstruct(ransomNote, magazine string) bool {
	// Principle is - given a magazine can you construct the ransomNote you wish
	// So do i have all of the characters available from the magazine
	// Intiial thought is to create a hashmap from magazine char -> numOccurences
	// If it turns out that you only need the character once in the magazine then the changes are:
	// -> then we just remove the decrement in the second loop
	// then iterate through ransomeNote checking for existing map value and then decrementing the count in the map
	// Perfomance questions would center around if i could add in some performance gains for edge cases
	// For example first checking if ransomNote < magazine
	// IF magazine >>> ransomeNote. Would it be better to just search magazine directly without conversion
	// to hashmap
	magMap := make(map[rune]int)
	for _, r := range magazine {
		magMap[r]++
	}
	fmt.Printf("mag map is %v", magMap)
	for _, r := range ransomNote {
		if countR, ok := magMap[r]; countR < 1 || !ok {
			fmt.Printf("Issue found character %s not in magMap %v", string(r), magMap)
			return false
		}
		magMap[r]--
	}
	return true
}

func IsIsomorphic(s, t string) bool {
	// are they isomorphic i.e. s <=> t
	// i.e. egg, add as e <-> a, g <-> d
	// initial thoyghts is to loop through the string and store the map in a hashmap and then if there is already an occurence
	// of that value in the key then we return false otherwise we return true
	// Hashmap  e -> a
	// g -> d

	//  f -> b
	// o -> a
	// o -> r // fails

	// b -> f
	// a -> o
	// r -> o // falis
	//

	// isnt it just a matter of if i have n unique characters on both sides? probably same O(n)
	m := make(map[byte]byte, len(s))
	m2 := make(map[byte]byte, len(t))
	for i := 0; i < len(s); i++ {
		currS := s[i]
		currT := t[i]
		r, ok := m[currS]
		u, ok2 := m2[currT]
		if (ok && r != currT) || (ok2 && u != currS) {
			return false
		}
		m[currS] = currT
		m2[currT] = currS
	}
	return true
}

func WordPattern(pattern, s string) bool {
	currWord := ""
	patternToWord := make(map[byte]string)
	wordToPattern := make(map[string]byte)
	patternTracker := 0
	maxP := len([]byte(pattern))
	for _, sChar := range s + " " {
		if patternTracker >= maxP {
			return false
		}
		if sChar == ' ' {
			w, wOk := patternToWord[pattern[patternTracker]]
			p, pOk := wordToPattern[currWord]

			if wOk && w != currWord || pOk && p != pattern[patternTracker] {
				return false
			}

			patternToWord[pattern[patternTracker]] = currWord
			wordToPattern[currWord] = pattern[patternTracker]

			currWord = ""
			patternTracker++
		} else {
			currWord += string(sChar)
		}
	}
	fmt.Println(patternTracker, maxP)
	return patternTracker == maxP
}

func IsAnagram(s, t string) bool {
	// is t an anagram of s
	// every letter of s is used exactly once in t
	// there should be a map such that i have index_s -> index_t
	// fail if that isnt the case
	tMap := make(map[rune]int)
	sMap := make(map[rune]int)
	for _, c := range t {
		tMap[c]++
	}
	for _, c := range s {
		sMap[c]++
	}
	return maps.Equal(sMap, tMap)
}

func GroupAnagrams(strs []string) [][]string {
	// So we want to be able take an array of words and group together the anagrams.

	return [][]string{{}}
}

func HappyNumber(n int) bool {
	// a number is happy if
	// the sum of squares of its digtis -> repeat until either it equals 1 or it infinitely loops
	// ex 2 -> 4 -> 16 -> 1 + 36 = 37 -> 9 + 49 = 58 -> 25 + 64 -> 89 -> 64 + 81 -> 145 -> 1 + 16 + 25 -> 42 -> 16 + 4 -> 20 -> 4 (loop)
	// initial thought is that we would need to do this until either it finds a number that already has been had, or it will hit 1
	// There may be a better mathematical reprs of this however? i.e. immediately tell if two digits are happy or not
	sumMap := make(map[int]int)
	for {
		digits := []int{}
		for n > 0 {
			digits = append(digits, n%10)
			n /= 10
		}
		sumD := 0
		for _, d := range digits {
			sumD += d * d
		}
		fmt.Printf("Digits: %v sumD: %v", digits, sumD)
		if sumD == 1 {
			return true
		}
		if _, ok := sumMap[sumD]; ok {
			return false
		} else {
			sumMap[sumD] = 1
			n = sumD
		}
	}
}

func ContainsDuplicate(nums []int, k int) bool {
	// Actually its. Given an array are there two indentical values that are within k positions
	// Intial idea is to have a map that maps value to position.
	// Then we can check that map and calculate the delta and if in bound return true
	posMap := make(map[int]int)
	for i, v := range nums {
		m, ok := posMap[v]
		if delta := i - m; ok && delta <= k {
			return true
		}
		posMap[v] = i
		fmt.Printf("v: %v is updated with %v\n", v, i)
	}
	return false
}

func PlusOne(digits []int) []int {
	r := 1
	result := digits
	dTracker := len(digits) - 1
	for r > 0 {
		result[dTracker] += r
		if result[dTracker] > 9 {
			result[dTracker] = 0
			if dTracker == 0 {
				result = append([]int{0}, result...)
			} else {
				dTracker--
			}
			r = 1
		} else {
			r = 0
			dTracker--
		}
	}
	return result
}

func MySqrt(x int) int {
	// TODO: Can be cleaned up to reduce control logic complexity. For example can add check in for statement
	if x <= 1 {
		return x
	}
	yUpper := x
	yLower := 0
	for {
		d := (yUpper - yLower) / 2
		m := yLower + d
		if d == 0 {
			fmt.Println("delta is zero yUpper yLower m", yUpper, yLower, m)
			return m
		}
		fmt.Println("for x ", x, "we have yUppper yLower d m ", yUpper, yLower, d, m)
		if m*m < x {
			fmt.Println("LESS m m*m x", m, m*m, x)
			yLower = m
		} else if m*m > x {
			fmt.Println("MORE m m*m x", m, m*m, x)
			yUpper = m
		} else {
			return m
		}
	}
}

func MaxDepth(root *TreeNode) int {
	// Given a binary tree find the maxDepth of it, i.e. how many steps can i take before hitting nil tree
	if root == nil {
		return 0
	}
	return 1 + max(MaxDepth(root.Left), MaxDepth(root.Right))
}

func IterativeDFSMaxDepth(root *TreeNode) int {
	// ok so we are going to use a stack and add children and cross off
	// need to think of a wayt to track the depth, sure it will become obvious when we create iteration
	// stack is lifo.  Again go until the stack is empty
	stack := []struct {
		pos   *TreeNode
		depth int
	}{{root, 1}}
	maxDepth := 0
	for len(stack) > 0 {
		curr := stack[0]
		if curr.pos.Left != nil {
			stack = append(stack, struct {
				pos   *TreeNode
				depth int
			}{curr.pos.Left, curr.depth + 1})
		}
		if curr.pos.Right != nil {
			stack = append(stack, struct {
				pos   *TreeNode
				depth int
			}{curr.pos.Right, curr.depth + 1})
		}

		maxDepth = max(maxDepth, curr.depth)
		stack = stack[1:]
	}
	return maxDepth
}

func IterativeBFSMaxDepth(root *TreeNode) int {
	/// fill a queue replacing eleements with its children
	// until there is nothing in queue. Somehow keep track of level
	if root == nil {
		return 0
	}
	q := []*TreeNode{root}
	levels := 0
	for len(q) > 0 {
		l := len(q)
		for range l {
			// this is a level/ Go through the whole queue
			n := q[0]
			q = q[1:]
			if n.Left != nil {
				q = append(q, n.Left)
			}
			if n.Right != nil {
				q = append(q, n.Right)
			}
		}

		levels++
	}
	return levels
}

func LegnthOfLastWord(s string) int {
	// find the position of the last whitespace and the second last whitespace
	e := len(s) - 1
	for s[e] == ' ' {
		e--
	}
	b := e
	for (b >= 0) && (s[b] != ' ') {
		b--
	}
	return e - b

}

func SearchInsert(nums []int, target int) int {
	// Use binary search to find the index where target would be or is
	// general idea is to go through the array using bin search
	// if i find target then return that index otherwise ensure we land on the
	// index it would be - think it will be more obvious when we do it
	// TODO: fix for edge case [1, 3], 2 <- currently returns 0 instead of 1 as it returns iLower
	l := len(nums)
	iUpper := l
	iLower := 0
	r := 0
	for delta := 0; delta <= 0; {
		delta = iUpper - iLower
		m := iLower + (delta / 2)
		if nums[m] == target {
			return m
		} else if nums[m] < target {
			// look at uppper half
			iLower = m + 1
			r = iUpper
		} else {
			// look at lower half
			iUpper = m - 1
			r = iLower
		}
	}
	return r
}

func LinkedListCycle(head *ListNode) bool {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}

func DeleteDuplicates(head *ListNode) *ListNode {
	// given a sorted linked list, delete all duplicates and return sorted
	// immediate thought is to just check if next is equal and only set next if not equal
	copy := head
	for head != nil && head.Next != nil {
		for head.Val == head.Next.Val && head.Next.Next != nil {
			println(head.Val, head.Next.Val)
			head.Next = head.Next.Next
		}
		head = head.Next
	}
	return copy
}

func InvertBinaryTree(root *TreeNode) *TreeNode {
	// invert in this case means flip left to right
	// so swap left and right
	// Can try both iteratively and recurisevly
	if root == nil {
		return root
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		curr.Right, curr.Left = curr.Left, curr.Right

		if curr.Left != nil {
			q = append(q, curr.Left)
		}
		if curr.Right != nil {
			q = append(q, curr.Right)
		}

	}
	return root
}

func CountNodes(root *TreeNode) int {
	// so this counts the number of nodes but we are also told that the tree is complete
	// sucht that if you consider the tree from the top left there are no empty vals
	// i think that means it is most efficient to do BFS and then just
	// count the number of levels and the number of nodes in last level
	// and retrun Sum_i{2^i} + a . where i is number of full rows and a number in last
	// assume that the summation has a numeric formula 2^i+1 - 1j

	if root == nil {
		return 0
	}
	q := []*TreeNode{root}
	nFullLevels := 0

	for len(q) > 0 {
		l := len(q)
		if l < (1 << nFullLevels) {
			// reached the first not full level
			// there are nFullLevels
			// and l leaves
			return (1 << (nFullLevels)) - 1 + l
		}
		for range l {
			curr := q[0]
			q = q[1:]
			if curr.Left != nil {
				q = append(q, curr.Left)
			}
			if curr.Right != nil {
				q = append(q, curr.Right)
			}
		}
		nFullLevels++

	}
	// they are all full levels
	fmt.Println("n full levels found", nFullLevels)
	return (1 << (nFullLevels)) - 1
}

func GetMinimumDifference(root *TreeNode) int {
	// do a left trav
	// pop and then set to right
	stack := []*TreeNode{}
	curr := root
	prevVal := -1
	minDelta := math.MaxInt
	for len(stack) > 0 || curr != nil {
		for curr != nil {
			stack = append([]*TreeNode{curr}, stack...)
			curr = curr.Left
		}
		fmt.Println("pop! ", stack[0])
		if prevVal >= 0 {
			d := prevVal - stack[0].Val
			if d < 0 {
				d *= -1
			}
			minDelta = min(minDelta, d)
		}
		curr = stack[0]
		prevVal = curr.Val
		stack = stack[1:]

		if curr != nil {
			curr = curr.Right
		}

	}
	return minDelta
}

func KthSmallest(root *TreeNode, k int) int {
	// in this BST find the kth smallest (1-indexed) node
	// feels pretty easy. Just do In order traversal and count pops
	stack := []*TreeNode{}
	curr := root
	count := 1
	for len(stack) > 0 || curr != nil {
		for curr != nil {
			stack = append([]*TreeNode{curr}, stack...)
			curr = curr.Left
		}
		curr = stack[0]
		stack = stack[1:]
		if count == k {
			return curr.Val
		}
		count++
		if curr != nil {
			curr = curr.Right
		}
	}
	return -1
}

func AverageOfLevels(root *TreeNode) []float32 {

	// so we want to be able to go through each level (BFS) and average the numbers
	// will use iteration so we will want a queue and then just add levels to the queue and then
	// average as we pop them
	queue := []*TreeNode{root}
	res := []float32{}
	for len(queue) > 0 {
		// everything in the queue is a level so pop it and its children
		l := len(queue)
		sum := float32(0.0)
		for range l {
			pop := queue[0]
			queue = queue[1:]
			if pop.Left != nil {
				queue = append(queue, pop.Left)
			}
			if pop.Right != nil {
				queue = append(queue, pop.Right)
			}
			sum += float32(pop.Val)
		}
		res = append(res, sum/float32(l))
		fmt.Println("queue ", queue)
	}
	return res
}

func IsValidBST(root *TreeNode) bool {
	// determine whether the tree is a valid BST
	// strictly less to the left and strictly greater to the right
	// thought is to implement in order traversal and then just check
	// whether they are in order i..e store prev int and check next int is bigger
	// Node vals can be negative so just need to account for that in the init of prev
	// we want to loop through the left pop at nil and then add right if its there
	stack := []*TreeNode{}
	curr := root
	prev := math.MinInt64
	for len(stack) > 0 || curr != nil {
		for curr != nil {
			stack = append([]*TreeNode{curr}, stack...)
			curr = curr.Left
		}
		curr = stack[0]

		if curr.Val <= prev {
			return false
		}
		prev = curr.Val
		stack = stack[1:]
		curr = curr.Right
	}
	return true
}

func RotateListRight(head *ListNode, k int) *ListNode {
	// so we want to find the n + 1 - kth node. Set the parent of that nodes Next to nil and then
	// go to then end of the list and set Next to the original head
	// We dont know how long the list is ahead of time so actually we are going to have to have
	// a buffer that stores the last k values so when we hit the end we can then reference that Node
	if k == 0 {
		return head
	}
	if head == nil || head.Next == nil {
		return head
	}

	mod := func(a int, b int) int {
		tmp := a % b
		if tmp < 0 {
			tmp += b
		}
		return tmp
	}

	curr := head
	kBuffer := make([]*ListNode, min(k+1, 500)) // need to account for very large ks as reqs state numNodes <= 500 but k can be 10^9
	index := 0
	count := 0
	for curr != nil {
		kBuffer[index] = curr
		if curr.Next == nil {
			if k > count {
				k = k % (count + 1)
				if k == 0 {
					return head
				}
				kBuffer = kBuffer[index-(k) : index+1]
				index = k
			}
			kth := mod(index+2, k+1)
			parentKth := mod(index+1, k+1)
			oldEnd := index

			kBuffer[parentKth].Next = nil // set parent of kth node to be the end of the linked list
			kBuffer[oldEnd].Next = head   // set the old end of list to point at the old head
			return kBuffer[kth]           // return the new head which is the old kth position which is at index - 1
		}
		index = (index + 1) % (k + 1)
		count++
		curr = curr.Next
	}
	return nil
}

func IsSymmetric(root *TreeNode) bool {
	// find whether the tree is symmetric or not
	// a tree is symmetric if when using BFS
	// that each level is smmetric i.e. [1], [2, 2], [3, 4, 4, 3] each array is reversible
	// [nil, 3, nil, 3] is not
	// inital thought is that iterative BFS and check  whether the level has
	// equality when reversing the slice
	// although it feels like there is probably a recursive way that would be cleaner

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		levelLength := len(queue)
		levelValues := []int{}
		for range levelLength {

			curr := queue[0]
			queue = queue[1:]
			if curr.Left != nil {
				queue = append(queue, curr.Left)
				levelValues = append(levelValues, curr.Left.Val)
			} else {
				levelValues = append(levelValues, -1000) // This only works because node vals are > -100
			}
			if curr.Right != nil {
				queue = append(queue, curr.Right)
				levelValues = append(levelValues, curr.Right.Val)
			} else {
				levelValues = append(levelValues, -1000) // This only works because node vals are > 0
			}
		}
		// now check that if i reverse the level vals they are the same
		for i := range len(levelValues) {
			if levelValues[i] != levelValues[len(levelValues)-(i+1)] {
				return false
			}
		}

	}
	return true
}

func IsSameTree(p *TreeNode, q *TreeNode) bool {

	// base case is when both are nil we return true
	// if then one is not nil and the other is return false
	// if then they are both not nil return if they are the same value and recurse
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	return p.Val == q.Val && IsSameTree(p.Left, q.Left) && IsSameTree(p.Right, q.Right)
}

func RightSideView(root *TreeNode) []int {
	// given a binary tree return the values of the nodes you would see if you were viewing
	// the tree from the rhs perspective
	// my first thought is that for each level of the tree it is the last value
	// so do a BFS and return the last value when ordering left then right nodes
	// will do it iteratively which means we need a quueue which we will
	// add to and then a loop that will assess the current level

	if root == nil {
		return []int{}
	}
	queue := []*TreeNode{root}
	l := len(queue)
	output := []int{}
	for l > 0 {
		rightVal := 0
		for _, node := range queue[:l] {
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			rightVal = node.Val
		}
		output = append(output, rightVal)
		queue = queue[l:]
		l = len(queue)

	}

	return output
}

func Flatten(root *TreeNode) {
	// this should take a root of a tree node and inplace turn it into
	// a faux linked list such that all of the left child nodes are nil
	// and the right child nodes are the pointer to the next in the list
	// the order should follow that of pre-order traversal
	// Solution: the crux of this will be to implement iterative DFS
	// in a pre-order way. It will amount to similar to in-order just
	// placing the items on the stack in the right order.
	if root == nil {
		return
	}
	stack := []*TreeNode{root}
	prev := root
	skip := true
	for len(stack) > 0 {
		curr := stack[0]
		stack = stack[1:]
		if !skip {
			prev.Left = nil
			prev.Right = curr
		} else {
			skip = false
		}

		prev = curr
		if curr.Right != nil {
			stack = append([]*TreeNode{curr.Right}, stack...)
		}
		if curr.Left != nil {
			stack = append([]*TreeNode{curr.Left}, stack...)
		}

	}

}

func buildTree(ts *TreeStack) (*TreeStack, *TreeStack) {

	if len(ts.inorder) < 2 {
		return &TreeStack{}, &TreeStack{}
	}

	rootVal := ts.preorder[0]
	var leftInorder, rightInorder []int
	leftValMap := make(map[int]bool)
	rightValMap := make(map[int]bool)

	// i need to split the inorder array around the rootVal as this will give me the left and right subtrees
	// this is O(n) where n is the number of nodes in the tree
	rootFlag := false
	for _, i := range ts.inorder {
		if !rootFlag && i != rootVal {
			leftInorder = append(leftInorder, i)
			leftValMap[i] = true
		} else {
			rootFlag = true
			if i != rootVal {
				rightInorder = append(rightInorder, i)
				rightValMap[i] = true
			}
		}
	}

	var leftPreorder, rightPreorder []int
	for _, i := range ts.preorder {
		if leftValMap[i] {
			leftPreorder = append(leftPreorder, i)
		} else if rightValMap[i] {
			rightPreorder = append(rightPreorder, i)
		}
	}

	leftNode := TreeNode{}
	if len(leftPreorder) > 0 {
		leftNode.Val = leftPreorder[0]
		ts.node.Left = &leftNode
	}
	rightNode := TreeNode{}
	if len(rightPreorder) > 0 {
		rightNode.Val = rightPreorder[0]
		ts.node.Right = &rightNode
	}

	leftTreeStack := TreeStack{&leftNode, leftPreorder, leftInorder}
	rightTreeStack := TreeStack{&rightNode, rightPreorder, rightInorder}

	return &leftTreeStack, &rightTreeStack

}

type TreeStack struct {
	node     *TreeNode
	preorder []int
	inorder  []int
}

func BuildTree(preorder []int, inorder []int) *TreeNode {
	// given the two arrays construct the given tree
	// initilally cant think why one would need both as they on their own define
	// a unique tree but maybe there is an efficieny im missing
	// ah yes! nulls dont appear in the orders and so im guessing having the two allows you to unpick that
	// will need to draw it out
	// unique values constraint feels like it will be pretty significant
	// probably that way i can then find the indexes in the two and work the tree struct from there

	// i need to split the inorder array around the rootVal as this will give me the left and right subtrees
	// this is O(n) where n is the number of nodes in the tree
	if len(preorder) == 0 {
		return &TreeNode{}
	}
	rootNode := TreeNode{Val: preorder[0]}
	baseStack := TreeStack{&rootNode, preorder, inorder}

	stack := []*TreeStack{&baseStack}

	for len(stack) > 0 {
		curr := stack[0]

		leftTreeStack, rightTreeStack := buildTree(curr)

		stack = stack[1:]
		if len(leftTreeStack.preorder) > 0 {
			stack = append(stack, leftTreeStack)
		}
		if len(rightTreeStack.preorder) > 0 {

			stack = append(stack, rightTreeStack)
		}
	}

	return &rootNode

}
