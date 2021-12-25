package helpers

import (
	"fmt"
	"math/rand"
	"time"
)

// Common ...
type Common struct {
	NameType    string
	VersionType int
}

// TestGreater ...
func TestGreater(a int, b int) {
	if a > b {
		fmt.Println(a, "is greater")
	} else {
		fmt.Println(b, "is greater")
	}
}

// RandomNum ...
func RandomNum(n int) int {

	rand.Seed(time.Now().UnixNano())
	val := rand.Intn(n)
	return val
}
