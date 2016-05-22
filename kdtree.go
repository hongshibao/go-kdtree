package kdtree

import (
	"github.com/hongshibao/go-algo"
)

type Point interface {
	Dim() int
	GetValue(dim int) float64
	Distance(p Point) float64
}

type kdTreeNode struct {
	axis           int
	splittingPoint Point
	leftChild      *kdTreeNode
	rightChild     *kdTreeNode
}

type KDTree struct {
	root *kdTreeNode
	dim  int
}

func (t *KDTree) Dim() int {
	return t.dim
}

func (t *KDTree) KNN(k int) []Point {
	// TODO
	return nil
}

func NewKDTree(points []Point) *KDTree {
	if len(points) == 0 {
		return nil
	}
	ret := &KDTree{
		dim:  points[0].Dim(),
		root: createKDTree(points, 0),
	}
	return ret
}

func createKDTree(points []Point, depth int) *kdTreeNode {
	if len(points) == 0 {
		return nil
	}
	dim := points[0].Dim()
	ret := &kdTreeNode{
		axis: depth % dim,
	}
	if len(points) == 1 {
		ret.splittingPoint = points[0]
		return ret
	}
	idx := selectSplittingPoint(points, ret.axis)
	if idx == -1 {
		return nil
	}
	ret.splittingPoint = points[idx]
	ret.leftChild = createKDTree(points[0:idx-1], depth+1)
	ret.rightChild = createKDTree(points[idx+1:len(points)], depth+1)
	return ret
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

func selectSplittingPoint(points []Point, axis int) int {
	helper := &selectionHelper{
		axis:   axis,
		points: points,
	}
	mid := len(points)/2 + 1
	err := algo.QuickSelect(helper, mid)
	if err != nil {
		return -1
	}
	return mid - 1
}

type kNNHeapNode struct {
	point    Point
	distance float64
}
