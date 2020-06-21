package main

import (
	"log"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func shuffle(n int) []int {
	b := rand.Perm(n)
	return b
}

func main() {
	log.Println(shuffle(3))
}
