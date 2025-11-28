/*
  (40 points)

  Implement the `distance` function such that it returns the distance
  between two 2D points in a plane, or the distance between two 3D
  spatial points. If the input parameters are not compatible point
  types, return `0.0, errors.New("Unsupported")`.

  The distance between two 2D points (x1, y1) and (x2, y2) is:
  `math.Sqrt(math.Pow(x1 - x2) + math.Pow(y1 - y2))`.

  The distance between two 3D spatial spatial points
  (x1, y1, z1) and (x2, y2, z2) is:
  `math.Sqrt(math.Pow(x1 - x2) + math.Pow(y1 - y2) + math.Pow(z1 - z2))`.

  A 2D point (x, y) is stored as a Point2D{x, y} while a 3D
  point (x, y, z) is stored as a Point3D{x, y, z}.

  Hint:
  - Use type assertion or type switch.
*/

package main

import (
	"errors"
	"fmt"
	"math"
)

type Point2D struct {
	x float64
	y float64
}

type Point3D struct {
	x float64
	y float64
	z float64
}

func distance(x any, y any) (float64, error) {
	// ----- <code begin> -----
	switch p1 := x.(type) {
	case Point2D:
		if p2, ok := y.(Point2D); ok {
			return math.Sqrt(math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y-p2.y, 2)), nil
		}
	case Point3D:
		if p2, ok := y.(Point3D); ok {
			return math.Sqrt(math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y-p2.y, 2) + math.Pow(p1.z-p2.z, 2)), nil
		}
	}
	return 0.0, errors.New("Unsupported")
	// ----- <code end> -----
}

func main() {
	p1 := Point2D{3.0, 4.0}
	p2 := Point2D{6.0, 8.0}
	if d, err := distance(p1, p2); err == nil {
		fmt.Printf("%.2f\n", d) // expected output: 5.00
	}

	q1 := Point3D{4.0, 2.0, 1.0}
	q2 := Point3D{1.0, 6.0, 3.0}
	if d, err := distance(q1, q2); err == nil {
		fmt.Printf("%.2f\n", d) // expected output: 5.39
	}
}
