package main

import (
	"fmt"
)

type NegativeSqrtError float64

// This shows that types other than struct can also have methods.
func (f NegativeSqrtError) Error() string {
	return fmt.Sprintf("math: square root of negative number %g", float64(f))
}

func main() {
	errOfNegativeSqrt := NegativeSqrtError(-4)
	fmt.Println(errOfNegativeSqrt.Error())
}
