package main

import (
	"fmt"
	"runtime"
)

const m int64 = 4
const mm int64 = m * m

var random [m * m]int64 //this is only 16MB for m = 2000
var R [m][m]int64

type Res struct {
	sum       int64
	threadNUm int
	x         int
	y         int
	direction string
}

func assignRandomNumbers(a *[m * m]int64) {
	for i := 0; i < len(a); i++ {
		k := int64(i + 1)
		if i < 55 {
			a[i] = (100003-200003*k+300007*k*k*k)%1000000 - 500000
		} else {
			a[i] = (a[i-24]+a[i-55]+1000000)%1000000 - 500000
		}
	}
}
func arrayToMatrix(ar *[m * m]int64, R *[m][m]int64) {
	for k := 0; k < len(ar); k++ {
		j := int64(k)
		R[j/m][j%m] = ar[k]
	}
}
func getSumV(R *[m][m]int64, start int64, step int64, ch chan Res) { //vertical sum
	var r Res
	r.sum = -1000000
	for i := 0; i < len(R[0]); i += int(step) {
		var sum int64 = 0
		for j := 0; j < len(R); j++ {
			sum += R[j][i]
		}
		if sum > r.sum {
			r.sum = sum
			r.direction = "V"
			r.threadNUm = int(start)
			r.x = 0
			r.y = i
		}
	}
	ch <- r
}
func getSumH(R *[m][m]int64, start int64, step int64, ch chan Res) { //horizontal sum
	var r Res
	r.sum = -1000000
	for i := 0; i < len(R); i += int(step) {
		var sum int64 = 0
		for j := 0; j < len(R[i]); j++ {
			sum += R[i][j]
		}
		if sum > r.sum {
			r.sum = sum
			r.direction = "H"
			r.threadNUm = int(start)
			r.x = i
			r.y = 0
		}
	}
	ch <- r
}
func getSumD(R *[m][m]int64, start int64, step int64, ch chan Res) { //diagonal sum
	var r Res
	r.sum = -1000000
	var sum int64
	for i := 0; i < len(R); i += int(step) { //start at left column going right and down
		sum = 0
		for j := 0; j < len(R); j++ {
			if i+j >= len(R) {
				break
			}
			sum += R[i+j][j]
		}
		if sum > r.sum {
			r.sum = sum
			r.direction = "D"
			r.threadNUm = int(start)
			r.x = i
			r.y = 0
		}
	}
	for i := 0; i < len(R[0]); i += int(step) { //start at to top row going right and down
		sum = 0
		for j := 0; j < len(R); j++ {
			if i+j >= len(R[i]) {
				break
			}
			sum += R[j][j+i]
		}
		if sum > r.sum {
			r.sum = sum
			r.direction = "D"
			r.threadNUm = int(start)
			r.x = 0
			r.y = i
		}
	}
	ch <- r
}
func getSumAD(R *[m][m]int64, start int64, step int64, ch chan Res) { // anti diagnonal sum
	var r Res
	r.sum = -1000000
	var sum int64 = 0
	for i := 0; i < len(R); i += int(step) { //start at the left column going right and up
		sum = 0
		for j := 0; j < len(R); j++ {
			if i-j < 0 {
				break
			}
			sum += R[i-j][j]
		}
		if sum > r.sum {
			r.sum = sum
			r.direction = "AD"
			r.threadNUm = int(start)
			r.x = i
			r.y = 0
		}
	}
	for i := 0; i < len(R[0]); i += int(step) { //start at bottom row going right and up
		sum = 0
		for j := 0; j < len(R); j++ {
			if i+j >= len(R) {
				break
			}
			sum += R[len(R)-1-j][i+j]
		}
		if sum > r.sum {
			r.sum = sum
			r.direction = "AD"
			r.threadNUm = int(start)
			r.x = len(R)
			r.y = i
		}
	}
	ch <- r
}
func getAllSums(R *[m][m]int64) Res {
	ch := make(chan Res)
	var ret Res
	ret.sum = -1000000
	numThread := runtime.NumCPU()
	for i := 0; i < numThread; i++ {
		go getSumH(R, int64(i), int64(numThread), ch)
	}
	for i := 0; i < numThread; i++ {
		r := <-ch
		fmt.Printf("sum = %d\n", r.sum)
		if r.sum > ret.sum {
			ret = r
		}
		go getSumV(R, int64(i), int64(numThread), ch)
	}
	fmt.Printf("largest horizontal sum = %d\n", ret.sum)
	printRes(ret)
	for i := 0; i < numThread; i++ {
		r := <-ch
		if r.sum > ret.sum {
			ret = r
		}
		go getSumD(R, int64(i), int64(numThread), ch)
	}
	fmt.Printf("largest vertical sum = %d\n", ret.sum)
	for i := 0; i < numThread; i++ {
		r := <-ch
		if r.sum > ret.sum {
			ret = r
		}
		go getSumAD(R, int64(i), int64(numThread), ch)
	}
	fmt.Printf("largest diagonal sum = %d\n", ret.sum)
	for i := 0; i < numThread; i++ {
		r := <-ch
		if r.sum > ret.sum {
			ret = r
		}
	}
	return ret

}
func main() {
	assignRandomNumbers(&random)
	arrayToMatrix(&random, &R)

	fmt.Println(R)
	sum := getAllSums(&R)
	ch := make(chan Res)
	go getSumH(&R, 0, 1, ch)
	hsum := <-ch
	printRes(hsum)
	fmt.Printf("Max sum is: %d\n", sum.sum)

}

func printRes(r Res) {
	fmt.Printf("sum=%d dir=%s at x=%d y=%d from thread %d\n", r.sum, r.direction, r.x, r.y, r.threadNUm)
}
