package t

import (
	"fmt"
	"testing"
)

//comment test
//Second Comment
func Hellosweeties(nameoh, nextnamesssss int) {
	vns := "testString"
	fmt.Println(vns)
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
