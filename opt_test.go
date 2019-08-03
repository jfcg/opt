package opt

import "testing"

func f1(x, y int) float64 {
	a := float64(x) - 3.3
	b := float64(y) - 4.4
	return a*a + b*b
}

func Test1(t *testing.T) {
	x, y, v, n := FindMin(1, 0, 0, 2, 2, f1, nil)
	if x != 4 || y != 4 || n != 19 {
		t.Fatal("not 4,4:", x, y, v, n)
	}

	x, y, v, n = FindMin(2, 0, 0, 2, 2, f1, nil)
	if x != 3 || y != 4 || n != 30 {
		t.Fatal("not 3,4:", x, y, v, n)
	}
}
