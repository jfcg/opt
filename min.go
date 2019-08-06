// Package opt provides some optimization routines
package opt

/*	Rectangular grid of fn() values centered at x0,y0
	^ 6 7 8
	y 3 4 5 // neighbor indices
	  0 1 2
	    x >
*/
func shiftRectGrid(k int, fv *[9]float64) {
	switch k {
	case 0: // down left
		fv[4], fv[5], fv[7], fv[8] = fv[0], fv[1], fv[3], fv[4]
		fv[0], fv[1], fv[2], fv[3], fv[6] = 0, 0, 0, 0, 0
	case 2: // down right
		fv[3], fv[4], fv[6], fv[7] = fv[1], fv[2], fv[4], fv[5]
		fv[0], fv[1], fv[2], fv[5], fv[8] = 0, 0, 0, 0, 0
	case 6: // up left
		fv[1], fv[2], fv[4], fv[5] = fv[3], fv[4], fv[6], fv[7]
		fv[0], fv[3], fv[6], fv[7], fv[8] = 0, 0, 0, 0, 0
	case 8: // up right
		fv[0], fv[1], fv[3], fv[4] = fv[4], fv[5], fv[7], fv[8]
		fv[2], fv[5], fv[6], fv[7], fv[8] = 0, 0, 0, 0, 0
	case 1: // down
		fv[6], fv[7], fv[8] = fv[3], fv[4], fv[5]
		fv[3], fv[4], fv[5] = fv[0], fv[1], fv[2]
		fv[0], fv[1], fv[2] = 0, 0, 0
	case 3: // left
		fv[2], fv[5], fv[8] = fv[1], fv[4], fv[7]
		fv[1], fv[4], fv[7] = fv[0], fv[3], fv[6]
		fv[0], fv[3], fv[6] = 0, 0, 0
	case 5: // right
		fv[0], fv[3], fv[6] = fv[1], fv[4], fv[7]
		fv[1], fv[4], fv[7] = fv[2], fv[5], fv[8]
		fv[2], fv[5], fv[8] = 0, 0, 0
	case 7: // up
		fv[0], fv[1], fv[2] = fv[3], fv[4], fv[5]
		fv[3], fv[4], fv[5] = fv[6], fv[7], fv[8]
		fv[6], fv[7], fv[8] = 0, 0, 0
	}
}

// FindMin minimizes positive fn() over (x,y) rectangular grid starting at (x0,y0) with
// (±dx,±dy) steps for up to |r| runs, halving step sizes between runs. If given pr(), calls it
// at every new optimal value (could be used for printing progress). Calls fn() at least once.
// If r<0, diagonal steps are also allowed.
// Returns optimal point & fn() value, and number of calls to fn().
func FindMin(r, x0, y0, dx, dy int, fn func(x, y int) float64,
	pr func(int, int, float64)) (int, int, float64, uint) {

	var fv [9]float64  // fn() values
	fv[4] = fn(x0, y0) // center
	nc := uint(1)      // number of calls to fn()
	if pr != nil {
		pr(x0, y0, fv[4])
	}

	var Ix, Iy [9]int          // increment arrays
	start, end, inc := 1, 7, 2 // loop vars
	if r < 0 {
		start, end, inc = 0, 8, 1 // allow diagonal steps
		r = -r
	}

	for ; r > 0; r-- {

		// check dx/dy, fix loop vars
		if dy == -dy {
			if dx == -dx { // halt if both are 0/minInt
				return x0, y0, fv[4], nc
			}
			start, end, inc, dy = 3, 5, 1, 0 // x only

		} else if dx == -dx {
			start, end, inc, dx = 1, 7, 3, 0 // y only
		}

		// set increments
		Ix[0], Ix[3], Ix[6] = -dx, -dx, -dx
		Ix[2], Ix[5], Ix[8] = dx, dx, dx

		Iy[0], Iy[1], Iy[2] = -dy, -dy, -dy
		Iy[6], Iy[7], Iy[8] = dy, dy, dy

		for {
			k := 4
			for i := start; i <= end; i += inc {

				if fv[i] > 0 { // center or known non-optimal?
					continue
				}

				fv[i] = fn(x0+Ix[i], y0+Iy[i])
				nc++

				if fv[i] < fv[k] { // better neighbor?
					k = i
				}
			}

			if k == 4 {
				fv[8], fv[7], fv[6], fv[5] = 0, 0, 0, 0
				fv[3], fv[2], fv[1], fv[0] = 0, 0, 0, 0
				break // zero all except center which is best
			}

			x0 += Ix[k] // switch to best neighbor
			y0 += Iy[k]
			if pr != nil {
				pr(x0, y0, fv[k])
			}

			shiftRectGrid(k, &fv)
		}

		dx /= 2 // halve steps
		dy /= 2
	}
	return x0, y0, fv[4], nc
}

/*	Triangular grid of fn() values centered at x0,y0
	^  5 6
	y 2 3 4 // neighbor indices
	   0 1
	    x >
*/
func shiftTriGrid(k int, fv *[7]float64) {
	switch k {
	case 0: // down left
		fv[3], fv[4], fv[5], fv[6] = fv[0], fv[1], fv[2], fv[3]
		fv[0], fv[1], fv[2] = 0, 0, 0
	case 1: // down right
		fv[2], fv[3], fv[5], fv[6] = fv[0], fv[1], fv[3], fv[4]
		fv[0], fv[1], fv[4] = 0, 0, 0
	case 4: // right
		fv[0], fv[2], fv[3], fv[5] = fv[1], fv[3], fv[4], fv[6]
		fv[1], fv[4], fv[6] = 0, 0, 0
	case 6: // up right
		fv[0], fv[1], fv[2], fv[3] = fv[3], fv[4], fv[5], fv[6]
		fv[4], fv[5], fv[6] = 0, 0, 0
	case 5: // up left
		fv[0], fv[1], fv[3], fv[4] = fv[2], fv[3], fv[5], fv[6]
		fv[2], fv[5], fv[6] = 0, 0, 0
	case 2: // left
		fv[1], fv[3], fv[4], fv[6] = fv[0], fv[2], fv[3], fv[5]
		fv[0], fv[2], fv[5] = 0, 0, 0
	}
}

// FindMinTri minimizes positive fn() over (x,y) triangular grid starting at (x0,y0) with
// (±dx,±dy) steps for up to r runs, halving step sizes between runs. If given pr(), calls it
// at every new optimal value (could be used for printing progress). Calls fn() at least once.
// Choose dy ~ 0.866*dx for equilateral grid. Faster than FindMin() via lesser calls to fn().
// Returns optimal point & fn() value, and number of calls to fn().
func FindMinTri(r, x0, y0, dx, dy int, fn func(x, y int) float64,
	pr func(int, int, float64)) (int, int, float64, uint) {

	var fv [7]float64  // fn() values
	fv[3] = fn(x0, y0) // center
	nc := uint(1)      // number of calls to fn()
	if pr != nil {
		pr(x0, y0, fv[3])
	}

	var Ix, Iy [7]int          // increment arrays
	start, end, inc := 0, 6, 1 // loop vars

	for ; r > 0; r-- {

		hx := dx / 2
		// check dx/dy, fix loop vars
		if dy == -dy {
			if dx == -dx { // halt if both are 0/minInt
				return x0, y0, fv[3], nc
			}
			start, end, inc, dy = 2, 4, 1, 0 // x only

		} else if dx == -dx {
			start, end, inc, dx, hx = 1, 5, 2, 0, 0 // y only
		} else if hx == 0 {
			start, end, inc = 1, 5, 1 // cross
		}

		// set increments
		Ix[4], Ix[5], Ix[6] = dx, -hx, dx-hx
		Ix[2], Ix[1], Ix[0] = -dx, hx, hx-dx

		Iy[0], Iy[1], Iy[5], Iy[6] = -dy, -dy, dy, dy

		for {
			k := 3
			for i := start; i <= end; i += inc {

				if fv[i] > 0 { // center or known non-optimal?
					continue
				}

				fv[i] = fn(x0+Ix[i], y0+Iy[i])
				nc++

				if fv[i] < fv[k] { // better neighbor?
					k = i
				}
			}

			if k == 3 {
				fv[6], fv[5], fv[4] = 0, 0, 0
				fv[2], fv[1], fv[0] = 0, 0, 0
				break // zero all except center which is best
			}

			x0 += Ix[k] // switch to best neighbor
			y0 += Iy[k]
			if pr != nil {
				pr(x0, y0, fv[k])
			}

			shiftTriGrid(k, &fv)
		}

		dx = hx // halve steps
		dy /= 2
	}
	return x0, y0, fv[3], nc
}
