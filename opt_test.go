/*	Copyright (c) 2019, Serhat Şevki Dinçer.
	This Source Code Form is subject to the terms of the Mozilla Public
	License, v. 2.0. If a copy of the MPL was not distributed with this
	file, You can obtain one at http://mozilla.org/MPL/2.0/.
*/

package opt

import "testing"

// to be minimized
func f1(x, y int) float64 {
	a := float64(x) - 3.3
	b := float64(y) - 4.4
	return a*a + b*b
}

var (
	px, py int
	po     float64
	pn     uint
)

// called at every new optima
func pro(x, y int, o float64) {
	px, py, po = x, y, o
	pn++
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
		pn = 0
		x, y, o, n := v.f(v.r, 0, 0, 2, 2, f1, pro)
		if x != v.x || y != v.y || n != v.n || // expected result
			x != px || y != py || o != po || // last pro() call records result
			pn == 0 || n < pn { // f1() calls >= pro() calls > 0
			t.Fatal(i, "gave:", x, y, n, "expected:", v.x, v.y, v.n)
		}
	}
}
