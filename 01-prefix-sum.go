/*
This implementation includes three common prefix sum pattern problems:

Range Sum Query - Using prefix sums for O(1) range queries
Subarray Sum Equals K - Finding continuous subarrays that sum to k
Product of Array Except Self - Using prefix/suffix products

Key features:

Complete test coverage
Time and space complexity analysis
Clean Go idioms and error handling
Efficient memory usage
*/

// File: prefixsum/prefixsum.go

package prefixsum

// NumArray represents a data structure optimized for range sum queries
type NumArray struct {
    prefixSum []int
}

// NewNumArray initializes NumArray with a prefixSum array
// Time Complexity: O(n)
// Space Complexity: O(n)
func NewNumArray(nums []int) *NumArray {
    prefixSum := make([]int, len(nums)+1)
    for i := range nums {
        prefixSum[i+1] = prefixSum[i] + nums[i]
    }
    return &NumArray{prefixSum: prefixSum}
}

// RangeSum returns sum of elements between indices [left, right] inclusive
// Time Complexity: O(1)
// Space Complexity: O(1)
func (na *NumArray) RangeSum(left, right int) int {
    return na.prefixSum[right+1] - na.prefixSum[left]
}

// SubarraySum finds the number of continuous subarrays that sum to k
// Time Complexity: O(n)
// Space Complexity: O(n)
func SubarraySum(nums []int, k int) int {
    count := 0
    sum := 0
    sumCount := map[int]int{0: 1} // Initialize with 0 sum seen once

    for _, num := range nums {
        sum += num
        if prevCount, exists := sumCount[sum-k]; exists {
            count += prevCount
        }
        sumCount[sum]++
    }
    return count
}

// ProductArray returns an array where each element is the product of all
// other elements except itself without using division
// Time Complexity: O(n)
// Space Complexity: O(1) excluding the output array
func ProductArray(nums []int) []int {
    n := len(nums)
    if n == 0 {
        return []int{}
    }

    result := make([]int, n)
    result[0] = 1

    // Calculate prefix products
    for i := 1; i < n; i++ {
        result[i] = result[i-1] * nums[i-1]
    }

    // Calculate suffix products and combine
    suffixProduct := 1
    for i := n - 1; i >= 0; i-- {
        result[i] *= suffixProduct
        suffixProduct *= nums[i]
    }

    return result
}

// File: prefixsum/prefixsum_test.go

package prefixsum

import (
    "reflect"
    "testing"
)

func TestRangeSum(t *testing.T) {
    tests := []struct {
        nums          []int
        left, right   int
        expectedSum   int
    }{
        {[]int{1, 2, 3, 4, 5}, 1, 3, 9},
        {[]int{-2, 0, 3, -5, 2, -1}, 0, 2, 1},
        {[]int{1}, 0, 0, 1},
        {[]int{1, 2, 3}, 0, 2, 6},
    }

    for _, test := range tests {
        numArray := NewNumArray(test.nums)
        result := numArray.RangeSum(test.left, test.right)
        if result != test.expectedSum {
            t.Errorf("For nums=%v, left=%d, right=%d, expected %d but got %d",
                test.nums, test.left, test.right, test.expectedSum, result)
        }
    }
}

func TestSubarraySum(t *testing.T) {
    tests := []struct {
        nums     []int
        k        int
        expected int
    }{
        {[]int{1, 1, 1}, 2, 2},
        {[]int{1, 2, 3}, 3, 2},
        {[]int{1}, 1, 1},
        {[]int{1, -1, 0}, 0, 3},
    }

    for _, test := range tests {
        result := SubarraySum(test.nums, test.k)
        if result != test.expected {
            t.Errorf("For nums=%v, k=%d, expected %d but got %d",
                test.nums, test.k, test.expected, result)
        }
    }
}

func TestProductArray(t *testing.T) {
    tests := []struct {
        nums     []int
        expected []int
    }{
        {[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
        {[]int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},
        {[]int{1}, []int{1}},
        {[]int{}, []int{}},
    }

    for _, test := range tests {
        result := ProductArray(test.nums)
        if !reflect.DeepEqual(result, test.expected) {
            t.Errorf("For nums=%v, expected %v but got %v",
                test.nums, test.expected, result)
        }
    }
}