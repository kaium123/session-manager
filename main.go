package main

import (
	"fmt"
)

// Custom deque implementation
type Deque struct {
	data [][]int
}

func (d *Deque) PushBack(val []int) {
	d.data = append(d.data, val)
}

func (d *Deque) PushFront(val []int) {
	d.data = append([][]int{val}, d.data...)
}

func (d *Deque) PopFront() []int {
	if len(d.data) == 0 {
		return nil
	}
	val := d.data[0]
	d.data = d.data[1:]
	return val
}

func (d *Deque) Len() int {
	return len(d.data)
}

// Solve the problem using the custom deque
func numSubIslands(grid1 [][]int, grid2 [][]int) int {
	m, n := len(grid1), len(grid1[0])
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	// Helper function to check if a cell is valid
	isValid := func(x, y int) bool {
		return x >= 0 && x < m && y >= 0 && y < n
	}

	// BFS function using the custom deque
	bfs := func(x, y int) bool {
		queue := &Deque{}
		queue.PushBack([]int{x, y})
		isSubIsland := true

		for queue.Len() > 0 {
			cell := queue.PopFront()
			cx, cy := cell[0], cell[1]

			for _, dir := range directions {
				nx, ny := cx+dir[0], cy+dir[1]
				if isValid(nx, ny) && grid2[nx][ny] == 1 {
					if grid1[nx][ny] == 0 {
						isSubIsland = false
					}
					grid2[nx][ny] = 0 // Mark as visited
					queue.PushBack([]int{nx, ny})
				}
			}
		}
		return isSubIsland
	}

	// Iterate through grid2 and count sub-islands
	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid2[i][j] == 1 {
				grid2[i][j] = 0 // Mark as visited
				if grid1[i][j] == 0 || !bfs(i, j) {
					continue
				}
				count++
			}
		}
	}

	return count
}

func main() {
	grid1 := [][]int{
		{1, 1, 1, 0, 0},
		{1, 1, 0, 0, 0},
		{1, 1, 0, 1, 1},
		{0, 0, 0, 1, 1},
	}
	grid2 := [][]int{
		{1, 1, 1, 0, 0},
		{0, 1, 0, 0, 0},
		{1, 1, 0, 1, 1},
		{0, 0, 0, 1, 1},
	}

	fmt.Println("Number of Sub-Islands:", numSubIslands(grid1, grid2))
}
