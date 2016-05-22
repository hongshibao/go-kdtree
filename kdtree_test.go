package kdtree

import (
	"fmt"
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

func TestKNN(t *testing.T) {
	points := make([]Point, 0)
	points = append(points, NewEuclideanPoint(0.0, 0.0, 0.0))
	points = append(points, NewEuclideanPoint(0.0, 0.0, 1.0))
	points = append(points, NewEuclideanPoint(0.0, 1.0, 0.0))
	points = append(points, NewEuclideanPoint(1.0, 0.0, 0.0))
	tree := NewKDTree(points)
	ans := tree.KNN(NewEuclideanPoint(0.0, 0.0, 0.1), 3)
	fmt.Println("len of ans:", len(ans))
	for _, p := range ans {
		for i := 0; i < p.Dim(); i++ {
			fmt.Print(p.GetValue(i), ", ")
		}
		fmt.Println()
	}
}
