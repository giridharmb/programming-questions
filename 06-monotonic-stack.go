/*

This implementation includes five classic Monotonic Stack problems:

Next Greater Element - Find next greater element for each array element
Daily Temperatures - Find days until warmer temperature
Largest Rectangle in Histogram - Find largest rectangular area in histogram
Maximal Rectangle - Find largest rectangle in binary matrix
Online Stock Span - Calculate stock price spans in real-time

Key features:

Efficient stack operations
Both array and matrix problems
Online/streaming data handling
Comprehensive test cases
Time/space complexity analysis
Clean Go idioms

*/

// File: monotonicstack/monotonic.go

package monotonicstack

// NextGreaterElement finds the next greater element for each array element
// Time Complexity: O(n)
// Space Complexity: O(n)
func NextGreaterElement(nums []int) []int {
    n := len(nums)
    result := make([]int, n)
    stack := make([]int, 0) // Stack stores indices
    
    // Initialize result with -1
    for i := range result {
        result[i] = -1
    }
    
    // Process all elements
    for i := 0; i < n*2; i++ {
        // Use modulo to handle circular array if needed
        idx := i % n
        
        // Pop elements from stack while current element is greater
        for len(stack) > 0 && nums[idx] > nums[stack[len(stack)-1]] {
            result[stack[len(stack)-1]] = nums[idx]
            stack = stack[:len(stack)-1]
        }
        
        if i < n {
            stack = append(stack, idx)
        }
    }
    
    return result
}

// DailyTemperatures finds number of days to wait for a warmer temperature
// Time Complexity: O(n)
// Space Complexity: O(n)
func DailyTemperatures(temperatures []int) []int {
    n := len(temperatures)
    result := make([]int, n)
    stack := make([]int, 0) // Stack stores indices
    
    for i := 0; i < n; i++ {
        // Pop elements while current temperature is higher
        for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
            prevDay := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            result[prevDay] = i - prevDay
        }
        stack = append(stack, i)
    }
    
    return result
}

// LargestRectangleArea finds the largest rectangle area in histogram
// Time Complexity: O(n)
// Space Complexity: O(n)
func LargestRectangleArea(heights []int) int {
    n := len(heights)
    if n == 0 {
        return 0
    }
    
    // Add 0 at the end to handle remaining elements in stack
    heights = append(heights, 0)
    stack := make([]int, 0) // Stack stores indices
    maxArea := 0
    
    for i := 0; i <= n; i++ {
        // Pop while current height is smaller than height at stack top
        for len(stack) > 0 && heights[i] < heights[stack[len(stack)-1]] {
            height := heights[stack[len(stack)-1]]
            stack = stack[:len(stack)-1]
            
            width := i
            if len(stack) > 0 {
                width = i - stack[len(stack)-1] - 1
            }
            
            area := height * width
            if area > maxArea {
                maxArea = area
            }
        }
        stack = append(stack, i)
    }
    
    return maxArea
}

// MaximalRectangle finds the largest rectangle containing only 1's in binary matrix
// Time Complexity: O(rows * cols)
// Space Complexity: O(cols)
func MaximalRectangle(matrix [][]byte) int {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return 0
    }
    
    rows, cols := len(matrix), len(matrix[0])
    heights := make([]int, cols)
    maxArea := 0
    
    // Process row by row
    for i := 0; i < rows; i++ {
        // Update heights
        for j := 0; j < cols; j++ {
            if matrix[i][j] == '1' {
                heights[j]++
            } else {
                heights[j] = 0
            }
        }
        
        // Find largest rectangle for current row
        area := LargestRectangleArea(heights)
        if area > maxArea {
            maxArea = area
        }
    }
    
    return maxArea
}

// StockSpanner implements an online stock span calculator
type StockSpanner struct {
    prices []int
    spans  []int
}

// NewStockSpanner initializes StockSpanner
func NewStockSpanner() StockSpanner {
    return StockSpanner{
        prices: make([]int, 0),
        spans:  make([]int, 0),
    }
}

// Next calculates span of stock price today
// Time Complexity: O(n) amortized
// Space Complexity: O(n)
func (ss *StockSpanner) Next(price int) int {
    span := 1
    
    // Look back while previous prices are smaller or equal
    for len(ss.prices) > 0 && price >= ss.prices[len(ss.prices)-1] {
        span += ss.spans[len(ss.spans)-1]
        ss.prices = ss.prices[:len(ss.prices)-1]
        ss.spans = ss.spans[:len(ss.spans)-1]
    }
    
    ss.prices = append(ss.prices, price)
    ss.spans = append(ss.spans, span)
    
    return span
}

// File: monotonicstack/monotonic_test.go

package monotonicstack

import (
    "reflect"
    "testing"
)

func TestNextGreaterElement(t *testing.T) {
    tests := []struct {
        nums     []int
        expected []int
    }{
        {[]int{4, 5, 2, 25}, []int{5, 25, 25, -1}},
        {[]int{2, 1, 2, 4}, []int{4, 2, 4, -1}},
        {[]int{1, 2, 3, 4}, []int{2, 3, 4, -1}},
        {[]int{4, 3, 2, 1}, []int{-1, -1, -1, -1}},
    }
    
    for _, test := range tests {
        result := NextGreaterElement(test.nums)
        if !reflect.DeepEqual(result, test.expected) {
            t.Errorf("For nums=%v, expected %v but got %v",
                test.nums, test.expected, result)
        }
    }
}

func TestDailyTemperatures(t *testing.T) {
    tests := []struct {
        temperatures []int
        expected    []int
    }{
        {[]int{73, 74, 75, 71, 69, 72, 76, 73}, []int{1, 1, 4, 2, 1, 1, 0, 0}},
        {[]int{30, 40, 50, 60}, []int{1, 1, 1, 0}},
        {[]int{30, 60, 90}, []int{1, 1, 0}},
    }
    
    for _, test := range tests {
        result := DailyTemperatures(test.temperatures)
        if !reflect.DeepEqual(result, test.expected) {
            t.Errorf("For temperatures=%v, expected %v but got %v",
                test.temperatures, test.expected, result)
        }
    }
}

func TestLargestRectangleArea(t *testing.T) {
    tests := []struct {
        heights  []int
        expected int
    }{
        {[]int{2, 1, 5, 6, 2, 3}, 10},
        {[]int{2, 4}, 4},
        {[]int{1}, 1},
        {[]int{}, 0},
    }
    
    for _, test := range tests {
        result := LargestRectangleArea(test.heights)
        if result != test.expected {
            t.Errorf("For heights=%v, expected %d but got %d",
                test.heights, test.expected, result)
        }
    }
}

func TestMaximalRectangle(t *testing.T) {
    tests := []struct {
        matrix   [][]byte
        expected int
    }{
        {
            [][]byte{
                {'1', '0', '1', '0', '0'},
                {'1', '0', '1', '1', '1'},
                {'1', '1', '1', '1', '1'},
                {'1', '0', '0', '1', '0'},
            },
            6,
        },
        {
            [][]byte{{'0'}},
            0,
        },
        {
            [][]byte{{'1'}},
            1,
        },
    }
    
    for _, test := range tests {
        result := MaximalRectangle(test.matrix)
        if result != test.expected {
            t.Errorf("For matrix=%v, expected %d but got %d",
                test.matrix, test.expected, result)
        }
    }
}

func TestStockSpanner(t *testing.T) {
    tests := []struct {
        prices   []int
        expected []int
    }{
        {
            []int{100, 80, 60, 70, 60, 75, 85},
            []int{1, 1, 1, 2, 1, 4, 6},
        },
        {
            []int{10, 20, 30, 40, 50},
            []int{1, 2, 3, 4, 5},
        },
    }
    
    for _, test := range tests {
        spanner := NewStockSpanner()
        for i, price := range test.prices {
            result := spanner.Next(price)
            if result != test.expected[i] {
                t.Errorf("For price=%d (index %d), expected %d but got %d",
                    price, i, test.expected[i], result)
            }
        }
    }
}