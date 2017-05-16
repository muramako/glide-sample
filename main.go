package main

import (
	"fmt"
	"github.com/cznic/mathutil"
	"math"
)

/*
$ # Usage e.g.:
$ go run example.go -max 1024 > mathutil.dat # generate 1kB of "random" data
*/
func main() {
	r, _ := mathutil.NewFC32(math.MinInt32, math.MaxInt32, true)
  fmt.Printf("%#v", r)
}
