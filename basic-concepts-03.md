# 3. Practical Examples with Complexity Analysis

## Example 1: Linear Search vs Binary Search

### Linear Search
```go
func linearSearch(arr []int, target int) int {
    for i, num := range arr {
        if num == target {
            return i
        }
    }
    return -1
}
```

**Analysis:**
- Time Complexity: O(n)
  - Must check each element in worst case
  - Number of operations grows linearly with input size
- Space Complexity: O(1)
  - Only uses a few variables regardless of input size
  - No additional data structures created

### Binary Search
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

**Analysis:**
- Time Complexity: O(log n)
  - Input space is halved in each iteration
  - Takes log₂(n) steps to find target
- Space Complexity: O(1)
  - Only uses three variables (left, right, mid)
  - Space usage doesn't grow with input

## Example 2: Matrix Operations

### Matrix Sum
```go
func matrixSum(matrix [][]int) int {
    sum := 0
    for i := 0; i < len(matrix); i++ {
        for j := 0; j < len(matrix[0]); j++ {
            sum += matrix[i][j]
        }
    }
    return sum
}
```

**Analysis:**
- Time Complexity: O(m*n)
  - Visits each element once
  - m rows * n columns total operations
- Space Complexity: O(1)
  - Only uses one variable for sum
  - No additional space needed

### Matrix Multiplication
```go
func matrixMultiply(a, b [][]int) [][]int {
    m, n, p := len(a), len(a[0]), len(b[0])
    result := make([][]int, m)
    for i := range result {
        result[i] = make([]int, p)
    }
    
    for i := 0; i < m; i++ {
        for j := 0; j < p; j++ {
            for k := 0; k < n; k++ {
                result[i][j] += a[i][k] * b[k][j]
            }
        }
    }
    return result
}
```

**Analysis:**
- Time Complexity: O(n³) for square matrices
  - Three nested loops
  - Each loop runs n times
- Space Complexity: O(n²)
  - Result matrix size is m*p
  - Space grows quadratically for square matrices

## Example 3: Tree Traversal

### Recursive DFS (Depth-First Search)
```go
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func dfs(root *TreeNode) []int {
    result := []int{}
    dfsHelper(root, &result)
    return result
}

func dfsHelper(node *TreeNode, result *[]int) {
    if node == nil {
        return
    }
    *result = append(*result, node.Val)
    dfsHelper(node.Left, result)
    dfsHelper(node.Right, result)
}
```

**Analysis:**
- Time Complexity: O(n)
  - Visits each node exactly once
  - n is number of nodes in tree
- Space Complexity: O(h)
  - h is height of tree (recursion stack)
  - Best case O(log n) for balanced tree
  - Worst case O(n) for skewed tree

### Iterative BFS (Breadth-First Search)
```go
func bfs(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }
    
    result := []int{}
    queue := []*TreeNode{root}
    
    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]
        result = append(result, node.Val)
        
        if node.Left != nil {
            queue = append(queue, node.Left)
        }
        if node.Right != nil {
            queue = append(queue, node.Right)
        }
    }
    return result
}
```

**Analysis:**
- Time Complexity: O(n)
  - Visits each node exactly once
  - Queue operations are O(1)
- Space Complexity: O(w)
  - w is maximum width of tree
  - Best/Average case O(n/2) ≈ O(n)
  - Worst case for last level of complete binary tree

## Example 4: String Manipulation

### String Reversal
```go
func reverseString(s string) string {
    runes := []rune(s)
    left, right := 0, len(runes)-1
    
    for left < right {
        runes[left], runes[right] = runes[right], runes[left]
        left++
        right--
    }
    return string(runes)
}
```

**Analysis:**
- Time Complexity: O(n)
  - Processes each character once
  - String conversion and creation is O(n)
- Space Complexity: O(n)
  - Creates new rune slice of size n
  - Returns new string of size n

### Finding All Substrings
```go
func findAllSubstrings(s string) []string {
    result := make([]string, 0)
    for i := 0; i < len(s); i++ {
        for j := i + 1; j <= len(s); j++ {
            result = append(result, s[i:j])
        }
    }
    return result
}
```

**Analysis:**
- Time Complexity: O(n²)
  - Nested loops over string length
  - String slicing is O(k) where k is substring length
  - Total operations ≈ n²/2
- Space Complexity: O(n³)
  - Stores all possible substrings
  - Number of substrings is n(n+1)/2
  - Each substring can be up to length n

## Key Insights from Examples

1. **Operation Counting:**
   - Count dominant operations
   - Consider hidden operations (string conversions, slice operations)
   - Look for nested iterations

2. **Space Analysis:**
   - Consider both auxiliary and input space
   - Account for recursive stack space
   - Look for data structure growth

3. **Optimization Opportunities:**
   - Trading space for time (caching)
   - Using appropriate data structures
   - Choosing iterative vs recursive approaches

4. **Common Patterns:**
   - Linear scanning: O(n)
   - Divide and conquer: O(log n)
   - Nested iterations: O(n²)
   - Tree/Graph traversal: O(V+E)