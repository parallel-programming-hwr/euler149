package main

import "fmt"

const m int64 = 10
const mm int64 = m * m

var random [m * m]int64 //this is only 16MB for m = 2000
var R[m][m]int64

func assignRandomNumbers(a *[m*m]int64) {
	for i := 0; i < len(a); i++ {
		k:=int64(i+1)
		if i < 55{
			a[i] = (100003-200003*k+300007*k*k*k)%1000000 - 500000
		}else{
			a[i]=(a[i-24] +a[i-55]+1000000)%1000000-500000
		}
	}
}
func arrayToMatrix(ar *[m*m]int64, R  *[m][m]int64 ){
	for k:=0;k<len(ar);k++{
		j:=int64(k)
		R[j/m][j%m]=ar[k]
	}
}
func getSumH(R *[m][m]int64,start int64, step int64){

}
func main() {
	assignRandomNumbers(&random)
	arrayToMatrix(&random,&R)
	fmt.Println(R)
	fmt.Println(R[0][9])
	fmt.Println(R[9][9])
}
