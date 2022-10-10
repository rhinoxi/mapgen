package cellularautomata

import (
	"math/rand"
)

type CaMap struct {
	tiles  [][]bool // true: floor, false: wall
	width  int
	height int
}

func NewCaMap(width, height int) *CaMap {
	m := make([][]bool, height)
	for i := 0; i < height; i++ {
		m[i] = make([]bool, width)
	}
	return &CaMap{
		tiles:  m,
		width:  width,
		height: height,
	}
}

func (m *CaMap) initNoise(density int) {
	for _, row := range m.tiles {
		for i := range row {
			if rand.Intn(100)+1 <= density {
				row[i] = true
			}
		}
	}
}

func (m *CaMap) getAroundWallCount(i, j int) int {
	count := 0
	for ii := i - 1; ii <= i+1; ii++ {
		for jj := j - 1; jj <= j+1; jj++ {
			if ii == i && jj == j {
				continue
			}
			if ii < 0 || ii >= m.width || jj < 0 || jj >= m.height {
				count++
				continue
			}
			if !m.tiles[jj][ii] {
				count++
				continue
			}
		}
	}
	return count
}

func (m *CaMap) generate(iters int) {
	if iters == 0 {
		return
	}

	tiles := make([][]bool, m.height)
	for i := 0; i < m.height; i++ {
		tiles[i] = make([]bool, m.width)
	}

	for j, row := range m.tiles {
		for i := range row {
			count := m.getAroundWallCount(i, j)
			if (row[i] && count >= 5) || (!row[i] && count >= 4) {
				tiles[j][i] = false
			} else {
				tiles[j][i] = true
			}
		}
	}
	m.tiles = tiles
	m.generate(iters - 1)
}

func Gen(width, height, noiseDensity, iters int, seed int64) [][]bool {
	rand.Seed(seed)

	m := NewCaMap(width, height)

	m.initNoise(noiseDensity)

	m.generate(iters)

	return m.tiles
}
