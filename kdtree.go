package kdtree

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

func newKDTree(points []Point, depth int) *KDTreeNode {
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

func selectSplittingPoint(points []Point, depth int) Point {
	// TODO
	return nil
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
