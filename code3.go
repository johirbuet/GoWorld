package main
import (
	"fmt"
	"strings"
)
func main() {
	lvalue := "hello"
	rvalue := "HELLO"
	fmt.Println("Equal? ",(lvalue == rvalue))
	fmt.Println("Equal? ",strings.EqualFold(lvalue,rvalue))
}