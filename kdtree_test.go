package kdtree

import (
	"math"
	"testing"
)

type EuclideanPoint struct {
	vec []float64
}

func (p *EuclideanPoint) Dim() int {
	return len(p.vec)
}

func (p *EuclideanPoint) GetValue(dim int) float64 {
	return p.vec[dim]
}

func (p *EuclideanPoint) Distance(ep Point) float64 {
	var ret float64
	for i := 0; i < len(p.vec); i++ {
		tmp := p.GetValue(i) - ep.GetValue(i)
		ret += tmp * tmp
	}
	return ret
}

func (p *EuclideanPoint) PlaneDistance(val float64, dim int) float64 {
	tmp := p.vec[dim] - val
	return math.Abs(tmp)
}

func NewEuclideanPoint(vals ...float64) *EuclideanPoint {
	ret := &EuclideanPoint{
		vec: []float64(vals),
	}
	return ret
}

func equal(p1 Point, p2 Point) bool {
	for i := 0; i < p1.Dim(); i++ {
		if p1.GetValue(i) != p2.GetValue(i) {
			return false
		}
	}
	return true
}

func checkKNNResult(t *testing.T, ans []Point, points ...Point) {
	if len(ans) != len(points) {
		t.Fatal("KNN result length error")
	}
	for i := 0; i < len(ans); i++ {
		if !equal(ans[i], points[i]) {
			t.Error("KNN results are wrong")
		}
	}
}

func TestKNN(t *testing.T) {
	// case 1
	points := make([]Point, 0)
	p1 := NewEuclideanPoint(0.0, 0.0, 0.0)
	p2 := NewEuclideanPoint(0.0, 0.0, 1.0)
	p3 := NewEuclideanPoint(0.0, 1.0, 0.0)
	p4 := NewEuclideanPoint(1.0, 0.0, 0.0)
	points = append(points, p1)
	points = append(points, p2)
	points = append(points, p3)
	points = append(points, p4)
	tree := NewKDTree(points)
	ans := tree.KNN(NewEuclideanPoint(0.0, 0.0, 0.1), 2)
	checkKNNResult(t, ans, p1, p2)
}
