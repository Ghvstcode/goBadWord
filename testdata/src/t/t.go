package t

import (
	"fmt"
	"testing"
)

//comment test
//Second Comment
func Hellosweeties(kill, nextnamesssss int, b func(a int)) {
	vns := "testString" + "hiiiiiiiiii"
	vxns := 3 * 2
	fmt.Println(vns)
	fmt.Println(vxns)
	//b.name = vns
	kill++
	nextnamesssss++
}

func TestIfTestFunctionsArentSkipped(t *testing.T) {
	i := 1
	if i > 2 {
		if i > 2 {
		}
		if i > 2 {
		}
		if i > 2 {
		}
		if i > 2 {
		}
	} else {
		if i > 2 {
		}
		if i > 2 {
		}
		if i > 2 {
		}
		if i > 2 {
		}
	}

	if i > 2 {
	}
}
