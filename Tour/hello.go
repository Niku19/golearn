package main

import (
	"fmt"
	"math"
	"time"

	"github.com/Niku19/golearn/Tour/bignumbers"
)

func main() {
	fmt.Println("Hello")
	fmt.Println(time.Now().Weekday())
	bignumbers.Biggy()
	math.Sqrt(5)
	var p *int
	i := 0
	p = &i
	fmt.Print(p)
}
