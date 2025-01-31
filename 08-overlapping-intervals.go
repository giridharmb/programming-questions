/*

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

*/

// File: intervals/intervals.go

package intervals

import (
    "sort"
)

// Interval represents a time interval with start and end times
type Interval struct {
    Start int
    End   int
}

// MergeIntervals merges all overlapping intervals
// Time Complexity: O(n log n)
// Space Complexity: O(n)
func MergeIntervals(intervals [][]int) [][]int {
    if len(intervals) <= 1 {
        return intervals
    }
    
    // Sort intervals based on start time
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    
    merged := make([][]int, 0)
    current := intervals[0]
    
    for i := 1; i < len(intervals); i++ {
        if current[1] >= intervals[i][0] {
            // Overlapping intervals, update end time
            if intervals[i][1] > current[1] {
                current[1] = intervals[i][1]
            }
        } else {
            // Non-overlapping interval, add current to result
            merged = append(merged, current)
            current = intervals[i]
        }
    }
    
    // Add the last interval
    merged = append(merged, current)
    return merged
}

// InsertInterval inserts a new interval and merges if necessary
// Time Complexity: O(n)
// Space Complexity: O(n)
func InsertInterval(intervals [][]int, newInterval []int) [][]int {
    result := make([][]int, 0)
    i := 0
    n := len(intervals)
    
    // Add all intervals that come before newInterval
    for i < n && intervals[i][1] < newInterval[0] {
        result = append(result, intervals[i])
        i++
    }
    
    // Merge overlapping intervals
    for i < n && intervals[i][0] <= newInterval[1] {
        newInterval[0] = min(newInterval[0], intervals[i][0])
        newInterval[1] = max(newInterval[1], intervals[i][1])
        i++
    }
    result = append(result, newInterval)
    
    // Add remaining intervals
    for i < n {
        result = append(result, intervals[i])
        i++
    }
    
    return result
}

// IntervalIntersection finds the intersection of two interval lists
// Time Complexity: O(n + m)
// Space Complexity: O(n + m)
func IntervalIntersection(firstList [][]int, secondList [][]int) [][]int {
    result := make([][]int, 0)
    i, j := 0, 0
    
    for i < len(firstList) && j < len(secondList) {
        // Find overlapping interval
        start := max(firstList[i][0], secondList[j][0])
        end := min(firstList[i][1], secondList[j][1])
        
        if start <= end {
            result = append(result, []int{start, end})
        }
        
        // Move pointer of interval that ends earlier
        if firstList[i][1] < secondList[j][1] {
            i++
        } else {
            j++
        }
    }
    
    return result
}

// CanAttendMeetings checks if a person can attend all meetings
// Time Complexity: O(n log n)
// Space Complexity: O(1)
func CanAttendMeetings(intervals [][]int) bool {
    if len(intervals) <= 1 {
        return true
    }
    
    // Sort by start time
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    
    // Check for overlap
    for i := 1; i < len(intervals); i++ {
        if intervals[i][0] < intervals[i-1][1] {
            return false
        }
    }
    
    return true
}

// MinMeetingRooms finds minimum number of meeting rooms required
// Time Complexity: O(n log n)
// Space Complexity: O(n)
func MinMeetingRooms(intervals [][]int) int {
    if len(intervals) == 0 {
        return 0
    }
    
    // Separate start and end times
    starts := make([]int, len(intervals))
    ends := make([]int, len(intervals))
    
    for i, interval := range intervals {
        starts[i] = interval[0]
        ends[i] = interval[1]
    }
    
    // Sort start and end times
    sort.Ints(starts)
    sort.Ints(ends)
    
    rooms := 0
    maxRooms := 0
    startPtr := 0
    endPtr := 0
    
    // Process events in chronological order
    for startPtr < len(intervals) {
        if starts[startPtr] < ends[endPtr] {
            rooms++
            startPtr++
        } else {
            rooms--
            endPtr++
        }
        maxRooms = max(maxRooms, rooms)
    }
    
    return maxRooms
}

// Helper functions
func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

// File: intervals/intervals_test.go

package intervals

import (
    "reflect"
    "testing"
)

func TestMergeIntervals(t *testing.T) {
    tests := []struct {
        intervals [][]int
        expected [][]int
    }{
        {
            [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
            [][]int{{1, 6}, {8, 10}, {15, 18}},
        },
        {
            [][]int{{1, 4}, {4, 5}},
            [][]int{{1, 5}},
        },
        {
            [][]int{{1, 4}},
            [][]int{{1, 4}},
        },
    }
    
    for _, test := range tests {
        result := MergeIntervals(test.intervals)
        if !reflect.DeepEqual(result, test.expected) {
            t.Errorf("For intervals=%v, expected %v but got %v",
                test.intervals, test.expected, result)
        }
    }
}

func TestInsertInterval(t *testing.T) {
    tests := []struct {
        intervals    [][]int
        newInterval []int
        expected    [][]int
    }{
        {
            [][]int{{1, 3}, {6, 9}},
            []int{2, 5},
            [][]int{{1, 5}, {6, 9}},
        },
        {
            [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
            []int{4, 8},
            [][]int{{1, 2}, {3, 10}, {12, 16}},
        },
    }
    
    for _, test := range tests {
        result := InsertInterval(test.intervals, test.newInterval)
        if !reflect.DeepEqual(result, test.expected) {
            t.Errorf("For intervals=%v, newInterval=%v, expected %v but got %v",
                test.intervals, test.newInterval, test.expected, result)
        }
    }
}

func TestIntervalIntersection(t *testing.T) {
    tests := []struct {
        firstList  [][]int
        secondList [][]int
        expected   [][]int
    }{
        {
            [][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}},
            [][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}},
            [][]int{{1, 2}, {5, 5}, {8, 10}, {15, 23}, {24, 24}, {25, 25}},
        },
        {
            [][]int{{1, 3}, {5, 9}},
            [][]int{},
            [][]int{},
        },
    }
    
    for _, test := range tests {
        result := IntervalIntersection(test.firstList, test.secondList)
        if !reflect.DeepEqual(result, test.expected) {
            t.Errorf("For firstList=%v, secondList=%v, expected %v but got %v",
                test.firstList, test.secondList, test.expected, result)
        }
    }
}

func TestCanAttendMeetings(t *testing.T) {
    tests := []struct {
        intervals [][]int
        expected bool
    }{
        {[][]int{{0, 30}, {5, 10}, {15, 20}}, false},
        {[][]int{{7, 10}, {2, 4}}, true},
        {[][]int{{1, 2}, {2, 3}, {3, 4}}, true},
    }
    
    for _, test := range tests {
        result := CanAttendMeetings(test.intervals)
        if result != test.expected {
            t.Errorf("For intervals=%v, expected %v but got %v",
                test.intervals, test.expected, result)
        }
    }
}

func TestMinMeetingRooms(t *testing.T) {
    tests := []struct {
        intervals [][]int
        expected int
    }{
        {[][]int{{0, 30}, {5, 10}, {15, 20}}, 2},
        {[][]int{{7, 10}, {2, 4}}, 1},
        {[][]int{{1, 5}, {8, 9}, {8, 9}}, 2},
    }
    
    for _, test := range tests {
        result := MinMeetingRooms(test.intervals)
        if result != test.expected {
            t.Errorf("For intervals=%v, expected %d but got %d",
                test.intervals, test.expected, result)
        }
    }
}