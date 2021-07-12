package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	people := [3]string{"Nico", "Luke", "John"}
	for _, person := range people {
		go isSexy(person, c)
	}
	for i := 0; i < len(people); i++ {
		fmt.Println("Waiting for ", i)
		fmt.Println(<-c)
	}
}

func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 3)
	c <- person + " is sexy."
}
