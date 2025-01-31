/*

This implementation includes six classic Dynamic Programming problems:

Fibonacci Numbers - Both top-down and bottom-up approaches
Longest Increasing Subsequence
0/1 Knapsack - Both approaches
Coin Change - Minimum coins needed
Edit Distance - String transformation
Longest Common Subsequence

Key features:

Both memoization (top-down) and tabulation (bottom-up) approaches
Space optimization techniques
State transition logic
Subproblem identification
Comprehensive test cases
Clean Go idioms

*/

// File: dp/dp.go

package dp

import "math"

// FibonacciTopDown calculates nth Fibonacci number using memoization
// Time Complexity: O(n)
// Space Complexity: O(n)
func FibonacciTopDown(n int) int {
    memo := make(map[int]int)
    return fibHelper(n, memo)
}

func fibHelper(n int, memo map[int]int) int {
    if n <= 1 {
        return n
    }
    
    if val, exists := memo[n]; exists {
        return val
    }
    
    memo[n] = fibHelper(n-1, memo) + fibHelper(n-2, memo)
    return memo[n]
}

// FibonacciBottomUp calculates nth Fibonacci number using tabulation
// Time Complexity: O(n)
// Space Complexity: O(1)
func FibonacciBottomUp(n int) int {
    if n <= 1 {
        return n
    }
    
    a, b := 0, 1
    for i := 2; i <= n; i++ {
        a, b = b, a+b
    }
    return b
}

// LongestIncreasingSubsequence finds length of longest increasing subsequence
// Time Complexity: O(n^2)
// Space Complexity: O(n)
func LongestIncreasingSubsequence(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    
    // dp[i] represents LIS ending at index i
    dp := make([]int, len(nums))
    for i := range dp {
        dp[i] = 1
    }
    
    maxLen := 1
    for i := 1; i < len(nums); i++ {
        for j := 0; j < i; j++ {
            if nums[i] > nums[j] {
                dp[i] = max(dp[i], dp[j]+1)
            }
        }
        maxLen = max(maxLen, dp[i])
    }
    
    return maxLen
}

// KnapsackTopDown solves 0/1 knapsack problem using memoization
// Time Complexity: O(n*w)
// Space Complexity: O(n*w)
func KnapsackTopDown(weights []int, values []int, capacity int) int {
    memo := make(map[string]int)
    return knapsackHelper(weights, values, capacity, len(weights)-1, memo)
}

func knapsackHelper(weights []int, values []int, capacity int, index int, 
    memo map[string]int) int {
    
    if index < 0 || capacity <= 0 {
        return 0
    }
    
    key := fmt.Sprintf("%d,%d", index, capacity)
    if val, exists := memo[key]; exists {
        return val
    }
    
    // Don't include current item
    result := knapsackHelper(weights, values, capacity, index-1, memo)
    
    // Include current item if possible
    if weights[index] <= capacity {
        include := values[index] + knapsackHelper(weights, values, 
            capacity-weights[index], index-1, memo)
        result = max(result, include)
    }
    
    memo[key] = result
    return result
}

// KnapsackBottomUp solves 0/1 knapsack problem using tabulation
// Time Complexity: O(n*w)
// Space Complexity: O(n*w)
func KnapsackBottomUp(weights []int, values []int, capacity int) int {
    n := len(weights)
    dp := make([][]int, n+1)
    for i := range dp {
        dp[i] = make([]int, capacity+1)
    }
    
    for i := 1; i <= n; i++ {
        for w := 0; w <= capacity; w++ {
            if weights[i-1] <= w {
                dp[i][w] = max(values[i-1]+dp[i-1][w-weights[i-1]], dp[i-1][w])
            } else {
                dp[i][w] = dp[i-1][w]
            }
        }
    }
    
    return dp[n][capacity]
}

// CoinChange finds minimum coins needed to make amount
// Time Complexity: O(amount * len(coins))
// Space Complexity: O(amount)
func CoinChange(coins []int, amount int) int {
    dp := make([]int, amount+1)
    for i := range dp {
        dp[i] = amount + 1
    }
    dp[0] = 0
    
    for i := 1; i <= amount; i++ {
        for _, coin := range coins {
            if coin <= i {
                dp[i] = min(dp[i], dp[i-coin]+1)
            }
        }
    }
    
    if dp[amount] > amount {
        return -1
    }
    return dp[amount]
}

// EditDistance calculates minimum operations to convert word1 to word2
// Time Complexity: O(m*n)
// Space Complexity: O(m*n)
func EditDistance(word1 string, word2 string) int {
    m, n := len(word1), len(word2)
    dp := make([][]int, m+1)
    for i := range dp {
        dp[i] = make([]int, n+1)
    }
    
    // Initialize first row and column
    for i := 0; i <= m; i++ {
        dp[i][0] = i
    }
    for j := 0; j <= n; j++ {
        dp[0][j] = j
    }
    
    // Fill dp table
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            if word1[i-1] == word2[j-1] {
                dp[i][j] = dp[i-1][j-1]
            } else {
                dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1])) + 1
            }
        }
    }
    
    return dp[m][n]
}

// LongestCommonSubsequence finds length of LCS
// Time Complexity: O(m*n)
// Space Complexity: O(m*n)
func LongestCommonSubsequence(text1 string, text2 string) int {
    m, n := len(text1), len(text2)
    dp := make([][]int, m+1)
    for i := range dp {
        dp[i] = make([]int, n+1)
    }
    
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            if text1[i-1] == text2[j-1] {
                dp[i][j] = dp[i-1][j-1] + 1
            } else {
                dp[i][j] = max(dp[i-1][j], dp[i][j-1])
            }
        }
    }
    
    return dp[m][n]
}

// Helper functions
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

// File: dp/dp_test.go

package dp

import "testing"

func TestFibonacci(t *testing.T) {
    tests := []struct {
        name     string
        n        int
        expected int
    }{
        {"Base case 0", 0, 0},
        {"Base case 1", 1, 1},
        {"Simple case", 5, 5},
        {"Larger case", 10, 55},
    }
    
    for _, test := range tests {
        t.Run(test.name+" Top Down", func(t *testing.T) {
            result := FibonacciTopDown(test.n)
            if result != test.expected {
                t.Errorf("Expected %d but got %d", test.expected, result)
            }
        })
        
        t.Run(test.name+" Bottom Up", func(t *testing.T) {
            result := FibonacciBottomUp(test.n)
            if result != test.expected {
                t.Errorf("Expected %d but got %d", test.expected, result)
            }
        })
    }
}

func TestLongestIncreasingSubsequence(t *testing.T) {
    tests := []struct {
        name     string
        nums     []int
        expected int
    }{
        {
            "Simple case",
            []int{10, 9, 2, 5, 3, 7, 101, 18},
            4,
        },
        {
            "All increasing",
            []int{1, 2, 3, 4, 5},
            5,
        },
        {
            "All decreasing",
            []int{5, 4, 3, 2, 1},
            1,
        },
    }
    
    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            result := LongestIncreasingSubsequence(test.nums)
            if result != test.expected {
                t.Errorf("Expected %d but got %d", test.expected, result)
            }
        })
    }
}

func TestKnapsack(t *testing.T) {
    tests := []struct {
        name     string
        weights  []int
        values   []int
        capacity int
        expected int
    }{
        {
            "Simple case",
            []int{1, 2, 3},
            []int{6, 10, 12},
            5,
            22,
        },
        {
            "No items fit",
            []int{5, 4, 6},
            []int{10, 8, 12},
            3,
            0,
        },
    }
    
    for _, test := range tests {
        t.Run(test.name+" Top Down", func(t *testing.T) {
            result := KnapsackTopDown(test.weights, test.values, test.capacity)
            if result != test.expected {
                t.Errorf("Expected %d but got %d", test.expected, result)
            }
        })
        
        t.Run(test.name+" Bottom Up", func(t *testing.T) {
            result := KnapsackBottomUp(test.weights, test.values, test.capacity)
            if result != test.expected {
                t.Errorf("Expected %d but got %d", test.expected, result)
            }
        })
    }
}

func TestCoinChange(t *testing.T) {
    tests := []struct {
        name     string
        coins    []int
        amount   int
        expected int
    }{
        {"Simple case", []int{1, 2, 5}, 11, 3},
        {"No solution", []int{2}, 3, -1},
        {"Zero amount", []int{1}, 0, 0},
    }
    
    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            result := CoinChange(test.coins, test.amount)
            if result != test.expected {
                t.Errorf("Expected %d but got %d", test.expected, result)
            }
        })
    }
}

func TestEditDistance(t *testing.T) {
    tests := []struct {
        name     string
        word1    string
        word2    string
        expected int
    }{
        {"Simple case", "horse", "ros", 3},
        {"Empty string", "", "a", 1},
        {"Same string", "abc", "abc", 0},
    }
    
    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            result := EditDistance(test.word1, test.word2)
            if result != test.expected {
                t.Errorf("Expected %d but got %d", test.expected, result)
            }
        })
    }
}

func TestLongestCommonSubsequence(t *testing.T) {
    tests := []struct {
        name     string
        text1    string
        text2    string
        expected int
    }{
        {"Simple case", "abcde", "ace", 3},
        {"No common", "abc", "def", 0},
        {"Same string", "abc", "abc", 3},
    }
    
    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            result := LongestCommonSubsequence(test.text1, test.text2)
            if result != test.expected {
                t.Errorf("Expected %d but got %d", test.expected, result)
            }
        })
    }
}