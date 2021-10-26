package t

import (
	"fmt"
	"math/rand"
	"time"
)

var thor string
var hulk string

func testtt(){
	for {
		go rand.Seed(time.Now().Unix())
	}


	thor = "newGreenEnergy"
	sth := "swing"
	fmt.Println(sth)
}

