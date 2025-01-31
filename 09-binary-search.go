/*

This implementation includes six classic Modified Binary Search problems:

Standard Binary Search - Basic binary search in sorted array
Search in Rotated Array - Binary search in rotated sorted array
Find Minimum in Rotated Array - Find minimum element in rotated sorted array
Search Range - Find first and last positions of target
Find Peak Element - Find peak element in array
Search in Bitonic Array - Search in array that increases then decreases

Key features:

Multiple binary search variations
Handling of rotated arrays
Finding boundaries and peaks
Search in different array patterns
Comprehensive test coverage
Clean Go idioms

*/

// File: binarysearch/binarysearch.go

package binarysearch

// Search finds target in sorted array
// Returns index of target, or -1 if not found
// Time Complexity: O(log n)
// Space Complexity: O(1)
func Search(nums []int, target int) int {
    left, right := 0, len(nums)-1
    
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] == target {
            return mid
        }
        if nums[mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    
    return -1
}

// SearchRotated searches target in rotated sorted array
// Time Complexity: O(log n)
// Space Complexity: O(1)
func SearchRotated(nums []int, target int) int {
    if len(nums) == 0 {
        return -1
    }
    
    left, right := 0, len(nums)-1
    
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] == target {
            return mid
        }
        
        // Check which half is sorted
        if nums[left] <= nums[mid] {
            // Left half is sorted
            if target >= nums[left] && target < nums[mid] {
                right = mid - 1
            } else {
                left = mid + 1
            }
        } else {
            // Right half is sorted
            if target > nums[mid] && target <= nums[right] {
                left = mid + 1
            } else {
                right = mid - 1
            }
        }
    }
    
    return -1
}

// FindMinRotated finds minimum element in rotated sorted array
// Time Complexity: O(log n)
// Space Complexity: O(1)
func FindMinRotated(nums []int) int {
    if len(nums) == 0 {
        return -1
    }
    
    left, right := 0, len(nums)-1
    
    for left < right {
        mid := left + (right-left)/2
        
        if nums[mid] > nums[right] {
            left = mid + 1
        } else {
            right = mid
        }
    }
    
    return nums[left]
}

// SearchRange finds first and last position of target
// Time Complexity: O(log n)
// Space Complexity: O(1)
func SearchRange(nums []int, target int) []int {
    result := []int{-1, -1}
    if len(nums) == 0 {
        return result
    }
    
    // Find first occurrence
    result[0] = findBound(nums, target, true)
    if result[0] == -1 {
        return result
    }
    
    // Find last occurrence
    result[1] = findBound(nums, target, false)
    return result
}

// findBound is a helper function to find first or last occurrence
func findBound(nums []int, target int, isFirst bool) int {
    left, right := 0, len(nums)-1
    
    for left <= right {
        mid := left + (right-left)/2
        
        if nums[mid] == target {
            if isFirst {
                if mid == 0 || nums[mid-1] != target {
                    return mid
                }
                right = mid - 1
            } else {
                if mid == len(nums)-1 || nums[mid+1] != target {
                    return mid
                }
                left = mid + 1
            }
        } else if nums[mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    
    return -1
}

// PeakElement finds a peak element in array
// (element greater than neighbors)
// Time Complexity: O(log n)
// Space Complexity: O(1)
func PeakElement(nums []int) int {
    left, right := 0, len(nums)-1
    
    for left < right {
        mid := left + (right-left)/2
        
        if nums[mid] > nums[mid+1] {
            right = mid
        } else {
            left = mid + 1
        }
    }
    
    return left
}

// SearchBitonicArray searches in bitonic array
// (increasing then decreasing)
// Time Complexity: O(log n)
// Space Complexity: O(1)
func SearchBitonicArray(nums []int, target int) int {
    // Find peak element
    peak := findPeak(nums)
    
    // Search in increasing part
    result := binarySearch(nums, target, 0, peak, true)
    if result != -1 {
        return result
    }
    
    // Search in decreasing part
    return binarySearch(nums, target, peak+1, len(nums)-1, false)
}

// findPeak finds peak in bitonic array
func findPeak(nums []int) int {
    left, right := 0, len(nums)-1
    
    for left < right {
        mid := left + (right-left)/2
        if nums[mid] > nums[mid+1] {
            right = mid
        } else {
            left = mid + 1
        }
    }
    
    return left
}

// binarySearch performs binary search in specified order
func binarySearch(nums []int, target int, left, right int, ascending bool) int {
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] == target {
            return mid
        }
        
        if ascending {
            if nums[mid] < target {
                left = mid + 1
            } else {
                right = mid - 1
            }
        } else {
            if nums[mid] < target {
                right = mid - 1
            } else {
                left = mid + 1
            }
        }
    }
    
    return -1
}

// File: binarysearch/binarysearch_test.go

package binarysearch

import (
    "reflect"
    "testing"
)

func TestSearch(t *testing.T) {
    tests := []struct {
        nums     []int
        target   int
        expected int
    }{
        {[]int{1, 3, 5, 6}, 5, 2},
        {[]int{1, 3, 5, 6}, 2, -1},
        {[]int{}, 5, -1},
    }
    
    for _, test := range tests {
        result := Search(test.nums, test.target)
        if result != test.expected {
            t.Errorf("For nums=%v, target=%d, expected %d but got %d",
                test.nums, test.target, test.expected, result)
        }
    }
}

func TestSearchRotated(t *testing.T) {
    tests := []struct {
        nums     []int
        target   int
        expected int
    }{
        {[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
        {[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
        {[]int{1}, 0, -1},
    }
    
    for _, test := range tests {
        result := SearchRotated(test.nums, test.target)
        if result != test.expected {
            t.Errorf("For nums=%v, target=%d, expected %d but got %d",
                test.nums, test.target, test.expected, result)
        }
    }
}

func TestFindMinRotated(t *testing.T) {
    tests := []struct {
        nums     []int
        expected int
    }{
        {[]int{3, 4, 5, 1, 2}, 1},
        {[]int{4, 5, 6, 7, 0, 1, 2}, 0},
        {[]int{11, 13, 15, 17}, 11},
    }
    
    for _, test := range tests {
        result := FindMinRotated(test.nums)
        if result != test.expected {
            t.Errorf("For nums=%v, expected %d but got %d",
                test.nums, test.expected, result)
        }
    }
}

func TestSearchRange(t *testing.T) {
    tests := []struct {
        nums     []int
        target   int
        expected []int
    }{
        {[]int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
        {[]int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1}},
        {[]int{}, 0, []int{-1, -1}},
    }
    
    for _, test := range tests {
        result := SearchRange(test.nums, test.target)
        if !reflect.DeepEqual(result, test.expected) {
            t.Errorf("For nums=%v, target=%d, expected %v but got %v",
                test.nums, test.target, test.expected, result)
        }
    }
}

func TestPeakElement(t *testing.T) {
    tests := []struct {
        nums     []int
        expected int
    }{
        {[]int{1, 2, 3, 1}, 2},
        {[]int{1, 2, 1, 3, 5, 6, 4}, 6},
    }
    
    for _, test := range tests {
        result := PeakElement(test.nums)
        if nums[result] < nums[result-1] || nums[result] < nums[result+1] {
            t.Errorf("For nums=%v, got invalid peak at index %d",
                test.nums, result)
        }
    }
}

func TestSearchBitonicArray(t *testing.T) {
    tests := []struct {
        nums     []int
        target   int
        expected int
    }{
        {[]int{1, 3, 8, 12, 4, 2}, 4, 4},
        {[]int{1, 3, 8, 12, 4, 2}, 10, -1},
        {[]int{1, 3, 8, 12}, 8, 2},
    }
    
    for _, test := range tests {
        result := SearchBitonicArray(test.nums, test.target)
        if result != test.expected {
            t.Errorf("For nums=%v, target=%d, expected %d but got %d",
                test.nums, test.target, test.expected, result)
        }
    }
}