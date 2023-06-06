package main 

import(
	"fmt"
	"time"
	"pi/internal/pi"
)

func main() {
	start := time.Now()
	
	// fmt.Println("Pi:", pi.ComputePi(1.0, 1_000_000_000)) // 51.59 s
	fmt.Println("Pi:", pi.ComputePi(1.0, 100_000_000)) // 5.18 s

	elapsed := time.Since(start)
	fmt.Println("Time: ", elapsed)
}
