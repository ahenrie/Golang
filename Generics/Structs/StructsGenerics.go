package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type CustomData interface {
	constraints.Ordered | []byte | []rune
}

type User[T CustomData] struct {
	ID   int
	Name string
	Data interface{}
}

func main() {
	u0 := User[int]{
		ID:   0,
		Name: "I am a User",
		Data: nil,
	}

	u1 := User[string]{
		ID:   1,
		Name: "I am also a user",
		Data: "hey",
	}

	u2 := User[byte]{
		ID:   1,
		Name: "I am also a user",
		Data: 11111111,
	}

	fmt.Printf("user: %+v\n", u0)
	fmt.Printf("user: %+v\n", u1)
	fmt.Printf("user: %+v\n", u2)
}
