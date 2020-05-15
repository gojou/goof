package main

import (
	"fmt"
	"log"
)

func main() {
	e := run()
	if e != nil {
		log.Fatalf("Fatal error:%v", e)
	}
}
func run() (e error) {
	u := "Mark"
	fmt.Printf("Hello %v!\n", u)
	return e
}
