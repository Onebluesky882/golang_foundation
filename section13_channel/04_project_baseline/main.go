// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"os"

// 	"time"
// )

// func running(track chan<- struct{}) {

// 	rand.Seed(time.Now().UnixNano())
// 	time.Sleep(time.Duration(10*time.Second + time.Duration(rand.Intn(100))*time.Millisecond))
// 	track <- struct{}{}
// }

// func main() {
// 	defer func() {
// 		fmt.Println("done")
// 	}()

// 	labChan := make(chan time.Time)

// 	go func() {
// 		lapTricker := time.NewTicker(1 * time.Second)
// 		for {
// 			labChan <- (<-lapTricker.C)
// 		}
// 	}()

// 	track1 := make(chan struct{})
// 	track2 := make(chan struct{})
// 	track3 := make(chan struct{})
// 	abort := make(chan struct{})

// 	go func() {
// 		os.Stdin.Read(make([]byte, 1))
// 		fmt.Println("abort to abort")
// 		abort <- struct{}{}
// 	}()

// 	go running(track1)
// 	go running(track2)
// 	go running(track3)

// 	for done := false; !done; {
// 		select {
// 		case <-track1:
// 			fmt.Println("winnter 1")
// 			done = true
// 		case <-track2:
// 			fmt.Println("winnter 2")
// 			done = true
// 		case <-track3:
// 			fmt.Println("winnter 3")
// 			done = true
// 		case v := <-labChan:
// 			fmt.Println("horse runing", v.Second())
// 		case <-abort:
// 			done = true
// 			break
// 		}
// 	}
// }

package main

import "fmt"

type Node struct {
	value string
	node  []Node
}

func main() {
	p := Node{value: "p"}
	g := Node{value: "g"}
	b := Node{value: "b", node: []Node{p, g}}
	q := Node{value: "q"}
	s := Node{value: "s"}
	k := Node{value: "k"}
	r := Node{value: "r", node: []Node{q}}
	a := Node{value: "a", node: []Node{r, s}}
	root := Node{value: "a", node: []Node{b, a, k}}

	outline([]string{}, &root)
}
func outline(stack []string, n *Node) {
	stack = append(stack, n.value)
	if len(n.node) == 0 {
		fmt.Println(stack)
	}

	for _, v := range n.node {
		outline(stack, &v)
	}
}
