package kdtree

type Point interface {
	Dim() int
	GetValue(dim int) float64
	Distance(p Point) float64
}

func NewKDTree(points []Point) *KDTree {
	return nil
}

type KDTreeNode struct {
	axis           int
	splittingPlane float64
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
