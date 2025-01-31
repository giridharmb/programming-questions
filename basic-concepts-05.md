# 5. Optimization Tips and Best Practices

## Time Optimization Strategies

### 1. Data Structure Selection
```go
// Bad: Using array for frequent lookups
func containsValue(arr []int, target int) bool {
    for _, v := range arr {
        if v == target {
            return true
        }
    }
    return false
}
// Time: O(n)

// Good: Using map for frequent lookups
func containsValue(m map[int]bool, target int) bool {
    _, exists := m[target]
    return exists
}
// Time: O(1) average
```

### 2. Precomputation and Caching
```go
// Without caching
func fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    return fibonacci(n-1) + fibonacci(n-2)
}
// Time: O(2ⁿ)

// With memoization
func fibonacci(n int, memo map[int]int) int {
    if n <= 1 {
        return n
    }
    if val, exists := memo[n]; exists {
        return val
    }
    memo[n] = fibonacci(n-1, memo) + fibonacci(n-2, memo)
    return memo[n]
}
// Time: O(n)
```

### 3. Early Termination
```go
// Bad: Always checking all elements
func findDuplicate(arr []int) bool {
    n := len(arr)
    for i := 0; i < n; i++ {
        for j := i + 1; j < n; j++ {
            if arr[i] == arr[j] {
                return true
            }
        }
    }
    return false
}

// Good: Return as soon as found
func findDuplicate(arr []int) bool {
    seen := make(map[int]bool)
    for _, num := range arr {
        if seen[num] {
            return true
        }
        seen[num] = true
    }
    return false
}
```

## Space Optimization Strategies

### 1. In-Place Operations
```go
// Using extra space
func reverse(arr []int) []int {
    result := make([]int, len(arr))
    for i := 0; i < len(arr); i++ {
        result[i] = arr[len(arr)-1-i]
    }
    return result
}
// Space: O(n)

// In-place operation
func reverse(arr []int) {
    left, right := 0, len(arr)-1
    for left < right {
        arr[left], arr[right] = arr[right], arr[left]
        left++
        right--
    }
}
// Space: O(1)
```

### 2. Bit Manipulation
```go
// Using array/map
func hasDuplicates(nums []int) bool {
    seen := make(map[int]bool)
    for _, num := range nums {
        if seen[num] {
            return true
        }
        seen[num] = true
    }
    return false
}
// Space: O(n)

// Using bit manipulation (for small integers)
func hasDuplicates(nums []int) bool {
    var seen uint64
    for _, num := range nums {
        if seen&(1<<uint(num)) != 0 {
            return true
        }
        seen |= 1 << uint(num)
    }
    return false
}
// Space: O(1)
```

### 3. Sliding Window
```go
// Using extra space
func maxSumSubarray(arr []int, k int) int {
    sums := make([]int, len(arr)-k+1)
    for i := 0; i <= len(arr)-k; i++ {
        sum := 0
        for j := 0; j < k; j++ {
            sum += arr[i+j]
        }
        sums[i] = sum
    }
    return max(sums...)
}
// Space: O(n)

// Using sliding window
func maxSumSubarray(arr []int, k int) int {
    windowSum := 0
    maxSum := 0
    for i := 0; i < k; i++ {
        windowSum += arr[i]
    }
    maxSum = windowSum
    
    for i := k; i < len(arr); i++ {
        windowSum = windowSum - arr[i-k] + arr[i]
        maxSum = max(maxSum, windowSum)
    }
    return maxSum
}
// Space: O(1)
```

## Code Organization Best Practices

### 1. Clear Variable Names
```go
// Bad
func p(a []int, t int) bool {
    for _, v := range a {
        if v == t {
            return true
        }
    }
    return false
}

// Good
func isPresent(array []int, target int) bool {
    for _, value := range array {
        if value == target {
            return true
        }
    }
    return false
}
```

### 2. Function Documentation
```go
// Bad
func process(nums []int) int {
    // ...
}

// Good
// calculateSum returns the sum of all positive numbers in the array
// Time Complexity: O(n)
// Space Complexity: O(1)
func calculateSum(nums []int) int {
    sum := 0
    for _, num := range nums {
        if num > 0 {
            sum += num
        }
    }
    return sum
}
```

### 3. Error Handling
```go
// Bad
func divide(a, b int) int {
    return a / b
}

// Good
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}
```

## Performance Monitoring Tips

### 1. Benchmarking
```go
func BenchmarkAlgorithm(b *testing.B) {
    input := generateTestData()
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        algorithm(input)
    }
}
```

### 2. Profiling
```go
import "runtime/pprof"

func main() {
    f, _ := os.Create("cpu.prof")
    pprof.StartCPUProfile(f)
    defer pprof.StopCPUProfile()
    
    // Your algorithm here
}
```

### 3. Memory Usage Tracking
```go
var m runtime.MemStats
runtime.ReadMemStats(&m)
fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
```

## Optimization Checklist

1. **Algorithm Selection**
   - Is this the best algorithm for the problem?
   - Are there any known better alternatives?
   - Does the input size justify a more complex solution?

2. **Data Structure Choice**
   - Is this the most efficient data structure?
   - Could a different structure improve performance?
   - Are we using too much memory?

3. **Code Implementation**
   - Are we using efficient operations?
   - Can we reduce the number of loops?
   - Are we handling edge cases efficiently?

4. **Memory Management**
   - Are we releasing resources properly?
   - Can we reduce memory allocations?
   - Are we using memory efficiently?

5. **Testing and Verification**
   - Have we tested with different input sizes?
   - Are we handling edge cases correctly?
   - Have we benchmarked the solution?