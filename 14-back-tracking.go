/*

This implementation includes six classic Backtracking problems:

Permutations - Generate all possible permutations
Combinations - Generate all possible k combinations
Subsets - Generate all possible subsets
N-Queens - Solve N-Queens puzzle
Palindrome Partitioning - Generate all possible palindrome partitions
Generate Parentheses - Generate all valid parentheses combinations

Key features:

Systematic backtracking approach
State management and restoration
Pruning invalid paths
Path tracking
Comprehensive test cases
Complex problem solving

*/

// File: backtracking/backtracking.go

package backtracking

import "sort"

// Permutations generates all possible permutations of nums
// Time Complexity: O(n!)
// Space Complexity: O(n!)
func Permutations(nums []int) [][]int {
    result := make([][]int, 0)
    permuteHelper(nums, 0, &result)
    return result
}

func permuteHelper(nums []int, start int, result *[][]int) {
    if start == len(nums) {
        permutation := make([]int, len(nums))
        copy(permutation, nums)
        *result = append(*result, permutation)
        return
    }
    
    for i := start; i < len(nums); i++ {
        // Swap elements
        nums[start], nums[i] = nums[i], nums[start]
        permuteHelper(nums, start+1, result)
        // Backtrack
        nums[start], nums[i] = nums[i], nums[start]
    }
}

// Combinations generates all possible k combinations from 1 to n
// Time Complexity: O(n choose k)
// Space Complexity: O(k)
func Combinations(n int, k int) [][]int {
    result := make([][]int, 0)
    current := make([]int, 0)
    combineHelper(n, k, 1, current, &result)
    return result
}

func combineHelper(n, k, start int, current []int, result *[][]int) {
    if len(current) == k {
        combination := make([]int, k)
        copy(combination, current)
        *result = append(*result, combination)
        return
    }
    
    for i := start; i <= n; i++ {
        current = append(current, i)
        combineHelper(n, k, i+1, current, result)
        current = current[:len(current)-1]
    }
}

// Subsets generates all possible subsets of nums
// Time Complexity: O(2^n)
// Space Complexity: O(2^n)
func Subsets(nums []int) [][]int {
    result := make([][]int, 0)
    current := make([]int, 0)
    subsetsHelper(nums, 0, current, &result)
    return result
}

func subsetsHelper(nums []int, index int, current []int, result *[][]int) {
    subset := make([]int, len(current))
    copy(subset, current)
    *result = append(*result, subset)
    
    for i := index; i < len(nums); i++ {
        current = append(current, nums[i])
        subsetsHelper(nums, i+1, current, result)
        current = current[:len(current)-1]
    }
}

// NQueens solves N-Queens puzzle
// Time Complexity: O(n!)
// Space Complexity: O(n)
func NQueens(n int) [][]string {
    result := make([][]string, 0)
    board := make([][]byte, n)
    for i := range board {
        board[i] = make([]byte, n)
        for j := range board[i] {
            board[i][j] = '.'
        }
    }
    
    solveNQueens(board, 0, &result)
    return result
}

func solveNQueens(board [][]byte, row int, result *[][]string) {
    if row == len(board) {
        solution := make([]string, len(board))
        for i := range board {
            solution[i] = string(board[i])
        }
        *result = append(*result, solution)
        return
    }
    
    for col := 0; col < len(board); col++ {
        if isValidQueenPlacement(board, row, col) {
            board[row][col] = 'Q'
            solveNQueens(board, row+1, result)
            board[row][col] = '.'
        }
    }
}

func isValidQueenPlacement(board [][]byte, row, col int) bool {
    // Check column
    for i := 0; i < row; i++ {
        if board[i][col] == 'Q' {
            return false
        }
    }
    
    // Check left diagonal
    for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
        if board[i][j] == 'Q' {
            return false
        }
    }
    
    // Check right diagonal
    for i, j := row-1, col+1; i >= 0 && j < len(board); i, j = i-1, j+1 {
        if board[i][j] == 'Q' {
            return false
        }
    }
    
    return true
}

// PalindromePartitioning generates all possible palindrome partitioning
// Time Complexity: O(n * 2^n)
// Space Complexity: O(n)
func PalindromePartitioning(s string) [][]string {
    result := make([][]string, 0)
    current := make([]string, 0)
    partitionHelper(s, 0, current, &result)
    return result
}

func partitionHelper(s string, start int, current []string, result *[][]string) {
    if start >= len(s) {
        partition := make([]string, len(current))
        copy(partition, current)
        *result = append(*result, partition)
        return
    }
    
    for end := start; end < len(s); end++ {
        if isPalindrome(s[start:end+1]) {
            current = append(current, s[start:end+1])
            partitionHelper(s, end+1, current, result)
            current = current[:len(current)-1]
        }
    }
}

func isPalindrome(s string) bool {
    left, right := 0, len(s)-1
    for left < right {
        if s[left] != s[right] {
            return false
        }
        left++
        right--
    }
    return true
}

// GenerateParentheses generates all valid parenthesis combinations
// Time Complexity: O(4^n / sqrt(n))
// Space Complexity: O(n)
func GenerateParentheses(n int) []string {
    result := make([]string, 0)
    generateParenHelper("", 0, 0, n, &result)
    return result
}

func generateParenHelper(current string, open, close, n int, result *[]string) {
    if len(current) == 2*n {
        *result = append(*result, current)
        return
    }
    
    if open < n {
        generateParenHelper(current+"(", open+1, close, n, result)
    }
    if close < open {
        generateParenHelper(current+")", open, close+1, n, result)
    }
}

// File: backtracking/backtracking_test.go

package backtracking

import (
    "reflect"
    "sort"
    "testing"
)

func TestPermutations(t *testing.T) {
    tests := []struct {
        name     string
        nums     []int
        expected [][]int
    }{
        {
            "Three numbers",
            []int{1, 2, 3},
            [][]int{
                {1, 2, 3}, {1, 3, 2}, {2, 1, 3},
                {2, 3, 1}, {3, 1, 2}, {3, 2, 1},
            },
        },
        {
            "Two numbers",
            []int{1, 2},
            [][]int{{1, 2}, {2, 1}},
        },
    }
    
    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            result := Permutations(test.nums)
            if !reflect.DeepEqual(result, test.expected) {
                t.Errorf("Expected %v but got %v", test.expected, result)
            }
        })
    }
}

func TestCombinations(t *testing.T) {
    tests := []struct {
        name     string
        n, k     int
        expected [][]int
    }{
        {
            "Four choose two",
            4, 2,
            [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}},
        },
    }
    
    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            result := Combinations(test.n, test.k)
            if !reflect.DeepEqual(result, test.expected) {
                t.Errorf("Expected %v but got %v", test.expected, result)
            }
        })
    }
}

func TestSubsets(t *testing.T) {
    tests := []struct {
        name     string
        nums     []int
        expected [][]int
    }{
        {
            "Three numbers",
            []int{1, 2, 3},
            [][]int{
                {}, {1}, {2}, {3},
                {1, 2}, {1, 3}, {2, 3},
                {1, 2, 3},
            },
        },
    }
    
    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            result := Subsets(test.nums)
            // Sort for comparison
            sortNestedSlices(result)
            sortNestedSlices(test.expected)
            if !reflect.DeepEqual(result, test.expected) {
                t.Errorf("Expected %v but got %v", test.expected, result)
            }
        })
    }
}

func TestNQueens(t *testing.T) {
    tests := []struct {
        name     string
        n        int
        expected int // Number of solutions
    }{
        {"4x4 board", 4, 2},
        {"8x8 board", 8, 92},
    }
    
    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            result := NQueens(test.n)
            if len(result) != test.expected {
                t.Errorf("Expected %d solutions but got %d",
                    test.expected, len(result))
            }
        })
    }
}

func TestPalindromePartitioning(t *testing.T) {
    tests := []struct {
        name     string
        s        string
        expected [][]string
    }{
        {
            "Simple string",
            "aab",
            [][]string{{"a", "a", "b"}, {"aa", "b"}},
        },
    }
    
    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            result := PalindromePartitioning(test.s)
            if !reflect.DeepEqual(result, test.expected) {
                t.Errorf("Expected %v but got %v", test.expected, result)
            }
        })
    }
}

func TestGenerateParentheses(t *testing.T) {
    tests := []struct {
        name     string
        n        int
        expected []string
    }{
        {
            "Three pairs",
            3,
            []string{
                "((()))", "(()())", "(())()",
                "()(())", "()()()",
            },
        },
    }
    
    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            result := GenerateParentheses(test.n)
            sort.Strings(result)
            sort.Strings(test.expected)
            if !reflect.DeepEqual(result, test.expected) {
                t.Errorf("Expected %v but got %v", test.expected, result)
            }
        })
    }
}

// Helper function to sort nested slices
func sortNestedSlices(slices [][]int) {
    for _, slice := range slices {
        sort.Ints(slice)
    }
    sort.Slice(slices, func(i, j int) bool {
        if len(slices[i]) != len(slices[j]) {
            return len(slices[i]) < len(slices[j])
        }
        for k := 0; k < len(slices[i]); k++ {
            if slices[i][k] != slices[j][k] {
                return slices[i][k] < slices[j][k]
            }
        }
        return false
    })
}