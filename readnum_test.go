package readnum

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func test1(text string) float64 {
	value, ok, err := Float(bufio.NewReader(strings.NewReader(text)))
	fmt.Printf("'%s' -> %f(%v)%v\n", text, value, ok, err)
	return value
}

func TestMain(t *testing.T) {
	test1("")
	test1(".")
	test1(".1")
	test1("3")
	test1("3 ")
	test1("-3")
	test1("-3 ")
	test1("3.14")
	test1("3.14 ")
	test1("3.14e2")
	test1("3.14e2 ")
	test1("3.14e-2")
	test1("3.14e-2 ")
	test1("-3.14e+1")
	test1("-3.14e+1 ")
	test1("-0.314e+1")
	test1("-0.314e+1 ")
}
