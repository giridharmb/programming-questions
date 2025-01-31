/*

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

*/

// File: matrix/matrix.go

package matrix

// Direction vectors for 4-directional movement
var directions = [][]int{
    {-1, 0},  // up
    {1, 0},   // down
    {0, -1},  // left
    {0, 1},   // right
}

// NumIslands counts number of islands in binary matrix
// Time Complexity: O(m*n)
// Space Complexity: O(m*n)
func NumIslands(grid [][]byte) int {
    if len(grid) == 0 || len(grid[0]) == 0 {
        return 0
    }
    
    rows, cols := len(grid), len(grid[0])
    count := 0
    
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            if grid[i][j] == '1' {
                count++
                dfs(grid, i, j)
            }
        }
    }
    
    return count
}

// dfs helper function to mark connected land cells
func dfs(grid [][]byte, row, col int) {
    rows, cols := len(grid), len(grid[0])
    
    // Base cases
    if row < 0 || row >= rows || col < 0 || col >= cols || grid[row][col] != '1' {
        return
    }
    
    // Mark as visited
    grid[row][col] = '0'
    
    // Visit all adjacent cells
    for _, dir := range directions {
        dfs(grid, row+dir[0], col+dir[1])
    }
}

// FloodFill performs flood fill on image
// Time Complexity: O(m*n)
// Space Complexity: O(m*n)
func FloodFill(image [][]int, sr int, sc int, newColor int) [][]int {
    if len(image) == 0 || len(image[0]) == 0 {
        return image
    }
    
    oldColor := image[sr][sc]
    if oldColor != newColor {
        floodFillDFS(image, sr, sc, oldColor, newColor)
    }
    return image
}

func floodFillDFS(image [][]int, row, col, oldColor, newColor int) {
    rows, cols := len(image), len(image[0])
    
    // Base cases
    if row < 0 || row >= rows || col < 0 || col >= cols || image[row][col] != oldColor {
        return
    }
    
    // Color current cell
    image[row][col] = newColor
    
    // Visit adjacent cells
    for _, dir := range directions {
        floodFillDFS(image, row+dir[0], col+dir[1], oldColor, newColor)
    }
}

// MaxAreaIsland finds the maximum area of an island
// Time Complexity: O(m*n)
// Space Complexity: O(m*n)
func MaxAreaIsland(grid [][]int) int {
    if len(grid) == 0 || len(grid[0]) == 0 {
        return 0
    }
    
    maxArea := 0
    rows, cols := len(grid), len(grid[0])
    
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            if grid[i][j] == 1 {
                maxArea = max(maxArea, getIslandArea(grid, i, j))
            }
        }
    }
    
    return maxArea
}

func getIslandArea(grid [][]int, row, col int) int {
    rows, cols := len(grid), len(grid[0])
    
    if row < 0 || row >= rows || col < 0 || col >= cols || grid[row][col] != 1 {
        return 0
    }
    
    // Mark as visited
    grid[row][col] = 0
    area := 1
    
    // Visit adjacent cells
    for _, dir := range directions {
        area += getIslandArea(grid, row+dir[0], col+dir[1])
    }
    
    return area
}

// PacificAtlantic finds cells that can flow to both oceans
// Time Complexity: O(m*n)
// Space Complexity: O(m*n)
func PacificAtlantic(heights [][]int) [][]int {
    if len(heights) == 0 || len(heights[0]) == 0 {
        return [][]int{}
    }
    
    rows, cols := len(heights), len(heights[0])
    pacific := make([][]bool, rows)
    atlantic := make([][]bool, rows)
    
    for i := range pacific {
        pacific[i] = make([]bool, cols)
        atlantic[i] = make([]bool, cols)
    }
    
    // DFS from Pacific edges
    for i := 0; i < rows; i++ {
        oceanDFS(heights, i, 0, pacific, math.MinInt32)
    }
    for j := 0; j < cols; j++ {
        oceanDFS(heights, 0, j, pacific, math.MinInt32)
    }
    
    // DFS from Atlantic edges
    for i := 0; i < rows; i++ {
        oceanDFS(heights, i, cols-1, atlantic, math.MinInt32)
    }
    for j := 0; j < cols; j++ {
        oceanDFS(heights, rows-1, j, atlantic, math.MinInt32)
    }
    
    // Find cells that can reach both oceans
    result := make([][]int, 0)
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            if pacific[i][j] && atlantic[i][j] {
                result = append(result, []int{i, j})
            }
        }
    }
    
    return result
}

func oceanDFS(heights [][]int, row, col int, visited [][]bool, prevHeight int) {
    rows, cols := len(heights), len(heights[0])
    
    if row < 0 || row >= rows || col < 0 || col >= cols || 
       visited[row][col] || heights[row][col] < prevHeight {
        return
    }
    
    visited[row][col] = true
    
    for _, dir := range directions {
        oceanDFS(heights, row+dir[0], col+dir[1], visited, heights[row][col])
    }
}

// RottenOranges finds minimum time until all oranges rot
// Time Complexity: O(m*n)
// Space Complexity: O(m*n)
func RottenOranges(grid [][]int) int {
    if len(grid) == 0 || len(grid[0]) == 0 {
        return -1
    }
    
    rows, cols := len(grid), len(grid[0])
    queue := make([][2]int, 0)
    freshCount := 0
    
    // Find all rotten oranges and count fresh ones
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            if grid[i][j] == 2 {
                queue = append(queue, [2]int{i, j})
            } else if grid[i][j] == 1 {
                freshCount++
            }
        }
    }
    
    if freshCount == 0 {
        return 0
    }
    
    minutes := 0
    
    // Process rotten oranges level by level
    for len(queue) > 0 {
        size := len(queue)
        for i := 0; i < size; i++ {
            pos := queue[0]
            queue = queue[1:]
            
            for _, dir := range directions {
                newRow, newCol := pos[0]+dir[0], pos[1]+dir[1]
                
                if newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols &&
                   grid[newRow][newCol] == 1 {
                    grid[newRow][newCol] = 2
                    queue = append(queue, [2]int{newRow, newCol})
                    freshCount--
                }
            }
        }
        minutes++
    }
    
    if freshCount > 0 {
        return -1
    }
    return minutes - 1
}

// Helper function
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

// File: matrix/matrix_test.go

package matrix

import (
    "reflect"
    "testing"
)

func TestNumIslands(t *testing.T) {
    tests := []struct {
        name     string
        grid     [][]byte
        expected int
    }{
        {
            "Multiple islands",
            [][]byte{
                {'1', '1', '0', '0', '0'},
                {'1', '1', '0', '0', '0'},
                {'0', '0', '1', '0', '0'},
                {'0', '0', '0', '1', '1'},
            },
            3,
        },
        {
            "Single island",
            [][]byte{
                {'1', '1', '1'},
                {'1', '1', '1'},
                {'1', '1', '1'},
            },
            1,
        },
        {
            "No islands",
            [][]byte{
                {'0', '0', '0'},
                {'0', '0', '0'},
            },
            0,
        },
    }
    
    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            result := NumIslands(test.grid)
            if result != test.expected {
                t.Errorf("Expected %d but got %d", test.expected, result)
            }
        })
    }
}

func TestFloodFill(t *testing.T) {
    tests := []struct {
        name      string
        image     [][]int
        sr, sc    int
        newColor  int
        expected  [][]int
    }{
        {
            "Simple fill",
            [][]int{
                {1, 1, 1},
                {1, 1, 0},
                {1, 0, 1},
            },
            1, 1, 2,
            [][]int{
                {2, 2, 2},
                {2, 2, 0},
                {2, 0, 1},
            },
        },
    }
    
    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            result := FloodFill(test.image, test.sr, test.sc, test.newColor)
            if !reflect.DeepEqual(result, test.expected) {
                t.Errorf("Expected %v but got %v", test.expected, result)
            }
        })
    }
}

func TestMaxAreaIsland(t *testing.T) {
    tests := []struct {
        name     string
        grid     [][]int
        expected int
    }{
        {
            "Multiple islands",
            [][]int{
                {0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
                {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
                {0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
                {0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
                {0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
                {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
                {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
            },
            6,
        },
    }
    
    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            result := MaxAreaIsland(test.grid)
            if result != test.expected {
                t.Errorf("Expected %d but got %d", test.expected, result)
            }
        })
    }
}

func TestRottenOranges(t *testing.T) {
    tests := []struct {
        name     string
        grid     [][]int
        expected int
    }{
        {
            "All oranges rot",
            [][]int{
                {2, 1, 1},
                {1, 1, 0},
                {0, 1, 1},
            },
            4,
        },
        {
            "Cannot rot all",
            [][]int{
                {2, 1, 1},
                {0, 1, 1},
                {1, 0, 1},
            },
            -1,
        },
    }
    
    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            result := RottenOranges(test.grid)
            if result != test.expected {
                t.Errorf("Expected %d but got %d", test.expected, result)
            }
        })
    }
}