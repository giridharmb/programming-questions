/*

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

*/

// File: bfs/bfs.go

package bfs

// TreeNode represents a node in a binary tree
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

// MinDepth finds minimum depth of binary tree
// Time Complexity: O(n)
// Space Complexity: O(w) where w is max width of tree
func MinDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    
    queue := []*TreeNode{root}
    depth := 1
    
    for len(queue) > 0 {
        levelSize := len(queue)
        
        for i := 0; i < levelSize; i++ {
            node := queue[0]
            queue = queue[1:]
            
            // Check if leaf node
            if node.Left == nil && node.Right == nil {
                return depth
            }
            
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
        depth++
    }
    
    return depth
}

// RightSideView returns right side view of binary tree
// Time Complexity: O(n)
// Space Complexity: O(w)
func RightSideView(root *TreeNode) []int {
    result := make([]int, 0)
    if root == nil {
        return result
    }
    
    queue := []*TreeNode{root}
    
    for len(queue) > 0 {
        levelSize := len(queue)
        
        for i := 0; i < levelSize; i++ {
            node := queue[0]
            queue = queue[1:]
            
            // If last node in level, add to result
            if i == levelSize-1 {
                result = append(result, node.Val)
            }
            
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
    }
    
    return result
}

// AverageLevels returns average value of nodes at each level
// Time Complexity: O(n)
// Space Complexity: O(w)
func AverageLevels(root *TreeNode) []float64 {
    result := make([]float64, 0)
    if root == nil {
        return result
    }
    
    queue := []*TreeNode{root}
    
    for len(queue) > 0 {
        levelSize := len(queue)
        levelSum := 0
        
        for i := 0; i < levelSize; i++ {
            node := queue[0]
            queue = queue[1:]
            levelSum += node.Val
            
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
        
        result = append(result, float64(levelSum)/float64(levelSize))
    }
    
    return result
}

// Graph represents a graph using adjacency list
type Graph struct {
    Vertices int
    AdjList  [][]int
}

// NewGraph creates a new graph
func NewGraph(vertices int) *Graph {
    return &Graph{
        Vertices: vertices,
        AdjList:  make([][]int, vertices),
    }
}

// AddEdge adds an edge to the graph
func (g *Graph) AddEdge(v1, v2 int) {
    g.AdjList[v1] = append(g.AdjList[v1], v2)
    g.AdjList[v2] = append(g.AdjList[v2], v1) // For undirected graph
}

// ShortestPath finds shortest path between source and destination
// Time Complexity: O(V + E)
// Space Complexity: O(V)
func (g *Graph) ShortestPath(source, dest int) []int {
    if source == dest {
        return []int{source}
    }
    
    visited := make([]bool, g.Vertices)
    parent := make([]int, g.Vertices)
    for i := range parent {
        parent[i] = -1
    }
    
    // BFS
    queue := []int{source}
    visited[source] = true
    
    for len(queue) > 0 {
        current := queue[0]
        queue = queue[1:]
        
        for _, neighbor := range g.AdjList[current] {
            if !visited[neighbor] {
                visited[neighbor] = true
                parent[neighbor] = current
                queue = append(queue, neighbor)
                
                if neighbor == dest {
                    return reconstructPath(parent, source, dest)
                }
            }
        }
    }
    
    return []int{} // No path found
}

// reconstructPath helper function to build path from parent array
func reconstructPath(parent []int, source, dest int) []int {
    path := make([]int, 0)
    current := dest
    
    for current != -1 {
        path = append([]int{current}, path...)
        current = parent[current]
    }
    
    return path
}

// WordLadder finds length of shortest transformation sequence
func WordLadder(beginWord string, endWord string, wordList []string) int {
    // Create word set for O(1) lookup
    wordSet := make(map[string]bool)
    for _, word := range wordList {
        wordSet[word] = true
    }
    
    if !wordSet[endWord] {
        return 0
    }
    
    queue := []string{beginWord}
    visited := make(map[string]bool)
    visited[beginWord] = true
    transformations := 1
    
    for len(queue) > 0 {
        levelSize := len(queue)
        
        for i := 0; i < levelSize; i++ {
            word := queue[0]
            queue = queue[1:]
            
            // Try all possible transformations
            wordChars := []byte(word)
            for j := 0; j < len(wordChars); j++ {
                original := wordChars[j]
                for c := byte('a'); c <= 'z'; c++ {
                    wordChars[j] = c
                    transformed := string(wordChars)
                    
                    if transformed == endWord {
                        return transformations + 1
                    }
                    
                    if wordSet[transformed] && !visited[transformed] {
                        visited[transformed] = true
                        queue = append(queue, transformed)
                    }
                }
                wordChars[j] = original
            }
        }
        transformations++
    }
    
    return 0
}

// ConnectedComponents finds all connected components in undirected graph
// Time Complexity: O(V + E)
// Space Complexity: O(V)
func (g *Graph) ConnectedComponents() [][]int {
    components := make([][]int, 0)
    visited := make([]bool, g.Vertices)
    
    for v := 0; v < g.Vertices; v++ {
        if !visited[v] {
            component := make([]int, 0)
            queue := []int{v}
            visited[v] = true
            
            for len(queue) > 0 {
                current := queue[0]
                queue = queue[1:]
                component = append(component, current)
                
                for _, neighbor := range g.AdjList[current] {
                    if !visited[neighbor] {
                        visited[neighbor] = true
                        queue = append(queue, neighbor)
                    }
                }
            }
            
            components = append(components, component)
        }
    }
    
    return components
}

// File: bfs/bfs_test.go

package bfs

import (
    "reflect"
    "testing"
)

// Helper function to create a sample tree
func createSampleTree() *TreeNode {
    /*
           1
          / \
         2   3
        / \   \
       4   5   6
    */
    root := &TreeNode{Val: 1}
    root.Left = &TreeNode{Val: 2}
    root.Right = &TreeNode{Val: 3}
    root.Left.Left = &TreeNode{Val: 4}
    root.Left.Right = &TreeNode{Val: 5}
    root.Right.Right = &TreeNode{Val: 6}
    return root
}

func TestMinDepth(t *testing.T) {
    tests := []struct {
        name     string
        root     *TreeNode
        expected int
    }{
        {"Sample tree", createSampleTree(), 2},
        {"Empty tree", nil, 0},
        {"Single node", &TreeNode{Val: 1}, 1},
    }
    
    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            result := MinDepth(test.root)
            if result != test.expected {
                t.Errorf("Expected %d but got %d", test.expected, result)
            }
        })
    }
}

func TestRightSideView(t *testing.T) {
    tests := []struct {
        name     string
        root     *TreeNode
        expected []int
    }{
        {"Sample tree", createSampleTree(), []int{1, 3, 6}},
        {"Empty tree", nil, []int{}},
        {"Single node", &TreeNode{Val: 1}, []int{1}},
    }
    
    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            result := RightSideView(test.root)
            if !reflect.DeepEqual(result, test.expected) {
                t.Errorf("Expected %v but got %v", test.expected, result)
            }
        })
    }
}

func TestAverageLevels(t *testing.T) {
    tests := []struct {
        name     string
        root     *TreeNode
        expected []float64
    }{
        {
            "Sample tree",
            createSampleTree(),
            []float64{1.0, 2.5, 5.0},
        },
        {"Empty tree", nil, []float64{}},
    }
    
    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            result := AverageLevels(test.root)
            if !reflect.DeepEqual(result, test.expected) {
                t.Errorf("Expected %v but got %v", test.expected, result)
            }
        })
    }
}

func TestShortestPath(t *testing.T) {
    g := NewGraph(6)
    g.AddEdge(0, 1)
    g.AddEdge(1, 2)
    g.AddEdge(2, 3)
    g.AddEdge(3, 4)
    g.AddEdge(4, 5)
    g.AddEdge(0, 2)
    
    tests := []struct {
        name     string
        source   int
        dest     int
        expected []int
    }{
        {"Direct path", 0, 1, []int{0, 1}},
        {"Longer path", 0, 5, []int{0, 2, 3, 4, 5}},
        {"Same vertex", 1, 1, []int{1}},
    }
    
    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            result := g.ShortestPath(test.source, test.dest)
            if !reflect.DeepEqual(result, test.expected) {
                t.Errorf("Expected %v but got %v", test.expected, result)
            }
        })
    }
}

func TestWordLadder(t *testing.T) {
    tests := []struct {
        name      string
        beginWord string
        endWord   string
        wordList  []string
        expected  int
    }{
        {
            "Valid transformation",
            "hit",
            "cog",
            []string{"hot", "dot", "dog", "lot", "log", "cog"},
            5,
        },
        {
            "No transformation",
            "hit",
            "cog",
            []string{"hot", "dot", "dog", "lot", "log"},
            0,
        },
    }
    
    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            result := WordLadder(test.beginWord, test.endWord, test.wordList)
            if result != test.expected {
                t.Errorf("Expected %d but got %d", test.expected, result)
            }
        })
    }
}

func TestConnectedComponents(t *testing.T) {
    g := NewGraph(7)
    g.AddEdge(0, 1)
    g.AddEdge(1, 2)
    g.AddEdge(3, 4)
    g.AddEdge(5, 6)
    
    expected := [][]int{
        {0, 1, 2},
        {3, 4},
        {5, 6},
    }
    
    result := g.ConnectedComponents()
    if len(result) != len(expected) {
        t.Errorf("Expected %d components but got %d", len(expected), len(result))
    }
    
    // Check each component (order may vary)
    foundComponents := make(map[int]bool)
    for _, comp := range result {
        size := len(comp)
        foundComponents[size] = true
    }
    
    for _, comp := range expected {
        if !foundComponents[len(comp)] {
            t.Errorf("Missing component of size %d", len(comp))
        }
    }
}