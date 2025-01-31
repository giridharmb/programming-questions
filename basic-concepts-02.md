# 2. Step-by-Step Guide for Calculating Complexity

## Time Complexity Calculation Steps

### Step 1: Identify Core Operations
Count operations that scale with input size:
- Assignments
- Comparisons
- Arithmetic operations
- Function calls
- Array/slice access
- Loop iterations

```go
// Example: Finding sum of array
func arraySum(arr []int) int {
    sum := 0           // 1 operation
    for i := 0; i < len(arr); i++ {  // n iterations
        sum += arr[i]  // 2 operations per iteration (access + addition)
    }
    return sum         // 1 operation
}
// Total: 1 + n*2 + 1 = 2n + 2 operations
// Therefore, O(n) time complexity
```

### Step 2: Analyze Loops and Recursion
- Single loop: Multiply by iterations
- Nested loops: Multiply nested iterations
- Recursion: Consider branching factor and depth

```go
// Example: Nested loops
func printPairs(arr []int) {
    n := len(arr)
    for i := 0; i < n; i++ {         // n iterations
        for j := 0; j < n; j++ {     // n iterations for each i
            fmt.Println(arr[i], arr[j]) // 1 operation
        }
    }
}
// Total: n * n * 1 = n² operations
// Therefore, O(n²) time complexity
```

### Step 3: Simplify Expression
1. Keep highest order terms
2. Remove constants
3. Remove coefficients

```go
// Example: Multiple operations
func complexOperation(arr []int) {
    n := len(arr)
    
    // Part 1: O(n) operation
    for i := 0; i < n; i++ {
        fmt.Print(arr[i])
    }
    
    // Part 2: O(n²) operation
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            fmt.Print(arr[i] + arr[j])
        }
    }
    
    // Part 3: O(1) operation
    fmt.Print(arr[0])
}
// Total: O(n) + O(n²) + O(1) = O(n²)
// Keep only highest order term
```

## Space Complexity Calculation Steps

### Step 1: Identify Additional Space
Count extra space that scales with input:
- Variables
- Data structures
- Recursive call stack
- Temporary buffers

```go
// Example: Creating new array
func doubleValues(arr []int) []int {
    result := make([]int, len(arr))  // O(n) space
    for i := 0; i < len(arr); i++ {
        result[i] = arr[i] * 2
    }
    return result
}
// Space Complexity: O(n)
```

### Step 2: Analyze Recursive Space
Consider maximum depth of recursion:
- Each recursive call typically adds a stack frame
- Count maximum number of simultaneous recursive calls

```go
// Example: Recursive factorial
func factorial(n int) int {
    if n <= 1 {
        return 1
    }
    return n * factorial(n - 1)
}
// Space Complexity: O(n) due to recursive stack
```

### Step 3: Special Considerations

#### In-Place Operations
```go
// Example: In-place array reversal
func reverse(arr []int) {
    left, right := 0, len(arr)-1
    for left < right {
        arr[left], arr[right] = arr[right], arr[left]
        left++
        right--
    }
}
// Space Complexity: O(1) as no extra space scales with input
```

#### Temporary Space
```go
// Example: Merge sort with temporary array
func mergeSort(arr []int) []int {
    if len(arr) <= 1 {
        return arr
    }
    
    mid := len(arr) / 2
    left := mergeSort(arr[:mid])
    right := mergeSort(arr[mid:])
    
    // Temporary array for merging
    result := make([]int, len(arr))
    // Merge operation...
    
    return result
}
// Space Complexity: O(n) for temporary arrays
```

## Practical Tips for Analysis

### 1. Common Patterns
- Single loop → O(n)
- Nested loops → O(n^depth)
- Halving size → O(log n)
- Recursive branching → O(branches^depth)

### 2. Space Optimization Techniques
- Use in-place operations when possible
- Reuse existing space
- Clear unused references
- Consider iterative vs recursive solutions

### 3. Verification Steps
1. Test with small inputs
2. Verify scaling behavior
3. Consider edge cases
4. Check both time and space impact

### 4. Performance vs Readability
Sometimes it's better to choose:
- Clearer code over minor optimizations
- Maintainable solutions over complex optimizations
- Standard library over custom implementations