package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)

	randomNum := random.Intn(10)
	fmt.Println(randomNum)

}
