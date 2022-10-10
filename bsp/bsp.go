package bsp

import (
	"math/rand"

	"github.com/rhinoxi/mapgen/util"
)

/*
coordinate:
(0,0) - (0,1)
(1,0) - (1,1)
*/
type BspTree struct {
	left   int
	right  int
	top    int
	bottom int
	vsplit bool
	Left   *BspTree // left, top
	Right  *BspTree // right, bottom
}

func NewBspTree(left, right, top, bottom int) *BspTree {
	tree := &BspTree{
		left:   left,
		right:  right,
		top:    top,
		bottom: bottom,
	}

	if (right - left) < (bottom - top) {
		tree.vsplit = true
	}
	return tree
}

func (t *BspTree) RowCenter() int {
	return (t.top + t.bottom) / 2
}

func (t *BspTree) ColumnCenter() int {
	return (t.left + t.right) / 2
}

func (t *BspTree) RowCount() int {
	return t.bottom - t.top + 1
}

func (t *BspTree) ColumnCount() int {
	return t.right - t.left + 1
}

func (t *BspTree) Split() {
	if t.vsplit {
		if t.RowCount() < 8 {
			return
		}
		start := (t.RowCount() - 1) / 3
		stop := (t.RowCount() - 1) / 3 * 2
		pivot := util.RandInt(start, stop+1) + t.top

		t.Left = NewBspTree(t.left, t.right, t.top, pivot)
		t.Right = NewBspTree(t.left, t.right, pivot+1, t.bottom)
	} else {
		if t.ColumnCount() < 8 {
			return
		}
		start := (t.ColumnCount() - 1) / 3
		stop := (t.ColumnCount() - 1) / 3 * 2
		pivot := util.RandInt(start, stop+1) + t.left

		t.Left = NewBspTree(t.left, pivot, t.top, t.bottom)
		t.Right = NewBspTree(pivot, t.right, t.top, t.bottom)
	}
}

/*
depth: 1 -> no split
depth: 2 -> split once
*/
func gen(tree *BspTree, depth int) {
	if tree == nil {
		return
	}
	depth--
	if depth == 0 {
		return
	}
	tree.Split()
	gen(tree.Left, depth)
	gen(tree.Right, depth)
}

func shrinkLeaf(tree *BspTree) {
	if tree.Left == nil {
		left := util.RandInt(tree.left, tree.left+(tree.right-tree.left)/4+1)
		right := util.RandInt(tree.right-(tree.right-tree.left)/4, tree.right+1)
		top := util.RandInt(tree.top, tree.top+(tree.bottom-tree.top)/4+1)
		bottom := util.RandInt(tree.bottom-(tree.bottom-tree.top)/4, tree.bottom+1)

		tree.left = left
		tree.right = right
		tree.top = top
		tree.bottom = bottom
		return
	}
	shrinkLeaf(tree.Left)
	shrinkLeaf(tree.Right)
}

func dig(tree *BspTree, m [][]bool) {
	if tree.Left == nil {
		return
	}
	if tree.vsplit {
		yTop := tree.Left.RowCenter()
		yBottom := tree.Right.RowCenter()
		for i := yTop; i < yBottom; i++ {
			m[i][tree.ColumnCenter()] = true
		}
	} else {
		xLeft := tree.Left.ColumnCenter()
		xRight := tree.Right.ColumnCenter()
		for i := xLeft; i < xRight; i++ {
			m[tree.RowCenter()][i] = true
		}
	}

	dig(tree.Left, m)
	dig(tree.Right, m)
}

func fillMap(tree *BspTree, m [][]bool) {
	if tree.Left == nil {
		for j := tree.top; j <= tree.bottom; j++ {
			for i := tree.left; i <= tree.right; i++ {
				m[j][i] = true
			}
		}
		return
	}
	fillMap(tree.Left, m)
	fillMap(tree.Right, m)
}

func Gen(width, height, depth int, seed int64) [][]bool {
	rand.Seed(seed)
	m := make([][]bool, height)
	for i := 0; i < height; i++ {
		m[i] = make([]bool, width)
	}
	tree := NewBspTree(0, width-1, 0, height-1)
	gen(tree, depth)

	shrinkLeaf(tree)

	fillMap(tree, m)
	dig(tree, m)
	return m
}
