package main
import (
	"fmt"
	"time"
)
func main() {
	t := time.Date(2009,time.November,10,23,0,0,0,time.UTC)
	fmt.Printf("Go launced on %s",t)
	now := time.Now()
	fmt.Printf("Current time %s",now)
}