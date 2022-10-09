package bsp

import (
	"math/rand"

	"github.com/rhinoxi/mapgen/util"
)

type BspTree struct {
	left   int
	right  int
	bottom int
	top    int
	vsplit bool
	Left   *BspTree // left, bottom
	Right  *BspTree // right, top
}

func NewBspTree(left, right, bottom, top int) *BspTree {
	tree := &BspTree{
		left:   left,
		right:  right,
		bottom: bottom,
		top:    top,
	}

	if (right - left) < (top - bottom) {
		tree.vsplit = true
	}
	return tree
}

func (t *BspTree) Split() {
	if t.vsplit {
		if t.top-t.bottom < 5 {
			return
		}
		start := (t.top - t.bottom) / 3
		stop := (t.top - t.bottom) / 3 * 2
		pivot := util.RandInt(start, stop+1) + t.bottom

		t.Left = NewBspTree(t.left, t.right, t.bottom, pivot)
		t.Right = NewBspTree(t.left, t.right, pivot+1, t.top)
	} else {
		if t.right-t.left < 5 {
			return
		}
		start := (t.right - t.left) / 3
		stop := (t.right - t.left) / 3 * 2
		pivot := util.RandInt(start, stop+1) + t.left

		t.Left = NewBspTree(t.left, pivot, t.bottom, t.top)
		t.Right = NewBspTree(pivot, t.right, t.bottom, t.top)
	}
}

/*
depth: 1 -> no split
depth: 2 -> split once
*/
func gen(tree *BspTree, depth int) {
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
		bottom := util.RandInt(tree.bottom, tree.bottom+(tree.top-tree.bottom)/4+1)
		top := util.RandInt(tree.top-(tree.top-tree.bottom)/4, tree.top+1)

		tree.left = left
		tree.right = right
		tree.bottom = bottom
		tree.top = top
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
		xCenter := (tree.left + tree.right) / 2
		yTop := (tree.Right.bottom + tree.Right.top) / 2
		yBottom := (tree.Left.bottom + tree.Left.top) / 2
		for i := yBottom; i < yTop; i++ {
			m[i][xCenter] = true
		}
	} else {
		yCenter := (tree.bottom + tree.top) / 2
		xLeft := (tree.Left.left + tree.Left.right) / 2
		xRight := (tree.Right.left + tree.Right.right) / 2
		for i := xLeft; i < xRight; i++ {
			m[yCenter][i] = true
		}
	}

	dig(tree.Left, m)
	dig(tree.Right, m)
}

func fillMap(tree *BspTree, m [][]bool) {
	if tree.Left == nil {
		for j := tree.bottom; j <= tree.top; j++ {
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
