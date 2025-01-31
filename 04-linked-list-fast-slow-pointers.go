/*

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

*/

// File: fastslow/fastslow.go

package fastslow

// ListNode represents a node in a linked list
type ListNode struct {
    Val  int
    Next *ListNode
}

// HasCycle determines if a linked list has a cycle
// Time Complexity: O(n)
// Space Complexity: O(1)
func HasCycle(head *ListNode) bool {
    if head == nil || head.Next == nil {
        return false
    }

    slow := head
    fast := head.Next

    for fast != nil && fast.Next != nil && slow != fast {
        slow = slow.Next
        fast = fast.Next.Next
    }

    return fast != nil && fast.Next != nil
}

// FindCycleStart finds the node where the cycle begins
// Returns nil if no cycle exists
// Time Complexity: O(n)
// Space Complexity: O(1)
func FindCycleStart(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return nil
    }

    // Find meeting point
    slow, fast := head, head
    hasCycle := false

    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
        if slow == fast {
            hasCycle = true
            break
        }
    }

    if !hasCycle {
        return nil
    }

    // Move pointer1 to head and find cycle start
    pointer1 := head
    pointer2 := slow

    for pointer1 != pointer2 {
        pointer1 = pointer1.Next
        pointer2 = pointer2.Next
    }

    return pointer1
}

// FindMiddle finds the middle node of the linked list
// For even length, returns the second middle node
// Time Complexity: O(n)
// Space Complexity: O(1)
func FindMiddle(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    slow, fast := head, head

    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }

    return slow
}

// IsPalindrome checks if a linked list is a palindrome
// Time Complexity: O(n)
// Space Complexity: O(1)
func IsPalindrome(head *ListNode) bool {
    if head == nil || head.Next == nil {
        return true
    }

    // Find middle
    slow, fast := head, head
    for fast.Next != nil && fast.Next.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }

    // Reverse second half
    secondHalf := reverseList(slow.Next)

    // Compare first and second half
    firstHalf := head
    for secondHalf != nil {
        if firstHalf.Val != secondHalf.Val {
            return false
        }
        firstHalf = firstHalf.Next
        secondHalf = secondHalf.Next
    }

    return true
}

// reverseList reverses a linked list and returns the new head
func reverseList(head *ListNode) *ListNode {
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

// ReorderList reorders list to L0 → Ln → L1 → Ln-1 → L2 → Ln-2 ...
// Time Complexity: O(n)
// Space Complexity: O(1)
func ReorderList(head *ListNode) {
    if head == nil || head.Next == nil {
        return
    }

    // Find middle
    slow, fast := head, head
    for fast.Next != nil && fast.Next.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }

    // Reverse second half
    second := reverseList(slow.Next)
    slow.Next = nil
    first := head

    // Merge lists
    for second != nil {
        temp1 := first.Next
        temp2 := second.Next
        first.Next = second
        second.Next = temp1
        first = temp1
        second = temp2
    }
}

// File: fastslow/fastslow_test.go

package fastslow

import "testing"

// createLinkedList helper function to create linked list from slice
func createLinkedList(nums []int) *ListNode {
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

// createCyclicList creates a linked list with a cycle
func createCyclicList(nums []int, pos int) *ListNode {
    if len(nums) == 0 {
        return nil
    }

    head := createLinkedList(nums)
    if pos == -1 {
        return head
    }

    // Find cycle point and last node
    var cycleNode *ListNode
    current := head
    count := 0
    for current.Next != nil {
        if count == pos {
            cycleNode = current
        }
        current = current.Next
        count++
    }
    current.Next = cycleNode
    return head
}

func TestHasCycle(t *testing.T) {
    tests := []struct {
        nums     []int
        pos      int
        expected bool
    }{
        {[]int{3, 2, 0, -4}, 1, true},
        {[]int{1, 2}, 0, true},
        {[]int{1}, -1, false},
        {[]int{}, -1, false},
    }

    for _, test := range tests {
        list := createCyclicList(test.nums, test.pos)
        result := HasCycle(list)
        if result != test.expected {
            t.Errorf("For list=%v, pos=%d, expected %v but got %v",
                test.nums, test.pos, test.expected, result)
        }
    }
}

func TestFindMiddle(t *testing.T) {
    tests := []struct {
        nums     []int
        expected int
    }{
        {[]int{1, 2, 3, 4, 5}, 3},
        {[]int{1, 2, 3, 4}, 3},
        {[]int{1}, 1},
        {[]int{1, 2}, 2},
    }

    for _, test := range tests {
        list := createLinkedList(test.nums)
        result := FindMiddle(list)
        if result.Val != test.expected {
            t.Errorf("For list=%v, expected %d but got %d",
                test.nums, test.expected, result.Val)
        }
    }
}

func TestIsPalindrome(t *testing.T) {
    tests := []struct {
        nums     []int
        expected bool
    }{
        {[]int{1, 2, 2, 1}, true},
        {[]int{1, 2}, false},
        {[]int{1}, true},
        {[]int{1, 2, 3, 2, 1}, true},
    }

    for _, test := range tests {
        list := createLinkedList(test.nums)
        result := IsPalindrome(list)
        if result != test.expected {
            t.Errorf("For list=%v, expected %v but got %v",
                test.nums, test.expected, result)
        }
    }
}

func compareLinkedLists(l1, l2 *ListNode) bool {
    for l1 != nil && l2 != nil {
        if l1.Val != l2.Val {
            return false
        }
        l1 = l1.Next
        l2 = l2.Next
    }
    return l1 == nil && l2 == nil
}

func TestReorderList(t *testing.T) {
    tests := []struct {
        input    []int
        expected []int
    }{
        {[]int{1, 2, 3, 4}, []int{1, 4, 2, 3}},
        {[]int{1, 2, 3, 4, 5}, []int{1, 5, 2, 4, 3}},
        {[]int{1, 2}, []int{1, 2}},
        {[]int{1}, []int{1}},
    }

    for _, test := range tests {
        list := createLinkedList(test.input)
        expected := createLinkedList(test.expected)
        ReorderList(list)
        if !compareLinkedLists(list, expected) {
            t.Errorf("For input=%v, expected %v but got different result",
                test.input, test.expected)
        }
    }
}