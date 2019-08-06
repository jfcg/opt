package opt

import "testing"

func f1(x, y int) float64 {
	a := float64(x) - 3.3
	b := float64(y) - 4.4
	return a*a + b*b
}

type res struct {
	f func(r, x0, y0, dx, dy int, fn func(x, y int) float64,
		pr func(int, int, float64)) (int, int, float64, uint)
	r, x, y int
	n       uint
}

var rs = [...]res{
	{FindMin, 1, 4, 4, 14},
	{FindMin, -1, 4, 4, 19},
	{FindMin, 2, 3, 4, 21},
	{FindMin, -2, 3, 4, 30},
	{FindMinTri, 1, 4, 4, 16},
	{FindMinTri, 2, 3, 4, 23},
}

func Test1(t *testing.T) {
	for i, v := range rs {
		x, y, _, n := v.f(v.r, 0, 0, 2, 2, f1, nil)
		if x != v.x || y != v.y || n != v.n {
			t.Fatal(i, "gave:", x, y, n, "expected:", v.x, v.y, v.n)
		}
	}
}
