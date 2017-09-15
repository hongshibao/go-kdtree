package kdtree

import (
	"testing"
)

type EuclideanPoint struct {
	PointBase
}

func (p *EuclideanPoint) Distance(other Point) float64 {
	var ret float64
	for i := 0; i < p.Dim(); i++ {
		tmp := p.GetValue(i) - other.GetValue(i)
		ret += tmp * tmp
	}
	return ret
}

func (p *EuclideanPoint) PlaneDistance(val float64, dim int) float64 {
	tmp := p.GetValue(dim) - val
	return tmp * tmp
}

func NewEuclideanPoint(vals ...float64) *EuclideanPoint {
	ret := &EuclideanPoint{
		PointBase: NewPointBase(vals),
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
	{
		p1 := NewEuclideanPoint(0.0, 0.0, 0.0)
		p2 := NewEuclideanPoint(0.0, 0.0, 1.0)
		p3 := NewEuclideanPoint(0.0, 1.0, 0.0)
		p4 := NewEuclideanPoint(1.0, 0.0, 0.0)
		points := make([]Point, 0)
		points = append(points, p1)
		points = append(points, p2)
		points = append(points, p3)
		points = append(points, p4)
		tree := NewKDTree(points)
		ans := tree.KNN(NewEuclideanPoint(0.0, 0.0, 0.1), 2)
		checkKNNResult(t, ans, p1, p2)
	}
	// case 2
	{
		p1 := NewEuclideanPoint(0.0, 0.0, 0.0)
		p2 := NewEuclideanPoint(0.0, 0.0, 1.0)
		p3 := NewEuclideanPoint(0.0, 1.0, 0.0)
		p4 := NewEuclideanPoint(1.0, 0.0, 0.0)
		p5 := NewEuclideanPoint(0.0, 0.0, 0.0)
		p6 := NewEuclideanPoint(0.0, 0.0, 0.1)
		p7 := NewEuclideanPoint(1.0, 1.0, 1.0)
		points := make([]Point, 0)
		points = append(points, p1)
		points = append(points, p2)
		points = append(points, p3)
		points = append(points, p4)
		points = append(points, p5)
		points = append(points, p6)
		points = append(points, p7)
		tree := NewKDTree(points)
		ans := tree.KNN(NewEuclideanPoint(0.0, 0.0, 0.0), 3)
		checkKNNResult(t, ans, p1, p5, p6)
		ans = tree.KNN(NewEuclideanPoint(0.0, 0.0, 0.0), 4)
		if !equal(ans[3], p2) && !equal(ans[3], p3) && !equal(ans[3], p4) {
			t.Error("KNN results are wrong")
		}
		ans = tree.KNN(NewEuclideanPoint(0.0, 0.0, 0.0), 7)
		if !equal(ans[6], p7) {
			t.Error("KNN results are wrong")
		}
	}
	// case 3
	{
		points := []Point{
			NewEuclideanPoint(0.0, 0.0, 0.0),
			NewEuclideanPoint(0.0, 0.0, 1.0),
			NewEuclideanPoint(0.0, 1.0, 0.0),
			NewEuclideanPoint(1.0, 0.0, 0.0),
			NewEuclideanPoint(0.0, 0.0, 0.0),
			NewEuclideanPoint(0.0, 0.0, 0.1),
			NewEuclideanPoint(1.0, 1.0, 1.0),
			NewEuclideanPoint(0.1, 0.1, 0.1),
		}
		tree := NewKDTree(points)
		ans := tree.KNN(NewEuclideanPoint(0.0, 0.0, 0.0), 7)
		if len(ans) != 7 {
			t.Errorf("expected 7 points, actual: %v", len(ans))
		}
	}
	// case 4
	{
		points := []Point{
			NewEuclideanPoint(0.0, 0.0, 0.0),
			NewEuclideanPoint(0.0, 0.0, 0.0),
			NewEuclideanPoint(0.0, 0.0, 0.0),
		}
		tree := NewKDTree(points)
		ans := tree.KNN(NewEuclideanPoint(0.0, 0.0, 0.0), 3)
		if len(ans) != 3 {
			t.Errorf("expected 3 points, actual: %v", len(ans))
		}
	}
}
