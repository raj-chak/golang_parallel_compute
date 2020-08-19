package main

import (
	"fmt"
	"runtime"
)

const (
	np = 4
	n = 10000
)

var h float64
var pi float64	
	/*
//!+output
$ go build gopl.io/raj/parallel
$ ./parallel
*/


func f(a float64) float64 {
   return 4.0/(1.0 + a * a)
} 



func chunk(start, end int64, c chan float64) { 
	var sum float64 = 0.0
	for i:= start; i < end; i++ {
		x := h * (float64(i) + 0.5)
		sum += f(x)
	}
	c <- sum * h 
}


func main()  {
	runtime.GOMAXPROCS(np); 
	h = 1.0/float64(n) 
	
	c := make(chan float64, np)
	for i:=0;i<np;i++{
		go chunk(int64(i)*n/int64(np), (int64(i)+1)*n/int64(np), c)
	}
	for i:=0; i < np; i++ {
		pi += <-c 
	}
	fmt.Println(pi)
}