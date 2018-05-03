package kdtree

import (
	"container/heap"

	"github.com/hongshibao/go-algo"
)

type Point interface {
	// Return the total number of dimensions
	Dim() int
	// Return the value X_{dim}, dim is started from 0
	GetValue(dim int) float64
	// Return the distance between two points
	Distance(p Point) float64
	// Return the distance between the point and the plane X_{dim}=val
	PlaneDistance(val float64, dim int) float64
}

type PointBase struct {
	Point
	Vec []float64
}

func (b PointBase) Dim() int {
	return len(b.Vec)
}

func (b PointBase) GetValue(dim int) float64 {
	return b.Vec[dim]
}

func NewPointBase(vals []float64) PointBase {
	ret := PointBase{}
	for _, val := range vals {
		ret.Vec = append(ret.Vec, val)
	}
	return ret
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

func (t *KDTree) KNN(target Point, k int) []Point {
	hp := &kNNHeapHelper{}
	t.search(t.root, hp, target, k)
	ret := make([]Point, 0, hp.Len())
	for hp.Len() > 0 {
		item := heap.Pop(hp).(*kNNHeapNode)
		ret = append(ret, item.point)
	}
	for i := len(ret)/2 - 1; i >= 0; i-- {
		opp := len(ret) - 1 - i
		ret[i], ret[opp] = ret[opp], ret[i]
	}
	return ret
}

func (t *KDTree) search(p *kdTreeNode,
	hp *kNNHeapHelper, target Point, k int) {
	stk := make([]*kdTreeNode, 0)
	for p != nil {
		stk = append(stk, p)
		if target.GetValue(p.axis) < p.splittingPoint.GetValue(p.axis) {
			p = p.leftChild
		} else {
			p = p.rightChild
		}
	}
	for i := len(stk) - 1; i >= 0; i-- {
		cur := stk[i]
		dist := target.Distance(cur.splittingPoint)
		if hp.Len() < k || (*hp)[0].distance >= dist {
			heap.Push(hp, &kNNHeapNode{
				point:    cur.splittingPoint,
				distance: dist,
			})
			if hp.Len() > k {
				heap.Pop(hp)
			}
		}
		if hp.Len() < k || target.PlaneDistance(
			cur.splittingPoint.GetValue(cur.axis), cur.axis) <=
			(*hp)[0].distance {
			if target.GetValue(cur.axis) < cur.splittingPoint.GetValue(cur.axis) {
				t.search(cur.rightChild, hp, target, k)
			} else {
				t.search(cur.leftChild, hp, target, k)
			}
		}
	}
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
	ret.leftChild = createKDTree(points[0:idx], depth+1)
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

type kNNHeapHelper []*kNNHeapNode

func (h kNNHeapHelper) Len() int {
	return len(h)
}

func (h kNNHeapHelper) Less(i, j int) bool {
	return h[i].distance > h[j].distance
}

func (h kNNHeapHelper) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *kNNHeapHelper) Push(x interface{}) {
	item := x.(*kNNHeapNode)
	*h = append(*h, item)
}

func (h *kNNHeapHelper) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}
