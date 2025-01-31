/*

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

*/

// File: treetraversal/traversal.go

package treetraversal

// TreeNode represents a node in a binary tree
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

// InorderTraversal performs inorder traversal (Left-Root-Right)
// Time Complexity: O(n)
// Space Complexity: O(h) where h is height of tree
func InorderTraversal(root *TreeNode) []int {
    var result []int
    inorderHelper(root, &result)
    return result
}

func inorderHelper(node *TreeNode, result *[]int) {
    if node == nil {
        return
    }
    inorderHelper(node.Left, result)
    *result = append(*result, node.Val)
    inorderHelper(node.Right, result)
}

// InorderIterative performs inorder traversal iteratively
// Time Complexity: O(n)
// Space Complexity: O(h) where h is height of tree
func InorderIterative(root *TreeNode) []int {
    result := make([]int, 0)
    stack := make([]*TreeNode, 0)
    current := root

    for current != nil || len(stack) > 0 {
        // Reach the leftmost node
        for current != nil {
            stack = append(stack, current)
            current = current.Left
        }

        // Process current node and move to right
        current = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        result = append(result, current.Val)
        current = current.Right
    }

    return result
}

// PreorderTraversal performs preorder traversal (Root-Left-Right)
// Time Complexity: O(n)
// Space Complexity: O(h)
func PreorderTraversal(root *TreeNode) []int {
    var result []int
    preorderHelper(root, &result)
    return result
}

func preorderHelper(node *TreeNode, result *[]int) {
    if node == nil {
        return
    }
    *result = append(*result, node.Val)
    preorderHelper(node.Left, result)
    preorderHelper(node.Right, result)
}

// PostorderTraversal performs postorder traversal (Left-Right-Root)
// Time Complexity: O(n)
// Space Complexity: O(h)
func PostorderTraversal(root *TreeNode) []int {
    var result []int
    postorderHelper(root, &result)
    return result
}

func postorderHelper(node *TreeNode, result *[]int) {
    if node == nil {
        return
    }
    postorderHelper(node.Left, result)
    postorderHelper(node.Right, result)
    *result = append(*result, node.Val)
}

// LevelOrderTraversal performs level order traversal
// Time Complexity: O(n)
// Space Complexity: O(w) where w is maximum width of tree
func LevelOrderTraversal(root *TreeNode) [][]int {
    result := make([][]int, 0)
    if root == nil {
        return result
    }

    queue := []*TreeNode{root}

    for len(queue) > 0 {
        levelSize := len(queue)
        currentLevel := make([]int, 0)

        for i := 0; i < levelSize; i++ {
            node := queue[0]
            queue = queue[1:]
            currentLevel = append(currentLevel, node.Val)

            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }

        result = append(result, currentLevel)
    }

    return result
}

// ZigzagLevelOrder performs zigzag level order traversal
// Time Complexity: O(n)
// Space Complexity: O(w)
func ZigzagLevelOrder(root *TreeNode) [][]int {
    result := make([][]int, 0)
    if root == nil {
        return result
    }

    queue := []*TreeNode{root}
    leftToRight := true

    for len(queue) > 0 {
        levelSize := len(queue)
        currentLevel := make([]int, levelSize)

        for i := 0; i < levelSize; i++ {
            node := queue[0]
            queue = queue[1:]

            // Fill current level based on direction
            if leftToRight {
                currentLevel[i] = node.Val
            } else {
                currentLevel[levelSize-1-i] = node.Val
            }

            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }

        result = append(result, currentLevel)
        leftToRight = !leftToRight
    }

    return result
}

// BoundaryTraversal returns boundary of binary tree
// Time Complexity: O(n)
// Space Complexity: O(n)
func BoundaryTraversal(root *TreeNode) []int {
    result := make([]int, 0)
    if root == nil {
        return result
    }

    // Add root
    result = append(result, root.Val)
    if root.Left == nil && root.Right == nil {
        return result
    }

    // Add left boundary
    addLeftBoundary(root.Left, &result)

    // Add leaves
    addLeaves(root, &result)

    // Add right boundary
    rightBoundary := make([]int, 0)
    addRightBoundary(root.Right, &rightBoundary)
    result = append(result, rightBoundary...)

    return result
}

func addLeftBoundary(node *TreeNode, result *[]int) {
    if node == nil || (node.Left == nil && node.Right == nil) {
        return
    }
    *result = append(*result, node.Val)
    if node.Left != nil {
        addLeftBoundary(node.Left, result)
    } else {
        addLeftBoundary(node.Right, result)
    }
}

func addLeaves(node *TreeNode, result *[]int) {
    if node == nil {
        return
    }
    if node.Left == nil && node.Right == nil {
        *result = append(*result, node.Val)
        return
    }
    addLeaves(node.Left, result)
    addLeaves(node.Right, result)
}

func addRightBoundary(node *TreeNode, result *[]int) {
    if node == nil || (node.Left == nil && node.Right == nil) {
        return
    }
    if node.Right != nil {
        addRightBoundary(node.Right, result)
    } else {
        addRightBoundary(node.Left, result)
    }
    *result = append(*result, node.Val)
}

// File: treetraversal/traversal_test.go

package treetraversal

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

func TestInorderTraversal(t *testing.T) {
    tests := []struct {
        name     string
        root     *TreeNode
        expected []int
    }{
        {
            "Sample tree",
            createSampleTree(),
            []int{4, 2, 5, 1, 3, 6},
        },
        {
            "Empty tree",
            nil,
            []int{},
        },
    }

    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            result := InorderTraversal(test.root)
            if !reflect.DeepEqual(result, test.expected) {
                t.Errorf("Expected %v but got %v", test.expected, result)
            }
        })
    }
}

func TestPreorderTraversal(t *testing.T) {
    tests := []struct {
        name     string
        root     *TreeNode
        expected []int
    }{
        {
            "Sample tree",
            createSampleTree(),
            []int{1, 2, 4, 5, 3, 6},
        },
        {
            "Empty tree",
            nil,
            []int{},
        },
    }

    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            result := PreorderTraversal(test.root)
            if !reflect.DeepEqual(result, test.expected) {
                t.Errorf("Expected %v but got %v", test.expected, result)
            }
        })
    }
}

func TestPostorderTraversal(t *testing.T) {
    tests := []struct {
        name     string
        root     *TreeNode
        expected []int
    }{
        {
            "Sample tree",
            createSampleTree(),
            []int{4, 5, 2, 6, 3, 1},
        },
        {
            "Empty tree",
            nil,
            []int{},
        },
    }

    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            result := PostorderTraversal(test.root)
            if !reflect.DeepEqual(result, test.expected) {
                t.Errorf("Expected %v but got %v", test.expected, result)
            }
        })
    }
}

func TestLevelOrderTraversal(t *testing.T) {
    tests := []struct {
        name     string
        root     *TreeNode
        expected [][]int
    }{
        {
            "Sample tree",
            createSampleTree(),
            [][]int{{1}, {2, 3}, {4, 5, 6}},
        },
        {
            "Empty tree",
            nil,
            [][]int{},
        },
    }

    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            result := LevelOrderTraversal(test.root)
            if !reflect.DeepEqual(result, test.expected) {
                t.Errorf("Expected %v but got %v", test.expected, result)
            }
        })
    }
}

func TestZigzagLevelOrder(t *testing.T) {
    tests := []struct {
        name     string
        root     *TreeNode
        expected [][]int
    }{
        {
            "Sample tree",
            createSampleTree(),
            [][]int{{1}, {3, 2}, {4, 5, 6}},
        },
        {
            "Empty tree",
            nil,
            [][]int{},
        },
    }

    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            result := ZigzagLevelOrder(test.root)
            if !reflect.DeepEqual(result, test.expected) {
                t.Errorf("Expected %v but got %v", test.expected, result)
            }
        })
    }
}

func TestBoundaryTraversal(t *testing.T) {
    tests := []struct {
        name     string
        root     *TreeNode
        expected []int
    }{
        {
            "Sample tree",
            createSampleTree(),
            []int{1, 2, 4, 5, 6, 3},
        },
        {
            "Empty tree",
            nil,
            []int{},
        },
    }

    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            result := BoundaryTraversal(test.root)
            if !reflect.DeepEqual(result, test.expected) {
                t.Errorf("Expected %v but got %v", test.expected, result)
            }
        })
    }
}