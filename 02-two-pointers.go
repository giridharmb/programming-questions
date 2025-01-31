/*

This implementation includes four classic Two Pointer pattern problems:

Two Sum - Finding two numbers that sum to target in sorted array
Three Sum - Finding all unique triplets that sum to zero
Remove Duplicates - Removing duplicates from sorted array in-place
Container With Most Water - Finding maximum area between vertical lines

Key features:

Comprehensive test coverage
Time/space complexity analysis for each function
Efficient in-place operations
Handle edge cases and duplicates
Clean Go style with helper functions

*/

// File: twopointers/twopointers.go

package twopointers

import "sort"

// TwoSum finds two numbers in a sorted array that sum up to target
// Returns indices (1-based) of the two numbers
// Time Complexity: O(n)
// Space Complexity: O(1)
func TwoSum(numbers []int, target int) []int {
    left, right := 0, len(numbers)-1
    
    for left < right {
        currentSum := numbers[left] + numbers[right]
        if currentSum == target {
            return []int{left + 1, right + 1}
        }
        if currentSum < target {
            left++
        } else {
            right--
        }
    }
    return []int{}
}

// ThreeSum finds all unique triplets that sum to zero
// Time Complexity: O(n²)
// Space Complexity: O(1) excluding the output array
func ThreeSum(nums []int) [][]int {
    sort.Ints(nums)
    result := make([][]int, 0)
    
    for i := 0; i < len(nums)-2; i++ {
        // Skip duplicates for i
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        
        left, right := i+1, len(nums)-1
        target := -nums[i]
        
        for left < right {
            sum := nums[left] + nums[right]
            
            if sum == target {
                result = append(result, []int{nums[i], nums[left], nums[right]})
                // Skip duplicates for left
                for left < right && nums[left] == nums[left+1] {
                    left++
                }
                // Skip duplicates for right
                for left < right && nums[right] == nums[right-1] {
                    right--
                }
                left++
                right--
            } else if sum < target {
                left++
            } else {
                right--
            }
        }
    }
    return result
}

// RemoveDuplicates removes duplicates in-place from sorted array
// Returns length of array after removing duplicates
// Time Complexity: O(n)
// Space Complexity: O(1)
func RemoveDuplicates(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    
    writePointer := 1
    for readPointer := 1; readPointer < len(nums); readPointer++ {
        if nums[readPointer] != nums[readPointer-1] {
            nums[writePointer] = nums[readPointer]
            writePointer++
        }
    }
    return writePointer
}

// ContainerWithMostWater finds maximum area of water that can be contained
// Time Complexity: O(n)
// Space Complexity: O(1)
func ContainerWithMostWater(height []int) int {
    maxArea := 0
    left, right := 0, len(height)-1
    
    for left < right {
        width := right - left
        h := min(height[left], height[right])
        area := width * h
        maxArea = max(maxArea, area)
        
        if height[left] < height[right] {
            left++
        } else {
            right--
        }
    }
    return maxArea
}

// min returns the minimum of two integers
func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

// max returns the maximum of two integers
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

// File: twopointers/twopointers_test.go

package twopointers

import (
    "reflect"
    "testing"
)

func TestTwoSum(t *testing.T) {
    tests := []struct {
        numbers  []int
        target   int
        expected []int
    }{
        {[]int{2, 7, 11, 15}, 9, []int{1, 2}},
        {[]int{2, 3, 4}, 6, []int{1, 3}},
        {[]int{-1, 0}, -1, []int{1, 2}},
        {[]int{1, 2, 3, 4, 5}, 10, []int{}},
    }
    
    for _, test := range tests {
        result := TwoSum(test.numbers, test.target)
        if !reflect.DeepEqual(result, test.expected) {
            t.Errorf("For numbers=%v, target=%d, expected %v but got %v",
                test.numbers, test.target, test.expected, result)
        }
    }
}

func TestThreeSum(t *testing.T) {
    tests := []struct {
        nums     []int
        expected [][]int
    }{
        {
            []int{-1, 0, 1, 2, -1, -4},
            [][]int{{-1, -1, 2}, {-1, 0, 1}},
        },
        {[]int{0, 0, 0}, [][]int{{0, 0, 0}}},
        {[]int{1, 2, 3}, [][]int{}},
    }
    
    for _, test := range tests {
        result := ThreeSum(test.nums)
        if !reflect.DeepEqual(result, test.expected) {
            t.Errorf("For nums=%v, expected %v but got %v",
                test.nums, test.expected, result)
        }
    }
}

func TestRemoveDuplicates(t *testing.T) {
    tests := []struct {
        nums            []int
        expectedLen     int
        expectedNums    []int
    }{
        {[]int{1, 1, 2}, 2, []int{1, 2}},
        {[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5, []int{0, 1, 2, 3, 4}},
        {[]int{1}, 1, []int{1}},
        {[]int{}, 0, []int{}},
    }
    
    for _, test := range tests {
        nums := make([]int, len(test.nums))
        copy(nums, test.nums)
        result := RemoveDuplicates(nums)
        
        if result != test.expectedLen {
            t.Errorf("For nums=%v, expected length %d but got %d",
                test.nums, test.expectedLen, result)
        }
        
        if !reflect.DeepEqual(nums[:result], test.expectedNums) {
            t.Errorf("For nums=%v, expected first %d elements to be %v but got %v",
                test.nums, test.expectedLen, test.expectedNums, nums[:result])
        }
    }
}

func TestContainerWithMostWater(t *testing.T) {
    tests := []struct {
        height   []int
        expected int
    }{
        {[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
        {[]int{1, 1}, 1},
        {[]int{4, 3, 2, 1, 4}, 16},
        {[]int{1, 2, 1}, 2},
    }
    
    for _, test := range tests {
        result := ContainerWithMostWater(test.height)
        if result != test.expected {
            t.Errorf("For height=%v, expected %d but got %d",
                test.height, test.expected, result)
        }
    }
}