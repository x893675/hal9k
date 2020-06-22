package main

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
)

func main() {
	for i:=0; i < 10; i++ {
		fmt.Println(RangeRand(0, 0))
	}
}

func RangeRand(min, max int64) int64 {
	if min > max {
		panic("the min is greater than max!")
	}
	if min < 0 {
		f64Min := math.Abs(float64(min))
		i64Min := int64(f64Min)
		result, _ := rand.Int(rand.Reader, big.NewInt(max+1+i64Min))
		return result.Int64() - i64Min
	} else {
		result, _ := rand.Int(rand.Reader, big.NewInt(max-min+1))
		return min + result.Int64()
	}
}
