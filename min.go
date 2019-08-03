// Package opt provides some optimization routines
package opt

/*	FindMin minimizes positive fn() over (x,y) grid starting at (x0,y0) with (±dx,±dy)
	steps for up to r runs, halving step sizes between runs. If given pr(), calls it at
	every new optimal value (could be used for printing progress). Calls fn() at least once.
	Returns optimal point & fn() value, and number of calls to fn().
*/
func FindMin(r, x0, y0, dx, dy int, fn func(x, y int) float64,
	pr func(int, int, float64)) (int, int, float64, uint) {

	/*	3x3 grid of fn() values centered at x0,y0
		^ 6 7 8
		y 3 4 5 // neighbor indices
		  0 1 2
		    x >
	*/
	var fv [9]float64
	fv[4] = fn(x0, y0) // center
	nc := uint(1)      // number of calls to fn()
	if pr != nil {
		pr(x0, y0, fv[4])
	}

	var Ix, Iy [9]int          // increment arrays
	start, end, inc := 0, 8, 1 // loop vars

	for ; r > 0; r-- {

		// check dx/dy and fix loop vars
		if dy == -dy {
			if dx == -dx { // halt if both are 0/minInt
				return x0, y0, fv[4], nc
			}

			start, end = 3, 5
		} else if dx == -dx {
			start, end, inc = 1, 7, 3
		}

		// initialize increment arrays
		for i := 6; i >= 0; i -= 3 {
			Ix[i], Ix[i+2] = -dx, dx
		}
		for i := 2; i >= 0; i-- {
			Iy[i], Iy[i+6] = -dy, dy
		}

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

			// update fv grid
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
			default: // up
				fv[0], fv[1], fv[2] = fv[3], fv[4], fv[5]
				fv[3], fv[4], fv[5] = fv[6], fv[7], fv[8]
				fv[6], fv[7], fv[8] = 0, 0, 0
			}
		}

		dx /= 2 // halve steps
		dy /= 2
	}
	return x0, y0, fv[4], nc
}
