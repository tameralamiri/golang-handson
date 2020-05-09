package main

import (
	"fmt"
)

func Abs(x float64) float64{
	if x < 0 {
		return -x
	}
	return x
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string{
	return fmt.Sprintf("can not Sqrt negative number:%v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0{
		return 0, ErrNegativeSqrt(x)
	}
	z := float64(1)
	for i:= 0;i<10;i++{
	    prev_z := z
		z -= (z*z - x) / (2*z)
		if Abs(prev_z - z) < 0.0001 {
			break
		} 
	}
	return z, nil
	
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
