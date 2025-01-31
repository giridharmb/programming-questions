/*

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

*/

// File: topk/topk.go

package topk

import (
    "container/heap"
)

// MinHeap implementation for integers
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

// MaxHeap implementation using MinHeap
type MaxHeap struct{ MinHeap }
func (h MaxHeap) Less(i, j int) bool { return h.MinHeap[i] > h.MinHeap[j] }

// Pair structure for frequency counting
type Pair struct {
    Num   int
    Count int
}

// PairHeap for frequency-based problems
type PairHeap []Pair
func (h PairHeap) Len() int           { return len(h) }
func (h PairHeap) Less(i, j int) bool { return h[i].Count > h[j].Count }
func (h PairHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *PairHeap) Push(x interface{}) { *h = append(*h, x.(Pair)) }
func (h *PairHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

// FindKthLargest finds the kth largest element in an array
// Time Complexity: O(n log k)
// Space Complexity: O(k)
func FindKthLargest(nums []int, k int) int {
    h := &MinHeap{}
    heap.Init(h)
    
    for _, num := range nums {
        heap.Push(h, num)
        if h.Len() > k {
            heap.Pop(h)
        }
    }
    
    return (*h)[0]
}

// TopKFrequent finds the k most frequent elements
// Time Complexity: O(n log k)
// Space Complexity: O(n)
func TopKFrequent(nums []int, k int) []int {
    // Count frequencies
    freqMap := make(map[int]int)
    for _, num := range nums {
        freqMap[num]++
    }
    
    // Use heap to find top k elements
    h := &PairHeap{}
    heap.Init(h)
    
    for num, count := range freqMap {
        heap.Push(h, Pair{num, count})
    }
    
    // Extract top k elements
    result := make([]int, k)
    for i := 0; i < k; i++ {
        result[i] = heap.Pop(h).(Pair).Num
    }
    
    return result
}

// KClosestPoints finds k points closest to origin
// Time Complexity: O(n log k)
// Space Complexity: O(k)
type Point struct {
    x, y     int
    distance int
}

type PointHeap []Point
func (h PointHeap) Len() int           { return len(h) }
func (h PointHeap) Less(i, j int) bool { return h[i].distance > h[j].distance }
func (h PointHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *PointHeap) Push(x interface{}) { *h = append(*h, x.(Point)) }
func (h *PointHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

func KClosestPoints(points [][]int, k int) [][]int {
    h := &PointHeap{}
    heap.Init(h)
    
    for _, point := range points {
        distance := point[0]*point[0] + point[1]*point[1]
        p := Point{point[0], point[1], distance}
        heap.Push(h, p)
        if h.Len() > k {
            heap.Pop(h)
        }
    }
    
    result := make([][]int, k)
    for i := k - 1; i >= 0; i-- {
        p := heap.Pop(h).(Point)
        result[i] = []int{p.x, p.y}
    }
    
    return result
}

// KthSmallestInMatrix finds kth smallest element in sorted matrix
// Time Complexity: O(k log n)
// Space Complexity: O(n)
type MatrixCell struct {
    val, row, col int
}

type MatrixHeap []MatrixCell
func (h MatrixHeap) Len() int           { return len(h) }
func (h MatrixHeap) Less(i, j int) bool { return h[i].val < h[j].val }
func (h MatrixHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MatrixHeap) Push(x interface{}) { *h = append(*h, x.(MatrixCell)) }
func (h *MatrixHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

func KthSmallestInMatrix(matrix [][]int, k int) int {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return 0
    }
    
    n := len(matrix)
    h := &MatrixHeap{}
    heap.Init(h)
    
    // Add first element from each row
    for i := 0; i < n; i++ {
        heap.Push(h, MatrixCell{matrix[i][0], i, 0})
    }
    
    // Process k elements
    for i := 0; i < k-1; i++ {
        cell := heap.Pop(h).(MatrixCell)
        if cell.col < n-1 {
            heap.Push(h, MatrixCell{
                matrix[cell.row][cell.col+1],
                cell.row,
                cell.col + 1,
            })
        }
    }
    
    return heap.Pop(h).(MatrixCell).val
}

// File: topk/topk_test.go

package topk

import (
    "reflect"
    "testing"
)

func TestFindKthLargest(t *testing.T) {
    tests := []struct {
        nums     []int
        k        int
        expected int
    }{
        {[]int{3, 2, 1, 5, 6, 4}, 2, 5},
        {[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4},
        {[]int{1}, 1, 1},
    }
    
    for _, test := range tests {
        result := FindKthLargest(test.nums, test.k)
        if result != test.expected {
            t.Errorf("For nums=%v, k=%d, expected %d but got %d",
                test.nums, test.k, test.expected, result)
        }
    }
}

func TestTopKFrequent(t *testing.T) {
    tests := []struct {
        nums     []int
        k        int
        expected []int
    }{
        {[]int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
        {[]int{1}, 1, []int{1}},
        {[]int{1, 2}, 2, []int{1, 2}},
    }
    
    for _, test := range tests {
        result := TopKFrequent(test.nums, test.k)
        if !reflect.DeepEqual(result, test.expected) {
            t.Errorf("For nums=%v, k=%d, expected %v but got %v",
                test.nums, test.k, test.expected, result)
        }
    }
}

func TestKClosestPoints(t *testing.T) {
    tests := []struct {
        points   [][]int
        k        int
        expected [][]int
    }{
        {
            [][]int{{1, 3}, {-2, 2}, {5, -1}},
            2,
            [][]int{{-2, 2}, {1, 3}},
        },
        {
            [][]int{{3, 3}, {5, -1}, {-2, 4}},
            2,
            [][]int{{3, 3}, {-2, 4}},
        },
    }
    
    for _, test := range tests {
        result := KClosestPoints(test.points, test.k)
        if !reflect.DeepEqual(result, test.expected) {
            t.Errorf("For points=%v, k=%d, expected %v but got %v",
                test.points, test.k, test.expected, result)
        }
    }
}

func TestKthSmallestInMatrix(t *testing.T) {
    tests := []struct {
        matrix   [][]int
        k        int
        expected int
    }{
        {
            [][]int{
                {1, 5, 9},
                {10, 11, 13},
                {12, 13, 15},
            },
            8,
            13,
        },
        {
            [][]int{
                {1, 2},
                {3, 4},
            },
            2,
            2,
        },
    }
    
    for _, test := range tests {
        result := KthSmallestInMatrix(test.matrix, test.k)
        if result != test.expected {
            t.Errorf("For matrix=%v, k=%d, expected %d but got %d",
                test.matrix, test.k, test.expected, result)
        }
    }
}