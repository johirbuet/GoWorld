package main
import (
	"fmt"
	"math/big"
)
func main() {
	var b1,b2,b3,bigSum big.Float;
	b1.SetFloat64(23.5)
	b2.SetFloat64(100.5)
	b3.SetFloat64(100.0)
	bigSum.Add(&b1,&b2).Add(&bigSum,&b3)
	fmt.Printf("Bigsum = %0.10g\n",&bigSum)
}