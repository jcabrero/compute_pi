package pi 

import(
	"fmt"
	"math/rand"
	"math"
)

type Point struct{
	x, y float64;
}

// This function generates 'n' random floating point points, in range 'min' - 'max'.
func randPoints(min, max float64, n int) []Point {
    res := make([]Point, 2 * n)
    for i := range res {
        res[i].x = min + rand.Float64() * (max - min)
		res[i].y = min + rand.Float64() * (max - min)
    }	
    return res
}

// This function determines with Pithagoras whether point is in a circle of given radius.
// For compatibility, returns 1, if it is inside, 0 otherwise.
func isInCircle(point Point, radius float64) (int64) {
	x_square := point.x * point.x
	y_square := point.y * point.y

	d := math.Sqrt(x_square + y_square)
	if d <= radius {
		return 1
	} else {
		return 0
	}
}

func min(x, y int) (int){
	if x < y {
		return x
	} else {
		return y
	}
}

// I cannot keep a lot of Points in memory, that's why we generate points in batches of 1M.
const MAX_MEM_RAND = 1_000_000

func ComputePi(radius float64, n int) (pi float64){
	

	var in, total int64 = 0,0 

	for n_ := 0; n_ < n; n_ += MAX_MEM_RAND {
		// For the last iteration, we only produce the remaining numbers asked in n
		points := randPoints(-radius, radius, min(MAX_MEM_RAND, n - n_))
		total += int64(len(points))
		for i := 0; i < len(points); i += 1{
			in += isInCircle(points[i], radius)
		}
	}
	fmt.Println("Points (In): ", in, " (Total): ", total)
	
	area_ratio := float64(in) / float64(total)
	pi = 4 * area_ratio
	return pi
}
