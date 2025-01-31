/*

This implementation includes four classic Sliding Window pattern problems:

Maximum Sum Subarray of size K - Fixed-size window
Longest Substring with K Distinct Characters - Variable-size window with frequency map
Minimum Size Subarray Sum - Variable-size window with target sum
Find All Anagrams - Fixed-size window with character frequency matching

Key features:

Both fixed and variable size window implementations
Efficient frequency map usage
Comprehensive test cases
Clear complexity analysis
Clean Go idioms and error handling

*/

// File: slidingwindow/slidingwindow.go

package slidingwindow

// MaxSubArraySum finds maximum sum subarray of fixed size k
// Time Complexity: O(n)
// Space Complexity: O(1)
func MaxSubArraySum(nums []int, k int) int {
    if len(nums) < k {
        return 0
    }

    // Calculate sum of first window
    windowSum := 0
    for i := 0; i < k; i++ {
        windowSum += nums[i]
    }

    maxSum := windowSum

    // Slide window and track maximum sum
    for i := k; i < len(nums); i++ {
        windowSum = windowSum - nums[i-k] + nums[i]
        if windowSum > maxSum {
            maxSum = windowSum
        }
    }

    return maxSum
}

// LongestSubstringKDistinct finds the longest substring with at most k distinct characters
// Time Complexity: O(n) where n is the length of the string
// Space Complexity: O(k) where k is the number of distinct characters
func LongestSubstringKDistinct(s string, k int) int {
    if k == 0 || len(s) == 0 {
        return 0
    }

    charFreq := make(map[byte]int)
    maxLength := 0
    windowStart := 0

    for windowEnd := 0; windowEnd < len(s); windowEnd++ {
        // Add current character to frequency map
        charFreq[s[windowEnd]]++

        // Shrink window while we have more than k distinct characters
        for len(charFreq) > k {
            charFreq[s[windowStart]]--
            if charFreq[s[windowStart]] == 0 {
                delete(charFreq, s[windowStart])
            }
            windowStart++
        }

        // Update maximum length
        currentLength := windowEnd - windowStart + 1
        if currentLength > maxLength {
            maxLength = currentLength
        }
    }

    return maxLength
}

// MinSubArraySum finds the minimum length subarray with sum greater than or equal to target
// Time Complexity: O(n)
// Space Complexity: O(1)
func MinSubArraySum(target int, nums []int) int {
    minLength := len(nums) + 1
    windowSum := 0
    windowStart := 0

    for windowEnd := 0; windowEnd < len(nums); windowEnd++ {
        windowSum += nums[windowEnd]

        for windowSum >= target {
            currentLength := windowEnd - windowStart + 1
            if currentLength < minLength {
                minLength = currentLength
            }
            windowSum -= nums[windowStart]
            windowStart++
        }
    }

    if minLength > len(nums) {
        return 0 // No subarray found
    }
    return minLength
}

// FindAnagrams finds all start indices of anagrams of pattern in string s
// Time Complexity: O(n) where n is the length of string s
// Space Complexity: O(k) where k is the size of character set
func FindAnagrams(s string, p string) []int {
    result := []int{}
    if len(s) < len(p) {
        return result
    }

    // Create frequency maps
    pCount := make([]int, 26)
    windowCount := make([]int, 26)

    // Initialize pattern frequency
    for i := 0; i < len(p); i++ {
        pCount[p[i]-'a']++
    }

    // Initialize first window
    for i := 0; i < len(p); i++ {
        windowCount[s[i]-'a']++
    }

    // Check first window
    if isAnagram(windowCount, pCount) {
        result = append(result, 0)
    }

    // Slide window
    for i := len(p); i < len(s); i++ {
        // Remove leftmost character
        windowCount[s[i-len(p)]-'a']--
        // Add new character
        windowCount[s[i]-'a']++
        
        if isAnagram(windowCount, pCount) {
            result = append(result, i-len(p)+1)
        }
    }

    return result
}

// isAnagram helper function to compare character frequencies
func isAnagram(count1, count2 []int) bool {
    for i := 0; i < 26; i++ {
        if count1[i] != count2[i] {
            return false
        }
    }
    return true
}

// File: slidingwindow/slidingwindow_test.go

package slidingwindow

import (
    "reflect"
    "testing"
)

func TestMaxSubArraySum(t *testing.T) {
    tests := []struct {
        nums     []int
        k        int
        expected int
    }{
        {[]int{1, 4, 2, 10, 23, 3, 1, 0, 20}, 4, 39},
        {[]int{100, 200, 300, 400}, 2, 700},
        {[]int{1, 2}, 3, 0},
        {[]int{1, 1, 1, 1, 1}, 2, 2},
    }

    for _, test := range tests {
        result := MaxSubArraySum(test.nums, test.k)
        if result != test.expected {
            t.Errorf("For nums=%v, k=%d, expected %d but got %d",
                test.nums, test.k, test.expected, result)
        }
    }
}

func TestLongestSubstringKDistinct(t *testing.T) {
    tests := []struct {
        s        string
        k        int
        expected int
    }{
        {"eceba", 2, 3},
        {"aa", 1, 2},
        {"abcabcabc", 2, 2},
        {"", 2, 0},
        {"a", 0, 0},
    }

    for _, test := range tests {
        result := LongestSubstringKDistinct(test.s, test.k)
        if result != test.expected {
            t.Errorf("For s=%s, k=%d, expected %d but got %d",
                test.s, test.k, test.expected, result)
        }
    }
}

func TestMinSubArraySum(t *testing.T) {
    tests := []struct {
        nums     []int
        target   int
        expected int
    }{
        {[]int{2, 3, 1, 2, 4, 3}, 7, 2},
        {[]int{1, 4, 4}, 4, 1},
        {[]int{1, 1, 1, 1, 1, 1, 1, 1}, 11, 0},
        {[]int{1, 2, 3, 4, 5}, 15, 5},
    }

    for _, test := range tests {
        result := MinSubArraySum(test.target, test.nums)
        if result != test.expected {
            t.Errorf("For nums=%v, target=%d, expected %d but got %d",
                test.nums, test.target, test.expected, result)
        }
    }
}

func TestFindAnagrams(t *testing.T) {
    tests := []struct {
        s        string
        p        string
        expected []int
    }{
        {"cbaebabacd", "abc", []int{0, 6}},
        {"abab", "ab", []int{0, 1, 2}},
        {"aa", "bb", []int{}},
        {"", "a", []int{}},
    }

    for _, test := range tests {
        result := FindAnagrams(test.s, test.p)
        if !reflect.DeepEqual(result, test.expected) {
            t.Errorf("For s=%s, p=%s, expected %v but got %v",
                test.s, test.p, test.expected, result)
        }
    }
}