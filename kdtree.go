package kdtree

import (
	"github.com/hongshibao/go-algo"
)

type Point interface {
	Dim() int
	GetValue(dim int) float64
	Distance(p Point) float64
}

func NewKDTree(points []Point) *KDTree {
	if len(points) == 0 {
		return nil
	}
	ret := &KDTree{
		dim: points[0].Dim(),
	}
	// TODO
	return ret
}

func createKDTree(points []Point, depth int) *KDTreeNode {
	if len(points) == 0 {
		return nil
	}
	dim := points[0].Dim()
	if len(points) == 1 {
		return &KDTreeNode{
			axis:           depth % dim,
			splittingPoint: points[0],
			leftChild:      nil,
			rightChild:     nil,
		}
	}
	// TODO
	return nil
}

type selectionHelper struct {
	axis   int
	points []Point
}

func (h *selectionHelper) Len() int {
	return len(h.points)
}

func (h *selectionHelper) Less(i, j int) bool {
	return h.points[i].GetValue(h.axis) < h.points[j].GetValue(h.axis)
}

func (h *selectionHelper) Swap(i, j int) {
	h.points[i], h.points[j] = h.points[j], h.points[i]
}

func selectSplittingPoint(points []Point, axis int) Point {
	helper := &selectionHelper{
		axis:   axis,
		points: points,
	}
	mid := len(points)/2 + 1
	err := algo.QuickSelect(helper, mid)
	if err != nil {
		return nil
	}
	return points[mid-1]
}

type KDTreeNode struct {
	axis           int
	splittingPoint Point
	leftChild      *KDTreeNode
	rightChild     *KDTreeNode
}

type KDTree struct {
	root *KDTreeNode
	dim  int
}

func (t *KDTree) Dim() int {
	return t.dim
}
