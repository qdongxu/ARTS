package w52

import (
	"fmt"
	"testing"
)

func TestShortestPath(t *testing.T) {
	// graph extracted from https://www.cnblogs.com/mcomco/p/10339251.html
	g := NewGraph(8)
	g.AddEdge(0, 1, 5)
	g.AddEdge(0, 4, 9)
	g.AddEdge(0, 7, 8)
	g.AddEdge(1, 2, 12)
	g.AddEdge(1, 3, 15)
	g.AddEdge(1, 7, 4)
	g.AddEdge(2, 3, 3)
	g.AddEdge(2, 6, 11)
	g.AddEdge(3, 6, 9)
	g.AddEdge(4, 5, 4)
	g.AddEdge(4, 6, 20)
	g.AddEdge(4, 7, 5)
	g.AddEdge(5, 2, 1)
	g.AddEdge(5, 6, 13)
	g.AddEdge(7, 2, 7)
	g.AddEdge(7, 5, 6)

	vs := ShortestPath(g, 0)
	for _, path := range vs {
		fmt.Printf("w: %d,  path:  %v\n", path.Weight, path.Path)

	}
}
