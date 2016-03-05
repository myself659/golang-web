package main

import (
	"fmt"
	"encoding/json"
	"os"
)

type model struct {
	state bool
	pictures []string
}

func main() {
	m := model{
		state: true,
		pictures: []string{
			"one.jpg",
			"two.jpg",
			"three.jpg",
		},
	}

	bs, err := json.Marshal(m)
	if err != nil {
		fmt.Println("error: ", err)
	}

	os.Stdout.Write(bs)
}

// What went wrong?
// See the next file for the answer.