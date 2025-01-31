/*

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

*/

// File: dfs/dfs.go

package dfs

// TreeNode represents a node in a binary tree
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

// PathSum determines if root-to-leaf path with given sum exists
// Time Complexity: O(n)
// Space Complexity: O(h) where h is height of tree
func PathSum(root *TreeNode, targetSum int) bool {
    if root == nil {
        return false
    }
    
    // If leaf node, check if remaining sum equals node value
    if root.Left == nil && root.Right == nil {
        return targetSum == root.Val
    }
    
    // Recursively check left and right subtrees
    return PathSum(root.Left, targetSum-root.Val) || 
           PathSum(root.Right, targetSum-root.Val)
}

// AllPathSum finds all root-to-leaf paths with given sum
// Time Complexity: O(n)
// Space Complexity: O(h)
func AllPathSum(root *TreeNode, targetSum int) [][]int {
    result := make([][]int, 0)
    currentPath := make([]int, 0)
    findPaths(root, targetSum, currentPath, &result)
    return result
}

func findPaths(node *TreeNode, targetSum int, currentPath []int, result *[][]int) {
    if node == nil {
        return
    }
    
    // Add current node to path
    currentPath = append(currentPath, node.Val)
    
    // Check if leaf node and sum matches
    if node.Left == nil && node.Right == nil && targetSum == node.Val {
        // Create a copy of current path
        pathCopy := make([]int, len(currentPath))
        copy(pathCopy, currentPath)
        *result = append(*result, pathCopy)
    }
    
    // Recursively search left and right subtrees
    findPaths(node.Left, targetSum-node.Val, currentPath, result)
    findPaths(node.Right, targetSum-node.Val, currentPath, result)
}

// CountPaths counts number of paths with given sum
// Paths can start and end anywhere
// Time Complexity: O(n)
// Space Complexity: O(h)
func CountPaths(root *TreeNode, targetSum int) int {
    return countPathsHelper(root, targetSum, []int{})
}

func countPathsHelper(node *TreeNode, targetSum int, path []int) int {
    if node == nil {
        return 0
    }
    
    // Add current node to path
    path = append(path, node.Val)
    count := 0
    sum := 0
    
    // Check all possible paths ending at current node
    for i := len(path) - 1; i >= 0; i-- {
        sum += path[i]
        if sum == targetSum {
            count++
        }
    }
    
    // Recursively count paths in left and right subtrees
    count += countPathsHelper(node.Left, targetSum, path)
    count += countPathsHelper(node.Right, targetSum, path)
    
    return count
}

// Graph represented as adjacency list
type Graph struct {
    Vertices int
    AdjList  [][]int
}

// NewGraph creates a new graph with given number of vertices
func NewGraph(vertices int) *Graph {
    return &Graph{
        Vertices: vertices,
        AdjList:  make([][]int, vertices),
    }
}

// AddEdge adds an edge to the graph
func (g *Graph) AddEdge(v1, v2 int) {
    g.AdjList[v1] = append(g.AdjList[v1], v2)
}

// HasPath determines if path exists between source and destination
// Time Complexity: O(V + E)
// Space Complexity: O(V)
func (g *Graph) HasPath(source, dest int) bool {
    visited := make([]bool, g.Vertices)
    return g.dfsPath(source, dest, visited)
}

func (g *Graph) dfsPath(current, dest int, visited []bool) bool {
    if current == dest {
        return true
    }
    
    visited[current] = true
    
    for _, neighbor := range g.AdjList[current] {
        if !visited[neighbor] {
            if g.dfsPath(neighbor, dest, visited) {
                return true
            }
        }
    }
    
    return false
}

// AllPaths finds all paths between source and destination
// Time Complexity: O(V^V)
// Space Complexity: O(V)
func (g *Graph) AllPaths(source, dest int) [][]int {
    result := make([][]int, 0)
    visited := make([]bool, g.Vertices)
    currentPath := []int{source}
    g.findAllPaths(source, dest, visited, currentPath, &result)
    return result
}

func (g *Graph) findAllPaths(current, dest int, visited []bool, 
    currentPath []int, result *[][]int) {
    
    if current == dest {
        pathCopy := make([]int, len(currentPath))
        copy(pathCopy, currentPath)
        *result = append(*result, pathCopy)
        return
    }
    
    visited[current] = true
    
    for _, neighbor := range g.AdjList[current] {
        if !visited[neighbor] {
            currentPath = append(currentPath, neighbor)
            g.findAllPaths(neighbor, dest, visited, currentPath, result)
            currentPath = currentPath[:len(currentPath)-1]
        }
    }
    
    visited[current] = false
}

// Clone performs deep copy of a graph
type Node struct {
    Val       int
    Neighbors []*Node
}

func CloneGraph(node *Node) *Node {
    if node == nil {
        return nil
    }
    
    visited := make(map[*Node]*Node)
    return cloneHelper(node, visited)
}

func cloneHelper(node *Node, visited map[*Node]*Node) *Node {
    if node == nil {
        return nil
    }
    
    // If already visited, return the clone
    if clone, exists := visited[node]; exists {
        return clone
    }
    
    // Create new node and add to visited map
    clone := &Node{Val: node.Val, Neighbors: make([]*Node, 0)}
    visited[node] = clone
    
    // Clone all neighbors
    for _, neighbor := range node.Neighbors {
        clone.Neighbors = append(clone.Neighbors, cloneHelper(neighbor, visited))
    }
    
    return clone
}

// File: dfs/dfs_test.go

package dfs

import (
    "reflect"
    "testing"
)

// Helper function to create a sample tree
func createSampleTree() *TreeNode {
    /*
           10
          /  \
         5    15
        / \     \
       3   7     18
    */
    root := &TreeNode{Val: 10}
    root.Left = &TreeNode{Val: 5}
    root.Right = &TreeNode{Val: 15}
    root.Left.Left = &TreeNode{Val: 3}
    root.Left.Right = &TreeNode{Val: 7}
    root.Right.Right = &TreeNode{Val: 18}
    return root
}

func TestPathSum(t *testing.T) {
    tests := []struct {
        name      string
        root      *TreeNode
        targetSum int
        expected  bool
    }{
        {
            "Path exists",
            createSampleTree(),
            18,
            true,
        },
        {
            "Path does not exist",
            createSampleTree(),
            100,
            false,
        },
        {
            "Empty tree",
            nil,
            0,
            false,
        },
    }

    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            result := PathSum(test.root, test.targetSum)
            if result != test.expected {
                t.Errorf("Expected %v but got %v", test.expected, result)
            }
        })
    }
}

func TestAllPathSum(t *testing.T) {
    tests := []struct {
        name      string
        root      *TreeNode
        targetSum int
        expected  [][]int
    }{
        {
            "Multiple paths",
            createSampleTree(),
            18,
            [][]int{{10, 5, 3}, {10, 8}},
        },
        {
            "No paths",
            createSampleTree(),
            100,
            [][]int{},
        },
    }

    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            result := AllPathSum(test.root, test.targetSum)
            if !reflect.DeepEqual(result, test.expected) {
                t.Errorf("Expected %v but got %v", test.expected, result)
            }
        })
    }
}

func TestCountPaths(t *testing.T) {
    tests := []struct {
        name      string
        root      *TreeNode
        targetSum int
        expected  int
    }{
        {
            "Multiple paths",
            createSampleTree(),
            18,
            3,
        },
        {
            "No paths",
            createSampleTree(),
            100,
            0,
        },
    }

    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            result := CountPaths(test.root, test.targetSum)
            if result != test.expected {
                t.Errorf("Expected %d but got %d", test.expected, result)
            }
        })
    }
}

func TestGraphHasPath(t *testing.T) {
    g := NewGraph(4)
    g.AddEdge(0, 1)
    g.AddEdge(1, 2)
    g.AddEdge(2, 3)
    
    tests := []struct {
        name     string
        source   int
        dest     int
        expected bool
    }{
        {"Path exists", 0, 3, true},
        {"No path", 3, 0, false},
        {"Same vertex", 1, 1, true},
    }

    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            result := g.HasPath(test.source, test.dest)
            if result != test.expected {
                t.Errorf("Expected %v but got %v", test.expected, result)
            }
        })
    }
}

func TestGraphAllPaths(t *testing.T) {
    g := NewGraph(4)
    g.AddEdge(0, 1)
    g.AddEdge(1, 2)
    g.AddEdge(2, 3)
    g.AddEdge(0, 2)
    
    tests := []struct {
        name     string
        source   int
        dest     int
        expected [][]int
    }{
        {
            "Multiple paths",
            0,
            3,
            [][]int{{0, 1, 2, 3}, {0, 2, 3}},
        },
        {
            "No paths",
            3,
            0,
            [][]int{},
        },
    }

    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            result := g.AllPaths(test.source, test.dest)
            if !reflect.DeepEqual(result, test.expected) {
                t.Errorf("Expected %v but got %v", test.expected, result)
            }
        })
    }
}
