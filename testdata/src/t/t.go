package t

import (
	"fmt"
	"testing"
)

type vic struct {
	name string
}

//comment test
//Second Comment
func Hellosweeties(nameoh, nextnamesssss int, b func(a int)) {
	vns := "testString"
	fmt.Println(vns)
	//b.name = vns
	nameoh++
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
