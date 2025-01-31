## Programming Questions

## [Time and Space Complexity](https://github.com/giridharmb/programming-questions/blob/master/basic.md)

- [Basic Concepts of Time and Space Complexity](https://github.com/giridharmb/programming-questions/blob/master/basic-concepts-01.md)
- [Step-by-Step Guide for Calculating Complexity](https://github.com/giridharmb/programming-questions/blob/master/basic-concepts-02.md)
- [Practical Examples with Complexity Analysis](https://github.com/giridharmb/programming-questions/blob/master/basic-concepts-03.md)
- [Quick Reference Tables and Guidelines](https://github.com/giridharmb/programming-questions/blob/master/basic-concepts-04.md)
- [Optimization Tips and Best Practices](https://github.com/giridharmb/programming-questions/blob/master/basic-concepts-05.md)
- [Common Pitfalls in Time and Space Complexity](https://github.com/giridharmb/programming-questions/blob/master/basic-concepts-06.md)

## [Basic Leet Code Questions](https://github.com/giridharmb/programming-questions/blob/master/fundamentals.md)

## Code

#### 1: Prefix & Sum

```
This implementation includes three common prefix sum pattern problems:

Range Sum Query - Using prefix sums for O(1) range queries
Subarray Sum Equals K - Finding continuous subarrays that sum to k
Product of Array Except Self - Using prefix/suffix products

Key features:

Complete test coverage
Time and space complexity analysis
Clean Go idioms and error handling
Efficient memory usage
```

#### [Code](https://github.com/giridharmb/programming-questions/blob/master/01-prefix-sum.go)

#### 2: Two Pointers

```
This implementation includes four classic Two Pointer pattern problems:

Two Sum - Finding two numbers that sum to target in sorted array
Three Sum - Finding all unique triplets that sum to zero
Remove Duplicates - Removing duplicates from sorted array in-place
Container With Most Water - Finding maximum area between vertical lines

Key features:

Comprehensive test coverage
Time/space complexity analysis for each function
Efficient in-place operations
Handle edge cases and duplicates
Clean Go style with helper functions
```

#### [Code](https://github.com/giridharmb/programming-questions/blob/master/02-two-pointers.go)

#### 3: Sliding Window

```
This implementation includes four classic Sliding Window pattern problems:

Maximum Sum Subarray of size K - Fixed-size window
Longest Substring with K Distinct Characters - Variable-size window with frequency map
Minimum Size Subarray Sum - Variable-size window with target sum
Find All Anagrams - Fixed-size window with character frequency matching

Key features:

Both fixed and variable size window implementations
Efficient frequency map usage
Comprehensive test cases
Clear complexity analysis
Clean Go idioms and error handling
```

#### [Code](https://github.com/giridharmb/programming-questions/blob/master/03-sliding-window.go)

#### 4: Linked List & Fast Slow Pointers

```
This implementation includes five classic Fast & Slow Pointers pattern problems:

Detect Cycle - Determine if a linked list has a cycle
Find Cycle Start - Find the node where the cycle begins
Find Middle - Find the middle node of a linked list
Check Palindrome - Check if a linked list is palindrome
Reorder List - Reorder list to L0→Ln→L1→Ln-1→L2→Ln-2

Key features:

Complete LinkedList node implementation
Helper functions for testing
Cycle detection and handling
In-place operations for O(1) space complexity
Comprehensive test cases with edge cases
Clean Go idioms
```

#### [Code](https://github.com/giridharmb/programming-questions/blob/master/04-linked-list-fast-slow-pointers.go)

#### 5: Linked List Reversal

```
This implementation includes four variations of the LinkedList reversal pattern:

Reverse Entire List - Basic complete reversal
Reverse Sublist - Reverse portion between given positions
Reverse K-Group - Reverse every k nodes
Reverse Alternate K-Group - Reverse alternate groups of k nodes

Key features:

In-place reversal operations
Comprehensive test suite
Edge case handling
Helper functions for testing
Multiple reversal variations
Detailed time/space complexity analysis
```

#### [Code](https://github.com/giridharmb/programming-questions/blob/master/05-linked-list-reversal.go)

#### 6: Monotonic Stack

```
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
```

#### [Code](https://github.com/giridharmb/programming-questions/blob/master/06-monotonic-stack.go)

#### 7: Top-K Elements

```
This implementation includes four classic Top 'K' Elements problems:

Kth Largest Element - Find kth largest element in array
Top K Frequent Elements - Find k most frequent elements
K Closest Points - Find k closest points to origin
Kth Smallest in Matrix - Find kth smallest element in sorted matrix

Key features:

Heap implementations using Go's container/heap interface
Custom heap types for different problems
Min/Max heap variations
Efficient priority queue operations
Comprehensive test coverage
Clean Go interfaces
```

#### [Code](https://github.com/giridharmb/programming-questions/blob/master/07-top-k-elements.go)

#### 8: Overlapping Intervals

```
This implementation includes five classic Overlapping Intervals problems:

Merge Intervals - Merge all overlapping intervals
Insert Interval - Insert and merge a new interval into existing intervals
Interval Intersection - Find intersection of two interval lists
Meeting Rooms I - Check if a person can attend all meetings
Meeting Rooms II - Find minimum number of meeting rooms required

Key features:

Interval structure representation
Efficient sorting-based solutions
Two-pointer techniques
Chronological event processing
Comprehensive test coverage
Clean Go idioms
```

#### [Code](https://github.com/giridharmb/programming-questions/blob/master/08-overlapping-intervals.go)

#### 9: Binary Search

```
This implementation includes six classic Modified Binary Search problems:

Standard Binary Search - Basic binary search in sorted array
Search in Rotated Array - Binary search in rotated sorted array
Find Minimum in Rotated Array - Find minimum element in rotated sorted array
Search Range - Find first and last positions of target
Find Peak Element - Find peak element in array
Search in Bitonic Array - Search in array that increases then decreases

Key features:

Multiple binary search variations
Handling of rotated arrays
Finding boundaries and peaks
Search in different array patterns
Comprehensive test coverage
Clean Go idioms
```

#### [Code](https://github.com/giridharmb/programming-questions/blob/master/09-binary-search.go)

#### 10: Binary Tree Traversal

```
This implementation includes six types of binary tree traversals:

Inorder Traversal (Left-Root-Right) - Both recursive and iterative
Preorder Traversal (Root-Left-Right)
Postorder Traversal (Left-Right-Root)
Level Order Traversal (BFS)
Zigzag Level Order Traversal
Boundary Traversal

Key features:

TreeNode structure definition
Both recursive and iterative implementations
Queue-based level order traversal
Stack-based iterative traversal
Comprehensive test cases
Clean Go idioms
```

#### [Code](https://github.com/giridharmb/programming-questions/blob/master/10-binary-tree-traversal.go)

#### 11: Depth First Search + Graph + Tree

```
This implementation includes both tree and graph DFS problems:
Tree DFS Problems:

Path Sum - Find if path exists with given sum
All Path Sum - Find all paths with given sum
Count Paths - Count paths with given sum (can start and end anywhere)

Graph DFS Problems:

Has Path - Check if path exists between two vertices
All Paths - Find all possible paths between two vertices
Clone Graph - Create deep copy of graph

Key features:

Both tree and graph implementations
Path finding algorithms
Backtracking techniques
Visited node tracking
Comprehensive test cases
Clean Go idioms
```

#### [Code](https://github.com/giridharmb/programming-questions/blob/master/11-dfs-graph-tree.go)

#### 12: BFS + Graph + Tree

```
This implementation includes both tree and graph BFS problems:
Tree BFS Problems:

Minimum Depth - Find minimum depth of binary tree
Right Side View - Get right side view of binary tree
Average of Levels - Calculate average value at each level

Graph BFS Problems:

Shortest Path - Find shortest path between two vertices
Word Ladder - Find shortest transformation sequence
Connected Components - Find all connected components in graph

Key features:

Level-wise processing
Queue-based implementations
Path finding algorithms
Word transformation
Component discovery
Comprehensive test cases
```

#### [Code](https://github.com/giridharmb/programming-questions/blob/master/12-bfs-graph-tree.go)

#### 13: Matrix Traversal

```
This implementation includes five classic Matrix Traversal problems:

Number of Islands - Count distinct islands in binary matrix
Flood Fill - Perform flood fill operation starting from given cell
Max Area Island - Find island with maximum area
Pacific Atlantic Water Flow - Find cells that can flow to both oceans
Rotting Oranges - Find minimum time until all oranges rot

Key features:

DFS and BFS approaches
Direction vectors for movement
Grid boundary checking
Visited cell marking
Queue-based processing
Comprehensive test cases
```

#### [Code](https://github.com/giridharmb/programming-questions/blob/master/13-matrix-traversal.go)

#### 14: Back Tracking

```
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
```

#### [Code](https://github.com/giridharmb/programming-questions/blob/master/14-back-tracking.go)

#### 15: Dynamic Programming

```
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
```

#### [Code](https://github.com/giridharmb/programming-questions/blob/master/15-dynamic-programming.go)