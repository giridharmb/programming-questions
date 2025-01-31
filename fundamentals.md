# Fundamental LeetCode Problems in Go

This repository contains implementations of fundamental LeetCode problems in Go, along with detailed explanations and multiple approaches where applicable.

## Table of Contents
1. [Two Sum](#two-sum)
2. [Valid Parentheses](#valid-parentheses)
3. [Reverse Linked List](#reverse-linked-list)
4. [Maximum Subarray](#maximum-subarray)
5. [Binary Search](#binary-search)

## Two Sum
**Problem**: Given an array of integers `nums` and an integer `target`, return indices of the two numbers in the array that add up to the target.

### Approach 1: Brute Force
- Time Complexity: O(n²)
- Space Complexity: O(1)

```go
func twoSum(nums []int, target int) []int {
    for i := 0; i < len(nums); i++ {
        for j := i + 1; j < len(nums); j++ {
            if nums[i] + nums[j] == target {
                return []int{i, j}
            }
        }
    }
    return nil
}
```

### Approach 2: Hash Map
- Time Complexity: O(n)
- Space Complexity: O(n)

```go
func twoSum(nums []int, target int) []int {
    seen := make(map[int]int)
    
    for i, num := range nums {
        complement := target - num
        if j, exists := seen[complement]; exists {
            return []int{j, i}
        }
        seen[num] = i
    }
    return nil
}
```

## Valid Parentheses
**Problem**: Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

### Solution using Stack
- Time Complexity: O(n)
- Space Complexity: O(n)

```go
func isValid(s string) bool {
    stack := make([]rune, 0)
    pairs := map[rune]rune{
        ')': '(',
        '}': '{',
        ']': '[',
    }
    
    for _, char := range s {
        if char == '(' || char == '{' || char == '[' {
            stack = append(stack, char)
        } else {
            if len(stack) == 0 {
                return false
            }
            if stack[len(stack)-1] != pairs[char] {
                return false
            }
            stack = stack[:len(stack)-1]
        }
    }
    
    return len(stack) == 0
}
```

## Reverse Linked List
**Problem**: Given the head of a singly linked list, reverse the list, and return the reversed list.

### LinkedList Definition
```go
type ListNode struct {
    Val int
    Next *ListNode
}
```

### Approach 1: Iterative
- Time Complexity: O(n)
- Space Complexity: O(1)

```go
func reverseList(head *ListNode) *ListNode {
    var prev *ListNode
    current := head
    
    for current != nil {
        nextTemp := current.Next
        current.Next = prev
        prev = current
        current = nextTemp
    }
    
    return prev
}
```

### Approach 2: Recursive
- Time Complexity: O(n)
- Space Complexity: O(n) due to recursion stack

```go
func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    
    reversedRest := reverseList(head.Next)
    head.Next.Next = head
    head.Next = nil
    
    return reversedRest
}
```

## Maximum Subarray
**Problem**: Given an integer array nums, find the contiguous subarray with the largest sum and return its sum.

### Approach 1: Kadane's Algorithm
- Time Complexity: O(n)
- Space Complexity: O(1)

```go
func maxSubArray(nums []int) int {
    maxSum := nums[0]
    currentSum := nums[0]
    
    for i := 1; i < len(nums); i++ {
        currentSum = max(nums[i], currentSum + nums[i])
        maxSum = max(maxSum, currentSum)
    }
    
    return maxSum
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```

### Approach 2: Divide and Conquer
- Time Complexity: O(n log n)
- Space Complexity: O(log n) due to recursion

```go
func maxSubArray(nums []int) int {
    return maxSubArrayDivideConquer(nums, 0, len(nums)-1)
}

func maxSubArrayDivideConquer(nums []int, left, right int) int {
    if left == right {
        return nums[left]
    }
    
    mid := (left + right) / 2
    
    leftSum := maxSubArrayDivideConquer(nums, left, mid)
    rightSum := maxSubArrayDivideConquer(nums, mid+1, right)
    crossSum := maxCrossingSum(nums, left, mid, right)
    
    return max(max(leftSum, rightSum), crossSum)
}

func maxCrossingSum(nums []int, left, mid, right int) int {
    leftSum := math.MinInt32
    sum := 0
    
    for i := mid; i >= left; i-- {
        sum += nums[i]
        if sum > leftSum {
            leftSum = sum
        }
    }
    
    rightSum := math.MinInt32
    sum = 0
    
    for i := mid+1; i <= right; i++ {
        sum += nums[i]
        if sum > rightSum {
            rightSum = sum
        }
    }
    
    return leftSum + rightSum
}
```

## Binary Search
**Problem**: Given an array of integers nums which is sorted in ascending order, and an integer target, write a function to search target in nums. If target exists, then return its index. Otherwise, return -1.

### Approach 1: Classical Binary Search
- Time Complexity: O(log n)
- Space Complexity: O(1)

```go
func search(nums []int, target int) int {
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
```

### Approach 2: Recursive Binary Search
- Time Complexity: O(log n)
- Space Complexity: O(log n) due to recursion

```go
func search(nums []int, target int) int {
    return binarySearch(nums, target, 0, len(nums)-1)
}

func binarySearch(nums []int, target, left, right int) int {
    if left > right {
        return -1
    }
    
    mid := left + (right-left)/2
    
    if nums[mid] == target {
        return mid
    }
    
    if nums[mid] < target {
        return binarySearch(nums, target, mid+1, right)
    }
    
    return binarySearch(nums, target, left, mid-1)
}
```

## Testing the Solutions

Each solution can be tested using Go's testing framework. Here's an example test file structure:

```go
package leetcode

import (
    "testing"
)

func TestTwoSum(t *testing.T) {
    tests := []struct {
        nums     []int
        target   int
        expected []int
    }{
        {[]int{2, 7, 11, 15}, 9, []int{0, 1}},
        {[]int{3, 2, 4}, 6, []int{1, 2}},
        {[]int{3, 3}, 6, []int{0, 1}},
    }

    for _, test := range tests {
        result := twoSum(test.nums, test.target)
        if !equalSlices(result, test.expected) {
            t.Errorf("For nums=%v and target=%d, expected %v but got %v",
                test.nums, test.target, test.expected, result)
        }
    }
}

func equalSlices(a, b []int) bool {
    if len(a) != len(b) {
        return false
    }
    for i := range a {
        if a[i] != b[i] {
            return false
        }
    }
    return true
}
```

## Time and Space Complexity Guide

For each problem, consider these aspects when choosing an approach:

1. **Time Complexity**: 
   - O(1): Constant time
   - O(log n): Logarithmic time (Binary Search)
   - O(n): Linear time (Single pass)
   - O(n log n): Log-linear time (Efficient sorting)
   - O(n²): Quadratic time (Nested loops)

2. **Space Complexity**:
   - O(1): Constant space
   - O(n): Linear space (Additional data structures)
   - O(log n): Logarithmic space (Recursive call stack in balanced trees)

Choose the appropriate approach based on:
- Input size constraints
- Time requirements
- Space requirements
- Code readability and maintainability

## Contributing

Feel free to contribute additional problems, alternative solutions, or improvements to existing solutions. Please ensure:

1. All code is properly formatted using `go fmt`
2. Include comprehensive test cases
3. Document time and space complexity
4. Provide clear explanations of the approach