package main

import (
	"fmt"
	"math"
	"testing"
)

//question1
type triangle struct {
	base   float64
	height float64
}

//question2
func (t triangle) Area() float64 {
	return 0.5 * (t.base * t.height)
}

//question3
func (t triangle) Perimeter() float64 {
	p := math.Sqrt(math.Pow(t.base, 2) + math.Pow(t.height, 2))
	return p + t.base + t.height
}

// question 6
type circle struct {
	radius float64
}

//question 7
func (c circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

//question 8
func (c circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

//question 9
func square(s float64) (float64, float64) {
	area := math.Pow(s, 2)
	perimeter := 4 * s
	return area, perimeter
}

//question 10
func TestSquare(t *testing.T) {
	a, p := square(4)
	if a != 16 {
		t.Errorf("Expected 24, got %v", a)
	}
	if p != 16 {
		t.Errorf("Expected 16, got &v", p)
	}
}

//question 11
/*func TestArea(t *testing.T) {
	t1 := triangle{
		base:   10,
		height: 12,
	}

	result := t1.area()

	if result != 100 {
		t.Errorf(" expected %v", result)
	}
}
*/
func main() {
	//questrion 4 -5
	tri1 := triangle{
		base:   10,
		height: 5,
	}

	tri2 := triangle{
		base:   12,
		height: 6,
	}

	circ1 := circle{
		radius: 7,
	}
	fmt.Println(tri1.Area())
	fmt.Println(tri2.Perimeter())
	fmt.Println(circ1.Area())

}
