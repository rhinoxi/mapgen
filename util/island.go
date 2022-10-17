package util

func CountIsland(m [][]bool) []int {
	var islands []int
	var visited [][]bool

	mapHeight := len(m)
	mapWidth := len(m[0])

	var tobeVisited [][2]int // [(x1, y1), (x2, y2), ...]

	for _, row := range m {
		visited = append(visited, make([]bool, len(row)))
	}

	islandCount := 0
	for j, row := range m {
		for i, isFloor := range row {
			if visited[j][i] {
				continue
			}

			visited[j][i] = true

			if isFloor {
				islandCount++
				islands = append(islands, 1)
				neighbors := findNeighbors(i, j, mapWidth-1, mapHeight-1)
				for _, loc := range neighbors {
					if !visited[loc[1]][loc[0]] {
						tobeVisited = append(tobeVisited, loc)
					}
				}

				islands[islandCount-1] = walkIsland(m, visited, tobeVisited, 1)
			}
		}
	}

	return islands
}

func walkIsland(m [][]bool, visited [][]bool, tobeVisited [][2]int, area int) int {
	if len(tobeVisited) == 0 {
		return area
	}
	mapHeight := len(m)
	mapWidth := len(m[0])

	var tobeVisitedLater [][2]int

	for _, loc := range tobeVisited {
		x := loc[0]
		y := loc[1]
		if visited[y][x] {
			continue
		}
		visited[y][x] = true
		if m[y][x] {
			area++
			neighbors := findNeighbors(x, y, mapWidth-1, mapHeight-1)
			for _, loc := range neighbors {
				if !visited[loc[1]][loc[0]] {
					tobeVisitedLater = append(tobeVisitedLater, loc)
				}
			}
		}
	}

	return walkIsland(m, visited, tobeVisitedLater, area)
}

func findNeighbors(x, y, maxX, maxY int) [][2]int {
	var neighbors [][2]int
	if x > 0 {
		neighbors = append(neighbors, [2]int{x - 1, y})
	}
	if x < maxX {
		neighbors = append(neighbors, [2]int{x + 1, y})
	}

	if y > 0 {
		neighbors = append(neighbors, [2]int{x, y - 1})
	}
	if y < maxY {
		neighbors = append(neighbors, [2]int{x, y + 1})
	}

	return neighbors
}
