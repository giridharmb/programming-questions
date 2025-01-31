/*

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

*/

// File: linkedlistreversal/reversal.go

package linkedlistreversal

// ListNode represents a node in a linked list
type ListNode struct {
    Val  int
    Next *ListNode
}

// ReverseList reverses an entire linked list
// Time Complexity: O(n)
// Space Complexity: O(1)
func ReverseList(head *ListNode) *ListNode {
    var prev *ListNode
    current := head
    
    for current != nil {
        nextTemp := current.Next
        current.Next = prev
        prev = current
        current = nextTemp
    }
    
    return prev
}

// ReverseSubList reverses a sublist between positions left and right (1-based indexing)
// Time Complexity: O(n)
// Space Complexity: O(1)
func ReverseSubList(head *ListNode, left, right int) *ListNode {
    if head == nil || left == right {
        return head
    }
    
    dummy := &ListNode{Next: head}
    prev := dummy
    
    // Move to the node before reversal starts
    for i := 1; i < left; i++ {
        prev = prev.Next
    }
    
    current := prev.Next
    // Reverse the sublist
    for i := 0; i < right-left; i++ {
        nextTemp := current.Next
        current.Next = nextTemp.Next
        nextTemp.Next = prev.Next
        prev.Next = nextTemp
    }
    
    return dummy.Next
}

// ReverseKGroup reverses every k nodes in the list
// Time Complexity: O(n)
// Space Complexity: O(1)
func ReverseKGroup(head *ListNode, k int) *ListNode {
    if head == nil || k <= 1 {
        return head
    }
    
    dummy := &ListNode{Next: head}
    prev := dummy
    
    for head != nil {
        // Check if remaining nodes are >= k
        tail := head
        for i := 1; i < k; i++ {
            tail = tail.Next
            if tail == nil {
                return dummy.Next
            }
        }
        
        nextHead := tail.Next
        // Reverse k nodes
        current := head
        var prevTemp *ListNode
        for i := 0; i < k; i++ {
            nextTemp := current.Next
            current.Next = prevTemp
            prevTemp = current
            current = nextTemp
        }
        
        // Connect with rest of the list
        prev.Next = tail
        head.Next = nextHead
        prev = head
        head = nextHead
    }
    
    return dummy.Next
}

// ReverseAlternateKGroup reverses alternate k nodes in the list
// Time Complexity: O(n)
// Space Complexity: O(1)
func ReverseAlternateKGroup(head *ListNode, k int) *ListNode {
    if head == nil || k <= 1 {
        return head
    }
    
    current := head
    var prev *ListNode
    var next *ListNode
    count := 0
    
    // Check if we have k nodes
    tail := current
    for i := 1; i < k && tail != nil; i++ {
        tail = tail.Next
    }
    if tail == nil {
        return head
    }
    
    // Reverse first k nodes
    while := current
    for current != nil && count < k {
        next = current.Next
        current.Next = prev
        prev = current
        current = next
        count++
    }
    
    // Connect with rest of the list
    if next != nil {
        head.Next = next
        // Skip k nodes
        count = 0
        while := next
        for count < k-1 && while.Next != nil {
            while = while.Next
            count++
        }
        if while.Next != nil {
            while.Next = ReverseAlternateKGroup(while.Next, k)
        }
    }
    
    return prev
}

// File: linkedlistreversal/reversal_test.go

package linkedlistreversal

import (
    "testing"
    "reflect"
)

// createList creates a linked list from a slice of integers
func createList(nums []int) *ListNode {
    if len(nums) == 0 {
        return nil
    }
    
    head := &ListNode{Val: nums[0]}
    current := head
    for i := 1; i < len(nums); i++ {
        current.Next = &ListNode{Val: nums[i]}
        current = current.Next
    }
    return head
}

// toSlice converts a linked list to a slice for testing
func toSlice(head *ListNode) []int {
    result := []int{}
    current := head
    for current != nil {
        result = append(result, current.Val)
        current = current.Next
    }
    return result
}

func TestReverseList(t *testing.T) {
    tests := []struct {
        input    []int
        expected []int
    }{
        {[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
        {[]int{1, 2}, []int{2, 1}},
        {[]int{1}, []int{1}},
        {[]int{}, []int{}},
    }
    
    for _, test := range tests {
        list := createList(test.input)
        result := ReverseList(list)
        resultSlice := toSlice(result)
        if !reflect.DeepEqual(resultSlice, test.expected) {
            t.Errorf("For input=%v, expected %v but got %v",
                test.input, test.expected, resultSlice)
        }
    }
}

func TestReverseSubList(t *testing.T) {
    tests := []struct {
        input    []int
        left     int
        right    int
        expected []int
    }{
        {[]int{1, 2, 3, 4, 5}, 2, 4, []int{1, 4, 3, 2, 5}},
        {[]int{5}, 1, 1, []int{5}},
        {[]int{3, 5}, 1, 2, []int{5, 3}},
    }
    
    for _, test := range tests {
        list := createList(test.input)
        result := ReverseSubList(list, test.left, test.right)
        resultSlice := toSlice(result)
        if !reflect.DeepEqual(resultSlice, test.expected) {
            t.Errorf("For input=%v, left=%d, right=%d, expected %v but got %v",
                test.input, test.left, test.right, test.expected, resultSlice)
        }
    }
}

func TestReverseKGroup(t *testing.T) {
    tests := []struct {
        input    []int
        k        int
        expected []int
    }{
        {[]int{1, 2, 3, 4, 5}, 2, []int{2, 1, 4, 3, 5}},
        {[]int{1, 2, 3, 4, 5}, 3, []int{3, 2, 1, 4, 5}},
        {[]int{1}, 1, []int{1}},
        {[]int{1, 2}, 3, []int{1, 2}},
    }
    
    for _, test := range tests {
        list := createList(test.input)
        result := ReverseKGroup(list, test.k)
        resultSlice := toSlice(result)
        if !reflect.DeepEqual(resultSlice, test.expected) {
            t.Errorf("For input=%v, k=%d, expected %v but got %v",
                test.input, test.k, test.expected, resultSlice)
        }
    }
}

func TestReverseAlternateKGroup(t *testing.T) {
    tests := []struct {
        input    []int
        k        int
        expected []int
    }{
        {[]int{1, 2, 3, 4, 5, 6, 7, 8}, 2, []int{2, 1, 3, 4, 6, 5, 7, 8}},
        {[]int{1, 2, 3, 4, 5}, 3, []int{3, 2, 1, 4, 5}},
        {[]int{1, 2}, 1, []int{1, 2}},
        {[]int{1}, 2, []int{1}},
    }
    
    for _, test := range tests {
        list := createList(test.input)
        result := ReverseAlternateKGroup(list, test.k)
        resultSlice := toSlice(result)
        if !reflect.DeepEqual(resultSlice, test.expected) {
            t.Errorf("For input=%v, k=%d, expected %v but got %v",
                test.input, test.k, test.expected, resultSlice)
        }
    }
}