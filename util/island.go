package util

type location struct {
	x int
	y int
}

type Island struct {
	locations []location
}

func (l *Island) Add(x, y int) {
	l.locations = append(l.locations, location{x: x, y: y})
}

func (l *Island) Area() int {
	return len(l.locations)
}

func RemoveSmallIslands(m [][]bool, islands []*Island, minArea int) []*Island {
	if minArea <= 1 {
		return islands
	}

	var largeIslands []*Island
	for _, island := range islands {
		if island.Area() < minArea {
			for _, loc := range island.locations {
				m[loc.y][loc.x] = false
			}
		} else {
			largeIslands = append(largeIslands, island)
		}
	}
	return largeIslands
}

func DetectIslands(m [][]bool) []*Island {
	var visited [][]bool

	mapHeight := len(m)
	mapWidth := len(m[0])

	var tobeVisited []location

	for _, row := range m {
		visited = append(visited, make([]bool, len(row)))
	}

	var islands []*Island
	for j, row := range m {
		for i, isFloor := range row {
			if visited[j][i] {
				continue
			}

			visited[j][i] = true

			if isFloor {
				neighbors := findNeighbors(i, j, mapWidth-1, mapHeight-1)
				for _, loc := range neighbors {
					if !visited[loc.y][loc.x] {
						tobeVisited = append(tobeVisited, loc)
					}
				}

				island := &Island{
					locations: []location{
						{
							x: i,
							y: j,
						},
					},
				}
				walkIsland(m, visited, tobeVisited, island)
				islands = append(islands, island)
			}
		}
	}

	return islands
}

func walkIsland(m [][]bool, visited [][]bool, tobeVisited []location, island *Island) {
	if len(tobeVisited) == 0 {
		return
	}
	mapHeight := len(m)
	mapWidth := len(m[0])

	var tobeVisitedLater []location

	for _, loc := range tobeVisited {
		if visited[loc.y][loc.x] {
			continue
		}
		visited[loc.y][loc.x] = true
		if m[loc.y][loc.x] {
			island.Add(loc.x, loc.y)
			neighbors := findNeighbors(loc.x, loc.y, mapWidth-1, mapHeight-1)
			for _, loc := range neighbors {
				if !visited[loc.y][loc.x] {
					tobeVisitedLater = append(tobeVisitedLater, loc)
				}
			}
		}
	}

	walkIsland(m, visited, tobeVisitedLater, island)
	return
}

func findNeighbors(x, y, maxX, maxY int) []location {
	var neighbors []location
	if x > 0 {
		neighbors = append(neighbors, location{x: x - 1, y: y})
	}
	if x < maxX {
		neighbors = append(neighbors, location{x: x + 1, y: y})
	}

	if y > 0 {
		neighbors = append(neighbors, location{x: x, y: y - 1})
	}
	if y < maxY {
		neighbors = append(neighbors, location{x: x, y: y + 1})
	}

	return neighbors
}
