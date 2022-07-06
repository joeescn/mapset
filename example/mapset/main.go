package main

import (
	"fmt"

	"github.com/joeescn/mapset"
)

func main() {
	s := mapset.NewSet[int]()
	s.Add(1)
	clone := s.Clone()
	clone.Add(2)

	fmt.Println(clone.IsSuperset(s), clone)
	// Output: true [1,2]

	m := mapset.NewMap[string, string]()
	m.Set("key", "value")
	fmt.Println(m.Get("key"))
	// Output: value true

}
