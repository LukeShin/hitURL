package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan bool)
	people := [2]string{"Nico", "Luke"}
	for _, person := range people {
		go isSexy(person, c)
	}
	fmt.Println(<-c)
	fmt.Println(<-c)
}

func isSexy(person string, c chan bool) {
	time.Sleep(time.Second * 1)
	fmt.Println(person)
	c <- true
}
