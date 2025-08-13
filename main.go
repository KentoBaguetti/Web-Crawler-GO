package main

import (
	"fmt"

	"github.com/KentoBaguetti/Web-Crawler-GO/datastructures"
	"github.com/KentoBaguetti/Web-Crawler-GO/utils"
)

func main() {
	fmt.Println("Init")

	q := datastructures.Queue{Elements: make([] string, 0), Length: 0}

	q.Enqueue("Hello")
	q.Enqueue("Kentaro")
	q.Dequeue()

	fmt.Println(q.Elements)

	s := datastructures.Set{Elements: make(map[string]bool, 0), Length: 0}

	s.Add("testUrl")
	s.Add("testUrl")

	fmt.Println(s.Elements)
	fmt.Println(s.Elements["testUrl"])
	fmt.Println(s.Elements["kentaro"])

	utils.TestFunction()

}




