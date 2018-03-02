package main

import (
	"testing"
	"math/rand"
	"time"
	"strconv"
	"fmt"
)

func TestSomething(t *testing.T) {
	// test stuff here...

	now := time.Now();
	fmt.Printf("%s\n", now)
	len:=len(now.String())
	var x string = now.String()[len-1:len]
	var y string = now.String()[len-2:len]
	z, u := strconv.Atoi(x)
	k, m := strconv.Atoi(y)
	fmt.Printf("%d,%d\n", z, k)
	if avg(z, k) < 20 {
		print(rand.Int())
		t.Error("????")
	} else {
		t.Error("hahah error")
	}

	if false {
		fmt.Printf("%s,%s", u, m)
	}
}