# KDTree
Golang implementation of KD tree (https://en.wikipedia.org/wiki/K-d_tree) data structure

## Getting started

Use go tool to install the package in your packages tree:
```
go get github.com/hongshibao/go-kdtree
```
Then you can use it in import section of your Go programs:
```go
import "github.com/hongshibao/go-kdtree"
```
The package name is ```kdtree```.

## Basic example

First you need to implement the ```Point``` interface:
```go
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
```
Here is an example of implementing ```Point``` interface with square of Euclidean distance as the ```Distance``` definition:
```go
type EuclideanPoint struct {
	Point
	Vec []float64
}

func (p *EuclideanPoint) Dim() int {
	return len(p.Vec)
}

func (p *EuclideanPoint) GetValue(dim int) float64 {
	return p.Vec[dim]
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
```
Now you can create KD-tree from a list of points and get a list of k nearest neighbours for a target point:
```go
func NewEuclideanPoint(vals ...float64) *EuclideanPoint {
	ret := &EuclideanPoint{}
	for _, val := range vals {
		ret.Vec = append(ret.Vec, val)
	}
	return ret
}

func main() {
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
	targetPoint := NewEuclideanPoint(0.0, 0.0, 0.1)
	neighbours := tree.KNN(targetPoint, 2)
	for idx, p := range neighbours {
		fmt.Printf("Point %d: (%f", idx, p.GetValue(0))
		for i := 1; i < p.Dim(); i++ {
			fmt.Printf(", %f", p.GetValue(i))
		}
		fmt.Println(")")
	}
}
```
The returned k nearest neighbours are sorted by their distance with the target point.
