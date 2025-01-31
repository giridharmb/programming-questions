# 1. Basic Concepts of Time and Space Complexity

## Understanding Big O Notation

### What is Big O?
- Big O notation represents the upper bound (worst-case scenario) of algorithm's growth rate
- Written as O(f(n)) where n is the input size
- Ignores constants and lower-order terms
- Focuses on how algorithm scales with larger inputs

### Why Use Big O?
1. Standardized way to compare algorithms
2. Focuses on long-term growth rate
3. Helps make architectural decisions
4. Identifies potential scalability issues

## Common Time Complexities

### O(1) - Constant Time
- Execution time stays same regardless of input size
- Examples:
  ```go
  func getFirst(arr []int) int {
      return arr[0]  // Always one operation
  }
  ```

### O(log n) - Logarithmic Time
- Execution time increases logarithmically
- Time doubles when input size squares
- Examples: Binary search
  ```go
  func binarySearch(arr []int, target int) int {
      left, right := 0, len(arr)-1
      for left <= right {
          mid := (left + right) / 2
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

### O(n) - Linear Time
- Execution time grows linearly with input
- Examples: Linear search
  ```go
  func linearSearch(arr []int, target int) int {
      for i, v := range arr {
          if v == target {
              return i
          }
      }
      return -1
  }
  ```

### O(n log n) - Linearithmic Time
- Common in efficient sorting algorithms
- Examples: Merge sort, Quick sort
  ```go
  func mergeSort(arr []int) []int {
      if len(arr) <= 1 {
          return arr
      }
      mid := len(arr) / 2
      left := mergeSort(arr[:mid])
      right := mergeSort(arr[mid:])
      return merge(left, right)
  }
  ```

### O(n²) - Quadratic Time
- Execution time grows quadratically
- Examples: Nested loops, bubble sort
  ```go
  func bubbleSort(arr []int) {
      n := len(arr)
      for i := 0; i < n; i++ {
          for j := 0; j < n-i-1; j++ {
              if arr[j] > arr[j+1] {
                  arr[j], arr[j+1] = arr[j+1], arr[j]
              }
          }
      }
  }
  ```

## Common Space Complexities

### O(1) - Constant Space
- Fixed amount of memory regardless of input
- Example: In-place array operations
  ```go
  func findMax(arr []int) int {
      max := arr[0]
      for _, v := range arr {
          if v > max {
              max = v
          }
      }
      return max
  }
  ```

### O(n) - Linear Space
- Memory usage grows linearly with input
- Example: Creating new array proportional to input
  ```go
  func doubleArray(arr []int) []int {
      result := make([]int, len(arr))
      for i, v := range arr {
          result[i] = v * 2
      }
      return result
  }
  ```

### O(n²) - Quadratic Space
- Memory usage grows quadratically
- Example: 2D matrix operations
  ```go
  func createMatrix(n int) [][]int {
      matrix := make([][]int, n)
      for i := range matrix {
          matrix[i] = make([]int, n)
      }
      return matrix
  }
  ```

## Important Concepts to Remember

1. **Best vs Average vs Worst Case**
   - Best Case (Ω - Omega notation)
   - Average Case (θ - Theta notation)
   - Worst Case (O - Big O notation)

2. **Space Complexity Includes**
   - Auxiliary space (extra space)
   - Input space (space needed for input)

3. **When Analyzing Complexity**
   - Focus on dominant terms
   - Drop constants
   - Consider both time and space
   - Think about scalability

4. **Common Mistakes to Avoid**
   - Forgetting about hidden loops
   - Not considering the input space
   - Overlooking auxiliary space requirements
   - Focusing only on time complexity