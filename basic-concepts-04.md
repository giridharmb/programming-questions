# 4. Quick Reference Tables and Guidelines

## Common Data Structure Operations

### Array/Slice Operations
| Operation               | Average Case | Worst Case | Space  |
|------------------------|--------------|------------|--------|
| Access                 | O(1)         | O(1)       | O(1)   |
| Search                 | O(n)         | O(n)       | O(1)   |
| Insert/Delete at end   | O(1)         | O(1)       | O(1)   |
| Insert/Delete at start | O(n)         | O(n)       | O(1)   |
| Insert/Delete at index | O(n)         | O(n)       | O(1)   |

### Map Operations
| Operation     | Average Case | Worst Case | Space  |
|---------------|--------------|------------|--------|
| Insert        | O(1)         | O(n)       | O(n)   |
| Delete        | O(1)         | O(n)       | O(1)   |
| Search        | O(1)         | O(n)       | O(1)   |
| Iteration     | O(n)         | O(n)       | O(1)   |

### Heap Operations
| Operation     | Time Complexity | Space   |
|---------------|----------------|---------|
| Insert        | O(log n)       | O(1)    |
| Delete Max    | O(log n)       | O(1)    |
| Get Max       | O(1)           | O(1)    |
| Build Heap    | O(n)           | O(1)    |

## Sorting Algorithms

| Algorithm      | Best Case  | Average Case | Worst Case | Space      |
|---------------|------------|--------------|------------|------------|
| Quick Sort    | O(n log n) | O(n log n)   | O(n²)      | O(log n)   |
| Merge Sort    | O(n log n) | O(n log n)   | O(n log n) | O(n)       |
| Heap Sort     | O(n log n) | O(n log n)   | O(n log n) | O(1)       |
| Bubble Sort   | O(n)       | O(n²)        | O(n²)      | O(1)       |
| Insertion Sort| O(n)       | O(n²)        | O(n²)      | O(1)       |
| Selection Sort| O(n²)      | O(n²)        | O(n²)      | O(1)       |

## Search Algorithms

| Algorithm        | Time Complexity | Space Complexity | Prerequisites    |
|-----------------|----------------|------------------|------------------|
| Linear Search   | O(n)           | O(1)             | None            |
| Binary Search   | O(log n)       | O(1)             | Sorted array    |
| DFS             | O(V + E)       | O(V)             | Graph/Tree      |
| BFS             | O(V + E)       | O(V)             | Graph/Tree      |
| Dijkstra        | O((V+E)log V)  | O(V)             | Weighted Graph  |

## Common Graph Algorithms

| Algorithm           | Time Complexity | Space Complexity |
|--------------------|----------------|------------------|
| DFS                | O(V + E)       | O(V)             |
| BFS                | O(V + E)       | O(V)             |
| Dijkstra           | O((V+E)log V)  | O(V)             |
| Bellman-Ford       | O(VE)          | O(V)             |
| Floyd-Warshall     | O(V³)          | O(V²)            |
| Prim's MST         | O((V+E)log V)  | O(V)             |
| Kruskal's MST      | O(E log E)     | O(V)             |

## Common Problem-Solving Patterns

### Sliding Window
```go
func slidingWindowPattern(arr []int, k int) {
    windowSum := 0
    windowStart := 0
    
    for windowEnd := 0; windowEnd < len(arr); windowEnd++ {
        windowSum += arr[windowEnd]      // Add element
        
        if windowEnd >= k-1 {
            // Process window
            windowSum -= arr[windowStart] // Remove element
            windowStart++                 // Slide window
        }
    }
}
// Time: O(n), Space: O(1)
```

### Two Pointers
```go
func twoPointerPattern(arr []int) {
    left, right := 0, len(arr)-1
    
    for left < right {
        // Process elements at left and right
        // Move pointers based on condition
        if condition {
            left++
        } else {
            right--
        }
    }
}
// Time: O(n), Space: O(1)
```

### Fast & Slow Pointers
```go
func fastSlowPattern(head *ListNode) {
    slow, fast := head, head
    
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    // Slow is at middle/meeting point
}
// Time: O(n), Space: O(1)
```

## Memory Complexity Guidelines

### Stack Space
- Recursive calls: O(depth)
- System stack size is limited
- Consider iterative alternatives

### Heap Space
- Dynamic allocations
- Objects, arrays, maps
- Persists until garbage collected

### Cache Considerations
- L1/L2/L3 cache effects
- Memory locality
- Array vs linked structures

## Common Optimization Techniques

1. **Space-Time Tradeoffs**
   - Caching/Memoization
   - Precomputation
   - Hash tables for lookups

2. **Algorithm Choice**
   - Consider input size
   - Consider constraints
   - Consider requirements

3. **Data Structure Choice**
   - Access patterns
   - Operation frequency
   - Memory constraints

4. **Code Optimization**
   - Avoid unnecessary allocations
   - Use efficient loops
   - Consider compiler optimizations

## Performance Tips

1. **When to Use Arrays**
   - Fixed size
   - Random access needed
   - Cache locality important

2. **When to Use Maps**
   - Key-value lookups
   - Dynamic size
   - Unique keys needed

3. **When to Use Recursion**
   - Tree/Graph traversal
   - Divide and conquer
   - Simple implementation needed

4. **When to Use Iteration**
   - Stack space limited
   - Simple linear processing
   - Better performance needed