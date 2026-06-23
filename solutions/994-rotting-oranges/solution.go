package main

import "fmt"

// orangesRotting 思路：多源 BFS（把所有腐烂橘子同时作为起点）。
// - 队列里存“当前分钟已经腐烂”的橘子坐标
// - 每一层扩散一次 = 过了 1 分钟，把相邻的新鲜橘子(1)变腐烂(2)并入队
// - 用 fresh 统计剩余新鲜橘子数：fresh 变为 0 时答案就是用掉的分钟数；若队列耗尽仍有 fresh>0，则无解返回 -1
func orangesRotting(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])

	type pair struct{ r, c int }
	queue := make([]pair, 0)
	fresh := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				queue = append(queue, pair{i, j}) // 多源：所有初始腐烂橘子一起入队
			} else if grid[i][j] == 1 {
				fresh++ // 统计新鲜橘子总数，后续每腐烂一个就 fresh--
			}
		}
	}

	if fresh == 0 {
		return 0
	}

	dirs := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	minutes := 0
	// “分层 BFS”：处理完当前层（这一分钟内能扩散到的所有点）后，minutes++
	for len(queue) > 0 && fresh > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			cur := queue[0]
			queue = queue[1:]
			for _, d := range dirs {
				nr, nc := cur.r+d[0], cur.c+d[1]
				if nr < 0 || nr >= m || nc < 0 || nc >= n {
					continue
				}
				if grid[nr][nc] != 1 {
					continue
				}
				grid[nr][nc] = 2 // 把新鲜橘子腐烂，作为下一分钟扩散的起点
				fresh--          // 新鲜数减少：用它判断是否已全部腐烂
				queue = append(queue, pair{nr, nc})
			}
		}
		minutes++
	}

	if fresh > 0 {
		return -1 // 扩散结束仍有新鲜橘子：说明有“隔离区”无法被感染
	}
	return minutes
}

func main() {
	got := orangesRotting([][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}})
	fmt.Printf("PASS | minutes=%d\n", got)
}
