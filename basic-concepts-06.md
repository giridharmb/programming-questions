# 6. Common Pitfalls in Time and Space Complexity

## Hidden Time Complexity Issues

### 1. String Operations
```go
// Pitfall: Hidden O(n) operation in string concatenation
func buildString(n int) string {
    result := ""
    for i := 0; i < n; i++ {
        result += "a"  // Each concatenation creates new string
    }
    return result
}
// Actual Time: O(n²)

// Better: Use strings.Builder
func buildString(n int) string {
    var builder strings.Builder
    for i := 0; i < n; i++ {
        builder.WriteRune('a')
    }
    return builder.String()
}
// Time: O(n)
```

### 2. Slice Operations
```go
// Pitfall: Hidden copy operations
func removeFirst(slice []int) []int {
    return slice[1:]  // Creates new slice header but shares array
}

// Better: Clear memory reference
func removeFirst(slice []int) []int {
    if len(slice) == 0 {
        return slice
    }
    copy(slice, slice[1:])
    return slice[:len(slice)-1]
}
```

### 3. Map Access in Loops
```go
// Pitfall: Repeated map lookups
func processMap(m map[string]int, key string) int {
    sum := 0
    for i := 0; i < 1000; i++ {
        sum += m[key]  // Map lookup in each iteration
    }
    return sum
}

// Better: Cache the lookup
func processMap(m map[string]int, key string) int {
    value := m[key]  // Single map lookup
    sum := 0
    for i := 0; i < 1000; i++ {
        sum += value
    }
    return sum
}
```

## Hidden Space Complexity Issues

### 1. Recursive Stack Space
```go
// Pitfall: Unbounded recursion
func factorial(n int) int {
    if n <= 1 {
        return 1
    }
    return n * factorial(n-1)
}
// Space: O(n) due to call stack

// Better: Iterative solution
func factorial(n int) int {
    result := 1
    for i := 2; i <= n; i++ {
        result *= i
    }
    return result
}
// Space: O(1)
```

### 2. Slice Capacity Growth
```go
// Pitfall: Unnecessary capacity
func buildSlice() []int {
    result := make([]int, 0, 1000000)  // Over-allocation
    for i := 0; i < 10; i++ {
        result = append(result, i)
    }
    return result
}

// Better: Appropriate initial capacity
func buildSlice() []int {
    result := make([]int, 0, 10)
    for i := 0; i < 10; i++ {
        result = append(result, i)
    }
    return result
}
```

### 3. Goroutine Leaks
```go
// Pitfall: Goroutine leak
func processChannel(ch chan int) {
    go func() {
        // This goroutine might never exit
        for val := range ch {
            fmt.Println(val)
        }
    }()
}

// Better: Ensure cleanup
func processChannel(ch chan int, done chan bool) {
    go func() {
        defer close(done)
        for val := range ch {
            fmt.Println(val)
        }
    }()
}
```

## Algorithm Analysis Mistakes

### 1. Ignoring Input Distribution
```go
// Pitfall: Assuming worst case is common
func quickSort(arr []int) {
    // Basic quicksort implementation
    // Worst case: O(n²)
    // Average case: O(n log n)
}

// Better: Consider randomization
func quickSortRandomized(arr []int) {
    // Randomized pivot selection
    // Makes worst case very unlikely
}
```

### 2. Overlooking Constants
```go
// Pitfall: Choosing based on Big O only
func search(arr []int, target int) bool {
    // Binary search: O(log n)
    // But with complex calculations...
}

// Sometimes better for small n
func linearSearch(arr []int, target int) bool {
    // Linear search: O(n)
    // But very simple operations
}
```

### 3. Incorrect Complexity Addition
```go
// Pitfall: Adding complexities incorrectly
func process(arr []int) {
    // O(n) operation
    for i := 0; i < len(arr); i++ {
        fmt.Println(arr[i])
    }
    
    // O(n²) operation
    for i := 0; i < len(arr); i++ {
        for j := 0; j < len(arr); j++ {
            fmt.Println(arr[i] + arr[j])
        }
    }
}
// Total is O(n²), not O(n + n²)
```

## Best Practices to Avoid Pitfalls

### 1. Performance Testing
```go
// Always test with different input sizes
func TestPerformance(t *testing.T) {
    sizes := []int{100, 1000, 10000, 100000}
    for _, size := range sizes {
        input := generateInput(size)
        start := time.Now()
        process(input)
        duration := time.Since(start)
        t.Logf("Size %d took %v", size, duration)
    }
}
```

### 2. Memory Profiling
```go
// Use memory profiling to detect leaks
import "runtime/pprof"

func main() {
    f, _ := os.Create("mem.prof")
    defer f.Close()
    
    // Run your code
    
    pprof.WriteHeapProfile(f)
}
```

### 3. Code Review Checklist
```go
// Before submitting code, check:
// 1. Are there any hidden loops?
// 2. Is memory being freed properly?
// 3. Are goroutines properly managed?
// 4. Are string operations optimized?
// 5. Are map/slice operations efficient?
```

## General Guidelines

1. **Always Measure**
   - Don't rely solely on theoretical analysis
   - Use benchmarks and profiling
   - Test with realistic data sizes

2. **Consider Trade-offs**
   - Time vs Space complexity
   - Readability vs Performance
   - Simplicity vs Optimization

3. **Document Assumptions**
   - Input size ranges
   - Expected data distributions
   - Performance requirements

4. **Review Edge Cases**
   - Empty inputs
   - Large inputs
   - Invalid inputs
   - Concurrent access

5. **Maintain Balance**
   - Don't over-optimize prematurely
   - Keep code maintainable
   - Document complex optimizations