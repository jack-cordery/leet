package leet

import (
	"fmt"
	"log"
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
    Val int
    Next *ListNode
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
                return &ListNode{Val: carry, Next:lPrev}
            } else {
                return lPrev
            }
        }
        a = currentl1.Val
        b = currentl2.Val
        digit := ( a + b + carry) % 10
        carry = ( a + b + carry) / 10
        lPrev = &ListNode{Val: digit, Next: lPrev}
    }
}


func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    dummy, sum := new(ListNode), 0
    for curr := dummy; l1 != nil || l2 != nil || sum != 0 ;curr = curr.Next {
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
            left +=1 
        }
        result = max(result, right - left + 1)

        right +=1
    }
    return result
}

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    
    var ans float64
    l1, l2 := 0, 0
    r1, r2 := len(nums1) - 1, len(nums2) - 1
    for {
    if r1 - l1 == 0 && r2 - l2 == 0 {
        ans = float64(nums1[r1] + nums2[r2]) / 2 
        break
    }
    // Past left
    if r1 - l1 < 0 {
        if r2 - l2 < 2 {
            if r2 - l2 < 1 {
                ans = float64(nums2[r2])
            } else {
                ans = float64(nums2[r2] + nums2[l2]) / 2
            }
            break
        }
        l2 ++
        r2 --
        continue
    } 
    // Past right 
    if r2 - l2 < 0 {
        if r1 - l1 < 2 {
            if r1 - l1 < 1 {
                ans = float64(nums1[r1])
            } else {
                ans = float64(nums1[r1] + nums1[l1]) / 2
            }
            break
        }
        l1 ++
        r1 --
        continue
    }

    if nums1[l1] <= nums2[l2] {
        l1 ++
    } else {
        l2 ++
    }

    if nums1[r1] >= nums2[r2] {
        r1 --
    } else {
        r2 --
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
    for r:= len(arr);r > 0; r-- {
        OuterLoop:
            for i:= 0;  i <= len(arr) - r ; i++ {
                window := arr[i: i + r ]
                for j:=0; j <= len(window)/2 ; j++ {
                    if window[j] != window[len(window) - 1 - j] {
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
    for i:= 0; i < numRows ;  i++ {
        fmt.Println("i", i)
        for j:=0; float32(j) < float32(l - i) / float32((2 * (numRows - 1))); j++ {
            step := i + (2 * j * (numRows -  1))
            fmt.Println("j", j)
            o = append(o, b[step])
            if step + 2*(numRows - ( i  + 1)) < l  && i!= 0 && i!= numRows - 1 {
                o = append(o, b[step + 2*(numRows - ( i  + 1))])
            }
            fmt.Println("cond", float32(l - i) / float32((2 * (numRows - 1))))
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
    if float64(x) >= math.Pow(2,31)|| float64(x) < -math.Pow(2, 31) {
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
    for j:= len(xArray); j > 0; j-- {
        o = append(o, xArray[j - 1])
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


    for i:=0; i < bigBase ; i++ {
        // So this loop goes up to half way 
        lowerBase := int(math.Pow(10, float64(i + 1)))
        upperBase := int(math.Pow(10, float64(bigBase - i)))
        lDigit := (x / upperBase) % 10
        rDigit := (x % lowerBase ) / (lowerBase / 10)
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
    for windowSize:=len(height) -  1; windowSize > 0; windowSize-- {
        for i:=0; i <= (len(height) - 1) - windowSize; i++{
            j:= i + windowSize
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
    i, j := 0, len(height) - 1
    for i < j{
        maxArea = max(maxArea, (j- i) * min(height[i], height[j]))
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
              digitString = romans[base / 10] + romans[5 * base / 10] 
        case 5:
            digitString = romans[5 * base / 10]
        case 6:
            digitString = romans[5 * base / 10] + romans[base / 10]
        case 7:
            digitString = romans[5 * base / 10] + romans[base / 10] + romans[base / 10]
        case 8:
            digitString = romans[5 * base / 10] + romans[base / 10] + romans[base / 10] + romans[base / 10]
        case 9:
            digitString = romans[base / 10] + romans[base] 
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
        chars :=  string(rommanArray[l - 2]) + string(rommanArray[l -1])
        value := romans[string(chars)]
        if value == 0 {
            // its invalid therefore its a single value 
             total += romans[string(rommanArray[l -1])]
             l --
        } else  {
            total += value
            l -= 2
        }
        fmt.Println("LoopE",l, chars, value, total)
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

    for i:=0; i < len(s); i++ {
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
            if state[len(state) - 1] != bracketMap[s[i]] { // not sure if s[i] is the byte representation
                // not closed corectectly so return false
                return false
            }
            state = state[:len(state) - 1] 
        }
    }
    return len(state) == 0 
}
