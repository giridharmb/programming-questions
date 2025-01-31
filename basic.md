# Understanding Time and Space Complexity

## Table of Contents
- [Introduction](#introduction)
- [Big O Notation](#big-o-notation)
- [Common Time Complexities](#common-time-complexities)
- [Common Space Complexities](#common-space-complexities)
- [How to Calculate Complexity](#how-to-calculate-complexity)
- [Practical Examples](#practical-examples)
- [Reference Guide](#reference-guide)

## Introduction

Time and space complexity are metrics used to analyze algorithm efficiency:
- **Time Complexity**: How running time increases with input size
- **Space Complexity**: How memory usage increases with input size

## Big O Notation

Big O notation describes the upper bound (worst case) of growth rate:
- Ignores constants and lower-order terms
- Focuses on scalability as input size grows
- Written as O(f(n)) where n is input size

Common notations from fastest to slowest:
- O(1) - Constant time
- O(log n) - Logarithmic time
- O(n) - Linear time
- O(n log n) - Linearithmic time
- O(n²) - Quadratic time
- O(2ⁿ) - Exponential time
- O(n!) - Factorial time

## Common Time Complexities

1. **O(1) - Constant Time**
   - Array access
   - Basic arithmetic operations
   - Hash table insertion/lookup

2. **O(log n) - Logarithmic Time**
   - Binary search
   - Balanced tree operations
   - Number of digits in a number

3. **O(n) - Linear Time**
   - Array traversal
   - Linear search
   - String traversal

4. **O(n log n) - Linearithmic Time**
   - Efficient sorting (Merge sort, Quick sort)
   - Heap operations

5. **O(n²) - Quadratic Time**
   - Nested loops
   - Bubble sort
   - Matrix traversal

## Common Space Complexities

1. **O(1) - Constant Space**
   - Fixed number of variables
   - In-place array operations

2. **O(n) - Linear Space**
   - Arrays/lists proportional to input
   - Stack calls in linear recursion

3. **O(n²) - Quadratic Space**
   - 2D arrays
   - Adjacency matrix for graphs

## How to Calculate Complexity

### Steps for Time Complexity:

1. Identify basic operations (assignments, comparisons, etc.)
2. Count how operations grow with input size
3. Keep highest order term
4. Remove constants and coefficients

### Steps for Space Complexity:

1. Identify additional space needed (variables, data structures)
2. Determine how space grows with input
3. Include recursive call stack if applicable
4. Keep highest order term

## Practical Examples

### Example 1: Find Maximum in Array
```go
func findMax(arr []int) int {
    max := arr[0]
    for i := 1; i < len(arr); i++ {
        if arr[i] > max {
            max = arr[i]
        }
    }
    return max
}
```
**Analysis**:
- Time Complexity: O(n)
  - Single loop through array of size n
  - Each iteration is O(1)
- Space Complexity: O(1)
  - Only one additional variable (max)
  - Space doesn't grow with input

### Example 2: Binary Search
```go
func binarySearch(arr []int, target int) int {
    left, right := 0, len(arr)-1
    
    for left <= right {
        mid := left + (right-left)/2
        if arr[mid] == target {
            return mid
        }
        if arr[mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return -1
}
```
**Analysis**:
- Time Complexity: O(log n)
  - Array size is halved in each iteration
  - Number of iterations is log₂(n)
- Space Complexity: O(1)
  - Only uses a few variables
  - Space doesn't grow with input

### Example 3: Matrix Multiplication
```go
func multiplyMatrix(a, b [][]int) [][]int {
    n := len(a)
    result := make([][]int, n)
    for i := 0; i < n; i++ {
        result[i] = make([]int, n)
        for j := 0; j < n; j++ {
            for k := 0; k < n; k++ {
                result[i][j] += a[i][k] * b[k][j]
            }
        }
    }
    return result
}
```
**Analysis**:
- Time Complexity: O(n³)
  - Three nested loops
  - Each loop runs n times
- Space Complexity: O(n²)
  - Result matrix size is n×n
  - Space grows quadratically with input

### Example 4: Recursive Fibonacci
```go
func fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    return fibonacci(n-1) + fibonacci(n-2)
}
```
**Analysis**:
- Time Complexity: O(2ⁿ)
  - Each call spawns two more calls
  - Forms a binary tree of depth n
- Space Complexity: O(n)
  - Maximum depth of recursion tree is n
  - Stack space is proportional to depth

## Reference Guide

### Quick Reference Table

| Operation                 | Time Complexity | Space Complexity |
|--------------------------|-----------------|------------------|
| Array Access             | O(1)            | O(1)            |
| Linear Search            | O(n)            | O(1)            |
| Binary Search            | O(log n)        | O(1)            |
| Bubble/Insertion Sort    | O(n²)           | O(1)            |
| Merge Sort              | O(n log n)      | O(n)            |
| Quick Sort              | O(n log n)      | O(log n)        |
| Hash Table Operation    | O(1) average    | O(n)            |
| BFS/DFS on Graph        | O(V + E)        | O(V)            |

### Tips for Optimization

1. **Time Optimization**:
   - Use appropriate data structures
   - Avoid nested loops when possible
   - Consider space-time tradeoffs
   - Use efficient algorithms (sorting, searching)

2. **Space Optimization**:
   - Use in-place algorithms when possible
   - Clear unused memory/references
   - Consider iterative vs recursive solutions
   - Reuse existing space

### Common Pitfalls

1. **Time Complexity**:
   - Forgetting hidden loops (string concatenation)
   - Not considering average vs worst case
   - Ignoring constant factors for small inputs

2. **Space Complexity**:
   - Forgetting recursive stack space
   - Not considering temporary space
   - Overlooking system overhead