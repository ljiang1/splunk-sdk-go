package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func sdfTestSomething(t *testing.T) {
	// test stuff here...

	now := time.Now()
	fmt.Printf("%s\n", now)
	len := len(now.String())
	var x string = now.String()[len-1 : len]
	var y string = now.String()[len-2 : len]
	z, u := strconv.Atoi(x)
	k, m := strconv.Atoi(y)
	fmt.Printf("%d,%d,%d,%d\n", z, k, u, m)
	if avg(z, k) < 20 {
		print(rand.Int())
		t.Error("????")
	} else {
		t.Error("hahah error")
	}

}

func TestWeb(t *testing.T) {
	// test stuff here...
	if webrequest() {
		print("=======okkkk")
	} else {
		t.Error("===can't visist web")
	}
}
