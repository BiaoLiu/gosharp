package utils

import (
	"fmt"
	"testing"
)

func TestCastDecimalFromInt(t *testing.T) {
	d := CastDecimalFromInt(10000000)
	fmt.Println(d.String())
}

func TestFloat64Mul(t *testing.T) {
	d := Float64Mul(1.00, 100)
	fmt.Println(d)
}
